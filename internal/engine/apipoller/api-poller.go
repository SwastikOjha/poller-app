package apipoller

import (
	"fmt"
	"log"
	"poller-app/cmd/config"
	"poller-app/internal/base"
	"poller-app/internal/domain"
	"poller-app/internal/engine"
	"poller-app/internal/service/filehandler"
	"poller-app/internal/service/geturi/repo"
	"time"
)

type ApiPoller struct {
	Logger *log.Logger
	Config config.Config
}

func NewApiPoller(log *log.Logger, conf config.Config) engine.Engine {
	return &ApiPoller{Logger: log, Config: conf}
}

func (ap ApiPoller) Run() error {
	ticker := time.NewTicker(time.Duration(*ap.Config.TimeInterval) * time.Second)
	done := make(chan bool)
	errChan := make(chan error)
	go func() {
		for {
			select {
			case <-done:
				errChan <- nil
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				err := runLogic(ap)
				if err != nil {
					ap.Logger.Printf("%v", err)
					errChan <- err
					return
				}
			}
		}
	}()
	// get the error from the goroutine
	if e := <-errChan; e != nil {
		ap.Logger.Println(e)
		return e
	}

	return nil
}

func runLogic(ap ApiPoller) error {
	params := new(domain.RequestParams)
	params.BaseCurrency = ap.Config.BaseCurrency
	params.ApiKey = ap.Config.ApiKey
	response, err := repo.GetURI(base.ApiURI, *params, ap.Logger)
	if err != nil {
		ap.Logger.Println(err)
		return err
	}
	err = filehandler.WriteFile(*response, *ap.Config.TargetFilePath, ap.Logger)
	if err != nil {
		ap.Logger.Println(err)
		return err
	}
	return nil
}
