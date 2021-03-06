# CPS 633: Lab 2

## Introduction

The purpose of this lab is to implement an Online Keystroke-based Access Control Model.
This OKAM consists of two parts. 
The first part is an authentication module that verifies the identity of the online user using biometrics.
The second part is an authorization module that manages access to files using permissions.

## Usage
**Before running the project, you must run the command `go get` in the project root directory.**

### Part 2
After adding the User*.txt files to the working directory, you can run part 2 of the lab by running `go run main.go`.
This program will output the deviations of the samples for each user, and the threshold value for each.
After outputting the statistics, you will be able to login as a user and request permissions for a file.
The users, files and permissions are located in the file `.user_and_permissions.db`.

### Part 3
To run this portion of the lab, you must be in the folder: `$GOHOME/src/github.com/mohoromitch/cps633-lab2`
and run the command `go run access.go`.
This will start the terminal application that allowes you to edit the read permissions for each user.
All user data is saved to the file `.user_and_permissions.db`.