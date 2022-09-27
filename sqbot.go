package sqbot

import (
	"encoding/json"
	"log"
	"os"
	"time"

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
		Data struct {
			Type       string `json:"type"`
			ID         string `json:"id"`
			Attributes struct {
				ID         string      `json:"id"`
				Name       string      `json:"name"`
				Address    interface{} `json:"address"`
				IP         string      `json:"ip"`
				Port       int         `json:"port"`
				Players    int         `json:"players"`
				MaxPlayers int         `json:"maxPlayers"`
				Rank       int         `json:"rank"`
				Location   []float64   `json:"location"`
				Status     string      `json:"status"`
				Details    struct {
					Map                     string `json:"map"`
					GameMode                string `json:"gameMode"`
					Version                 string `json:"version"`
					Secure                  int    `json:"secure"`
					LicensedServer          bool   `json:"licensedServer"`
					LicenseID               string `json:"licenseId"`
					NumPubConn              int    `json:"numPubConn"`
					NumPrivConn             int    `json:"numPrivConn"`
					NumOpenPrivConn         int    `json:"numOpenPrivConn"`
					SquadPlayerReserveCount int    `json:"squad_playerReserveCount"`
					SquadPlayTime           int    `json:"squad_playTime"`
					SquadPublicQueueLimit   int    `json:"squad_publicQueueLimit"`
					SquadPublicQueue        int    `json:"squad_publicQueue"`
					SquadReservedQueue      int    `json:"squad_reservedQueue"`
					SquadTeamOne            string `json:"squad_teamOne"`
					SquadTeamTwo            string `json:"squad_teamTwo"`
					Modded                  bool   `json:"modded"`
					ServerSteamID           string `json:"serverSteamId"`
				} `json:"details"`
				Private     bool      `json:"private"`
				CreatedAt   time.Time `json:"createdAt"`
				UpdatedAt   time.Time `json:"updatedAt"`
				PortQuery   int       `json:"portQuery"`
				Country     string    `json:"country"`
				QueryStatus string    `json:"queryStatus"`
			} `json:"attributes"`
			Relationships struct {
				Game struct {
					Data struct {
						Type string `json:"type"`
						ID   string `json:"id"`
					} `json:"data"`
				} `json:"game"`
			} `json:"relationships"`
		} `json:"data"`
		Included []interface{} `json:"included"`
	}
	jsonPath := "C:\\info.json"

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Printf("open json file %v error [ %v ]", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	var conf Info
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("decode error [ %v ]", err)
		return
	} else {
		ctx.Send(message.At(ctx.Event.UserID).String() + message.Text("\n", "数据更新时间：(美国时间，中国加8小时）", conf.Data.Attributes.UpdatedAt, "\n", "服务器名称：", conf.Data.Attributes.Name, "\n", "当前人数：", conf.Data.Attributes.Players, "/", conf.Data.Attributes.MaxPlayers, "\n", "排队人数(不是实时)：", conf.Data.Attributes.Details.SquadPublicQueue, "\n", "服务器模式：", conf.Data.Attributes.Details.GameMode, "，", conf.Data.Attributes.Details.SquadTeamOne, "   VS   ", conf.Data.Attributes.Details.SquadTeamTwo, "\n", "服务器版本：", conf.Data.Attributes.Details.Version, "\n", "服务器地图：", conf.Data.Attributes.Details.Map).String() + sendMsg.String())
	}
}
