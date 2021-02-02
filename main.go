package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
)


func main() {
	token := "1658559528:AAEq3Pd2XukSeFj0bVpPDE_p4SaPMXuhf7Q"
	sendto := "8550799"
	args := os.Args[1:]
	message := strings.Join(args, " ")
	fmt.Println("Sending:\n" + message + "\nTo:" + sendto)
	constURL := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + sendto + "&text=" + message +"&"
	http.Post(constURL, "image/jpeg", nil)
}
