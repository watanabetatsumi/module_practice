package main

import (
	"fmt"
	"sync"
)

// チャネルを受信専用で受け取る
func task(txtCh <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    // 1秒待つ
    // time.Sleep(2 * time.Second)
    // チャネルから値を受け取る
    for  i :=  0; i < 5; i++{
        // val := <-txtCh
        fmt.Printf("%d-process2\n", <-txtCh)
    }
}
