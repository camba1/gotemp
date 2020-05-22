package globalUtils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"goTemp/globalerrors"
	"log"
	"time"
)

//glErr: Holds the service global errors that are shared cross services
var glErr globalerrors.SrvError

func TimeStampPPBToTime(timeStamps ...*timestamp.Timestamp) ([]time.Time, error) {
	var returnTimes []time.Time
	for _, timeStamp := range timeStamps {
		newTime, err := ptypes.Timestamp(timeStamp)
		if err != nil {
			log.Printf(glErr.DtProtoTimeStampToTimeStamp(timeStamp, err))
			//log.Printf("Unable to convert proto timestamp %v to timestamp  when trying to update promotion. Error: %v \n", timeStamp, err)
			return nil, err
		}
		returnTimes = append(returnTimes, newTime)
	}
	return returnTimes, nil

}

func TimeToTimeStampPPB(times ...time.Time) ([]*timestamp.Timestamp, error) {
	var returnStamps []*timestamp.Timestamp
	for _, currentTime := range times {
		newStamp, err := ptypes.TimestampProto(currentTime)
		if err != nil {
			log.Printf(glErr.DtTimeStampToProtoTimeStamp(currentTime, err))
			//log.Printf("Unable to convert timestamp %v to proto timestamp  when trying to update promotion. Error: %v \n", currentTime, err)
			return nil, err
		}
		returnStamps = append(returnStamps, newStamp)
	}
	return returnStamps, nil
}
