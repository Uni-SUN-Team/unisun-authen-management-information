package repositories

import "unisun/api/unisun-authen-inquiry/src/entities"

type RepositoriesUserAuthPermission interface {
	FindbyUserid(userId int) entities.UserAuthPermission
	Create(Data entities.UserAuthPermission)
	UpdateVersionToken(versionToken int, Data entities.UserAuthPermission)
}
