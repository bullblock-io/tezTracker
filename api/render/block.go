package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

// Block renders an app level model to a gennerated OpenAPI model.
func Block(b models.Block) *genModels.BlocksRow {
	ts := b.Timestamp.Unix()
	return &genModels.BlocksRow{
		Level:                    b.Level.Ptr(),
		Proto:                    b.Proto.Ptr(),
		Predecessor:              b.Predecessor.Ptr(),
		Timestamp:                &ts,
		ValidationPass:           b.ValidationPass.Ptr(),
		Fitness:                  b.Fitness.Ptr(),
		Context:                  b.Context,
		Signature:                b.Signature,
		Protocol:                 b.Protocol.Ptr(),
		ChainID:                  b.ChainID,
		Hash:                     b.Hash.Ptr(),
		OperationsHash:           b.OperationsHash,
		PeriodKind:               b.PeriodKind,
		CurrentExpectedQuorum:    b.CurrentExpectedQuorum,
		ActiveProposal:           b.ActiveProposal,
		Baker:                    b.Baker,
		NonceHash:                b.NonceHash,
		ConsumedGas:              b.ConsumedGas,
		MetaLevel:                b.MetaLevel,
		MetaLevelPosition:        b.MetaLevelPosition,
		MetaCycle:                b.MetaCycle,
		MetaCyclePosition:        b.MetaCyclePosition,
		MetaVotingPeriod:         b.MetaVotingPeriod,
		MetaVotingPeriodPosition: b.MetaVotingPeriodPosition,
		ExpectedCommitment:       b.ExpectedCommitment,
	}
}

// Blocks renders a slice of app level Blocks into a slice of OpenAPI models.
func Blocks(bs []models.Block) []*genModels.BlocksRow {
	blocks := make([]*genModels.BlocksRow, len(bs))
	for i := range bs {
		blocks[i] = Block(bs[i])
	}
	return blocks
}
