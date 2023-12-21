# Take Home Test

Thanks for your interest in BCF.
Please implement the requirements below and treat this as you would a real project.

When complete please tar or zip the project and send it to us via email or
create a git repo that we can pull.

## Requirements
Given a file path and two date templates, write a program that can transform 
the dates in the file from one date format to the other.

- Implement in your language of choice, Scala is preferred but all are valid
  - If your language supports functional constructs that is also preferred.
- Include steps on how to build and run your solution, including how to run tests.
- Allow us to interact with your code, how we interact is up to you.
  - eg via cli, http request, test, etc
  - include steps on how to use the program
- Any date matching the input format should be transformed to the output format. 
- Any date not matching the input format can be left alone. 
- All error conditions should be handled.
- If a requirement is ambiguous please make a decision and note the ambiguity.
- If you make some assumptions, make note of them.

A possible version could implement the following interface
```scala
def transformDates(filePath: String, sourcePattern: String, destinationPattern: String) = ???
```

## Examples:
```text
Dear diary, 

On 9/6/78 a really neat thing happened.  But on 10/10/03 thing went poorly.
I am hopeful that 10/11 will be better.
```

Would become 
```text
Dear diary, 

On September 06, 1978 a really neat thing happened.  But on October 10, 2003  thing went poorly.
I am hopeful that 10/11 will be better.
```

```text
April 26th, 2008 was a day in time.
```
would become
```text
2008-04-26 was a day in time.
```
