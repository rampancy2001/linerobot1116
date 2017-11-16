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
				if (message.Text == "老大") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(" 歡迎COC的朋友們!")).Do()}
				if (message.Text == "豪璘是誰") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("NG的啦")).Do()}				
				if (message.Text == "雨靖是誰") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("X大香又甜")).Do()}	
				if (message.Text == "彥增是誰") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("天龍國人")).Do()}
				if (message.Text == "早安") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()}
				if (message.Text == "罵是誰") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("黑人")).Do()}
				if (message.Text == "幹") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話")).Do()}	
				if (message.Text == "哈") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()}	
				if (message.Text == "...") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不說話就是承認")).Do()}	
			}
		}
	}
}
