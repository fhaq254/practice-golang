// https://tutorialedge.net/golang/the-go-init-function/

// init()
// It’s incredibly important to note that you cannot rely upon the order of execution of your init() functions.
// It’s instead better to focus on writing your systems in such a way that the order does not matter.

package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	} else {
		fmt.Println("It's always the truth.")
	}
}

// you’ll see that AnswerToLife() will run before our init() function
// as our WhatIsThe variable is declared before our init() function is called.
