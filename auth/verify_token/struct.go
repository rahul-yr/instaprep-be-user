package verifytoken

type Response struct {
	Status bool `json:"status"`
}

type RequestParams struct {
	Token string `json:"token"  xml:"token" form:"token"`
}
