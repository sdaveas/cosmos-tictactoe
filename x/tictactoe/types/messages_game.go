package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, guest string) *MsgCreateGame {
  return &MsgCreateGame{
    Creator: creator,
    Guest: guest,
	}
}

func (msg *MsgCreateGame) Route() string {
  return RouterKey
}

func (msg *MsgCreateGame) Type() string {
  return "CreateGame"
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}


var _ sdk.Msg = &MsgAcceptGame{}

func NewMsgAcceptGame(guest string, id string) *MsgAcceptGame {
  return &MsgAcceptGame{
    Guest: guest,
    Id: id,
	}
}

func (msg *MsgAcceptGame) Route() string {
  return RouterKey
}

func (msg *MsgAcceptGame) Type() string {
  return "AcceptGame"
}

func (msg *MsgAcceptGame) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Guest)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptGame) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Guest)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdateGame{}

func NewMsgUpdateGame(caller string, id string, cell uint64) *MsgUpdateGame {

  return &MsgUpdateGame{
    Id: id,
    Caller: caller,
    Cell: cell,
	}
}

func (msg *MsgUpdateGame) Route() string {
  return RouterKey
}

func (msg *MsgUpdateGame) Type() string {
  return "UpdateGame"
}

func (msg *MsgUpdateGame) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Caller)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGame) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Caller)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgDeleteGame(creator string, id string) *MsgDeleteGame {
  return &MsgDeleteGame{
        Id: id,
		Creator: creator,
	}
}
func (msg *MsgDeleteGame) Route() string {
  return RouterKey
}

func (msg *MsgDeleteGame) Type() string {
  return "DeleteGame"
}

func (msg *MsgDeleteGame) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGame) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}


