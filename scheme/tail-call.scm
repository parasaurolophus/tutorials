#lang racket

;; Copyright Kirk Rader 2024

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Tail Call Optimization
;;
;;
;; Together with lexical and first-class continuations, tail call
;; optimiation is a core building block of CPS (Continuation Passing
;; Style).
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Self-calling implementation of the factorial function that is
;; susceptible to stack overflow.
(define (naive-factorial x)
  (if (< x 1)
      1
      (* x (naive-factorial (- x 1)))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Mathematically equivalent implementation of factorial that benefits
;; from tail-call optimization.
;;
;; The stack does not grow so the only limit on x is the heap memory
;; available to represent x! as a bignum at each step.
(define (factorial x)
  (let f ((n x)
          (a 1))
    (if (< n 1)
        a
        (f (- n 1) (* n a)))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Similar implementation of the Fibonacci (a.k.a. Pingala) series.
;;
;; This shows that the approach shown here generalizes simply by
;; adjusting the number of parameters and details of the calculation
;; at each iteration.
(define (fibonacci x)
  (let f ((n x)
          (a 0)
          (b 1))
    (if (< n 1)
        a
        (f (- n 1) b (+ a b)))))
