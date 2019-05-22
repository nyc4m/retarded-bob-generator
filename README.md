# Retarded sponge bob generator

Generate this mainstream meme easier, with one command line ! :joy::100::ok_hand:

![Spongebob](./res/bob_source.png)



## Installation

This program uses the impact font, as well as a png file. to pack this two resources easily, I used [`go.rice`](https://github.com/GeertJohan/go.rice). To **install the program** you need to run 2 commands

```bash
$ rice go-embbed //compile resources
$ go install //actually install the program
```

## Usage

For now you can generate a *retarded* sentence just like sponge bob do, and either generate the meme, or getting the text output

```bash
./retarded-bob-generator -input="I'm not retarded, leave me alone" -text=false -out="retarded_bob.png"
```

this command will create an image `retarded_bob.png` with the string provided in `input`. If you prefer to only get the string, just put `text` to true.

All the commands are described in the help section

```bash
$ ./retarded-bob-generator --help
```

The text is auto centered, you really just have to run one command :sunglasses:
