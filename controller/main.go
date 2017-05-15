/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SiCo-DevOps/cfg"
)

var (
	needAAA bool = true
	config       = cfg.Config
	errcode int8
	err     error
	RpcAddr = map[string]string{
		"He": "He.SiCo" + config.Rpc.He,
		"Li": "Li.SiCo" + config.Rpc.Li,
		"Be": "Be.SiCo" + config.Rpc.Be,
		"B":  "B.SiCo" + config.Rpc.B,
		"C":  "C.SiCo" + config.Rpc.C,
		"N":  "N.SiCo" + config.Rpc.N,
	}
)

func AuthPostData(rw http.ResponseWriter, req *http.Request) ([]byte, bool) {
	header := req.Header.Get("Content-Type")
	if header != "application/json" {
		rsp, _ := json.Marshal(&ResponseData{2, "request must follow application/json"})
		httprsp(rw, rsp)
		return nil, false
	}
	body, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	return body, true
}
