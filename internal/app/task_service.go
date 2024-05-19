package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	Edit(t domain.Task) (domain.Task, error)
	FindById(id uint64, userId uint64) (domain.Task, error)
	Delete(id uint64, userId uint64) error
	FindByUser(userId uint64) ([]domain.Task, error)
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Edit(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Edit(t)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) FindById(id uint64, userId uint64) (domain.Task, error) {
	task, err := s.taskRepo.FindById(id, userId)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Delete(id uint64, userId uint64) error {
	err := s.taskRepo.Delete(id, userId)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return err
	}
	return nil
}

func (s taskService) FindByUser(userId uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.FindByUser(userId)
	if err != nil {
		log.Printf("TaskService: %s", err)
		return []domain.Task{}, err
	}
	return tasks, nil
}
