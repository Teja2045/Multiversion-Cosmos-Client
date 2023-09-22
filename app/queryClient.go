package app

import (
	tmserviceCli "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	bankcl "github.com/cosmos/cosmos-sdk/x/bank/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
	slash "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type TotalDelegationRewards struct {
	Rewards []distribution.DelegationDelegatorReward
	Total   []sdkTypes.DecCoin
}

type TxInfo struct {
	Tx         tx.Tx
	TxResponse sdkTypes.TxResponse
}

// QueryClient is a consolidated client for making RPC calls to Cosmos.
type QueryClient interface {
	// TX RPC
	GetTransactionData(txHash string) (TxInfo, error)
	// Bank RPC
	GetAccountBalances(address string) ([]sdkTypes.Coin, error)
	GetAllDenomsMetadata(key string) (*bankcl.QueryDenomsMetadataResponse, error)
	GetTotalSupply(key string) ([]sdkTypes.Coin, error)
	GetDenomSupply(denom string) (*bankcl.QuerySupplyOfResponse, error)
	GetValidatorsList(status string, key string) (
		*stakingTypes.QueryValidatorsResponse, error)
	GetStakingPool() (stakingTypes.Pool, error)
	GetDelegatorDelegations(delegatorAddr string) (
		stakingTypes.DelegationResponses, error)
	GetValidatorDelegatorDels(validatorAddr string,
		delegatorAddr string) (*stakingTypes.DelegationResponse, error)
	GetValidatorDelegatorUnDels(validatorAddr string,
		delegatorAddr string) (*stakingTypes.QueryUnbondingDelegationResponse, error)

	GetDelegatorUnbondingDelegations(delegatorAddr string) (
		stakingTypes.UnbondingDelegations, error)
	GetDelegatorRewards(delegatorAddr string) (TotalDelegationRewards,
		error)
	GetDelegatorRewardsOfValidator(delegatorAddr,
		validatorAddr string) ([]sdkTypes.DecCoin, error)
	// Governance RPC
	// depends on build tag
	GovernanceInterface
	// Minting RPC
	GetInflation() (sdkTypes.Dec, error)
	// GRPC requests counter
	GetRequestCount() int64
	GetTx(hash string) (*tx.GetTxResponse, error)
	GetDistributionParams() (*distribution.QueryParamsResponse, error)
	GetMintParams() (*mint.QueryParamsResponse, error)
	GetSlashingParams() (*slash.QueryParamsResponse, error)

	GetNodeInfo() (*tmserviceCli.GetNodeInfoResponse, error)
	GetCommunityPool() (*distribution.QueryCommunityPoolResponse, error)
}
