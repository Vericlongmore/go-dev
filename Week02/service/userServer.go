package service

import (
	"errors"
	"go-dev/Week02/dao"
	"net/http"
	"strconv"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.FormValue("userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := getUser(id); err != nil {

		if errors.Is(err, dao.ErrNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(""))
}

func getUser(ID int) error {
	_, err := dao.QueryUserByID(ID)

	return err
}
