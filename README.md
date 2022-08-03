# lambda-calculus-interpreter
A lambda calculus interpreter written in golang

# USAGE
Run the lambdacalc.go in the root directory of the project to use the REPL. The greek lambda character is not supported and only the case-insensitive string "lambda" 
can be used for lambdas. Lambda expressions only take one argument but multi-argument lambda expressions can be simulated using nested lambda expressions (ie lambda x.lambda y.x)

# EXAMPLE USAGE
>>true=lambda x.lambda y.x <br/>
Final value of expression (λx.(λy.x)) <br/>
>>false = lambda x.lambda y.y <br/>
Final value of expression (λx.(λy.y)) <br/>
>>not = lambda b.b false true <br/>
Final value of expression (λb.((b (λx.(λy.y))) (λx.(λy.x)))) <br/>
>>not true <br/>
Final value of expression (λx.(λy.y)) <br/>
>>not false <br/>
Final value of expression (λx.(λy.x)) <br/>
