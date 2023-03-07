# CHANGELOG

## v1.2.1

### Added
- Added PaginationLimit, PaginationOffset query params

### Removed
- Removed User state enum

## v1.2.0

### Added
- Added `Health` model that indicates the current health status reported by the cluster

### Removed

- Removed `database` field for `clusters_users_delete`,  `clusters_users_find_by_id`, `clusters_users_patch`  method in `UsersApi`


## v1.1.0

### Added

- Added `ClustersPatch` and its underlying data types, requests and execution functions
- Added `ClustersUsersPatch` and its underlying data types, requests and execution functions
- Added `Name`, `Edition` for templates
- Added `Id` for TemplateList
- Added `Snapshot` as a ResourceType
- Added `Pagination`: Offset, Limit, Links

### Removed

- Removed `database` field for `UserProperties`

## v1.0.0
### Overview
This is the first release of the SDK Go DBaaS Mongo. With IONOS Cloud Database as a service, you can quickly setup and manage MongoDB database clusters
