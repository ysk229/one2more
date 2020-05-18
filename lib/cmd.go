package lib

import (
    "bufio"
    "io"
    "log"
    "os/exec"
)

type Cmd struct {

}

func NewCmd() *Cmd {
    return &Cmd{}
}
func (c *Cmd) Run (name string,arg ...string) error {
    cmd:= exec.Command(name ,arg...)
    log.Println(cmd.Args)
    stdout,err := cmd.StdoutPipe()

    if err != nil {
        log.Println(cmd.Args)
        log.Println("输出管道出错")
        log.Fatalln(err)
    }

    if err = cmd.Start(); err != nil {
        log.Println(cmd.Args)
        log.Println("执行脚本出错")
        log.Fatalln(err)
        return err
    }
    //创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
    reader := bufio.NewReader(stdout)
    //实时循环读取输出流中的一行内容
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
            break
        }
        r := []rune(line)
        log.Println(string(r))
    }

    //阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
    if err = cmd.Wait(); err != nil {
        log.Println(cmd.Args)
        log.Println("执行中出错")
        log.Fatalln(err)
        return err
    }
    return nil
}