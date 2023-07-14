package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/variety-jones/cfstress/pkg/models"
)

func (srv *Server) HomeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) GetStressTestHandler(c echo.Context) error {
	contestID, err := strconv.Atoi(strings.TrimSpace(c.Param("contestID")))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Contest ID should be an integer")
	}

	problemIndex := strings.TrimSpace(c.Param("problemIndex"))
	if problemIndex == "" {
		return c.JSON(http.StatusBadRequest, "Problem should not be empty")
	}

	zap.S().Info(contestID, problemIndex)
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) PostStressTestHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) GetGlobalStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) PostGlobalStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) TicketStatusHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) MailingListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (srv *Server) SimulateConcurrentUsers(c echo.Context) error {
	usersCount, err := strconv.Atoi(c.FormValue("usersCount"))
	if err != nil {
		return c.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	go func() {
		// Use a semaphore to ensure that no more than 100 requests are sent
		// to the database at the same time.
		// Idea taken from: http://jmoiron.net/blog/limiting-concurrency-in-go/
		concurrency := 100
		sem := make(chan bool, concurrency)
		for loopCount := 0; loopCount < usersCount; {
			// This line will keep blocking until there are less than 100 workers
			// alive.
			sem <- true
			ticket := &models.Ticket{
				TicketID: loopCount,
				Type:     "stress",
				Problem: models.Problem{
					ContestID: 123,
					Index:     "a",
				},
				Parameters: "--n_high 10 --n_low 5 --t_high 2 --t_low 1",
			}

			go func() {
				// Once done, free up the resources so other goroutines can
				// be spawned.
				defer func() {
					<-sem
				}()
				var ticketNumber int
				if ticketNumber, err = srv.ticketStore.Add(ticket); err != nil {
					zap.S().Fatal(err)
				}
				if _, err := srv.judge.ProcessTicket(ticket); err != nil {
					zap.S().Fatal(err)
				} else {
					zap.S().Infof("Queued: %d", ticketNumber)
				}
			}()
			loopCount++
		}

		// There might still be goroutines running. We wait for them to complete.
		for i := 0; i < cap(sem); i++ {
			sem <- true
		}
		zap.S().Info("All tickets have been successfully queued.")
	}()
	return c.String(http.StatusOK,
		fmt.Sprintf("Initiated %d stress test requests", usersCount))

}

func (srv *Server) ListenAndServe(addr string) error {
	zap.S().Infof("Starting web server at %s hello", addr)
	return srv.ec.Start(addr)
}
