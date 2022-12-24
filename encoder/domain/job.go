package domain

import "time"

// * -> Ponteiro -> variavel global, alteravel de qualquer local | aponta pro mesmo local da memória
type Job struct {
	ID               string
	OutputBucketPath string
	Status           string
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
