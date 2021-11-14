package user

import (
	"merchant-service/entity"
	"merchant-service/query"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindByID(ID string) (entity.User, error)
	DeleteByID(ID string) (string, error)
	UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	CreateOutletUser(Outlet entity.Outlet) (entity.Outlet, error)
	FindOutletUserByID(ID string) (entity.Outlet, error)
	ShowAllOutletUser() ([]entity.Outlet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) UpdateByID(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CreateOutletUser(outlet entity.Outlet) (entity.Outlet, error) {
	qry := query.CreateOutletbyUser

	err := r.db.Raw(qry,
		outlet.Id,
		outlet.OutletName,
		outlet.Picture,
		outlet.UserId,
		outlet.CreatedAt,
		outlet.UpdatedAt).Scan(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) FindOutletUserByID(ID string) (entity.Outlet, error) {
	var outlet entity.Outlet

	qry := query.FindOutletUserByID

	err := r.db.Raw(qry, ID).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}

func (r *repository) ShowAllOutletUser() ([]entity.Outlet, error) {
	var outlet []entity.Outlet

	qry := query.GetAllOutlets

	err := r.db.Raw(qry).Scan(&outlet).Error

	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
