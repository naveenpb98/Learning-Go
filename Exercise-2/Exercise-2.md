# Exercise 34

## Problem Statement

Write a program that reads integers from standard input until it encounters an end-of-file (EOF) marker. As it reads each integer, the program should add the integer to a running total. Once the program encounters an EOF marker, it should print the final sum of all the integers it read.

Input

The input consists of a sequence of integers, each terminated by a newline character. The input terminates when the user types an end-of-file (EOF) marker, which typically occurs by pressing Ctrl+D on Unix-like systems.

Output

The output consists of a single line containing the sum of all the integers that were read from standard input.

Examples

Input

10

20

30

Output

60

Constraints

    The input may contain any number of integers.
    The integers can be positive, negative, or zero.

Hints

    Use the fmt.Fscanln() function to read an integer from standard input.
    Use an infinite loop to read integers until an EOF marker is encountered.
    Use a variable to accumulate the sum of the integers.
    Use the fmt.Println() function to print the final sum.