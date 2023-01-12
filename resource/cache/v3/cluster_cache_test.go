package cachev3_test

import (
	"fmt"
	"testing"

	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	cachev3 "github.com/yshdzw/go-control-plane-kit/resource/cache/v3"
)

func TestClusterCacheAllMethod(t *testing.T) {
	var (
		clusterCache = cachev3.NewCache[*cachev3.Cluster]()
		err          error
		cluster      *cachev3.Cluster
	)

	// Set
	clusterCache.Set(&cachev3.Cluster{Cluster: &clusterv3.Cluster{Name: "cluster-01"}})
	cluster, err = clusterCache.GetByName("cluster-01")
	if err != nil {
		t.Fatal(err)
	} else {
		cluster.Name = "changed name"
		t.Log(cluster)

		c, _ := clusterCache.GetByName("cluster-01")
		t.Log(c)
	}

	// SetList
	clusterCache.SetList([]*cachev3.Cluster{
		{
			Cluster: &clusterv3.Cluster{Name: "cluster-02"},
		}, {
			Cluster: &clusterv3.Cluster{Name: "cluster-03"},
		}, {
			Cluster: &clusterv3.Cluster{Name: "cluster-04"},
		}, {
			Cluster: &clusterv3.Cluster{Name: "cluster-05"},
		},
	})
	cluster, err = clusterCache.GetByName("cluster-02")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(cluster)
		clusterCache.PrintNames()
	}

	// GetData
	items, err := clusterCache.GetData(cluster)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(items)
	}

	// GetList
	list, err := clusterCache.GetList(cluster)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(list)
	}

	// GetResources
	res, err := clusterCache.GetResources(cluster)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}

	// Clone
	cacheTmp, err := clusterCache.Clone()
	if err != nil {
		t.Fatal(err)
	} else {
		cacheTmp.PrintNames()
		cacheTmp2, _ := cacheTmp.Clone(cluster)
		cacheTmp2.PrintNames()
	}

	// Delete
	cacheTmp.Delete(cluster)
	cacheTmp.PrintNames()
	clusterCache.PrintNames()
	cacheTmp.DeleteList([]*cachev3.Cluster{
		{
			Cluster: &clusterv3.Cluster{Name: "cluster-03"},
		}, {
			Cluster: &clusterv3.Cluster{Name: "cluster-04"},
		},
	})
	cacheTmp.PrintNames()
	cacheTmp.Clean()
	cacheTmp.PrintNames()
	clusterCache.PrintNames()
}

// goos: linux
// goarch: amd64
// cpu: 12th Gen Intel(R) Core(TM) i5-12400
// BenchmarkSet
// BenchmarkSet-4           1000000              1642 ns/op             667 B/op          4 allocs/op
func BenchmarkSet(b *testing.B) {
	clusterCache := cachev3.NewCache[*cachev3.Cluster]()
	for i := 0; i < b.N; i++ {
		clusterCache.Set(&cachev3.Cluster{Cluster: &clusterv3.Cluster{Name: fmt.Sprintf("cluster-%d", i)}})
	}
}
