package masterdata

type Carrier struct {
	Id         *int64  `json:"id"`
	SharedId   *string `json:"sharedId"`
	InternalId *string `json:"internalId"`
	Name       *string `json:"name"`
	ShortName  *string `json:"short_name"`
}
