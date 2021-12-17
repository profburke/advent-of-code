package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const codelen = 12

var masks [codelen]int

func computeMasks() {
	for i := 0; i < codelen; i++ {
		masks[i] = 1 << i
	}
}

func readReport() (codes []int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		code, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		codes = append(codes, int(code))
	}

	return
}

func moreOnes(count, total int) bool {
	return (count > (total - count))
}

func equalOnes(count, total int) bool {
	return (count == (total - count))
}

func computeCounts(items []int) (counts []int) {
	counts = make([]int, codelen)

	for _, item := range items {
		for i := 0; i < codelen; i++ {
			if item&masks[i] > 0 {
				counts[i]++
			}
		}
	}

	return
}

func part1(codes []int) {
	var gamma, epsilon int

	counts := computeCounts(codes)

	for i := 0; i < codelen; i++ {
		if moreOnes(counts[i], len(codes)) {
			gamma |= masks[i]
		} else {
			epsilon |= masks[i]
		}
	}

	fmt.Println(gamma, epsilon, gamma*epsilon)

	return
}

func sliceCopy(s []int) (r []int) {
	r = make([]int, len(s))
	copy(r, s)

	return
}

func part2(codes []int) {
	var orating, crating int

	ocandidates := sliceCopy(codes)
	ccandidates := sliceCopy(codes)

	for i := codelen - 1; i >= 0; i-- {
		var newcandidates []int

		if len(ocandidates) > 1 {
			counts := computeCounts(ocandidates)

			for _, code := range ocandidates {
				if moreOnes(counts[i], len(ocandidates)) ||
					equalOnes(counts[i], len(ocandidates)) {
					if code&masks[i] > 0 {
						newcandidates = append(newcandidates, code)
					}
				} else {
					if !(code&masks[i] > 0) {
						newcandidates = append(newcandidates, code)
					}
				}
			}

			ocandidates = sliceCopy(newcandidates)
		}

		newcandidates = make([]int, 0)

		if len(ccandidates) > 1 {
			counts := computeCounts(ccandidates)

			for _, code := range ccandidates {
				if moreOnes(counts[i], len(ccandidates)) ||
					equalOnes(counts[i], len(ccandidates)) {
					if code&masks[i] == 0 {
						newcandidates = append(newcandidates, code)
					}
				} else {
					if code&masks[i] > 0 {
						newcandidates = append(newcandidates, code)
					}
				}
			}

			ccandidates = sliceCopy(newcandidates)
		}
	}

	orating = ocandidates[0]
	crating = ccandidates[0]

	fmt.Println(orating, crating, orating*crating)
}

func main() {
	computeMasks()
	codes := readReport()
	part1(codes)
	part2(codes)
}

// Local Variables:
// compile-command: "go build"
// End:
