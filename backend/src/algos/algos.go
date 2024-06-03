package algos

import (
	"math/rand"
)

func Algo(arr []int, edgeIndex, lastIndex int) []int {
	for i := edgeIndex; i > lastIndex; i-- {
		rand_index := rand.Intn(i + 1) // i+1 из-за того, что rand.Intn возвращает псевдослучайное число из полуинтервала [0, n)
		arr[rand_index], arr[i] = arr[i], arr[rand_index]
	}

	return arr
}
