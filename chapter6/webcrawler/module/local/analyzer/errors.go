package analyzer

import "gopcp.v2/chapter6/webcrawler/errors"

// genError 用于生成爬虫错误值。
func genError(errMsg string) error {
	return errors.NewCrawlerError(errors.ErrorTypeAnalyzer, errMsg)
}

// genParameterError 用于生成爬虫参数错误值。
func genParameterError(errMsg string) error {
	return errors.NewCrawlerErrorBy(errors.ErrorTypeAnalyzer,
		errors.NewIllegalParameterError(errMsg))
}
