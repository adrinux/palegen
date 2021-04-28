# Palegen

CLI based automated CSS colour palette generator. (Pronounced "pal-ee-jen".)

Palegen is a command line application that accepts a single colour, generates a harmonius palette and outputs the colours to a file as CSS variables.

![screenshot of palegen html page showing output colours](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300_thumb.png)

Inspired by Javascript based [Palx](https://github.com/jxnblk/palx), Palegen (for the moment) lacks some of Palx's sophistication.

Palegen has one advantage in that it runs on the command line and can be part of an automated workflow (but note that output colours may shift slightly between versions).

Written in Golang Palegen uses the [go-colorful](https://github.com/lucasb-eyer/go-colorful) library for colour conversions.

## Usage

Palegen reads from a palegen.ini configuration file, see palegen.ini for a commented example.
Currently you define a single input color and a destination for the output CSS. Then run palegen from the directory containing the ini file.

```bash
palegen
```

## Output

Colours are output in hsl format as CSS variables to a file named 'colors.css'. Example CSS output from the above command and default palegen.ini is [here](https://github.com/adrinux/palegen/blob/master/example/colors.css).

By default colors.css is output to the example folder, you can view the results by opening palegen.html in a browser.
An example screenshot of that page is [here](https://github.com/adrinux/palegen/blob/master/example/screenshot-Palegen-ff3300.png).

## Installation

Using go get to download, compile and install will give your the very latest code:

```bash
go get github.com/adrinux/palegen
```

It's more sensible to check out a specific version and compile it yourself:

```bash
git clone https://github.com/adrinux/palegen.git
cd palegen
git checkout tags/0.2.0
go build
```

And copy the resulting binary somewhere in your path (check it's executable).

(At some point I should set up binary releases via github...)

## Flaws, Issues and other things of note

1. Colour Names: For certain input colours the names of output colours can seem a little off. Red may be perceptually orange, yellow perceptually green. See issue #4
2. Limited to hsl() CSS output. See #8
