# tgme
Send yourself a Telegram message from the CLI
Started off in Python, migrated to Go.
Just a simple message sending tool for use in CLI
E.G. `brew update | tgme 'Done updating brew'`

# How to use
Download the go file and compile with

`go build tgme.go`

then 

`sudo cp tgme /usr/local/bin`

To run, simply enter

`tgme some text here`

in a terminal.

The first run will ask you for a Token and userID
Token = Token for Bot (ask @BotFather)
UserID = Your user id (ask @userinfobot)
Make sure you start the conversation with your bot by pressing the start button or you will messages won't come through.
