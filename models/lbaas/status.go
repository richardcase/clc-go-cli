package lbaas

type StatusReq struct {
	DataCenter string `json:"-" valid:"required" URIParam:"yes"`
	RequestId  string `json:"-" valid:"required" URIParam:"yes"`
}
