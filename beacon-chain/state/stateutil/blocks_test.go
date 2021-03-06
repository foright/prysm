package stateutil_test

import (
	"testing"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/state/stateutil"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/prysmaticlabs/prysm/shared/testutil/assert"
	"github.com/prysmaticlabs/prysm/shared/testutil/require"
)

func TestBlockRoot(t *testing.T) {
	genState, keys := testutil.DeterministicGenesisState(t, 100)
	blk, err := testutil.GenerateFullBlock(genState, keys, testutil.DefaultBlockGenConfig(), 10)
	require.NoError(t, err)
	expectedRoot, err := ssz.HashTreeRoot(blk.Block)
	require.NoError(t, err)
	receivedRoot, err := stateutil.BlockRoot(blk.Block)
	require.NoError(t, err)
	require.Equal(t, expectedRoot, receivedRoot)
	blk, err = testutil.GenerateFullBlock(genState, keys, testutil.DefaultBlockGenConfig(), 100)
	require.NoError(t, err)
	expectedRoot, err = ssz.HashTreeRoot(blk.Block)
	require.NoError(t, err)
	receivedRoot, err = stateutil.BlockRoot(blk.Block)
	require.NoError(t, err)
	require.Equal(t, expectedRoot, receivedRoot)
}

func TestBlockBodyRoot_NilIsSameAsEmpty(t *testing.T) {
	a, err := ssz.HashTreeRoot(&ethpb.BeaconBlockBody{})
	require.NoError(t, err)
	b, err := stateutil.BlockBodyRoot(nil)
	require.NoError(t, err)
	assert.Equal(t, a, b, "A nil and empty block body do not generate the same root")
}
