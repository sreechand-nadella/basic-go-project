# Basic Go Project

This is a basic Go project that demonstrates the following concepts:
- **Basic Control Flow:** Using loops and if statements.
- **Arrays:** Declaring and printing fixed-size arrays.
- **Structures (Structs):** Defining and using a struct.
- **Slices:** Creating and dynamically extending slices.

## Project Overview

The project consists of a simple Go program located in `main.go` along with the necessary Go module file (`go.mod`) and a `.gitignore` file to ignore build artifacts. The project was then initialized as a Git repository and pushed to GitHub.

### Project Structure

basic-go-project/ ├── main.go # Contains the main Go code ├── go.mod # Go module file ├── .gitignore # Git ignore file for build artifacts └── README.md # This file


## How to Run the Project

To run the project locally, execute the following command in your project root:

```bash
go run main.go
# 1. Create the project directory and navigate into it:
mkdir basic-go-project
cd basic-go-project

# 2. Initialize a Go module:
go mod init basic-go-project
go mod tidy

# 3. Create the main.go file and add your code (open in your favorite editor):
# (Use your editor to create main.go and paste in your Go code.)

# 4. Create a .gitignore file:
touch .gitignore

# 5. Add content to .gitignore (you can do this manually or use echo commands):
echo "*.exe" >> .gitignore
echo "*.out" >> .gitignore
echo "*.o" >> .gitignore
echo "*.a" >> .gitignore
echo "*.so" >> .gitignore
echo "/go/" >> .gitignore
echo ".DS_Store" >> .gitignore

# 6. Initialize a Git repository:
git init

# 7. Stage all files:
git add .

# 8. Commit your changes:
git commit -m "Initial commit - basic go project demonstrating basic control flow, struct, slice, and array"

# 9. Add the remote repository (replace with your GitHub URL):
git remote add origin https://github.com/sreechand-nadella/basic-go-project.git

# 10. Rename your branch to main (if not already done):
git branch -M main

# 11. Push your code to GitHub:
git push -u origin main
