# Mutex and Channel basics

### What is an atomic operation?
> An operation that is guaranteed to not have race conditions, meaning it cant be interrupted

### What is a semaphore?
> A atmoic variable used for synchronization, typically by incrementing and decrementing (V and P)

### What is a mutex?
> Mutual Exclusion, meaning that only one thread can lock the mutex at a time, and the same thread must also unlock it. Used for protecting resources

### What is the difference between a mutex and a binary semaphore?
> Mutex protects while binary semaphore signals

### What is a critical section?
> The access to the shared resource protected by a mutex

### What is the difference between race conditions and data races?
 > Race condition is caused by incorrect order of events. This needs synchronization/redesign to ensure correctness. A data race is when two or more threads attept to access the same memory area at the same time, at least one of them writing. 

### List some advantages of using message passing over lock-based synchronization primitives.
> Simpler to debug, only one thread will access data structures.
> Scalability
> Threads will inherently block

### List some advantages of using lock-based synchronization primitives over message passing.
> Performance?
