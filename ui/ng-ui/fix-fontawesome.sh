#!/bin/bash
# Script to replace FontAwesome Pro classes with free alternatives
echo "Replacing FontAwesome Pro classes with free alternatives..."

# Replace "fad" (duotone) with "fas" (solid)
find ./src -type f -name "*.html" -exec sed -i '' 's/class="fad /class="fas /g' {} \;
find ./src -type f -name "*.html" -exec sed -i '' 's/fa-spinner-third/fa-spinner/g' {} \;
find ./src -type f -name "*.html" -exec sed -i '' 's/fa-rocket-launch/fa-rocket/g' {} \;
find ./src -type f -name "*.html" -exec sed -i '' 's/fa-user-visor/fa-user/g' {} \;

echo "Done!"
