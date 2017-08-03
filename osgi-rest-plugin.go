package main
import (
	"log"
	"github.com/milak/tools/osgi"
	"github.com/milak/tools/network"
)
var SymbolicName 	string
var Version 		string
var logger *log.Logger
func init(){
	SymbolicName 	= "osgi-rest"
	Version 	= "1.0"
}
func Start(context osgi.BundleContext){
	logger = osgi.GetLoggerFromContext(context)
	logger.Println("INFO Starting osgi-rest plugin")
	var objectMap map[string]interface{}
	context.SetProperty("rest-map",&objectMap)
	network.Listen("/","80",objectMap)
}
func Stop(context osgi.BundleContext){
	logger.Println("INFO Stopping osgi-rest plugin")
}
