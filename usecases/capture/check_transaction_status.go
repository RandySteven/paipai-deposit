package capture

import "context"

func (c *captureWorkflow) checkTransactionStatus(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	return executionData, nil
}
