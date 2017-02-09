all: jim.bin

jim.bin : ../jim.py
	( cd .. ; python ../rye/rye.py build jim.py )
	-rm -f jim.bin
	cp ../jim.bin .
	chmod 755 jim.bin
	chmod 644 yak.conf jimbo.boot.sh
	chmod 755 .

push: all
	(cd .. ; sh ../hodor/assimilate.sh jimbo.hodor aegon.yak.net)

clean:
	-rm -f jim.bin
