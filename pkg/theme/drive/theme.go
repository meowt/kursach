package drive

import (
	"Diploma/pkg/user"
)

type ThemeDriveImpl struct {
}

type ThemeDriveModule struct {
	user.Gateway
}

func SetupThemeDrive() ThemeDriveModule {
	return ThemeDriveModule{
		Gateway: &ThemeDriveImpl{},
	}
}
