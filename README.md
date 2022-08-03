# lambda-calculus-interpreter
A lambda calculus interpreter written in golang

#USAGE
Run the lambdacalc.go in the root directory of the project to use the REPL. The greek lambda character is not supported and only the case-insensitive string "lambda" 
can be used for lambdas. Lambda expressions only take one argument but multi-argument lambda expressions can be simulated using nested lambda expressions (ie lambda x.lambda y.x)

#EXAMPLE USAGE
>>true=lambda x.lambda y.x
Final value of expression (λx.(λy.x))
>>false = lambda x.lambda y.y
Final value of expression (λx.(λy.y))
>>not = lambda b.b false true
Final value of expression (λb.((b (λx.(λy.y))) (λx.(λy.x))))
>>not true
Final value of expression (λx.(λy.y))
>>not false
Final value of expression (λx.(λy.x))
