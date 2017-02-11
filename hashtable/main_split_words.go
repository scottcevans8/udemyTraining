package main

import (
    "fmt"
    "net/http"
    "bufio"
    "log"
    "os"
)

func main() {
    res := getResource()
    words := parseResults(res)
    printWords(words)
}

func getResource() (res *http.Response) {
    res,err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
    if(err != nil) {
        log.Fatalln(err)
    }
    return
}

func parseResults(res *http.Response) (words map[string]string) {
    words = make(map[string]string)

    sc := bufio.NewScanner(res.Body)
    sc.Split(bufio.ScanWords)

    for sc.Scan() {
        words[sc.Text()] = ""
    }
    if err := sc.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "ERROR: reading input:", err)
    }
    return
}

func printWords (words map[string]string) {
    i := 0
    for k,_ := range words {
        fmt.Println(k)
        if(i >= 10) {
            break
        }
        i++
    }
}

