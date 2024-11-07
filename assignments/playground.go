package assignments

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	//c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func Playground() {
	//var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{} // to implement interface with pointer reviever methods, you need to dereference

	valueCounter := Counter{}

	//myStringer = pointerCounter    // ok
	//myStringer = valueCounter      // ok
	myIncrementer = pointerCounter // ok
	myIncrementer.Increment()

	// myIncrementer = valueCounter // compile-time error!

	valueCounter.Increment() // go automatically dereferences
	valueCounter.Increment()
	valueCounter.Increment()
	fmt.Println(valueCounter.String())

	fmt.Println(myIncrementer)
}
