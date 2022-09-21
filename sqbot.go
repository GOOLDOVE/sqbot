package sqbot

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/RicheyJang/PaimengBot/basic/nickname"
	"github.com/RicheyJang/PaimengBot/manager"
	"github.com/RicheyJang/PaimengBot/utils"

	log "github.com/sirupsen/logrus"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)
func sqbot(ctx *zero.Ctx) {
	// 打开json文件
	jsonFile, err := os.Open("info.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	// 发送
	ctx.Send(message.Text("服务器信息为：\n").String() + message.At(ctx.Event.UserID).String() + sendMsg.String())
}

//var proxy *manager.PluginProxy
//var info = manager.PluginInfo{
//	Name: "服务器状态查询",
//	Usage: `
//查询服务器信息
//	服务器信息 主服
//`,
//	Classify: "实用工具",
//}

//技术优先以下封存
//func init() {
//	proxy = manager.RegisterPlugin(info)
//	if proxy == nil {
//		return
//	}
//	proxy.OnCommands([]string{"服务器", "服务器查询"}, zero.OnlyToMe).SetBlock(true).ThirdPriority().Handle(covid19Handler)
//	proxy.OnRegex(`^(\S{1,10})服务器信息$`, zero.OnlyToMe).SetBlock(true).SetPriority(3).Handle(covid19Handler)
//}
