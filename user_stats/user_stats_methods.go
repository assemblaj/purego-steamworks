package userstats

import (
	. "github.com/assemblaj/purego-steamworks"
)

type LeaderboardEntry struct {
	SteamIDUser CSteamID
	GlobalRank  int32
	Score       int32
	Details     int32
	UGC         UGCHandle
}
type ELeaderboardDataRequest int32

const (
	ELeaderboardDataRequestGlobal           ELeaderboardDataRequest = 0
	ELeaderboardDataRequestGlobalAroundUser ELeaderboardDataRequest = 1
	ELeaderboardDataRequestFriends          ELeaderboardDataRequest = 2
	ELeaderboardDataRequestUsers            ELeaderboardDataRequest = 3
)

type ELeaderboardSortMethod int32

const (
	ELeaderboardSortMethodNone       ELeaderboardSortMethod = 0
	ELeaderboardSortMethodAscending  ELeaderboardSortMethod = 1
	ELeaderboardSortMethodDescending ELeaderboardSortMethod = 2
)

type ELeaderboardDisplayType int32

const (
	ELeaderboardDisplayTypeNone             ELeaderboardDisplayType = 0
	ELeaderboardDisplayTypeNumeric          ELeaderboardDisplayType = 1
	ELeaderboardDisplayTypeTimeSeconds      ELeaderboardDisplayType = 2
	ELeaderboardDisplayTypeTimeMilliSeconds ELeaderboardDisplayType = 3
)

type ELeaderboardUploadScoreMethod int32

const (
	ELeaderboardUploadScoreMethodNone        ELeaderboardUploadScoreMethod = 0
	ELeaderboardUploadScoreMethodKeepBest    ELeaderboardUploadScoreMethod = 1
	ELeaderboardUploadScoreMethodForceUpdate ELeaderboardUploadScoreMethod = 2
)

type SteamLeaderboard uint64
type SteamLeaderboardEntries uint64

const (
	flatAPI_SteamUserStats                                      = "SteamAPI_SteamUserStats_v013"
	flatAPI_ISteamUserStats_GetStatInt32                        = "SteamAPI_ISteamUserStats_GetStatInt32"
	flatAPI_ISteamUserStats_GetStatFloat                        = "SteamAPI_ISteamUserStats_GetStatFloat"
	flatAPI_ISteamUserStats_SetStatInt32                        = "SteamAPI_ISteamUserStats_SetStatInt32"
	flatAPI_ISteamUserStats_SetStatFloat                        = "SteamAPI_ISteamUserStats_SetStatFloat"
	flatAPI_ISteamUserStats_UpdateAvgRateStat                   = "SteamAPI_ISteamUserStats_UpdateAvgRateStat"
	flatAPI_ISteamUserStats_GetAchievement                      = "SteamAPI_ISteamUserStats_GetAchievement"
	flatAPI_ISteamUserStats_SetAchievement                      = "SteamAPI_ISteamUserStats_SetAchievement"
	flatAPI_ISteamUserStats_ClearAchievement                    = "SteamAPI_ISteamUserStats_ClearAchievement"
	flatAPI_ISteamUserStats_GetAchievementAndUnlockTime         = "SteamAPI_ISteamUserStats_GetAchievementAndUnlockTime"
	flatAPI_ISteamUserStats_StoreStats                          = "SteamAPI_ISteamUserStats_StoreStats"
	flatAPI_ISteamUserStats_GetAchievementIcon                  = "SteamAPI_ISteamUserStats_GetAchievementIcon"
	flatAPI_ISteamUserStats_GetAchievementDisplayAttribute      = "SteamAPI_ISteamUserStats_GetAchievementDisplayAttribute"
	flatAPI_ISteamUserStats_IndicateAchievementProgress         = "SteamAPI_ISteamUserStats_IndicateAchievementProgress"
	flatAPI_ISteamUserStats_GetNumAchievements                  = "SteamAPI_ISteamUserStats_GetNumAchievements"
	flatAPI_ISteamUserStats_GetAchievementName                  = "SteamAPI_ISteamUserStats_GetAchievementName"
	flatAPI_ISteamUserStats_RequestUserStats                    = "SteamAPI_ISteamUserStats_RequestUserStats"
	flatAPI_ISteamUserStats_GetUserStatInt32                    = "SteamAPI_ISteamUserStats_GetUserStatInt32"
	flatAPI_ISteamUserStats_GetUserStatFloat                    = "SteamAPI_ISteamUserStats_GetUserStatFloat"
	flatAPI_ISteamUserStats_GetUserAchievement                  = "SteamAPI_ISteamUserStats_GetUserAchievement"
	flatAPI_ISteamUserStats_GetUserAchievementAndUnlockTime     = "SteamAPI_ISteamUserStats_GetUserAchievementAndUnlockTime"
	flatAPI_ISteamUserStats_ResetAllStats                       = "SteamAPI_ISteamUserStats_ResetAllStats"
	flatAPI_ISteamUserStats_FindOrCreateLeaderboard             = "SteamAPI_ISteamUserStats_FindOrCreateLeaderboard"
	flatAPI_ISteamUserStats_FindLeaderboard                     = "SteamAPI_ISteamUserStats_FindLeaderboard"
	flatAPI_ISteamUserStats_GetLeaderboardName                  = "SteamAPI_ISteamUserStats_GetLeaderboardName"
	flatAPI_ISteamUserStats_GetLeaderboardEntryCount            = "SteamAPI_ISteamUserStats_GetLeaderboardEntryCount"
	flatAPI_ISteamUserStats_GetLeaderboardSortMethod            = "SteamAPI_ISteamUserStats_GetLeaderboardSortMethod"
	flatAPI_ISteamUserStats_GetLeaderboardDisplayType           = "SteamAPI_ISteamUserStats_GetLeaderboardDisplayType"
	flatAPI_ISteamUserStats_DownloadLeaderboardEntries          = "SteamAPI_ISteamUserStats_DownloadLeaderboardEntries"
	flatAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers  = "SteamAPI_ISteamUserStats_DownloadLeaderboardEntriesForUsers"
	flatAPI_ISteamUserStats_GetDownloadedLeaderboardEntry       = "SteamAPI_ISteamUserStats_GetDownloadedLeaderboardEntry"
	flatAPI_ISteamUserStats_UploadLeaderboardScore              = "SteamAPI_ISteamUserStats_UploadLeaderboardScore"
	flatAPI_ISteamUserStats_AttachLeaderboardUGC                = "SteamAPI_ISteamUserStats_AttachLeaderboardUGC"
	flatAPI_ISteamUserStats_GetNumberOfCurrentPlayers           = "SteamAPI_ISteamUserStats_GetNumberOfCurrentPlayers"
	flatAPI_ISteamUserStats_RequestGlobalAchievementPercentages = "SteamAPI_ISteamUserStats_RequestGlobalAchievementPercentages"
	flatAPI_ISteamUserStats_GetMostAchievedAchievementInfo      = "SteamAPI_ISteamUserStats_GetMostAchievedAchievementInfo"
	flatAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo  = "SteamAPI_ISteamUserStats_GetNextMostAchievedAchievementInfo"
	flatAPI_ISteamUserStats_GetAchievementAchievedPercent       = "SteamAPI_ISteamUserStats_GetAchievementAchievedPercent"
	flatAPI_ISteamUserStats_RequestGlobalStats                  = "SteamAPI_ISteamUserStats_RequestGlobalStats"
	flatAPI_ISteamUserStats_GetGlobalStatInt64                  = "SteamAPI_ISteamUserStats_GetGlobalStatInt64"
	flatAPI_ISteamUserStats_GetGlobalStatDouble                 = "SteamAPI_ISteamUserStats_GetGlobalStatDouble"
	flatAPI_ISteamUserStats_GetGlobalStatHistoryInt64           = "SteamAPI_ISteamUserStats_GetGlobalStatHistoryInt64"
	flatAPI_ISteamUserStats_GetGlobalStatHistoryDouble          = "SteamAPI_ISteamUserStats_GetGlobalStatHistoryDouble"
	flatAPI_ISteamUserStats_GetAchievementProgressLimitsInt32   = "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsInt32"
	flatAPI_ISteamUserStats_GetAchievementProgressLimitsFloat   = "SteamAPI_ISteamUserStats_GetAchievementProgressLimitsFloat"
)

var (
	steamUserStats_init                                 func() uintptr
	iSteamUserStats_GetStatInt32                        func(steamUserStats uintptr, pchName string, pData *int32) bool
	iSteamUserStats_GetStatFloat                        func(steamUserStats uintptr, pchName string, pData *float32) bool
	iSteamUserStats_SetStatInt32                        func(steamUserStats uintptr, pchName string, nData int32) bool
	iSteamUserStats_SetStatFloat                        func(steamUserStats uintptr, pchName string, fData float32) bool
	iSteamUserStats_UpdateAvgRateStat                   func(steamUserStats uintptr, pchName string, flCountThisSession float32, dSessionLength float64) bool
	iSteamUserStats_GetAchievement                      func(steamUserStats uintptr, pchName string, pbAchieved *bool) bool
	iSteamUserStats_SetAchievement                      func(steamUserStats uintptr, pchName string) bool
	iSteamUserStats_ClearAchievement                    func(steamUserStats uintptr, pchName string) bool
	iSteamUserStats_GetAchievementAndUnlockTime         func(steamUserStats uintptr, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool
	iSteamUserStats_StoreStats                          func(steamUserStats uintptr) bool
	iSteamUserStats_GetAchievementIcon                  func(steamUserStats uintptr, pchName string) int32
	iSteamUserStats_GetAchievementDisplayAttribute      func(steamUserStats uintptr, pchName string, pchKey string) string
	iSteamUserStats_IndicateAchievementProgress         func(steamUserStats uintptr, pchName string, nCurProgress uint32, nMaxProgress uint32) bool
	iSteamUserStats_GetNumAchievements                  func(steamUserStats uintptr) uint32
	iSteamUserStats_GetAchievementName                  func(steamUserStats uintptr, iAchievement uint32) string
	iSteamUserStats_RequestUserStats                    func(steamUserStats uintptr, steamIDUser Uint64SteamID) SteamAPICall
	iSteamUserStats_GetUserStatInt32                    func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pData *int32) bool
	iSteamUserStats_GetUserStatFloat                    func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pData *float32) bool
	iSteamUserStats_GetUserAchievement                  func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool) bool
	iSteamUserStats_GetUserAchievementAndUnlockTime     func(steamUserStats uintptr, steamIDUser Uint64SteamID, pchName string, pbAchieved *bool, punUnlockTime *uint32) bool
	iSteamUserStats_ResetAllStats                       func(steamUserStats uintptr, bAchievementsToo bool) bool
	iSteamUserStats_FindOrCreateLeaderboard             func(steamUserStats uintptr, pchLeaderboardName string, eLeaderboardSortMethod ELeaderboardSortMethod, eLeaderboardDisplayType ELeaderboardDisplayType) SteamAPICall
	iSteamUserStats_FindLeaderboard                     func(steamUserStats uintptr, pchLeaderboardName string) SteamAPICall
	iSteamUserStats_GetLeaderboardName                  func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) string
	iSteamUserStats_GetLeaderboardEntryCount            func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) int32
	iSteamUserStats_GetLeaderboardSortMethod            func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardSortMethod
	iSteamUserStats_GetLeaderboardDisplayType           func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard) ELeaderboardDisplayType
	iSteamUserStats_DownloadLeaderboardEntries          func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) SteamAPICall
	iSteamUserStats_DownloadLeaderboardEntriesForUsers  func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, prgUsers []CSteamID, cUsers int32) SteamAPICall
	iSteamUserStats_GetDownloadedLeaderboardEntry       func(steamUserStats uintptr, hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, pLeaderboardEntry *LeaderboardEntry, pDetails *[]int32, cDetailsMax int32) bool
	iSteamUserStats_UploadLeaderboardScore              func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32, cScoreDetailsCount int32) SteamAPICall
	iSteamUserStats_AttachLeaderboardUGC                func(steamUserStats uintptr, hSteamLeaderboard SteamLeaderboard, hUGC UGCHandle) SteamAPICall
	iSteamUserStats_GetNumberOfCurrentPlayers           func(steamUserStats uintptr) SteamAPICall
	iSteamUserStats_RequestGlobalAchievementPercentages func(steamUserStats uintptr) SteamAPICall
	iSteamUserStats_GetMostAchievedAchievementInfo      func(steamUserStats uintptr, pchName *[]byte, unNameBufLen uint32, pflPercent *float32, pbAchieved *bool) int32
	iSteamUserStats_GetNextMostAchievedAchievementInfo  func(steamUserStats uintptr, iIteratorPrevious int32, pchName *[]byte, unNameBufLen uint32, pflPercent *float32, pbAchieved *bool) int32
	iSteamUserStats_GetAchievementAchievedPercent       func(steamUserStats uintptr, pchName string, pflPercent *float32) bool
	iSteamUserStats_RequestGlobalStats                  func(steamUserStats uintptr, nHistoryDays int32) SteamAPICall
	iSteamUserStats_GetGlobalStatInt64                  func(steamUserStats uintptr, pchStatName string, pData *int64) bool
	iSteamUserStats_GetGlobalStatDouble                 func(steamUserStats uintptr, pchStatName string, pData *float64) bool
	iSteamUserStats_GetGlobalStatHistoryInt64           func(steamUserStats uintptr, pchStatName string, pData *[]int64, cubData uint32) int32
	iSteamUserStats_GetGlobalStatHistoryDouble          func(steamUserStats uintptr, pchStatName string, pData *[]float64, cubData uint32) int32
	iSteamUserStats_GetAchievementProgressLimitsInt32   func(steamUserStats uintptr, pchName string, pnMinProgress *int32, pnMaxProgress *int32) bool
	iSteamUserStats_GetAchievementProgressLimitsFloat   func(steamUserStats uintptr, pchName string, pfMinProgress *float32, pfMaxProgress *float32) bool
)

var steamUserStats uintptr

func userStats() uintptr {
	if steamUserStats == 0 {
		steamUserStats = steamUserStats_init()
		return steamUserStats
	}
	return steamUserStats
}

func GetStat(name string) (data int32, success bool) {
	success = iSteamUserStats_GetStatInt32(userStats(), name, &data)
	return data, success
}

func GetStatFloat(name string) (data float32, success bool) {
	success = iSteamUserStats_GetStatFloat(userStats(), name, &data)
	return data, success
}

func SetStat(name string, data int32) bool {
	return iSteamUserStats_SetStatInt32(userStats(), name, data)
}

func SetStatFloat(name string, data float32) bool {
	return iSteamUserStats_SetStatFloat(userStats(), name, data)
}

func UpdateAvgRateStat(name string, countThisSession float32, sessionLength float64) bool {
	return iSteamUserStats_UpdateAvgRateStat(userStats(), name, countThisSession, sessionLength)
}

func GetAchievement(name string) (achieved bool, success bool) {
	success = iSteamUserStats_GetAchievement(userStats(), name, &achieved)
	return achieved, success
}

func SetAchievement(name string) bool {
	return iSteamUserStats_SetAchievement(userStats(), name)
}

func ClearAchievement(name string) bool {
	return iSteamUserStats_ClearAchievement(userStats(), name)
}

func GetAchievementAndUnlockTime(name string) (achieved bool, unlockTime uint32, success bool) {
	success = iSteamUserStats_GetAchievementAndUnlockTime(userStats(), name, &achieved, &unlockTime)
	return achieved, unlockTime, success
}

func StoreStats() bool {
	return iSteamUserStats_StoreStats(userStats())
}

func GetAchievementIcon(name string) int32 {
	return iSteamUserStats_GetAchievementIcon(userStats(), name)
}

func GetAchievementDisplayAttribute(name string, key string) string {
	return iSteamUserStats_GetAchievementDisplayAttribute(userStats(), name, key)
}

func IndicateAchievementProgress(name string, curProgress uint32, maxProgress uint32) bool {
	return iSteamUserStats_IndicateAchievementProgress(userStats(), name, curProgress, curProgress)
}

func GetNumAchievements() uint32 {
	return iSteamUserStats_GetNumAchievements(userStats())
}

func GetAchievementName(achievement uint32) string {
	return iSteamUserStats_GetAchievementName(userStats(), achievement)
}

func RequestUserStats(userSteamID Uint64SteamID) CallResult[UserStatsReceived] {
	handle := iSteamUserStats_RequestUserStats(userStats(), userSteamID)
	return CallResult[UserStatsReceived]{Handle: handle}
}

func GetUserStat(userSteamID Uint64SteamID, name string) (data int32, success bool) {
	success = iSteamUserStats_GetUserStatInt32(userStats(), userSteamID, name, &data)
	return data, success
}

func GetUserStatFloat(userSteamID Uint64SteamID, name string) (data float32, success bool) {
	success = iSteamUserStats_GetUserStatFloat(userStats(), userSteamID, name, &data)
	return data, success
}

func GetUserAchievement(userSteamID Uint64SteamID, name string) (achieved bool, success bool) {
	success = iSteamUserStats_GetUserAchievement(userStats(), userSteamID, name, &achieved)
	return achieved, success
}

func GetUserAchievementAndUnlockTime(userSteamID Uint64SteamID, name string) (achieved bool, unlockTime uint32, success bool) {
	success = iSteamUserStats_GetUserAchievementAndUnlockTime(userStats(), userSteamID, name, &achieved, &unlockTime)
	return achieved, unlockTime, success
}

func ResetAllStats(achievementsToo bool) bool {
	return iSteamUserStats_ResetAllStats(userStats(), achievementsToo)
}

func FindOrCreateLeaderboard(leaderboardName string, leaderboardSortMethod ELeaderboardSortMethod, leaderboardDisplayType ELeaderboardDisplayType) CallResult[LeaderboardFindResult] {
	handle := iSteamUserStats_FindOrCreateLeaderboard(userStats(), leaderboardName, leaderboardSortMethod, leaderboardDisplayType)
	return CallResult[LeaderboardFindResult]{Handle: handle}
}

func FindLeaderboard(leaderboardName string) CallResult[LeaderboardFindResult] {
	handle := iSteamUserStats_FindLeaderboard(userStats(), leaderboardName)
	return CallResult[LeaderboardFindResult]{Handle: handle}
}

func GetLeaderboardName(steamLeaderboard SteamLeaderboard) string {
	return iSteamUserStats_GetLeaderboardName(userStats(), steamLeaderboard)
}

func GetLeaderboardEntryCount(steamLeaderboard SteamLeaderboard) int32 {
	return iSteamUserStats_GetLeaderboardEntryCount(userStats(), steamLeaderboard)
}

func GetLeaderboardSortMethod(steamLeaderboard SteamLeaderboard) ELeaderboardSortMethod {
	return iSteamUserStats_GetLeaderboardSortMethod(userStats(), steamLeaderboard)
}

func GetLeaderboardDisplayType(steamLeaderboard SteamLeaderboard) ELeaderboardDisplayType {
	return iSteamUserStats_GetLeaderboardDisplayType(userStats(), steamLeaderboard)
}

func DownloadLeaderboardEntries(steamLeaderboard SteamLeaderboard, eLeaderboardDataRequest ELeaderboardDataRequest, nRangeStart int32, nRangeEnd int32) CallResult[LeaderboardScoresDownloaded] {
	handle := iSteamUserStats_DownloadLeaderboardEntries(userStats(), steamLeaderboard, eLeaderboardDataRequest, nRangeStart, nRangeEnd)
	return CallResult[LeaderboardScoresDownloaded]{Handle: handle}
}

func DownloadLeaderboardEntriesForUsers(steamLeaderboard SteamLeaderboard, prgUsers []CSteamID) CallResult[LeaderboardScoresDownloaded] {
	handle := iSteamUserStats_DownloadLeaderboardEntriesForUsers(userStats(), steamLeaderboard, prgUsers, int32(len(prgUsers)))
	return CallResult[LeaderboardScoresDownloaded]{Handle: handle}
}

func GetDownloadedLeaderboardEntry(hSteamLeaderboardEntries SteamLeaderboardEntries, index int32, cDetailsMax int32) (pLeaderboardEntry LeaderboardEntry, pDetails []int32, success bool) {
	pDetails = make([]int32, 0, cDetailsMax)
	success = iSteamUserStats_GetDownloadedLeaderboardEntry(userStats(), hSteamLeaderboardEntries, index, &pLeaderboardEntry, &pDetails, cDetailsMax)
	return pLeaderboardEntry, pDetails, success
}

func UploadLeaderboardScore(steamLeaderboard SteamLeaderboard, eLeaderboardUploadScoreMethod ELeaderboardUploadScoreMethod, nScore int32, pScoreDetails []int32) CallResult[LeaderboardScoreUploaded] {
	handle := iSteamUserStats_UploadLeaderboardScore(userStats(), steamLeaderboard, eLeaderboardUploadScoreMethod, nScore, pScoreDetails, int32(len(pScoreDetails)))
	return CallResult[LeaderboardScoreUploaded]{Handle: handle}
}

func AttachLeaderboardUGC(steamLeaderboard SteamLeaderboard, hUGC UGCHandle) CallResult[LeaderboardUGCSet] {
	handle := iSteamUserStats_AttachLeaderboardUGC(userStats(), steamLeaderboard, hUGC)
	return CallResult[LeaderboardUGCSet]{Handle: handle}
}

func GetNumberOfCurrentPlayers() CallResult[NumberOfCurrentPlayers] {
	handle := iSteamUserStats_GetNumberOfCurrentPlayers(userStats())
	return CallResult[NumberOfCurrentPlayers]{Handle: handle}
}

func RequestGlobalAchievementPercentages() CallResult[GlobalAchievementPercentagesReady] {
	handle := iSteamUserStats_RequestGlobalAchievementPercentages(userStats())
	return CallResult[GlobalAchievementPercentagesReady]{Handle: handle}
}

func GetMostAchievedAchievementInfo(nameBufLen uint32) (name string, percent float32, achieved bool, index int32) {
	buf := make([]byte, 0, nameBufLen)

	index = iSteamUserStats_GetMostAchievedAchievementInfo(userStats(), &buf, nameBufLen, &percent, &achieved)
	return string(buf), percent, achieved, index
}

func GetNextMostAchievedAchievementInfo(iteratorPrevious int32, nameBufLen uint32) (name string, percent float32, achieved bool, index int32) {
	buf := make([]byte, 0, nameBufLen)

	index = iSteamUserStats_GetNextMostAchievedAchievementInfo(userStats(), iteratorPrevious, &buf, nameBufLen, &percent, &achieved)
	return string(buf), percent, achieved, index
}

func GetAchievementAchievedPercent(name string) (percent float32, success bool) {
	success = iSteamUserStats_GetAchievementAchievedPercent(userStats(), name, &percent)
	return percent, success
}

func RequestGlobalStats(historyDays int32) CallResult[GlobalStatsReceived] {
	handle := iSteamUserStats_RequestGlobalStats(userStats(), historyDays)
	return CallResult[GlobalStatsReceived]{Handle: handle}
}

func GetGlobalStat(pchStatName string) (data int64, success bool) {
	success = iSteamUserStats_GetGlobalStatInt64(userStats(), pchStatName, &data)
	return data, success
}

func GetGlobalStatDouble(pchStatName string) (data float64, success bool) {
	success = iSteamUserStats_GetGlobalStatDouble(userStats(), pchStatName, &data)
	return data, success
}

func GetGlobalStatHistory(pchStatName string, dataSize uint32) (data []int64) {
	data = make([]int64, 0, dataSize)
	result := iSteamUserStats_GetGlobalStatHistoryInt64(userStats(), pchStatName, &data, dataSize)
	return data[:result]
}

func GetGlobalStatHistoryDouble(statName string, dataSize uint32) (data []float64) {
	data = make([]float64, 0, dataSize)
	result := iSteamUserStats_GetGlobalStatHistoryDouble(userStats(), statName, &data, dataSize)
	return data[:result]
}

func GetAchievementProgressLimits(name string) (minProgress int32, maxProgress int32, success bool) {
	success = iSteamUserStats_GetAchievementProgressLimitsInt32(userStats(), name, &minProgress, &maxProgress)
	return minProgress, maxProgress, success
}

func GetAchievementProgressLimitsFloat(name string) (minProgress float32, maxProgress float32, success bool) {
	success = iSteamUserStats_GetAchievementProgressLimitsFloat(userStats(), name, &minProgress, &maxProgress)
	return minProgress, maxProgress, success
}
