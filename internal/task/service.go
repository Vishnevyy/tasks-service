package task

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service { return &Service{repo: r} }

func (s *Service) CreateTask(t Task) (*Task, error) {
	if err := s.repo.Create(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Service) GetTask(id uint32) (*Task, error) { return s.repo.GetByID(id) }

func (s *Service) ListTasks() ([]Task, error) { return s.repo.List() }

func (s *Service) ListTasksByUser(uid uint32) ([]Task, error) { return s.repo.ListByUser(uid) }

func (s *Service) UpdateTask(id uint32, title string, isDone bool) (*Task, error) {
	t := &Task{ID: id}
	// загрузим текущую, чтобы сохранить остальные поля как есть
	cur, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	t.UserID = cur.UserID
	t.Title = title
	t.IsDone = isDone
	if err := s.repo.Update(t); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) DeleteTask(id uint32) error { return s.repo.Delete(id) }
