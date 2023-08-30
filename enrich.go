package blunder

func (b *Blunder) Enrich(httpErr HTTPError) HTTPError {
	if b.typeURI != nil {
		httpErr.Type = *b.typeURI
	}

	return httpErr
}
