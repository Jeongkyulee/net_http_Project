package main

import (
	"fmt"
	"net/http"
	"net_http_Project/godule"
)

func main() {
	godule.CheckVariable("hello")
	s := "Hello, world!"

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// HTML로 웹 페이지 작성
		html := `
		<html>
		<head>
			<title>Hello</title>
		</head>
		<body>
			<span class="hello">` + s + `</span>
		</body>
		</html>
		`

		res.Header().Set("Content-Type", "text/html") // HTML 헤더 설정
		res.Write([]byte(html))                       // 웹 브라우저에 응답
	})

	http.Handle( // /assets/ 경로에 접근했을 때
		// 파일 서버를 동작시킴
		"/assets/",
		http.StripPrefix( // 파일 서버를 실행할 때 assets
			// 디렉터리를 지정했으므로 URL 경로에서 /assets/ 삭제
			"/assets/",
			http.FileServer(http.Dir("assets")), // 웹 서버에서 assets
			// 디렉터리 아래의 파일을 표시
		),
	)
	defer func() {
		http.ListenAndServe(":80", nil) // 80번 포트에서 웹 서버 실행
	}()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
