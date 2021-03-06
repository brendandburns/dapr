/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// CronJobStatus represents the current state of a cron job.
type V1beta1CronJobStatus struct {

	// A list of pointers to currently running jobs.
	Active []V1ObjectReference `json:"active,omitempty"`

	// Information when was the last time the job was successfully scheduled.
	LastScheduleTime time.Time `json:"lastScheduleTime,omitempty"`
}
