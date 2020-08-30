package pkg

type EventHandler interface {
	FireEvent(stockData *StockData)
}
