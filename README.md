# *NOTE*
Palegen is currently in development and output colours may shift slightly between versions. It's useable but for the moment I recommoned a one time generation rather than including it as part of an automated build.

# Palegen

CLI based automated CSS colour palette generator. (Pronounced "pal-ee-jen".)

Palegen is a command line application that accepts a single colour, generates a harmonius palette and outputs the colours to a file as CSS variables.

![screenshot of palegen html page showing output colours](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300_thumb.png)

Inspired by Javascript based [Palx](https://github.com/jxnblk/palx), Palegen (for the moment) lacks some of Palx's sohpistication.

Palegen has one advantage in that it runs on the command line and can be part of an automated workflow.

Written in Golang Palegen uses the [go-colorful](https://github.com/lucasb-eyer/go-colorful) library for colour conversions.

## Usage

A single color in CSS hex style without the hash symbol should be passed when calling Palegen. For example using #ff3300

```bash
palegen ff3300
```

## Output

Colours are output in hsl format as CSS variables to a file named 'colors.css'. Example CSS output from the above command is [here](https://github.com/adrinux/palegen/blob/master/example/colors.css). 'colors.css' will be saved where the command is run from.

If you run palegen in the example folder you can view the results by opening palegen.html in your browser.
An example screenshot of that page is [here](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300.png).

## Installation

The easiest way at present is using go get to download, compile and install:

```bash
go get github.com/adrinux/palegen
```

(At some point I should set up binary releases via github.)

## Flaws, Issues and other things of note

1. Colour Names: For certain input colours the names of output colours can seem a little off. Red may be perceptually orange, yellow perceptually green. See issue #4
2. HSL vs HCL colour space. HCL can produce better colour distribution but is not directly supported by CSS. HSL is used by Palx and Palegen follows that. For full discussion see issue #5
3. Limited to hsl() CSS output. See #8
4. Limited to hex input. See #9
