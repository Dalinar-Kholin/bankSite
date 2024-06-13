package endpoints

import (
	"WDB/views"
	"fmt"
	"net/http"
	"time"
)

func (h *Handlers) ThreadTest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("start")
	time.Sleep(5 * time.Second)
	fmt.Printf("end")
	views.ResponseWithError(w, 400, "bada")
}
