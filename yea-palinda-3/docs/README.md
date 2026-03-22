#### What happens if you remove the `go-command` from the `Seek` call in the `main` function?
Because it is no longer running the multiple of the seek function at the same time, the order of who sends and recieves will be the exact same everytime.

#### What happens if you switch the declaration `wg := new(sync.WaitGroup)` to `var wg sync.WaitGroup` and the parameter `wg *sync.WaitGroup` to `wg sync.WaitGroup`?
If you switch the decleration and parameter like that then you are no longer sending in a reference to the same waitgroup but instead making a copy for each goroutine, which means that every wg.Done() doesn't actually affect the waitgroup in main, so it gets stuck on wg.Wait() and it turns into a deadlock.

#### What happens if you remove the buffer on the channel match?
The select will cause the last persons goroutine to wait for a reciever forever because there is no buffer and it will enter a deadlock.

#### What happens if you remove the default-case from the case-statement in the `main` function?
Nothing different will happen if you don't change anything else but if you make the people slice have an even number of people then the program will enter into a deadlock because there is no case for no extra message being unsent.


|Variant       | Runtime (ms) |
| ------------ | ------------:|
| singleworker |          676 |
| mapreduce    |          500 |