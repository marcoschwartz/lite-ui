#!/bin/bash

# Lite UI Release Script
# Single source of truth: VERSION file

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Read version from VERSION file (single source of truth)
VERSION=$(cat VERSION | tr -d '\n')

echo -e "${GREEN}🚀 Lite UI Release v${VERSION}${NC}"
echo ""

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}❌ Error: Not a git repository${NC}"
    exit 1
fi

# Get commit message
read -p "Commit message: " COMMIT_MSG

if [[ -z "$COMMIT_MSG" ]]; then
    echo -e "${RED}❌ Commit message cannot be empty${NC}"
    exit 1
fi

# Add all changes
echo -e "${GREEN}📦 Adding files...${NC}"
git add .

# Commit
echo -e "${GREEN}💾 Committing: ${COMMIT_MSG}${NC}"
git commit -m "$COMMIT_MSG"

# Create tag
echo -e "${GREEN}🏷️  Creating tag v${VERSION}...${NC}"
git tag -a "v${VERSION}" -m "Release v${VERSION}" 2>/dev/null || {
    echo -e "${YELLOW}⚠️  Tag v${VERSION} already exists, deleting and recreating...${NC}"
    git tag -d "v${VERSION}"
    git push origin ":refs/tags/v${VERSION}" 2>/dev/null || true
    git tag -a "v${VERSION}" -m "Release v${VERSION}"
}

# Push
echo -e "${GREEN}⬆️  Pushing to origin...${NC}"
git push origin $(git branch --show-current)
git push origin "v${VERSION}"

echo ""
echo -e "${GREEN}✅ Release v${VERSION} complete!${NC}"
echo ""
