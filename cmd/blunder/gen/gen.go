package gen

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/georgi-georgiev/blunder"
	"github.com/spf13/cobra"
)

var CmdGen = &cobra.Command{
	Use:   "run",
	Short: "Run project",
	Long:  "Run project. Example: kratos run",
	Run:   Gen,
}

func Gen(cmd *cobra.Command, args []string) {
	tmpl, err := template.New("blunder.tmpl").ParseFiles("blunder.tmpl")
	if err != nil {
		panic(err)
	}

	errorCodes := make([]blunder.ErrorCode, 0)
	for reasonCode, reason := range blunder.Reasons {
		errorCodes = append(errorCodes, blunder.ErrorCode{
			Status:      strconv.Itoa(reason.ReasonGroup.Status) + " " + http.StatusText(reason.ReasonGroup.Status),
			Title:       reason.ReasonGroup.Title,
			Description: reason.ReasonGroup.Description,
			Resolution:  reason.ReasonGroup.Resolution,
			Code:        int(reasonCode),
			Reason:      reasonCode.String(),
			Message:     reason.Message,
			Tip:         reason.Tip,
		})
	}

	var f *os.File
	f, err = os.Create("blunder.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, errorCodes)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
