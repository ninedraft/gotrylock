# gotrylock

TryLock mutex for Golang

Example:

```go
import(
    "fmt"

    "github.com/ninedraft/gotrylock"
)

func main() {
    tlm := &gotrylock.TryMutex{}
	tlm.Lock()
	log.Println("main start!")
	go func() {
		for !tlm.TryLock(time.Second) {
			log.Println("timeout!")
		}
		log.Println("goroutine start!")
		time.Sleep(4 * time.Second)
		log.Println("goroutine stop!")
		tlm.Unlock()
	}()
	time.Sleep(4 * time.Second)
	log.Println("main stop!")
	tlm.Unlock()
	tlm.Lock()
	log.Println("main start!")
	log.Println("main stop!")
	tlm.Unlock()
	log.Println("Yay!")
}
```

Output:
```
2016/10/26 11:46:47 main start!
2016/10/26 11:46:48 timeout!
2016/10/26 11:46:49 timeout!
2016/10/26 11:46:50 timeout!
2016/10/26 11:46:51 main stop!
2016/10/26 11:46:51 goroutine start!
2016/10/26 11:46:55 goroutine stop!
2016/10/26 11:46:55 main start!
2016/10/26 11:46:55 main stop!
2016/10/26 11:46:55 Yay!
```
