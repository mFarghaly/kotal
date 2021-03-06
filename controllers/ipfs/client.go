package controllers

import (
	"fmt"

	ipfsv1alpha1 "github.com/kotalco/kotal/apis/ipfs/v1alpha1"
	"github.com/kotalco/kotal/controllers/shared"
	"k8s.io/apimachinery/pkg/runtime"
)

// IPFSClient is IPFS peer client
type IPFSClient interface {
	shared.Client
}

// NewIPFSClient creates a new client for ipfs peer or cluster peer
func NewIPFSClient(obj runtime.Object) (IPFSClient, error) {
	switch peer := obj.(type) {
	case *ipfsv1alpha1.Peer:
		return &GoIPFSClient{peer}, nil
	case *ipfsv1alpha1.ClusterPeer:
		return &GoIPFSClusterClient{peer}, nil
	}
	return nil, fmt.Errorf("no client support for %s", obj)
}
