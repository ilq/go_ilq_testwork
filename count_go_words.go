package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type wordCount struct {
	url   string
	count int
}

func main() {
	var totalCount int = 0
	c_qnt := make(chan wordCount)
	gr_count := 0
	printed_counts := 0
	k := 2
	urls := getUrls()

	for _, url := range urls {
		go getWordCount(url, c_qnt, &gr_count)
		gr_count++
		if gr_count >= k {
			receiveCQnt(c_qnt, &totalCount)
			printed_counts++
			gr_count--
		}
	}

	for i := 0; i < (len(urls) - printed_counts); i++ {
		receiveCQnt(c_qnt, &totalCount)
	}

	fmt.Printf("Total: %d\n", totalCount)
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

func getWordCount(url string, c_qnt chan wordCount, gr_count *int) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	count_go := strings.Count(string(body), "go")
	*gr_count--
	c_qnt <- wordCount{url, count_go}
}

func receiveCQnt(c_qnt chan wordCount, totalCount *int) {
	words := <-c_qnt
	fmt.Printf("Count for %s: %d\n", words.url, words.count)
	*totalCount += words.count
}
