package mtga

import (
	"encoding/json"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/connect_resp"
)

type ConnectResp struct {
	onConnectResp func(resp connect_resp.Response)
}

func (parser *Parser) parseConnectRespThreadLog(l thread.Log) {
	if parser.onConnectResp != nil {
		var resp connect_resp.Response
		err := json.Unmarshal(l.Raw, &resp)
		if err != nil {
			panic.Fatalln(err)
		}
		parser.onConnectResp(resp)
	}
}

func (resp *ConnectResp) OnConnectResp(callback func(resp connect_resp.Response)) {
	resp.onConnectResp = callback
}
