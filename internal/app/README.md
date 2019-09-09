# Description

This is the backend server of the application.

## Design

The backend uses the Clean Architecture model. It allows the abstraction between business rules, presentation I/O and the drivers/frameworks.

This is essential for an application's need for expansion and additions/deletions that will not have an impact on places where it shouldn't.

## Folder structure

- **Controllers** - Use cases and business rules. 
You can find out all of the possible application's behaviors, features and functions by simply reading the file names.

- **Drivers** - Frameworks and drivers that connect to background essential tools for the application
eg. Databases

- **Handlers** - Control the input/output and presentation behavior.
eg. HTTP, RPC, CLI

- **Models** - Business entities. They describe all the possible entities that exist in the application.
eg. Users

- **Repositories** - Functional structures that allow the control of database storage for each model.
