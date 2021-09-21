package profile

import (
	"errors"
)

type Profile struct {
	Destination string `json:"destination,omitempty" gorm:"not null"`
	Protocol    string `json:"protocol,omitempty" gorm:"not null"`
	FilePath    string `json:"file_path,omitempty" gorm:"not null"`
	Name        string `json:"name,omitempty" gorm:"not null;primarykey"`
}

func (profile *Profile) Validate() error {
	if profile.Destination == "" {
		return errors.New("destination is mandatory")
	}

	if profile.Protocol != "tcp" && profile.Protocol != "udp" {
		return errors.New("protocol needs to be tcp or udp")
	}

	if profile.FilePath == "" {
		return errors.New("file path is mandatory")
	}

	return nil
}
