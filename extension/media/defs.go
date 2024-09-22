package media

const (
	ErrUnimplemented         Error = "unimplemented method"
	ErrForbidden             Error = "forbidden"
	ErrMediaNotFound         Error = "media not found"
	ErrEpisodeNotFound       Error = "episode not found"
	ErrAuthorizationRequired Error = "authorization required"
)

func (e Error) Error() string {
	return string(e)
}
