package pkg

type StocksRetriever interface {
	GetStockData(symbol string, interval int) StockData
}
