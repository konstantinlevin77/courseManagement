package handlers

import "net/http"

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	message := "Hello, world!"

	_, _ = w.Write([]byte(message))

}
