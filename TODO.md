# TODO list for Goment
* goment.go
    * fromISOString
        * uses whatever timezone is parsed for now, need to figure out if we should convert to Local. 
    * fromExistingTime
        * should the method convert the time to Local?
* iso.go
    * need to handle time formats with commas, like 15:04:05,9999
    * need to investigate nanosecond parsing - when parsing something like 2011-04-02 03:04:05.10
        * becomes 2011-04-02 03:04:05.1 +0000 UTC
        * if construct same date with time.Date(2011, 4, 2, 3, 4, 5, 10, time.UTC), it becomes 2011-04-02 03:04:05.00000001 +0000 UTC
    * combine isoDateFormat & isoTimeFormat structs to common struct
    * need better test coverage
    * examine regexs and figure out if all are needed
* add_subtract.go
    * make sure the normalization of dates is well documented and consistent, whether it defaults to how Go handles it (Nov 31st becomes Dec 1st) or how Moment.js handles it (Nov 31st becomes Nov 30th)
* get_set.go
    * make sure the normalization of dates is well documented and consistent, whether it defaults to how Go handles it (Nov 31st becomes Dec 1st) or how Moment.js handles it (Nov 31st becomes Nov 30th)
    * figure out how to handle values greater than what is valid for field. Does it overflow to the next field? Handle negative values also. Related to point above.
    * [implement Minimum method](https://momentjs.com/docs/#/get-set/min/)
    * [implement Maximum method](https://momentjs.com/docs/#/get-set/max/)
* start_end_of.go
    * simplify repeated units across files
* format.go
    * support fractional seconds, like S SS SSS SSSS...
* parse.go
    * support fractional seconds, S SS SSS
* _test.go
    * add messages to asserts