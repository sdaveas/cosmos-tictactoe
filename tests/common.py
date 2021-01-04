"""
Common functions for funtional tests
"""

import re
import json
from starport_bindings import starport_bindings

USER1 = 1
USER2 = 2


def count_games():
    """
    Retrieve the total number of active and inactive games
    """

    res = starport_bindings(["query_games"])
    return int(res.split("total: ")[1].strip('"'))


def get_winner_id(game_id, x_player_id, o_player_id):
    """
    Retrieve the id of the winner
    """

    game = starport_bindings(["query_game", str(game_id)])
    winner = re.search(r"winner: (\w+)", game).group().split(" ")[1]
    if winner == "NONE":
        return None

    xplayer = re.search(r"xplayer: (\w+)", game).group().split(" ")[1]
    oplayer = re.search(r"oplayer: (\w+)", game).group().split(" ")[1]
    if winner == xplayer:
        return x_player_id
    if winner == oplayer:
        return o_player_id
    return None


def get_status(game_id):
    """
    Retrieve the status of the game
    """

    game = starport_bindings(["query_game", str(game_id)])
    status = re.search(r"status: (\w+)", game).group().split(" ")[1]
    return status


def create_game(host_id, guest_id):
    """
    Create a game
    """

    res = starport_bindings(["create_game", str(host_id), str(guest_id)])
    return json.loads(res.split()[3])


def accept_game(guest_id, game_id):
    """
    Accept a game
    """

    res = starport_bindings(["accept_game", str(guest_id), str(game_id)])
    return json.loads(res.split()[3])
