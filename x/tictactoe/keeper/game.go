package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sdaveas/tictactoe/x/tictactoe/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
    "crypto/sha1"
    "strings"
)

// GetGameCount get the total number of game
func (k Keeper) GetGameCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameCountKey))
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetGameCount set the total number of game
func (k Keeper) SetGameCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameCountKey))
	byteKey := types.KeyPrefix(types.GameCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateGame(ctx sdk.Context, msg types.MsgCreateGame) {
	// Create the game
    count := k.GetGameCount(ctx)
    var game = types.Game{
        Creator: msg.Creator,
        Id: strconv.FormatInt(count, 10),
        Guest: msg.Guest,
        Oplayer: types.Player_NONE,
        Xplayer: types.Player_NONE,
        Winner: types.Player_NONE,
        Board: "         ",
        Status: types.GameStatus_OPEN,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
    key := types.KeyPrefix(types.GameKey + game.Id)
    value := k.cdc.MustMarshalBinaryBare(&game)
    store.Set(key, value)

    // Update game count
    k.SetGameCount(ctx, count+1)
}

func assert_valid_accept(game types.Game, msg types.MsgAcceptGame) {
    if game.Status != types.GameStatus_OPEN {
		panic("Game is not Open at the moment.")
	}
    if game.Guest != msg.Guest {
        panic("This is not your game")
    }
}

func assign_players(creator string, guest string) (types.Player, types.Player) {
    h := sha1.New()
    h.Write([]byte(creator + guest))
    bs := h.Sum(nil)
    var xplayer types.Player
    var oplayer types.Player
    if bs[0] >> 7 == 1 {
        xplayer = types.Player_CREATOR
        oplayer = types.Player_GUEST
    } else {
        oplayer = types.Player_CREATOR
        xplayer = types.Player_GUEST
    }
    return xplayer, oplayer
}

func (k Keeper) AcceptGame(ctx sdk.Context, msg types.MsgAcceptGame) {
    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
    game := k.GetGame(ctx, msg.Id)

    assert_valid_accept(game, msg)

    game.Xplayer, game.Oplayer = assign_players(game.Creator, game.Guest)
    game.Status = types.GameStatus_RUNNING

    keyPref := types.KeyPrefix(types.GameKey + msg.Id)
    b := k.cdc.MustMarshalBinaryBare(&game)
    store.Set(keyPref, b)
}

func expected_role(game types.Game) types.Player {
    x_counter := strings.Count(game.Board, X)
    o_counter := strings.Count(game.Board, O)
    if x_counter == o_counter {
        return game.Xplayer
    } else if x_counter > o_counter {
        return game.Oplayer
    } else {
        panic("Illegal board state")
    }
}

func caller_role(game types.Game, msg types.MsgMakeMove) types.Player {
    if msg.Caller == game.Creator {
        return types.Player_CREATOR
    }
    return types.Player_GUEST
}

func assert_valid_move(game types.Game, msg types.MsgMakeMove, caller_role types.Player) {
    if game.Creator != msg.Caller && game.Guest != msg.Caller {
        panic("This is not your game")
    }
    if game.Status != types.GameStatus_RUNNING {
		panic("Game is not Running at the moment.")
	}
    if msg.Cell > 8 {
        panic("Cell index out of range")
    }
    if game.Board[msg.Cell] != ' ' {
        panic("Cell already allocated")
    }
    if caller_role != expected_role(game) {
        panic("Not your turn")
    }
}

func get_move(role types.Player, game types.Game) string {
    if role == game.Xplayer {
        return "X"
    } else if role == game.Oplayer {
        return "O"
    }
    panic("Unknown role")
}

type State string

const(
    Running = "Running"
    Draw = "Draw"
    Win = "Win"
)

type Token string

const(
    X = "X"
    O = "O"
    Empty = " "
)

// Retrieve game's state
func game_state(board string, player string, cell64 uint64) State {
    col := 0
    row := 0
    diag := 0
    rdiag := 0
    n := 3
    cell := int(cell64)
    x := cell/n
    y := cell%n

    for i:=0; i<n; i++ {
        if string(board[x*n + i]) == player { col+=1 }
        if string(board[i*n + y]) == player { row+=1 }
        if string(board[i*n + i]) == player { diag+=1 }
        if string(board[i*n + n-i-1]) == player { rdiag+=1 }
    }
    if row == n || col == n || diag == n || rdiag == n {
        return Win
    }

    empty_cells := strings.Count(board, Empty)
    if empty_cells == 0 {
        return Draw
    }

    return Running
}

func (k Keeper) MakeMove(ctx sdk.Context, msg types.MsgMakeMove) {
    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
    game := k.GetGame(ctx, msg.Id)

    caller_role := caller_role(game, msg)
    assert_valid_move(game, msg, caller_role)

    move := get_move(caller_role, game)
    game.Board = game.Board[:msg.Cell] + move + game.Board[msg.Cell+1:]

    state := game_state(game.Board, move, msg.Cell)
    if state == Win {
        game.Winner = caller_role
        game.Status = types.GameStatus_CLOSED
    } else if state == Draw {
        game.Status = types.GameStatus_CLOSED
    }

    keyPref := types.KeyPrefix(types.GameKey + msg.Id)
    b := k.cdc.MustMarshalBinaryBare(&game)
    store.Set(keyPref, b)
}

func (k Keeper) GetGame(ctx sdk.Context, key string) types.Game {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	var game types.Game
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GameKey + key)), &game)
	return game
}

func (k Keeper) HasGame(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	return store.Has(types.KeyPrefix(types.GameKey + id))
}

func (k Keeper) GetGameOwner(ctx sdk.Context, key string) string {
    return k.GetGame(ctx, key).Creator
}

// DeleteGame deletes a game
func (k Keeper) DeleteGame(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	store.Delete(types.KeyPrefix(types.GameKey + key))
}

func (k Keeper) GetAllGame(ctx sdk.Context) (msgs []types.Game) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.GameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Game
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
