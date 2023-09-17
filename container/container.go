package container

import (
	"fmt"
	"glcharge/storage"
)

type container struct {
	storage storage.IDal
}

var containerinstance *container

func (c container) Storage() storage.IDal {
	return c.storage
}

func (c *container) SetStorage(storage storage.IDal) {
	c.storage = storage
}

func GetContainer() *container {
	if containerinstance == nil {
		if containerinstance == nil {
			fmt.Println("Creating single instance now.")
			containerinstance = &container{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return containerinstance
}

func ResetContainer() *container {
	containerinstance = &container{}
	return containerinstance
}
