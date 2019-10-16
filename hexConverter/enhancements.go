
/// FLOWS and DEFINITIONS 


// LOG POWER DATA PAGE 10

string.concat("Power message received - page # ", string.substring($flow.message, 0, 2))

//conditional 

string.substring($flow.message, 0, 2) == "10"


// LOG POWER BATTERY PAGE 

string.concat("Power sensor battery status received - page # ", string.substring($flow.message, 1, 1), "Battery level is low.")

// conditional 

(string.substring($flow.message, 0, 2) == "52") && (string.substring($flow.message, 15, 1) == "4")


// LOG NON DATA POWER PAGE 

string.concat("Non-data page - Power Message = ", $flow.message) 

// conditional 

(string.substring($flow.message, 0, 2) != "10") && (string.substring($flow.message, 0, 2) != "52" && (string.substring($flow.message, 0, 2) != "12")) // changing to add DATA PAGE 12 for cadence 



MQTT PUBLISH 

string.concat("{\"deviceType\":", string.tostring($flow.deviceType) ,  ", 
 \"deviceTime\": ",  string.tostring($activity[DecodeEventCountPower].result), ",
   \"accPower\": ",  string.tostring($activity[DecodeAccPower].result), "}")


   string.concat("{\"deviceType\":", string.tostring($flow.deviceType) ,  ", 
 \"dataType\": ",  string.tostring(XXXXX), ",
   \"data"\": [{", ", \"name\:", string.tostring(XXXX),  ", \"value\:", string.tostring(XXXX), 
   "\},", "\{", \"name\:", string.tostring(XXXX),  ", \"value\:", string.tostring(XXXX), "\}," "\]", "}")


   {
	"deviceType": "11",
	"dataType": "Cadence",

	"data": [{
			"name": "AvgCadence",
			"value": 60
		},
		{
			"name": "AvgCadence2",
			"value": 65
		}
	]

}

EDITING NEW MQTT PUBLISH 



string.concat("{\"deviceType\":", string.tostring($flow.deviceType), ",\"dataType\":", "Cadence", ",\"data\": [{", ",\"name\":", "AvgCadence", ",\"value\":", string.tostring($activity[DecodeInstCadence].result), "},", "{\"name\": AvgCadence2,",  "\"value\":", 65, "}]", "}")



string.concat("{\"deviceType\":", string.tostring($flow.deviceType), ",\"dataType\":", "Cadence", ",\"data\": [{", ",\"name\":", "AvgCadence", ",\"value\":", string.tostring($activity[DecodeInstCadence].result), "},", "{\"name\": AvgCadence2,",  "\"value\":", 65, "}]", "}")

Final Edit in BikeDemov4 // Friday June 28th 11:19 AM 

string.concat("{\"deviceType\":", string.tostring($flow.deviceType), ",\"dataType\":", "Cadence", ",\"data\": [{", ",\"name\":", "AvgCadence", ",\"value\":", string.tostring($activity[DecodeInstCadence].result), "},", "{\"name\": AvgCadence2,",  "\"value\": 65 }]}")


Tibco.123


192.168.0.100


In log take out 2nd sample value 


string.concat("{\"deviceType\":", string.tostring($flow.deviceType), ",\"dataType\":", "Cadence", ",\"data\": [{", ",\"name\":", "AvgCadence", ",\"value\":", string.tostring($activity[DecodeInstCadence].result), "}]}")



In the Data Page 10 for Power we have the accPower & instPower info in the payload 


