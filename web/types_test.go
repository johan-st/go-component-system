package web

import (
	"log"
	"os"
	"testing"
)

func Test_Example(t *testing.T) {
	t.Log("Example test")

	ses:= Session{}

	

}


func Test_FarFuture(t *testing.T){
	ses := Session{}
	ses.Add(CompLogger(log.New(os.Stderr, "test farFuture", log.LstdFlags )))
	ses.Add(CompCart)
}