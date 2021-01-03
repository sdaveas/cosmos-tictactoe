"""
Functional test for gameplay
"""

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


def get_addr_of_role(role, host_addr, guest_addr):
    """
    Returns the address of a role (creator/host)
    """

    if role == "CREATOR":
        return host_addr
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


def find_roles(game_id, host_addr, guest_addr):
    """
    Returns the roles (1, 2) for X and O player, respectively
    """

    game = starport_bindings(["query_game", str(game_id)])
    xplayer = re.search(r"xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search(r"oplayer: (\w+)", game).group().split(" ")[1]
    return (
        get_addr_of_role(xplayer, host_addr, guest_addr),
        get_addr_of_role(oplayer, host_addr, guest_addr),
    )


def make_move(user_id, game_id, cell):
    """
    Makes a move
    """

    res = starport_bindings(["make_move", str(user_id), str(game_id), str(cell)])
    return json.loads(res.split()[3])


def test_game_x_wins():
    """
    Test a game in which X wins
    """

    # Assign addresses
    creator_addr = get_addr_of(USER1)
    host_addr = get_addr_of(USER2)

    # create a game
    create_game(USER1, USER2)

    # get game id
    game_id = count_games() - 1

    # accept game
    accept_game(USER2, game_id)

    # get players' roles
    x_player_addr, o_player_addr = find_roles(game_id, creator_addr, host_addr)
    assert (
        x_player_addr is not None and o_player_addr is not None
    ), "Failed to assign players"

    x_player, o_player = get_corresp_role(x_player_addr, creator_addr)

    x_moves = [0, 2, 4, 7, 8]
    o_moves = [1, 3, 5, 6]

    # X|O|X
    # O|X|O
    # O|X|X

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


test_game_x_wins()
