package main


import (
	"github.com/sevenNt/rocketmq"
	//"time"
	"fmt"
	//"strconv"
	"github.com/gin-gonic/gin/json"
)

type UserVo struct {
	UID      string `json:"uid"`
	Name     string `json:"userName"`
	PassWord string `json:"passWord"`
	Oper     string `json:"oper"`
	Groups   string `json:"groups"`
}

func main() {

	group := "dev-VodHotClacSrcData"
	topic := "canal_vod_collect__video_collected_count_live"
	//conf := &rocketmq.Config{
	//	Namesrv:   "192.168.7.101:9876;192.168.7.102:9876;192.168.7.103:9876",
	//	ClientIp:     "192.168.1.23",
	//	InstanceName: "DEFAULT",
	//}
	conf := &rocketmq.Config{
		Namesrv:   "rocketmq-cs.zxbike.top:32075",
		ClientIp:"192.168.60.94",
		InstanceName: "DEFAULT",
	}

	producer, err := rocketmq.NewDefaultProducer(group, conf)
	producer.Start()
	if err != nil {
		return
	}
	//msg := rocketmq.NewMessage(topic, []byte("Hello world!"))
	//if sendResult, err := producer.Send(msg); err != nil {
	//	fmt.Println("-----------------",err)
	//	return
	//} else {
	//	fmt.Println("Sync sending success!, ", sendResult)
	//}
	var ttt  UserVo
	ttt.UID="123123"
	ttt.Name="daolin"
	ttt.PassWord="dadfasdf"
	ttt.Oper="ADD"
	ttt.Groups="1,2,3,4,5"
	data,_ :=json.Marshal(ttt)
	for i := 0; i < 3; i++ {
		msg := rocketmq.NewMessage(topic, data)
		if sendResult, err := producer.Send(msg); err != nil {
			fmt.Println("-----------------",err)
		} else {
			fmt.Println("Sync sending success!, ", sendResult)
		}
	}

}
