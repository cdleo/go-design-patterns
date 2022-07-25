**Mediator Pattern in Go**

It is a behavioral design pattern that reduces the coupling between the components of a program by causing them to communicate indirectly through a special mediator object.

It makes it easy to modify, extend, and reuse individual components because they are no longer dependent on all other classes.

Let us get an example to show how is implemented with Golang.

For this case image an airport with only one runway available for all the fly. As we know only one airplane should take off and landing at time, knowing that to control this synchronization the control tower works as a mediator communicating the runway availability.


Additional Information
- It allows us to pull communications between various components into a single site, making it easier to understand and maintain that part, doing that we are applying the single responsibility principle.
- By implementing this pattern, you could introduce new mediators without having to change the components themselves and cover the open-close principle.
- This pattern helps us to reduce the coupling between various components of a program.