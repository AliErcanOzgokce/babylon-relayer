package keeper_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BtcstakingKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestFetchStakingData(t *testing.T) {
	k, ctx := keepertest.BtcstakingKeeper(t)

	// Mock server to simulate API response
	mockData := types.StakingData{
		// Fill with appropriate mock data
	}
	mockResponse, _ := json.Marshal(mockData)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponse)
	}))
	defer server.Close()

	// Override the URL in FetchStakingData
	originalURL := "https://staking-api.staging.babylonchain.io/v1/delegation/"
	k.FetchStakingData = func(ctx sdk.Context, txHash string) (types.StakingData, error) {
		url := server.URL + "/" + txHash
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

	// Call FetchStakingData
	txHash := "sample-tx-hash"
	data, err := k.FetchStakingData(ctx, txHash)
	require.NoError(t, err)
	require.Equal(t, mockData, data)
}package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/AliErcanOzgokce/babylon-relayer/testutil/keeper"
	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
)

func TestVerifyStakingData(t *testing.T) {
	k, ctx := keepertest.BtcstakingKeeper(t)

	// Mock response for FetchStakingData
	mockStakingData := types.StakingData{
		// Fill with appropriate mock data
	}

	// Mock the FetchStakingData function
	k.FetchStakingData = func(ctx sdk.Context, txHash string) (types.StakingData, error) {
		return mockStakingData, nil
	}

	// Call VerifyStakingData
	data, err := k.VerifyStakingData(ctx)
	require.NoError(t, err)
	require.Equal(t, mockStakingData, data)
}