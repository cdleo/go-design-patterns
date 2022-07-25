State Pattern in Go

This pattern allows an object to change behavior when its internal state changes. This pattern is related to a finite-state machine, and each state represents a different behavior executed by the state handler.

Let us explain this pattern with an example. suppose that you have a vending machine and there are defined the following states, choose a product, insert money and give the product. Based on that let us go ahead with the code.

In the first image, we can see defined the interface to abstract all the methods implemented by the different states. In this example we only have two, select the item and dispense the item.


In the images below we can see both states and the method implementations. And there we can see the logic depending on the state.



After we have defined the needed states we can go ahead and define the structure that going to manage each one. In the image below we can see that logic.

As we can see in the struct “vendingMachine” we are defining the needed states as property and we are defining one more to keep the current status.

Another this that we can observe in the code is the implementation of a “state” interface, in this case inside of each method we call the methods defined by the current state.


At this point, we can see how it is executed.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

State and Strategy Patterns are similar to each other because both patterns allow changing the behavior and their main differences are the following:
On one hand Strategy Pattern makes these objects completely independent and unknown to each other, in the other hand, State Pattern does not restrict dependencies between particular states, allowing them to alter the state of the context at will.
Pros and Cons

Single responsibility principle. Organize code related to particular states in separate structs.
Helps to implement the Open-closed Principle. Allows to Introduce new states without changing the existing state struct or context.
Allows simplifying the context code by removing bulky state machine conditionals.

Applying the pattern may be overkill if a state machine has only a few states or rarely changes.