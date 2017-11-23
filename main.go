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
//	"github.com/PuerkitoBio/goquery"
)

var bot *linebot.Client

func main() {
	var err error
//	var err2 error
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
					case "你好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嘿你好")).Do()
					case "早安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()
					case "午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("睡醒了嗎")).Do()
					case "下午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾點了,昨晚在幹嘛")).Do()
					case "晚安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝好夢")).Do()
					case "hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "Hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "嗨": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "看": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("別以為不知道你在說髒話")).Do()
					case "喔": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("敷衍誰")).Do()
					case "哦": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("敷衍誰")).Do()
					case "幹": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()
					case "抽": bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()
					case "老大是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你是在問那個玉樹臨風的帥哥?")).Do()
					case "抱歉": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出口部是常識吧")).Do()
					case "對不起": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出口部是常識吧")).Do()
					case "傻眼": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("雨靖的口頭禪")).Do()
					case "好呀": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一言既出 駟馬難追")).Do()
					case "感謝老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("讚嘆老大")).Do()
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
					case "機器人": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()
					case "/help": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()
					case "#1": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這你也相信喔..幾歲了")).Do()
					case "#2": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("系統建置中~~預計2050年上線")).Do()
					case "#9": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("動動腰起床了, 生意上門喔!")).Do()
				}
				if strings.Contains(message.Text,"生日快樂") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家一起唱~祝你生日快樂~祝你生日快樂~祝你生日快樂~祝你生日快樂")).Do()} 
				if strings.Contains(message.Text,"哈哈哈哈哈哈") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()}
				if strings.Contains(message.Text,"......") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很無言")).Do()}
				if strings.Contains(message.Text,"幹你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"幹他") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"幹我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"吉利山莊") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有正妹")).Do()}
				if strings.Contains(message.Text,"死我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家要身體健康")).Do()}
				if strings.Contains(message.Text,"幫忙捐兵") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又來了,請大家幫忙捐兵喔")).Do()}
				if strings.Contains(message.Text,"請問") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我想這是個好問題!請各位幫忙回答喔")).Do()}
				if strings.Contains(message.Text,"換手機") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哇～好野人")).Do()}
				if strings.Contains(message.Text,"excel") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"Excel") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"EXCEL") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"新年") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝福大家新年快樂！")).Do()}
				if strings.Contains(message.Text,"聖誕節") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("叮叮噹叮叮噹鈴聲多響亮～聖誕節快樂！")).Do()}
				if strings.Contains(message.Text,"祝福") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我用誠摯的心給與最大的祝福")).Do()}
				if strings.Contains(message.Text,"死了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家要身體健康")).Do()}
				if strings.Contains(message.Text,"怎麼辦") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("擲茭喔")).Do()}
				if strings.Contains(message.Text,"羨慕") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()}
				if strings.Contains(message.Text,"PS4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"Ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"野生") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歐買尬!本人比較帥")).Do()}
				if strings.Contains(message.Text,"媽的") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽知道你在這裡po廢文嗎")).Do()}
				if strings.Contains(message.Text,"ㄇㄉ") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽知道你在這裡po廢文嗎")).Do()}
				if strings.Contains(message.Text,"吃屎") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又油又香 粉好吃")).Do()}
				if strings.Contains(message.Text,"妹妹") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("醒醒吧~你沒有妹妹")).Do()}
				if strings.Contains(message.Text,"謝謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"感謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"累了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("對不起我只是個機器人, 可以原諒我嗎")).Do()}
				if strings.Contains(message.Text,"星期一") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("每個星期一從醒來那一刻便覺得心裡有一種難以擺脫的壓抑，覺得很鬱悶。")).Do()}
				if strings.Contains(message.Text,"星期五") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("T.G.I.F.喔喔喔喔喔喔耶")).Do()}
				if strings.Contains(message.Text,"關鍵字") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("想知道更多服務查詢請回傳 /help")).Do()}
				if strings.Contains(message.Text,"地震") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("媽呀, 嚇死人, 以下開始報平安!")).Do()}
				if strings.Contains(message.Text,"http") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("謝謝分享!愛你喔")).Do()}
				if strings.Contains(message.Text,"八兩金") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("新莊八兩金在此")).Do()}
				if strings.Contains(message.Text,"我朋友") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("先承認你就是你朋友吧")).Do()}
				if strings.Contains(message.Text,"我可以") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你有不可以的嗎")).Do()}
				if strings.Contains(message.Text,"點名") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大 1")).Do()}
				if strings.Contains(message.Text,"1,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1",strings.Trim(message.Text,"1,"))).Do()}
				if strings.Contains(message.Text,"2,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2",strings.Trim(message.Text,"2,"))).Do()}
				if strings.Contains(message.Text,"3,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("3",strings.Trim(message.Text,"3,"))).Do()}
				if strings.Contains(message.Text,"4,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("4",strings.Trim(message.Text,"4,"))).Do()}
				if strings.Contains(message.Text,"測試") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2","18")).Do()}
//				if strings.Contains(message.Text,"joke") {
//					doc, err2 := goquery.NewDocument("http://www.qiushibaike.com")
//  					if err2 != nil {log.Fatal(err2)}
//  					doc.Find(".content").Each(func(i int, s *goquery.Selection){
//    					bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage(s)).Do()
// 					 })
//				}
				//if strings.Contains(message.Text,"測試2") {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg","http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg")).Do()}
				if strings.Contains(message.Text,"老大") {
					if strings.Contains(message.Text,"靠你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大有寬闊肩膀可以依偎")).Do()}
					if strings.Contains(message.Text,"問你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("來來來！都來")).Do()}
					if strings.Contains(message.Text,"問題") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("來來來！都來")).Do()}
					if strings.Contains(message.Text,"最愛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一生只愛你一個")).Do()}
					if strings.Contains(message.Text,"不行") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你才不行")).Do()}
					if strings.Contains(message.Text,"老了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大今年18歲")).Do()}
					if strings.Contains(message.Text,"老人") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大今年18歲")).Do()}
					if strings.Contains(message.Text,"愛你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請排隊領取號碼牌")).Do()}
					if strings.Contains(message.Text,"你好") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("お元氣てすか？私は元氣てす")).Do()}
					if strings.Contains(message.Text,"帥") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("帥不過是一個字 卻跟著我一輩子")).Do()}
					if strings.Contains(message.Text,"在嗎") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好 請問有什麼事")).Do()}
					if strings.Contains(message.Text,"在哪") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在這裡在那裡 在昨天在未來 原來已在你心裡")).Do()}
					if strings.Contains(message.Text,"去哪") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("去尋找未來 追求夢想")).Do()}
					if strings.Contains(message.Text,"幹嘛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在你後面喜滋滋地看著你")).Do()}
					if strings.Contains(message.Text,"幾歲") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("18歲")).Do()}
					if strings.Contains(message.Text,"厲害") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("鵝樂告無大計")).Do()}
					if strings.Contains(message.Text,"單身") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("報名專線: 控八控控-控啾里控控控")).Do()}
					if strings.Contains(message.Text,"閒") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在忙也要跟你喝一杯咖啡")).Do()}
				}
				if strings.Contains(message.Text,"機器人") {
					if strings.Contains(message.Text,"沒") {
						if strings.Contains(message.Text,"用") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
						if strings.Contains(message.Text,"幫助") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哼, 罷工中")).Do()}
					}
					if strings.Contains(message.Text,"廢物") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"爛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"功能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"幹嘛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"會什麼") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"技能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"服務") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
				}
			}
		}
	}
}
