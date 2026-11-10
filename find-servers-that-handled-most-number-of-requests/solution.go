// Package findserversthathandledmostnumberofrequests solves leetcode problem
package findserversthathandledmostnumberofrequests

import (
	"container/heap"
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func busiestServers(k int, arrival []int, load []int) []int {
	availableServers := redblacktree.NewWithIntComparator()
	for server := range k {
		availableServers.Put(server, struct{}{})
	}

	busyServers := busyServers{}

	serverReqCout := make([]int, k)

	maxRequestCount := math.MinInt32

	for req := range len(arrival) {
		for len(busyServers) > 0 && busyServers[0].BusyTill <= arrival[req] {
			availableServers.Put(busyServers[0].ServerIndex, struct{}{})
			heap.Pop(&busyServers)
		}

		if availableServers.Size() == 0 {
			// drop reque
			continue
		}

		preferredServer := req % k
		node, ok := availableServers.Ceiling(preferredServer)

		var allocatedServer int
		if ok {
			allocatedServer = node.Key.(int)
		} else {
			allocatedServer = availableServers.Left().Key.(int)
		}

		heap.Push(&busyServers, busyServer{allocatedServer, arrival[req] + load[req]})
		availableServers.Remove(allocatedServer)
		serverReqCout[allocatedServer] += 1

		if maxRequestCount < serverReqCout[allocatedServer] {
			maxRequestCount = serverReqCout[allocatedServer]
		}

	}

	var busiestServers []int
	for server, reqCount := range serverReqCout {
		if reqCount == maxRequestCount {
			busiestServers = append(busiestServers, server)
		}
	}

	return busiestServers
}

type busyServer struct {
	ServerIndex int
	BusyTill    int
}

type busyServers []busyServer

func (servers busyServers) Len() int {
	return len(servers)
}

func (servers busyServers) Less(i, j int) bool {
	return servers[i].BusyTill < servers[j].BusyTill
}

func (servers busyServers) Swap(i, j int) {
	servers[i], servers[j] = servers[j], servers[i]
}

func (servers *busyServers) Push(x any) {
	server := x.(busyServer)
	*servers = append(*servers, server)
}

func (servers *busyServers) Pop() any {
	old := *servers
	n := len(old)
	server := old[n-1]
	*servers = old[0 : n-1]
	return server
}
