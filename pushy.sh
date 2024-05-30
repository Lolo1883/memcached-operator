#!/bin/bash

# Check if a commit message was provided
if [ -z "$1" ]; then
  echo "Usage: $0 <commit-message>"
  exit 1
fi

# Assign the first argument to a variable
COMMIT_MESSAGE=$1

# Add all changes to the staging area
git add .

# Commit the changes with the provided commit message
git commit -m "$COMMIT_MESSAGE"

# Push the changes to the remote repository (origin main)
git push origin main

# Print a success message
echo "Changes have been successfully pushed to origin main."

