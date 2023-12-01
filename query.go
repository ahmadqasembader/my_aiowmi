package main


// import "fmt"
import (
	"fmt"
	"strings"
)


type Query struct{
	query string
	namespace string
	language string

	qc QContext
}


// Acts as the constructor for the Query
func (q *Query) NewQuery(query string) *Query{
	// Defautling the querying language and the namespace
	q.language = "WQL"
	q.namespace = "root/cimv2"
	q.query = query

	// Prefixing the namespace with the appropriate format if not already
	if !strings.HasPrefix(q.namespace, "//"){
		q.namespace = "//./" + q.namespace
	}
	return q
}

func (q Query) Context(conn Connection, proto Protocol) QContext{
	flags := WBEM_FLAG_RETURN_IMMEDIATELY | WBEM_FLAG_FORWARD_ONLY
	timeout :=  60
	skip_optimize := false
	return QContext{q.query, conn, proto, flags, timeout, skip_optimize}
}

// TODO: should be an aysnc function
func (q *Query) Start(conn Connection, proto Protocol, flags, timeout int){
	timeout = 60
	qc := QContext{q.query, conn, proto, flags, timeout, true}
	fmt.Println(qc)
	// TODO: await for this function
	// qc.start()
}
