# Preface

> Test-driven development is a technique for designing and structuring code
> that encourages simplicity and increases one's confidence in code, even as
> its size increases (xv).

## A Technique

As a technique, TDD is a skill.  That is, it's a practice that people have to
learn.  In addition, the skill is based on a set of belief. Siddiqui lists
three beliefs.

+ Simplicity is essential. (Siddiqui describes simplicity as "the art of
  maximizing the work *not* done" (xv).
+ Obviousness and clarity are better than cleverness.
+ Writing uncluttered code is a large part of success.

The technique itself is a three-step process.

+ Write a test and see it fail.
+ Write just enough code to make the test pass.
+ Refactor the code you've written to improve it—while making sure that it
  still passes the initial test.

## Designing and Structuring Code

According to Siddiqui, "TDD is not fundamentally about testing code" (xvi).
Instead, "the purpose of TDD is to improve the design and structure" of code
(xvi).  Siddiqui argues that if tests were the center of TDD, then we could
write them after rather than before we write the code itself.  Instead, we
write the tests first in order to help us write better designed and structured
software.  (Do I buy this? Maybe, but let's see.)

## Simplicity

Siddiqui lays out several concrete, measurable indicators of simplicity.

+ Fewer lines of code per feature
+ Lower cyclomatic complexity
+ Fewer side effects
+ Smaller runtime or memory requirements

TDD encourages all of these indicators because it pushes us to do "the
simplest thing that works"—where *works* means *passes the tests* (xvi).  TDD
helps to guard against YAGNI violations because you are not allowed to write
code in anticipation of a problem.  First, you must write a test that fails.

## Confidence

Human beings, according to Siddiqui, naturally seek and enjoy predictability.
TDD soothes and pleases us because we trust well-tested code.  We can see all
the ways that we have checked for failure, and as a result we have greater
confidence in what we've written.
