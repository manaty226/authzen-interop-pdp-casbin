// Code generated by ogen, DO NOT EDIT.

package server

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// EvaluateAccess implements evaluateAccess operation.
//
// Evaluate a user based on their profile.
//
// POST /access/v1/evaluation
func (UnimplementedHandler) EvaluateAccess(ctx context.Context, req *EvaluateAccessReq) (r *EvaluateAccessOK, _ error) {
	return r, ht.ErrNotImplemented
}
