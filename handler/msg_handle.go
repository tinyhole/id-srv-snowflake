package handler

import (
	"github.com/tinyhole/id-srv-snowflake/idl/base"
	"github.com/tinyhole/id-srv-snowflake/idl/platform/id/srv/snowflake"
	//"github.com/micro/go-micro/config"
	"context"
	"time"
)


type SnowFlake struct{
	key           string  // etcd key
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
}

func (sf *SnowFlake)Init() {
	sf.lastTimeStamp = time.Now().UnixNano() / 1000000
	// 把机器 id 左移 12 位,让出 12 位空间给序列号使用
	//sf.machineID = int64(config.Get("srv","srvId").Int(1024))
	sf.machineID = 1024
	sf.getLastTimeStamp()
}

//获得etcd时间
func (sf *SnowFlake)getLastTimeStamp(){
	sf.lastTimeStamp = time.Now().UnixNano()

}


//更新etcd 时间记录
func (sf *SnowFlake)setLastTimeStamp(){
}


func (sf *SnowFlake)GetSnowflakeId() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000

	// 同一毫秒
	if curTimeStamp == sf.lastTimeStamp {
		sf.sn++
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sf.sn > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			sf.lastTimeStamp = curTimeStamp
			sf.sn = 0
		}

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作

		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位

		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF

		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22

		id := rightBinValue | sf.machineID | sf.sn
		return id
	}

	if curTimeStamp > sf.lastTimeStamp {
		sf.sn = 0
		sf.lastTimeStamp = curTimeStamp
		sf.setLastTimeStamp()

		// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作

		// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位

		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF

		// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
		rightBinValue <<= 22

		id := rightBinValue | sf.machineID | sf.sn

		return id

	}

	if curTimeStamp < sf.lastTimeStamp {
		return 0
	}

	return 0
}

var globalSnow SnowFlake
func init(){
	globalSnow := SnowFlake{}
	globalSnow.Init()
}

type Handle struct{}

func (h *Handle) GetID(ctx context.Context, req *base.Empty, rsp *snowflake.GetIDRsp)error {

	rsp.ID = globalSnow.GetSnowflakeId()

	return nil
}
