# Exercise 3: Replace Strings
## Problem Statement

Write a program that reads lines of text from standard input, replaces all occurrences of the first command-line argument with the second command-line argument, and prints the modified lines to standard output.

Input

The program is invoked with two command-line arguments:

    The old string to be replaced
    The new string to be substituted

Output

The output consists of the modified lines of text, with all occurrences of the old string replaced with the new string. Each line of the output is terminated by a newline character.

Examples


Input

This is an example line of text.
This is another example line of text.

Output

This is an example line of new-string.
This is another example line of new-string.


Hints

    Use the strings.ReplaceAll() function to replace all occurrences of the old string with the new string.
    Read the input lines using the bufio.Scanner type.
    Write the modified lines to standard output using the fmt.Println() function.