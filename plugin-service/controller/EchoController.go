package controller

import (
	"encoding/json"
	"github.com/goosmesh/goos/core/utils"
	"net/http"
)

func ServiceEcho(w http.ResponseWriter, r *http.Request)  {
	resp := utils.Succeed(nil)
	if err := json.NewEncoder(w).Encode(resp); err != nil{
		panic(err)
	}
}
