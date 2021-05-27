# Functional Go

## Overview

It is planned to incorporate generics into the language, and when that happens the way we program in golang now will be history, and most of the existing code that pretended to be reusable will be obsolete.

This library is an attempt to allow gophers to start programming the way we will do after the incorporation of generics, helping our new code be reusable for longer time.

Despite the name, it is not a goal of this library to be fundamentally functional, but to offer some data types that programmers use frequently in languages with generics support.

To start with, I decide to include type constructors for data types that golang does not have:
- `Stream` a monadic immutable collection that evaluates lazily  
- `ValiResu` A monadic validation result with combinators that accumulate the errors. 
- `Validate` A monadic validation action that behaves like `ValiResu` but evaluates lazily.
- some other data types needed to support the above.

Other type constructors will be added over time with your active help ;)

>~~Note: The `ValiResult` and `Validate` data types are not fully monads because they use a map (instead of a set or sorted sequence) to accumulate the errors and therefore the associative rule is not satisfied when two instances contain different errors indexed with the same key.
> But they behave as monads provided distinct errors are associated to distinct keys, which is easy to achieve assigning different keys to errors coming from different places.~~ 

# Type constructor instantiation
This library contains not only a handful of type constructors, but also a tool to instantiate them and any other type constructor you add.

The type instantiation is implemented making a copy of the type constructor source code files (named templates), replacing all occurrences of the type parameters with the type arguments you desire.

Type constructors templates are normal golang source files that include directives inside comments to help the instantiator to resolve internal and external dependencies and control code expansion convergence. 

# Usage



     


