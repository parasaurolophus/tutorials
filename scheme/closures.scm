#lang racket

;; Copyright Kirk Rader 2024

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Lexical Closures in Scheme
;;
;;
;; Together with tail-call optimization and first-class continuations,
;; lexical closures created by the lambda special form are a core
;; building block of CPS (Continuation Passing Style). But they can be
;; used for many purposes more loosely related to functional
;; programming as well as other paradigms. (It is a maxim of Lisp
;; programmers that, "objects are a poor man's closures."
;; Historically, the worlds first commercially signficant
;; object-oriented programming system -- Flavors, on what became the
;; Symbolics Lisp Machine -- was implemented using lexical closures as
;; the underlying representation of objects and methods.)
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Use a pair of lexical closures to implement an "object" with a
;; "field," x, with getter and setter "methods."
(define (make-parent x)
  (let ((get-x (lambda () x))
        (set-x! (lambda (y) (set! x y))))
    (values get-x set-x!)))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Use another set of lexical closures to implement object-oriented
;; "inheritance."
(define (make-child x y)
  (let-values (((get-x set-x!) (make-parent x)))
    (let* ((get-y (lambda () y))
           (set-y! (lambda (z) (set! y z)))
           (multiply (lambda () (* (get-x) (get-y)))))
      (values get-x set-x! get-y set-y! multiply))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Demonstrate that each set of closures returned by make-child (and,
;; therefore, make-parent) represent distinct instances.
(define (test-children)
  (let-values (((get-x-1 set-x-1! get-y-1 set-y-1! multiply-1) (make-child 0 0))
               ((get-x-2 set-x-2! get-y-2 set-y-2! multiply-2) (make-child 0 0)))
    (printf "~a ~a ~a ~a ~a\n" (get-x-1) (get-y-1) (get-x-2) (get-y-2) (multiply-1))
    (set-x-1! 2)
    (set-y-1! 42)
    (printf "~a ~a ~a ~a ~a\n" (get-x-1) (get-y-1) (get-x-2) (get-y-2) (multiply-1))))

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;; Simple functional composition.
(define (compose value . closures)
  (let loop ((v value)
             (c closures))
    (if (pair? c)
        (loop ((first c) value) (rest c))
        v)))
