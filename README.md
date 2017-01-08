#Learn Go

This is a public repository for those wanting to learn the basics of Go.

##Course Outline
1. [Basics](basics/basics.md)
	1. [Printing](basics/printing/printing.md)
	1. [Values](basics/values/values.md)
	1. [Variables and Constants](basics/vars-consts/vars-consts.md)
	1. [Operators](basics/operators/operators.md)
	1. [If Else Statements](basics/if-else/if-else.md)
	1. [For Loops](basics/for/for.md)
	1. [Challenge 1: Count to 20 with Even and Odd](challenges/basics/20-even-odd/20-even-odd.md)
	1. [Arrays](basics/arrays/arrays.md)
	1. [Slices](basics/slices/slices.md)
	1. [Switches](basics/switches/switches.md)
	1. [Maps](basics/maps/maps.md)
	1. [Range](basics/range/range.md)
	1. [Functions](basics/functions/functions.md)
	1. [Challenge 2: Check if Prime](challenges/basics/check-prime/check-prime.md)
1. [Intermediate](intermediate/intermediate.md)
	1. [Advanced Functions](intermediate/adv-func/adv-func.md)

##Getting Setup with Go

Installing Go isn't hard but it can be a bit frustrating for newcomers. This short guide will take you through downloading, installing and setting up Go for the first time.

###Download Go

Downloading Go is the first step in this process, and it's also the easiest. To get started, visit <https://golang.org/dl/> and download the featured download appropriate for your operating system.
The featured downloads are easy installers that will do most of the hard work for you. If you are on Windows, download the .msi file, on Mac, download the .pkg file and on linux, download the .tar.gz file.
Now that you have the correct installer downloaded for your OS, run the installer and follow the simple steps to get Go installed.

Finally, to check that you have Go installed, open up a new command line prompt and just type Go, you should get a message listing all the Go commands. If that doesn't work, make sure you installed Go properly then restart your computer and try the Go command again.

###Configuring Go

There are only a few things that you will need to manually configure when installing Go. Start off by entering into the command line, `go env`. This command should print out the details of your Go environment. Settings such as the GOROOT should already be set, the Go Root is where the files for running Go are stored. Importantly, you need to set the GOPATH to wherever you will store your Go programs. To do this, you have to change
