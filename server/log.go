package server

import (
    "log"
    "os"
    "io"
    "io/ioutil"
    "strings"
    "time"
)

var (
    LogTrace   *log.Logger
    LogInfo    *log.Logger
    LogWarning *log.Logger
    LogError   *log.Logger
    LogToFile	*log.Logger
)

func LogInit(traceHandle io.Writer,
    infoHandle io.Writer,
    warningHandle io.Writer,
    errorHandle io.Writer, prefix string){

	LogTrace = log.New(traceHandle,
        "TRACE: ",
        log.Ldate|log.Ltime)

    LogInfo = log.New(infoHandle,
        "INFO: ",
        log.Ldate|log.Ltime)

    LogWarning = log.New(warningHandle,
        "WARNING: ",
        log.Ldate|log.Ltime)

    LogError = log.New(errorHandle,
        "ERROR: ",
        log.Ldate|log.Ltime)

    DateNow := time.Now().Local().Format("20060102")
    FileName := "Log-"+DateNow+".txt"

    file, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalln("Failed to open log file")
	}

	LogToFile = log.New(file,
	    prefix+": ",
	    log.Ldate|log.Ltime)


}

func Logging(msg string,prefix string){
	//log record of everything

	upperPrefix := strings.ToUpper(prefix)

	LogInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, upperPrefix)

	if(upperPrefix == "ERROR"){
		LogError.Println(msg)
	} else if(upperPrefix == "INFO"){
		LogInfo.Println(msg)
	} else if(upperPrefix == "WARNING"){
		LogWarning.Println(msg)
	} else {
		LogTrace.Println(msg)
	}
	
    LogToFile.Println(msg)
}

