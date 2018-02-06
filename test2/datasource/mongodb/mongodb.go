package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/kataras/golog"
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
	dbName      = "calorii"
)

type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	Col     *mgo.Collection
}

func init() {
	if mainSession == nil {
		var err error
		mainSession, err = mgo.Dial("127.0.0.1")

		if err != nil {
			golog.Error(err)
		}

		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(dbName)
	}

}

func New() *MgoDb {
	session := mainSession.Copy()
	db := session.DB(dbName)

	return &MgoDb{Session: session, Db: db}
}

func (m *MgoDb) C(collection string) *mgo.Collection {
	m.Col = m.Session.DB(dbName).C(collection)
	return m.Col
}

func (m *MgoDb) Close() bool {
	defer m.Session.Close()
	return true
}

func (m *MgoDb) DropoDb() {
	err := m.Session.DB(dbName).DropDatabase()
	if err != nil {
		golog.Error(err)
	}
}

func (m *MgoDb) RemoveAll(collection string) bool {
	m.Session.DB(dbName).C(collection).RemoveAll(nil)

	m.Col = m.Session.DB(dbName).C(collection)
	return true
}

func (m *MgoDb) Index(collection string, keys []string) bool {

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := m.Db.C(collection).EnsureIndex(index)
	if err != nil {
		golog.Error(err)
		return false
	}

	return true
}

func (m *MgoDb) IsDup(err error) bool {

	if mgo.IsDup(err) {
		return true
	}

	return false
}
