# CSCI4030U: Big Data Project Part 1
# makefile
# Author: Michael Valdron
# Date: Feb 23, 2018
build:
	go build
clean:
	go clean
rebuild:
	go clean
	go build
run:
	./frequent-item-sets -alg=$(alg) -f="$(file)" -t=$(t)
run-apriori-retail:
	./frequent-item-sets -alg=a -f="retail.dat" -t=$(t)
run-pcy-retail:
	./frequent-item-sets -alg=p -f="retail.dat" -t=$(t)
run-apriori-netflix:
	./frequent-item-sets -alg=a -f="netflix.data" -t=$(t)
run-pcy-netflix:
	./frequent-item-sets -alg=p -f="netflix.data" -t=$(t)
