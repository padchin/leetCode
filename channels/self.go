package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	// Функция, которую нужно выполнить только один раз
	initialize := func() {
		fmt.Println("Инициализация выполняется только один раз")
	}
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(5)
	// Несколько горутин, которые вызывают функцию initialize через sync.Once
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer waitGroup.Done()
			once.Do(initialize) // Эта функция будет выполнена только один раз
			fmt.Println(i)
		}(i)
	}

	// Ожидание завершения всех горутин

	waitGroup.Wait()

	fmt.Println("Горутины завершили выполнение")
}
