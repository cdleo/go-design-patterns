Flyweight Pattern in Go

Flyweight is a structural design pattern that allows programs to support large numbers of objects while keeping memory usage low.

The pattern accomplishes this by sharing parts of the object’s state between multiple objects. In other words, the Flyweight saves RAM memory by caching the same information used by different objects.

Let us see an example:






Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

You could have a Composite tree and define shared nodes implementing Flyweights in order to save RAM.
Flyweight helps create many small objects, while Facade shows how to create a single object that represents an entire subsystem.
Flyweight could be similar to Singleton if you could somehow reduce all the shared states of the objects to a single flyweight object. But there are two fundamental differences between these patterns:
1. There should only be one Singleton instance, whereas Flyweight can have multiple instances with different intrinsic states.
2.- Singleton object can be mutable. Flyweight objects are immutable.
Pros and Cons

It helps save a lot of RAM, as long as your program has a lot of similar objects.

May be trading RAM for CPU cycles when some of the context information must be recomputed every time someone calls a flyweight method.
The code gets very complicated. New team members will always be wondering why an entity’s state was separated in such a way.