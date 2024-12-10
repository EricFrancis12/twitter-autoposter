#!/bin/bash

input_file_path="config.example.json"
output_file_path="config.json"

file_content=$(cat "$input_file_path")

echo $file_content > "$output_file_path"

echo "created new config file at: $output_file_path"
