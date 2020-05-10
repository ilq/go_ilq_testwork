package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

type TotalCount struct {
    sync.Mutex
    count int
}

func (tc *TotalCount) add(n int) {
    tc.Lock()
    tc.count += n
    tc.Unlock()
}

type GoRoutineCount struct {
    sync.Mutex
    count int
}

func (gt *GoRoutineCount) increment() {
    gt.Lock()
    gt.count++
    gt.Unlock()
}

func (gt *GoRoutineCount) decrement() {
    gt.Lock()
    gt.count--
    gt.Unlock()
}

type wordCount struct {
	url   string
	count int
}

func main() {
	totalCount := TotalCount{count: 0}
	cQnt := make(chan wordCount)
	grCount := GoRoutineCount{count: 0}
	printedCounts := 0
	k := 2
	urls := getUrls()

	for _, url := range urls {
		go getWordCount(url, cQnt, &grCount)

		grCount.increment()
		if grCount.count >= k {
			receiveCQnt(cQnt, &totalCount)
			printedCounts++
			grCount.decrement()
		}
	}

	for i := 0; i < (len(urls) - printedCounts); i++ {
		receiveCQnt(cQnt, &totalCount)
	}

	fmt.Printf("Total: %d\n", totalCount.count)
}

func getUrls() []string {
	reader := bufio.NewReader(os.Stdin)
	urls := make([]string, 0)
	for {
		url, err := reader.ReadString('\n')
		if url == "\n" || err == io.EOF {
			return urls
		}
		url = strings.Replace(url, "\n", "", -1)
		urls = append(urls, url)
	}
}

func getWordCount(url string, cQnt chan wordCount, grCount *GoRoutineCount) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	count := strings.Count(string(body), "go")
	grCount.decrement()
	cQnt <- wordCount{url, count}
}

func receiveCQnt(cQnt chan wordCount, totalCount *TotalCount) {
	words := <-cQnt
	fmt.Printf("Count for %s: %d\n", words.url, words.count)
	totalCount.add(words.count)
}
