#!/bin/bash

#
# Copyright (c) 2024 OrigAdmin. All rights reserved.
#

# Record the original working directory
ORIGINAL_DIR=$(pwd)
MODULE_NAME=$1

# Define a function to check for the existence of a go.mod file in a directory and perform corresponding actions
check_go_mod_and_act() {
    local dir="$1"
    local go_mod_name="go.mod"
    local updated=1 # Assume not updated by default
    
    # Change to the directory
    cd "$dir" || return 1
    
    # Check if the go.mod file exists in the directory
    if [ -f "$go_mod_name" ]; then
        # If it exists, perform specific operations
        echo "Updating packages in: $dir"
        go get -u ./...
        go mod tidy
        
        go test -race ./... || true
        local test_status=$?
        # Only mark as updated if tests pass
        if [ $test_status -eq 0 ]; then
            updated=0
        fi
    fi
    
    # Return to the original working directory
    cd "$ORIGINAL_DIR" || return 1
    # Return the update status
    return $updated
}

# Define a function to commit go.mod and go.sum files to Git
git_commit_changes() {
    local dir="$1"
    local module_name=""
    local commit_message=""
    
    # Change to the directory
    cd "$dir" || return
    
    # Construct the commit message, including the module name
    # the module name must be the directory name without the beginning './'
    # otherwise, the commit message will be incorrect
    module_name=$(echo "$dir" | sed "s/^.\///") # Drop the beginning './'
    
    echo "Checking update ${module_name}/(go.mod|go.sum)"
    if git status --porcelain | grep -q -E "^[ ]?(M)[ ]? ${module_name}/(go\.mod|go\.sum)$"; then
        # Add go.mod and go.sum to the Git staging area
        git add go.mod go.sum
        
        commit_message="feat($module_name): Update go.mod and go.sum for ${module_name}"
        echo "Committing changes in [$module_name] with message: $commit_message"
        # Commit the changes
        git commit -m "$commit_message"
    else
        echo "No changes to commit in [$module_name]"
    fi
    
    # Return to the original working directory
    cd "$ORIGINAL_DIR" || return
}

# Define a function to traverse directories and apply the check_go_mod_and_act function
update_go_mod() {
    local module_name="$1"
    
    # If a module_name is specified, process only that directory
    if [ -n "$module_name" ]; then
        if check_go_mod_and_act "./$module_name"; then
            echo "Processing directories in: ./$module_name"
            local updated=$?
            if [ $updated -eq 0 ]; then
                git_commit_changes "./$module_name"
            fi
        fi
    else
    	  # Specify root directory
        start_dir="."
        find "$start_dir" -mindepth 1 -type d \
  				-not -path "$start_dir/.*" \
					-not -path "$start_dir/example*" \
					-not -path "$start_dir/examples*" \
					-not -path "$start_dir/test*" \
					-not -path "$start_dir/tests*" \
        	-print0 | while IFS= read -r -d '' dir; do
            if check_go_mod_and_act "$dir"; then
                echo "Processing directories in: $dir"
                local updated=$?
                if [ $updated -eq 0 ]; then
                    git_commit_changes "$dir"
                fi
            fi
        done
    fi
}

update_go_mod "$MODULE_NAME"
