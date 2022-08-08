package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riku-smile/go-stations/model"
	"github.com/riku-smile/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req *model.CreateTODORequest
	// メソッドの判定
	if r.Method == "POST" {
		// CreateTODORequestをデコード
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Subjectが空であることを判定
	if req.Subject == "" {
		// BadRequestを渡す
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Subjectが空でない場合
	ctx := r.Context()
	res, err := h.Create(
		ctx,
		req,
	)

	if err != nil {
		fmt.Println(err)
	}

	// CreateTODOをエンコード
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Println(err)
	}

	// StatusOKを返す
	w.WriteHeader(http.StatusOK)
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		fmt.Println(err)
	}

	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
