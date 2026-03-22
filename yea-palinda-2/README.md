If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. What is wrong:
Because you give the channel a value it pauses the current goroutine until the value is used,
but since it's on the line after it never happens, causing a deadlock
1. How it was fixed:
Giving the channel a value in a seperate goroutine means that the program can continue to the next line and no deadlock occurs
You could also just give the channel a buffer

##### Buggy Code 2
1. What is wrong:
Main ends before the final number can be printed
2. How it was fixed:
Add a waitgroup which tells main to wait until
the function is finished to exit the program

#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|-|-|-|-|
| What happens if you do X? |  Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function? | I think it might crash | It crashed | It crashed because the producers were still sending to the channel, and you can't send to a closed channel |
| What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?  | I think it will crash again | It crashed | It crashed because one of the producers finished faster than the others, and the rest can't send to a closed channel |
| What happens if you remove the statement `close(ch)` completely?  | It will be stuck infinitely | It ran like normal | I forgot that it will wait until the producers are finished, so while the consumers would continue infinitely if not stopped, main finished and exited the program |
| What happens if you increase the number of consumers from 2 to 4?  | the program finishes faster | It finished faster | More consumers mean more lines can be taken from the channels at the same time |
| Can you be sure that all strings are printed before the program stops?  | No | No | The main program exits before the consumers can print everything because the waitgroup only waits for the producers. |