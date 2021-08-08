package models

type transactions struct {
	data map[int]int
}

// i.e. go test is green
func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for i, v := range d {
		data[i] = v
	}

	return &transactions{data: data}
}

// i.e. go test is green
func (t *transactions) Get(idx int) int {
	v, ok := t.data[idx]
	if ok {
		return v
	}
	return 0
}

func (t *transactions) GetTotal() int {
	total := 0

	for _, v := range t.data {
		total += v
	}

	return total
}

func (t *transactions) GetTotalWithinRange(i, j int) int {
	total := 0

	for ; i <= j && i < len(t.data); i++ {
		total += t.data[i]
	}

	return total
}
