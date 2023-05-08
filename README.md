go-log
====

go-log provides a generic logging interface with various out-of-the-box implementations.
As the interface is minimal, Developers can also provide their own implementation with relative ease

This aproach provides:

1) Abstraction

    By using a generic logging interface, you can abstract away the underlying logging implementation details.
    This means that you can switch out the logging implementation without changing the rest of your code.

1) Consistency

    A generic logging interface ensures that the logging calls in your code are consistent.
    This makes it easier to read and understand the code, and makes it easier to maintain in the long term.

1) Testability

    A generic logging interface makes it easier to write tests for your code.
    You can use a mock logging implementation to test your code and even assert that certain logs were emitted.

1) Modularity

    You can provide separate implementations to different application components with pre-baked fields,
    or different configuration of the underlying implementation.


