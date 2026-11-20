//go:build windows

package cli

import (
	"context"
	"io"
	"os"

	"golang.org/x/sys/windows"
)

// Poll before reading the raw terminal so a completed loopback callback can
// cancel its outstanding paste prompt without leaving a blocked input goroutine.
func readAuthLoginInput(ctx context.Context, buffer []byte) (int, error) {
	handle := windows.Handle(os.Stdin.Fd())
	for {
		if err := ctx.Err(); err != nil {
			return 0, context.Cause(ctx)
		}
		event, err := windows.WaitForSingleObject(handle, 50)
		if err != nil {
			return os.Stdin.Read(buffer)
		}
		switch event {
		case uint32(windows.WAIT_TIMEOUT):
		case windows.WAIT_OBJECT_0:
			if ctx.Err() != nil {
				return 0, context.Cause(ctx)
			}
			return os.Stdin.Read(buffer)
		default:
			return 0, io.EOF
		}
	}
}
