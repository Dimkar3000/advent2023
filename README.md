# My solutions for Advent 2023

This is another year when I took the time to solve the advent of code. My language of choice was `go`, and I managed to solve most of the problems with fast, at most cases, answers. I managed to solve all 25 problem for part 1, but I missed 3 part 2 solution where I consider the proposed solutions insane or too boring to implement.

# The problems I came across

A lot of the latest challenges essentially wanted you to solve part1 again but with insane requirements. I understand that I am not the audience for such problems, but maybe there would be more people solving the second half of advent if there were less of them. The ways to solve them varied for problem to problem, but it came down to 3 common patterns:

## Finding a shortcut by specializing your code to the input.

For example, there were problems where you would have to solve the same problem but for a huge grid. You could do it, but it would take hours. The input was crafted in a way where you could predict the output without calculating everything. I don't understand what is the value of such a problem, because normally in software you don't get to make such assumptions.

## Deep knowledge math and algorithms

Don't get me wrong, I love to implement algorithms but, scaling a problem to huge number only to make me use the LCM mathematical formula is not a fan. Or making me implement the problem in a certain way in order to avoid rounding errors was also very cumbersome. It essentially makes you write the formulas in a way that avoid division and is not always possible, to be fair it was, this time. Also, I had to implement custom Dijkstra implementations twice... TWICE😂

## Need for specialized libraries

In one of the problems, it was required that I have to solve a 2D linear equation and I did, following a guide 😅. On part 2 I had to solve multiple systems of 3d linear algebra equations, and that wasn't possible in a language without a SAT solver. All the solutions I saw were using such a tool to solve it for them. But then what the point of the exercise, to use a premade solution?

# My progress learning go

I grew a lot as a new `go` user this month. I when from someone was good at writing code in general, to some who now know how to use the language, mostly. Of course, because I spent most of my time solving programming exercises it doesn't mean I am a `go` expert, but at least I don't need to remind my self how to write a struct every day. The main pain points were essentially, getting used to the syntax, which was quirky at best, but you get used to it fast, and the luck of strong tools. Coming from languages with great feature to manage, process and edit lists/arrays, I can tell you that this process was painful in `go`. Rust, C++, Python, C#, JavaScript, java, all have better mechanisms to process data without me having to implement everything about the needed mutations, logic, iteration and so on. Maybe it's a skill issue and there is a strong package out there that will be the solution I needed, but I didn't find it. It is painfully obvious that the language is design to be minimal but some feature so common in modern languages, especially the garbage collected ones, that I expect them to be there. Just to name some, null access of fields, API functions on lists and maps, etc... Also, sorting was hard to figure out (I did, but the only reason I did was that the VS Code extension for `go` has a shortcut for sorting an array.) One last thing, does anybody understand how the module system work? Does the filename need to be the same with the folder on a new package?

# My final thoughts

I use advent as a collection of problems to get to understand a new language more in depth. This was the third year I try this approach and I love it. Specifically, the early days are fantastic, with bite sized problems that force you to get to understand your language better. Creating new data structure, parsing, code organization, even new libraries to use. Even on problems where I didn't get the solution on my own, I got to read other people's solutions and transpile them to `go`. This process made understand what is good, and what is bad in this language, at least in a CLI, problem-solving environment. The previous years I used rust, this year I used `go`, I don't believe it will be the same language again, but what will it be?
