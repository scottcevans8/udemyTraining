package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
)

func main() {
    res := getBook()
    bucket := scanThePage(res)
    printResults(bucket)
}

func getBook() (res *http.Response) {
    res,err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
    if(err != nil) {
        log.Fatalln(err)
    }
    return
}

func scanThePage(res *http.Response) (buckets map[int]map[string]int) {
    scanner := bufio.NewScanner(res.Body)
    defer res.Body.Close()
    scanner.Split(bufio.ScanWords)
    buckets = make(map[int]map[string]int)
    for i := 0; i < 12; i++ {
        buckets[i] = make(map[string]int)
    }
    for scanner.Scan() {
        word := scanner.Text()
        n := hash(word,12)
        buckets[n][word]++
    }
    return
}

func hash(word string, buckets int) int {
    var sum int
    for _,v := range word {
        sum += int(v)
    }
    return sum % buckets
}

func printResults (buckets map[int]map[string]int) {
    for k,v := range buckets[6] {
        fmt.Println(v, " \t- ", k)
    }
}
