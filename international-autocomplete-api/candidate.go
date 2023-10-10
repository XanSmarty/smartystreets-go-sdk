package international_autocomplete_api

type Candidate struct {
	Street                  string `json:"street"`
	Locality                string `json:"locality"`
	AdministrativeArea      string `json:"administrative_area"`
	SuperAdministrativeArea string `json:"super_administrative_area"`
	PostalCode              string `json:"postal_code"`
	CountryIso3             string `json:"country_iso3"`

	Entries     int    `json:"entries"`
	AddressText string `json:"address_text"`
	AddressID   string `json:"address_id"`
}
