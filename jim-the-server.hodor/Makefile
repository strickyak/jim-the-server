all: jim.bin

jim.bin : ../jim.py
	( cd .. ; python ../rye/rye.py build jim.py )
	cp ../jim.bin .

push:
	(cd .. ; sh ../hodor/assimilate.sh jim-the-server.hodor aegon.yak.net)

clean:
	-rm jim.bin
