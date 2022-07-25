Visitor Pattern in Go

Is a behavioral design pattern. It is used when we have to perform an operation on a group of similar kinds of Objects. With the help of visitor patterns, we can move the operational logic from the objects to another class.

Also, the Visitor pattern allows you to add behaviors to a structure without modifying the structure.

The visitor pattern consists of two parts:

a method called Visit() which is implemented by the visitor and is called for every element in the data structure
visitable classes providing Accept() methods that accept a visitor
Let us have an example to explain more and see how it is implemented in Golang.

Imagine that you are working on an application to calculate the location and the area occupied by some buildings in the city. Each one of those places will form part of the building interface.

Now we have to add a behavior to get the build area, we can add directly the method to the “build” interface but that approach is not good enough because you want to ensure the open-close principle for future implementations.

As we can see in the following image, we are defining an interface named “visitor” and inside three methods that will help us to add new functionalities to our building objects.

Also, we can see in the interfaces that we are covering the rules for the visitor pattern, we have defined the “accept” and the “visitor” methods.


In the image below is defined the needed code for the concrete type “church”.


In the image below is defined the needed code for the concrete type “stadium”.


In the image below is defined the needed code for the concrete type “hotel”.


The following image shows the “visitor” that is used to calculate the area and as we can see we have three methods to manage each one of the buildings.

Note: Here we could implement the strategy pattern in order to handle all the building types


The following image shows the “visitor” that is used to calculate the coordinates and again here needs to be defined the methods to handle each one of the building types.

Note: Here we could implement the strategy pattern in order to handle all the building types


And in the last image, we can see the usage of this pattern.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

You can treat Visitor as a powerful version of the Command pattern.
You can use the Visitor pattern to perform an operation on an entire tree defined by the Composite pattern.
You can implement Visitor and Iterator patterns at the same time, in order to iterate through a complex data structure and perform some operations on its elements.
Pros and Cons

Helps to implement the open-close principle because with this pattern we can add/modify the behavior with out affect the main struct.
This pattern allows to implement single responsibility principle by declaring all the similar logic in different structs and interfaces.
A visiting object can accumulate some useful information while working with various objects. This can be useful when you want to traverse a complex object structure, such as an object tree, and apply the visitor to each object in that structure.

You must update all visitors each time a new struct is added to or removed from the element hierarchy.
Visitors may lack the necessary access to the private fields and methods of the elements they are supposed to work with.