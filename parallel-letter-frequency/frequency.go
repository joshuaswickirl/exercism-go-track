package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns
// this data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates the total frequency of each rune in a
// given array of strings and returns this data as a FreqMap.
func ConcurrentFrequency(input []string) FreqMap {
	freqMap := FreqMap{}
	channel := make(chan FreqMap, 10)
	for _, s := range input {
		go func(s string) {
			channel <- Frequency(s) // send partial freqmap
		}(s)
	}
	for range input {
		// combine freqmaps
		for k, v := range <-channel {
			freqMap[k] += v
		}
	}
	return freqMap
}
