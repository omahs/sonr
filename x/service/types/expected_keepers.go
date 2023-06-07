package types

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/go-webauthn/webauthn/protocol"
	identitytypes "github.com/sonrhq/core/x/identity/types"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
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
	AssignIdentity(credential *identitytypes.VerificationMethod, primary vaulttypes.Account, alias string, accounts ...vaulttypes.Account) (*identitytypes.DIDDocument, error)
	CheckAlsoKnownAs(ctx sdk.Context, alias string) error
	GetIdentityByPrimaryAlias(ctx sdk.Context, alias string) (val identitytypes.Identification, found bool)
	ResolveIdentityByPrimaryAlias(ctx sdk.Context, alias string) (val identitytypes.DIDDocument, err error)
	ResolveIdentity(ctx sdk.Context, did string) (identitytypes.DIDDocument, error)
}

// VaultKeeper defines the expected interface for managing Keys on IPFS Vaults
type VaultKeeper interface {
	// Methods imported from vault should be defined here
	AssignVault(ctx sdk.Context, ucw uint64, credential *WebauthnCredential) (vaulttypes.Account, error)
	GetClaimableWallet(ctx sdk.Context, id uint64) (val vaulttypes.ClaimableWallet, found bool)
	NextUnclaimedWallet(ctx sdk.Context) (*vaulttypes.ClaimableWallet, protocol.URLEncodedBase64, error)
	GetAccount(accDid string) (vaulttypes.Account, error)
}
