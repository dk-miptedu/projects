package logs

import (
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, msg string) {
	logHandlers := make(map[string]interface{})
	logHandlers["status"] = status
	logHandlers["method"] = r.Method
	logHandlers["path"] = r.URL.Path
	logHandlers["error"] = msg
	w.WriteHeader(status)
	w.Write([]byte(msg))
	LogConfigsParams("handlers", logHandlers)
}
func SuccessHandler(w http.ResponseWriter, r *http.Request, status int, msg string) {
	logHandlers := make(map[string]interface{})
	logHandlers["method"] = r.Method
	logHandlers["path"] = r.URL.Path
	logHandlers["message"] = msg
	w.WriteHeader(status)
	w.Write([]byte(msg))
	LogConfigsParams("handlers", logHandlers)
}
