package sqbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/RicheyJang/PaimengBot/manager"
	"github.com/RicheyJang/PaimengBot/utils"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var info = manager.PluginInfo{ // [1] 声明插件信息结构变量
	Name: "机器人",
	Usage: `
用法：
	机器人
`,
}
var proxy *manager.PluginProxy // [2] 声明插件代理变量

func init() {
	proxy = manager.RegisterPlugin(info) // [3] 使用插件信息初始化插件代理
	if proxy == nil {                    // 若初始化失败，请return，失败原因会在日志中打印
		return
	}
	proxy.OnCommands([]string{"机器人", "sqbot"}).SetBlock(true).SecondPriority().Handle(sqbot) // [4] 注册事件处理函数
	proxy.AddConfig("times", 2)                                                              // proxy提供的统一配置项管理功能，此函数新增一个配置项times，默认值为2
}

// EchoHandler [5] Handler实现
func sqbot(ctx *zero.Ctx) {
	// 打开json文件
	jsonFile, err := os.Open("info.json")

	str := utils.GetArgs(ctx)           // 派蒙Bot提供的工具函数，用于获取此次事件的消息参数内容
	tm := proxy.GetConfigInt64("times") // proxy提供的统一配置项管理功能，此函数用于获取int64类型的times配置项值

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
	for i := int64(0); i < tm; i++ {
		ctx.Send(message.Text("服务器信息为：\n").String() + message.At(ctx.Event.UserID).String() + sendMsg.String())
	}
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
