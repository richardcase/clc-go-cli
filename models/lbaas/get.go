package lbaas

type GetReq struct {
	DataCenter     string `json:"-" valid:"required" URIParam:"yes"`
	LoadBalancerID string `json:"-" valid:"required" URIParam:"yes"`
}
