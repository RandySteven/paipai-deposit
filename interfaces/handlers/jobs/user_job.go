package job_interfaces

import "context"

type IUserJob interface {
	UpdateUserStatus(ctx context.Context) error
}
