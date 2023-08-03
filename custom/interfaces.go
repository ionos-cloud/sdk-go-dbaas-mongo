package ionoscloud

import (
	ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mongo"
)

type Identifiable interface {
	GetId() *string
	GetIdOk() (*string, bool)
	SetId(v string)
	HasId() bool

	GetType() *ionoscloud.ResourceType
	GetTypeOk() (*ionoscloud.ResourceType, bool)
	SetType(ionoscloud.ResourceType)
	HasType() bool
}

type Listable[T Identifiable] interface {
	Identifiable
	GetItems() *[]T
	GetItemsOk() (*[]T, bool)
	SetItems(v []T)
	HasItems() bool
}

type Offsetable[T Identifiable] interface {
	Identifiable
	GetOffset() *int32
	GetOffsetOk() (*int32, bool)
	SetOffset(int32)
	HasOffset() bool
}

type Limitable[T Identifiable] interface {
	Identifiable
	GetLimit() *int32
	GetLimitOk() (*int32, bool)
	SetLimit(int32)
	HasLimit() bool
}

type Stateful interface {
	GetState() *ionoscloud.State
	GetStateOk() (*ionoscloud.State, bool)
	SetState(ionoscloud.State)
}

type Executable[T Identifiable] interface {
	Execute() (T, *ionoscloud.APIResponse, error)
}
