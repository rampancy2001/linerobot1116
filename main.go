// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
					case "你是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好!我是自動回覆的機器人,在CoC台灣英雄聯盟為您服務,開發者為Ryan Chang")).Do()
					case "老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歡迎COC的朋友們!")).Do()
					case "豪璘是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("NG的啦")).Do()
					case "雨靖是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("X大香又甜")).Do()
					case "彥增是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("天龍國人")).Do()
					case "早安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()
					case "午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("睡醒了嗎")).Do()
					case "下午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾點了")).Do()
					case "晚安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝好夢")).Do()
					case "罵是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("黑人")).Do()
					case "幹": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話")).Do()
					case "哈": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()
					case "哈哈": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()
					case "哈哈哈": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()
					case "哈哈哈哈": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()
					case "...": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不說話就是承認")).Do()
					case "喔": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("敷衍誰")).Do()
					case "抽": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想幹嘛, 以為有圖?")).Do()
					case "老大是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你是在問那個玉樹臨風的帥哥?")).Do()
					case "抱歉": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出  部是常識吧")).Do()
					case "對不起": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出  部是常識吧")).Do()
					case "傻眼": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("雨靖的口頭禪")).Do()
					case "好呀": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一言既出 駟馬難追")).Do()
					case "感謝老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("讚嘆老大")).Do()
				}
			}
		}
	}
}
