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


def get_add_of_role(role, host_addr, guest_addr):
    if role == "CREATOR":
        return host_addr
    if role == "GUEST":
        return guest_addr
    return None


def find_roles(game_id, host_addr, guest_addr):

    game = starport_bindings(["query_game", str(game_id)])
    xplayer = re.search("xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search("oplayer: (\w+)", game).group().split(" ")[1]
    return get_add_of_role(xplayer, host_addr, guest_addr), get_add_of_role(
        oplayer, host_addr, guest_addr
    )


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

    # accept the game again; this should fail
    res = accept_game(USER2, game_id)
    assert res["raw_log"] == "panic"

    # accept the game as not privilaged user; this should fail
    res = accept_game(USER1, game_id)
    assert res["raw_log"] == "panic"


test_create_game()
test_accept_game()
