package main 


import "net/http"

func main() {
 http.Handle("/", Hello{message: "Run..."})
 http.ListenAndServe(":3000", nil)
}

type Hello struct {
 message string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte(h.message))
}
