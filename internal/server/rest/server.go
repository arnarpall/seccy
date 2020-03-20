package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arnarpall/seccy/internal/log"
	"github.com/arnarpall/seccy/internal/server"
	"github.com/arnarpall/seccy/internal/version"
	"github.com/arnarpall/seccy/pkg/client"
	"github.com/julienschmidt/httprouter"
)

type restApi struct {
	logger        *log.Logger
	listenAddress string
	client        client.Client
}

type keyValue struct {
	Key string
	Value string
}

func (ra restApi) Serve() error {
	router := httprouter.New()
	router.GET("/healthz", ra.health)
	router.GET("/live", ra.health)
	router.GET("/version", ra.version)
	router.GET("/keys", ra.listKeys)
	router.POST("/keys", ra.addKeyValue)
	router.GET("/keys/:key", ra.getValue)

	return http.ListenAndServe(ra.listenAddress, router)
}

func (ra restApi) version(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, version.BuildVersion)
}

func (ra restApi) health(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "OK")
}

func (ra restApi) listKeys(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	keyChan, err := ra.client.ListKeys()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var keys = make([]string, 0, 10)
	for key := range keyChan {
		keys = append(keys, key)
	}

	ra.json(w, keys)
}

func (ra restApi) json(w http.ResponseWriter, v interface{} )  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func (ra restApi) getValue(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	key := p.ByName("key")

	if key == "" {
		http.Error(w, "Required parameter 'key' missing", http.StatusBadRequest)
		return
	}

	val, err := ra.client.Get(key)
	if err != nil {
		ra.logger.Error(err)
		http.Error(w, fmt.Sprintf("unable to get value for key %s", key), http.StatusNotFound)
		return
	}

	ra.json(w, val)
}

func (ra restApi) addKeyValue(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	kv := new(keyValue)

	err := json.NewDecoder(r.Body).Decode(kv)
	if err != nil {
		ra.logger.Error(err)
		http.Error(w, "unable to read key value from request", http.StatusBadRequest)
		return
	}

	if kv.Key == "" {
		http.Error(w, "key is missing from request", http.StatusBadRequest)
		return
	}

	if kv.Value == "" {
		http.Error(w, "value is missing from request", http.StatusBadRequest)
		return
	}

	err = ra.client.Set(kv.Key, kv.Value)
	if err != nil {
		ra.logger.Error(err)
		http.Error(w, "unable to save key value", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func New(logger *log.Logger, listenAddress string, c client.Client) server.Server {
	return &restApi{
		logger:        logger,
		listenAddress: listenAddress,
		client:        c,
	}
}
