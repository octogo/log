/*
Package octolog is a very thin and flexible asynchronous logging library.
It is thread-safe and can be used from many different go-routines at once.

	import "github.com/octogo/log"

	func main() {
			router := octolog.New()
			defer router.Drain()

			logger := router.NewLogger()
			logger.Info("Hello world!")
	}
*/
package octolog
