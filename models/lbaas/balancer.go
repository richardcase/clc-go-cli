package lbaas

import (
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/base"
)

type LoadBalancerIdentifier struct {
	LoadBalancerId   string
	LoadBalancerName string
}

func (b *LoadBalancerIdentifier) Validate() error {
	if (b.LoadBalancerId == "") == (b.LoadBalancerName == "") {
		return fmt.Errorf("Exactly one of the load-balancer-id and load-balancer-name properties must be specified.")
	}
	return nil
}

func (b *LoadBalancerIdentifier) InferID(cn base.Connection) error {
	if b.LoadBalancerName == "" {
		return nil
	}

	id, err := IDByName(cn, b.LoadBalancerName)
	if err != nil {
		return err
	}
	b.LoadBalancerId = id
	return nil
}

func (b *LoadBalancerIdentifier) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "LoadBalancerName" {
		return nil, nil
	}

	return GetNames(cn)
}
