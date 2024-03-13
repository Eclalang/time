# Time library


[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/time)](https://goreportcard.com/report/github.com/Eclalang/time)
[![codecov](https://codecov.io/gh/Eclalang/time/graph/badge.svg?token=YNCIYERVBO)](https://codecov.io/gh/Eclalang/time)

## Candidate functions :

|   Func Name   |                      Prototype                       |                                Description                                | Comments |
|:-------------:|:----------------------------------------------------:|:-------------------------------------------------------------------------:|:--------:|
| ConvertRoman  |          ConvertRoman(str string) string {}          |                    Converts a number to a roman number                    |   N/A    |
|     Date      | Date(year, month, day, hour, min, sec int) string {} |                 Returns a string representation of a date                 | UTC Time |
|      Now      |                   Now() string {}                    |            Returns a string representation of the current time            |   N/A    |
|     Sleep     |             Sleep(sec int \| float64) {}             |      Pauses the current goroutine for a specified number of seconds       |   N/A    |
|   Strftime    |       Strftime(format, date string) string {}        | Returns a string representation of a date according to a specified format |   N/A    |

## Supported Conversion Formats :
| Pattern |                   Description                    |
|:-------:|:------------------------------------------------:|
|   %d    |   Day of the month as a decimal number (01-31)   |
|   %H    | Hour (24-hour clock) as a decimal number (00-23) |
|   %M    |        Minute as a decimal number (00-59)        |
|   %m    |        Month as a decimal number (01-12)         |
|   %S    |        Second as a decimal number (00-59)        |
|   %Y    |                Year with century                 |
|   %y    |           Year without century (00-99)           |
|   %%    |                    A '%' sign                    |