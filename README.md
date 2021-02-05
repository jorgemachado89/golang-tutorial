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
* Empty for loop conditions equal are considered as never ending ones unless explicitly terminated.
* For expression, range and post statement types exist
* Such as in Switch statements break statements stop the execution of the next execution block within the current scope
* Breakpoint and continue clauses are also available in Go
* Arrays elements need all to be of the same type
* Arrays have fixed lengths although Slices have a variable length
* Slices holds pointers to data elements in Array. Slices are references and as such are passed as references when dealing with function arguments.
* Slices hold the starting, length of Slice and optional capacity. Capacity sets the maximum size of the Slice.
* "make" keyword allocates and initializes a Slice.
* It is possible to recursively create Slices based on Slices.
* The capacity of the Slice if equal to the length of its underlying Array
* When appending values to a Slice the underlying Array will hold its capacity until it can no longer support additional values. Moreover, when adding a capacity + 1 value a new underlying Array will be created with double the previous capacity. The reference to the Slice does not mutate when appending new values.
* Maps are unordered and represent the implementation of a Hash table. Key value pairs.
* Maps are unsafe for concurrency
* When dealing with large sized maps, try to allocate before hand the expected size in order to improve performance.  
* Maps are passed by reference such as pointers and slices
* Concurrency deals with the aspect of having multiple processes executing independently not the same as in parallel execution
* Goroutines are not scheduled by the OS but instead via the Go runtime.
* Goroutines are lighter than OS Threads
* Less switching since it does not block due to OS dependencies as when dealing with OS Threads.
* Multiple Goroutines can be running on a single OS Thread.
* Goroutines allow for safe communication via channels.
* Go's concurrency uses the Actor model. Communicating Sequential Processes.
* Unbuffered channels locks the sender Goroutine till another Goroutine is ready to receive the payload.
* Buffered channels does not block the sender Goroutine unless the buffer is full thus locking the sender until a receiver comes and grabs data to release the channel 

## Questions

* Considering Golangs scripting, simplicity and concurrency habilities coupled with web application development how does compare when matched against NodeJS ?
    * https://github.com/golang/go/wiki/FromXToGo
* How do channels work through example ? None is provided during this course.

## TODO

- Checkout the following links and channels:
    - [ ] https://github.com/golang/go/wiki#contributing-to-the-go-project
    - [ ] https://blog.rubylearning.com/* -est-practices-for-a-new-go-developer-8660384302fc
    - [ ] https://github.com/golang/go/wiki/LearnServerProgramming
    - [ ] https://golang.org/doc/contribute.html
    - [ ] https://rakyll.org/go-tool-flags/
    - [ ] https://github.com/golang/go/wiki/GoUsers
    - [ ] https://github.com/golang/go/wiki/GoUserGroups
    - [ ] https://rakyll.org/cross-compilation/
    - [ ] https://github.com/golang/go/wiki/whygo

- Projects worth checking out:
    - [ ] https://github.com/satellity/satellity

