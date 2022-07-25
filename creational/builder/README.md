Builder Pattern in Go

We can describe Builder as a creational design pattern that allows us to build complex objects step by step setting only the properties that we need.

Now, I will show you how we could implement it with Golang. In the following example, we going to build two robot objects, for each one we will set the needed properties using the builder pattern.


In the image above we define a struct type for the robot that we want to create. It contains a type, autonomy, and five “optional” properties that will help us to manage different object behaviors.

Apart to define the robot struct we need to define the “Builder Pattern”, in Golang we can do it using a struct that has a property of the type of the object that we going to manage, in this case, a robot.

With the new “newRobotBuilder” function we going to create the builder and use each one of the functions to set the needed properties.


The following image shows the code needed to set by step each one of the robot properties using the builder (lines 30 to 58). And in line 60 we can read a “build()” function that function will construct the object with the specific values.


We already saw how to define the builder and in the image below we can see how is executed.

First, let us take a look at lines 65 to 76. We create two robots one fighter and one singer, how do we do it? calling to newRobotBuilder and then we start to set the needed properties for each one.

Lines ahead we could find the call to build() function in order to create the robot.


Output:


Note: In this example we just print the properties values, but let us imagine a more complicated implementation, imaging that you need to build a SQL query with multiple optional fields and conditions, using the “Builder Pattern” we can create a SQL Builder and append the needed columns or conditions, something like:

“ sqlQuery := newQueryBuilder().withName().withCount().withLocation().WhereName(“someValue”).WhereLocation(“someValue”).build()”

Additional Information
After learning how this pattern could be implemented with Golang and before ending this article let me give you some bonus information, I will mention a couple of pros and cons and how this pattern could be related to other patterns.

Relationship with other patterns

You can use Builder when creating complex Composite trees because you can program its construction steps to work recursively.
You can combine Builder with Bridge: the director “struct” plays the role of the abstraction, while different constructors act as implementations.
Builder can be implemented as part of a Singleton.
The builder could be combined with the Factory pattern.
Pros and Cons

You can build objects step by step, defer construction steps, or execute steps recursively.
You can reuse the same build code when building multiple product renderings.
Helps to apply the Single Responsibility Principle. You can isolate complex build code from the business logic of the product.

The overall complexity of the code increases as the pattern requires the creation of several new classes.
