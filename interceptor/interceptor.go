/**
 * @Author: evnxo
 * @Description:
 * @File:  interceptor
 * @Version: 1.0.0
 * @Date: 2021/7/3 3:52 下午
 */

package interceptor

import (
	"fmt"
	"net/http"
	"time"
)

func NewElapsedTimeInterceptor() MiddlewareInterceptor {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		defer func() {
			endTime := time.Now()
			elapsed := endTime.Sub(startTime)
			fmt.Println("logTime", elapsed)
		}()
		next(w, r)
	}
}

//func NewRequestIdInterceptor() MiddlewareInterceptor {
//	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
//		if r.Headers.Get("X-Request-Id") == "" {
//			r.Headers.Set("X-Request-Id", generateRequestId())
//		}
//		next(w, r)
//	}
//}
