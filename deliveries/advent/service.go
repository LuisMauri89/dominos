package advent

import (
	"context"
	"errors"
	"sync"
)

type AdvertService interface {
	List(ctx context.Context) (map[string]Delivery, error)
	Update(ctx context.Context, id string, advert Delivery) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (Delivery, error)
	GetByStatus(ctx context.Context, status string) (Delivery, error)
}

type AdvertRepository interface {
	Find(ctx context.Context) (map[string]Delivery, error)
	Update(ctx context.Context, id string, advert Delivery) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (Delivery, error)
	GetByStatus(ctx context.Context, status string) (Delivery, error)
}

type Delivery struct {
	ID          string `json:"id"`
	OrderID     string `json:"orderid,omitempty"`
	Status      string `json:"status,omitempty"` // PENDING - READY - DELIVERED
	To          string `json:"to,omitempty"`
	FinalPrice  int    `json:"finalprice,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inMemoryRepository struct {
	mtx sync.RWMutex
	ads map[string]Delivery
}

func NewInMemoryRepository() AdvertRepository {
	return &inMemoryRepository{
		ads: map[string]Delivery{},
	}
}

func (r *inMemoryRepository) Find(ctx context.Context) (map[string]Delivery, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.ads, nil
}

func (r *inMemoryRepository) Update(ctx context.Context, id string, advert Delivery) error {
	if id != advert.ID {
		return ErrInconsistentIDs
	}
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.ads[id] = advert
	return nil
}

func (r *inMemoryRepository) Delete(ctx context.Context, id string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.ads[id]; !ok {
		return ErrNotFound
	}
	delete(r.ads, id)
	return nil
}

func (r *inMemoryRepository) GetById(ctx context.Context, id string) (Delivery, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ad, ok := r.ads[id]
	if !ok {
		return Delivery{}, ErrNotFound
	}
	return ad, nil
}

func (r *inMemoryRepository) GetByStatus(ctx context.Context, id string) (Delivery, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ad, ok := r.ads[status]
	if !ok {
		return Delivery{}, ErrNotFound
	}
	return ad, nil
}

type advertService struct {
	repository AdvertRepository
}

func NewService(repository AdvertRepository) AdvertService {
	return &advertService{
		repository: repository,
	}
}

func (s *advertService) List(ctx context.Context) (map[string]Delivery, error) {
	ads, err := s.repository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (s *advertService) Update(ctx context.Context, id string, advert Delivery) error {
	ad, err := s.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	ad.OrderID = advert.OrderID
	ad.Status = advert.Status
	ad.To = advert.To
	ad.FinalPrice = advert.FinalPrice
	ad.Address = advert.Address
	ad.Description = advert.Description

	if err := s.repository.Update(ctx, id, ad); err != nil {
		return err
	}
	return nil
}

func (s *advertService) Delete(ctx context.Context, id string) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *advertService) GetById(ctx context.Context, id string) (Delivery, error) {
	ad, err := s.repository.GetById(ctx, id)
	if err != nil {
		return ad, err
	}
	return ad, nil
}
func (s *advertService) GetByStatus(ctx context.Context, status string) (Delivery, error) {
	ad, err := s.repository.GetByStatus(ctx, status)
	if err != nil {
		return ad, err
	}
	return ad, nil
}

//BUSCAR POR ID CAMBIAR A BUSCAR POR ESTADO
