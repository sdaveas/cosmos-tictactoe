package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreateGame{}, "tictactoe/CreateGame", nil)
cdc.RegisterConcrete(&MsgAcceptGame{}, "tictactoe/AcceptGame", nil)
cdc.RegisterConcrete(&MsgMakeMove{},   "tictactoe/MakeMove", nil)
cdc.RegisterConcrete(&MsgDeleteGame{}, "tictactoe/DeleteGame", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreateGame{},
	&MsgAcceptGame{},
	&MsgMakeMove{},
	&MsgDeleteGame{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
