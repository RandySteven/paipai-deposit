package capture

import "context"

func (c *captureWorkflow) balanceDeduct(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	return executionData, nil
}
