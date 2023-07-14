package web

import (
	"github.com/variety-jones/cfstress/pkg/executioner"
	"github.com/variety-jones/cfstress/pkg/store"

	"github.com/labstack/echo/v4"
)

type Server struct {
	ec          *echo.Echo
	ticketStore store.TicketStore
	judge       executioner.Judge
}

func CreateWebServer(tickeStore store.TicketStore, judge executioner.Judge) *Server {
	srv := &Server{
		ec:          echo.New(),
		ticketStore: tickeStore,
		judge:       judge,
	}

	srv.ec.GET(kHome, srv.HomeHandler)

	srv.ec.GET(kTest, srv.GetStressTestHandler)
	srv.ec.POST(kTest, srv.PostStressTestHandler)

	srv.ec.GET(kGlobalStatus, srv.GetGlobalStatusHandler)
	srv.ec.GET(kGlobalStatus, srv.PostGlobalStatusHandler)

	srv.ec.GET(kTicketStatus, srv.TicketStatusHandler)

	srv.ec.GET(kMailingList, srv.MailingListHandler)

	srv.ec.POST(kSimulateStressTest, srv.SimulateConcurrentUsers)

	return srv
}
