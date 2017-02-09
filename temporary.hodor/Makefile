all: jim.bin

jim.bin : ../jim.py
	( cd .. ; python ../rye/rye.py build jim.py )
	-rm -f jim.bin
	cp ../jim.bin .
	chmod 755 jim.bin
	chmod 644 temporary.conf temporary.boot.sh
	chmod 755 .

push: all
	(cd .. ; sh ../hodor/assimilate.sh temporary.hodor forth.yak.net)

clean:
	-rm -f jim.bin
