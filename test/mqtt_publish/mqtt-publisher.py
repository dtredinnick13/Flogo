import sys  
import os
import paho.mqtt.client as paho


def main():  
   filepath = "/Users/amberjaura/Downloads/mqtt_publish/bike-data.txt"

   broker="192.168.254.22"
   port=1883 
   mqttclient = paho.Client("control1")
   mqttclient.connect(broker,port)
   


   if not os.path.isfile(filepath):
       print("File path {} does not exist. Exiting...".format(filepath))
       sys.exit()

   with open(filepath) as fp:
       cnt = 0
       for line in fp:
           print(line)
           mqttclient.publish("org.teamtibco.bikesensor.rawdata",line)
           cnt+=1
           #print("line {} contents {}".format(cnt, line))
        
   print("total sent : ",cnt)

if __name__ == '__main__':  
   main()
