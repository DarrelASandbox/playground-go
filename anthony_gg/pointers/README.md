# About The Project

- [Everything You Need To Know About Pointers In Golang](https://www.youtube.com/watch?v=mqH21m0MsWk)

&nbsp;

# Pointers

## Functional

- In this version, `takeDamage` still takes a `Player` struct as an argument and modifies it.
- However, instead of modifying the original struct, it returns a new `Player` struct that represents the state of the player after the damage has been taken.
- This new `Player` struct is then reassigned to the `player` variable in the `main` function.

## Pointer

## Comparison

**Functional Approach:**

The functional approach emphasizes immutability, and thus, operations typically return a new value rather than modifying an existing one.

_Pros:_

1. _Safety_: It avoids the risk of unintended side effects because data is not mutated.
2. _Testability_: It is often easier to test functional code due to its deterministic nature.
3. _Concurrency_: In multi-threaded environments, you often don't need to worry about race conditions because you're not mutating shared state.

_Cons:_

1. _Performance_: Creating a new copy of large structures every time can lead to performance overhead.
2. _Memory Usage_: The need to frequently create copies can lead to higher memory usage.

**Pointer Approach:**

The pointer approach allows you to directly modify the original data structure.

_Pros:_

1. _Efficiency_: Pointers provide a more efficient way to handle large data structures as you are not creating a copy.
2. _Control_: You have more control over the memory you are working with.

_Cons:_

1. _Mutability_: There's a risk of mutating data unintentionally, which can cause bugs.
2. _Concurrency Issues_: In a multi-threaded environment, if multiple threads are mutating a shared state simultaneously, it can lead to race conditions.

Which is objectively better depends on the specific circumstances:

1. If you're dealing with small data structures or prioritizing the benefits of immutability such as safety, testability, and simplicity, especially in a multi-threaded environment, then a functional approach might be more suitable.
2. If you're working with large data structures where copying would be too expensive, or if you require the efficiency and control that pointers provide, then the pointer approach might be better.

Finally, it's also worth noting that the distinction between functional and pointer approaches can sometimes be a bit blurred in practice. For example, even in functional programming, you may sometimes use techniques that give you the efficiency benefits of pointers (like persistent data structures), and likewise, in pointer-based programming, you may use techniques to minimize mutation and manage side effects.
