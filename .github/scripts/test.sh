#!/bin/bash

#
# Copyright (c) 2024 OrigAdmin. All rights reserved.
#

# Record the original working directory
ORIGINAL_DIR=$(pwd)
ROOT_DIR=$1

# Define a function to check for the existence of a go.mod file in a directory and perform corresponding actions
check_go_mod_and_test() {
    local dir="$1"
    local go_mod_name="go.mod"
    local need_test=1 # Assume not need_test by default
    
    # Change to the directory
    cd "$dir" || return 1
    
    # Check if the go.mod file exists in the directory
    if [ -f "$go_mod_name" ]; then
        # If it exists, perform specific operations
        
        #		go get -u ./...
        #		go mod tidy
        
        #		local test_status=$?
        # Only mark as need_test if tests pass
        #		if [ $test_status -eq 0 ]; then
        need_test=0
        #		fi
    fi
    
    # Return to the original working directory
    cd "$ORIGINAL_DIR" || return 1
    # Return the update status
    return $need_test
}

# Define a function to traverse directories and apply the check_go_mod_and_act function
check_all_go_mod() {
    local start_dir="$1"
    
    # If a module_name is specified, process only that directory
    if [ "$start_dir" != "." ]; then
        #		if check_go_mod_and_test "./$module_name"; then
        #			echo "Processing directories in: ./$module_name"
        #			local need_test=$?
        #			if [ $need_test -eq 0 ]; then
        #				git_commit_changes "./$module_name"
        #			fi
        start_dir="./$start_dir"
    else
        start_dir="."
    fi
    
    # Skip the root directory ('.')
    find "$start_dir" -mindepth 1 -type d -not -path "./.*" -print0 | while IFS= read -r -d '' dir; do
        
        cd "$ORIGINAL_DIR" || return 1
        if check_go_mod_and_test "$dir"; then
            local need_test=$?
            if [ $need_test -eq 0 ]; then
                echo "Testing packages in: $dir"
                cd "$dir" || continue
                go fmt ./...
                go mod tidy || return 1
                go test -race ./... || return 1
            fi
        fi
    done
}

check_all_go_mod "$ROOT_DIR"
