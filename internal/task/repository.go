package task

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Create(t *Task) error { return r.db.Create(t).Error }
func (r *Repository) GetByID(id uint32) (*Task, error) {
	var x Task
	return &x, r.db.First(&x, id).Error
}
func (r *Repository) List() ([]Task, error) { var xs []Task; return xs, r.db.Find(&xs).Error }
func (r *Repository) ListByUser(uid uint32) ([]Task, error) {
	var xs []Task
	return xs, r.db.Where("user_id = ?", uid).Find(&xs).Error
}
func (r *Repository) Update(t *Task) error   { return r.db.Save(t).Error }
func (r *Repository) Delete(id uint32) error { return r.db.Delete(&Task{}, id).Error }
