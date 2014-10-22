package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		fmt.Println("listening on :" + port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			panic(err)
		}
	}()

	port2 := "7070"
	go func() {
		fmt.Println("listening on :" + port2)
		err := http.ListenAndServe(":"+port2, nil)
		if err != nil {
			panic(err)
		}
	}()
	select {}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	var sb bytes.Buffer
	res.Header().Set("Content-Type", "text/html")
	sb.WriteString("Hello,Go!<br><br>")

	sb.WriteString("Host=" + req.Host + "<br>")
	sb.WriteString("Referer=" + req.Referer() + "<br>")
	reqHeader := req.Header

	for n, v := range reqHeader {
		sb.WriteString(n + "=")
		for _, av := range v {
			sb.WriteString(av + ",")
		}
		sb.WriteString("<br>")
	}
	env := os.Environ()
	sb.WriteString("<br><br>System Env Variables<br>===============<br>")
	for _, v := range env {
		sb.WriteString(v)
		sb.WriteString("<br>")
	}

	res.Write(sb.Bytes())
}

