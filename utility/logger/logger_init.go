package logger

import (
	"github.com/op/go-logging"
)


var LOGGER *logging.Logger

func init() {

	LOGGER = logging.MustGetLogger("main")
	//outmode := os.Getenv("DATUMCHI_LOGGER_OUT")
	//if outmode == "" {
	//
	//	format := logging.MustStringFormatter(
	//		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	//	)
	//
	//	backend := logging.NewLogBackend(os.Stderr, "", 0)
	//	backendFormatter := logging.NewBackendFormatter(backend, format)
	//
	//
	//
	//}


}