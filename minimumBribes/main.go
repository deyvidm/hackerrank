// HackerRank :: New Year Chaos
//
// It is New Year's Day and people are in line for the Wonderland rollercoaster ride. Each person wears a sticker indicating their initial position in the queue from to . Any person can bribe the person directly in front of them to swap positions, but they still wear their original sticker. One person can bribe at most two others.
// Determine the minimum number of bribes that took place to get to a given queue order. Print the number of bribes, or, if anyone has bribed more than two people, print Too chaotic.

// q=[1, 2, 3, 4, 5, 6, 7, 8]
// If person 5 bribes person 4 , the queue will look like this: q=[1, 2, 3, 5, 4, 6, 7, 8]
// Only bribe is required. Print 1

// Input Format
// - The first line contains an integer N, the number of people in the queue
// - The second line has N space-separated integers describing the final state of the queue.

// Example 1 Input:
// 5
// 2 1 5 3 4
// Example 1 Output:
// 3

// Example 2 Input
// 5
// 2 5 1 3 4
// Example 2 Output
// Too chaotic

// Too chaotic

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func minimumBribes(q []int32) {
	bribes := int32(0)
	tooChaoticFlag := false
	if len(q) < 1 {
		return
	}
	// i+1   is a person's current position
	// q[i]  is a person's original position (+1 realtive to i, because they're 1-indexed)
	for i := range q {
		// taking diffs in position reveals if this person jumped (bribed) more than the 2 allowed spots
		diff := q[i] - int32(i+1)
		if diff > 2 {
			fmt.Println("Too chaotic")
			tooChaoticFlag = true
			break
		}

		// check q[0:i] for numbers bigger than q[i]
		// this shows us how many folks bribed q[i] to get in front, since all folks with numbers bigger than q[i] come after q[i] in a regular, unbribed queue
		// an optimization we can make :
		//      since folks can only bribe twice, the person behind q[i] (i.e. q[i+1]) can only advance two spots closer to q[0]
		//		which means, q[i+1] can never get any farther than q[i-1]
		//      which means, in    stead of scanning q[0:i] for numbers more than q[i], we can just scan q[q[i]-1:i]
		originalPosition := q[i] - 1

		// this is all just slightly confusing because q[0] != 0, but instead q[0] == 1, (i.e  q[i] = i-1)
		// so to calculate the final farther possible jump, we must adjust the original position by -1 to acount for 1-indexing
		maxJump := originalPosition - 1

		// we use maxJump to construct a slice, so we gotta set a floor at 0 since -1 and below are bad slice params
		if maxJump < 0 {
			maxJump = 0
		}

		// in the case a person gives more bribes than they take, they will advance their position in the queue
		// this can produce an inverted slice interval (constructed by comparing starting vs ending position)
		// 		ex:  q[3:2] will throw a syntax error
		// we make sure to check start and end for sanity before constructing a slice
		start, end := math.Min(float64(maxJump), float64(i)), math.Max(float64(maxJump), float64(i))

		// with this loop, we're simply checking how many of the folks infront of the one at q[i] have a ticket number that comes after q[i]
		// this essentially reveals how many bribes q[i] took, as no one is able to overtake q[i] without bribing them first
		for _, k := range q[int(start):int(end)] {
			if k > q[i] {
				bribes++
			}
		}
	}
	if !tooChaoticFlag {
		fmt.Println(bribes)
	}
}

// duplicate function, but slightly altered to be testable.
// awful practice in production -- will never do.
// i need to keep minimumBribes(...) as it is, in order to keep HackerRank happy
func minimumBribesTestable(q []int32) (bribes int32) {
	if len(q) < 1 {
		return
	}
	for i := range q {
		diff := q[i] - int32(i+1)
		if diff > 2 {
			fmt.Println("Too chaotic")
			break
		}
		originalPosition := q[i] - 1
		maxJump := originalPosition - 1
		if maxJump < 0 {
			maxJump = 0
		}
		start, end := math.Min(float64(maxJump), float64(i)), math.Max(float64(maxJump), float64(i))
		for _, k := range q[int(start):int(end)] {
			if k > q[i] {
				bribes++
			}
		}
	}
	return bribes
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var q []int32

	for i := 0; i < int(n); i++ {
		qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
		checkError(err)
		qItem := int32(qItemTemp)
		q = append(q, qItem)
	}

	minimumBribes(q)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
