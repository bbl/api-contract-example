package main

import (
	"context"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api-contracts-example/generated/api"
)

type StoreServer struct {
	stores map[string]api.Store
	mu     sync.RWMutex
}

func NewStoreServer() *StoreServer {
	return &StoreServer{
		stores: make(map[string]api.Store),
	}
}

func (s *StoreServer) StoresList(ctx context.Context, request api.StoresListRequestObject) (api.StoresListResponseObject, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []api.Store
	for _, store := range s.stores {
		if strings.Contains(store.Name, request.Params.Filter) {
			result = append(result, store)
		}
	}

	if result == nil {
		result = []api.Store{}
	}

	return api.StoresList200JSONResponse(result), nil
}

func (s *StoreServer) StoresCreate(ctx context.Context, request api.StoresCreateRequestObject) (api.StoresCreateResponseObject, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	store := *request.Body
	s.stores[id] = store

	return api.StoresCreate200JSONResponse(store), nil
}

func (s *StoreServer) StoresRead(ctx context.Context, request api.StoresReadRequestObject) (api.StoresReadResponseObject, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	store, ok := s.stores[request.Id]
	if !ok {
		return api.StoresRead200JSONResponse{}, nil
	}

	return api.StoresRead200JSONResponse(store), nil
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := NewStoreServer()
	strictHandler := api.NewStrictHandler(server, nil)
	api.RegisterHandlers(e, strictHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
