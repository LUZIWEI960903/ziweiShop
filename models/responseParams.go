package models

type NewRoleList struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerListRole struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerList struct {
	Id       int                  `json:"id"`       // manager的id
	Username string               `json:"username"` // manager的名字
	Mobile   string               `json:"mobile"`   // manager的电话
	Email    string               `json:"email"`    // manager的邮箱
	AddTime  string               `json:"add_time"` // manager的创建时间
	Role     IndexManagerListRole `json:"role"`     // manager对应role的信息
}

type EditManagerList struct {
	Id       int           `json:"id"`        // manager的id
	Username string        `json:"username"`  // manager的名字
	Password string        `json:"password"`  // manager的密码
	Mobile   string        `json:"mobile"`    // manager的电话
	Email    string        `json:"email"`     // manager的邮箱
	RoleList []NewRoleList `json:"role_list"` // 所有role的信息
}
