package tictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sdaveas/tictactoe/x/tictactoe/types"
	"github.com/sdaveas/tictactoe/x/tictactoe/keeper"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateGame) (*sdk.Result, error) {
	k.CreateGame(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgAcceptGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgAcceptGame) (*sdk.Result, error) {
	k.AcceptGame(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgMakeMove(ctx sdk.Context, k keeper.Keeper, msg *types.MsgMakeMove) (*sdk.Result, error) {

	k.MakeMove(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteGame(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteGame) (*sdk.Result, error) {
    if !k.HasGame(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
    }
    if msg.Creator != k.GetGameOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
    }

	k.DeleteGame(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
