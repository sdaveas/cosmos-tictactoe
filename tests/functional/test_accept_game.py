"""
Functional test for game accept
"""

from common import USER1, USER2, count_games, create_game, accept_game, get_status


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
