#!/bin.bash
#!/bin/bash
# Script: looksagain.sh
# Description: Find all .sh files (recursively), remove the .sh extension, and show names in ascending order

find . -type f -name "*.sh" 
    | rev | cut -d'/' -f1 | rev \
    | cut -d'.' -f1 \
    | sort -r