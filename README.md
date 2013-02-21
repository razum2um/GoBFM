WTF?
====
Proof of concept, that we can fast translate XML => JSON using webservices in Go

HOWTO
=====

1) Install GVM

    bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)

2) Install Go-1.0.3

    gvm install go1
    gvm use go1

3) Setup envorinment

    ./install.sh

4) Launch webserver

    ./run.sh

5) Now let's curl!

    time curl http://127.0.0.1:9000/bfm/ --data @bfm100.xml

Results on my ASUS K43E
=======================

```
$ time curl http://127.0.0.1:9000/bfm/ --data @bfm.xml &> /dev/null

real0m0.028s
user0m0.005s
sys0m0.005s

$ time curl http://127.0.0.1:9000/bfm/ --data @bfm50.xml &> /dev/null

real0m0.148s
user0m0.011s
sys0m0.003s

$ time curl http://127.0.0.1:9000/bfm/ --data @bfm100.xml &> /dev/null

real0m0.281s
user0m0.018s
sys0m0.004s

$ time curl http://127.0.0.1:9000/bfm/ --data @bfm200.xml &> /dev/null

real0m0.537s
user0m0.537s016s
sys0m0.007s
```
