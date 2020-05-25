package frequencysort

// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/537/week-4-may-22nd-may-28th/

type cdata struct {
	freq int
	c    string
}

// V3 (coming soon)-- reduce runtime complexity using []cdata and map[string]*cdata
//   (where *cdata corresponds to an element in a slice) to reduce runtime complexity

func frequencySort(s string) string {
	// chararray will always contain sorted (decreasing) cdata structs by frequency
	chararray := []cdata{}
	chararraypmap := map[string]*cdata{}
	for _, r := range s {
		currchar := string(r)
		if val, ok := chararraypmap[currchar]; ok {
			// increment counter of character
			val.freq++
			// TODO: while previous element in charraray has less frequency, swap elements

		} else {
			// if character hasn't been encountered
			newchar := cdata{c: currchar, freq: 1}
			chararray = append(chararray, newchar)
			chararraypmap[currchar] = &chararray[len(chararray)-1]
		}
	}
	// assemble return string
	var rstring string
	for _, cstruct := range chararray {
		for i := 0; i < cstruct.freq; i++ {
			rstring += cstruct.c
		}
	}
	return rstring
}

// V2 (not finalized yet)-- uses always ordered array of structs

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
