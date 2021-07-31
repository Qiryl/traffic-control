### traffic-control

#### List of what was used:
* [gorilla/mux](https://github.com/gorilla/mux)
* [gocsv](https://github.com/gocarina/gocsv) 
* [viper](https://github.com/spf13/viper)
* [logrus](https://github.com/Sirupsen/logrus)

#### Some concepts used in development:
* During the development process, a clean architecture approach was used.
* The file system was used as data storage. Data is stored in csv files. The storage format is as follows:
      
      date,number,velocity
      ...
      ...
      date1,number1,velocity1
      
* Ability to set access time to requests. In the configuration file, you can set the `start` and `end` time for accessing requests.

#### Endpoints
* `/create/{date}/{number}/{velocity}` - Create a new entry
* `/all` - Get all entries
* `/number/{number}` - Get entries by car number
* `/date/{date}` - Get entries by date
* `/velocity/{velocity}` - Get entries by car velocity
* `/limit/{date}/{velocity}` - Get entries in which cars have exceeded the speed for the specified date
* `/minmax/{date}` - Get the minimum and maximum velocity for the specified date

#### Build
    make
#### Start
    ./api
