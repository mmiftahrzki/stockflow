package product

import (
	"backend/internal/logger"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type product struct {
	service Service
	log     *logrus.Entry
}

func New(db *pgxpool.Pool) product {
	return product{
		service: newService(newRepo(db)),
		log:     logger.GetLogger().WithField("component", "Product"),
	}

}

func (h *product) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (h *product) Post(w http.ResponseWriter, r *http.Request) {
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

	command := createCommand{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&command)
	if err != nil {
		h.log.Error(err)

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = command.Validate()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	err = h.service.CreateProduct(r.Context(), command)
	if err != nil {
		if errors.Is(err, errProductSKUAlreadyExists) {
			w.WriteHeader(http.StatusConflict)

			return
		}

		h.log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.log.Info("product created successfully")
	w.WriteHeader(http.StatusCreated)
}

func (h *product) Put(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (h *product) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (h *product) PostStock(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}

func (h *product) GetMovements(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}
