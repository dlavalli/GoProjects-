package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
)

// To use a subpackage variable in other code outside of this subfolder,
// we must name the variable with the first letter as an uppercase letter
var Session *gocql.Session

/*
When you import submodules in Go, any packages containing a method called init() will have all of those
init() functions called immediately before the process’ main() function. The caution here is that if you have
several submodules with init() functions declared, there is no guarantee of run order so you must ensure
you don’t inadvertently impose a required order by having mixed dependencies. If you must have these
initializations happen in a specific order, it’s best to declare an exported method in each submodule
and call them explicitly from your process’ main() function
*/
func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "streamdemoapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}