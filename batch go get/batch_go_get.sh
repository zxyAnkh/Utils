#!/bin/bash
for d1 in `ls ${GOPATH}/src`
do
	for d2 in `ls ${GOPATH}/src/${d1}`
	do
		fn=`ls ${GOPATH}/src/${d1}/${d2} | grep .go`
		if [ -z "$fn" ]; 
			then
				for d3 in `ls ${GOPATH}/src/${d1}/${d2}`
				do
					echo ${GOPATH}/src/${d1}/${d2}/${d3}
					`cd ${GOPATH}/src/${d1}/${d2}/${d3}|go get -u`
				done
			else
				echo ${GOPATH}/src/${d1}/${d2}
				`cd ${GOPATH}/src/${d1}/${d2}|go get -u`
		fi
	done
done