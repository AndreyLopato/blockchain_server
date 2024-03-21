package service

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Service struct {
	RepoIf repoInterface
}

func NewService(repo repoInterface) *Service {
	return &Service{RepoIf: repo}
}

type repoInterface interface {
	Set(height int, blockchain string) error
	Get(height int) (error, string)
}

func (srv *Service) PerformWcReq(height int) (error, string) {

	queryResp, err := http.Get("https://api.whatsonchain.com/v1/bsv/test/block/height/" + strconv.Itoa(height))
	if err != nil {
		return err, ""
	}
	defer queryResp.Body.Close()

	body, err := io.ReadAll(queryResp.Body)
	if err != nil {
		return err, ""
	}

	err, bdData := srv.RepoIf.Get(height)
	if err != nil {
		return err, ""
	}
	if bdData == "" {
		err := srv.RepoIf.Set(height, string(body))
		if err != nil {
			return err, ""
		}
	} else {
		fmt.Printf("bc with height %d already exists\n", height)
	}

	return nil, string(body)
}
