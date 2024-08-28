// Copyright (C) 2018 Betalo AB - All Rights Reserved

package main

import (
	"fmt"
	"net/http"
	"os"
	"realIp/realips"
	"runtime"
	"strings"
)

func main() {
	http.HandleFunc("/realip", handlerRealIp)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("handlerRealIp server error: ", err)
	} else {
		fmt.Println("handlerRealIp server run ok at: 0.0.0.0:8888")
	}
}
func handlerRealIp(w http.ResponseWriter, r *http.Request) {
	ip := realips.RequestIp(r)
	if runtime.GOOS == "linux" {
		if len(strings.Split(ip, ".")) == 4 {
			//ReplaceGodaddyDNSIP("mybbs.vip", ip)
			_ = os.WriteFile("/xxx/publicip.txt", []byte(ip), 0644)
			w.Write([]byte(ip))
		} else {
			w.Write([]byte(ip))
		}
	} else {
		w.Write([]byte(ip))
	}
}
