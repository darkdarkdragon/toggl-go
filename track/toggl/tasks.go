package toggl

import (
	"context"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Task represents the properties of a task.
type Task struct {
	Active           *bool      `json:"active,omitempty"`
	At               *time.Time `json:"at,omitempty"`
	EstimatedSeconds *int       `json:"estimated_seconds,omitempty"`
	ID               *int       `json:"id,omitempty"`
	Name             *string    `json:"name,omitempty"`
	ProjectID        *int       `json:"project_id,omitempty"`
	Recurring        *bool      `json:"recurring,omitempty"`
	ServerDeletedAt  *time.Time `json:"server_deleted_at,omitempty"`
	TrackedSeconds   *int       `json:"tracked_seconds,omitempty"`
	UserID           *int       `json:"user_id,omitempty"`
	WorkspaceID      *int       `json:"workspace_id,omitempty"`
}

// GetTags lists workspace tags.
func (c *APIClient) GetTasks(ctx context.Context, workspaceID int, projectID int) ([]*Task, error) {
	var tasks []*Task
	apiSpecificPath := path.Join(workspacesPath, strconv.Itoa(workspaceID), "projects", strconv.Itoa(projectID), "tasks")
	if err := c.httpGet(ctx, apiSpecificPath, nil, &tasks); err != nil {
		return nil, errors.Wrap(err, "failed to get tags")
	}
	return tasks, nil
}
