#!/bin/bash

# Exit immediately if a command fails
set -e

# Define variables
BASE_URL="${TOLGEE_BASE_URL}"
API_KEY="${TOLGEE_API_KEY}"
OUTPUT="data.zip"
DEST_DIR="locals"

# Check if BASE_URL is set
if [ -z "$BASE_URL" ]; then
  echo "❌ Error: TOLGEE_BASE_URL environment variable is not set."
  echo "Please set it first, e.g.:"
  echo "  export TOLGEE_BASE_URL=http://localhost:10280"
  exit 1
fi

# Check if API_KEY is set
if [ -z "$API_KEY" ]; then
  echo "❌ Error: TOLGEE_API_KEY environment variable is not set."
  echo "Please set it first, e.g.:"
  echo "  export TOLGEE_API_KEY=tgpak_abcdef1234567890"
  exit 1
fi

# Create destination directory if not exists
mkdir -p "$DEST_DIR"

# Download the ZIP file
echo "Downloading translation data from $BASE_URL ..."
curl -s "${BASE_URL}/v2/projects/export?ak=${API_KEY}" --output "$OUTPUT"

# Unzip into locals directory
echo "Extracting files to $DEST_DIR ..."
unzip -o "$OUTPUT" -d "$DEST_DIR"

# Remove the ZIP file
rm "$OUTPUT"

echo "✅ Export completed. Files extracted to '$DEST_DIR'."
