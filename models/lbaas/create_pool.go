package lbaas

type CreatePoolReq struct {
	DataCenter          string `json:"-" valid:"required" URIParam:"yes"`
	LoadBalancerID      string `json:"-" valid:"required" URIParam:"yes"`
	Port                string `json:"port" valid:"required"`
	LoadBalancingMethod string `json:"loadBalancingMethod" oneOf:"roundrobin,leastconn"`
	LoadBalancingMode   string `json:"loadBalancingMode" valid:"required" oneOf:"http,tcp"`
	HealthCheck         string `json:"healthCheck,omitempty"`
	Persistence         string `json:"persistence" oneOf:"none,source_ip"`
	IdleTimeout         int    `json:"idleTimeout"`
}
