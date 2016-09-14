package mongo

import (
	"time"

	"github.com/andreandradecosta/rpimonitor"

	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type SampleService struct {
	mongoSession *mgo.Session
}

func (s *SampleService) Query(start, end time.Time) ([]rpimonitor.Sample, error) {
	return nil, nil
}
