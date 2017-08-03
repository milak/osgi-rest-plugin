package main
import (
	"log"
	"github.com/milak/tools/osgi"
)
var SymbolicName 	string
var Version 		string
const SERVICE_NAME 	= "AMQP"
var logger *log.Logger
func init(){
	SymbolicName 	= "amqp-connector"
	Version 		= "1.0"
}
func Start(context osgi.BundleContext){
	logger = osgi.GetLoggerFromContext(context)
	logger.Println("INFO Starting osgi-rest plugin")
}
func Stop(context osgi.BundleContext){
  logger.Println("INFO Stopping osgi-rest plugin")
}
