package sqbot

import (
	"encoding/json"
	"log"
	"os"

	"github.com/RicheyJang/PaimengBot/manager"

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
	proxy.OnCommands([]string{"sqbot", "服务器多少人", "现在服务器多少人了"}).SetBlock(true).SecondPriority().Handle(sqbot) // [4] 注册事件处理函数
	proxy.AddConfig("times", 2)                                                                              // proxy提供的统一配置项管理功能，此函数新增一个配置项times，默认值为2
}

var sendMsg message.Message

// EchoHandler [5] Handler实现
func sqbot(ctx *zero.Ctx) {
	type Info struct {
		All      int64  `json:"max_players"`
		NOW      int64  `json:"currentPlayers"`
		VISION   string `json:"version"`
		MAP      string `json:"map"`
		T1       string `json:"squad_teamOne"`
		T2       string `json:"squad_teamTwo"`
		TIME     string `json:"updatedAt"`
		GAMEMODE string `json:"gameMode"`
		NAME     string `json:"name"`
		//time   string `json:"updatedAt"`
		//time   string `json:"updatedAt"`
	}
	jsonPath := "C:\\info.json"

	fmt, err := os.Open("C:\\RCON\\app.js")
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Printf("open json file %v error [ %v ]", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	_ = fmt

	var conf Info
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("decode error [ %v ]", err)
		return
	} else {
		ctx.Send(message.At(ctx.Event.UserID).String() + message.Text("服务器信息为：\n", "服务器名称：", conf.NAME, "\n", "当前人数：", conf.NOW, "/", conf.All, "\n", "服务器模式：", conf.GAMEMODE, "，", conf.T1, "vs", conf.T2, "\n", "服务器版本：", conf.VISION, "\n", "服务器地图：", conf.MAP).String() + sendMsg.String())
	}
}
