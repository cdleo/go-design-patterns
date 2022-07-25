Bridge Pattern in Go

An official definition is the following:

The Bridge is a structural design pattern that allows you to split a large class, or a group of closely related classes, into two separate hierarchies (abstraction and implementation) that can be developed independently of one another.

Let us explain it with an example, suppose that you have two computers one desktop and a Laptop, and apart to have two different printers, you need to handle both printers on each one of your PCs.

Here is where this pattern comes into action because you could define the needed code in separate structs and interfaces to break the dependency to code each business logic for similar places.

Also applying this pattern could reduce the code coupling that is good for our software architecture.

In a few words, the Bridge pattern helps us to extend the business logic without the need to modify the previous code. But to do that we have to think about the implementation of this pattern and prepare our code to work in this way since we are defining the functionality.

I split the code into three images to make it easier to explain.

In the image below we can see the interface to abstract the computer actions. Also, we can see the implementation of the interface by the two kinds of computers used in this example, note that each one has a property named “printer” which is an interface type that allows to implementation of the bridge and call “print” method from the different printer.


In the image below, now we can see the printer interface and the implementation for the two printers.


And here the Bridge Pattern is in action, as we can see in the image in order to switch between printers we use the “setPrinter” method, and by doing that we can adapt the code to execute the different implementations.


Output


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Bridge Pattern is usually designed in advance, which allows you to develop modules independently. On the other hand, Adapter Pattern is often used with existing implementations to make them work together seamlessly.
Bridge, State, Strategy, and to some extent Adapter. They have very similar structures. In fact, all of these patterns are based on composition, which consists of delegating work to other objects. However, they all solve different problems.
Pros and Cons

Open/closed principle. You can introduce new abstractions and implementations independent of each other.
Single responsibility principle. You can focus on high-level logic in the abstraction and platform details in the implementation.

It may be that the code gets complicated if you apply the pattern to a very cohesive class.