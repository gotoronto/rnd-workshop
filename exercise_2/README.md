Exercise 2

# Goal
This exercise teaches you how to write a class. You will refactor the solution of exercise 1 into a class.

# Task
Create a new file `exercise_2/crawler/list/url_list.go`. You should implement a class called `URLList` in the file that you just created. The class will have one method `Add`, you are welcome to copy the code you write in exercise 1 into the new class.

After you have implemented the `URLList` class and all the tests pass, you should refactor the code in `exercise_2/crawler.go` so that the `Add` function in `exercise_2/crawler.go` calls the the `Add` function of `URLList`.

To run the test:
```
cd exercise_2/crawler/list
dep ensure
go test
<once all the tests pass>
cd ..   // to go to exercise_2/crawler
<refactor code>
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
