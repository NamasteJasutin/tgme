# tgme
Send yourself a Telegram message from the CLI

E.G. `brew update | tgme 'Done updating brew'`

# How to use
Download `tgme.go` and build binary with

`go build tgme.go`

then 

`sudo cp tgme /usr/local/bin`

To run, simply enter

`tgme some text here`

in a terminal.

The first run will ask you for a Token and userID

Token = Token for Bot (ask @BotFather)

UserID = Your user id (ask @userinfobot)

Make sure you start the conversation with your bot from Telegram by pressing the start button or messages won't come through.
