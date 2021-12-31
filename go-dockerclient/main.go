package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func main() {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}


	//imgs, err := client.ListImages(docker.ListImagesOptions{All: false})
	//if err != nil {
	//	panic(err)
	//}
	//for _, img := range imgs {
	//	fmt.Println("ID: ", img.ID)
	//	fmt.Println("RepoTags: ", img.RepoTags)
	//	fmt.Println("Created: ", img.Created)
	//	fmt.Println("Size: ", img.Size)
	//	fmt.Println("VirtualSize: ", img.VirtualSize)
	//	fmt.Println("ParentId: ", img.ParentID)
	//}


	//func (c *Client) CreateContainer(opts CreateContainerOptions) (*Container, error)
	//opts = CreateContainerOptions.
	//_ := CreateContainer()


	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "Ubuntu Container",
		Config: &docker.Config{
			Image: "ubuntu:latest",
		},
	})


	if err != nil {
		fmt.Println("============FAILED ================")
	}else {
		fmt.Println("================", container)
	}

}