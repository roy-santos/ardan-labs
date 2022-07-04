# Ultimate Go Advanced - Advanced Concepts

**Design Philosophy**

*"Good engineering is less about finding the "perfect" solution and more about understanding the tradeoffs and being able to explain them."* - JBD

**4 major categories for basis of code review, prioritized in this order**:

1. **Integrity** - reliablility, all reads/writes/data transformations are accurate, consistent, and efficient. (costs some performance)

*"Failure is expected, failure is not an odd case. Design systems that help you identify failure. Design systems that can recover from failure."* - JBD

2. **Readability** - comprehensible codebase, easy to understand mental model, doesnt hide the "cost" with lots of abstractions

*"Making things easy to do is a false economy. Focus on making things easy to understand and the rest will follow."* - Peter Bourgon

3. **Simplicity** - hide complexity w/o losing readibility, achieved through refactoring, hard to design and complicated to build. (costs some readibility)
4. **Performance** - less computations to reach needed results

Things to consider for performance:
<ul>
    <li> external latency, performance daeth (network calls, system calls, milliseconds of latancy)
    <li> internal latency, microseconds of latency (garbage collector, synchronization, orchestration)
    <li> data access on the machine, access and storage
    <li> algorithm efficiencies, tight loops
</ul>

Question: Can we write Go code that is **fast enough**?
   <ul>
      <li> 99% of the time, answer is YES!
   </ul>

