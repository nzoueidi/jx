// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	versioned0 "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	"github.com/petergtz/pegomock"
)

func AnyVersioned0Interface() versioned0.Interface {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(versioned0.Interface))(nil)).Elem()))
	var nullValue versioned0.Interface
	return nullValue
}

func EqVersioned0Interface(value versioned0.Interface) versioned0.Interface {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue versioned0.Interface
	return nullValue
}
