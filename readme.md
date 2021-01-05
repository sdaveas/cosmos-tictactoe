
# tictactoe

**tictactoe** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

In case of the following error,
```
$(go env GOPATH)/bin must be added to your $PATH. See https://golang.org/doc/gopath_code.html#GOPATH
```

try:
```
export PATH=$PATH:$HOME/go/bin
```

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Game

This project creates a blockchain that hosts [tictactoe](https://en.wikipedia.org/wiki/Tic-tac-toe) games.
The interaction with the blockchain is actualized via cli commands are supported. Each command is described thoroughly below.

### Rules
* All state of the game live on-chain. State includes open games, games currently in progress and completed games.
* Any user can submit a transaction to the network to invite others to start a game (i.e. create an open game).
* Other users may submit transactions to accept invitations. When an invitation is accepted, the game starts.
* The roles of “X” and “O” are decided as follows:
    * users' public keys are concatenated and the result is hashed
    * If the first bit of the output is 0, then the guest plays "O" and the host plays "X" and vice versa.
    * “X” has the first move.
* Both users are issuing transactions to the network to make their moves until the game is complete.
* The blockchain supports multiple concurrent games sessions/players.

### Create game

A registered user creates a game and indicates a guest
```
tictactoed tx tictactoe create-game <userX address> --from <userY>
```
Then, you should get a preview of the message you are about to submit:
```
{"body":{"messages":[{"@type":"/sdaveas.tictactoe.tictactoe.MsgCreateGame","creator":"cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st","guest":"cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]:
```
Press `y` to submit the message. There should follow a response that includes various information for your transaction, including the `txhash`, actions and published block height:
```
{"height":"3840","txhash":"E4EBB5A578E3614ACC31F65C0D28C96CE8968F2F67F7E571F22153F29666A598","codespace":"","code":0,"data":"0A0C0A0A43726561746547616D65","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"CreateGame\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"CreateGame"}]}]}],"info":"","gas_wanted":"200000","gas_used":"46697","tx":null,"timestamp":""}
```

You can verify that a game has been created by either querying all games:
```
tictactoed query tictactoe list-game
```
response:
```
- board: '         '
  creator: cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st
  guest: cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k
  id: "0"
  oplayer: NONE
  status: OPEN
  winner: NONE
  xplayer: NONE
pagination:
  next_key: null
  total: "1"
```

or by querying a game by `id`

```
tictactoed query tictactoe show-game 0
```
response:
```
Game:
  board: '         '
  creator: cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st
  guest: cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k
  id: "0"
  oplayer: NONE
  status: OPEN
  winner: NONE
  xplayer: NONE
```

### Game status:
**board** : the current state of the game as a string with size 9. Inside of each non-void cell of the string, **X** or **O** tokens are placed.

**creator** : the public address of the creator

**guest**: the public address of the guest

**id**: the id of the game

**xplayer**: The player that is assigned with *X* token (values from [HOST/GUEST/NONE])

**oplayer**: The player that is assigned with *O* token (values from [HOST/GUEST/NONE])

**status**: The current status of the game (values from [OPEN/RUNNING/CLOSED])

**winner**: The winner of the game (values from [HOST/ GUEST/NONE])

### Accept game
Next, the guest must accept the game in order to start. This is done by the following message:
```
tictactoed tx tictactoe accept-game <id> --from <guest address>
```

Request:
```
{"body":{"messages":[{"@type":"/sdaveas.tictactoe.tictactoe.MsgAcceptGame","guest":"cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k","id":"0"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}
confirm transaction before signing and broadcasting [y/N]:
```
Response:
```
{"height":"4416","txhash":"F75419351DEDA3117761039097497B7B666712A2E31FCEB63C89A7ADC36240AB","codespace":"undefined","code":111222,"data":"","raw_log":"panic","logs":[],"info":"","gas_wanted":"200000","gas_used":"39333","tx":null,"timestamp":""}
```
Now, the status of the game changes as such:
```
Game:
  board: '         '
  creator: cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st
  guest: cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k
  id: "0"
  xplayer: GUEST
  oplayer: CREATOR
  status: RUNNING
  winner: NONE
```
Note that, field `xplayer` and `oplayer` are assigned, and field `status` is set to `RUNNING`. This means that the first player, in our case the `guest`, can start playing.

### Making moves

Each player can make a move in their turn. If a player attempts to move during their opponent's turn, the move will not be rejected.

Players move by submitting the following message:
```
tictactoed tx tictactoe make-move <id> <cell> --from <X player address>```
```
Request:
```
{"body":{"messages":[{"@type":"/sdaveas.tictactoe.tictactoe.MsgMakeMove","caller":"cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k","id":"11","cell":"0"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}
confirm transaction before signing and broadcasting [y/N]:
```
Response:
```
{"height":"4877","txhash":"1ED6E966B93A66CC491797438F42437A9D11212B32995186D748E953EBBE3D99","codespace":"","code":0,"data":"0A0A0A084D616B654D6F7665","raw_log":"[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"MakeMove\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"MakeMove"}]}]}],"info":"","gas_wanted":"200000","gas_used":"44776","tx":null,"timestamp":""}
```
Now, the game status changes as such:
```
Game:
  board: 'X        '
  creator: cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st
  guest: cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k
  id: "0"
  oplayer: CREATOR
  status: RUNNING
  winner: NONE
  xplayer: GUEST
```
Note that, the content of the board includes an **X** token in the first index of the board.

Suppose we have the following game:
```
X | O | X
O | X | O
O | X | X
```
The messages that can potentially produce the above state are the following:

```
tictactoed tx tictactoe make-move 0 0 --from user2 # plays X at (0,0)
tictactoed tx tictactoe make-move 0 1 --from user1 # plays O at (0,1)
tictactoed tx tictactoe make-move 0 2 --from user2 # plays X at (0,2)
tictactoed tx tictactoe make-move 0 3 --from user2 # plays O at (1,0)
tictactoed tx tictactoe make-move 0 4 --from user2 # plays X at (1,1)
tictactoed tx tictactoe make-move 0 5 --from user1 # plays O at (1,2)
tictactoed tx tictactoe make-move 0 7 --from user2 # plays X at (2,1)
tictactoed tx tictactoe make-move 0 6 --from user1 # plays O at (2,0)
tictactoed tx tictactoe make-move 0 8 --from user2 # plays X at (2,2)
```
Note that, at the last move wins the game for the X player. Now, the status of the game is the following:
```
Game:
  board: 'XOXOXOOXX'
  creator: cosmos1gvfssxktnalf73mqpuu6pac4mtzrpw9lyg92st
  guest: cosmos1t8xz0vw857lk270vttuvjtpflnean0endard3k
  id: "0"
  oplayer: CREATOR
  status: RUNNING
  winner: GUEST
  xplayer: GUEST
```
All moves have been registered in **board** field, and the **winner** value has changed to **GUEST**, which indicates that the guest won the game.

## Tests
In the directory
```
tictactoe/tests/
```
functional test can be found. Specifically:
* **test_create_game**: Asserts that a game is created successfully.
* **test_accept_game**: Asserts that a game is accepted successfully only under the correct circumstances.
* **test_play_game**: Asserts that a game is played successfully only under the correct circumstances.

Run the tests by starting a `tictactoe` instance and then:
```
$ cd tests
$ pipenv shell
$ pytest
```

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
