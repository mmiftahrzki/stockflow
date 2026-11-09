package order

import (
	"backend/internal/logger"
	"backend/internal/product"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type order struct {
	Service Service
	log     *logrus.Entry
}

func New(db *pgxpool.Pool, productService product.Service) order {
	return order{
		Service: newService(newRepo(db), productService),
		log:     logger.GetLogger().WithField("component", "Order"),
	}
}

func (o *order) Post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	contentLenStr := r.Header.Get("Content-Length")
	contentLen, err := strconv.Atoi(contentLenStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if contentLen == 0 {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	command := commandCreate{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&command)
	if err != nil {
		o.log.Error(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = command.Validate()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	err = o.Service.Create(r.Context(), command)
	if err != nil {
		if errors.Is(err, nil) {
			w.WriteHeader(http.StatusConflict)

			return
		}

		o.log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	o.log.Info("order created successfully")
	w.WriteHeader(http.StatusCreated)
}

func (o *order) PostCancel(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (o *order) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}
