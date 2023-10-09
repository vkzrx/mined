package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	vm "github.com/vkzrx/mined/vm"
	proto "github.com/vkzrx/mined/vm/proto"
)

type minecraftApi struct {
	Context context.Context
	Client  *vm.MinecraftClient
	Router  chi.Router
}

func newMinecraftApi(ctx context.Context) (*minecraftApi, error) {
	client, err := vm.NewMinecraftClient("localhost:5001")
	if err != nil {
		return nil, err
	}

	api := &minecraftApi{
		Context: ctx,
		Client:  client,
		Router:  chi.NewRouter(),
	}

	api.Router.Post("/start", api.start)

	return api, nil
}

type StartResponse struct {
	Message string
}

func (api *minecraftApi) start(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(api.Context, time.Second)
	defer cancel()

	result, err := api.Client.Service.Start(ctx, &proto.StartRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("result", result)

	res := &StartResponse{Message: "success"}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
