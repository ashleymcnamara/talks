package templates

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Course struct {
	Name              string    `json:"name"`
	ShortName         string    `json:"short_name"`
	Instructor        string    `json:"instructor"`
	InstructorEmail   string    `json:"instructor_email"`
	InstructorTwitter string    `json:"instructor_twitter"`
	OutputDirectory   string    `json:"output_directory"`
	Modules           []*Module `json:"modules"`
}

func NewCourse(cmd *cobra.Command) Course {
	return Course{
		Name:              cmd.Flag("course").Value.String(),
		ShortName:         cmd.Flag("shortname").Value.String(),
		Instructor:        viper.GetString("author"),
		InstructorEmail:   viper.GetString("email"),
		InstructorTwitter: viper.GetString("twitter"),
		OutputDirectory:   filepath.Join(viper.GetString("coursedir"), cmd.Flag("shortname").Value.String()),
		Modules:           make([]*Module, 0),
	}

}
