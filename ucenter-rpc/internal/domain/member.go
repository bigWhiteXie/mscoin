package domain

import (
	"coin-common/tools"
	"gorm.io/gorm"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	repo.MemberRepo
}

func NewMemberDomain(db *gorm.DB) *MemberDomain {
	return &MemberDomain{
		dao.NewMemberDao(db),
	}
}

func (d *MemberDomain) Register(username string, password string, phone string, country string, promotion string, partner string) error {
	//构建用户信息,生成盐并进行加密
	mem := &model.Member{
		Username:      username,
		MobilePhone:   phone,
		Country:       country,
		PromotionCode: promotion,
	}
	mem.FillSuperPartner(partner)
	mem.Salt, mem.Password = tools.Encode(password, nil)
	if err := tools.Default(mem); err != nil {
		return err
	}
	if err := d.Save(mem); err != nil {
		return err
	}
	return nil
}

func (d *MemberDomain) UpdateLoginCount(id int64, incr int) {
	d.MemberRepo.IncrLoginCount(id, incr)
}
