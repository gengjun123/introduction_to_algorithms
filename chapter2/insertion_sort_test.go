package chapter2

import (
  "testing"
)

func TestInsertionSortAsec(t *testing.T) {
  instance := []int{5, 3, 9, 2, 0, 1, 4, 6, 8, 7}
  result := insertionSortAsec(instance)
  for i, e := range result {
    if e != i {
      t.Error("failed")
    }
  }
}

func TestInsertionSortDesc(t *testing.T) {
  instance := []int{5, 3, 9, 2, 0, 1, 4, 6, 8, 7}
  result := insertionSortDesc(instance)
  for i, e := range result {
    if i + e != 9 {
      t.Error("failed")
    }
  }
}
