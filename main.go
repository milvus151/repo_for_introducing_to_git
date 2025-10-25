package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // добавляем горутину в группу ожидания
		go worker(i, &wg)
	}

	wg.Wait() // ждем завершения всех горутин
	fmt.Println("Все горутины завершили выполнение")
}
