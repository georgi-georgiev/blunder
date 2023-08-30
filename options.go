package blunder

type BlunderOptions struct {
	Environment    Environment //stacktrace
	TypeURI        string
	Domain         string
	IsTraceable    bool //traceid,correlationid
	IsIdentifiable bool //include
	IsTimeable     bool //timestamp
	IsRecovarable  bool //recoverable, action
}
