package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Create(w, tDto)
	}
}

func (c TaskController) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.EditTaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		task.UserId = user.Id
		task, err = c.taskService.Edit(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Create(w, tDto)
	}
}

func (c TaskController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var TaskFindById requests.TaskFindById
		if err := json.NewDecoder(r.Body).Decode(&TaskFindById); err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		id := TaskFindById.Id
		user := r.Context().Value(UserKey).(domain.User)
		task, err := c.taskService.FindById(id, user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}
		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Success(w, tDto)

	}

}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var TaskFindById requests.TaskFindById
		if err := json.NewDecoder(r.Body).Decode(&TaskFindById); err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		id := TaskFindById.Id
		user := r.Context().Value(UserKey).(domain.User)
		err := c.taskService.Delete(id, user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}
		Ok(w)

	}

}

func (c TaskController) GetByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := r.Context().Value(UserKey).(domain.User)
		tasks, err := c.taskService.FindByUser(user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		var tDtos []resources.TaskDto
		for _, task := range tasks {
			tDto = resources.TaskDto{}.DomainToDto(task)
			tDtos = append(tDtos, tDto)
		}
		Success(w, interface{}(tDtos))

	}

}
