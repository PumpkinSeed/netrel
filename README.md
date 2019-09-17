# netrel
Internet reliability check - CLI tool

netrel analyze the reliability of the internet connection. It's sending ICMP packets to the trusted hosts many time, and based on the analyzed result determine a percentage how good the connection is.

### Usage

```
go install github.com/PumpkinSeed/netrel

sudo netrel
```

##### Flags

`--print-meta`: Print the complete analyzed data.

##### Sample output

```
$ sudo netrel
===============================================

Tested hosts:
	1.1.1.1
	google-public-dns-a.google.com
	google-public-dns-b.google.com
	139.130.4.5
Final score of internet reliability: 75.137545% 
Test spent: 43.574424145s
```

