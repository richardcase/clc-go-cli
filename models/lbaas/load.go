package lbaas

import (
	"fmt"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/base"
)

type loadBalancerList struct {
	LoadBalancers []LoadBalancer `json:"values"`
}

func Load(cn base.Connection) ([]LoadBalancer, error) {
	var balancers loadBalancerList

	err := cn.ExecuteRequest("GET", "https://api.loadbalancer.ctl.io/{accountAlias}/loadbalancers", nil, &balancers)
	if err != nil {
		return nil, err
	}
	return balancers.LoadBalancers, nil
}

func IDByName(cn base.Connection, name string) (string, error) {
	balancers, err := Load(cn)
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, b := range balancers {
		if strings.ToLower(b.Name) == strings.ToLower(name) {
			matched = append(matched, b.ID)
		}
	}

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no LBaaS frameworks with name %s.", name)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There are more than one LBaaS frameworks with name %s. Please, specify an ID.", name)
	}
}

func GetNames(cn base.Connection) ([]string, error) {
	balancers, err := Load(cn)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, b := range balancers {
		names = append(names, b.Name)
	}
	return names, nil
}
