package main

import (
	"fmt"
	"sync"
	"github.com/watanabetatsumi/module_practice/greeting"
)

func hello() {
    fmt.Println("Hello World!")
}

func main() {
    // チャネルを作成
    // バッファを10まで設定
    var wg sync.WaitGroup
    wg.Add(1)
    txtCh := make(chan string, 10)
    defer close(txtCh)//mallocみたいなもので、close（free）しとく
    // task を起動
    go task(txtCh,&wg)
    // txtChに値を10個送信
    // バッファに収まっているのでblockingが発生しない
    for i := 0; i < 10; i++ {
        txtCh <- fmt.Sprintf("push txt%d.", i)
    }
    wg.Wait()

	hello()
	message:= greeting.Hello("渡邉")
	fmt.Println(message)
    // チャネルへの送信がブロックされないので
    // goroutineの結果を待たずに処理が終了する

    //もしバッファが10未満なら、fmt.Sprintf("push txt%d.", i)の受け取り手がいなくなり、//blockingが生じる
}