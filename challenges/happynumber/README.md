# Happy Number

A happy number is defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle that does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Write a function that determines if a number is happy.

## Example

**Input:** 19  
**Output:** true  
**Explanation:**  
1² + 9² = 1 + 81 = 82  
8² + 2² = 64 + 4 = 68  
6² + 8² = 36 + 64 = 100  
1² + 0² + 0² = 1 + 0 + 0 = 1  

**Input:** 4  
**Output:** false  
**Explanation:**  
4² = 16  
1² + 6² = 1 + 36 = 37  
3² + 7² = 9 + 49 = 58  
5² + 8² = 25 + 64 = 89  
8² + 9² = 64 + 81 = 145  
1² + 4² + 5² = 1 + 16 + 25 = 42  
4² + 2² = 16 + 4 = 20  
2² + 0² = 4 + 0 = 4  
... (loops back to 4)

## Constraints

- Assume the input is a positive integer.
- The function should return `true` if the number is happy, `false` otherwise.