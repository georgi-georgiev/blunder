package blunder

import (
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func Generate() {
	var tmplFile = "blunder.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	errorCodes := make([]ErrorCode, 0)
	for reasonCode, reason := range Reasons {
		errorCodes = append(errorCodes, ErrorCode{
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
