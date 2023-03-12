package messages

import (
	"github.com/google/uuid"
	goecs "github.com/oneforx/go-ecs"
	goserge "github.com/oneforx/go-serge"
)

var (
	SC_CONNECT_FAILED = func(data interface{}) goserge.Message {
		return goserge.Message{
			MessageType: "CONNECT_FAILED",
			Data:        data,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_HYB,
		}
	}
	SC_CONNECT_SUCCESS = func(data interface{}) goserge.Message {
		return goserge.Message{
			MessageType: "CONNECT_SUCCESS",
			Data:        data,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_HYB,
		}
	}
	SC_PING = func(latence int) goserge.Message {
		return goserge.Message{
			MessageType: "PING",
			Data:        latence,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	SC_CREATE_ENTITY = func(entityData goecs.EntityNoCycle) goserge.Message {
		return goserge.Message{
			MessageType: "CREATE_ENTITY",
			Data:        entityData,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	SC_UPDATE_ENTITY = func(entityId uuid.UUID, entityData []*goecs.Component) goserge.Message {
		return goserge.Message{
			MessageType: "UPDATE_ENTITY",
			Data: UpdateMessageData{
				Id:         entityId,
				Components: entityData,
			},
			Target:  goserge.CLIENT_TARGET,
			NetMode: goserge.NET_UDP,
		}
	}
	SC_DELETE_ENTITY = func(id string) goserge.Message {
		return goserge.Message{
			MessageType: "DELETE_ENTITY",
			Data:        id,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
	SC_CREATE_WORLD = func(entitiesData []goecs.EntityNoCycle) goserge.Message {
		return goserge.Message{
			MessageType: "CREATE_WORLD",
			Data:        entitiesData,
			Target:      goserge.CLIENT_TARGET,
			NetMode:     goserge.NET_TCP,
		}
	}
)
