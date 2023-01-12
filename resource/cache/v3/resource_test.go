package cachev3_test

import (
	"testing"

	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	cachev3 "github.com/yshdzw/go-control-plane-kit/resource/cache/v3"
)

func TestGetType(t *testing.T) {
	e := cachev3.Endpoint{}
	if e.GetType() != resourcev3.EndpointType {
		t.Fatal(e.GetType())
	}

	c := cachev3.Cluster{}
	if c.GetType() != resourcev3.ClusterType {
		t.Fatal(c.GetType())
	}

	r := cachev3.Route{}
	if r.GetType() != resourcev3.RouteType {
		t.Fatal(r.GetType())
	}

	sr := cachev3.ScopedRoute{}
	if sr.GetType() != resourcev3.ScopedRouteType {
		t.Fatal(sr.GetType())
	}

	vh := cachev3.VirtualHost{}
	if vh.GetType() != resourcev3.VirtualHostType {
		t.Fatal(vh.GetType())
	}

	l := cachev3.Listener{}
	if l.GetType() != resourcev3.ListenerType {
		t.Fatal(l.GetType())
	}

	sec := cachev3.Secret{}
	if sec.GetType() != resourcev3.SecretType {
		t.Fatal(sec.GetType())
	}

	ec := cachev3.ExtensionConfig{}
	if ec.GetType() != resourcev3.ExtensionConfigType {
		t.Fatal(ec.GetType())
	}

	run := cachev3.Runtime{}
	if run.GetType() != resourcev3.RuntimeType {
		t.Fatal(run.GetType())
	}
}
