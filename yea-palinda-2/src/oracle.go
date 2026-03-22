// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// TODO: Answer questions.
	go func() {
		for question := range questions {
			go prophecy(question, answers)
		}

	}()
	// TODO: Make prophecies.
	go func() {
		for {
			time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
			go prophecy("", answers)
		}
	}()
	// TODO: Print answers.
	go func() {
		for answer := range answers {
			for _, letter := range answer {
				time.Sleep(time.Duration(50+rand.Intn(200)) * time.Millisecond)
				fmt.Print(string(letter))
			}
			fmt.Print("\n" + prompt)
		}
	}()
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	// Most of these are ai-generated
	reactiveWords := map[string]string{
		"life":       "Ah, life! Every life is like a candle, closer to burning out every second.",
		"death":      "Death is but the next great adventure.",
		"love":       "Love? A powerful magic, yet a heavy burden.",
		"future":     "The future is a mist that hides many paths.",
		"weather":    "The clouds reflect the turmoil of the soul.",
		"money":      "Gold shines, but it cannot light the way.",
		"time":       "Time is a river that flows only one way.",
		"moon":       "A cold eye watching the world dream.",
		"stars":      "Ancient whispers written in a language we have long forgotten.",
		"sea":        "A hungry cradle that gives and takes in equal measure.",
		"fire":       "A fickle friend that warms the hand but devours the house.",
		"wind":       "The breath of the world, carrying secrets from one shore to the next.",
		"mountain":   "A monument to patience, indifferent to the blinking of a human life.",
		"fate":       "A tapestry woven by hands we cannot see and cannot stay.",
		"truth":      "A blade that cuts the one who swings it as deeply as the one it strikes.",
		"silence":    "The loudest sound in a room full of ghosts.",
		"memory":     "A library where the ink fades a little more every time a book is opened.",
		"hope":       "The last ember in a hearth, stubbornly refusing to turn to ash.",
		"fear":       "A shadow that grows taller as the sun begins to set.",
		"regret":     "The echo of a door that was closed too soon.",
		"power":      "A crown of thorns hidden beneath a layer of gold.",
		"friendship": "Two travelers sharing a cloak against a storm.",
		"war":        "A monster that eats the young and leaves the old to starve.",
		"wisdom":     "The scars left behind after the foolishness has burned away.",
		"lies":       "Coins minted in darkness that lose their value in the light.",
		"road":       "A ribbon that ties the beginning to the end.",
		"mirror":     "A thief that steals your youth while you watch.",
		"city":       "A hive of stone where every cell holds a different tragedy.",
		"home":       "The place where the heart finally stops looking for a way out.",
		"shadow":     "A constant companion that never speaks a word.",
		"dream":      "A glimpse into a life you were never meant to lead.",
		"hunger":     "A repeating fire that no amount of wood can never truly quench.",
	}
	// Find the longest word.
	longestWord := ""
	reactiveWord := ""
	found := false
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		reactive, ok := reactiveWords[w]
		if ok && !found {
			reactiveWord = reactive
			if rand.Intn(10) < 5 {
				found = true
			}
		}
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"The dark reunion is near.",
		"Slow is precise, and precise is fast.",
		"Beware of the man who speaks in hands.",
		"The next full moon will be an auspicious night.",
		"What a wonderful world.",
		"The samurai brothers come back stronger after death.",
		"Yin and yang are necessary for balance.",
		"It's dangerous to go alone, take this!",
		"The victor will not always have the last laugh.",
		"Many speak in riddles, as do I.",
		"The answer lies in your heart.",
		"The word friend is subjective for everyone.",
		"May the moonlight guide us.",
		"Praise the sun!",
		"The world's strongest fear no one.",
		"Many seek me, but fates answers are rarely kind.",
		"The moment your fate is read, it is set in stone.",
		"I knew you would come here before you knew it yourself.",
		"Nothing can withstand time.",
	}

	if longestWord == "" {
		answer <- nonsense[rand.Intn(len(nonsense))]
	} else if reactiveWord == "" {
		answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
	} else {
		answer <- reactiveWord
	}
}
