package repository_interfaces

import "github.com/RandySteven/go-kopi/entities/models"

type UserRepository interface {
	Saver[models.User]
	Finder[models.User]
	Updater[models.User]
	Deleter[models.User]
}
