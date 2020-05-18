package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "one2more/lib"
    "os"
    "sync"
    "sync/atomic"
    "time"
)

func waitGroup(files []os.FileInfo) {
    wg := &sync.WaitGroup{}
    wg.Add(len(files))
    for _, file := range files {
        go func(file os.FileInfo) {
            if !file.IsDir() {
                fileName := fmt.Sprintf("./shell/%s", file.Name())
                err := lib.NewCmd().Run("/usr/bin/bash", "-c", fileName+" 2>&1 ")
                if err != nil {
                    log.Println("执行失败")
                    log.Fatal(err)
                }
            }
            defer wg.Done()
        }(file)
    }
    wg.Wait()
}

//原子记数器
func atom(files []os.FileInfo) {
    var total int32
    for _, file := range files {
        go func(file os.FileInfo) {
            if !file.IsDir() {
                fileName := fmt.Sprintf("./shell/%s", file.Name())
                err := lib.NewCmd().Run("/usr/bin/bash", "-c", fileName+" 2>&1 ")
                if err != nil {
                    log.Println("执行失败")
                    log.Fatal(err)
                }
            }
            defer atomic.AddInt32(&total, 1)

        }(file)
    }
    for atomic.LoadInt32(&total) < int32(len(files)) {
        time.Sleep(time.Microsecond)
    }
}
func chanel(files []os.FileInfo) {
    ch := make(chan struct{})
    for _, file := range files {
        go func(file os.FileInfo) {
            if !file.IsDir() {
                fileName := fmt.Sprintf("./shell/%s", file.Name())
                err := lib.NewCmd().Run("/usr/bin/bash", "-c", fileName+" 2>&1 ")
                if err != nil {
                    log.Println("执行失败")
                    log.Fatal(err)
                }
                ch <- struct{}{}
                log.Println("执行完毕")
            }
        }(file)
    }
    //读取chan
    for _, file := range files {
        <-ch
        log.Printf("flie %s,处理完毕\n", file.Name())
    }

}
func selects(files []os.FileInfo) {
    ch := make(chan struct{})
    for _, file := range files {
        go func(file os.FileInfo) {
            if !file.IsDir() {
                fileName := fmt.Sprintf("./shell/%s", file.Name())
                err := lib.NewCmd().Run("/usr/bin/bash", "-c", fileName+" 2>&1 ")
                if err != nil {
                    log.Println("执行失败")
                    log.Fatal(err)
                }
                ch <- struct{}{}
                log.Println("执行完毕")
            }
        }(file)
    }
    //读取chan
    for _, file := range files {
        tm := time.NewTimer(time.Second * 60) //给通道创建容忍时间，如果5s内无法读写，就即刻返回

        select {
        case <-ch:
            log.Println("任务",file.Name(),"处理完毕")
        case <-tm.C:
            fmt.Println("send data timeout!")
            break
        }
    }

}
func main() {
    files, _ := ioutil.ReadDir("./shell/")
    //waitGroup(files)
    //atom(files)
    //chanel(files)
    selects(files)
}
