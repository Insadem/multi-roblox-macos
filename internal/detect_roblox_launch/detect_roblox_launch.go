package detect_roblox_launch

import (
	"time"
)

// TODO: if roblox not opened after add new instance more than once - display zenity message to switch to fallback

// channel resolves if launched any instance
func ShouldLaunchIn(duration time.Duration) <-chan bool {
	resultChan := make(chan bool)
	defer close(resultChan)

	go func() {
		// ctx, _ := context.WithTimeout(context.Background(), duration)

		// select {
		// case <-ctx.Done():
		// 	resultChan <- false;
		// case // check every 50 ms if roblox is opened (alr, we need isolated package to track all roblox instance and when some closes or opens)
		// }
	}()

	return resultChan
}
