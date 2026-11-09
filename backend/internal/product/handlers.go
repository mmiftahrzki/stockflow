package product

import (
	"backend/internal/logger"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type product struct {
	Service Service
	log     *logrus.Entry
}

func New(db *pgxpool.Pool) product {
	return product{
		Service: newService(newRepo(db)),
		log:     logger.GetLogger().WithField("component", "Product"),
	}
}

func (h *product) Get(w http.ResponseWriter, r *http.Request) {
	query := QueryList{}
	queryParams := r.URL.Query()

	search := queryParams.Get("search")
	pageStr := queryParams.Get("page")
	limitStr := queryParams.Get("limit")

	if search != "" {
		query.search = search
	}

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		query.page = page
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		query.limit = limit
	}

	products, err := h.Service.List(r.Context(), query)
	if err != nil {
		h.log.Error(err)

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	buffer := bytes.NewBuffer(nil)
	json_encoder := json.NewEncoder(buffer)
	err = json_encoder.Encode(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(buffer.Bytes())

	h.log.Info("products data retrieved successfully")
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

	command := commandCreate{}
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

	err = h.Service.Create(r.Context(), command)
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
