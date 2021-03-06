package jira

type User struct {
	AvatarUrls  AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
}

type AvatarUrls struct {
	Four8X48  string `json:"48x48,omitempty" structs:"48x48,omitempty"`
	Two4X24   string `json:"24x24,omitempty" structs:"24x24,omitempty"`
	One6X16   string `json:"16x16,omitempty" structs:"16x16,omitempty"`
	Three2X32 string `json:"32x32,omitempty" structs:"32x32,omitempty"`
}
