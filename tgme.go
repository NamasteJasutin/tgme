package tgme

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"os/user"
	"path"
	"bufio"
)

type Config struct {
	Token, Recipient string
}

func noCfg() {
	fmt.Println("No CFG file found")
	fmt.Println("Enter your Telegram Bot Token (You can ask the @BotFather)")
	reader := bufio.NewReader(os.Stdin)
	newToken, _ := reader.ReadString('\n')
	newToken = strings.Replace(newToken, "\n", "", -1)
	fmt.Println("Enter your own Telegram Recipient ID (ask @userinfobot)")
	newRecipient, _ := reader.ReadString('\n')
	newRecipient = strings.Replace(newRecipient, "\n", "", -1)
	fmt.Println("Token:", newToken, "\n", "Recipient", newRecipient, "\nAre you sure? [y\\n]")
	saveFile, _ := reader.ReadString('\n')
	saveFile = strings.ToLower(strings.Replace(saveFile, "\n", "", -1))
	
	data := Config{
		Token: newToken,
		Recipient: newRecipient,
	}

	if saveFile == "y" {
		userDir, _ := user.Current()
		cfgFile := path.Join(userDir.HomeDir, `tgme.json`)
		file, _ := json.MarshalIndent(data, "", "  ")
		ioutil.WriteFile(cfgFile, file, 0644)
	} else {
		fmt.Println("Operation cancelled. Please start again.")
		os.Exit(2)
	}
}

func readFile(cfg *Config) {
	userDir, _ := user.Current()
	cfgFile := path.Join(userDir.HomeDir, `tgme.json`)
	f, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		noCfg()
		f, err = ioutil.ReadFile(cfgFile) // Read again after CFG is made
	}

	json := json.Unmarshal(f, &cfg)
	if json != nil {
		fmt.Println("Error:", json)
	}
}

func main() {
	var cfg Config
	readFile(&cfg)
	token := cfg.Token
	sendto := cfg.Recipient
	// Import CL Arguments as message
	args := os.Args[1:]

	if args != nil {
		message := strings.Join(args, " ")
		fmt.Println("sent:\n" + message + "\nTo:" + sendto)
		constURL := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + sendto + "&text=" + message +"&"
		http.Post(constURL, "image/jpeg", nil)
		os.Exit(0)
	}
}
