//go:build 45
// +build 45

package app

import (
	"context"
	"log"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// governance module queries
type GovernanceInterface interface {
	GetProposalInfo(proposalID uint64) (gov.Proposal, error)
	//GetProposalV1Info(proposalID uint64) (*govTypes1.QueryProposalResponse, error)
	GetProposalTally(proposalID uint64) (gov.TallyResult, error)
	GetProposalVotes(proposalID uint64) ([]gov.Vote, error)
	GetProposalDeposits(proposalID uint64) ([]gov.Deposit, error)
	GetAllProposals() ([]gov.Proposal, error)
	GetGovParams() (*gov.QueryParamsResponse, error)
}

type GovernanceClient struct {
	govCl gov.QueryClient
}

func NewGovernanceClient(conn *grpc.ClientConn) *GovernanceClient {
	return &GovernanceClient{
		govCl: gov.NewQueryClient(conn),
	}
}

func GrpcDial(address string, enc EncodingConfig) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions())
}

// governance module queries implementation
func (c *queryClient) GetProposalInfo(proposalID uint64) (gov.Proposal, error) {
	c.count++
	resp, err := c.gov.govCl.Proposal(context.Background(), &gov.QueryProposalRequest{ProposalId: proposalID})
	if err != nil {
		return gov.Proposal{}, err
	}

	return resp.Proposal, nil
}

func (c *queryClient) GetProposalTally(proposalID uint64) (gov.TallyResult, error) {
	c.count++
	resp, err := c.gov.govCl.TallyResult(context.Background(), &gov.QueryTallyResultRequest{ProposalId: proposalID})
	if err != nil {
		return gov.EmptyTallyResult(), err
	}

	return resp.Tally, nil
}

func (c *queryClient) GetProposalVotes(proposalID uint64) ([]gov.Vote, error) {
	c.count++
	resp, err := c.gov.govCl.Votes(context.Background(), &gov.QueryVotesRequest{ProposalId: proposalID})
	if err != nil {
		return nil, err
	}

	return resp.Votes, nil
}

func (c *queryClient) GetProposalDeposits(proposalID uint64) ([]gov.Deposit, error) {
	c.count++
	resp, err := c.gov.govCl.Deposits(context.Background(), &gov.QueryDepositsRequest{ProposalId: proposalID})
	if err != nil {
		return nil, err
	}

	return resp.Deposits, nil
}

func (c *queryClient) GetAllProposals() ([]gov.Proposal, error) {
	// NOTE: The QueryProposalsRequest can take certain parameters, but I didn't
	// find any reference in the NodeJS implementation that any params are being
	// passed. Need to confirm usage.
	c.count++
	resp, err := c.gov.govCl.Proposals(context.Background(), &gov.QueryProposalsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Proposals, nil
}

func (c *queryClient) GetGovParams() (*gov.QueryParamsResponse, error) {
	var response gov.QueryParamsResponse
	resp1, err := c.gov.govCl.Params(context.Background(), &gov.QueryParamsRequest{
		ParamsType: "tallying",
	})

	log.Println("tallying ==== ", resp1)

	if err != nil {
		return nil, err
	}

	response.TallyParams = resp1.TallyParams

	resp2, err := c.gov.govCl.Params(context.Background(), &gov.QueryParamsRequest{
		ParamsType: "voting",
	})
	if err != nil {
		return nil, err
	}

	log.Println("voting ==== ", resp2)
	response.VotingParams = resp2.VotingParams

	resp3, err := c.gov.govCl.Params(context.Background(), &gov.QueryParamsRequest{
		ParamsType: "deposit",
	})
	if err != nil {
		return nil, err
	}

	log.Println("deposit ==== ", resp3)

	response.DepositParams = resp3.DepositParams

	return &response, nil
}

// codec
func RegisterCodecForGov(encConfig *EncodingConfig) {
	gov.RegisterLegacyAminoCodec(encConfig.Amino)
	gov.RegisterInterfaces(encConfig.InterfaceRegistry)
}
