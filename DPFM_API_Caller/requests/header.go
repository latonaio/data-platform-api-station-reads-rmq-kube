package requests

type Header struct {
	Station							int		`json:"Station"`
	StationType						string	`json:"StationType"`
	StationOwner					int		`json:"StationOwner"`
	StationOwnerBusinessPartnerRole	string	`json:"StationOwnerBusinessPartnerRole"`
	PrimaryLine						int		`json:"PrimaryLine"`
	PersonResponsible				*string	`json:"PersonResponsible"`
	URL								*string	`json:"URL"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	AvailabilityOfParking			*bool	`json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			*int	`json:"NumberOfParkingSpaces"`
	Site							int		`json:"Site"`
	SiteType						string	`json:"SiteType"`
	AirPort							*int	`json:"AirPort"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	CreateUser						int		`json:"CreateUser"`
	LastChangeUser					int		`json:"LastChangeUser"`
	IsReleased						*bool	`json:"IsReleased"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
