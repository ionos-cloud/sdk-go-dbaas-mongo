package custom_test

import (
	"context"
	"fmt"

	"github.com/ionos-cloud/sdk-go-dbaas-mongo"
	"github.com/ionos-cloud/sdk-go-dbaas-mongo/custom"
)

func createSomeResources(client *ionoscloud.APIClient) {
	props := &ionoscloud.CreateClusterProperties{
		DisplayName:    ionoscloud.PtrString("TestCluster1"),
		Edition:        ionoscloud.PtrString("playground"),
		Instances:      ionoscloud.PtrInt32(1),
		TemplateID:     ionoscloud.PtrString("33457e53-1f8b-4ed2-8a12-2d42355aa759"),
		MongoDBVersion: ionoscloud.PtrString("6.0"),
		Location:       ionoscloud.PtrString("de/txl"),
		Connections: &[]ionoscloud.Connection{
			{
				DatacenterId: ionoscloud.PtrString("ade108ff-9daf-407a-9e72-361415fbd900"),
				LanId:        ionoscloud.PtrString("1"),
				CidrList: &[]string{
					"192.168.1.190/24",
				},
			},
		},
	}

	// Create a cluster
	c1, _, err := client.ClustersApi.ClustersPost(context.Background()).
		CreateClusterRequest(ionoscloud.CreateClusterRequest{Properties: props}).
		Execute()
	if err != nil {
		panic(fmt.Errorf("bad create: %w", err))
	}

	// Create a user
	uProps := &ionoscloud.UserProperties{
		Username: ionoscloud.PtrString("HelloWorld1"),
		Password: ionoscloud.PtrString("NGErwayncwqcynw0omwg"),
		Roles: &[]ionoscloud.UserRoles{
			{
				Role:     ionoscloud.PtrString("read"),
				Database: ionoscloud.PtrString("inexistant"),
			},
		},
	}
	_, _, err = client.UsersApi.ClustersUsersPost(context.Background(), *c1.Id).User(ionoscloud.User{Properties: uProps}).Execute()
	if err != nil {
		panic(fmt.Errorf("bad create: %w", err))
	}

	// Create another user with a different username
	uProps.Username = ionoscloud.PtrString("HelloWorld2")
	_, _, err = client.UsersApi.ClustersUsersPost(context.Background(), *c1.Id).User(ionoscloud.User{Properties: uProps}).Execute()
	if err != nil {
		panic(fmt.Errorf("bad create: %w", err))
	}

	// Create another cluster with a different DisplayName and a different Lan
	props.DisplayName = ionoscloud.PtrString("TestCluster2")
	(*props.Connections)[0].LanId = ionoscloud.PtrString("2")
	_, _, err = client.ClustersApi.ClustersPost(context.Background()).
		CreateClusterRequest(ionoscloud.CreateClusterRequest{Properties: props}).
		Execute()
	if err != nil {
		panic(fmt.Errorf("bad create: %w", err))
	}

}

func ExampleIdentifiable() {
	client := ionoscloud.NewAPIClient(ionoscloud.NewConfigurationFromEnv())
	// createSomeResources(client)

	// key: cluster name
	// value: slice of children of that cluster (snapshots, etc)
	clusterChildren := make(map[string][]custom.Resource)

	ls, _, err := client.ClustersApi.ClustersGet(context.Background()).Execute()
	if err != nil {
		return
	}

	for _, cluster := range *ls.Items {
		var children []custom.Resource

		snapshots, _, err := client.SnapshotsApi.ClustersSnapshotsGet(context.Background(), *cluster.Id).Execute()
		if err != nil {
			return
		}
		for _, u := range *snapshots.GetItems() {
			children = append(children, &u)
		}

		users, _, err := client.UsersApi.ClustersUsersGet(context.Background(), *cluster.Id).Execute()
		if err != nil {
			return
		}
		for _, u := range *users.GetItems() {
			children = append(children, &u)
		}

		clusterChildren[*cluster.Properties.DisplayName] = children
	}

	for clusterName, children := range clusterChildren {
		fmt.Printf("The cluster %s has the following children:\n", clusterName)

		typeCounts := make(map[string]int)

		for _, child := range children {
			typeCounts[string(*child.GetType())]++
		}

		// Display the counts
		for typeName, count := range typeCounts {
			fmt.Printf("  - %d %s(s)\n", count, typeName)
		}
	}

	// Output:
	// The cluster TestCluster1 has the following children:
	//   - 2 user(s)
	// The cluster TestCluster2 has the following children:

}
