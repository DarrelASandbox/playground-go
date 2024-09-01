- [META](#meta)
  - [Why unit tests and how to make them work for you](#why-unit-tests-and-how-to-make-them-work-for-you)
    - [Refactoring](#refactoring)
    - [Unit tests](#unit-tests)
    - [(Well designed) units](#well-designed-units)
    - [Small steps](#small-steps)
    - [In a Nutshell](#in-a-nutshell)

# META

## [Why unit tests and how to make them work for you](https://quii.gitbook.io/learn-go-with-tests/meta/why)

- [LondonGophers 12/12/2018: Chris James - How to not build legacy systems that everyone hates](https://www.youtube.com/watch?v=Kwtit8ZEK7U)

> Any software system used in the real-world must change or become less and less useful in the environment

> As a system evolves, its complexity increases unless work is done to reduce it

> However the term "refactoring" is often used when it's not appropriate. If somebody talks about a system being broken for a couple of days while they are refactoring, you can be pretty sure they are not refactoring.

- In order to **safely refactor you need unit tests** because they provide
  - Confidence you can reshape code without worrying about changing behavior
  - Documentation for humans as to how the system should behave
  - Much faster and more reliable feedback than manual testing

### Refactoring

- Gives us signals about our unit tests. If we have to do manual checks, we need more tests. If tests are wrongly failing then our tests are at the wrong abstraction level (or have no value and should be deleted).
- Helps us handle the complexities within and between our units.

### Unit tests

- Give a safety net to refactor.
- Verify and document the behavior of our units.

### (Well designed) units

- Easy to write meaningful unit tests.
- Easy to refactor.

### Small steps

- Write a small test for a small amount of desired behaviour
- Check the test fails with a clear error (red)
- Write the minimal amount of code to make the test pass (green)
- Refactor
- Repeat

### In a Nutshell

- The strength of software is that we can change it. Most software will require change over time in unpredictable ways; but don't try and over-engineer because it's too hard to predict the future.
- Instead we need to make it so we can keep our software malleable. In order to change software we have to refactor it as it evolves or it will turn into a mess
- A good test suite can help you refactor quicker and in a less stressful manner
- Writing good unit tests is a design problem so think about structuring your code so you have meaningful units that you can integrate together like Lego bricks.
- TDD can help and force you to design well factored software iteratively, backed by tests to help future work as it arrives.
