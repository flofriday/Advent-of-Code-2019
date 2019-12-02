import os

print("--- Add a day ---")
print("Enter the number of the day you want to create")
day = input("> ")

# Check if the folder is already there
foldername = "day" + day
if os.path.exists(foldername):
    print("Sorry, that day already exits")
    exit()

# Create the folder
os.makedirs(foldername)

# Create go file
go_content = """
package main

func main() {
    // TODO: Solve puzzle here
}
"""
go_file = open(foldername + "/main.go","w")
go_file.write(go_content.strip())
go_file.close()

# Create input file
go_file = open(foldername + "/input.txt","w")
go_file.close()

# Create .gitignore file
git_content = f"""
# Exclude binaries
day{day}.exe
day{day}
main.exe
main
"""
git_file = open(foldername + "/.gitignore","w")
git_file.write(git_content.strip())
git_file.close()