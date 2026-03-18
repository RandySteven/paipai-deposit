package jobs

import job_interfaces "github.com/RandySteven/go-kopi/interfaces/handlers/jobs"

type (
	Job struct {
		UserJob job_interfaces.IUserJob
	}
)

func NewJob(userJob job_interfaces.IUserJob) *Job {
	return &Job{
		UserJob: userJob,
	}
}

