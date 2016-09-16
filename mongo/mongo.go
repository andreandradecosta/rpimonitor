package mongo

import (
	"errors"
	"time"

	"github.com/andreandradecosta/rpimonitor"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SampleService struct {
	mongoSession *mgo.Session
}

func NewSampleService(url string) (*SampleService, error) {
	session, err := mgo.Dial(url)
	return &SampleService{
		mongoSession: session,
	}, err
}

func (s *SampleService) Query(start, end time.Time) ([]rpimonitor.Sample, error) {
	if s.mongoSession == nil {
		return nil, errors.New("No session to Mongo")
	}
	session := s.mongoSession.Copy()
	defer session.Close()
	c := session.DB("rpimonitor").C("samples")
	var result []rpimonitor.Sample
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

	return result, err
}

func (s *SampleService) Write(sample rpimonitor.Sample) error {
	if s.mongoSession == nil {
		return errors.New("No session to Mongo")
	}
	defer s.mongoSession.Refresh()
	return s.mongoSession.DB("rpimonitor").C("samples").Insert(sample)
}