package repo

import "ucenter/internal/model"

type MemberRepo interface {
	Save(m *model.Member) error
	FindByPhone(phone string) (*model.Member, error)
	FindByUsername(username string) (*model.Member, error)
	IncrLoginCount(id int64, incr int) error
}
