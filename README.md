# Golang Tutorial

Followed Pluralsight tutorial [Go Fundamentals](https://app.pluralsight.com/library/courses/go-fundamentals)

## Notes

* Started as a *systems* language
* Syntax similar to C
* Garbage collecting
* Very small list of reserved keywords -> ~ =< 25
* Compiled language
* Strongly typed
* Supports Type Inference
* Go is case sensitive
* Publicly accessible methods start with Uppercase.
* Non declarations, at package level, outside methods are not allowed.
* Package level declarations even if not used are allowed but should be avoided.
* Go passes arguments by value
* Go supports passing multiple return values and naming those same values as part of the function signature
* Go supports first class functions.
* Go does not type cast strings or int to boolean for evaluating conditionals.
* Variables initializations can be made part of the if statements and are scoped to it.
* Conditional ifs expect a single else statement although multiple else if statements may be added. Guarantees single first match execution in lexical order.
* Switch statements accept like Conditional ifs an initializing statement and its case value can be of any basic type.
* Switch statements have implicit breaks.
* Switch case fallthrough statements allow for transfering control to the next clause although grouping case statements will lead to the same result.
* Switch case statements may have multiple values comma separated values as an alternative to grouping case statements
* Idiomatic to return errors as the last return. nil is used to indicate success.
* Underscore is used to ignore left hand values in an assignment and can be useful for importing a module disregarding the exported values making use only of its side effects. Can also be used to silence compiler errors on unused imports and variables.