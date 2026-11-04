package entities

// ValidatedUser for User object has passed validation.
type ValidatedUser struct {
	User

	IsValidated bool
}

// NewValidatedUser to copy user object to be ValidatedUser
// which has passed validation and retains its validated state safely.
func NewValidatedUser(user *User) (*ValidatedUser, error) {
	if err := user.validate(); err != nil {
		return nil, err
	}

	return &ValidatedUser{
		User:        *user,
		IsValidated: true,
	}, nil
}

// IsValid to chect value object IsValidated.
func (v *ValidatedUser) IsValid() bool {
	return v.IsValidated
}
