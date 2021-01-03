import json
import re
from starport_bindings import starport_bindings

"""
Functional test for game creation
"""

USER1 = 1
USER2 = 2


def count_games():
    res = starport_bindings(["query_games"])
    return int(res.split("total: ")[1].strip('"'))


def create_game(host_id, guest_id):
    res = starport_bindings(["create_game", str(host_id), str(guest_id)])
    return json.loads(res.split()[3])


def accept_game(guest_id, game_id):
    res = starport_bindings(["accept_game", str(guest_id), str(game_id)])
    return json.loads(res.split()[3])


def get_addr_of(user_id):
    return starport_bindings(["get_addr", str(user_id)]).split()[3]


def get_addr_of_role(role, host_addr, guest_addr):
    if role == "CREATOR":
        return host_addr
    if role == "GUEST":
        return guest_addr
    return None


def get_id_of_role(x_player_addr, o_player_addr, creator_addr):
    if x_player_addr == creator_addr:
        return USER1, USER2
    return USER2, USER1


def find_roles(game_id, host_addr, guest_addr):

    game = starport_bindings(["query_game", str(game_id)])
    xplayer = re.search("xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search("oplayer: (\w+)", game).group().split(" ")[1]
    return (
        get_addr_of_role(xplayer, host_addr, guest_addr),
        get_addr_of_role(oplayer, host_addr, guest_addr),
    )


def get_winner_id(game_id, x_player_id, o_player_id):

    game = starport_bindings(["query_game", str(game_id)])
    winner = re.search("winner: (\w+)", game).group().split(" ")[1]
    xplayer = re.search("xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search("oplayer: (\w+)", game).group().split(" ")[1]
    if winner == xplayer:
        return x_player_id
    elif winner == oplayer:
        return o_player_id
    return None


def get_status(game_id):

    game = starport_bindings(["query_game", str(game_id)])
    status = re.search("status: (\w+)", game).group().split(" ")[1]
    return status


def make_move(user_id, game_id, cell):

    res = starport_bindings(["make_move", str(user_id), str(game_id), str(cell)])
    return json.loads(res.split()[3])


def test_create_game():
    """
    Test create
    """

    # count games
    count = count_games()

    # create a game
    res = create_game(USER1, USER2)
    assert res["raw_log"] != "panic"

    # count games again
    count += 1
    assert count_games() == count, "Game not created"
    assert get_status(count-1) == "OPEN", "Invalid game state"


def test_accept_game():
    """
    Test accept
    """

    # create a game
    res = create_game(USER1, USER2)
    assert res["raw_log"] != "panic"

    # get game id
    game_id = count_games() - 1

    # accept a game as provilages; note that ids start from 0
    res = accept_game(USER2, game_id)
    assert res["raw_log"] != "panic"
    assert get_status(game_id) == "RUNNING", "Invalid game state"

    # accept the game again; this should fail
    res = accept_game(USER2, game_id)
    assert res["raw_log"] == "panic"

    # accept the game as not privilaged user; this should fail
    res = accept_game(USER1, game_id)
    assert res["raw_log"] == "panic"


def test_game_X_wins():
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

    x_player, o_player = get_id_of_role(x_player_addr, o_player_addr, creator_addr)

    x_moves = [0, 2, 4, 7, 8]
    o_moves = [1, 3, 5, 6]

    """
    X | O | X
    ---------
    O | X | O
    ---------
    O | X | X
    """

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


test_create_game()
test_accept_game()
test_make_move()
