# Between Two Integers

## Problem Statement
You are given three integers _A_, _B_ and _C_. Determine whether _C_ is not less than _A_ and not greater than _B_.

### Constraints
- _−100 ≤ A, B, C ≤ 100_
- _A_, _B_ and _C_ are all integers

### Input
Input is given from Standard Input in the following format:
```
A B C 
```
### Output
If the condition is satisfied, print `Yes`; otherwise, print `No`.

### Sample 1
| Input	 | Output |
|--------|--------|
|1 3 2   |  Yes  |

C=2 is not less than A=1 and not greater than B=3, and thus the output should be `Yes`.

### Sample 2
| Input	 | Output |
|--------|--------|
|6 5 4   |  No  |

C=4 is less than A=1 and and thus the output should be `No`.

### Sample 3
| Input	 | Output |
|--------|--------|
|2 2 2   |  Yes  |