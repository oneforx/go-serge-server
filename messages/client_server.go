package messages

import (
	"github.com/google/uuid"
	goserge "github.com/oneforx/go-serge"
)

var (
	CS_CONNECT_TOKEN = func(message string) goserge.Message {
		return goserge.Message{
			MessageType: "CONNECT_TOKEN",
			Data:        message,
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_HYB,
		}
	}
	CS_PONG = func() goserge.Message {
		return goserge.Message{
			MessageType: "PONG",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_CREATE_CHARACTER = func() goserge.Message {
		return goserge.Message{
			MessageType: "CREATE_CHARACTER",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_PLAY_CHARACTER = func(worldId, characterId uuid.UUID) goserge.Message {
		return goserge.Message{
			MessageType: "PLAY_CHARACTER",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_DELETE_CHARACTER = func(id string) goserge.Message {
		return goserge.Message{
			MessageType: "DELETE_CHARACTER",
			Data:        id,
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_CONNECT_WORLD = func() goserge.Message {
		return goserge.Message{
			MessageType: "CONNECT_WORLD",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_DISCONNECT_WORLD = func() goserge.Message {
		return goserge.Message{
			MessageType: "DISCONNECT_WORLD",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_SHOOT = func() goserge.Message {
		return goserge.Message{
			MessageType: "SHOOT",
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	CS_ORIENTATION = func(x int, y int) goserge.Message {
		return goserge.Message{
			MessageType: "ORIENTATION",
			Data:        map[string]interface{}{"x": x, "y": y},
			Target:      goserge.SERVER_TARGET,
			NetMode:     goserge.NET_UDP,
		}
	}
)
