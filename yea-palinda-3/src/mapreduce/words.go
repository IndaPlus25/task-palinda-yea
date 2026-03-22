package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	words := strings.Fields(text)

	processes := 6
	total_length := len(words)
	length := total_length / processes

	var wg sync.WaitGroup
	ch := make(chan map[string]int, processes)

	for i := 0; i < total_length; i += length {
		wg.Add(1)

		end := i + length
		if end > total_length {
			end = total_length
		}

		go func(slice []string) {
			defer wg.Done()
			freq := make(map[string]int)

			for _, word := range slice {
				word = strings.ToLower(strings.Trim(word, ",."))
				freq[word]++
			}
			ch <- freq
		}(words[i:end])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	freqs := make(map[string]int)
	for freq := range ch {
		for word, count := range freq {
			freqs[word] += count
		}
	}

	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Milliseconds()

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := os.ReadFile(DataFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
