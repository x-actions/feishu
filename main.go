// Copyright 2022 xiexianbin<me@xiexianbin.cn>
//
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

// ref https://github.com/larksuite/oapi-sdk-go/blob/v2_main/sample/customer_bot/customer_bot.go

package main

import (
	"context"
	"flag"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v2"
	"os"
)

var webhook string
var secret string

var msgtype string
var title string
var feishuContent string

func init() {
	flag.StringVar(&webhook, "webhook", "", "feishu CustomerBot webhook url, like https://open.feishu.cn/open-apis/bot/v2/hook/<uuid>")
	flag.StringVar(&secret, "secret", "", "feishu CustomerBot webhook secret")

	flag.StringVar(&msgtype, "msgtype", "text", "message type (text, interactive)")
	flag.StringVar(&title, "title", "foo title", "message title send by bot")
	flag.StringVar(&feishuContent, "content", "bar message content", "message content send by bot, "+
		"if msgtype is interactive, content is markdown Text "+
		"ref https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#4996824a")

	flag.Parse()
}

func usage() {
	flag.Usage()
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	var err error
	var resp *lark.CustomerBotSendMessageResp

	customerBot := lark.NewCustomerBot(webhook, secret)

	var sendFeishuContent interface{}
	switch msgtype {
	case "text":
		sendFeishuContent = &lark.MessageText{Text: feishuContent}
	case "interactive":
		sendFeishuContent = &lark.MessageCard{
			Header: &lark.MessageCardHeader{
				Title: &lark.MessageCardPlainText{
					Content: title,
				},
			},
			Config: &lark.MessageCardConfig{WideScreenMode: lark.BoolPtr(true)},
			Elements: []lark.MessageCardElement{
				&lark.MessageCardMarkdown{
					Content: feishuContent,
				},
			},
		}

	default:
		panic(fmt.Sprintf("un-support msgtype: %s", msgtype))
	}

	resp, err = customerBot.SendMessage(context.TODO(), msgtype, sendFeishuContent)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("request id: %s\n", resp.RequestId())
		fmt.Printf("response is: %s\n", lark.Prettify(resp))
	}
}
