/*
Package octolog is a very thin and flexible asynchronous logging library.
It is thread-safe and can be used from many different go-routines at once.

	import "github.com/octogo/log"

	func main() {
			defer log.Drain()

			log.Println("this goes to stdout")
			log.Fatal("this goes to stderr")
			log.Panic("this goes to stderr")

			logger := log.NewLogger("my-example-logger")

			logger.Debug("this goes to stdout (teal)")
			logger.Info("this goes to stdout (default)")
			logger.Notice("this goes to stdout (green)")
			logger.Alert("this goes to stderr (magenta)")
			logger.Warning("this goes to stderr (yellow)")
			logger.Error("this goes to stderr (red)")
	}
*/
package log
