package capture

import "context"

func (c *captureWorkflow) updateTransactionStatus(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	return executionData, nil
}
