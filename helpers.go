
import "sync"
var (
    mutex sync.Mutex
    counter int
)
func Increment() int {
    mutex.Lock()
    defer mutex.Unlock()
    counter++
    return counter
}

