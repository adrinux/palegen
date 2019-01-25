# Palegen

Command line based automated CSS colour palette generator. (I pronounce it "pal-ee-jen".) 

Palegen is a command line application that accepts a single colour, generates a harmonius palette and outputs the colours to a file as CSS variables.

![screenshot](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300.png)

Inspired by javascript based [Palx](https://github.com/jxnblk/palx) Palegen (for the moment) lacks some of Palx's sohpistication.

Palegen has one advantage in that it runs on the command line and can be part of an automated workflow.

Written in Golang Palegen uses the [go-colorful](https://github.com/lucasb-eyer/go-colorful) library for colour conversions.

## Usage

A single color in CSS hex style without the hash symbol should be passed when calling Palegen. For example using #ff3300

```bash
palegen ff3300
```

## Output

Colours are output in hsl format as CSS variables to a file names 'colors.css'. Example CSS output from the above command is [here](https://github.com/adrinux/palegen/blob/master/example/colors.css). 'colors.css' will be saved whereever the command is run from.

If you run palegen in the example folder you can view the results by opening palegen.html in your browser.
A screenshot of that page is ![here](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300.png).


## Flaws, Issues and other things of note.

There are some minor issues with the output.