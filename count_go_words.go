package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	// 	"time"
)

type counts struct {
	url   string
	count int
}

func main() {
	var totalCount int = 0
	c_counts := make(chan counts)
	reader := bufio.NewReader(os.Stdin)
	gr_count := 0
	count_lost_gr := 0
	k := 5
	for {
		url, err := reader.ReadString('\n')
		if url == "\n" || err == io.EOF {
			for i := 0; i < count_lost_gr; i++ {
				count := <-c_counts
				fmt.Printf("%s: %d\n", count.url, count.count)
				totalCount += count.count
			}
			fmt.Printf("Total count: %d\n", totalCount)
			break
		}
		url = strings.Replace(url, "\n", "", -1)
		gr_count++
		count_lost_gr++

		go func(url string, gr_count *int) {
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			count_go := strings.Count(string(body), "go")
			*gr_count--
			c_counts <- counts{url, count_go}
		}(url, &gr_count)

		if gr_count >= k {
			count := <-c_counts
			count_lost_gr--
			fmt.Printf("%s: %d\n", count.url, count.count)
			totalCount += count.count
			gr_count--
		}
	}
}