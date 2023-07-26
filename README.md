# Command Line Todo App (Golang)


## Description

The Command Line Todo App is a simple task management application built using Golang. It allows users to create and manage their todos directly from the command line. Todos can be marked as done or undone, and completed todos can be cleared from the list. The application provides a visual distinction for undone todos in red and done todos in green, making it easy to identify their status.

## Features

- Create new todos with custom todo strings.
- Mark todos as done or undone.
- Delete todos by index.
- Clear completed todos from the list.
- Visual representation of undone todos in red and done todos in green.

## Installation

To use this application, you need to have Golang installed on your system. Follow these steps to get started:

1. Clone this repository to your local machine:

```bash
git clone https://github.com/Kishan-Kulkarni/todo.git
```

2. Navigate to the project directory:

```bash
cd todo
```

3. Build the application using the following command:

```bash
go build
```

4. Once the build process is complete, install the application globally:

```bash
go install
```

Now, you can use the "todo" command anywhere in your terminal to access the Todo App.

## Usage

The Todo App supports the following commands:

1. **Insert Todo:**

   To create a new todo, use the following command:

   ```bash
   todo insert Todo string
   ```

   Replace `Todo string` with the actual task you want to add. The new todo will be added to the end of the list.

2. **Delete Todo:**

   To delete a todo by its index, use the following command:

   ```bash
   todo delete index
   ```

   Replace `index` with the index number of the todo you want to delete. The index is 1-based, so the first todo will have an index of 1, the second todo will have an index of 2, and so on. If no index is specified, the entire list will be deleted.

3. **Complete Todo:**

   To mark a todo as done, use the following command:

   ```bash
   todo done index
   ```

   Replace `index` with the index number of the todo you want to mark as done. The index is 1-based, and if no index is specified, the entire list will be marked as done

4. **Clear Completed Todos:**

   To remove all completed todos from the list, use the following command:

   ```bash
   todo clear
   ```

   This will clear all the todos that are marked as done.

## Example

Here's an example of how you can use the Todo App:

```bash
# Add new todos
todo insert Buy groceries
todo insert Prepare presentation
todo insert Go for a run

# Mark a todo as done
todo done 1

# View the updated list with color-coded todos
todo

# Delete a todo
todo delete 2

# Clear all completed todos
todo clear
```

## Contributions

Contributions to this Todo App are welcome! If you find any bugs, have feature requests, or want to improve the code, feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/Kishan-Kulkarni/todo).


Happy task managing!