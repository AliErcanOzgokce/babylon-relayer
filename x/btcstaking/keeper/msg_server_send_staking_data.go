package keeper

import (
	"context"

	"github.com/AliErcanOzgokce/babylon-relayer/x/btcstaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendStakingData(goCtx context.Context, msg *types.MsgSendStakingData) (*types.MsgSendStakingDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendStakingDataResponse{}, nil
}
