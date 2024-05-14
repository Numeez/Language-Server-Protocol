package lsp

type Request struct{
	RPC string `json:"jsonrpc"`
	ID int `json:"id"`
	Method string `json:"method"`

	// We will add params later

}

type Response struct{
	RPC string `json:"jsonrpc"`
	ID *int `json:"id,omitempty"`
	Result string `json:"result"`

	//Result
	//Error

}

type Notification struct{
	RPC string `json:"jsonrpc"`
	Method string `json:"method"`
}