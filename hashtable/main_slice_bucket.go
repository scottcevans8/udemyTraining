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

func scanThePage(res *http.Response) (buckets [][]string) {
    scanner := bufio.NewScanner(res.Body)
    defer res.Body.Close()
    scanner.Split(bufio.ScanWords)
    buckets = make([][]string, 12)
    for scanner.Scan() {
        word := scanner.Text()
        n := hash(word,12)
        buckets[n] = append(buckets[n], word)
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

func printResults (buckets [][]string) {
    for i := 0; i < 12; i++ {
        fmt.Println(i," -", len(buckets[i]))
    }
    //fmt.Println(buckets[6])
}
