"""
Functional test for gameplay
"""

import pytest
import json
import re
from starport_bindings import starport_bindings
from common import (
    USER1,
    USER2,
    count_games,
    create_game,
    accept_game,
    get_winner_id,
    get_status,
)


def get_addr_of(user_id):
    """
    Return the address from a user id
    """

    return starport_bindings(["get_addr", str(user_id)]).split()[3]


def get_addr_of_role(role, creator_addr, guest_addr):
    """
    Returns the address of a role (creator/guest)
    """

    if role == "CREATOR":
        return creator_addr
    if role == "GUEST":
        return guest_addr
    return None


def get_corresp_role(x_player_addr, creator_addr):
    """
    Returns the id (1,2) of an address
    """

    if x_player_addr == creator_addr:
        return USER1, USER2
    return USER2, USER1


def find_roles(game_id, creator_addr, guest_addr):
    """
    Returns the roles (1, 2) for X and O player, respectively
    """

    game = starport_bindings(["query_game", str(game_id)])
    xplayer = re.search(r"xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search(r"oplayer: (\w+)", game).group().split(" ")[1]
    return (
        get_addr_of_role(xplayer, creator_addr, guest_addr),
        get_addr_of_role(oplayer, creator_addr, guest_addr),
    )


@pytest.fixture(scope="session", autouse=True)
def set_up_game(request):
    """
    Set up a game
    """

    global creator
    global guest
    global creator_addr
    global guest_addr
    global game_id

    creator = USER1
    guest = USER2
    creator_addr = get_addr_of(creator)
    guest_addr = get_addr_of(guest)


def make_move(user_id, game_id, cell):
    """
    Makes a move
    """

    res = starport_bindings(["make_move", str(user_id), str(game_id), str(cell)])
    return json.loads(res.split()[3])


def test_assign_addresses():
    """
    Test playing at a game providing an invalid id
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(guest, game_id)

    # get players' roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    assert (
        x_player_addr is not None and o_player_addr is not None
    ), "Failed to assign players"


def test_assign_addresses():
    """
    Test playing at a game providing an invalid id
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(guest, game_id)

    # move into an invalid game
    res = make_move(creator, game_id + 1, 0)
    assert res["raw_log"] == "panic", "Move in invalid game"
    res = make_move(guest, game_id + 1, 0)
    assert res["raw_log"] == "panic", "Move in invalid game"


def test_assign_tokens():
    """
    Test if game has correct xplayer and oplayer
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(USER2, game_id)

    # get players' roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    assert (
        x_player_addr is not None
        and o_player_addr is not None
        and x_player_addr != o_player_addr
    ), "Failed to assign players"


def test_invalid_moves():
    """
    Test invalid moves;
    1. Play an already allocated cell
    2. Play a cell out of range
    3. Play on opponent's turn
    4. O player plays first
    5. Player plays a move in opponent's turn
    """

    # create game
    create_game(creator, guest)
    game_id = count_games() - 1
    # accept game
    accept_game(USER2, game_id)

    # get roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    x_player, o_player = get_corresp_role(x_player_addr, creator_addr)

    # play 0,0
    res = make_move(x_player, game_id, 0)
    assert res["raw_log"] != "panic", "Move failed"

    # play an already allocated cell
    res = make_move(o_player, game_id, 0)
    assert res["raw_log"] == "panic", "Illegal move"

    # O player plays first
    res = make_move(o_player, game_id, 9)
    assert res["raw_log"] == "panic", "Illegal move"

    # play in opponent's turn
    res = make_move(x_player, game_id, 1)
    assert res["raw_log"] == "panic", "Illegal move"

    # O plays and attempts to play again
    res = make_move(o_player, game_id, 1)
    assert res["raw_log"] != "panic", "Move failed"
    res = make_move(o_player, game_id, 2)
    assert res["raw_log"] == "panic", "Illegal move"


def test_x_wins():
    """
    Test a game in which X wins

    X|O|X
    O|X|O
    O|X|X
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(guest, game_id)

    # get roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    x_player, o_player = get_corresp_role(x_player_addr, creator_addr)

    x_moves = [0, 2, 4, 7, 8]
    o_moves = [1, 3, 5, 6]

    # play the game
    count = 1
    for x_move, y_move in zip(x_moves, o_moves):
        res = make_move(x_player, game_id, x_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for X failed"
        res = make_move(o_player, game_id, y_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for O failed"
        count += 1
    # last move by X
    res = make_move(x_player, game_id, x_moves[-1])

    assert res["raw_log"] != "panic", "Move " + str(count) + " for X failed"
    assert get_winner_id(game_id, x_player, o_player) == x_player, "Invalid winner"
    assert get_status(game_id) == "CLOSED", "Invalid game status"


def test_o_wins():
    """
    Test a game in which O wins

    X| |X
     |X|
    O|O|O
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(guest, game_id)

    # get roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    x_player, o_player = get_corresp_role(x_player_addr, creator_addr)

    x_moves = [0, 2, 4]
    o_moves = [6, 7, 8]

    # play the game
    count = 1
    for x_move, y_move in zip(x_moves, o_moves):
        res = make_move(x_player, game_id, x_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for X failed"
        res = make_move(o_player, game_id, y_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for O failed"
        count += 1
    # last move by X
    res = make_move(x_player, game_id, x_moves[-1])

    assert get_winner_id(game_id, x_player, o_player) == o_player, "Invalid winner"
    assert get_status(game_id) == "CLOSED", "Invalid game status"


def test_draw():
    """
    Test a draw game

    X|O|X
    X|X|O
    O|X|O
    """

    # make game
    create_game(creator, guest)
    game_id = count_games() - 1
    accept_game(guest, game_id)

    # get roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, guest_addr)
    x_player, o_player = get_corresp_role(x_player_addr, creator_addr)

    x_moves = [0, 2, 3, 4, 7]
    o_moves = [1, 5, 6, 8]

    # play the game
    count = 1
    for x_move, y_move in zip(x_moves, o_moves):
        res = make_move(x_player, game_id, x_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for X failed"
        res = make_move(o_player, game_id, y_move)
        assert res["raw_log"] != "panic", "Move " + str(count) + " for O failed"
        count += 1
    # last move by X
    res = make_move(x_player, game_id, x_moves[-1])

    assert get_winner_id(game_id, x_player, o_player) == None, "Invalid winner"
    assert get_status(game_id) == "CLOSED", "Invalid game status"
