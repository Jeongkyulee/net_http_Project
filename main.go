package net_http_Project

import (
	"net/http"
)

func main() {
	initconst()
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!")) // 웹 브라우저에 응답
	}) // / 경로에 접속했을 때 실행할 함수 설정

	http.ListenAndServe(":80", nil) // 80번 포트에서 웹 서버 실행
}

func nethttp_CanonicalHeaderKey() {
	//http.CanonicalHeaderKey()

}
