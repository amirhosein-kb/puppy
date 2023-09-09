package puppy

import (
	"github.com/amirhosein-kb/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() string{
	return "i am from version 1.1.0"
}
