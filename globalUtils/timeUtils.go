package globalUtils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"goTemp/globalerrors"
	"log"
	"time"
)

// DateLayoutISO defines the default time format for dates entered as strings
const DateLayoutISO = "2006-01-02"

// glErr holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

// TimeStampPPBToTime converts gRPC timestamp to time.time. Can take multiple timestamps at once. Returns a slice with
// the resulting time.times
func TimeStampPPBToTime(timeStamps ...*timestamp.Timestamp) ([]time.Time, error) {
	var returnTimes []time.Time
	for _, timeStamp := range timeStamps {
		newTime, err := ptypes.Timestamp(timeStamp)
		if err != nil {
			log.Printf(glErr.DtProtoTimeStampToTimeStamp(timeStamp, err))
			// log.Printf("Unable to convert proto timestamp %v to timestamp  when trying to update promotion. Error: %v \n", timeStamp, err)
			return nil, err
		}
		returnTimes = append(returnTimes, newTime)
	}
	return returnTimes, nil

}

// TimeToTimeStampPPB converts time.time to gRPC timestamp. Can take multiple time.times at once. Returns a slice with
// the resulting timestamps
func TimeToTimeStampPPB(times ...time.Time) ([]*timestamp.Timestamp, error) {
	var returnStamps []*timestamp.Timestamp
	for _, currentTime := range times {
		newStamp, err := ptypes.TimestampProto(currentTime)
		if err != nil {
			log.Printf(glErr.DtTimeStampToProtoTimeStamp(currentTime, err))
			// log.Printf("Unable to convert timestamp %v to proto timestamp  when trying to update promotion. Error: %v \n", currentTime, err)
			return nil, err
		}
		returnStamps = append(returnStamps, newStamp)
	}
	return returnStamps, nil
}

// CheckValidityDates checks if the validFrom - ValidThru timestamp based time frame is valid
func CheckValidityDates(validFrom *timestamp.Timestamp, validThru *timestamp.Timestamp) ([]string, error) {
	var FailureDesc []string
	validDates := true
	if validFrom == nil {
		FailureDesc = append(FailureDesc, glErr.MissingField("valid from"))
		validDates = false
	}
	if validThru == nil {
		FailureDesc = append(FailureDesc, glErr.MissingField("valid thru"))
		validDates = false
	}
	if validDates {
		vd, err := TimeStampPPBToTime(validFrom, validThru)
		if err != nil {
			return nil, err
		}
		if vd[0].After(vd[1]) || vd[1].Equal(vd[0]) {
			FailureDesc = append(FailureDesc, glErr.DtInvalidValidityDates(vd[0], vd[1]))
		}
	}
	return FailureDesc, nil
}

// GetNextYearTimeStamp returns a timestamp that is exactly one year from today
func GetNextYearTimeStamp() *timestamp.Timestamp {
	myDates, _ := TimeToTimeStampPPB(time.Now().AddDate(1, 0, 0))
	return myDates[0]
}

// TimeStringToTimestamp converts the string representation of a date to a gRPC timestamp. String must match the DateLayoutISO
func TimeStringToTimestamp(dateStr string) (error, *timestamp.Timestamp) {
	dateTime, err := time.Parse(DateLayoutISO, dateStr)
	if err != nil {
		log.Fatalf("Unable to Format date %v", dateStr)
	}
	dateTsProto, err := ptypes.TimestampProto(dateTime)
	if err != nil {
		log.Fatalf("Unable to convert time to timestamp %v", dateTime)
	}
	return err, dateTsProto
}
