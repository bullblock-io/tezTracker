package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

// OperationGroup renders an app level model to a gennerated OpenAPI model.
func OperationGroup(b models.OperationGroup) *genModels.OperationGroupsRow {
	return &genModels.OperationGroupsRow{
		Protocol:  &b.Protocol,
		ChainID:   b.ChainID,
		Hash:      &b.Hash,
		Branch:    &b.Branch,
		Signature: b.Signature,
		BlockID:   &b.BlockID,
	}
}
