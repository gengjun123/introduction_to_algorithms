package chapter2

func insertionSortAsec(list []int) []int {
  for i := 1; i < len(list); i++ {
    key := list[i]

    //insert key into list[0...i-1] which already soted
    j := i - 1
    for j >= 0 && list[j] > key {
      //move list[j] backward
      list[j + 1] = list[j]
      j--
    }
    list[j + 1] = key
  }
  return list
}

func insertionSortDesc(list []int) []int {
  for i := 1; i < len(list); i++ {
    key := list[i]

    //insert key into list[0...i-1] which already soted
    j := i - 1
    for j >= 0 && list[j] < key {
      //move list[j] backward
      list[j + 1] = list[j]
      j--
    }
    list[j + 1] = key
  }
  return list
}
