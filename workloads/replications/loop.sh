#!/bin/sh

a=0

while [ $a -lt 1000 ]
do
   echo $a
   curl 192.168.49.2:32000
   #a=`expr $a + 1`
done