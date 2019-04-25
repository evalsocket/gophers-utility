package shutdown_test

import (
	"fmt"
	"syscall"
	"time"

	"github.com/evalsocket/gophers-utility/os/shutdown"
)

func Example() {
	// mock shutdown signall Ctrl + C
	go func() {
		time.Sleep(10 * time.Millisecond)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	shutdown.GracefulStop(func() {
		fmt.Println("shutdown")
	})

	// Output:
	// shutdown
}
