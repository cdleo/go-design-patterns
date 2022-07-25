Decorator Pattern in Go

Decorator is a structural design pattern that allows you to dynamically add new behavior to existing objects by placing them inside special wrapper objects. Those wrappers will work as an extension of our code that helps to reuse functionality and doesn’t edit the existing codes.

In the following example, we going to simulate a notification system. Imagine a system that has a notification logic that was written to support SMS notifications but now we want to send the same notification message (build inside of send notification function) but now it should work for slack and emails.

We don’t want to redo all the code logic because is stable, then here is the time to apply this pattern. Let us see the code example.


In the above image, we can see the code related to the SMS notification. We can find the interface implemented by the struct. Now in the image below we are defining the needed code to generate slack notifications, here we have the next:

* Defines a struct to wrap the notification functionality.
* Implements the notification interface methods.


And now the implementation. As we can see there we are taking the instance of “smsNotification” and is passed to the new slack struct. Another important thing here is that we should call the “send” method from the slack notification that is because we need to execute both, the sms and the slack notifications.


Output:


Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

Composite and Decorator have similar structure diagrams as they both rely on recursive composition to organize an indefinite number of objects.
A Decorator is like a Composite but it only has one child component. There is another important difference: Decorator adds additional responsibilities to the wrapped object, while Composite just “recaps” the results of its children. However, patterns can also help: you can use the Decorator to extend the behavior of a specific object in the Composite tree.
Decorator allows you to change the structure of an object, while Strategy allows you to change its innards.
Decorator and Proxy have similar structures, but very different purposes. Both patterns are based on the principle of composition, according to which one object should delegate part of the work to another. The difference is that normally a Proxy manages the lifecycle of its service object itself, while the Decorators composition is always controlled by the client.
Pros and Cons

The behavior of an object/struct could be extended/modified without creating a new struct.
Adds or removes responsibilities from an object at runtime.
Allows to combine multiple behaviors by wrapping an object with multiple decorators.

It is difficult to remove a specific wrapper from the wrapper stack.
It is difficult to implement a decorator in such a way that its behavior does not depend on the order in the decorator stack.
The initial configuration code for the layers can look ugly.
