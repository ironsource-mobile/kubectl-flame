// : Copyright Verizon Media
// : Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.
package details

import (
	"time"

	"github.com/ironsource-mobile/kubectl-flame/api"
)

type ProfilingJob struct {
	Duration          time.Duration
	ID                string
	ContainerID       string
	ContainerName     string
	PodUID            string
	Language          api.ProgrammingLanguage
	TargetProcessName string
	Event             api.ProfilingEvent
}
