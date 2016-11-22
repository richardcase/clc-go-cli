package lbaas

import "github.com/centurylinkcloud/clc-go-cli/base"

type LoadBalancerRequest struct {
	RequestID      string `json:"id"`
	Status         string `oneOf:"ACTIVE,COMPLETE"`
	Description    string
	RequestDate    base.JSTimeStamp
	CompletionDate *base.JSTimeStamp `json:",omitempty"`
	Links          []Link            `json:"links"`
}

type LoadBalancer struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	PublicIPAddress   string            `json:"publicIPAddress"`
	PrivateIPAddress  string            `json:"privateIPAddress"`
	Pools             []Pool            `json:"pools"`
	Status            string            `json:"status",oneOf:"ACTIVE,READY,DELETED,UNDER_CONSTRUCTION,UPDATING_CONFIGURATION,FAILED"`
	AccountAlias      string            `json:"accountAlias"`
	DataCenter        string            `json:"dataCenter"`
	KeepAliveRouterID string            `json:"keepalivedRouterId"`
	Version           string            `json:"version,omitempty"`
	CreationTime      base.JSTimeStamp  `json:"creationTime"`
	DeletionTime      *base.JSTimeStamp `json:"deletionTime,omitempty"`
}

type Pool struct {
	ID                  string `json:"id"`
	Port                string `json:"port"`
	LoadBalancingMethod string `json:"loadBalancingMethod",oneOf:"roundrobin,leastconn"`
	Persistence         string `json:"persistence" oneOf:"none,source_ip"`
	IdleTimeout         int    `json:"idleTimeout"`
	LoadBalancingMode   string `json:"loadBalancingMode",oneOf:"tcp,http"`
	Nodes               []Node `json:"nodes"`
	HealthCheck         string `json:"healthCheck"`
}

type Node struct {
	IPAddress   string `json:"ipAddress"`
	PrivatePort string `json:"privatePort"`
}

type Link struct {
	Rel        string `json:"rel"`
	Href       string `json:"href"`
	ResourceID string `json:"resourceId"`
}
