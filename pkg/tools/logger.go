package tools

import (
	"go.uber.org/zap"
)

var LoggerBase *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	LoggerBase, _ = zap.NewProduction()

	SugaredLogger = LoggerBase.Sugar()
}
