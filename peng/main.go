package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

var websites []string

func main() {
	//Crawl("http://golang.org/", 4, fetcher)

	/*
		resp, err := http.Get("http://www.vip.com")
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Print(body)
		fmt.Print(resp.Body)
	*/
	//runtime.GOMAXPROCS(4)
	// 各大电商网站首页数据量大小检测
	websites = []string{
		"http://www.51buy.com/", "http://www.360buy.com/", "http://www.tmall.com/", "http://www.taobao.com/",
		"http://china.alibaba.com/", "http://www.paipai.com/", "http://shop.qq.com/", "http://www.lightinthebox.com/",
		"http://www.amazon.cn/", "http://www.newegg.com.cn/", "http://www.vancl.com/", "http://www.yihaodian.com/",
		"http://www.dangdang.com/", "http://www.m18.com/", "http://www.suning.com/", "http://www.hstyle.com/",
		"http://shop.vipshop.com/home.php"}
	// 并发5个运行
	pnum := 10 // 默认设置10个并发测试
	parallelRequest(pnum, websites)
}

func parallelRequest(pnum int, websites []string) {
	total := len(websites)
	if pnum <= 0 {
		pnum = total
	}
	if pnum > total {
		pnum = total
	}
	startTime := time.Now().UnixNano()
	fetchData := make(map[string]string, total)
	execChans := make(chan bool, pnum)
	doneChans := make(chan bool, 1)
	for i := 0; i < total; i++ {
		go request(i, websites[i], execChans, doneChans, fetchData)
	}

	for i := 0; i < total; i++ {
		r := <-doneChans
		<-execChans
		if !r {
			log.Printf("第 %s 项获取失败", i)
		}
	}
	close(doneChans)
	close(execChans)
	processed := float32(time.Now().UnixNano()-startTime) / 1e9
	log.Printf("全部完成，并发数量：%d，共耗时：%.3fs", pnum, processed)
	log.Printf("data: %q", fetchData)
}

func request(i int, url string, execChans chan bool, doneChans chan bool, fetchData map[string]string) {
	execChans <- true
	log.Printf("NO:%02d,url:%s,start...", i, url)
	isOk := false
	startTime := time.Now().UnixNano()
	resp, _ := http.Get(url)
	defer (func() {
		resp.Body.Close()
		doneChans <- isOk
		processed := float32(time.Now().UnixNano()-startTime) / 1e9
		log.Printf("NO:%02d, url:%s, end, status:%t,time:%.3fs", i, url, isOk, processed)
	})()
	body, err := ioutil.ReadAll(resp.Body)
	len := len(body)
	fetchData[url] = fmt.Sprintf("len: %d", len)
	if err == nil {
		isOk = true
	}
}
