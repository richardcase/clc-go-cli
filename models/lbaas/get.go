package lbaas

type GetReq struct {
	LoadBalancerIdentifier `argument:"composed" URIParam:"LoadBalancerId" json:"-"`
	DataCenter             string `json:"-" valid:"required" URIParam:"yes"`
}
