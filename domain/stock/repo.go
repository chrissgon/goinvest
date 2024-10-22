package stock

type StockSearchRepo interface {
	Run(ID string) (*StockEntity, error)
}
