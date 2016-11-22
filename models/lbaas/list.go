package lbaas

type ListResp struct {
	LoadBalancers []LoadBalancer `json:"values"`
}
