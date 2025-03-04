package gameserver

import (
	. "github.com/assemblaj/purego-steamworks"
)

const (
	GSClientApproveID                     SteamCallbackID = 201
	GSClientDenyID                        SteamCallbackID = 202
	GSClientKickID                        SteamCallbackID = 203
	GSClientAchievementStatusID           SteamCallbackID = 206
	GSPolicyResponseID                    SteamCallbackID = 115
	GSGameplayStatsID                     SteamCallbackID = 207
	GSClientGroupStatusID                 SteamCallbackID = 208
	GSReputationID                        SteamCallbackID = 209
	AssociateWithClanResultID             SteamCallbackID = 210
	ComputeNewPlayerCompatibilityResultID SteamCallbackID = 211
)

type GSClientApprove struct {
	SteamID      CSteamID
	OwnerSteamID CSteamID
}
type GSClientDeny struct {
	SteamID      CSteamID
	DenyReason   EDenyReason
	OptionalText [128]byte
}
type GSClientKick struct {
	SteamID    CSteamID
	DenyReason EDenyReason
}
type GSClientAchievementStatus struct {
	SteamID     uint64
	Achievement [128]byte
	Unlocked    bool
}
type GSPolicyResponse struct {
	Secure uint8
}
type GSGameplayStats struct {
	Result             EResult
	Rank               int32
	TotalConnects      uint32
	TotalMinutesPlayed uint32
}
type GSClientGroupStatus struct {
	SteamIDUser  CSteamID
	SteamIDGroup CSteamID
	Member       bool
	Officer      bool
}
type GSReputation struct {
	Result          EResult
	ReputationScore uint32
	Banned          bool
	BannedIP        uint32
	BannedPort      uint16
	BannedGameID    uint64
	BanExpires      uint32
}
type AssociateWithClanResult struct {
	Result EResult
}
type ComputeNewPlayerCompatibilityResult struct {
	Result                           EResult
	PlayersThatDontLikeCandidate     int32
	PlayersThatCandidateDoesntLike   int32
	ClanPlayersThatDontLikeCandidate int32
	SteamIDCandidate                 CSteamID
}

func (cb GSClientApprove) CallbackID() SteamCallbackID {
	return GSClientApproveID
}

func (cb GSClientApprove) Name() string {
	return "GSClientApprove"
}

func (cb GSClientApprove) String() string {
	return CallbackString(cb)
}

func (cb GSClientDeny) CallbackID() SteamCallbackID {
	return GSClientDenyID
}

func (cb GSClientDeny) Name() string {
	return "GSClientDeny"
}

func (cb GSClientDeny) String() string {
	return CallbackString(cb)
}

func (cb GSClientKick) CallbackID() SteamCallbackID {
	return GSClientKickID
}

func (cb GSClientKick) Name() string {
	return "GSClientKick"
}

func (cb GSClientKick) String() string {
	return CallbackString(cb)
}

func (cb GSClientAchievementStatus) CallbackID() SteamCallbackID {
	return GSClientAchievementStatusID
}

func (cb GSClientAchievementStatus) Name() string {
	return "GSClientAchievementStatus"
}

func (cb GSClientAchievementStatus) String() string {
	return CallbackString(cb)
}

func (cb GSPolicyResponse) CallbackID() SteamCallbackID {
	return GSPolicyResponseID
}

func (cb GSPolicyResponse) Name() string {
	return "GSPolicyResponse"
}

func (cb GSPolicyResponse) String() string {
	return CallbackString(cb)
}

func (cb GSGameplayStats) CallbackID() SteamCallbackID {
	return GSGameplayStatsID
}

func (cb GSGameplayStats) Name() string {
	return "GSGameplayStats"
}

func (cb GSGameplayStats) String() string {
	return CallbackString(cb)
}

func (cb GSClientGroupStatus) CallbackID() SteamCallbackID {
	return GSClientGroupStatusID
}

func (cb GSClientGroupStatus) Name() string {
	return "GSClientGroupStatus"
}

func (cb GSClientGroupStatus) String() string {
	return CallbackString(cb)
}

func (cb GSReputation) CallbackID() SteamCallbackID {
	return GSReputationID
}

func (cb GSReputation) Name() string {
	return "GSReputation"
}

func (cb GSReputation) String() string {
	return CallbackString(cb)
}

func (cb AssociateWithClanResult) CallbackID() SteamCallbackID {
	return AssociateWithClanResultID
}

func (cb AssociateWithClanResult) Name() string {
	return "AssociateWithClanResult"
}

func (cb AssociateWithClanResult) String() string {
	return CallbackString(cb)
}

func (cb ComputeNewPlayerCompatibilityResult) CallbackID() SteamCallbackID {
	return ComputeNewPlayerCompatibilityResultID
}

func (cb ComputeNewPlayerCompatibilityResult) Name() string {
	return "ComputeNewPlayerCompatibilityResult"
}

func (cb ComputeNewPlayerCompatibilityResult) String() string {
	return CallbackString(cb)
}
