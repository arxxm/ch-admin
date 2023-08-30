package clickhousebulk

import (
	"context"
	"errors"
	"fmt"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/domain/task"
	"pingocean-front/clickhouse-admin-sms-consumer/pkg/logger"

	"github.com/valyala/fasthttp"
)

type Sms interface {
	InsertSmsTask(ctx context.Context, sms task.Sms) error
}

func (c *ClickhousebulkRepository) InsertSmsTask(ctx context.Context, sms task.Sms) error {
	logger.Info(fmt.Sprintf("InsertSmsTask new, task_id: %v", sms.ID))

	body := []byte(sms.GetInsertSqlString())

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)   // <- do not forget to release
	defer fasthttp.ReleaseResponse(resp) // <- do not forget to release

	req.SetRequestURI(c.cfg.CHBulk.Url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBodyRaw(body)

	if err := fasthttp.Do(req, resp); err != nil {
		logger.Error(fmt.Sprintf("InsertSmsTask http init err, task_id: %v, err: %v", sms.ID, err))
		return err
	}

	statusCode := resp.StatusCode()
	if statusCode != fasthttp.StatusOK {
		logger.Error(fmt.Sprintf("InsertSmsTask http response status err, task_id: %v, status_code: %v", sms.ID, statusCode))
		return errors.New("InsertSmsTask http response status is not 200")
	}

	logger.Info(fmt.Sprintf("InsertSmsTask sent, task_id: %v", sms.ID))

	return nil
}
