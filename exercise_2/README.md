Exercise 2

# Goal
This exercise teaches you how to write a class. You will refactor the solution of exercise 1 into a class.

# Task
Implement a class called `URLSet` in `url_set.go`. The class will have one method `Add`, you are welcome to copy the code you write in exercise 1 into the new class.

After you have implemented the `URLSet` class and all the tests pass, you should refactor the code in `exercise_2/crawler.go` so that the `Add` function in `exercise_2/crawler.go` calls the the `Add` function of `URLSet`.

To run the test:
```
cd exercise_2
dep ensure
go test
```

# Resources
- Pointers (succinct) https://gobyexample.com/pointers
- Pointers (more detailed) https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
- "Class" https://gobyexample.com/methods

# Note
PLEASE DON'T MODIFY TEST CASES (You can uncomment them and that's the only thing you can do).

There is a solution folder, but don't look at it until you have finished the exercise. You will NOT learn much if don't try to solve it yourself.

Feel free to google solutions when you are stuck.
