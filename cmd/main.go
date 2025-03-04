package main

import (
	"fmt"
	"os"

	"github.com/assemblaj/purego-steamworks/callback"
	"github.com/assemblaj/purego-steamworks/friends"
	networktypes "github.com/assemblaj/purego-steamworks/network_types"
	networkingmessages "github.com/assemblaj/purego-steamworks/networking_messages"

	steamworks "github.com/assemblaj/purego-steamworks"
	//matchmaking "github.com/assemblaj/purego-steamworks/matchmaking"
)

func main() {
	defer steamworks.Shutdown()

	if steamworks.RestartAppiIfNecessary(1431780) {
		os.Exit(1)
	}
	if steamworks.Init() != nil {
		panic("steamworks.Init failed")
	}

	go callback.DoForever()

	// if friends.GetPersonaName() == "assemblaj" {
	// 	handle := matchmaking.RequestLobbyList()
	// 	fmt.Println(handle)
	// 	fmt.Println("getting LOBBY list WITH  API call ")

	// 	callback.RegisterCallResult(handle, func(result matchmaking.LobbyMatchList, failed bool) {
	// 		fmt.Println("lobbies matching")
	// 		fmt.Println(result.LobbiesMatching)
	// 		callback.RegisterCallback(func(call matchmaking.LobbyEnter) {
	// 			fmt.Println("Entered Lobby")
	// 			fmt.Println(call.SteamIDLobby)
	// 		})
	// 		steamID := matchmaking.GetLobbyByIndex(0)
	// 		newHandle := matchmaking.JoinLobby(steamID)
	// 		fmt.Println("joining LOBBY WITH  API call ")
	// 		fmt.Println(newHandle)
	// 		callback.RegisterCallResult(newHandle, func(result matchmaking.LobbyEnter, failed bool) {
	// 			fmt.Println("lobby joined ")
	// 			fmt.Println(result.SteamIDLobby)
	// 		})
	// 	})
	// } else {
	// 	callback.RegisterCallback(func(call matchmaking.LobbyCreated) {
	// 		fmt.Println("Created Lobby")
	// 	})
	// 	callback.RegisterCallback(func(call matchmaking.LobbyEnter) {
	// 		fmt.Println("Entered Lobby")
	// 	})
	// 	handle := matchmaking.CreateLobby(2, 2)
	// 	fmt.Println("CREATING LOBBY WITH  API call ")
	// 	fmt.Println(handle)
	// 	callback.RegisterCallResult(handle, func(result matchmaking.LobbyCreated, failed bool) {
	// 		fmt.Println("lobby created")
	// 		fmt.Println(result.SteamIDLobby)
	// 	})
	// }

	// result := friends.GetPersonaName()
	// fmt.Println(result)
	// resultA := apps.GetCurrentGameLanguage() // works
	// fmt.Println(resultA)
	// resultB := apps.GetAvailableGameLanguages() // works
	// fmt.Println(resultB)
	// resultC := apps.GetAppInstallDir(1431780) // works
	// fmt.Println(resultC)

	for i := 0; i < int(friends.GetFriendCount(int32(friends.EFriendFlagImmediate))); i++ {
		id := friends.GetFriendByIndex(int32(i), int32(friends.EFriendFlagImmediate))
		name := friends.GetFriendPersonaName(id)
		if name == "assemblaj" || name == "Fantasma" {
			var snd networktypes.SteamNetworkingIdentity
			snd.SetSteamID64(uint64(id))
			callback.RegisterCallback(func(cb networkingmessages.SteamNetworkingMessagesSessionFailed) {
				fmt.Println("Network Session failed ")
			})
			callback.RegisterCallback(func(cb networkingmessages.SteamNetworkingMessagesSessionRequest) {
				fmt.Println("accepting message")
				networkingmessages.AcceptSessionWithUser(cb.IdentityRemote)

				msgs := networkingmessages.ReceiveMessagesOnChannel(1, 2)
				var msgNo int
				msgNo = len(msgs)
				for msgNo == 0 {
					fmt.Println("no messages recieved on channel")
					msgs = networkingmessages.ReceiveMessagesOnChannel(1, 2)
					msgNo = len(msgs)
				}
				fmt.Println("recieving message")
				fmt.Println("message amount recieved")
				fmt.Println(len(msgs))
				fmt.Println("messages:")
				for i := 0; i < len(msgs); i++ {
					fmt.Println(msgs[i])
				}
				networkingmessages.CloseSessionWithUser(cb.IdentityRemote)
			})
			networkingmessages.SendMessageToUser(snd, []byte("Hello World!"), 8, 1)
		}
	}

	for {

	}
}
