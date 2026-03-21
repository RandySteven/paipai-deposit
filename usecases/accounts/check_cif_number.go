package accounts

import "context"

func (a *accountWorkflow) checkCifNumber(ctx context.Context, executionData *ExecutionData) (err error) {
	executionData.SetActivity(activityCreateAccount)
	return nil
}
