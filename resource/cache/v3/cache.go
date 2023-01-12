package cachev3

import (
	"fmt"
	"sync"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	resourcev3 "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
)

type cacheT interface {
	*Endpoint | *Cluster | *Route | *ScopedRoute | *VirtualHost | *Listener | *Secret | *ExtensionConfig | *Runtime

	// 自带方法
	GetName() string

	// 扩展方法
	DeepCopy() (any, error)
	GetResource() types.Resource
	GetType() resourcev3.Type
}

type Cache[T cacheT] interface {
	Set(T)
	SetList([]T)

	Delete(T)
	DeleteList([]T)
	Clean()

	// 以Get开头的方法都具有Clone特性，对外部修改不可见。
	GetData(...T) (map[string]T, error)
	GetByName(string) (T, error)
	GetList(...T) ([]T, error)
	GetResources(...T) ([]types.Resource, error)

	// 当不传递参数时等效于克隆自身
	Clone(...T) (Cache[T], error)

	PrintNames()
}

type cache[T cacheT] struct {
	mu   sync.RWMutex
	data map[string]T // resource name: resource
}

func NewCache[T cacheT]() Cache[T] {
	return &cache[T]{data: make(map[string]T)}
}

func (c *cache[T]) Set(item T) {
	if item != nil {
		c.mu.Lock()
		c.data[item.GetName()] = item
		c.mu.Unlock()
	}
}

func (c *cache[T]) SetList(list []T) {
	if len(list) > 0 {
		c.mu.Lock()
		for _, item := range list {
			if item != nil {
				c.data[item.GetName()] = item
			}
		}
		c.mu.Unlock()
	}
}

func (c *cache[T]) Delete(item T) {
	if item != nil {
		c.mu.Lock()
		delete(c.data, item.GetName())
		c.mu.Unlock()
	}
}

func (c *cache[T]) DeleteList(list []T) {
	if len(list) > 0 {
		c.mu.Lock()
		for _, item := range list {
			if item != nil {
				delete(c.data, item.GetName())
			}
		}
		c.mu.Unlock()
	}
}

func (c *cache[T]) Clean() {
	if len(c.data) == 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = make(map[string]T)
}

func (c *cache[T]) GetData(removes ...T) (map[string]T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	rmNames := map[string]struct{}{}
	for _, rm := range removes {
		rmNames[rm.GetName()] = struct{}{}
	}

	items := map[string]T{}
	for name, item := range c.data {
		if _, ok := rmNames[name]; !ok {
			any, err := item.DeepCopy()
			if err != nil {
				return nil, err
			} else {
				items[name] = any.(T)
			}
		}
	}
	return items, nil
}

func (c *cache[T]) GetByName(name string) (T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	any, err := c.data[name].DeepCopy()
	if err != nil {
		return nil, err
	}
	return any.(T), nil
}

func (c *cache[T]) GetList(removes ...T) ([]T, error) {
	list := []T{}
	items, err := c.GetData(removes...)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		list = append(list, item)
	}
	return list, nil
}

func (c *cache[T]) GetResources(removes ...T) ([]types.Resource, error) {
	list := []types.Resource{}
	items, err := c.GetData(removes...)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		list = append(list, item.GetResource())
	}
	return list, nil
}

func (c *cache[T]) Clone(removes ...T) (Cache[T], error) {
	data, err := c.GetData(removes...)
	if err != nil {
		return nil, err
	}
	return &cache[T]{data: data}, nil
}

func (c *cache[T]) PrintNames() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	names := []string{}
	for name := range c.data {
		names = append(names, name)
	}
	fmt.Printf("%T %+v\n", c, names)
}
