package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	identitytypes "github.com/sonrhq/core/x/identity/types"
)

type GroupKeeper interface {
	CreateGroup(goCtx context.Context, req *group.MsgCreateGroup) (*group.MsgCreateGroupResponse, error)
	CreateGroupPolicy(goCtx context.Context, req *group.MsgCreateGroupPolicy) (*group.MsgCreateGroupPolicyResponse, error)
	CreateGroupWithPolicy(goCtx context.Context, req *group.MsgCreateGroupWithPolicy) (*group.MsgCreateGroupWithPolicyResponse, error)
	GroupMembers(goCtx context.Context, request *group.QueryGroupMembersRequest) (*group.QueryGroupMembersResponse, error)
	GroupPolicyInfo(goCtx context.Context, request *group.QueryGroupPolicyInfoRequest) (*group.QueryGroupPolicyInfoResponse, error)
	GroupsByAdmin(goCtx context.Context, request *group.QueryGroupsByAdminRequest) (*group.QueryGroupsByAdminResponse, error)
	GroupsByMember(goCtx context.Context, request *group.QueryGroupsByMemberRequest) (*group.QueryGroupsByMemberResponse, error)
	LeaveGroup(goCtx context.Context, req *group.MsgLeaveGroup) (*group.MsgLeaveGroupResponse, error)
	Proposal(goCtx context.Context, request *group.QueryProposalRequest) (*group.QueryProposalResponse, error)
	PruneProposals(ctx sdk.Context) error
	SubmitProposal(goCtx context.Context, req *group.MsgSubmitProposal) (*group.MsgSubmitProposalResponse, error)
	UpdateGroupMembers(goCtx context.Context, req *group.MsgUpdateGroupMembers) (*group.MsgUpdateGroupMembersResponse, error)
	UpdateGroupMetadata(goCtx context.Context, req *group.MsgUpdateGroupMetadata) (*group.MsgUpdateGroupMetadataResponse, error)
	Vote(goCtx context.Context, req *group.MsgVote) (*group.MsgVoteResponse, error)
	VotesByProposal(goCtx context.Context, request *group.QueryVotesByProposalRequest) (*group.QueryVotesByProposalResponse, error)
	WithdrawProposal(goCtx context.Context, req *group.MsgWithdrawProposal) (*group.MsgWithdrawProposalResponse, error)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// IdentityKeeper defines the expected interface needed to retrieve account balances.
type IdentityKeeper interface {
	GetDidDocument(ctx sdk.Context, did string) (identitytypes.DidDocument, bool)
	HasAuthentication(ctx sdk.Context, reference string) bool
	GetAuthentication(ctx sdk.Context, reference string) (identitytypes.VerificationRelationship, bool)
	ResolveIdentity(ctx sdk.Context, did string) (identitytypes.DIDDocument, error)
}
