#!/bin/bash

#
# Copyright (c) 2024 OrigAdmin. All rights reserved.
#

# shellcheck disable=SC1091
source "$(pwd)"/.github/scripts/git/git_tag.sh
# shellcheck disable=SC1091
source "$(pwd)"/.github/scripts/git/git_cmd.sh

# Record the original working directory
ORIGINAL_DIR=$(pwd)

# Function to check for the existence of a go.mod file in a directory
check_for_go_mod() {
    local dir="$1"
    local go_mod_name="go.mod"
    local updated=1 # Assume not updated by default
    
    # Change to the directory
    cd "$dir" || return 1
    
    # Check if the go.mod file exists in the directory
    if [ -f "$go_mod_name" ]; then
        echo "Find dependencies go.mod in directories:"
        echo " ->DIR: $dir"
        updated=0
    fi
    
    # Return to the original working directory for the next iteration
    cd "$ORIGINAL_DIR" || return 1
    
    # Return the update status
    return $updated
}

# Function to check if required functions are defined
function_checks() {
    if ! declare -f get_latest_tag >/dev/null; then
        echo "Error: get_latest_tag function is not defined"
        exit 1
    fi
    
    if ! declare -f get_matching_tags >/dev/null; then
        echo "Error: get_matching_tags function is not defined"
        exit 1
    fi
    
    if ! declare -f get_head_version_tag >/dev/null; then
        echo "Error: get_head_version_tag function is not defined"
        exit 1
    fi
    
    if ! declare -f get_next_module_version >/dev/null; then
        echo "Error: get_next_module_version function is not defined"
        exit 1
    fi
    
    if ! declare -f get_latest_module_tag >/dev/null; then
        echo "Error: get_latest_module_tag function is not defined"
        exit 1
    fi
}

# Function to check the go.mod file and perform actions
handle_go_mod_directory() {
    local dir="$1"
    local module="main"
    local module_name
    module_name="$(echo "$dir" | sed "s/^.\///")" # Drop the beginning './'
    
    if [ "$module_name" != "." ]; then
        module="$module_name"
    fi
    
    if check_for_go_mod "$dir"; then
        local HEAD_TAG
        local LATEST_TAG
        local NEXT_TAG
        echo " ->MODULE_NAME: $module"
        HEAD_TAG=$(get_head_version_tag "$module_name")
        if [ -n "$HEAD_TAG" ]; then
            echo " ->HEAD_TAG: $HEAD_TAG"
            echo ""
            return
        fi
        LATEST_TAG=$(get_latest_module_tag "$module_name")
        echo " ->LATEST_TAG: $LATEST_TAG"
        NEXT_TAG=$(get_next_module_version "$module_name")
        echo " ->NEXT_TAG: $NEXT_TAG"
        
        if [ -z "$HEAD_TAG" ]; then
            echo "Creating new tag: $NEXT_TAG"
            create_new_tag "$NEXT_TAG"
        fi
        
        echo ""
    fi
    
}

# Define a function to traverse directories and apply the handle_go_mod_directory  function
add_go_mod_tag() {
    echo "Checking for go.mod files..."
    function_checks
    local module_name="$1"
    
    echo ""
    echo "COMMIT_HASH: $(get_current_commit_hash "$module_name")"
    echo ""

    local current_branch="main"
    current_branch=$(git branch --show-current)
		echo "Current branch: $current_branch"

		# Save all local changes to the Git stash
		echo "Saving all local changes and change to the main branch..."
    git stash save --include-untracked
    git checkout -B "main" "origin/main"

    # If a module_name is specified, process only that directory
    if [ "$module_name" == "." ]; then
        handle_go_mod_directory "."
        elif [ -n "$module_name" ]; then
        handle_go_mod_directory "$module_name"
    else
        # Skip the root directory ('.')
        find . -mindepth 1 -type d -not -path './.*' -print0 | while IFS= read -r -d '' dir; do
            # Construct the commit message, including the module name
            # the module name must be the directory name without the beginning './'
            # otherwise, the commit message will be incorrect
            handle_go_mod_directory "$dir"
        done
    fi

    # Restore the original branch and pop the stash
    git checkout "$current_branch"
    git stash pop
}

add_go_mod_tag "$1"
