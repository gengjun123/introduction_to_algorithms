package chapter2

import (
  "strconv"
  "math"
)

func mergeSort(list []int) []int {
  n := len(list)

  //list with only one element is sorted certainly
  if n <= 1 {
    return list
  }

  //divide into [0 : n / 2), [n / 2 : n)
  left := mergeSort(list[0 : n / 2])
  right := mergeSort(list[n / 2 : n])
  return merge(left, right)
}

//merge left and right slices to be sorted into left[0 : len(left) + len(right)]
//left and right share the same base array
func merge(left, right []int) []int {
  //allocate two new slices to storage left and right, because we will use old slices to storage the merged result
  leftNew := make([]int, len(left), len(left) + 1)
  rightNew := make([]int, len(right), len(right) + 1)
  copy(leftNew, left)
  copy(rightNew, right)
  //append the MAXInt to the leftNew/rightNew end
  //making sure that looping len(left) + len(right) times will put all elements into the merged result
  leftNew = append(leftNew, getMAXInt())
  rightNew = append(rightNew, getMAXInt())

  //extend left lenght to len(left) + len(right) to storage the merged result
  left = left[ : len(left) + len(right)]

  li, ri := 0, 0
  for i := 0; i < len(left); i++ {
    if leftNew[li] < rightNew[ri] {
      left[i] = leftNew[li]
      li++
    } else {
      left[i] = rightNew[ri]
      ri++
    }
  }
  return left
}

func getMAXInt() int {
  if strconv.IntSize == 32 {
    return math.MaxInt32
  } else {
    return math.MaxInt64
  }
}
