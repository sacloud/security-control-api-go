// Copyright 2026- The security-control-api-go Authors
// SPDX-License-Identifier: Apache-2.0

package securitycontrol

import (
	"context"
	"errors"

	v1 "github.com/sacloud/security-control-api-go/apis/v1"
)

type ActivationsAPI interface {
	Create(ctx context.Context, servicePrincipalID string) (*v1.ActivationOutput, error)
	Read(ctx context.Context) (*v1.ActivationOutput, error)
	Update(ctx context.Context, servicePrincipalID string, isActive bool) (*v1.ActivationOutput, error)
}

var _ ActivationsAPI = (*activationsOp)(nil)

type activationsOp struct {
	client *v1.Client
}

func NewActivationsOp(client *v1.Client) ActivationsAPI {
	return &activationsOp{client: client}
}

func (op *activationsOp) Create(ctx context.Context, servicePrincipalID string) (*v1.ActivationOutput, error) {
	const methodName = "Activations.Create"
	res, err := op.client.ProjectActivationCreate(ctx, &v1.ActivationCreateInput{ServicePrincipalId: servicePrincipalID})
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.ActivationOutput:
		return r, nil
	case *v1.ProjectActivationCreateApplicationJSONBadRequest:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationCreateApplicationJSONUnauthorized:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationCreateApplicationJSONConflict:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationCreateApplicationJSONForbidden:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationCreateApplicationJSONTooManyRequests:
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	case *v1.UnexpectedErrorStatusCode:
		return nil, NewAPIError(methodName, r.StatusCode, errors.New(r.Response.Detail.Value))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}

func (op *activationsOp) Read(ctx context.Context) (*v1.ActivationOutput, error) {
	const methodName = "Activations.Read"

	res, err := op.client.ProjectActivationRead(ctx)
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.ActivationOutput:
		return r, nil
	case *v1.ProjectActivationReadApplicationJSONUnauthorized:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationReadApplicationJSONForbidden:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationReadApplicationJSONNotFound:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationReadApplicationJSONTooManyRequests:
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	case *v1.UnexpectedErrorStatusCode:
		return nil, NewAPIError(methodName, r.StatusCode, errors.New(r.Response.Detail.Value))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}

func (op *activationsOp) Update(ctx context.Context, servicePrincipalID string, isActive bool) (*v1.ActivationOutput, error) {
	const methodName = "Activations.Update"

	res, err := op.client.ProjectActivationUpdate(ctx, &v1.ActivationUpdateInput{
		ServicePrincipalId: servicePrincipalID,
		IsActive:           isActive,
	})
	if err != nil {
		return nil, NewAPIError(methodName, 0, err)
	}

	switch r := res.(type) {
	case *v1.ActivationOutput:
		return r, nil
	case *v1.ProjectActivationUpdateApplicationJSONBadRequest:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationUpdateApplicationJSONUnauthorized:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationUpdateApplicationJSONForbidden:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationUpdateApplicationJSONNotFound:
		return nil, NewAPIError(methodName, int(r.Status.Value), errors.New(string(r.Detail.Value)))
	case *v1.ProjectActivationUpdateApplicationJSONTooManyRequests:
		return nil, NewAPIError(methodName, int(r.Response.Status.Value), errors.New(string(r.Response.Detail.Value)))
	case *v1.UnexpectedErrorStatusCode:
		return nil, NewAPIError(methodName, r.StatusCode, errors.New(r.Response.Detail.Value))
	default:
		return nil, NewAPIError(methodName, 0, errors.New("unknown error"))
	}
}
