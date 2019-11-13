package rpc_client

import (
	"context"
	"strconv"
	"time"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client"
	"github.com/bullblock-io/tezTracker/services/rpc_client/client/baking_rights"
	genmodels "github.com/bullblock-io/tezTracker/services/rpc_client/models"
)

const headBlock = "head"

type Tezos struct {
	client  *client.Tezosrpc
	network string
}

func New(cfg client.TransportConfig, network string) *Tezos {
	return &Tezos{
		client:  client.NewHTTPClientWithConfig(nil, &cfg),
		network: network,
	}
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
