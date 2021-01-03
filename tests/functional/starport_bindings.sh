#!/bin/bash

set -e

export PATH=$PATH:$HOME/go/bin
APPLICATION="tictactoed"


# get address of a user
get_addr()
{
    $APPLICATION keys list | grep $1 -A 2 | tail -n 1 | cut -d" " -f 4;
}

# create a game from a host and invite guest
create_game()
{
    host=$1
    guest=$2
    yes | $APPLICATION tx ${APPLICATION::-1} create-game $(get_addr $guest) --from $host
}

# accept a game
accept_game()
{
    guest=$1
    id=$2
    yes | $APPLICATION tx ${APPLICATION::-1} accept-game $id --from $guest
}

# accept a game
make_move()
{
    player=$1
    id=$2
    cell=$3
    yes | $APPLICATION tx ${APPLICATION::-1} make-move $id $cell --from $player
}

query_game()
{
    id=$1
    $APPLICATION query ${APPLICATION::-1} show-game $id
}

query_games()
{
    $APPLICATION query ${APPLICATION::-1} list-game
}

echo cmd: $1 $2

if [ $1 == "get_addr" ]
then
    user=user$2
    get_addr $user
elif [ $1 == "create_game" ]
then
    host=user$2
    guest=user$3
    create_game $host $guest
elif [ $1 == "accept_game" ]
then
    guest=user$2
    id=$3
    accept_game $guest $id
elif [ $1 == "make_move" ]
then
    player=user$2
    id=$3
    cell=$4
    make_move $player $id $cell
elif [ $1 == "query_games" ]
then
    query_games
elif [ $1 == "query_game" ]
then
    id=$2
    query_game $id
fi
