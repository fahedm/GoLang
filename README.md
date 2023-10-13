
# GoLang

Learning GoLang

# Introduction

- GoLang was developed at Google in 2007 by Ken Thompson, Rob Pike and Robert Griesmer  and made open source in 2009.
- It is cross platform and open source pragramming language
- It has fast runtime and compile time.
- It is statically typed and compiled language.
- It supports concurrancy, which makes it suitable for systems where multi-processing is essential.
- It has garbage collector which automatically manages memory.

## Uses

- Web Development (Server-side)
- Network programming
- Cross-plaform enterprise applications
- Cloud-native Development

## Some disadvatages

- It has limited Object-Oriented features
- It doesn't support Classes and Objects.
- It doesn't support Inheritance.
- Comparatively new in the industry.

## Prducts build using GoLang

Docker, Kubernetes, openshift, DropBox, Netflix, Golang etc.

## Go is compiled

- It is compiled and statically type ( Variables have specific types).
- Go feels like a intereted language as Go tool can run a file without precompiling.
- In Background, Go will temporarily create a compiled executable, these compiled executables are OS specific.
- Compiled apps contains a statically linked runtime.
- No External Virtual machine is required.

## Go is partially object oriented

- Go have interfaces concept.
- we can define types. custom types can implement 1 or more interfaces.
- Custom Types can have member methods.
- custome structs (data structures) can have member fields.
- Types and Structs feels like class.
- Type Inheritence is not supported.
- mupltiple methods(in a type) or package with same name. Method or Operator Overloading is not supported.
- Go has Structured Exception Handling.
- Implicit numeric conversions are not allowed.

## Code Sematics

- Designed as next-generation language for C. Hence, GO borrows lot of syntax from C
- It also borrows syntax from Pascal, Modula and Oberon.
- Go tries to reducing amount of typing.

### Syntax

- Go is case sensitive.
- variables and packages names are lower case or mixed case
- Public Field name  has to start with uppercase character.
- Inital Upper case character means the symbol is exported.
- An initial uppercase character, means that that symbol is exported. means it is public
- A lowercase initial character means that the field or method isn't exported and isn't available to the rest of the application, or you could say that it's private.
- Semi colon is required to terminate the statement but we don't write it always as end of line means termination of statement, ; will be put automatically while compilation.
- Go is sensitive to white spaces, i.e. line feeds, tabs, etc.
- builtin package in always available.

### Variables

- Each variables must assigned type, and once assigned it can't be changed.
- 1. var keyword is used to create a variable. (var variableName type = value).
- 2. := signe is also used, in this we don't have to define type, compiler decides the type of the variable, based on the value. (variablename := value)
- It is not possible to declare a variable using :=, without assigning a value to it.
- var can be used inside and outside of functions. With var, Declaration and assignment can be done separately.
- := can only be used inside functions. with :=, declaration and assignment must be done in the same line.

### Multiple Variables and rules

- with GO, you can define multiple variables in one line.
- if you use any type, it is only possible to declare that type of variables.
- If type is not defined, then we can have different tyeps of variable, we just have to use :=.
- Naming rules, are same as any other programming language.
- Only starts with char or underscore, can't be a keyword, case-sensitive.
- We can use camelCase, PascalCase, snake_case etc.
- Constant variable can be created using "const" keyword. it will be unchangeable and read-only.
- Constant names are usually written in uppercase letters for easy identification and differentiation from variables. These can be declared both inside and outside of a function.
- you can declare constant with type or without type.
- Multiple constants can be grouped together into a block for readability.

### In Basic Hello World Program

- package: each program should be a part of a package. we need to define to which package current program belongs to.
- import: each program may need some packages to be imported. In basic program we will import 'fmt'/ format package to call Println and other functions.
- func: it is keyword which will be used to define for a function block.
- Execution will start from main package.
- Statements end with a newline or semicolon. left curly bracket '{' can't be first character of new line.
- If required we can write multiple statements in one line, seperated by ';'
- Single line comments can be given after //
- MultiLine comments can be given inside /*...*/

### Running the Program

- In package where you program reside, run "go run \<filename\>". you can run "go run ." if there is only one file. there can be only one main package. which will be executed by this command.
- If you want to create a build. run "go build ." it is create a build for that package.

### GO output using format package (fmt)

- Print() function is used to print only given variable.
- Printf() can be used to format the print statement using "%v" for value and %T for type or variable.
- Println() can print arguments,  whitespace is added between the arguments, and a newline is added at the end.

### Data Types

- int- stores integers (whole numbers), such as 123 or -123

  1. Signed integers (int) - can store both positive and negative values
  2. Unsigned integers (uint)- can only store non-negative values
  3. default is int type.

- float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
- string - stores text, such as "Hello World". String values are surrounded by double quotes
- bool- stores values with two states: true or false. Default value is False

### contants

- The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
  