package rpc_client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tzblock "github.com/bullblock-io/go-tezos/v2/block"
	tzc "github.com/bullblock-io/go-tezos/v2/client"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client/baking_rights"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client/snapshots"
	genmodels "github.com/bullblock-io/tezTracker/services/rpc_client/models"
)

const headBlock = "head"
const BlocksInCycle = 4096

type Tezos struct {
	client    *client.Tezosrpc
	network   string
	tzcClient *tzblock.BlockService
}

func New(cfg client.TransportConfig, network string) *Tezos {

	return &Tezos{
		client:    client.NewHTTPClientWithConfig(nil, &cfg),
		network:   network,
		tzcClient: tzblock.NewBlockService(tzc.NewClient(cfg.Host)),
	}
}
func (t *Tezos) RightsFor(ctx context.Context, blockFrom, blockTo int64) ([]models.BakingRight, error) {
	params := baking_rights.NewGetBakingRightsParamsWithContext(ctx)
	params.SetNetwork(t.network)
	params.SetBlock(strconv.FormatInt(blockFrom, 10))
	levels := []string{}
	for b := blockFrom; b <= blockTo; b++ {
		levels = append(levels, strconv.FormatInt(b, 10))
	}
	params.SetLevel(levels)
	resp, err := t.client.BakingRights.GetBakingRights(params)
	if err != nil {
		return nil, err
	}
	rights := make([]models.BakingRight, len(resp.Payload))
	for i := range resp.Payload {
		if resp.Payload[i] != nil {
			rights[i] = models.BakingRight{
				Level:         resp.Payload[i].Level,
				Priority:      int(resp.Payload[i].Priority),
				Delegate:      resp.Payload[i].Delegate,
				EstimatedTime: time.Time(resp.Payload[i].EstimatedTime),
			}
		}
	}
	return rights, nil
}

func (t *Tezos) FutureRightsFor(ctx context.Context, blockFrom, blockTo int64) ([]models.FutureBakingRight, error) {
	params := baking_rights.NewGetBakingRightsParamsWithContext(ctx)
	params.SetNetwork(t.network)
	params.SetBlock(headBlock)
	levels := []string{}
	for b := blockFrom; b <= blockTo; b++ {
		levels = append(levels, strconv.FormatInt(b, 10))
	}
	params.SetLevel(levels)
	resp, err := t.client.BakingRights.GetBakingRights(params)
	if err != nil {
		return nil, err
	}
	rights := make([]models.FutureBakingRight, len(resp.Payload))
	for i := range resp.Payload {
		if resp.Payload[i] != nil {
			rights[i] = genRightToModel(*resp.Payload[i])
		}
	}
	return rights, nil
}

func genRightToModel(m genmodels.BakingRight) models.FutureBakingRight {
	return models.FutureBakingRight{
		Level:         m.Level,
		Priority:      int(m.Priority),
		Delegate:      m.Delegate,
		EstimatedTime: time.Time(m.EstimatedTime),
	}
}

func (t *Tezos) SnapshotBlockForCycle(ctx context.Context, cycle int64, useHead bool) (int64, error) {
	params := snapshots.NewGetRollSnapshotParamsWithContext(ctx).WithCycle(cycle).WithNetwork(t.network)
	if useHead {
		params.SetBlock(headBlock)
	} else {
		level := cycle*BlocksInCycle + 1
		params.SetBlock(strconv.FormatInt(level, 10))
	}

	resp, err := t.client.Snapshots.GetRollSnapshot(params)
	if err != nil {
		return 0, err
	}
	snapshot := resp.Payload
	snapBlock := ((cycle-7)*4096 + 1) + (snapshot+1)*256 - 1
	return snapBlock, nil
}

func (t *Tezos) DoubleBakingEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleBakingEvidence, err error) {
	block, err := t.tzcClient.Get(blockLevel)
	if err != nil {
		return dee, err
	}
	for i := range block.Operations {
		for _, op := range block.Operations[i] {
			if strings.EqualFold(op.Hash, operationHash) {
				dee, err := ToDoubleBakingEvidence(op)
				if err != nil {
					return dee, err
				}
				dee.BlockLevel = int64(block.Header.Level)
				dee.BlockHash = block.Hash
				return dee, err
			}
		}
	}
	return dee, fmt.Errorf("not found")
}

func ToDoubleBakingEvidence(op tzblock.Operations) (dee models.DoubleBakingEvidence, err error) {
	for i := range op.Contents {
		if op.Contents[i].Bh1 != nil {
			dee.DenouncedLevel = int64(op.Contents[i].Bh1.Level)
			dee.Priority = op.Contents[i].Bh1.Priority
			if meta := op.Contents[i].Metadata; meta != nil {
				for _, bu := range meta.BalanceUpdates {
					if strings.EqualFold(bu.Kind, "freezer") {
						switch strings.ToLower(bu.Category) {
						case "deposits":
							dee.Offender = bu.Delegate
							change, err := strconv.ParseInt(bu.Change, 10, 64)
							if err != nil {
								return dee, err
							}
							dee.LostDeposits = -change
						case "rewards":
							change, err := strconv.ParseInt(bu.Change, 10, 64)
							if err != nil {
								return dee, err
							}
							if change < 0 {
								dee.LostRewards = -change
							} else {
								dee.BakerReward = change
								dee.EvidenceBaker = bu.Delegate
							}

						}
					}
				}
			}
			return dee, nil
		}
	}
	return dee, fmt.Errorf("not a double baking evidence")

}

func (t *Tezos) DoubleEndrsementEvidenceLevel(ctx context.Context, blockLevel int, operationHash string) (int64, error) {
	block, err := t.tzcClient.Get(blockLevel)
	if err != nil {
		return 0, err
	}
	for i := range block.Operations {
		for _, op := range block.Operations[i] {
			if strings.EqualFold(op.Hash, operationHash) {
				return GetDoubleEndrsementEvidenceLevel(op)
			}
		}
	}
	return 0, fmt.Errorf("not found")
}

func GetDoubleEndrsementEvidenceLevel(op tzblock.Operations) (int64, error) {
	for i := range op.Contents {
		if op.Contents[i].Op1 != nil {
			return int64(op.Contents[i].Op1.Operations.Level), nil
		}
	}
	return 0, fmt.Errorf("not a double endorsement evidence")
}
