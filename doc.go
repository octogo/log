/*
Package log is a very thin and flexible asynchronous logging library.
It is thread-safe and can be used from many different goroutines
simulatenously.

	import "github.com/octogo/log"

	func main() {
			defer log.Close()

			log.Println("this goes to stdout")
			log.Fatal("this goes to stderr")
			log.Panic("this goes to stderr")

			// create a logger and log some demo lines
			logger := log.NewLogger("my-example-logger")
			logger.Debug("this goes to stdout (teal)")
			logger.Info("this goes to stdout (default)")
			logger.Notice("this goes to stdout (green)")
			logger.Alert("this goes to stderr (magenta)")
			logger.Warning("this goes to stderr (yellow)")
			logger.Error("this goes to stderr (red)")

			// create a child logger and log some more demo lines
			childLogger := logger.NewLogger("child")
			childLogger.Debug("this goes to stdout (teal)")
			childLogger.Info("this goes to stdout (default)")
			childLogger.Notice("this goes to stdout (green)")
			childLogger.Alert("this goes to stderr (magenta)")
			childLogger.Warning("this goes to stderr (yellow)")
			childLogger.Error("this goes to stderr (red)")
	}
*/
package log
