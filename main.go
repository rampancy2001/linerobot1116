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
	"strings"
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
					case "你是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好!我是自動回覆的機器人,在CoC台灣英雄聯盟為您服務")).Do()
					case "rampancy": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好!我是自動回覆的機器人,在CoC台灣英雄聯盟為您服務")).Do()
					case "老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "豪璘是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("NG的啦")).Do()
					case "雨靖是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("網帥")).Do()
					case "彥增是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("天龍國人")).Do()
					case "罵是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("塔哥被黑了")).Do()
					case "瞬哥是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("澳洲打工仔")).Do()
					case "Jonas是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Jonas")).Do()
					case "大家好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好呀")).Do()
					case "早安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()
					case "午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("睡醒了嗎")).Do()
					case "下午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾點了,昨晚在幹嘛")).Do()
					case "晚安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝好夢")).Do()
					case "hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "看": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()
					case "喔": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("敷衍誰")).Do()
					case "哦": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("敷衍誰")).Do()
					case "抽": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想幹嘛, 以為有圖?")).Do()
					case "老大是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你是在問那個玉樹臨風的帥哥?")).Do()
					case "抱歉": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出口部是常識吧")).Do()
					case "對不起": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出口部是常識吧")).Do()
					case "傻眼": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("雨靖的口頭禪")).Do()
					case "好呀": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一言既出 駟馬難追")).Do()
					case "感謝老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("讚嘆老大")).Do()
					case "感謝": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真有禮貌!")).Do()
					case "謝謝": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真有禮貌!")).Do()
					case "誰在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "有誰在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "真可愛": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("啾咪")).Do()
					case "去洗澡": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("洗澡卡get!")).Do()
					case "@@": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("^_^")).Do()
					case "來下班": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很爽耶你")).Do()
					case "下班了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很爽耶你")).Do()
					case "上班去": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油加油")).Do()
					case "上班了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("上班還不是都在混")).Do()
					case "太強了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("跪")).Do()
					case "笑點在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽知道你在這裡po廢文嗎")).Do()
				}
				if strings.Contains(message.Text,"生日快樂") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家一起唱~祝你生日快樂~祝你生日快樂~祝你生日快樂~祝你生日快樂")).Do()} 
				if strings.Contains(message.Text,"哈哈哈") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()}
				if strings.Contains(message.Text,"...") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很無言")).Do()}
				if strings.Contains(message.Text,"幹") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"吉利") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有正妹")).Do()}
				if strings.Contains(message.Text,"死我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("RIP")).Do()}
				if strings.Contains(message.Text,"幫忙捐兵") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又來了,請大家幫忙捐兵喔")).Do()}
				if strings.Contains(message.Text,"請問") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請各位幫忙回答喔")).Do()}
				if strings.Contains(message.Text,"測試") {bot.ReplyMessage(event.ReplyToken, linebot.NewImagemapMessage("http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg","http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg",linebot.ImagemapBaseSize{200, 200})).Do()}
			}
		}
	}
}
