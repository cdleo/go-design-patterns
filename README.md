# Design Patterns in GO

## What a Design Pattern is?
The design pattern is a part of the software architecture that helps us to organize the code that we are writing because provides a general, reusable and applicable solution to different software design problems.

For example, you detect in your code the next situation: “You are detecting a lot of lines that are creating a struct instance, some of the assigned values are similar for some of the instances”. Here is when a design pattern comes to the help for example we can go ahead and define an approach to implement a “Factory Pattern” and then save all these similar/duplicated lines.

Think of another case: “you want to initialize a struct only once and prevent new instances related to that type, all in order to keep the value during execution”. So he remembers a design pattern called “Singleton” and makes the decision to implement it.

As you can see with those little examples, a design pattern could be described as a template that helps to identify problems in or software and offer an appropriate solution for each specific case.

All the examples and information in this repository has the intention of help you to understand what Pattern is and how could be implemented in GO.

## Design Patterns Types
We can detect three types that are grouped according to the pattern purpose, we have creational patterns, structural patterns, and behavior patterns. In the following section are described each one of the pattern groups.

**Creational Patterns**
These patterns provide object creation mechanisms that increase the flexibility and reuse of existing code.

- [Factory](https://github.com/cdleo/go-design-patterns/blob/main/creational/factory)
- [Builder](https://github.com/cdleo/go-design-patterns/blob/main/creational/builder)
- [Singleton](https://github.com/cdleo/go-design-patterns/blob/main/creational/singleton)

**Structural patterns**
These patterns explain how to assemble objects and classes into larger structures while maintaining the flexibility and efficiency of the structure.

- [Adapter](https://github.com/cdleo/go-design-patterns/blob/main/structural/adapter)
- [Bridge](https://github.com/cdleo/go-design-patterns/blob/main/structural/bridge)
- [Composite](https://github.com/cdleo/go-design-patterns/blob/main/structural/compisite)
- [Decorator](https://github.com/cdleo/go-design-patterns/blob/main/structural/decorator)
- [Proxy](https://github.com/cdleo/go-design-patterns/blob/main/structural/proxy)
- [Flyweight](https://github.com/cdleo/go-design-patterns/blob/main/structural/flyweight)

**Behavior patterns**
These patterns deal with algorithms and the assignment of responsibilities between functionalities.

- [Strategy](https://github.com/cdleo/go-design-patterns/blob/main/behavioral/strategy)
- [State](https://github.com/cdleo/go-design-patterns/blob/main/behavioral/state)
- [Iterator](https://github.com/cdleo/go-design-patterns/blob/main/behavioral/iterator)
- [Mediator](https://github.com/cdleo/go-design-patterns/blob/main/behavioral/mediator)
- [Visitor](https://github.com/cdleo/go-design-patterns/blob/main/behavioral/visitor)

## Contributing
Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
