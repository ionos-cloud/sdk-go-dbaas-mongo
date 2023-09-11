package custom

import (
	ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mongo"
)

type Resource interface {
	GetType() *ionoscloud.ResourceType
	GetTypeOk() (*ionoscloud.ResourceType, bool)
	SetType(ionoscloud.ResourceType)
	HasType() bool
}

type Identifiable interface {
	Resource

	GetId() *string
	GetIdOk() (*string, bool)
	SetId(v string)
	HasId() bool
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

type Paginated[T Identifiable] interface {
	Identifiable
	Offsetable[T]
	Listable[T]
}

type Stateful interface {
	GetState() *ionoscloud.State
	GetStateOk() (*ionoscloud.State, bool)
	SetState(ionoscloud.State)
}

// Can't use this:
//     Type does not implement 'custom.Executable[*ionoscloud.ClusterResponse]'
//     need the method: Execute() (*ionoscloud.ClusterResponse, *ionoscloud.APIResponse, error)
//     have the method: Execute() (ClusterResponse, *APIResponse, error)
//
// (removing ptr ends up with the unimplemented Resource interface instead)
//
// type Executable[T Identifiable] interface {
// 	Execute() (*T, *ionoscloud.APIResponse, error)
// }
