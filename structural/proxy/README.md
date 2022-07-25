Proxy Pattern in Go

The Proxy pattern is a structural pattern whose purpose is to provide a surrogate or intermediary of an object to control its access. Usually wraps an object to hide some of its characteristics.

Let us have an example:

In this example, we are going to control the access to the system that allows driving a car.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Facade Pattern and Proxy Pattern are similar because in both could buffer a complex entity and initialize it on their own.
The Adapter Pattern provides a different interface to the wrapped object, on the other hand, the Proxy Pattern provides the same interface, and the Decorator Pattern provides an enhanced interface.
Pros and Cons

Allows control of the service object (main struct) without the client knowing.
Allows managing the lifecycle of the service object when the clients don’t care.
The proxy works even if the service object is not ready or available.
Help to implement the open-close principle.

The code can get complicated since you have to introduce lots of new classes.
The service response may be affected because adds a delay in the execution time.