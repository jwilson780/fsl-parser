# FSL Interpreter 

FSL Interpreter is an interpreter that operates on the fictitious FSL programming language

## Installation

### Prerequisite 
Install [go](https://go.dev/doc/install) for your local environment.
On Mac with brew:
```bash
brew install go
```

Download this repo and navigate to the root folder.

## Usage
### Running
From the root folder of the repo run:
```bash
go run . corrected-sample-script.txt
```

This will run a single, well formatted, script. 

```bash
go run . basic-script.txt corrected-sample-script.txt
```
This will run two scripts, in sequential order.  Any variables used in the first script 
will be present in the second, but will be overwritten if they are present in the second.  In other words, 
everything persists between the scripts but it will be overwritten if there is a reference
to it in any subsequent script.

### What can it do?
This repo uses the [ANTLR](https://www.antlr.org/) project to build a parser and lexer based
on a provided grammar. The grammar in this project is in the [FSL.g4](FSL.g4) file.  It contains
all the rules needed to lex and parse the language.  I made a few concessions with this file based on the provided
sample-script.txt file.  It appears there is a standard library of sorts in there, with the add, create, update and 
delete functions. So I made those standard features in the language, as well as the ability to provide user defined
functions and variable declarations.

The interpreter supports the following:
* Variable declaration
```
var1: 1
var2: 2
var3: 0
```
* Standard functions that support both references and raw values
```
cmd:create, id: var1, value:3.5
cmd:update, id: var2, value:#var1
cmd:delete, id:var1

cmd:add, id:var1, operand1: 10, operand2: 5
cmd:subtract, id:var1, operand1: #var1, operand2: 5
cmd:multiply, id:var1, operand1: #var1, operand2: #var2
cmd:divide, id:var2, operand1: 10, operand2: 5

cmd:print, value:#var1
cmd:print, value:4
```
* User defined functions that support function parameters
```
sum: [
    cmd:add, id: $id, operand1: $value1, operand2: $value2
]

# calling with param values
cmd:#sum, id:var1, value1:#var1, value2:#var2

printAll: [
    cmd:print, value: #var1
    cmd:print, value: #var2
    cmd:print, value: #var3
]
```
For a script to work, you will need to ensure there is an init function within it.  You can run scripts 
without init functions, but they will need to call commands and functions directly.  Functions that are defined
will not just be called at the end of the script, only the init function.

### How it works

ANTLR will generate go code based on the provided .g4 file.  It provides files that can be seen in the /parser folder. 
These files are what allows me to walk the tree using my own code that is housed in the /interpreter folder. 

In the interpreter code, the program walks over the tree that is passed to it. Each of the methods housed there
allows the program to talk different actions based on the context that is passed in. 

Using ANTLR and its generated interfaces and methods, allows this program to easily extensible and easy to read.  Its 
easy to add new keywords in the grammar and react to them in the interpreter.

## License

[MIT](https://choosealicense.com/licenses/mit/)