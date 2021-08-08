package models

type transactions struct {
	data map[int]int
}

// i.e. go test is green
func NewTransactions(d []int) Transactions {
	data := make(map[int]int)

	for k, v := range d {
		data[k] = v
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
