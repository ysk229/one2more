package lib

import (
    "testing"
)

func TestCmd_Run(t *testing.T) {
    type args struct {
        name string
        arg  []string
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {"ping",args{"ping",[]string{"www.baidu.com","-n","3"}},false},
    }
    for _, tt := range tests {
       t.Run(tt.name, func(t *testing.T) {
            c := &Cmd{}
            if err := c.Run(tt.args.name, tt.args.arg...); (err != nil) != tt.wantErr {
                t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}