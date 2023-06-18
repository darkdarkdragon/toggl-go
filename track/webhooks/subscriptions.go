package webhooks

import (
	"context"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type EventFilter struct {
	Action string `json:"action,omitempty"`
	Entity string `json:"entity,omitempty"`
}

// EventFilters represents the properties of event filters.
type Subscription struct {
	CreatedAt        *time.Time     `json:"created_at,omitempty"`
	DeletedAt        *time.Time     `json:"deleted_at,omitempty"`
	Description      *string        `json:"description,omitempty"`
	Enabled          *bool          `json:"enabled,omitempty"`
	EventFilters     []*EventFilter `json:"event_filters,omitempty"`
	HasPendingEvents *bool          `json:"has_pending_events,omitempty"`
	Secret           *string        `json:"secret,omitempty"`
	SubscriptionID   *int           `json:"subscription_id,omitempty"`
	UpdatedAt        *time.Time     `json:"updated_at,omitempty"`
	UrlCallback      *string        `json:"url_callback,omitempty"`
	UserID           *int           `json:"user_id,omitempty"`
	ValidatedAt      *time.Time     `json:"validated_at,omitempty"`
	WorkspaceID      *int           `json:"workspace_id,omitempty"`
}

// GetSubscriptions gets the list of subscriptions.
func (c *APIClient) GetSubscriptions(ctx context.Context, worspaceID int) ([]*Subscription, error) {
	var subscriptions []*Subscription
	apiSpecificPath := path.Join(webhooksPath, "subscriptions", strconv.FormatInt(int64(worspaceID), 10))
	if err := c.httpGet(ctx, apiSpecificPath, nil, &subscriptions); err != nil {
		return nil, errors.Wrap(err, "failed to get subscriptions")
	}
	return subscriptions, nil
}

// CreateSubscriptionRequestBody represents a request body of CreateSubscription.
type CreateSubscriptionRequestBody struct {
	Description  *string       `json:"description,omitempty"`
	Enabled      *bool         `json:"enabled,omitempty"`
	URLCallback  *string       `json:"url_callback,omitempty"`
	Secret       *string       `json:"secret,omitempty"`
	EventFilters []EventFilter `json:"event_filters,omitempty"`
}

// SearchDetailedReport returns time entries for detailed report.
func (c *APIClient) CreteSubscription(ctx context.Context, workspaceID int, reqBody *CreateSubscriptionRequestBody) (*Subscription, error) {
	var sub = &Subscription{}
	apiSpecificPath := path.Join(webhooksPath, "subscriptions", strconv.Itoa(workspaceID))
	if err := c.httpPost(ctx, apiSpecificPath, reqBody, sub); err != nil {
		return nil, errors.Wrap(err, "failed to create webhook")
	}
	return sub, nil
}

// SearchDetailedReport returns time entries for detailed report.
func (c *APIClient) SetEnabled(ctx context.Context, workspaceID int, subsID int, enabled bool) (*Subscription, error) {
	var sub = &Subscription{}
	reqBody := struct {
		Enabled bool `json:"enabled"`
	}{Enabled: enabled}
	apiSpecificPath := path.Join(webhooksPath, "subscriptions", strconv.Itoa(workspaceID), strconv.Itoa(subsID))
	if err := c.httpPatch(ctx, apiSpecificPath, reqBody, sub); err != nil {
		return nil, errors.Wrap(err, "failed to set enabled for webhook")
	}
	return sub, nil
}
