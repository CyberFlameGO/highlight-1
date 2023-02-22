package clickhouse

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	modelInputs "github.com/highlight-run/highlight/backend/private-graph/graph/model"
)

func getLogsPayload(logs []*modelInputs.LogEdge) *modelInputs.LogsPayload {
	hasNextPage := len(logs) == LOG_LIMIT+1

	if hasNextPage {
		logs = logs[:LOG_LIMIT]
	}

	return &modelInputs.LogsPayload{
		Edges: logs,
		PageInfo: &modelInputs.PageInfo{
			HasNextPage: hasNextPage,
		},
	}
}

func encodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}
func decodeCursor(encodedCursor string) (res time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errors.New("cursor is invalid")
		return
	}

	res, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return
	}
	uuid = arrStr[1]
	return
}
