package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
	 	Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.

 func Same(t1, t2 *tree.Tree) bool {

	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0;i<10;i++ {

		if ( <-c1 != <-c2 )	{
	
				return false
		}
	}

	return true
 }

// Note - Always remember to close channel if using range function, otherwise not
// needed

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
	
	fmt.Println("Result of same tree ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Result of different tree ", Same(tree.New(1), tree.New(3)))

}
