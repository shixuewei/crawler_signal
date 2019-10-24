package engine

import (
	"learning/crawler_signal/fetcher"
	"log"
)

//单机版爬虫主引擎，负责调度各个任务的执行

func Run(seeds ...Request) { //将请求送入引擎（因为不知道有多少请求，所以定义为切片）
	//初始化请求切片
	var requests []Request
	//循环遍历将seeds中的Request送入到requests中
	for _, v := range seeds {
		requests = append(requests, v)
	}
	//开始执行具体的处理
	//具体的请求过程为：1、用切片形成一个队列，将所传入的Request按照FIFO排列好
	//2、当该切片中的长度为0时，说明没有要执行的Request了
	for len(requests) > 0 {
		//获取得到第一个Request
		r := requests[0]
		//重新切片，相当于将所有后续内容前移
		requests = requests[1:]
		//输出
		log.Println("Fetching %s", r.Url)
		//调用fetch方法，开始获取每一个url的具体内容
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher : error"+"fetching url %s: %v", r.Url, err)
			continue
		}

		//当得到需要的整个页面时，我们可以进行执行相应的提取出我们需要的信息
		parseResult := r.ParserFunc(body)
		//将我们在parseResult中获取到的接下来的url存入队列中（这块就体现了我们的广度优先的算法）
		requests = append(requests, parseResult.Requests...)
		//将parseResult中的item输出
		for _, v := range parseResult.Items {
			log.Printf("Got item %v", v)
		}

	}
}
