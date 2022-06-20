package auth

import (
	"fmt"

	"github.com/casbin/casbin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Authoriser struct {
	enforcer *casbin.Enforcer
}

func New(model, policy string) *Authoriser {
	enforcer := casbin.NewEnforcer(model, policy)
	return &Authoriser{
		enforcer: enforcer,
	}
}

func (a *Authoriser) Authorise(subject, object, action string) error {
	if !a.enforcer.Enforce(subject, object, action) {
		msg := fmt.Sprintf("%s not permitted to %s to %s", subject, action, object)
		st := status.New(codes.PermissionDenied, msg)
		return st.Err()
	}
	return nil
}
