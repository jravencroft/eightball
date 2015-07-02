package main

import "net/http"
import "flag"
import "fmt"
import "math/rand"
import "time"

var debug = false
var nocache bool = false
var port string = ""
var path string = ""
var messages []string 

type Object struct {
	key   string
	value []byte
}

func config() {
	flag.BoolVar(&debug, "debug", false, "display debug information")
	flag.BoolVar(&nocache, "nocache", false, "disable content cache")
	flag.StringVar(&port, "port", "80", "TCP port value for listener")
	flag.StringVar(&path, "path", "/tmp", "Root of content, default is \"tmp\"")
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Fprint(w, messages[random.Intn(len(messages) - 1)])
}

func fillmessages(count int) []string{
	var out = []string {
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	return out
}

func debuginfo() {
	fmt.Println("Configuration:")
	fmt.Println("     Listener port:", port)
	fmt.Println("         Path root:", path)
	fmt.Println("  Caching disabled:", nocache)
	fmt.Println("          Messages:", len(messages));
}

func main() {
	messages = fillmessages(20)
	config();
	if(debug == true) {
		debuginfo()
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
