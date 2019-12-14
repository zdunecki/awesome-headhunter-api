package cursor

import "net/http"

const DEFAULT_CURSOR = "https://fb.com/MONOstudio-393020864107976/"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
