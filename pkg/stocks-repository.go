package pkg

type StocksRepostory interface {
	Add(string)
	Delete(string)
	List() []string
}
