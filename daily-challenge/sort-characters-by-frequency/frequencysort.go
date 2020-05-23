package frequencysort

// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/

type cdata struct {
	freq int
	c    string
}

// V3 (coming soon)-- reduce runtime complexity using []cdata and map[string]*cdata
//   (where *cdata corresponds to an element in a slice) to reduce runtime complexity

// V2 (not finalized yet)-- uses always ordered array of structs

func frequencySort(s string) string {
	// chararray will always contain sorted (decreasing) cdata structs by frequency
	chararray := []cdata{}
	var charfound bool
	for _, r := range s {
		currchar := string(r)
		charfound = false
		for i, char := range chararray {
			if char.c == currchar {
				char.freq++
				charfound = true
				// if frequency of index-1 char is now less, swap the two to mantain sort
				tmpi := i
				for tmpi > 0 && chararray[tmpi-1].freq < char.freq {
					tmp := chararray[tmpi-1]
					chararray[tmpi-1] = char
					char = tmp
					tmpi--
				}
				// if i > 0 {
				// 	for swapindex := i; swapindex >= 0; swapindex-- {
				// 		if chararray[i-1].freq < char.freq {
				// 		}
				// 	}
				// }
			}
		}
		if charfound == false {
			chararray = append(chararray, cdata{freq: 1, c: currchar})
		}
	}
	return "s"
}

// v1 solution O(n^2)

// chars := make(map[string]int)
// for _, char := range s {
// 	if _, ok := chars[string(char)]; ok {
// 		chars[string(char)]++
// 	} else {
// 		chars[string(char)] = 1
// 	}
// }
// // extract frequencies into array and sort
// freqs := []int{}
// for _, freq := range chars {
// 	freqs = append(freqs, freq)
// }
// // sort frequencies
// sort.Ints(freqs)
// return "s"
// //for each frequency
// }
