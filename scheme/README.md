_Copyright &copy; Kirk Rader 2024_

# Examples and Tutorials for the Scheme Programming Language

- [./closures.scm](./closures.scm)

   ```scheme
  > (test-children)
  0 0 0 0 0
  2 42 0 0 84
  ```

- [./tail-call.scm](./tail-call.scm)

  ```scheme
  > (map naive-factorial '(0 1 2 3 4 5 6 7 8 9))
  '(1 1 2 6 24 120 720 5040 40320 362880)
  > (map factorial '(0 1 2 3 4 5 6 7 8 9))
  '(1 1 2 6 24 120 720 5040 40320 362880)
  > (map fibonacci '(0 1 2 3 4 5 6 7 8 9))
  '(0 1 1 2 3 5 8 13 21 34)
  ```

- [./continuations.scm](./continuations.scm)

  ```scheme
  > (return-early 42)
  invoking return with 42
  42
  > (return-early-with-protection 42)
  entering protected extent
    calling return with 42 as parameter
  exiting protected extent
  42
  > (test-resume 42)
  calling allow-resume with 42
    entering protected extent
        returning an inner continuation
    exiting protected extent
  received continuation, calling it with 43
    entering protected extent
        resumed continuation received 43, returning 85
    exiting protected extent
  received 85, returning it as final result
  85
  ```
