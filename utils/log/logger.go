package log

import(
	"go.uber.org/zap"
	"sync"
)

var mux sync.RWMutex
var logger * zap.SugaredLogger

func Logger() *zap.SugaredLogger {
	mux.Lock()
	defer mux.Unlock()

	if logger == nil {
		l , err := zap.NewProduction()
		//l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		logger = l.Sugar()
	}

	return logger
}
