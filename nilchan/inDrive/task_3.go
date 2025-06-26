package inDrive

import (
	"context"
	"fmt"
	"time"
)

// что выведет
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	time.Sleep(2950 * time.Millisecond)
	doDbRequest(ctx)
}

func doDbRequest(ctx context.Context) {
	newCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-newCtx.Done():
		fmt.Println("timeout")
	case <-timer.C:
		fmt.Println("Request Done")
	}
}

// timeout Потому что newCtx завершится через 50 мс, а timer — только через 1 секунду.

/*
Что происходит:
В main() создается ctx с таймаутом 3 секунды.
Программа "спит" 2.95 секунды, после чего вызывает doDbRequest(ctx).
В doDbRequest(ctx) создается newCtx, чей дедлайн — минимальный из родительского (3с) и текущего (10с). Поэтому newCtx тоже истекает через ~50 мс.
Запускается time.NewTimer(1 * time.Second) — он сработает через 1 секунду.

В select:
Если newCtx.Done() сработает раньше — будет выведено "timeout".
Если раньше сработает timer.C — будет "Request Done".

Тайминги:
В момент входа в doDbRequest, родительский ctx уже почти истёк (осталось ~50 мс).
newCtx унаследует этот дедлайн → сработает через ~50 мс.
timer.C сработает через 1 секунду.
*/
