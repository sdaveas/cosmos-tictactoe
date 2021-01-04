"""
Functional test for game accept
"""

import pytest
from common import USER1, USER2, count_games, create_game, accept_game, get_status


def test_accept_game():
    """
    Test running state
    """

    # create a game
    res = create_game(USER1, USER2)
    game_id = count_games() - 1

    # accept a game as provilages; note that ids start from 0
    res = accept_game(USER2, game_id)
    assert res["raw_log"] != "panic"


def test_running_state():
    """
    Test running state
    """
    # create a game
    res = create_game(USER1, USER2)
    game_id = count_games() - 1
    res = accept_game(USER2, game_id)

    # check status
    assert get_status(game_id) == "RUNNING", "Invalid game state"


def test_invalid_accept():
    """
    Test failing accepts
    """

    # create a game
    res = create_game(USER1, USER2)
    game_id = count_games() - 1
    res = accept_game(USER2, game_id)

    # accept the game again; this should fail
    res = accept_game(USER2, game_id)
    assert res["raw_log"] == "panic"

    # try to accept a game that does not exist; this should fail
    res = accept_game(USER2, game_id + 1)
    assert res["raw_log"] == "panic"

    # accept the game as not privilaged user; this should fail
    res = accept_game(USER1, game_id)
    assert res["raw_log"] == "panic"
