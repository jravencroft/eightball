package main

import "net/http"
import "flag"
import "fmt"
import "math/rand"
import "strings"
import "time"

var debug = false
var nocache bool = false
var port string = ""
var path string = ""
var messages map[string][]string 

func config() {
	flag.BoolVar(&debug, "debug", false, "display debug information")
	flag.BoolVar(&nocache, "nocache", false, "disable content cache")
	flag.StringVar(&port, "port", "80", "TCP port value for listener")
	flag.StringVar(&path, "path", "/tmp", "Root of content, default is \"tmp\"")
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	lang := getLang(r.Header.Get("Accept-Language"))
	random := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Fprint(w, messages[lang][random.Intn(len(messages[lang]) - 1)])
}

func getLang(header string) string {
	if strings.HasPrefix(header, "en") {
		return "en"
	} else if strings.HasPrefix(header, "it") {
		return "it"
	} else {
		return "en"
	}
}

func fillmessages(count int) map[string][]string{
	out := map[string][]string {}
	out["en"] = []string {
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

	out["it"] = []string {
		"Per quanto posso vedere, sì",
		"È certo",
		"È decisamente così",
		"Molto probabilmente",
		"Le prospettive sono buone",
		"I segni indicano di sì",
		"Senza alcun dubbio",
		"Sì",
		"Sì, senza dubbio",
		"Ci puoi contare",
		"È difficile rispondere, prova di nuovo",
		"Rifai la domanda più tardi",
		"Meglio non risponderti adesso",
		"Non posso predirlo ora",
		"Concentrati e rifai la domanda",
		"Non ci contare",
		"La mia risposta è no",
		"Le mie fonti dicono di no",
		"Le prospettive non sono buone",
		"Molto incerto",
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
