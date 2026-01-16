// Copyright 2026- The security-control-api-go Authors
// SPDX-License-Identifier: Apache-2.0

package securitycontrol

import (
	"context"
	"errors"
	"net/http"

	v1 "github.com/sacloud/security-control-api-go/apis/v1"
)

type EvaluationRulesAPI interface {
	List(ctx context.Context, parameter v1.EvaluationRulesListParams) (*v1.EvaluationRulesListOK, error)
	Read(ctx context.Context, id string) (*v1.EvaluationRule, error)
	Update(ctx context.Context, id string, request v1.EvaluationRuleInput) (*v1.EvaluationRule, error)
}

var _ EvaluationRulesAPI = (*evaluationRuleOp)(nil)

type evaluationRuleOp struct {
	client *v1.Client
}

func NewEvaluationRulesOp(client *v1.Client) EvaluationRulesAPI {
	return &evaluationRuleOp{client: client}
}

func (op *evaluationRuleOp) List(ctx context.Context, params v1.EvaluationRulesListParams) (*v1.EvaluationRulesListOK, error) {
	const methodName = "EvaluationRules.List"

	res, err := op.client.EvaluationRulesList(ctx, params)
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.EvaluationRulesListOK:
		return r, nil
	case *v1.EvaluationRulesListApplicationJSONBadRequest:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesListApplicationJSONUnauthorized:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesListApplicationJSONForbidden:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesListApplicationJSONNotFound:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesListApplicationJSONTooManyRequests:
		// TODO: Handle RetryAfter field
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	case *v1.UnexpectedErrorStatusCode:
		return nil, NewAPIError(methodName, r.StatusCode, errors.New(r.Response.Detail.Value))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}

func (op *evaluationRuleOp) Read(ctx context.Context, id string) (*v1.EvaluationRule, error) {
	const methodName = "EvaluationRules.Read"

	res, err := op.client.EvaluationRulesRead(ctx, v1.EvaluationRulesReadParams{EvaluationRuleId: v1.EvaluationRuleID(id)})
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.EvaluationRule:
		return r, nil
	case *v1.EvaluationRulesReadApplicationJSONNotFound:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesReadApplicationJSONForbidden:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesReadApplicationJSONTooManyRequests:
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}

func (op *evaluationRuleOp) Update(ctx context.Context, id string, req v1.EvaluationRuleInput) (*v1.EvaluationRule, error) {
	const methodName = "EvaluationRules.Update"

	res, err := op.client.EvaluationRulesUpdate(ctx, &req, v1.EvaluationRulesUpdateParams{EvaluationRuleId: v1.EvaluationRuleID(id)})
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.EvaluationRule:
		return r, nil
	case *v1.EvaluationRulesUpdateApplicationJSONBadRequest:
		return nil, NewAPIError(methodName, http.StatusBadRequest, errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesUpdateApplicationJSONUnauthorized:
		return nil, NewAPIError(methodName, http.StatusUnauthorized, errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesUpdateApplicationJSONForbidden:
		return nil, NewAPIError(methodName, http.StatusForbidden, errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesUpdateApplicationJSONNotFound:
		return nil, NewAPIError(methodName, http.StatusNotFound, errors.New(string(r.Detail.Value)))
	case *v1.EvaluationRulesUpdateApplicationJSONTooManyRequests:
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}
