![pm](https://socialify.git.ci/burntcarrot/pm/image?description=1&descriptionEditable=PM%3A%20The%20project%20management%20tool%20%5BWIP%5D&font=Inter&logo=https%3A%2F%2Fwww.svgrepo.com%2Fshow%2F367972%2Fwaves-outline.svg&pattern=Solid&theme=Light)

### Table of Contents

| [Tech Stack](#tech-stack) | [Unit Testing](#unit-testing) | [API Testing](#api-testing) | [API Docs](#api-docs) | [Continuous Integration](#ci) |
| :-----------------------: | :---------------------------: | ----------------------------------- | --------------------------------------------- | ------------------------- |

## Tech Stack

PM is built with:

- Go
- Echo
- Redis
- GORM
- Mockery
- Testify
- Insomnia
- Swagger
- GitHub Actions

## Unit Testing

Above 80% code coverage in each of the domains. (business domains)

```
ok      github.com/burntcarrot/pm/entity/project        0.012s  coverage: 83.3% of statements
?       github.com/burntcarrot/pm/entity/project/mocks  [no test files]
ok      github.com/burntcarrot/pm/entity/task   0.010s  coverage: 83.3% of statements
?       github.com/burntcarrot/pm/entity/task/mocks     [no test files]
ok      github.com/burntcarrot/pm/entity/user   0.012s  coverage: 100.0% of statements
?       github.com/burntcarrot/pm/entity/user/mocks     [no test files]
```

```
=== RUN   TestCreateProject
=== RUN   TestCreateProject/Valid_Project_Creation
=== RUN   TestCreateProject/Invalid_Project_Creation
--- PASS: TestCreateProject (0.00s)
    --- PASS: TestCreateProject/Valid_Project_Creation (0.00s)
    --- PASS: TestCreateProject/Invalid_Project_Creation (0.00s)
=== RUN   TestGetProjects
=== RUN   TestGetProjects/Get_Projects
--- PASS: TestGetProjects (0.00s)
    --- PASS: TestGetProjects/Get_Projects (0.00s)
=== RUN   TestGetProjectByID
=== RUN   TestGetProjectByID/Valid_Get_Project_by_ID
=== RUN   TestGetProjectByID/Invalid_Get_Project_by_ID
--- PASS: TestGetProjectByID (0.00s)
    --- PASS: TestGetProjectByID/Valid_Get_Project_by_ID (0.00s)
    --- PASS: TestGetProjectByID/Invalid_Get_Project_by_ID (0.00s)
PASS
ok      github.com/burntcarrot/pm/entity/project        (cached)
?       github.com/burntcarrot/pm/entity/project/mocks  [no test files]
=== RUN   TestCreateTask
=== RUN   TestCreateTask/Valid_Task_Creation
=== RUN   TestCreateTask/Invalid_Task_Creation
--- PASS: TestCreateTask (0.00s)
    --- PASS: TestCreateTask/Valid_Task_Creation (0.00s)
    --- PASS: TestCreateTask/Invalid_Task_Creation (0.00s)
=== RUN   TestGetTasks
=== RUN   TestGetTasks/Get_Tasks
--- PASS: TestGetTasks (0.00s)
    --- PASS: TestGetTasks/Get_Tasks (0.00s)
=== RUN   TestGetTaskByName
=== RUN   TestGetTaskByName/Valid_Get_Task_by_Name
=== RUN   TestGetTaskByName/Invalid_Get_Task_by_Name
--- PASS: TestGetTaskByName (0.00s)
    --- PASS: TestGetTaskByName/Valid_Get_Task_by_Name (0.00s)
    --- PASS: TestGetTaskByName/Invalid_Get_Task_by_Name (0.00s)
PASS
ok      github.com/burntcarrot/pm/entity/task   (cached)
?       github.com/burntcarrot/pm/entity/task/mocks     [no test files]
=== RUN   TestLogin
=== RUN   TestLogin/Valid_Login
=== RUN   TestLogin/Invalid_Login_(Empty_Email)
=== RUN   TestLogin/Invalid_Login_(Empty_Password)
--- PASS: TestLogin (0.00s)
    --- PASS: TestLogin/Valid_Login (0.00s)
    --- PASS: TestLogin/Invalid_Login_(Empty_Email) (0.00s)
    --- PASS: TestLogin/Invalid_Login_(Empty_Password) (0.00s)
=== RUN   TestRegister
=== RUN   TestRegister/Valid_Register
=== RUN   TestRegister/Invalid_Register
--- PASS: TestRegister (0.00s)
    --- PASS: TestRegister/Valid_Register (0.00s)
    --- PASS: TestRegister/Invalid_Register (0.00s)
=== RUN   TestGetByName
=== RUN   TestGetByName/Get_user_by_Name
--- PASS: TestGetByName (0.00s)
    --- PASS: TestGetByName/Get_user_by_Name (0.00s)
PASS
```

## API Testing

Coming soon.

## API Docs

Powered by Insomnia and GitHub Pages: https://burntcarrot.github.io/pm

## CI

Coming soon.
