package accounts

import "context"

func (a *accountWorkflow) checkCifNumber(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	executionData.SetActivity(activityCreateAccount)
	return executionData, nil
}
