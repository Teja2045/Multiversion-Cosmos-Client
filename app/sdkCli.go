package app

import (
	"context"
	"fmt"
	"log"

	tmserviceCli "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	mint "github.com/cosmos/cosmos-sdk/x/mint/types"
	slash "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/pkg/errors"
)

const limit = 200

type queryClient struct {
	bankCl  bank.QueryClient
	distCl  distribution.QueryClient
	mintCl  mint.QueryClient
	stkCl   stakingTypes.QueryClient
	txCl    tx.ServiceClient
	slashCl slash.QueryClient
	// tag
	gov   *GovernanceClient
	count int64
	tmCli tmserviceCli.ServiceClient
}

func NewClient(address string, enc EncodingConfig) (QueryClient, error) {

	//tag
	conn, err := GrpcDial(address, enc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to Service")
	}

	cc := queryClient{
		bankCl:  bank.NewQueryClient(conn),
		distCl:  distribution.NewQueryClient(conn),
		mintCl:  mint.NewQueryClient(conn),
		stkCl:   stakingTypes.NewQueryClient(conn),
		txCl:    tx.NewServiceClient(conn),
		slashCl: slash.NewQueryClient(conn),

		// tag
		gov:   NewGovernanceClient(conn),
		tmCli: tmserviceCli.NewServiceClient(conn),
	}

	return &cc, nil
}

/// --------------TRANSACTION MODULE RPC CALLS -----------------------

func (c *queryClient) GetTransactionData(txHash string) (TxInfo, error) {
	c.count++
	resp, err := c.txCl.GetTx(context.Background(), &tx.GetTxRequest{Hash: txHash})
	if err != nil {
		return TxInfo{}, err
	}

	return TxInfo{Tx: *resp.Tx, TxResponse: *resp.TxResponse}, nil
}

/// ------------------BANK MODULE RPC CALLS---------------------------

func (c *queryClient) GetAccountBalances(address string) ([]sdkTypes.Coin, error) {
	c.count++
	req := bank.QueryAllBalancesRequest{Address: address}
	resp, err := c.bankCl.AllBalances(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	// log.Println("result:", resp.Balances)
	return resp.Balances, nil
}

func (c *queryClient) GetAllDenomsMetadata(key string) (*bank.QueryDenomsMetadataResponse,
	error) {
	reqObj := bank.QueryDenomsMetadataRequest{Pagination: &query.PageRequest{Key: []byte(key),
		Limit:      limit,
		CountTotal: true}}
	resp, err := c.bankCl.DenomsMetadata(context.Background(), &reqObj)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetDenomSupply(denom string) (*bank.QuerySupplyOfResponse,
	error) {
	resp, err := c.bankCl.SupplyOf(context.Background(),
		&bank.QuerySupplyOfRequest{Denom: denom})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetTotalSupply(key string) ([]sdkTypes.Coin, error) {
	c.count++
	in := bank.QueryTotalSupplyRequest{
		Pagination: &query.PageRequest{Key: []byte(key),
			Limit: limit,
		}}
	resp, err := c.bankCl.TotalSupply(context.Background(), &in)
	if err != nil {
		return nil, err
	}

	return resp.Supply, nil
}

/// -------------------STAKING MODULE RPC CALLS ----------------------

func (c *queryClient) GetValidatorsList(status string, key string) (*stakingTypes.QueryValidatorsResponse, error) {
	c.count++
	resp, err := c.stkCl.Validators(context.Background(), &stakingTypes.QueryValidatorsRequest{
		Pagination: &query.PageRequest{
			Key:        []byte(key),
			Limit:      limit,
			CountTotal: true,
			Reverse:    true,
		},
		Status: status})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetDelegationInfo() {} // TODO: Figure out the RPC call to make

func (c *queryClient) GetStakingPool() (stakingTypes.Pool, error) {
	c.count++
	resp, err := c.stkCl.Pool(context.Background(), &stakingTypes.QueryPoolRequest{})
	if err != nil {
		return stakingTypes.Pool{}, err
	}

	return resp.Pool, nil
}

func (c *queryClient) GetAddressDelegations() {} // TODO: figure out the RPC call to make

func (c *queryClient) GetDelegatorUnbondingDelegations(
	delegatorAddr string) (stakingTypes.UnbondingDelegations, error) {
	c.count++
	resp, err := c.stkCl.DelegatorUnbondingDelegations(context.Background(),
		&stakingTypes.QueryDelegatorUnbondingDelegationsRequest{
			DelegatorAddr: delegatorAddr,
			Pagination:    nil,
		})

	if err != nil {
		return nil, err
	}

	return resp.UnbondingResponses, nil
}

func (c *queryClient) GetValidatorDelegatorDels(validatorAddr string,
	delegatorAddr string) (*stakingTypes.DelegationResponse, error) {
	c.count++

	resp, err := c.stkCl.Delegation(context.Background(), &stakingTypes.QueryDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})

	if err != nil {
		return nil, err
	}

	return resp.DelegationResponse, nil
}

func (c *queryClient) GetValidatorDelegatorUnDels(validatorAddr string,
	delegatorAddr string) (*stakingTypes.QueryUnbondingDelegationResponse, error) {
	c.count++

	resp, err := c.stkCl.UnbondingDelegation(context.Background(), &stakingTypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: delegatorAddr,
		ValidatorAddr: validatorAddr,
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetDelegatorDelegations(delegatorAddr string) (stakingTypes.DelegationResponses, error) {
	c.count++

	resp, err := c.stkCl.DelegatorDelegations(context.Background(), &stakingTypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: delegatorAddr,
		Pagination: &query.PageRequest{
			Limit: limit,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.DelegationResponses, nil
}

/// ---------------- DISTRIBUTION MODULE RPC CALLS -------------------

func (c *queryClient) GetDelegatorRewards(delegatorAddr string) (TotalDelegationRewards, error) {
	c.count++
	resp, err := c.distCl.DelegationTotalRewards(context.Background(),
		&distribution.QueryDelegationTotalRewardsRequest{DelegatorAddress: delegatorAddr})
	if err != nil {
		return TotalDelegationRewards{}, err
	}

	return TotalDelegationRewards{Rewards: resp.Rewards, Total: resp.Total}, nil
}

func (c *queryClient) GetValidatorDistributions(address string) {} // TODO: figure out the RPC call to make

func (c *queryClient) GetDelegatorRewardsOfValidator(delegatorAddr, validatorAddr string) ([]sdkTypes.DecCoin, error) {
	c.count++
	req := distribution.QueryDelegationRewardsRequest{DelegatorAddress: delegatorAddr, ValidatorAddress: validatorAddr}
	resp, err := c.distCl.DelegationRewards(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return resp.Rewards, nil
}

/// ------------------ GOVERNANCE MODULE RPC CALLS -------------------
// implemented in respective versions (tag)

/// --------------------- MINT MODULE RPC CALLS ----------------------

func (c *queryClient) GetInflation() (sdkTypes.Dec, error) {
	c.count++
	resp, err := c.mintCl.Inflation(context.Background(), &mint.QueryInflationRequest{})
	if err != nil {
		log.Println("Error while getting inflation-=================", err)
		return sdkTypes.Dec{}, err
	}

	return resp.Inflation, nil
}

/// ---------------------- GRPC CALL COUNTER -------------------------

func (c *queryClient) GetRequestCount() int64 {
	return c.count
}

// get tx

func (c *queryClient) GetTx(hash string) (*tx.GetTxResponse, error) {
	resp, err := c.txCl.GetTx(context.Background(), &tx.GetTxRequest{Hash: hash})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetDistributionParams() (*distribution.QueryParamsResponse, error) {
	resp, err := c.distCl.Params(context.Background(), &distribution.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *queryClient) GetMintParams() (*mint.QueryParamsResponse, error) {
	resp, err := c.mintCl.Params(context.Background(), &mint.QueryParamsRequest{})
	if err != nil {
		return nil, fmt.Errorf("mint params: %w", err)
	}

	return resp, nil
}

func (c *queryClient) GetSlashingParams() (*slash.QueryParamsResponse, error) {
	resp, err := c.slashCl.Params(context.Background(), &slash.QueryParamsRequest{})
	if err != nil {
		return nil, fmt.Errorf("slash params : %w", err)
	}

	return resp, nil
}

func (c *queryClient) GetStakingParams() (*stakingTypes.QueryParamsResponse, error) {
	resp, err := c.stkCl.Params(context.Background(), &stakingTypes.QueryParamsRequest{})
	if err != nil {
		return nil, fmt.Errorf("stake params: %w", err)
	}

	return resp, nil
}

func (c *queryClient) GetNodeInfo() (*tmserviceCli.GetNodeInfoResponse,
	error) {
	resp, err := c.tmCli.GetNodeInfo(context.Background(),
		&tmserviceCli.GetNodeInfoRequest{})
	if err != nil {
		return nil, fmt.Errorf("error while getting node info: %w", err)
	}

	return resp, nil
}

func (c *queryClient) GetCommunityPool() (*distribution.QueryCommunityPoolResponse,
	error) {
	resp, err := c.distCl.CommunityPool(context.Background(),
		&distribution.QueryCommunityPoolRequest{})
	if err != nil {
		return nil, fmt.Errorf("error while getting node info: %w", err)
	}

	return resp, nil
}
