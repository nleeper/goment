# TODO list for Goment
* goment.go
    * fromISOString
        * uses whatever timezone is parsed for now, need to figure out if we should convert to Local. 
    * fromExistingTime
        * currently converts the time to Local. Should it?
* iso.go
    * need to handle YYYYYY date formats, like +002006-01-02
    * need to handle time formats with commas, like 15:04:05,9999
    * combine isoDateFormat & isoTimeFormat structs to common struct
    * need better test coverage
    * examine regexs and figure out if all are needed
* add.go
    * make sure the normalization of dates is well documented and consistent, whether it defaults to how Go handles it (Nov 31st becomes Dec 1st) or how Moment.js handles it (Nov 31st becomes Nov 30th)
* subtract.go
    * make sure the normalization of dates is well documented and consistent, whether it defaults to how Go handles it (Nov 31st becomes Dec 1st) or how Moment.js handles it (Nov 31st becomes Nov 30th)
* getters.go
    * [implement Weekday method, make locale aware](https://momentjs.com/docs/#/get-set/weekday/)
    * [implement Week method, make locale aware](https://momentjs.com/docs/#/get-set/week/)
    * [implement WeekYear method, make locale aware](https://momentjs.com/docs/#/get-set/week-year/)
    * [implement WeeksInYear method, make locale aware](https://momentjs.com/docs/#/get-set/weeks-in-year/)
    * [implement ISOWeeksInYear method, should be number of weeks in year by ISO weeks](https://momentjs.com/docs/#/get-set/iso-weeks-in-year/)
* setters.go
    * make sure the normalization of dates is well documented and consistent, whether it defaults to how Go handles it (Nov 31st becomes Dec 1st) or how Moment.js handles it (Nov 31st becomes Nov 30th)
    * figure out how to handle values greater than what is valid for field. Does it overflow to the next field? Related to point above.
    * [implement SetWeekday method, make locale aware](https://momentjs.com/docs/#/get-set/weekday/)
    * [implement SetWeek method, make locale aware](https://momentjs.com/docs/#/get-set/week/)
    * [implement SetISOWeek method, should set the ISO week of the year](https://momentjs.com/docs/#/get-set/iso-week/)
    * [implement SetWeekYear method, make locale aware](https://momentjs.com/docs/#/get-set/week-year/)
    * [implement SetISOWeekYear method, should set the ISO week-year](https://momentjs.com/docs/#/get-set/iso-week-year/)