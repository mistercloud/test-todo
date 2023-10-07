package usecase

import "github.com/mistercloud/test-todo/goapi/internal/entity"

type (
	Todo interface {
		List() []entity.Item
		Add(title string) (entity.Item, error)
		Remove(id entity.ItemId) error
	}
)
