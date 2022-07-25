Singleton Design Pattern in Go
Singleton is a component which is instantiated only once in the application life-cycle.

It is used almost in every application built.

In this blog for Singleton Design Pattern we’ll look into when and how to implement Singleton using Go. We will also take a look into the problem that of singleton pattern consists.


Singleton Design Pattern
When To?
Component makes sense to initialized only once.
Construction call is too expensive.
Want to prevent from creating any additional copies.
Scenario
Let’s take the basic scenario where you want to create a connection between your application and the database. In such case you only want a single instances to be created and use the same instance all over the application(connection pooling).

There are two ways of creating a singleton object into Go.

init()
sync.Once
Both the ways are used based on the requirements.

init()
It is a package level function which is called while the application is booting up(even before main()) and it is thread safe.

Let’s have a database connection using init() function.


Here, if you look at the hierarchy of calling the function then you’ll get to see that the init() function has more priority than that of main() function.

init() function is for fast initialization. So if you want some logic to be executed even before the execution of main() function you can use init() function.

sync.Once
It’s Go’s inbuilt package providing you functionality of creating an instance only once.

Let’s have a database connection using sync.Once.


In the above implementation, GetDBConnection() is getting called from main(). So here is lazy-loading mechanism where the instance being created when required instead of on boot-up.

sync.Once internally maintains a flag that once the code is executed then it should not get called again. It is not only for creating instances but also can be used when the code is to be executed only once in the application life-cycle.