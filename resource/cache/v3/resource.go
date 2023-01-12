package cachev3

import (
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtimev3 "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	// 利用protojson序列化和反序列化实现DeepCopy
	protoDeepCopy = func(src, dst proto.Message) error {
		bs, err := protojson.Marshal(src)
		if err != nil {
			return err
		}
		return protojson.Unmarshal(bs, dst)
	}
)

// Resource types in xDS v3.
// APITypePrefix       = "type.googleapis.com/"
// EndpointType        = APITypePrefix + "envoy.config.endpoint.v3.ClusterLoadAssignment"
// ClusterType         = APITypePrefix + "envoy.config.cluster.v3.Cluster"
// RouteType           = APITypePrefix + "envoy.config.route.v3.RouteConfiguration"
// ScopedRouteType     = APITypePrefix + "envoy.config.route.v3.ScopedRouteConfiguration"
// VirtualHostType     = APITypePrefix + "envoy.config.route.v3.VirtualHost"
// ListenerType        = APITypePrefix + "envoy.config.listener.v3.Listener"
// SecretType          = APITypePrefix + "envoy.extensions.transport_sockets.tls.v3.Secret"
// ExtensionConfigType = APITypePrefix + "envoy.config.core.v3.TypedExtensionConfig"
// RuntimeType         = APITypePrefix + "envoy.service.runtime.v3.Runtime"
type (
	Endpoint struct {
		*endpointv3.ClusterLoadAssignment
	}
	Cluster     struct{ *clusterv3.Cluster }
	Route       struct{ *routev3.RouteConfiguration }
	ScopedRoute struct {
		*routev3.ScopedRouteConfiguration
	}
	VirtualHost     struct{ *routev3.VirtualHost }
	Listener        struct{ *listenerv3.Listener }
	Secret          struct{ *tlsv3.Secret }
	ExtensionConfig struct{ *corev3.TypedExtensionConfig }
	Runtime         struct{ *runtimev3.Runtime }
)

func (e *Endpoint) DeepCopy() (any, error) {
	var t Endpoint = Endpoint{ClusterLoadAssignment: &endpointv3.ClusterLoadAssignment{}}
	err := protoDeepCopy(e.ClusterLoadAssignment, t.ClusterLoadAssignment)
	return &t, err
}

func (e *Endpoint) GetResource() types.Resource {
	return e.ClusterLoadAssignment
}

func (e *Endpoint) GetType() resourcev3.Type {
	return resourcev3.EndpointType
}

func (c *Cluster) DeepCopy() (any, error) {
	var t Cluster = Cluster{Cluster: &clusterv3.Cluster{}}
	err := protoDeepCopy(c.Cluster, t.Cluster)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *Cluster) GetResource() types.Resource {
	return c.Cluster
}

func (c *Cluster) GetType() resourcev3.Type {
	return resourcev3.ClusterType
}

func (r *Route) DeepCopy() (any, error) {
	var t Route = Route{RouteConfiguration: &routev3.RouteConfiguration{}}
	err := protoDeepCopy(r.RouteConfiguration, t.RouteConfiguration)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Route) GetResource() types.Resource {
	return r.RouteConfiguration
}

func (r *Route) GetType() resourcev3.Type {
	return resourcev3.RouteType
}

func (sr *ScopedRoute) DeepCopy() (any, error) {
	var t ScopedRoute = ScopedRoute{ScopedRouteConfiguration: &routev3.ScopedRouteConfiguration{}}
	err := protoDeepCopy(sr.ScopedRouteConfiguration, t.ScopedRouteConfiguration)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (sr *ScopedRoute) GetResource() types.Resource {
	return sr.ScopedRouteConfiguration
}

func (sr *ScopedRoute) GetType() resourcev3.Type {
	return resourcev3.ScopedRouteType
}

func (vh *VirtualHost) DeepCopy() (any, error) {
	var t VirtualHost = VirtualHost{VirtualHost: &routev3.VirtualHost{}}
	err := protoDeepCopy(vh.VirtualHost, t.VirtualHost)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (vh *VirtualHost) GetResource() types.Resource {
	return vh.VirtualHost
}

func (vh *VirtualHost) GetType() resourcev3.Type {
	return resourcev3.VirtualHostType
}

func (l *Listener) DeepCopy() (any, error) {
	var t Listener = Listener{Listener: &listenerv3.Listener{}}
	err := protoDeepCopy(l.Listener, t.Listener)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (l *Listener) GetResource() types.Resource {
	return l.Listener
}

func (l *Listener) GetType() resourcev3.Type {
	return resourcev3.ListenerType
}

func (s *Secret) DeepCopy() (any, error) {
	var t Secret = Secret{Secret: &tlsv3.Secret{}}
	err := protoDeepCopy(s.Secret, t.Secret)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *Secret) GetResource() types.Resource {
	return s.Secret
}

func (s *Secret) GetType() resourcev3.Type {
	return resourcev3.SecretType
}

func (ec *ExtensionConfig) DeepCopy() (any, error) {
	var t ExtensionConfig = ExtensionConfig{TypedExtensionConfig: &corev3.TypedExtensionConfig{}}
	err := protoDeepCopy(ec.TypedExtensionConfig, t.TypedExtensionConfig)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (ec *ExtensionConfig) GetResource() types.Resource {
	return ec.TypedExtensionConfig
}

func (ec *ExtensionConfig) GetType() resourcev3.Type {
	return resourcev3.ExtensionConfigType
}

func (r *Runtime) DeepCopy() (any, error) {
	var t Runtime = Runtime{Runtime: &runtimev3.Runtime{}}
	err := protoDeepCopy(r.Runtime, t.Runtime)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *Runtime) GetResource() types.Resource {
	return r.Runtime
}

func (r *Runtime) GetType() resourcev3.Type {
	return resourcev3.RuntimeType
}
