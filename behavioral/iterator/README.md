Iterator Pattern in Go

The Iterator is a behavior pattern that helps us to iterate collections of the same or different types. The main idea is to abstract the iteration in a wrapper, that wrapper will define a generic method to iterate the collection.

Let us see an example:




Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Allows to loop maps that implement Composite Pattern
Allows to use the Factory Method pattern in conjunction with the Iterator Pattern In order to return different types of iterators that are compatible with collections.
Is possible to use Visitor together with Iterator to walk through a complex data structure and perform some operations on its elements, even though they all have different classes.
Pros and Cons

Allows to loop through the same collection in parallel because each iterator object contains its own iteration state.
Helps to apply the “Single responsibility principle”. You can clean up your code by extracting bulky traversal algorithms and putting them into interfaces and implementing them in separate structs.
Helps to work with the Open-closed principle. You can implement new types of collections and iterators and pass existing code without breaking anything.

Applying the pattern can be overkill if your application works only with simple collections.
Using an iterator can be less efficient than directly traversing the elements of some specialized collections.