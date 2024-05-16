package dpfm_api_output_formatter

import (
	"data-platform-api-station-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Station,
			&pm.StationType,
			&pm.StationOwner,
			&pm.StationOwnerBusinessPartnerRole,
			&pm.PrimaryLine,
			&pm.PersonResponsible,
			&pm.URL,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.DailyOperationStartTime,
			&pm.DailyOperationEndTime,
			&pm.Description,
			&pm.LongText,
			&pm.Introduction,
			&pm.OperationRemarks,
			&pm.PhoneNumber,
			&pm.AvailabilityOfParking,
			&pm.NumberOfParkingSpaces,
			&pm.Site,
			&pm.SiteType,
			&pm.Airport,
			&pm.Project,
			&pm.WBSElement,
			&pm.Tag1,
			&pm.Tag2,
			&pm.Tag3,
			&pm.Tag4,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Station:							data.Station,
			StationType:						data.StationType,
			StationOwner:						data.StationOwner,
			StationOwnerBusinessPartnerRole:	data.StationOwnerBusinessPartnerRole,
			PrimaryLine:						data.PrimaryLine,
			PersonResponsible:					data.PersonResponsible,
			URL:								data.URL,
			ValidityStartDate:					data.ValidityStartDate,
			ValidityEndDate:					data.ValidityEndDate,
			DailyOperationStartTime:			data.DailyOperationStartTime,
			DailyOperationEndTime:				data.DailyOperationEndTime,
			Description:						data.Description,
			LongText:							data.LongText,
			Introduction:						data.Introduction,
			OperationRemarks:					data.OperationRemarks,
			PhoneNumber:						data.PhoneNumber,
			AvailabilityOfParking:				data.AvailabilityOfParking,
			NumberOfParkingSpaces:				data.NumberOfParkingSpaces,
			Site:								data.Site,
			SiteType:							data.SiteType,
			Airport:							data.Airport,
			Project:							data.Project,
			WBSElement:							data.WBSElement,
			Tag1:								data.Tag1,
			Tag2:								data.Tag2,
			Tag3:								data.Tag3,
			Tag4:								data.Tag4,
			CreationDate:						data.CreationDate,
			CreationTime:						data.CreationTime,
			LastChangeDate:						data.LastChangeDate,
			LastChangeTime:						data.LastChangeTime,
			CreateUser:							data.CreateUser,
			LastChangeUser:						data.LastChangeUser,
			IsReleased:							data.IsReleased,
			IsMarkedForDeletion:				data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.Station,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
			&pm.EmailAddress,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
			Station:                   data.Station,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
			EmailAddress:            data.EmailAddress,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.Station,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalSubRegion,
			&pm.LocalRegion,
			&pm.Country,
			&pm.GlobalRegion,
			&pm.TimeZone,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
			&pm.XCoordinate,
			&pm.YCoordinate,
			&pm.ZCoordinate,
			&pm.Site,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

		data := pm
		address = append(address, Address{
			Station:       	data.Station,
			AddressID:   	data.AddressID,
			PostalCode:  	data.PostalCode,
			LocalSubRegion: data.LocalSubRegion,
			LocalRegion: 	data.LocalRegion,
			Country:     	data.Country,
			GlobalRegion: 	data.GlobalRegion,
			TimeZone:	 	data.TimeZone,
			District:    	data.District,
			StreetName:  	data.StreetName,
			CityName:    	data.CityName,
			Building:    	data.Building,
			Floor:       	data.Floor,
			Room:        	data.Room,
			XCoordinate: 	data.XCoordinate,
			YCoordinate: 	data.YCoordinate,
			ZCoordinate: 	data.ZCoordinate,
			Site:		 	data.Site,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToRailwayLine(rows *sql.Rows) (*[]RailwayLine, error) {
	defer rows.Close()
	railwayLine := make([]RailwayLine, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.RailwayLine{}

		err := rows.Scan(
			&pm.Station,
			&pm.RailwayLine,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &railwayLine, err
		}

		data := pm
		railwayLine = append(railwayLine, RailwayLine{
			Station:				data.Station,
			RailwayLine:			data.RailwayLine,
			ValidityStartDate:		data.ValidityStartDate,
			ValidityEndDate:		data.ValidityEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			CreateUser:				data.CreateUser,
			LastChangeUser:			data.LastChangeUser,
			IsReleased:				data.IsReleased,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &railwayLine, nil
	}

	return &railwayLine, nil
}

