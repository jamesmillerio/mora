package api

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/mora/session"
	"labix.org/v2/mgo/bson"
)

type StatisticsResource struct {
	SessMng *session.SessionManager
}

// GET http://localhost:8181/stats/local/landskape
func (s *StatisticsResource) GetDatabaseStatistics(req *restful.Request, resp *restful.Response) {
	session, needsClose, err := s.SessMng.Get(req.PathParameter("alias"))
	if err != nil {
		handleError(err, resp)
		return
	}
	if needsClose {
		defer session.Close()
	}
	dbname := req.PathParameter("database")
	result := bson.M{}
	err = session.DB(dbname).Run(bson.M{"dbStats": 1, "scale": 1}, &result)
	if err != nil {
		handleError(err, resp)
		return
	}
	fmt.Printf("result:%#v", result)
	resp.WriteEntity(result)
}

// GET http://localhost:8181/stats/local/landskape/systems
func (s *StatisticsResource) GetCollectionStatistics(req *restful.Request, resp *restful.Response) {
	session, needsClose, err := s.SessMng.Get(req.PathParameter("alias"))
	if err != nil {
		handleError(err, resp)
		return
	}
	if needsClose {
		defer session.Close()
	}
	dbname := req.PathParameter("database")
	colname := req.PathParameter("collection")
	result := bson.M{}
	err = session.DB(dbname).Run(bson.M{"collStats": colname, "scale": 1}, &result)
	if err != nil {
		handleError(err, resp)
		return
	}
	fmt.Printf("result:%#v", result)
	resp.WriteEntity(result)
}
