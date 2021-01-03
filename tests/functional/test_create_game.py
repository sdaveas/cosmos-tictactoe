"""
Functional test for game creation
"""

from common import USER1, USER2, count_games, create_game, get_winner_id, get_status


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
    assert get_status(count - 1) == "OPEN", "Invalid game state"
    winner = get_winner_id(count - 1, USER1, USER2)
    assert winner is None, "Invalid game winner"
