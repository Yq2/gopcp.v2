package scheduler

import (
	"gopcp.v2/chapter6/webcrawler/errors"
	"gopcp.v2/chapter6/webcrawler/module"
	"gopcp.v2/chapter6/webcrawler/toolkit/buffer"
)

// genError 用于生成爬虫错误值。
func genError(errMsg string) error {
	return errors.NewCrawlerError(errors.ErrorTypeScheduler,
		errMsg)
}

// genErrorByError 用于基于给定的错误值生成爬虫错误值。
func genErrorByError(err error) error {
	return errors.NewCrawlerError(errors.ErrorTypeScheduler,
		err.Error())
}

// genParameterError 用于生成爬虫参数错误值。
func genParameterError(errMsg string) error {
	return errors.NewCrawlerErrorBy(errors.ErrorTypeScheduler,
		errors.NewIllegalParameterError(errMsg))
}

// sendError 用于向错误缓冲池发送错误值。
func sendError(err error, mid module.MID, errorBufferPool buffer.Pool) bool {
	if err == nil || errorBufferPool == nil || errorBufferPool.Closed() {
		return false
	}
	var crawlerError errors.CrawlerError
	var ok bool
	crawlerError, ok = err.(errors.CrawlerError)
	if !ok {
		var moduleType module.Type
		var errorType errors.ErrorType
		ok, moduleType = module.GetType(mid)
		if !ok {
			errorType = errors.ErrorTypeScheduler
		} else {
			switch moduleType {
			case module.TypeDownloader:
				errorType = errors.ErrorTypeDownloader
			case module.TypeAnalyzer:
				errorType = errors.ErrorTypeAnalyzer
			case module.TypePipeline:
				errorType = errors.ErrorTypePipeline
			}
		}
		crawlerError = errors.NewCrawlerError(errorType, err.Error())
	}
	if errorBufferPool.Closed() {
		return false
	}
	go func(crawlerError errors.CrawlerError) {
		if err := errorBufferPool.Put(crawlerError); err != nil {
			logger.Warnln("The error buffer pool was closed. Ignore error sending.")
		}
	}(crawlerError)
	return true
}
