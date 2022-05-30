package wasmbindings

import sdk "github.com/cosmos/cosmos-sdk/types"

type OsmosisMsg struct {
	/// Contracts can create denoms, namespaced under the contract's address.
	/// A contract may create any number of independent sub-denoms.
	CreateDenom *CreateDenom `json:"create_denom,omitempty"`
	/// Contracts can change the admin of a denom that they are the admin of.
	ChangeAdmin *ChangeAdmin `json:"change_admin,omitempty"`
	/// Contracts can mint native tokens for an existing factory denom
	/// that they are the admin of.
	MintTokens *MintTokens `json:"mint_tokens,omitempty"`
	/// Contracts can burn native tokens for an existing factory denom
	/// that they are the admin of.
	/// Currently, the burn from address must be the admin contract.
	BurnTokens *BurnTokens `json:"burn_tokens,omitempty"`
	/// Swap over one or more pools
	Swap *SwapMsg `json:"swap,omitempty"`
	/// Add liquidity to a specified pool with only one of the required assets (i.e. Join pool 1 (50/50 ATOM-OSMO) with just ATOM).
	JoinSwapExactAmountIn *JoinSwapExactAmountIn `json:"join_swap_exact_amount_in,omitempty"`
}

/// CreateDenom creates a new factory denom, of denomination:
/// factory/{creating contract address}/{Subdenom}
/// Subdenom can be of length at most 44 characters, in [0-9a-zA-Z./]
/// The (creating contract address, subdenom) pair must be unique.
/// The created denom's admin is the creating contract address,
/// but this admin can be changed using the ChangeAdmin binding.
type CreateDenom struct {
	Subdenom string `json:"subdenom"`
}

/// ChangeAdmin changes the admin for a factory denom.
/// If the NewAdminAddress is empty, the denom has no admin.
type ChangeAdmin struct {
	Denom           string `json:"denom"`
	NewAdminAddress string `json:"new_admin_address"`
}

type MintTokens struct {
	Denom         string  `json:"denom"`
	Amount        sdk.Int `json:"amount"`
	MintToAddress string  `json:"mint_to_address"`
}

type BurnTokens struct {
	Denom  string  `json:"denom"`
	Amount sdk.Int `json:"amount"`
	// BurnFromAddress must be set to "" for now.
	BurnFromAddress string `json:"burn_from_address"`
}

type SwapMsg struct {
	First  Swap                `json:"first"`
	Route  []Step              `json:"route"`
	Amount SwapAmountWithLimit `json:"amount"`
}

type JoinSwapExactAmountIn struct {
	PoolId            uint64   `json:"pool_id"`
	ShareOutMinAmount sdk.Int  `json:"share_out_min_amount"`
	TokenIn           sdk.Coin `json:"token_in"`
}
