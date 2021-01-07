package mongoshard

import "testing"

func TestRouter(t *testing.T) {
	testCases := []struct {
		name           string
		shardedNumber  int
		nodesInput     []*node
		expectedRouter map[int]*node
	}{
		{"Without nodes", 5, []*node{}, map[int]*node{}},
		{"Less nodes than shards", 3, []*node{{id: 0}}, map[int]*node{0: {id: 0}, 1: {id: 0}, 2: {id: 0}}},
		{"More nodes than shards", 1, []*node{{id: 0}, {id: 1}, {id: 2}}, map[int]*node{0: {id: 0}}},
		{"Equal nodes and shards", 3, []*node{{id: 0}, {id: 1}, {id: 2}}, map[int]*node{0: {id: 0}, 1: {id: 1}, 2: {id: 2}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := make(map[int]*node)
			route(tc.nodesInput, router, tc.shardedNumber)
			for idShard, nodeExpected := range tc.expectedRouter {
				nodeOuptut := router[idShard]
				if nodeOuptut == nil || nodeOuptut.id != nodeExpected.id {
					t.Errorf("%v | ID SHARD: %v | EXPECTED: nodeID %v | OUTPUT: nodeID %v", tc.name, idShard, nodeExpected.id, nodeOuptut)
				}
			}
		})
	}
}
