package hotdog

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var (
	whatHeSaid = []string{
		"My chiropractor wanted to massage my glutes with a large spoon",
		"Miley Cyrus is generally accepted as attractive",
		"Bieber is a decent looking guy",
		"You gotta get some scented candles bro",
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(whatHeSaid))

	fmt.Fprintf(w, fmt.Sprintf(`
		<!doctype html>
		<html>
			<head>
				<title>What Did Alex Say Today?</title>
				<style>
					div {
						text-align: center; 
						font-family: Arial,sans-serif; 
						font-weight: bold; 
						font-size: 52t pt
					}
				</style>
			</head>
			<body>
				<div>%s</div>
			</body>
		</html>`, whatHeSaid[index]))
}

func init() {
	http.HandleFunc("/", handler)
}
