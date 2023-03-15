# TIME LIBRARY FOR ECLA

# Candidate functions :

| Func Name |                      Prototype                       |                          Description                           | Comments |
|:---------:|:----------------------------------------------------:|:--------------------------------------------------------------:|:--------:|
|   Date    | Date(year, month, day, hour, min, sec int) string {} |           Returns a string representation of a date            | UTC Time |
|    Now    |                   Now() string {}                    |      Returns a string representation of the current time       |   N/A    |
|   Sleep   |                  Sleep(sec int) {}                   | Pauses the current goroutine for a specified number of seconds |   N/A    |
|   Timer   |                  Timer(sec int) {}                   |            Waits for a specified number of seconds             |   N/A    |
