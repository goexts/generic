#!/bin/bash

#
# Copyright (c) 2024 OrigAdmin. All rights reserved.
#

# Helper function to get tags matching a pattern
# Args:
#   pattern (string): The pattern to match tags against
# Returns:
#   string: A list of tags matching the pattern
get_matching_tags() {
  local pattern=$1
  git tag -l "$pattern"
}

# Helper function to get the latest tag from a list of tags
# Args:
#   tags (string): A list of tags
# Returns:
#   string: The latest tag
get_latest_tag() {
  local tags=$1
  # Sort the tags and pick the last one as the latest
  sort -V <<<"$tags" | tail -n1
}

# Helper function to get the latest commit hash
# Returns:
#   string: The latest commit hash
get_current_commit_hash() {
  git rev-parse --short HEAD
}

# Helper function to get tags that point to a specific commit hash
# Args:
#   hash (string): The commit hash to look for
# Returns:
#   string: A list of tags that point to the commit hash
get_tags_for_commit() {
  local hash=$1
  git tag --points-at "$hash"
}

# Helper function to add a tag to a specific commit hash
# Args:
#   next_tag (string): The tag to add
# Returns:
#   None
create_new_tag() {
  local next_tag=$1
  git tag -a "$next_tag" -m "Bumped version to $next_tag"
}
