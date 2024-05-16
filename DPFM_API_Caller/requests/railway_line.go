package requests

type RailwayLine struct {
	Station				int		`json:"Station"`
	RailwayLine			int		`json:"RailwayLine"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	CreateUser			int		`json:"CreateUser"`
	LastChangeUser		int		`json:"LastChangeUser"`
	IsReleased			*bool	`json:"IsReleased"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
