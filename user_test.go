package girder

type testUser struct {
	id          string
	permissions []string
}

func (u *testUser) GetID() string {
	return u.id
}

func (u *testUser) GetPermissions() []string {
	return u.permissions
}
