package main

import "testing"

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(10)

	if len(data) != 10 {
		t.Errorf("ожидалась длина 10, получили %d", len(data))
	}
}

func TestGenerateRandomElementsZeroSize(t *testing.T) {
	data := generateRandomElements(0)

	if len(data) != 0 {
		t.Errorf("ожидался пустой слайс, получили длину %d", len(data))
	}
}

func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	data := generateRandomElements(-1)

	if len(data) != 0 {
		t.Errorf("ожидался пустой слайс, получили длину %d", len(data))
	}
}

func TestMaximum(t *testing.T) {
	data := []int{3, 15, 7, 1, 9}

	got := maximum(data)

	if got != 15 {
		t.Errorf("ожидалось 15, получили %d", got)
	}
}

func TestMaximumEmptySlice(t *testing.T) {
	data := []int{}

	got := maximum(data)

	if got != 0 {
		t.Errorf("ожидалось 0, получили %d", got)
	}
}

func TestMaximumOneElement(t *testing.T) {
	data := []int{5}

	got := maximum(data)

	if got != 5 {
		t.Errorf("ожидалось 5, получили %d", got)
	}
}

func TestMaximumSameElements(t *testing.T) {
	data := []int{4, 4, 4}

	got := maximum(data)

	if got != 4 {
		t.Errorf("ожидалось 4, получили %d", got)
	}
}
