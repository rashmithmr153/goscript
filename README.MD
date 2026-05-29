# goscript

goscript is a dynamically-typed scripting language built from scratch in Go, featuring variables, functions, loops, conditionals, lexical scoping, and an interactive REPL.

## Installation

```bash
git clone https://github.com/rashmithmr153/goscript.git
cd goscript
go build .
```

## Usage

### File mode

Write your program in a `.gs` file and run it:

```bash
go run . main.gs
```

### REPL mode

Start the interactive shell by running without arguments:

```bash
go run .
> let x = 10;
> print(x);
10
> exit
```

## Language Guide

### Variables

Variables are declared with the `let` keyword. goscript supports integers, strings, and booleans.

```
let x = 10;
let name = "goscript";
let flag = true;
```

### Arithmetic

goscript supports the standard arithmetic operators with correct precedence.

```
let x = 2 + 3 * 4;
let y = (2 + 3) * 4;
let z = 10 / 2 - 1;
```

### Comparisons

```
let a = 10 == 10;
let b = 5 != 3;
let c = 10 > 5;
let d = 3 <= 3;
```

### Conditionals

```
let x = 10;
if x > 5 {
    print("greater");
} else {
    print("lesser");
}
```

### Loops

goscript uses a while-style loop with the `for` keyword.

```
let i = 0;
for i < 5 {
    print(i);
    let i = i + 1;
}
```

### Functions

Functions are defined with the `func` keyword and support multiple parameters, return values, and lexical scoping.

```
func add(x, y) {
    return x + y;
}

let result = add(10, 5);
print(result);
```

Functions can access variables from their outer scope:

```
let base = 100;

func addBase(x) {
    return base + x;
}

print(addBase(5));
```

### Recursion

```
func fib(n) {
    if n < 2 {
        return n;
    }
    return fib(n - 1) + fib(n - 2);
}

print(fib(10));
```

### Types

| Type    | Example          |
|---------|------------------|
| int     | `let x = 42;`    |
| string  | `let s = "hi";`  |
| bool    | `let b = true;`  |

### Built-in Functions

`print(value)` - prints a value to stdout.

```
print(42);
print("hello");
print(true);
```

## Error Handling

goscript provides descriptive runtime errors for common mistakes:

```
> let x = 10 / 0;
runtime error: division by zero

> undefinedFunc();
runtime error: undefined function 'undefinedFunc'

> print();
runtime error: print() requires at least one argument
```

## Example Program

```
func greet(name) {
    print("hello");
    print(name);
}

func factorial(n) {
    if n <= 1 {
        return 1;
    }
    return n * factorial(n - 1);
}

greet("goscript");

let i = 1;
for i <= 5 {
    print(factorial(i));
    let i = i + 1;
}
```

## Project Structure

```
goscript/
    ast/        AST node types and evaluator
    lexer/      tokenizer
    parser/     recursive descent parser
    main.go     entry point, REPL and file execution
    main.gs     example program
```

## Implementation

goscript is built from scratch without any parsing libraries. The pipeline is:

```
source code
    -> lexer    (tokenizes raw text into tokens)
    -> parser   (builds an AST using recursive descent)
    -> evaluator (walks the AST and evaluates nodes)
```

Key concepts used: finite automata for lexing, context-free grammars for parsing, recursive tree evaluation, and lexical scoping with a parent environment chain.