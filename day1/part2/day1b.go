//
// --- Part Two ---
//
// Your analysis only confirmed what everyone feared: the two lists of location IDs are indeed very different.
//
// Or are they?
//
// The Historians can't agree on which group made the mistakes or how to read most of the Chief's handwriting,
// but in the commotion you notice an interesting detail: a lot of location IDs appear in both lists!
// Maybe the other numbers aren't location IDs at all but rather misinterpreted handwriting.
//
// This time, you'll need to figure out exactly how often each number from the left list appears in the right list.
// Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
//
// Here are the same example lists again:
//
// 3   4
// 4   3
// 2   5
// 1   3
// 3   9
// 3   3
//
// For these example lists, here is the process of finding the similarity score:
//
//     The first number in the left list is 3. It appears in the right list three times, so the similarity score increases by 3 * 3 = 9.
//     The second number in the left list is 4. It appears in the right list once, so the similarity score increases by 4 * 1 = 4.
//     The third number in the left list is 2. It does not appear in the right list, so the similarity score does not increase (2 * 0 = 0).
//     The fourth number, 1, also does not appear in the right list.
//     The fifth number, 3, appears in the right list three times; the similarity score increases by 9.
//     The last number, 3, appears in the right list three times; the similarity score again increases by 9.
//
// So, for these example lists, the similarity score at the end of this process is 31 (9 + 4 + 0 + 0 + 9 + 9).
//
// Once again consider your left and right lists. What is their similarity score?

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_file(file_path string) ([]int, map[int]int, error) {

  var list1 []int
  var lmap =  make(map[int]int)
  file, err := os.Open(file_path)
  if err != nil {
    return nil, nil, fmt.Errorf("Cannot read file")
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    var line string = scanner.Text()
    arr := strings.Fields(line)
    if len(arr) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		item1, err := strconv.Atoi(arr[0])
		if err != nil {
			return nil, nil, fmt.Errorf("cannot convert to int: %v", err)
		}

		item2, err := strconv.Atoi(arr[1])
		if err != nil {
			return nil, nil, fmt.Errorf("cannot convert to int: %v", err)
		}
    list1 = append(list1, item1)
      
    lmap[item2] ++

  }
  return list1, lmap, nil
}

func calculate_distance(list1 []int, list2 map[int]int) int {
  sort.Ints(list1)
  var total int = 0
  for i:=0; i< len(list1); i++{
    total += list2[list1[i]] * list1[i]
  }
  return total
}

func main() {
  list1, list2, err := read_file("day1.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  total := calculate_distance(list1, list2)
  fmt.Println("Total Distance is: " ,total)
}
