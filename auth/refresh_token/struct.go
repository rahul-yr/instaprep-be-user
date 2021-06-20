package refreshtoken

type Response struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type RequestParams struct {
	Token string `json:"token"  xml:"token" form:"token"`
}
