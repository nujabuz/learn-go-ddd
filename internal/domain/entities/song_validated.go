package entities

// ValidatedSong for Song object has passed validation.
type ValidatedSong struct {
	Song

	IsValidated bool
}

// NewValidatedSong to copy song object to be ValidatedSong
// which has passed validation and retains its validated state safely.
func NewValidatedSong(song *Song) (*ValidatedSong, error) {
	if err := song.validate(); err != nil {
		return nil, err
	}

	return &ValidatedSong{
		Song:        *song,
		IsValidated: true,
	}, nil
}

// IsValid to chect value object IsValidated.
func (v *ValidatedSong) IsValid() bool {
	return v.IsValidated
}
