package infra

import (
	"fmt"

	"github.com/chrissgon/goinvest/entity"
)

type PersistMemory[T any] struct {
	list map[string]T
}

func NewPersistMemory[T any]() entity.IPersist[T] {
	return &PersistMemory[T]{
		list: map[string]T{},
	}
}

func (p *PersistMemory[T]) Add(ID string, entity T) error {
	fmt.Println(p.list)
	p.list[ID] = entity
	return nil
}

func (p *PersistMemory[T]) Get(ID string) T {
	fmt.Println(p.list)
	return p.list[ID]
}

// // Add implements entity.IPersistStock.
// func (p *PersistMemory[T]) Add(stock.StockEntity) error {
// 	panic("unimplemented")
// }

// // Get implements entity.IPersistStock.
// func (p *PersistMemory[T]) Get(ID string) stock.StockEntity {
// 	panic("unimplemented")
// }

// func (s *StockPersistMemory) Add(stk stock.StockEntity) error {
// 	s.stocks[stk.ID] = stk
// 	return nil
// }

// func (s *StockPersistMemory) Get(ID string) stock.StockEntity {
// 	return s.stocks[ID]
// }
