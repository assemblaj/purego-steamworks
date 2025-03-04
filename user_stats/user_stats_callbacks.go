package userstats

import . "github.com/assemblaj/purego-steamworks"

const (
	UserStatsReceivedID                 SteamCallbackID = 1101
	UserStatsStoredID                   SteamCallbackID = 1102
	UserAchievementStoredID             SteamCallbackID = 1103
	LeaderboardFindResultID             SteamCallbackID = 1104
	LeaderboardScoresDownloadedID       SteamCallbackID = 1105
	LeaderboardScoreUploadedID          SteamCallbackID = 1106
	NumberOfCurrentPlayersID            SteamCallbackID = 1107
	UserStatsUnloadedID                 SteamCallbackID = 1108
	UserAchievementIconFetchedID        SteamCallbackID = 1109
	GlobalAchievementPercentagesReadyID SteamCallbackID = 1110
	LeaderboardUGCSetID                 SteamCallbackID = 1111
	GlobalStatsReceivedID               SteamCallbackID = 1112
)

type UserStatsReceived struct {
	GameID      uint64
	Result      EResult
	SteamIDUser CSteamID
}
type UserStatsStored struct {
	GameID uint64
	Result EResult
}
type UserAchievementStored struct {
	GameID           uint64
	GroupAchievement bool
	AchievementName  [128]byte
	CurProgress      uint32
	MaxProgress      uint32
}
type LeaderboardFindResult struct {
	SteamLeaderboard SteamLeaderboard
	LeaderboardFound uint8
}
type LeaderboardScoresDownloaded struct {
	SteamLeaderboard        SteamLeaderboard
	SteamLeaderboardEntries SteamLeaderboardEntries
	EntryCount              int32
}
type LeaderboardScoreUploaded struct {
	Success            uint8
	SteamLeaderboard   SteamLeaderboard
	Score              int32
	ScoreChanged       uint8
	GlobalRankNew      int32
	GlobalRankPrevious int32
}
type NumberOfCurrentPlayers struct {
	Success uint8
	Players int32
}
type UserStatsUnloaded struct {
	SteamIDUser CSteamID
}
type UserAchievementIconFetched struct {
	GameID          CGameID
	AchievementName [128]byte
	Achieved        bool
	IconHandle      int32
}
type GlobalAchievementPercentagesReady struct {
	GameID uint64
	Result EResult
}
type LeaderboardUGCSet struct {
	Result           EResult
	SteamLeaderboard SteamLeaderboard
}
type GlobalStatsReceived struct {
	GameID uint64
	Result EResult
}

func (cb UserStatsReceived) CallbackID() SteamCallbackID {
	return UserStatsReceivedID
}

func (cb UserStatsReceived) Name() string {
	return "UserStatsReceived"
}

func (cb UserStatsReceived) String() string {
	return CallbackString(cb)
}

func (cb UserStatsStored) CallbackID() SteamCallbackID {
	return UserStatsStoredID
}

func (cb UserStatsStored) Name() string {
	return "UserStatsStored"
}

func (cb UserStatsStored) String() string {
	return CallbackString(cb)
}

func (cb UserAchievementStored) CallbackID() SteamCallbackID {
	return UserAchievementStoredID
}

func (cb UserAchievementStored) Name() string {
	return "UserAchievementStored"
}

func (cb UserAchievementStored) String() string {
	return CallbackString(cb)
}

func (cb LeaderboardFindResult) CallbackID() SteamCallbackID {
	return LeaderboardFindResultID
}

func (cb LeaderboardFindResult) Name() string {
	return "LeaderboardFindResult"
}

func (cb LeaderboardFindResult) String() string {
	return CallbackString(cb)
}

func (cb LeaderboardScoresDownloaded) CallbackID() SteamCallbackID {
	return LeaderboardScoresDownloadedID
}

func (cb LeaderboardScoresDownloaded) Name() string {
	return "LeaderboardScoresDownloaded"
}

func (cb LeaderboardScoresDownloaded) String() string {
	return CallbackString(cb)
}

func (cb LeaderboardScoreUploaded) CallbackID() SteamCallbackID {
	return LeaderboardScoreUploadedID
}

func (cb LeaderboardScoreUploaded) Name() string {
	return "LeaderboardScoreUploaded"
}

func (cb LeaderboardScoreUploaded) String() string {
	return CallbackString(cb)
}

func (cb NumberOfCurrentPlayers) CallbackID() SteamCallbackID {
	return NumberOfCurrentPlayersID
}

func (cb NumberOfCurrentPlayers) Name() string {
	return "NumberOfCurrentPlayers"
}

func (cb NumberOfCurrentPlayers) String() string {
	return CallbackString(cb)
}

func (cb UserStatsUnloaded) CallbackID() SteamCallbackID {
	return UserStatsUnloadedID
}

func (cb UserStatsUnloaded) Name() string {
	return "UserStatsUnloaded"
}

func (cb UserStatsUnloaded) String() string {
	return CallbackString(cb)
}

func (cb UserAchievementIconFetched) CallbackID() SteamCallbackID {
	return UserAchievementIconFetchedID
}

func (cb UserAchievementIconFetched) Name() string {
	return "UserAchievementIconFetched"
}

func (cb UserAchievementIconFetched) String() string {
	return CallbackString(cb)
}

func (cb GlobalAchievementPercentagesReady) CallbackID() SteamCallbackID {
	return GlobalAchievementPercentagesReadyID
}

func (cb GlobalAchievementPercentagesReady) Name() string {
	return "GlobalAchievementPercentagesReady"
}

func (cb GlobalAchievementPercentagesReady) String() string {
	return CallbackString(cb)
}

func (cb LeaderboardUGCSet) CallbackID() SteamCallbackID {
	return LeaderboardUGCSetID
}

func (cb LeaderboardUGCSet) Name() string {
	return "LeaderboardUGCSet"
}

func (cb LeaderboardUGCSet) String() string {
	return CallbackString(cb)
}

func (cb GlobalStatsReceived) CallbackID() SteamCallbackID {
	return GlobalStatsReceivedID
}

func (cb GlobalStatsReceived) Name() string {
	return "GlobalStatsReceived"
}

func (cb GlobalStatsReceived) String() string {
	return CallbackString(cb)
}
