# PythonProblems

The best way to learn a language is by using it to solve problems. 

There are two ways to use python: running scripts, or interactively. You can start an interactive session by just typing `python` at the command line, and then entering individual commands. Using interactive python is mandatory - it’s essential for learning by quickly trying commands. You should probably get ipython (brew install ipython), and use that instead, because it has lots more features (like extensive documentation to tell you about individual python commands so you don’t even need to look them up on google).


For each problem, we describe the behaviour of the script when we run. The document is split in sections, with each section containing some problems and some things to read about.

For each section, make a folder (use mkdir from the command line). 
The answer to a problem, eg. Section 1, Problem 4, should be a standalone python file in the respective folder, eg. section1/problem4.py, that can be run as `python section1/problem4.py`. 

Before you do anything, learn vim. Vim is a text editor. No more TextEdit or XCode. You can have XCode back after a year of coding if you have really good reasons for it. Right now, you only have bad reasons for it. Learning vim will take some practice. But it makes you progressively more efficient over time. And it keeps you in the terminal, which is important. Use it to write a few diary entries. Don’t be afraid to learn more vim commands, ever. 

Side note, when you’re bored one day, learn how to encrypt the diary files you wrote in vim using GPG.

Ok, problems.

# Section 1 - Intro

Output “hello world”.  
The simplest program.
Ask the user how old they are. Output how many times older they are than a 4 year old (eg if the user enter 25, the output is 6.25)
requires prompting the user for input, and simple arithmetic.
Output the current price of bitcoin
requires fetching data from https://api.quadrigacx.com/v2/ticker and parsing the returned json value to get the “last” price

You should read up on the following:

- interpreted vs compiled languages
- stdin, stdout, stderr (how you send information and get it back from a process)
- python functions

# Section 2 - Functions

Repeat all of Section 1, but using functions in our scripts, and calling the function from a `main` function. You’ll have to read about functions and what a basic proper python program looks like. So redo Section 1 using programs that have a `main` function that calls the function doing the actual work. 

Read up on:

- command line programs. try out some of the command line programs you haven’t tried before
- grep. learn some grep. be able to find where a phrase is written in a directory of files.

# Section 3 - Command line arguments

Repeat problem 1.2 (ask the user how old they are), but instead of prompting them, take their age as a command line argument (so when you run the program, it’s just `python section3/problem1.py 25`)

Repeat problem 1.3 (bitcoin price), but take the currency pair as a command line argument (so when you run the program, it’s just `python section3/problem2.py btc_cad`). The url changes to: 
`https://api.quadrigacx.com/v2/ticker?book=btc_cad`
The other options are `btc_usd` and `btc_xau` (gold!). What happens if you enter something bogus?

Read up on:

- what’s an API and what’s it have to do with HTTP? learn to use curl. know about HTTP stuff like requests (GET, POST, DELETE, etc.), status codes (200, 201, 404, etc.), headers, etc.
- know the difference between a natural number, integer, rational number, and real number. 

# Section 4 - Basics

 From now on, we call command line programs `cli` s and we call their arguments `cli args`

1. Take an integer as a cli arg and output it’s factorial (the product of it and all integers smaller than it)
2. Take a list of numbers as cli args and output the average value of the list.
3. Take 3 ints ( the year, the month, the day) as cli args and output a pretty sentence. (eg. `python section4/problem3.py 2014 3 20`) would return “Today is March 20th, 2014”). If the date is invalid, it should tell the user there was an error.
4. Same as 3, but the cli arg is “yyyy-mm-dd” 
5. Take an array of integers as cli args and output them in sorted order
w.ll need to look up sorting in python
6. Redo 5 where we take an additional cli arg which determines whether we sort in increasing or decreasing order. 
7. Take a string as a cli arg and output its sha256 hash. You will have to import a crypto or hashing library in the script to do it. Try it out in the interactive python first.


Read up on:

- ASCII: how is text represented in a computer? What’s a byte? 
- Binary: how to convert numbers from base 10 to binary and back!
- Hex: how to convert numbers from base 10 to hex and back
- How to write a byte array as hex

