package dao

import (
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemeberDao struct {
	db *gorm.DB
}

func NewMemberDao(db *gorm.DB) *MemeberDao {
	return &MemeberDao{db: db}
}

func (dao *MemeberDao) Save(m *model.Member) error {
	return dao.db.Create(m).Error
}

func (dao *MemeberDao) FindByPhone(phone string) (*model.Member, error) {
	member := &model.Member{}
	if err := dao.db.Model(&model.Member{}).Where("mobile_phone = ?", phone).First(member).Error; err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return member, nil
}

func (dao *MemeberDao) FindByUsername(username string) (*model.Member, error) {
	mem := &model.Member{}
	if err := dao.db.Where("username = ?", username).First(mem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return mem, nil
}

func (dao *MemeberDao) IncrLoginCount(id int64, incr int) error {
	err := dao.db.Exec("update member set login_count=login_count+? where id = ?", incr, id).Error
	return err
}
