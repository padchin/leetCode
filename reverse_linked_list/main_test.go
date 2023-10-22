package main

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	// Создаем связанный список: 1 -> 2 -> 3 -> 4 -> 5
	head := &Node{Data: 1}
	second := &Node{Data: 2}
	third := &Node{Data: 3}
	fourth := &Node{Data: 4}
	fifth := &Node{Data: 5}
	head.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth

	// Обращаем связанный список
	reversedHead := reverseLinkedList(head)

	// Проверяем, что обращенный список имеет правильную структуру: 5 -> 4 -> 3 -> 2 -> 1
	if reversedHead.Data != 5 {
		t.Errorf("Ожидается, что голова обращенного списка будет равна 5, но получено %d", reversedHead.Data)
	}

	if reversedHead.Next.Data != 4 {
		t.Errorf("Ожидается, что следующий элемент после головы обращенного списка будет равен 4, но получено %d", reversedHead.Next.Data)
	}

	if reversedHead.Next.Next.Data != 3 {
		t.Errorf("Ожидается, что следующий элемент после второго элемента обращенного списка будет равен 3, но получено %d", reversedHead.Next.Next.Data)
	}

	if reversedHead.Next.Next.Next.Data != 2 {
		t.Errorf("Ожидается, что следующий элемент после третьего элемента обращенного списка будет равен 2, но получено %d", reversedHead.Next.Next.Next.Data)
	}

	if reversedHead.Next.Next.Next.Next.Data != 1 {
		t.Errorf("Ожидается, что следующий элемент после четвертого элемента обращенного списка будет равен 1, но получено %d", reversedHead.Next.Next.Next.Next.Data)
	}

	// Проверяем, что исходный список не изменился (1 -> 2 -> 3 -> 4 -> 5)
	if head.Data != 1 {
		t.Errorf("Ожидается, что голова исходного списка останется равной 1, но получено %d", head.Data)
	}

	if head.Next.Data != 2 {
		t.Errorf("Ожидается, что следующий элемент после головы исходного списка останется равным 2, но получено %d", head.Next.Data)
	}

	if head.Next.Next.Data != 3 {
		t.Errorf("Ожидается, что следующий элемент после второго элемента исходного списка останется равным 3, но получено %d", head.Next.Next.Data)
	}

	if head.Next.Next.Next.Data != 4 {
		t.Errorf("Ожидается, что следующий элемент после третьего элемента исходного списка останется равным 4, но получено %d", head.Next.Next.Next.Data)
	}

	if head.Next.Next.Next.Next.Data != 5 {
		t.Errorf("Ожидается, что следующий элемент после четвертого элемента исходного списка останется равным 5, но получено %d", head.Next.Next.Next.Next.Data)
	}
}
