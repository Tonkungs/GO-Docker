package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	// ถ้าเข้ามาที่ http://localhost:8080 ก็จะไปที่ฟั่งชั้นนี้
	http.Handle("/", logginMiddlware(http.HandlerFunc(handler)))
	// เปิดการทำงานและใส่ พอท 8080
	http.ListenAndServe(":8080", nil)
}

// คืนค่านี้กลับไป
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Hello Main Ton ต้นคุง")
}

// ตัว Middlware ไว้แสดงการทำงานต่างๆ ที่เข้ามา
func logginMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("uri: %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// https://gobyexample.com/
