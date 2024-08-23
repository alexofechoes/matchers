package matchers

import (
	"context"
)

func AnyContext() AnyContextMatcher {
	return AnyContextMatcher{}
}

type AnyContextMatcher struct{}

func (AnyContextMatcher) Matches(x interface{}) bool {
	if _, ok := x.(context.Context); ok {
		return true
	}
	return false
}

func (AnyContextMatcher) String() string {
	return "is context"
}
