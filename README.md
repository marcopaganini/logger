**NAME**

logger - A simple logging library for Go

**DESCRIPTION**

logger is a simple logging library for go. It extends the log package and
also provides 'verbose' and 'debug' messages support. To use:

    import "github.com/marcopaganini/logger"
    
    log *logger.Logger
    log = logger.New("prefix")
    log.SetVerboseLevel(1)
    log.SetDebugLevel(2)

    // All the usual log.* functions can be used normally, plus
    log.Verboseln(1, "This prints only if verbose level >= 1")
    log.Debugln(2, "This prints only if debug level >= 2")

    // log.Verbosef and logDebugf are also available

**AUTHOR**

(C) Aug/2014 by Marco Paganini <paganini AT paganini DOT net>
