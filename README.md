# cpp-code-validator
Code Validator for cpp code that tests the code for given input and delivers final verdict|| Written in Go for high speed performance
- Currently only tested for Linux

Testing:

1. Navigate to Executables directory [`cd Executables`]
2. Make sure both executables namely : -code_runner and judge_backend    have been given executing priviliges [run `chmod +x {filename}`]
3. Feel free to modify input.txt and judge.txt according to your needs.
4. Open Terminal and run ./judge_backend
4. Now navigate to the directory in which you have your cpp file eg. test.cpp
5. Hit a Post Request to judge with curl or any other tool along with your cpp file
    eg. `curl -X POST -F "code=/test.cpp" http://localhost:8080/judge/backend`
6. You will get a verdict as a response to your POST request

Development:

Directory - Checker is responsible for generating an application to check code.cpp file with input output requirements.
Directory - judge is responsible for generating an application to listen to POST requests and run the file in a temporary directory

1. Navigate to whichever directory you would like to experiment with.
2. Run :  `go mod tidy`
3. To build the executable run go build
4. Make sure app has sufficient priviliges chmod +x {executable_name}
5. Run the executable with `./{exe.._name}` in Terminal. Make sure that judge_backend has checker executable to actually check the provided code.

Further Improvements:
-Dockerise the code.cpp with checker to generate verdict
-Speed improvements
-Actual Deployment
