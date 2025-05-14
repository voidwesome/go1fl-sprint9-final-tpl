package main

import (
	"testing"
)

// TestGenerateRandomElements проверяет правильность работы функции generateRandomElements.
func TestGenerateRandomElements(t *testing.T) {
	// Проверяем случай, когда размер среза равен нулю
	t.Run("ZeroSize", func(t *testing.T) {
		result := generateRandomElements(0)
		if len(result) != 0 {
			t.Fatalf("Expected empty slice, got %d elements", len(result))
		}
	})

	// Проверяем случай, когда размер среза положительный
	t.Run("PositiveSize", func(t *testing.T) {
		result := generateRandomElements(100)
		if len(result) != 100 {
			t.Fatalf("Expected 100 elements, got %d", len(result))
		}
	})
}

// TestMaximum проверяет правильность работы функции maximum с различными входными данными.
func TestMaximum(t *testing.T) {
	// Тестируем срез с нулевым размером
	t.Run("EmptySlice", func(t *testing.T) {
		result := maximum([]int{})
		if result != 0 {
			t.Fatalf("Expected 0 for empty slice, got %d", result)
		}
	})

	// Тестируем срез, содержащий один элемент
	t.Run("SingleElement", func(t *testing.T) {
		result := maximum([]int{42})
		if result != 42 {
			t.Fatalf("Expected 42, got %d", result)
		}
	})

	// Тестируем срез с несколькими элементами
	t.Run("MultipleElements", func(t *testing.T) {
		data := []int{1, 3, 5, 2, 9, 7}
		result := maximum(data)
		if result != 9 {
			t.Fatalf("Expected 9, got %d", result)
		}
	})
}
