package kardinal

import (
	"context"
	"errors"
	"fmt"
	"time"

	openapi "github.com/dropon/kardinal-optimisation-client"
)

const defaultPollInterval = time.Second

// AuthWrapper lets callers inject authentication data into outgoing contexts.
type AuthWrapper func(context.Context) context.Context

// Client wraps the generated API client with Optiflow-compatible helpers.
type Client struct {
	API          *openapi.APIClient
	AgencyID     openapi.AgencyId
	PollInterval time.Duration
	AuthWrapper  AuthWrapper
}

// NewClient builds a helper client with sensible defaults.
func NewClient(api *openapi.APIClient, agencyID openapi.AgencyId, authWrapper AuthWrapper) *Client {
	return &Client{
		API:          api,
		AgencyID:     agencyID,
		PollInterval: defaultPollInterval,
		AuthWrapper:  authWrapper,
	}
}

// StartOptimization creates or updates a plan and returns its identifier.
func (c *Client) StartOptimization(ctx context.Context, req openapi.Plan) (string, error) {
	if err := c.validate(); err != nil {
		return "", err
	}

	planID, err := planIDFromPlan(req)
	if err != nil {
		return "", err
	}

	plan := req
	plan.SetAgencyId(string(c.AgencyID))

	ctx = c.wrapAuth(ctx)

	resp, _, err := c.API.PlanAPI.PutPlan(ctx, c.AgencyID, planID).Plan(plan).Execute()
	if err != nil {
		return "", err
	}

	if id := envelopePlanID(resp); id != "" {
		return id, nil
	}

	return planID, nil
}

// GetOptimization fetches the current state of a plan.
func (c *Client) GetOptimization(ctx context.Context, id string) (*openapi.Plan, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if id == "" {
		return nil, errors.New("plan id is required")
	}

	ctx = c.wrapAuth(ctx)

	env, _, err := c.API.PlanAPI.GetPlan(ctx, c.AgencyID, id).Execute()
	if err != nil {
		return nil, err
	}

	plan, ok := env.GetItemOk()
	if !ok || plan == nil {
		return nil, errors.New("plan payload missing from response")
	}

	return plan, nil
}

// MustGetOptimization polls until the plan reaches a terminal state.
func (c *Client) MustGetOptimization(ctx context.Context, id string) (*openapi.Plan, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	var (
		state openapi.PlanState = openapi.PLANSTATE_WAITING
		plan  *openapi.Plan
		err   error
	)

	for running := true; running; running = isRunning(state) {
		plan, err = c.GetOptimization(ctx, id)
		if err != nil {
			time.Sleep(c.pollDelay())
			continue
		}

		state = plan.GetState()
		if isRunning(state) {
			time.Sleep(c.pollDelay())
		}
	}

	if state != openapi.PLANSTATE_OPTIMIZED {
		return plan, fmt.Errorf("plan %s finished with state %s", id, state)
	}

	return plan, nil
}

func (c *Client) validate() error {
	if c == nil {
		return errors.New("client is nil")
	}
	if c.API == nil {
		return errors.New("api client is nil")
	}
	if c.API.PlanAPI == nil {
		return errors.New("plan api client is nil")
	}
	if c.AgencyID == "" {
		return errors.New("agency id is required")
	}
	return nil
}

func (c *Client) wrapAuth(ctx context.Context) context.Context {
	if c == nil || c.AuthWrapper == nil {
		return ctx
	}
	return c.AuthWrapper(ctx)
}

func (c *Client) pollDelay() time.Duration {
	if c != nil && c.PollInterval > 0 {
		return c.PollInterval
	}
	return defaultPollInterval
}

func isRunning(state openapi.PlanState) bool {
	switch state {
	case "", openapi.PLANSTATE_WAITING, openapi.PLANSTATE_PROCESSING, openapi.PLANSTATE_PRE_OPTIMIZING, openapi.PLANSTATE_OPTIMIZING:
		return true
	default:
		return false
	}
}

func planIDFromPlan(plan openapi.Plan) (string, error) {
	id, err := interfaceToString(plan.GetId())
	if err != nil || id == "" {
		return "", errors.New("plan id is required")
	}
	return id, nil
}

func envelopePlanID(env *openapi.EnvelopedPlan) string {
	if env == nil {
		return ""
	}
	if raw, ok := env.GetPlanIdOk(); ok && raw != nil {
		if id, err := interfaceToString(*raw); err == nil {
			return id
		}
	}
	if item, ok := env.GetItemOk(); ok && item != nil {
		if id, err := interfaceToString(item.GetId()); err == nil {
			return id
		}
	}
	return ""
}

func interfaceToString(value interface{}) (string, error) {
	if value == nil {
		return "", errors.New("value is nil")
	}

	switch v := value.(type) {
	case string:
		return v, nil
	case fmt.Stringer:
		return v.String(), nil
	case []byte:
		return string(v), nil
	default:
		s := fmt.Sprint(v)
		if s == "<nil>" {
			return "", errors.New("value is nil")
		}
		return s, nil
	}
}
