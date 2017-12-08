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
	"math/rand"
	"time"
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
	var a int
	rand.Seed(time.Now().UnixNano())
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
					case "老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "謝老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大姓張不姓謝喔")).Do()
					case "豪璘是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("NG的啦")).Do()
					case "雨靖是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("網帥")).Do()
					case "彥增是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("天龍國人")).Do()
					case "罵是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("塔哥被黑了")).Do()
					case "瞬哥是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("澳洲打工仔")).Do()
					case "Jonas是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("他高傲，但是宅心仁厚，他低調，但是受萬人景仰，他可以把神賜給人類的火，運用得出神入化，燒出堪稱火之藝術的超級菜式，他究竟是神仙的化身？還是地獄的使者？沒人知道，但是可以肯定，每個人都給他一個稱號——食~~神！——")).Do()
					case "大家好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好呀")).Do()
					case "大家早": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你也早!")).Do()
					case "你好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嘿你好")).Do()
					case "早安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()
					case "午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("睡醒了嗎?太晚了吧")).Do()
					case "下午安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾點了,老實說昨晚在幹嘛")).Do()
					case "晚安": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝好夢")).Do()
					case "hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "Hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "嗨": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "看": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("別以為不知道你在說髒話")).Do()
					case "喔": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("喔屁喔")).Do()
					case "哦": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哦屁喔")).Do()
					case "幹": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()					
					case "老大是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你是在問那個玉樹臨風的帥哥?")).Do()
					case "抱歉": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出X部是常識吧")).Do()
					case "對不起": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出X部是常識吧")).Do()
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
					case "好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("謝謝")).Do()
					case "不好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("為什麼有一點點難過")).Do()
					case "/help": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()
					case "#1": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這你也相信喔..幾歲了")).Do()
					case "#2": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("系統建置中~~預計2050年上線")).Do()
					case "#9": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("動動腰起床了, 生意上門喔!")).Do()
					case "抽": a=rand.Intn(4)
						if a == 0 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()}
						if a == 1 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg","https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg")).Do()}
						if a == 2 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()}
						if a == 3 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg","https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg")).Do()}
				}
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
					if strings.Contains(message.Text,"單身") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("曾经有一份真诚的爱情放在我面前，我没有珍惜，等我失去的时候我才后悔莫及，人世间最痛苦的事莫过于此。如果上天能够给我一个再来一次的机会，我会对那个女孩子说三个字：我爱你。如果非要在这份爱上加上一个期限，我希望是……一万年！")).Do()}
					if strings.Contains(message.Text,"閒") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在忙也要跟你喝一杯咖啡")).Do()}
					if strings.Contains(message.Text,"快") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不要一直催老大, 一直催一直催")).Do()}
					if strings.Contains(message.Text,"公分") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("相信我, 老大30公分")).Do()}
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
				if strings.Contains(message.Text,"跨年") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都沒揪")).Do()}
				if strings.Contains(message.Text,"聖誕節") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("叮叮噹叮叮噹鈴聲多響亮～祝大家聖誕節快樂！")).Do()}
				if strings.Contains(message.Text,"感恩節") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝上帝賜予一年度豐收的祝福")).Do()}
				if strings.Contains(message.Text,"祝福") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我用誠摯的心給與最大的祝福")).Do()}
				if strings.Contains(message.Text,"死了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家要身體健康")).Do()}
				if strings.Contains(message.Text,"怎麼辦") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("擲茭喔")).Do()}
				if strings.Contains(message.Text,"羨慕") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()}
				if strings.Contains(message.Text,"PS4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"Ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"野生") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歐買尬!本人比較帥")).Do()}
				if strings.Contains(message.Text,"媽的") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("人是人他媽生的，妖是妖他媽生的，只要你有一顆善良的心就不再是妖，是人妖？")).Do()}
				if strings.Contains(message.Text,"ㄇㄉ") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽貴姓？")).Do()}
				if strings.Contains(message.Text,"吃屎") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又油又香 粉好吃")).Do()}
				if strings.Contains(message.Text,"妹妹") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("醒醒吧~你沒有妹妹")).Do()}
				if strings.Contains(message.Text,"謝謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"感謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"累了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("對不起我只是個機器人, 可以原諒我嗎")).Do()}
				if strings.Contains(message.Text,"星期一") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("每個星期一從醒來那一刻便覺得心裡有一種難以擺脫的壓抑，覺得很鬱悶。")).Do()}
				if strings.Contains(message.Text,"星期五") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("T.G.I.F.喔喔喔喔喔喔耶")).Do()}
				if strings.Contains(message.Text,"解鎖") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰解鎖了?給你拍拍手的啦")).Do()}
				if strings.Contains(message.Text,"關鍵字") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油喔, 還有很多慢慢找")).Do()}
				if strings.Contains(message.Text,"地震") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("媽呀, 嚇死人, 以下開始報平安!")).Do()}
				if strings.Contains(message.Text,"http") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("謝謝分享!愛你喔")).Do()}
				if strings.Contains(message.Text,"八兩金") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("新莊八兩金在此")).Do()}
				if strings.Contains(message.Text,"我朋友") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("先承認你就是你朋友吧")).Do()}
				if strings.Contains(message.Text,"我可以") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你有不可以的嗎")).Do()}
				if strings.Contains(message.Text,"點名") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大 1")).Do()}
				if strings.Contains(message.Text,"打氣") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油!加油!加油加油加油!")).Do()}
				if strings.Contains(message.Text,"3D") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我家樓上正妹32D")).Do()}
				if strings.Contains(message.Text,"下面") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大下的麵很好吃")).Do()}
				if strings.Contains(message.Text,"扶") {
					if strings.Contains(message.Text,"過馬路") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("積積陰陰德")).Do()}
					}
				if strings.Contains(message.Text,"嗎") {
					if strings.Contains(message.Text,"會") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("既然你誠心誠意的發問了,我們就大發慈悲的告訴你,哇恩災")).Do()}
					if strings.Contains(message.Text,"是") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("既然你誠心誠意的發問了,我們就大發慈悲的告訴你,哇恩災")).Do()}
					if strings.Contains(message.Text,"可以") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("既然你誠心誠意的發問了,我們就大發慈悲的告訴你,哇恩災")).Do()}
					}
				if strings.Contains(message.Text,"中") {
					if strings.Contains(message.Text,"出來") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("簡稱中出來")).Do()}
					}
				if strings.Contains(message.Text,"好友") {
					if strings.Contains(message.Text,"加") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("也要加我!!好嗎")).Do()}
					}
				if strings.Contains(message.Text,"三星") {
					if strings.Contains(message.Text,"11本") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哇~太強啦,偶像")).Do()}
					}
				if strings.Contains(message.Text,"你") {
					if strings.Contains(message.Text,"慢") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("剛睡醒呀")).Do()}
					}
				if strings.Contains(message.Text,"期待") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我割你的臍帶")).Do()}
				if strings.Contains(message.Text,"更新") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("太好了, 快點來下載")).Do()}
				if strings.Contains(message.Text,"贏了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("喔耶, 爽呀")).Do()}
				if strings.Contains(message.Text,"淫了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你溼了嗎")).Do()}
				if strings.Contains(message.Text,"濕了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("濕主你好")).Do()}
				if strings.Contains(message.Text,"溼了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("溼主你好")).Do()}
				if strings.Contains(message.Text,"30cm") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("皮就佔29公分")).Do()}
				if strings.Contains(message.Text,"30公分") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("皮就佔29公分")).Do()}
				if strings.Contains(message.Text,"ininder") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("沒30cm就請低調點")).Do()}
				if strings.Contains(message.Text,"XD") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OGC")).Do()}
				if strings.Contains(message.Text,"愛你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("羞 >///<")).Do()}
				if strings.Contains(message.Text,"也知道") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("根本有預知的神力呀, 佩服")).Do()}
				if strings.Contains(message.Text,"超帥") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("別以為你長得帥我就不打你")).Do()}
				if strings.Contains(message.Text,"爽") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我也要一起爽")).Do()}
				if strings.Contains(message.Text,"打我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("打你應該，不打你悲哀!")).Do()}
				if strings.Contains(message.Text,"賠") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("聊得挺投緣，賠點醫藥費算了。")).Do()}
				if strings.Contains(message.Text,"氣勢") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這種陣勢難道就是人稱《神鵰俠侶》里的楊過、小龍女？")).Do()}
				if strings.Contains(message.Text,"沒說") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("難道我學過如來神掌也要說給你聽嗎？")).Do()}
				if strings.Contains(message.Text,"沒力") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("打人都沒力氣，還說是黑社會？")).Do()}
				if strings.Contains(message.Text,"好詩") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一鄉二里共三夫子不識四書五經六易竟教七八九子十分大膽，十室九貧尚且湊得八兩七錢六分五厘四毫且三心二意一等下流。")).Do()}
				if strings.Contains(message.Text,"調皮") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你又頑皮了，怎麼能這樣跟觀音姐姐說話呢？")).Do()}
				if strings.Contains(message.Text,"沒用") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("就算是一條內褲一卷衛生紙也有它的用處。")).Do()}
				if strings.Contains(message.Text,"有空") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我一秒鐘幾十萬上下，我會和你去踢球？")).Do()}
				if strings.Contains(message.Text,"掃地") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("掃地只不過是我的表面工作，我真正地身份是一位研究僧。")).Do()}
				if strings.Contains(message.Text,"夢遺") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("貧僧乃少林寺方丈，法號夢遺。阿彌陀佛，我隨風而來，隨風而去...")).Do()}
				if strings.Contains(message.Text,"好吃") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("從來沒有試過這麼清新脫俗的感覺，牛肉的鮮，撒尿蝦的甜，混在一起的味道，竟然比「老鼠般」有過之而不及，正如比我的初戀更加詩情畫意，所謂舉頭望明月，低頭思故鄉，好詩！好詩啊！")).Do()}
				if strings.Contains(message.Text,"騙") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("憑你的智慧，我唬得了你嗎？")).Do()}
				if strings.Contains(message.Text,"興趣") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("除暴安良是我們做市民的責任，行善積德是我本身的興趣，所以扶老太太過馬路我每個星期都做一次，如果是碰到國定假日的話我還做兩三次。")).Do()}
				if strings.Contains(message.Text,"責任") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("除暴安良是我們做市民的責任，行善積德是我本身的興趣，所以扶老太太過馬路我每個星期都做一次，如果是碰到國定假日的話我還做兩三次。")).Do()}
				if strings.Contains(message.Text,"普通") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("錯！這並不是個普通的箱子，它是箱中之神，簡稱箱神！")).Do()}
				if strings.Contains(message.Text,"服了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("I服了You")).Do()}
				if strings.Contains(message.Text,"危險") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你回火星吧，地球是很危險的")).Do()}
				if strings.Contains(message.Text,"尊敬") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我對你的敬仰真是如滔滔江水，連綿不絕，又有如黃河泛濫一發不可收拾。")).Do()}
				if strings.Contains(message.Text,"笑話") {
					a=rand.Intn(8)
					if a == 0 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有一個國家舉辦最噁心比賽，至最後決賽時剩下三個人爭冠，其中一個人拿了一杯尿，在嘴中漱了漱，吞下，哈的一聲，全場鼓掌想冠軍必落於此家。第二個人從袋中拿出一堆蟑螂，剝了翅膀就嘖嘖嘖的吸牠的肚子，不時還吐出一兩隻腳,吃到第十隻的時後,國王面有菜色的說不用吃了你這樣就第一名了。 此時見第三個人拿出一杯液體，說，這是我半年前感冒到現在，每次吐的痰都收集在裡面，我現在要把它喝完。只見一整杯白白黃黃還帶泡泡的，他搖了搖，試圖讓有些積太久快要凝固的化開， 國王眼淚都要掉下來了，說：不用了不用了你只要喝一口你就冠軍了～這人便拿起杯子咕嘟咕嘟地開始喝，因為很濃又很多過了五分多鐘才喝完，此時全場已淚流滿面，國王說幹嘛我不是叫你喝一口就冠軍了嗎？這人回答道，我也只是想喝一口，但是我一直咬不斷～～")).Do()}
					if a == 1 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有人很喜歡“麻辣粉絲煲”這道菜。有一次，他上飯館，又點了這道菜。但侍者告訴他，這道菜已經賣完了。“真的賣完了嗎？”他很失望地問。“先生，真的賣完了。你瞧，最後一份賣給那桌的先生了。”侍者回答道。那人順著侍者的指點，看見有個很體面的紳士坐在鄰座。紳士的飯菜已經吃得差不多了，但那份“麻辣粉絲煲”居然還是滿滿的。那人覺得紳士很浪費美味，所以他走到紳士旁邊，指著那份“麻辣粉絲煲”，很有禮貌地問：“先生，您這還要嗎？”紳士很有風度地搖搖頭。于是那人立刻坐下，拿起調羹狼吞虎咽起來。風卷殘雲，一會兒一半下肚了，突然間他發現在砂鍋底躺著一只很小很小但皮毛已長全的小老鼠。一陣惡心，那人把吃下去的所有粉絲通通吐回了砂鍋裏。當他在那兒翻胃不已的時候，那紳士用很同情的眼光看著他，說：“很惡心是嗎？剛才我也是這樣……”")).Do()}
					if a == 2 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這天，酒店老板正在大廳巡視。來了一乞丐上前說道：”老板給個牙簽行嗎？”老板給他一個打發走了。一會兒，又來一個乞丐，也是來要牙簽的。老板心想現在這乞丐怎麽不要飯改要牙簽了？也同樣給他一個打發走了，沒過多久，又來一個乞丐。老板對他說：”你也是來要牙簽的嗎？”乞丐說：”有個人吐了，可我晚了一步，已經被前面兩個乞丐把能吃的都吃了，現在只剩下湯了。你能給我個吸管嗎？”")).Do()}
					if a == 3 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大、老二乘坐飛機，老二暈機，不停嘔吐; 一袋吐滿;老大只好去取袋子，等他回來時，發覺全機人都在不停嘔吐。 老大問其原因，老二說：“我看到這只袋子也吐滿了，只好又喝進去了半袋，結果他們就全吐了。”")).Do()}
					if a == 4 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有個人去帶著朋友去探望他的外婆。當他和外婆說話時，他的朋友開始吃著咖啡桌上放的花生，把花生都吃完了。當他們離開時，他的朋友對外婆說：「謝謝您的花生」外婆回應說：「喔！嗯！唉！自從我牙齒掉光後，我就只能吸掉它們外層的巧克力而已。老了，咳..」")).Do()}
					if a == 5 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有個富豪找傭人,面試的題目是上廁所,前幾個上完後都沒有洗手就出來了,富豪因此把他們打發走了,只有一個洗了手,于是富豪留下了他;可是有一天,富豪卻發現他沒有洗手就出來了,富豪問他是爲什麽;傭人答到“偶今天帶了手紙”")).Do()}
					if a == 6 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("醫學院學生圍在蓋著白布的屍體周圍第一次上真人解剖課。教授開始講課了，“做爲醫生，必需具備兩項重要素質，第一要不怕惡心”。說完教授掀開白布，把手指插入屍體肛門，然後抽出並放在嘴裏吮吸。“學著做”，他告訴同學們。同學們都覺得很惡心，猶豫很久但最終不得不依次去做。當最後一個人做完後，教授又說 “第二個素質是觀察。我插入中指但吸食指。同學們，要注意觀察！”")).Do()}
					if a == 7 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你呀")).Do()}
				}
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

				if strings.Contains(message.Text,"機器人") {
					if strings.Contains(message.Text,"沒") {
						if strings.Contains(message.Text,"用") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
						if strings.Contains(message.Text,"幫助") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哼, 罷工中")).Do()}
					}
					if strings.Contains(message.Text,"廢物") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"爛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"笨") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"臭") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"功能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"幹嘛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"會什麼") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"技能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"服務") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"守則") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("第一:機器人不得傷害人類，或因不作為(袖手旁觀)使人類受到傷害；第二:除非違背第一法則，機器人必須服從人類的命令；第三:在不違背第一及第二法則下，機器人必須保護自己。")).Do()}
				}
			}
		}
	}
}
