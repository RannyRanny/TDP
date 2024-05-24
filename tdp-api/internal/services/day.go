package service

import (
	"tdp-api/internal/model"
	"tdp-api/internal/repository"
	"time"
)

// DayService encapsulates business logic for day management.
type DayService interface {
	CreateDay(day model.Day) (model.Day, error)
	GetDay(id uint) (model.Day, error)
	UpdateDay(day model.Day) (model.Day, error)
	DeleteDay(id uint) error
	GetDays(userID uint) ([]model.Day, error)
	GetDaysByDate(date time.Time, userID uint) ([]model.Day, error)
	CreateTemplateDay(day model.TemplateDay) (model.TemplateDay, error)
	GetTemplateDay(id uint) (model.TemplateDay, error)
	UpdateTemplateDay(day model.TemplateDay) (model.TemplateDay, error)
	DeleteTemplateDay(id uint) error
	GetTemplateDays(userID uint) ([]model.TemplateDay, error)
}

type dayService struct {
	repo         repository.DayRepository
	templateRepo repository.TemplateDayRepository
}

func NewDayService(repo repository.DayRepository,
	templateRepo repository.TemplateDayRepository) DayService {
	return &dayService{repo: repo, templateRepo: templateRepo}
}

func (s *dayService) CreateDay(day model.Day) (model.Day, error) {
	return s.repo.Create(day)
}

func (s *dayService) GetDay(id uint) (model.Day, error) {
	return s.repo.FindByID(id)
}

func (s *dayService) UpdateDay(day model.Day) (model.Day, error) {
	return s.repo.Update(day)
}

func (s *dayService) DeleteDay(id uint) error {
	return s.repo.Delete(id)
}

func (s *dayService) GetDays(userID uint) ([]model.Day, error) {
	return s.repo.GetDays(userID)
}

func (s *dayService) GetDaysByDate(date time.Time, userID uint) ([]model.Day, error) {
	return s.repo.FindByDate(date, userID)
}

func (s *dayService) CreateTemplateDay(day model.TemplateDay) (model.TemplateDay, error) {
	return s.templateRepo.Create(day)
}

func (s *dayService) GetTemplateDay(id uint) (model.TemplateDay, error) {
	return s.templateRepo.FindByID(id)
}

func (s *dayService) UpdateTemplateDay(day model.TemplateDay) (model.TemplateDay, error) {
	return s.templateRepo.Update(day)
}

func (s *dayService) DeleteTemplateDay(id uint) error {
	return s.templateRepo.Delete(id)
}

func (s *dayService) GetTemplateDays(userID uint) ([]model.TemplateDay, error) {
	return s.templateRepo.GetDays(userID)
}
