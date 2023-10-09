package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/go-chi/chi"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
)

func newVMApi(ctx context.Context) *vmApi {
	api := &vmApi{
		Router:  chi.NewRouter(),
		Context: ctx,
	}

	api.Router.Get("/{project}", api.list)
	api.Router.Get("/{project}/{name}", api.get)

	api.Router.Post("/start/{project}/{name}", api.start)
	api.Router.Post("/suspend/{project}/{name}", api.suspend)
	api.Router.Post("/stop/{project}/{name}", api.stop)

	return api
}

type vmApi struct {
	Router  chi.Router
	Context context.Context
}

type VMInstance struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	Zone         string `json:"zone"`
	NetworkIP    string `json:"networkIP"`
	MachineType  string `json:"machineType"`
	CpuPlatform  string `json:"cpuPlatform"`
	CreationTime string `json:"creationTime"`
}

func (api *vmApi) get(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")
	name := chi.URLParam(r, "name")
	zone := r.URL.Query().Get("zone")

	client, err := compute.NewInstancesRESTClient(api.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &computepb.GetInstanceRequest{
		Project:  project,
		Instance: name,
		Zone:     zone,
	}

	instance, err := client.Get(api.Context, req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res := VMInstance{
		Name:         instance.GetName(),
		Status:       instance.GetStatus(),
		Zone:         instance.GetZone(),
		NetworkIP:    instance.GetNetworkInterfaces()[0].GetAccessConfigs()[0].GetNatIP(),
		MachineType:  instance.GetMachineType(),
		CpuPlatform:  instance.GetCpuPlatform(),
		CreationTime: instance.GetCreationTimestamp(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (api *vmApi) list(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")
	zone := r.URL.Query().Get("zone")

	client, err := compute.NewInstancesRESTClient(api.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &computepb.ListInstancesRequest{
		Project:    project,
		Zone:       zone,
		MaxResults: proto.Uint32(10),
	}

	it := client.List(api.Context, req)

	var instances []VMInstance

	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error iterating: %v", err)
		}

		instances = append(instances, VMInstance{
			Name:         instance.GetName(),
			Status:       instance.GetStatus(),
			Zone:         instance.GetZone(),
			NetworkIP:    instance.GetNetworkInterfaces()[0].GetAccessConfigs()[0].GetNatIP(),
			MachineType:  instance.GetMachineType(),
			CpuPlatform:  instance.GetCpuPlatform(),
			CreationTime: instance.GetCreationTimestamp(),
		})
	}

	if err := json.NewEncoder(w).Encode(instances); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type startVMResponse struct {
	Message string `json:"message"`
}

func (api *vmApi) start(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")
	name := chi.URLParam(r, "name")
	zone := r.URL.Query().Get("zone")

	client, err := compute.NewInstancesRESTClient(api.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &computepb.StartInstanceRequest{
		Project:  project,
		Instance: name,
		Zone:     zone,
	}
	_, err = client.Start(api.Context, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &startVMResponse{
		Message: "success",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type suspendVMResponse struct {
	Message string `json:"message"`
}

func (api *vmApi) suspend(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")
	name := chi.URLParam(r, "name")
	zone := r.URL.Query().Get("zone")

	client, err := compute.NewInstancesRESTClient(api.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &computepb.SuspendInstanceRequest{
		Project:  project,
		Instance: name,
		Zone:     zone,
	}

	_, err = client.Suspend(api.Context, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &suspendVMResponse{
		Message: "success",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type stopVMResponse struct {
	Message string `json:"message"`
}

func (api *vmApi) stop(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")
	name := chi.URLParam(r, "name")
	zone := r.URL.Query().Get("zone")

	client, err := compute.NewInstancesRESTClient(api.Context)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &computepb.StopInstanceRequest{
		Project:  project,
		Instance: name,
		Zone:     zone,
	}

	_, err = client.Stop(api.Context, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := &stopVMResponse{
		Message: "success",
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
