# Palegen

(Pronounced: "pal-ee-jen".)

Command line based automated CSS colour palette generator.

Palegen is a command line application that accepts a single colour, generates a harmonius palette and outputs the colours to a file as CSS variables.

Inspired by the javascript based [Palx](https://github.com/jxnblk/palx) it lacks Palx's sohpistication.

Written in Golang Palegen has one advantage in that it will run on the command line and can be part of an automated workflow.

## Usage

A single color in CSS hex style minus the hash symbol (eg #ff3300 as ff3300) should be passed when calling Palegen.

```bash
palegen ff3300
```

## Output

Colours are output in hsl format as CSS variables.
