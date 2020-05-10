# Тестовое задание по golang.

## Ключи запуска.
```commandline
Usage of /tmp/go-build962375681/b001/exe/count_go_words:
  -k int
        The number of goroutines running at the same time. (default 5)
  -w string
        The number of goroutines running at the same time. (default "go")
  -word string
        The number of goroutines running at the same time. (default "go")
```

## Пример запуска.
```commandline
echo "https://golang.org/\nhttps://golang-blog.blogspot.com/2019/\nhttps://tproger.ru/translations/golang-basics/\nhttps://ru.wikipedia.org/wiki/Go\nhttps://yourbasic.org/golang/switch-statement/\nhttps://golang.org/\nhttps://vk.com/golang\nhttp://golang-book.ru/\n" | go run count_go_words.go
```

## Вывод.
```commandline
Count for https://golang.org/: 38
Count for https://tproger.ru/translations/golang-basics/: 143
Count for https://golang-blog.blogspot.com/2019/: 515
Count for https://golang.org/: 38
Count for https://ru.wikipedia.org/wiki/Go: 310
Count for https://yourbasic.org/golang/switch-statement/: 47
Count for http://golang-book.ru/: 17
Count for https://vk.com/golang: 27
Total: 1135
```
