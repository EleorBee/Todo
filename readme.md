# Todo

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/EleorBee/Todo.git
cd Todo
```

Run the following command to build and run the project:

```bash
go build -o Todo
todo --help # To see the list of available commands

# Adding a new task
todo add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating task
todo update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully (ID: 1)

# Deleting task
todo delete 1
# Output: Task deleted successfully (ID: 1)

# Marking a task as in progress or done
todo mark-in-todo 1
todo mark-in-progress 1
todo mark-done 1

# Listing all tasks
todo list

# Listing tasks by status
todo list done
todo list todo
todo list in-progress
```

