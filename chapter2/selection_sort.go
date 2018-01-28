package chapter2

func selectionSort(list []int) []int {
  n := len(list)
  //loop for inserting the ith smallest number into list[i]
  for i := 0; i < n - 1; i++ {

    //begin with the last number list[n - 1], if list[j] smaller than list[j - 1], then swip them
    for j := n - 1; j - 1 >= i; j-- {
      if list[j] < list[j - 1] {
        list[j], list[j - 1] = list[j - 1], list[j]
      }
    }
  }

  return list
}
