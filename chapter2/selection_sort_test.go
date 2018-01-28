package chapter2

import (
  "testing"
)

func TestSelectionSort(t *testing.T) {
  instance := []int{5, 3, 9, 2, 0, 1, 4, 6, 8, 7}
  result := selectionSort(instance)
  for i, e := range result {
    if e != i {
      t.Error("failed")
    }
  }
}
