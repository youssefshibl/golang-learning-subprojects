# Command-Line Quiz Application

A simple, timed quiz application that runs in your terminal. This program reads questions and answers from a CSV file and challenges users to answer within a time limit.

## Overview

This command-line quiz application is built in Go and offers the following features:

- Reads questions and answers from a CSV file
- Times the quiz with a default time limit of 20 seconds
- Tracks and displays the user's score at the end
- Customizable CSV file input

## Installation

```bash
# Clone the repository
# Clone the repository
git clone https://github.com/youssefshibl/golang-learning-subprojects.git

# Navigate to the project directory
cd golang-learning-subprojects

# Navigate to the quiz directory
cd quizGame

# Build the project
go build -o quiz

# Optional: Move the binary to your PATH
mv quiz /usr/local/bin/
```

## Usage

```bash
# Run with default settings (uses problems.csv and 20-second timer)
./quiz

# Specify a different CSV file
./quiz -csv=my-questions.csv
```

## CSV File Format

The application expects a CSV file with the following format:

- Each row contains a question and its answer
- The first column contains the question
- The second column contains the correct answer

Example CSV content:

```
5+5,10
7+3,10
1+1,2
8+3,11
```

## Default Settings

- **Default CSV file**: `problems.csv` (must be in the same directory as the executable)
- **Time limit**: 20 seconds for the entire quiz

## How It Works

1. The program reads questions and answers from the specified CSV file
2. A timer starts counting down from 20 seconds
3. Questions are presented one after another
4. The user inputs their answers
5. When the timer expires or all questions are answered, the final score is displayed

## Features

- **Timed quizzes**: Test your knowledge under pressure with a countdown timer
- **Custom question sets**: Use your own CSV files to create different quizzes
- **Simple scoring**: Get immediate feedback on your performance

## Technical Details

- Written in Go
- Uses Go's concurrency features (goroutines and channels) to handle the timer and quiz simultaneously
- Implements graceful error handling for file operations
