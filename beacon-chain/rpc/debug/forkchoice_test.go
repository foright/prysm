package debug

import (
	"context"
	"testing"

	ptypes "github.com/gogo/protobuf/types"
	mock "github.com/prysmaticlabs/prysm/beacon-chain/blockchain/testing"
	"github.com/prysmaticlabs/prysm/beacon-chain/forkchoice/protoarray"
	"github.com/prysmaticlabs/prysm/shared/testutil/assert"
	"github.com/prysmaticlabs/prysm/shared/testutil/require"
)

func TestServer_GetForkChoice(t *testing.T) {
	store := &protoarray.Store{
		PruneThreshold: 1,
		JustifiedEpoch: 2,
		FinalizedEpoch: 3,
		Nodes:          []*protoarray.Node{{Slot: 4, Root: [32]byte{'a'}}},
	}

	bs := &Server{
		HeadFetcher: &mock.ChainService{ForkChoiceStore: store},
	}

	res, err := bs.GetProtoArrayForkChoice(context.Background(), &ptypes.Empty{})
	require.NoError(t, err)
	assert.Equal(t, store.PruneThreshold, res.PruneThreshold, "Did not get wanted prune threshold")
	assert.Equal(t, store.JustifiedEpoch, res.JustifiedEpoch, "Did not get wanted justified epoch")
	assert.Equal(t, store.FinalizedEpoch, res.FinalizedEpoch, "Did not get wanted finalized epoch")
	assert.Equal(t, store.Nodes[0].Slot, res.ProtoArrayNodes[0].Slot, "Did not get wanted node slot")
}
