package keeper

import (
	"github.com/sdaveas/tictactoe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
