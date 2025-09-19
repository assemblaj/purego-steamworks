## purego-steamworks

Go bindings for the **Steamworks SDK** implemented with [`purego`](https://github.com/ebitengine/purego), allowing you to call the Steamworks C API from Go **without CGO**.

These bindings are designed for game developers who need **direct, low-overhead access** to Steam services such as achievements, leaderboards, multiplayer matchmaking, and inventory—while staying within pure Go toolchains.

### Features

* **Zero CGO**: Built on [`purego`](https://github.com/ebitengine/purego) for fast builds and simple cross-compilation.
* **Complete API coverage**: Wraps core Steamworks interfaces (e.g. `ISteamApps`, `ISteamFriends`, `ISteamGameServer`, `ISteamInput`) with idiomatic Go interfaces and types.
* **Callback & CallResult support**: Dispatches Steam callbacks and async call results using native Go generics and channels.
* **Windows and Linux**: Works wherever the Steamworks SDK is supported.
* **Permissive license**: Apache 2.0.

### Installation

```bash
go get github.com/assemblaj/purego-steamworks
```

This assumes you have the [Steamworks SDK](https://partner.steamgames.com/) downloaded and available on your build machine.

### Basic Usage

Initialize Steam and call any of the exposed interfaces.
Here’s a minimal example:

```go
package main

import (
    "fmt"
    sw "github.com/assemblaj/purego-steamworks"
)

func main() {
    // Ensure Steam is running and restart if necessary
    if sw.RestartAppIfNecessary(480) { // 480 = Spacewar test AppID
        return
    }

    if err := sw.Init(); err != nil {
        panic(err)
    }
    defer sw.Shutdown()

    name := sw.SteamFriends().GetPersonaName()
    fmt.Println("Logged in as:", name)

    // Pump callbacks regularly (e.g. each frame in a game loop)
    for {
        sw.RunCallbacks()
        // your game loop …
    }
}
```

### Repository Layout

* `interfaces.go` – Go interfaces mirroring Steamworks APIs (e.g. `ISteamFriends`, `ISteamApps`)
* `steam_methods.go` – Core Steam initialization, shutdown, and error handling
* `callback_dispatch.go` – Central dispatch for hundreds of Steam callback types
* `utils_*.go` – Platform-specific helpers (e.g. Windows string conversions)

### Development Notes

* The binding uses generated type definitions and flat API constants from `steam_api.json`.
* Callbacks are dispatched using `dispatchCallback` and `dispatchCallResult`, mapping the Steam `m_iCallback` IDs to strongly typed Go structs.
* Debug mode can be enabled with `sw.SetDebugMode(true)`.

### License

Apache License 2.0 — see [LICENSE](LICENSE).
