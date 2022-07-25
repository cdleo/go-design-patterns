Factory Patterns in Go

Is a creational design pattern and it is also one of the most commonly used patterns.

This pattern provides a way to create instances using a “factory” struct type and a method that helps to create the instances according to the provided values and takes as a base a struct with common attributes.

Example:
Suppose that we are developing a “Robot factory” software, then we need to build two types of robots, a teacher robot, and a fighter robot.

We could define two approaches for this learning article, the first one is a factory function, and the second one is the abstract factory. The following section will be described each one.

Implement a Factory Method.
“A factory defined by a function that helps us to create instances of a certain structure with determinate values or values that could be provided in the function arguments.”


In the above code, We can see the factory function in line 26, that function will return an instance of “Robot” with the provided values. Then each time that we need to create an instance we can call the function instead of to created directly with the struct.

Output:
As you can see the two Robots were created using the factory function


Implement an Abstract Factory
“An Abstract factory defines a function that helps us to create multiple objects of the same family, that creation is done calling the specific factory function related to the object that we want to create. each one of the family objects is defined with a struct and that struct is composed by a base struct that defines the family properties.”

I going to show you an example to make it clearer:

In this first code snipe, we can see the struct that defines the “family-type” in this case robot (line 18), that struct will be used by each one of the “object types” that we want to create.


In the image below we can see the other important part of the Abstract factory, which are the types that will be created and its factory function, each one with its specific values.


The following code shows the function factory, as we can see is used to create teacher and fighter robots depending on the value passed in the param


And here the factory usage


Output


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Most of the Design patterns start with implementing the Factory method which is one of the most straightforward and most personalizable, and when the project and the code are growing sometimes evolves to Abstract Factory, Prototype, o Builder whos are more flexible but more complicated.

Prototype for example requires the initialization of an object that will be implemented in its logic, to do that we can implement a factory method.

Abstract Factory could be implemented instead of Facade Pattern when we just need to hide the object creation.

The Factory Method can serve as a stepping stone to a larger Template Method.

You can implement a Prototype Pattern to define the Abstract Factory methods.

Pros and Cons

Avoid a strong anchor between the creator and the concrete products.
Single responsibility principle. You can move the product creation code to a place in the program, making the code easier to maintain.
Open-closed principle. You can incorporate new product types into the program without breaking the existing client code.
You can be sure that the products you get from a factory are compatible with each other.

It may be that the code gets complicated since you have to incorporate a multitude of structures and compositions to implement the pattern. The ideal situation would be to insert the pattern into an existing hierarchy of creator structs.