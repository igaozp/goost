package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// 获取数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建通道，接受匹配后的结果
	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	// waitGroup 数量设置为数据源数量，数据源 waitGroup goroutine 一一对应
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// waitGroup 计数器减一
			waitGroup.Done()
		}(matcher, feed)
	}

	// 监听 waitGroup 是否完成
	go func() {
		// 监听 waitGroup 计数器是否清零（阻塞），清零后继续执行
		waitGroup.Wait()
		// 关闭通道
		close(results)
	}()

	Display(results)
}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
