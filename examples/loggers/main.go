package main

import "github.com/octogo/log"

func main() {
	log.Init()
	logger := log.New("main", nil)
	logger.Info("Parent")

	child1 := logger.NewLogger("child-1")
	child1.Info("Child 1")

	child2 := logger.NewLogger("child-2")
	child2.Info("Child 2")

	grandChild1 := child1.NewLogger("grand-child-1")
	grandChild1.Info("Grandchild 1")

	grandChild2 := child2.NewLogger("grand-child-2")
	grandChild2.Info("Grandchild 2")
}
