package mongo

import (
	"time"

	"github.com/andreandradecosta/rpimonitor"
	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SampleService is responsible for interacting with Mongo DB.
type SampleService struct {
	mongoSession *mgo.Session
}

//NewSampleService connects do Mongo and returns a service instance.
func NewSampleService(url string) (*SampleService, error) {
	session, err := mgo.Dial(url)
	return &SampleService{
		mongoSession: session,
	}, errors.Wrap(err, "error connecting to mongo")
}

//Query returns a collections of samples stored in the specified time interval.
func (s *SampleService) Query(start, end time.Time) ([]rpimonitor.Sample, error) {
	if s.mongoSession == nil {
		return nil, errors.New("No session to Mongo")
	}
	session := s.mongoSession.Copy()
	defer session.Close()
	c := session.DB("rpimonitor").C("samples")
	result := make([]rpimonitor.Sample, 1)
	err := c.Find(bson.M{
		"localTime": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}).Select(bson.M{
		"localTime":             1,
		"timestamp":             1,
		"metrics.load":          1,
		"metrics.virtualMemory": 1,
		"metrics.swapMemory":    1,
		"metrics.temperature":   1,
	}).Sort(
		"-localTime",
	).All(&result)
	return result, errors.Wrapf(err, "query failed [%s]-[%s]", start, end)
}

//Write stores a sample metrics in the database.
func (s *SampleService) Write(sample *rpimonitor.Sample) error {
	if s.mongoSession == nil {
		return errors.New("No session to Mongo")
	}
	defer s.mongoSession.Refresh()
	return s.mongoSession.DB("rpimonitor").C("samples").Insert(sample)
}
