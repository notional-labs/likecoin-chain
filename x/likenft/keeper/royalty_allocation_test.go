package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/likecoin/likecoin-chain/v3/testutil/keeper"
	"github.com/likecoin/likecoin-chain/v3/x/likenft/types"
	"github.com/stretchr/testify/require"
)

func TestComputeRoyaltyAllocation(t *testing.T) {
	t.Run("single stakeholder", func(t *testing.T) {
		keeper, ctx := keepertest.LikenftKeeper(t)
		royaltyAmount, allocations, err := keeper.ComputeRoyaltyAllocation(ctx, 110, types.RoyaltyConfig{
			RateBasisPoints: 1000,
			Stakeholders: []types.RoyaltyStakeholder{
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
					Weight:  1,
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, uint64(11), royaltyAmount) // 110 * 10% = 11
		require.Equal(t, []types.RoyaltyAllocation{
			{
				Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
				Amount:  uint64(11), // 11 / 1 * 1 = 11
			},
		}, allocations)
	})
	t.Run("multiple stakeholder", func(t *testing.T) {
		keeper, ctx := keepertest.LikenftKeeper(t)
		royaltyAmount, allocations, err := keeper.ComputeRoyaltyAllocation(ctx, 110, types.RoyaltyConfig{
			RateBasisPoints: 1000,
			Stakeholders: []types.RoyaltyStakeholder{
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
					Weight:  1,
				},
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 1, 1}),
					Weight:  2,
				},
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 1, 1, 1}),
					Weight:  3,
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, uint64(9), royaltyAmount) // 1 + 3 + 5 = 9
		require.Subset(t, []types.RoyaltyAllocation{
			{
				Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
				Amount:  uint64(1), // 11 / 6 * 1 = 1.83... ~ 1
			},
			{
				Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 1, 1}),
				Amount:  uint64(3), // 11 / 6 * 2 = 3.66... ~ 3
			},
			{
				Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 1, 1, 1}),
				Amount:  uint64(5), // 11 / 6 * 3 = 5.5... ~ 5
			},
		}, allocations)
	})
	t.Run("no stakeholder", func(t *testing.T) {
		keeper, ctx := keepertest.LikenftKeeper(t)
		royaltyAmount, allocations, err := keeper.ComputeRoyaltyAllocation(ctx, 110, types.RoyaltyConfig{
			RateBasisPoints: 1000,
			Stakeholders:    []types.RoyaltyStakeholder{},
		})
		require.NoError(t, err)
		require.Equal(t, uint64(0), royaltyAmount)
		require.Empty(t, allocations)
	})
	t.Run("zero rate", func(t *testing.T) {
		keeper, ctx := keepertest.LikenftKeeper(t)
		royaltyAmount, allocations, err := keeper.ComputeRoyaltyAllocation(ctx, 110, types.RoyaltyConfig{
			RateBasisPoints: 0,
			Stakeholders: []types.RoyaltyStakeholder{
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
					Weight:  1,
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, uint64(0), royaltyAmount)
		require.Empty(t, allocations)
	})
	t.Run("zero weight", func(t *testing.T) {
		keeper, ctx := keepertest.LikenftKeeper(t)
		royaltyAmount, allocations, err := keeper.ComputeRoyaltyAllocation(ctx, 110, types.RoyaltyConfig{
			RateBasisPoints: 1000,
			Stakeholders: []types.RoyaltyStakeholder{
				{
					Account: sdk.AccAddress([]byte{0, 0, 0, 0, 0, 0, 0, 1}),
					Weight:  0,
				},
			},
		})
		require.NoError(t, err)
		require.Equal(t, uint64(0), royaltyAmount)
		require.Empty(t, allocations)
	})
}
