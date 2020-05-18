package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type MicroCMSWebhookRequestBody struct {
	Service string `json:"service"`
	Api     string `json:"api"`
	Id      string `json:"id"`
	Type    string `json:"type"`
}

func webhookHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストをバリデートする
	if req.Method != "POST" {
		log.Println("error: invalid request method")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Header.Get("Content-Type") != "application/json" {
		log.Println("error: Invalid request content type")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	microCMSWebhookRequestBody := MicroCMSWebhookRequestBody{}
	if err := json.Unmarshal(body, &microCMSWebhookRequestBody); err != nil {
		log.Fatal(err)
	}

	// TODO: MicroCMSのAPIを叩き最新記事情報を取得する
	// TODO: 短時間に連投された記事を見分けられたりすると尚良い

	// TODO: responseが5分返ってこないのはよくないので、リクエストが来たらジョブキューにいれるなどしてレスポンスを返す実装にするとよさそう
	time.Sleep(5 * time.Minute) // 記事が実際にデプロイされるまで3分程度かかるためスリープする

	// TODO: Twitter APIを叩き、ブログ記事についてのツイートする
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Starting server at Port %d", port)
	http.HandleFunc("/microcms_webhook", webhookHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
