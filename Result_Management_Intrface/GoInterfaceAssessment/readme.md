# Result Management System for Multiple Student Types

## To implement
 - Model students from different streams e.g. Engineering, Arts, each with its own grading logic.
 - Use interfaces to unify student types under a common grading interface.
 - Use composition — maybe a base student struct with embedded logic for common fields like name.

## Implementation Detail
 - ### Model Section
    - We have defined a interface Student with following methods:
        - GetID
        - GetName
        - GetDepartment
        - CalculateGrade
    - Every struct (engineering, arts) with implement this interface.
    - Engineering department has its own grading System (each subject has some credit and it is calculated accordingly)
    - Art department has normal grading system
 - ### Service Section
    - we have result service with following functionality:
        - AddStudent: it add a new student, if student with same id is present gives error.
        - GetStudentById: Return detail of student with that Id, if no student found gives error.
        - CalculateAllResult: Calculate the result of all student based on their grading system.

## How to Run the Program
- Unzip the folder
- Navigate into the folder 
    ``` bash 
    cd GoInterfaceAssessment
    ```
- Run the binary file: `main.exe`
    ``` bash 
    ./main.exe
    ```
## Sample Output
    
    Student: Vicky, Department: Engineering, Grade: A, GPA: 8.84, Grading System: Credit based
    Student: Shiv, Department: Arts, Grade: A, GPA: 8.37, Grading System: Normal
    