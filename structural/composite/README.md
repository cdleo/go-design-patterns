Composite Pattern in Go

It is a structural design pattern that allows you to combine objects from the same “family” in a tree-like structure and work with it as if it were a single object. Allowing access to each of the objects to execute their methods.

This pattern is commonly used to solve the cases when we need to handle tree structures because is easy to iterate each one of its child/inner objects.

To define this pattern, we can use a concrete type (the struct type) or we can use the interface type to define the child objects.

Let us have an example.

Imagine that we have a file system and there we have folders and files. Now Imagine that you need to develop a code to search in the file system by a given string, there is when the “Composite Pattern” comes to action. The pattern will help us to organize the data structure and help to handle the folder and files iterations.

In the first image, we can see two things defined. The first one (in line 6) is the interface that will be used to abstract the types (folders and files). The second one is lines 11 to 21 there is defined the “file” struct that implements the “component” interface.


In the image below, we can see defined the structure used to define “folders”, as we can see this struct has a defined property of type “component” and is an array, that is because a folder contains folders and files.


And finally the usage in this case the main function.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

You can use Builder when creating complex Composite trees because you can program its construction steps to work recursively.
Chain of Responsibility is often used in conjunction with Composite. In this case, when a leaf component receives a request, it can pass it along the chain of all parent components back to the root of the object tree.
Composite and Decorator have similar structure diagrams as they both rely on recursive composition to organize an indefinite number of objects.
You can use “Iterator Pattern” to iterate Composite trees.
You can use the Visitor pattern to perform an operation on an entire Composite tree.
Designs that make extensive use of Composite and Decorator can often benefit from the use of prototypes. Applying the pattern allows you to clone complex structures instead of rebuilding them from scratch.
Pros and Cons

You can work with complex tree structures more comfortably: use polymorphism and recursion to your advantage.
Open-closed principle. You can introduce new element types in the application without modifying the existing code, which now works with the object tree.

In some cases, you will have to overgeneralize the component interface, making it more difficult to understand.