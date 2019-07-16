package services

import (
	"github.com/bullblock-io/tezTracker/repos/block"
	"github.com/bullblock-io/tezTracker/repos/operation_groups"
)

type (
	// TezTracker is the main service for tezos tracker. It has methods to process all the user's requests.
	TezTracker struct {
		repoProvider Provider
	}

	// Provider is the abstract interface to get any repository.
	Provider interface {
		GetBlock() block.Repo
		GetOperationGroup() operation_groups.Repo
	}
)

// New creates a new TexTracker service using the repository provider.
func New(rp Provider) *TezTracker {
	return &TezTracker{repoProvider: rp}
}
