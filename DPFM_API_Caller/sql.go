package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-station-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-station-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner
	var railwayLine *[]dpfm_api_output_formatter.RailwayLine

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "HeadersByStations":
			func() {
				header = c.HeadersByStations(mtx, input, output, errs, log)
			}()
		case "HeadersBySite":
			func() {
				header = c.HeadersBySite(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "Partners":
			func() {
				partner = c.Partners(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Addresses":
			func() {
				address = c.Addresses(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegion":
			func() {
				address = c.AddressesByLocalSubRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegions":
			func() {
				address = c.AddressesByLocalSubRegions(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegion":
			func() {
				address = c.AddressesByLocalRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegions":
			func() {
				address = c.AddressesByLocalRegions(mtx, input, output, errs, log)
			}()
		case "RailwayLine":
			func() {
				railwayLine = c.RailwayLine(mtx, input, output, errs, log)
			}()
		case "RailwayLines":
			func() {
				railwayLine = c.RailwayLines(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                header,
		Partner:               partner,
		Address:               address,
		RailwayLine:           railwayLine,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.Station = %d", input.Header.Station)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.Station ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.Station ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByStations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	log.Info("HeadersByStations")
	in := ""

	for iHeader, vHeader := range input.Headers {
		station := vHeader.Station
		if iHeader == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				station,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			station,
		)
	}

	where := fmt.Sprintf(" WHERE ( Station ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.Station ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySite(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Site = %d", input.Header.Site)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.Station ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	station := input.Header.Station
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, station, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_partner_data
		WHERE (Station, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY Station ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partners(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	station := input.Header.Station
	partner := input.Header.Partner

	cnt := 0
	for _, _ = range partner {
		args = append(args, station)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_partner_data
		WHERE (Station) IN ( `+repeat+` ) 
		ORDER BY Station ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, v.Station, v.Station)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data
		WHERE (Station, AddressID) IN ( `+repeat+` ) 
		ORDER BY Station ASC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Addresses(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	station := input.Header.Station
	address := input.Header.Address

	cnt := 0
	for _, _ = range address {
		args = append(args, station)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data
		WHERE (Station) IN ( `+repeat+` ) 
		ORDER BY Station ASC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalSubRegion(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	where := "WHERE 1 = 1"

	if input.Header.Address[0].LocalSubRegion != nil {
		where = fmt.Sprintf("WHERE address.LocalSubRegion = \"%s\"", *input.Header.Address[0].LocalSubRegion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data AS address
		` + where + ` ORDER BY address.LocalSubRegion ASC, address.Station ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalSubRegions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {

	log.Info("AddressesByLocalSubRegions")
	in := ""

	for iAddress, vAddress := range input.Header.Address {
		localSubRegion := vAddress.LocalSubRegion
		if iAddress == 0 {
			in = fmt.Sprintf(
				"( '%s' )",
				localSubRegion,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%s' )",
			in,
			localSubRegion,
		)
	}

	where := fmt.Sprintf(" WHERE ( LocalSubRegion ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data AS address
		` + where + ` ORDER BY address.LocalSubRegion ASC, address.Station ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalRegion(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	where := "WHERE 1 = 1"

	if input.Header.Address[0].LocalRegion != nil {
		where = fmt.Sprintf("WHERE address.LocalRegion = \"%s\"", *input.Header.Address[0].LocalRegion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data AS address
		` + where + ` ORDER BY address.LocalRegion ASC, address.Station ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalRegions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {

	log.Info("AddressesByLocalRegions")
	in := ""

	for iAddress, vAddress := range input.Header.Address {
		localRegion := vAddress.LocalRegion
		if iAddress == 0 {
			in = fmt.Sprintf(
				"( '%s' )",
				localRegion,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%s' )",
			in,
			localRegion,
		)
	}

	where := fmt.Sprintf(" WHERE ( LocalRegion ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_address_data AS address
		` + where + ` ORDER BY address.LocalRegion ASC, address.Station ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) RailwayLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.RailwayLine {
	var args []interface{}
	where := fmt.Sprintf("WHERE Station = %d ", input.Header.Station)

	railwayLineIDs := ""
	for _, v := range input.Header.RailwayLine {
		railwayLineIDs = fmt.Sprintf("%s, %d", railwayLineIDs, v.RailwayLine)
	}

	if len(railwayLineIDs) != 0 {
		where = fmt.Sprintf("%s\nAND RailwayLine IN ( %s ) ", where, railwayLineIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_railwayLine_data
		`+where+` ORDER BY IsMarkedForDeletion ASC, IsReleased ASC, Station ASC, RailwayLine ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToRailwayLine(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) RailwayLines(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.RailwayLine {
	railwayLine := &dpfm_api_input_reader.RailwayLine{}
	if len(input.Header.RailwayLine) > 0 {
		railwayLine = &input.Header.RailwayLine[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Station != 0 {
		where = fmt.Sprintf("WHERE Station = %d", input.Header.Station)
	}

	if railwayLine != nil {
		if railwayLine.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND railwayLine.IsReleased = %v", where, *railwayLine.IsReleased)
		}
		if railwayLine.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND railwayLine.IsMarkedForDeletion = %v", where, *railwayLine.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_railwayLine_data as railwayLine
		` + where + ` ORDER BY railwayLine.IsMarkedForDeletion ASC, railwayLine.IsReleased ASC, railwayLine.Station ASC, railwayLine.RailwayLine ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToRailwayLine(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
