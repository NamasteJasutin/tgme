# tgme
Send yourself a Telegram message from the CLI

E.G. `brew update | tgme 'Done updating brew'`

# How to use
Clone the repo and build using

`go build`

then 

`sudo cp tgme /usr/local/bin`

To run, simply enter

`tgme some text here`

in a terminal.

On the first run, you'll be asked for a Token for the bot to send from and a userID to send messages to by default.

Token = Token for Bot (ask @BotFather)

UserID = Your user id (ask @userinfobot)

Make sure you start the conversation with your bot from Telegram by pressing the start button before trying to send messages from your bot, or messages won't come through!
