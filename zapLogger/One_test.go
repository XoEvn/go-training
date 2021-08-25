/**
 * @Author: evnxo
 * @Description:
 * @File:  One_Test
 * @Version: 1.0.0
 * @Date: 2021/7/12 9:07 下午
 */

package zapLogger

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestSugar(t *testing.T) {
	logger, _ := zap.NewProduction()
	// 默认 logger 不缓冲。
	// 但由于底层 api 允许缓冲，所以在进程退出之前调用 Sync 是一个好习惯。
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infof("Failed to fetch URL: %s", "https://baidu.com")
}

func TestLogger(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// 强类型字段
		zap.String("url", "https://baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
