package clickhouse

import (
	"context"
	"fmt"
	"testing"
	"time"

	modelInputs "github.com/highlight-run/highlight/backend/private-graph/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ReadLogsTestSuite struct {
	suite.Suite
	ctx    context.Context
	client *Client
}

func (suite *ReadLogsTestSuite) SetupTest() {
	suite.ctx = context.Background()
	client, err := setupClickhouseTestDB()

	suite.Require().NoError(err)

	suite.client = client
}

func (suite *ReadLogsTestSuite) TeardownTest() {
	suite.client.conn.Exec(context.Background(), fmt.Sprintf("TRUNCATE TABLE %s", LogsTable)) //nolint:errcheck
}

func (suite *ReadLogsTestSuite) TestReadLogs(t *testing.T) {
	now := time.Now()
	rows := []*LogRow{
		{
			Timestamp:    now,
			ProjectId:    1,
			Body:         "body",
			SeverityText: "info",
		},
	}

	require.NoError(t, suite.client.BatchWriteLogRows(suite.ctx, rows))
	logLines, err := suite.client.ReadLogs(suite.ctx, 1, modelInputs.LogsParamsInput{})
	require.NoError(t, err)

	assert.Len(t, logLines, 1)
	assert.Equal(t, now.UnixMilli(), logLines[0].Timestamp.UnixMilli())
	assert.Equal(t, "info", logLines[0].SeverityText)
	assert.Equal(t, "body", logLines[0].Body)
	assert.Equal(t, map[string]interface{}{"project_id": "7"}, logLines[0].LogAttributes)
}

func (suite *ReadLogsTestSuite) TestReadLogsWithTimeQuery(t *testing.T) {
	now := time.Now()
	rows := []*LogRow{
		{
			Timestamp: now,
			ProjectId: 1,
		},
	}

	require.NoError(t, suite.client.BatchWriteLogRows(suite.ctx, rows))

	logLines, err := suite.client.ReadLogs(suite.ctx, 1, modelInputs.LogsParamsInput{
		DateRange: &modelInputs.DateRangeRequiredInput{
			StartDate: now.Add(-time.Hour * 2),
			EndDate:   now.Add(-time.Hour * 1),
		},
	})
	require.NoError(t, err)

	assert.Len(t, logLines, 0)

	logLines, err = suite.client.ReadLogs(suite.ctx, 1, modelInputs.LogsParamsInput{
		DateRange: &modelInputs.DateRangeRequiredInput{
			StartDate: now.Add(-time.Hour * 1),
			EndDate:   now.Add(time.Hour * 1),
		},
	})
	require.NoError(t, err)

	assert.Len(t, logLines, 1)
}
