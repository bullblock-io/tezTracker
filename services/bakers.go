package services

import (
	"github.com/bullblock-io/tezTracker/models"
)

const (
	BlocksInCycle = 4096
	PreservedCycles = 5
	XTZ = 1000000
	BlockSecurityDeposit = 512 * XTZ
	EndorsementSecurityDeposit = 64 * XTZ
	BlockReward = 16 * XTZ
	EndorsementReward = 2 * XTZ
)


// BakerList retrives up to limit of bakers after the specified id.
func (t *TezTracker) BakerList(limit uint64, after string) ([]models.Baker, error) {
	r := t.repoProvider.GetBaker()
	return r.List(limit, after)
}

func (t *TezTracker) GetCurrentCycle() (int64, error){
	r := t.repoProvider.GetBlock()
	lastBlock, err := r.Last()
	if err != nil{
		return 0, err
	}
	return lastBlock.MetaCycle, nil
}

func GetFirstPreservedBlock(currentCycle int64) (fpb int64){
	fpc := currentCycle - PreservedCycles

	if fpc > 0 {
		fpb = fpc * BlocksInCycle + 1
	}
	return fpb
}

func (t *TezTracker) GetBakerInfo(accountID string) (bi *models.BakerInfo, err error){
	r := t.repoProvider.GetBaker()
	found, delegate, err := r.Find(accountID)
	if err != nil{
		return nil, err
	}
	if !found{
		return nil, nil
	}
	bi = &models.BakerInfo{Delegate: delegate}
	curCycle,err := t.GetCurrentCycle()
	if err != nil{
		return bi, err
	}
	fpb := GetFirstPreservedBlock(curCycle)
	counter,err := r.BlocksCountBakedBy([]string{accountID}, fpb)
	if err != nil{
		return bi, err
	}
	var blocksCount int64
	if len(counter) == 1{
		blocksCount = counter[0].Count
	}
	bi.BakingDeposits = blocksCount * BlockSecurityDeposit
	bi.BakingRewards = blocksCount * BlockReward

	endCounter,err := r.EndorsementsCountBy([]string{accountID}, fpb)
	if err != nil{
		return bi, err
	}
	var endorsementCount int64
	var endorsementWeight float64
	if len(endCounter) == 1{
		endorsementCount = endCounter[0].Count
		endorsementWeight = endCounter[0].Weight
	}
	bi.EndorsementDeposits = endorsementCount * EndorsementSecurityDeposit
	bi.EndorsementRewards = int64(endorsementWeight * EndorsementReward)

	return bi, nil
}