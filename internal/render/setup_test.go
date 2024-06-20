package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sseudes108/Course_Go_Web_App/internal/config"
	"github.com/sseudes108/Course_Go_Web_App/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (writer *myWriter) Header() http.Header {
	var header http.Header
	return header
}

func (writer *myWriter) WriteHeader(integer int) {

}

func (writer *myWriter) Write(bytes []byte) (int, error) {
	lenght := len(bytes)
	return lenght, nil
}
