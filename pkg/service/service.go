package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"main/pkg/repository"
	"net/http"
	"strconv"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (srv *Service) PerformWcReq(height int) error {

	queryResp, err := http.Get("https://api.whatsonchain.com/v1/bsv/test/block/height/" + strconv.Itoa(height))
	if err != nil {
		return err
	}
	defer queryResp.Body.Close()

	body, err := io.ReadAll(queryResp.Body)
	if err != nil {
		return err
	}

	err, bdData := srv.repo.Get(height)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}
	if bdData == "" {
		srv.repo.Set(height, string(body))
	} else {
		fmt.Printf("bc with height %d already exists\n", height)
	}

	return nil
}
