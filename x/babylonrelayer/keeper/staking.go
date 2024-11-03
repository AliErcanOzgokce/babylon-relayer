package keeper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/username/babylon-chain/x/staking/types"
)

func (k Keeper) FetchStakingData(ctx sdk.Context) (types.StakingData, error) {
	url := "https://babylonapi.com/staking-data"
	resp, err := http.Get(url)
	if err != nil {
		return types.StakingData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.StakingData{}, err
	}

	var data types.StakingData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return types.StakingData{}, err
	}

	return data, nil
}
