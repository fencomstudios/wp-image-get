This is a re-hash of a script that I originally wrote in C# for someone who needed to download all of the images from a Wordpress installation to migrate to another platform, and they needed to know what images went with what content. This program is a bit simpler in that it doesn't track what images were used where, but it does use the same input file that I used before - the Wordpress Export XML. I didn't use any of that old code when writing this - in fact, I didn't even open that old file. I wanted to do my best to try this one from scratch using as much of the knowledge as I could from the boot.dev Go course, and also learning how to read the Go package docs to try and piece together how to make this work.

I specifically tested this with the "Media" version of the export file, so I'm not sure if it will work with other versions of the file. I think the most that it will do is yell at you because it can't write a file thinking it's a directory or something like that.

## Get started
1. Clone the repo
2. Use "go build" to build the executable for your device.
3. Run the executable; it'll prompt you for the name of the XML file and the download rate in seconds, then it'll get going!

## Things to note:
* The downloads folder is created in the same directory as the executable, so make sure it's somewhere you want it to be before you start.
* It's currently just spitting out errors for anything other than files, but I'll add in those checks to make it a bit more graceful when I'm able.
* You can put the XML anywhere, but it's probably easiest to put it in the same directory as the executable so that you don't have to type or copy a lot of characters
