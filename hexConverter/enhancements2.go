cd flogo executables folder 

instead of running app ./appname_darwin_amd64

enabling metrics

FLOGO_HTTP_SERVICE_PORT=7777 FLOGO_APP_METRICS=true ./appname_darwin_amd64

enabling prometheus metrics 

FLOGO_HTTP_SERVICE_PORT=7779 FLOGO_APP_METRICS_PROMETHEUS=true ./<app-binary> //prometheus



> output.txt //saves output to that file 

> output.txt 2>&1 //throws in to file appending 



70,000 over 6.5 hours 

45 unique

7 per hour

65 rides 

3 dt pts/ sec

selected from 70,000 from flogo to streambase

