package in_memory

import (
	"fmt"
	"sync/atomic"

	"applicationDesignTest/pkg/storage/in_memory/collections"
)

// Попробовал сделать реляционную псевдо in-memory БД - понял, что для данного кейса это через чур и. Поставил себе в pet проекты. Ограничимся псевдо документной БД
// с блокировками на уровне конкретных таблиц с помощью атомиков

const (
	queryCounterStep uint64 = 1
)

type CollectionsList struct {
	Orders           Collection[collections.Order]
	RoomAvailability Collection[collections.RoomAvailability]
}

type storage struct {
	// Глобальный счетчик запросов, уникализирует каждый запрос (если отдельно) либо транзакцию
	queryIDCounter atomic.Uint64

	collections *CollectionsList
}

func newStorage() *storage {
	fmt.Println("Creating new storage!")

	return &storage{
		collections: &CollectionsList{
			Orders:           newCollection[collections.Order]("orders"),
			RoomAvailability: newCollection[collections.RoomAvailability]("room_available"),
		},
	}
}

func (s *storage) newQueryID() uint64 {
	// Пока просто инкрементим счетчик для получения ИД (uint64 довольно-таки большой, для наших целей хватит)
	// Можно рассмотреть вариант зациклить его, при достижении максимального значения
	return s.queryIDCounter.Add(queryCounterStep)
}
