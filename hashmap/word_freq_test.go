package hashmap

type WordsFrequency struct {
	wordFreq map[string]int
}

// leetcode: Constructor
func NewWordsFrequency(book []string) WordsFrequency {
	aMap := make(map[string]int)
	for _, w := range book {
		aMap[w] = aMap[w] + 1
	}
	return WordsFrequency{
		wordFreq: aMap,
	}
}

func (this *WordsFrequency) Get(word string) int {
	return this.wordFreq[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
