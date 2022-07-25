Adapter Pattern in Go

Is a structural design pattern that enables collaboration between objects with incompatible interfaces. Implementing this pattern we can adapt the behavior between objects and allow to works each other.

Imagine a car and a boat, both are different objects If we implement the Adapter pattern we can achieve that the boat goes to the road and the car goes to the water.

Let us see a code example:

In this first part of the code we can see a client struct that drives the “navigation”, the navigation is different by car and boat there is where we going to apply the adapter.

In this example, we have a car that goes to the destination using roads and a boat that navigates in the water, as we can see the action to move from the start point to the destination is different.


Now, we need to adapt the car to go ahead in the water to do that we need the adapter struct that defines a method “navigate To Destination” that going to adapt the car to do the needed actions.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Bridge, State, Strategy, and to some extent Adapter. They have very similar structures. In fact, all of these patterns are based on composition, which consists of delegating work to other objects. However, they all solve different problems.
The Adapter Pattern changes the interface of an existing object while the Decorator Pattern enhances an object without changing its interface. Also, Decorator supports recursive composition, which is not possible when using Adapter.
Bridge Pattern is usually designed in advance, which allows you to develop modules independently. On the other hand, Adapter Pattern is often used with existing implementations to make them work together seamlessly.
Pros and Cons

You can separate the interface or data conversion code from the primary business logic of the program.
Open-closed principle. You can introduce new types of adapters to the program without modifying existing code, by working with the adapters through the interface.

The complexity of the code increases since you have to introduce a bunch of new interfaces and structures.