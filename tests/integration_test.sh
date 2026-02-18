#!/bin/bash

# Integration Test (Happy Path)
# Usage: sh tests/integration_test.sh [PORT] (Default: 3000)

PORT=${1:-3000}
API_URL="http://localhost:$PORT/api/v1"
QUIZ_ENDPOINT="$API_URL/quizzes"

echo "=============================================="
echo "🧪 Starting Integration Test"
echo "Target: $API_URL"
echo "=============================================="

# 1. Create a Quiz
echo "1️⃣  Creating a new quiz..."
response=$(curl -s -X POST "$QUIZ_ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{
    "question": "Integration Test Question",
    "choice1": "A",
    "choice2": "B",
    "choice3": "C",
    "choice4": "D"
  }')

echo "Response: $response"

# Extract ID (simple grep since we might not have jq, or assume response format)
# Assuming response json has "id": "..." 
# Ideally better with jq but try to be dependency-lite
# Let's verify status code instead for robustness in shell
http_code=$(curl -s -o /dev/null -w "%{http_code}" -X POST "$QUIZ_ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{
    "question": "Integration Test Q2",
    "choice1": "1",
    "choice2": "2",
    "choice3": "3",
    "choice4": "4"
  }')

if [ "$http_code" == "201" ] || [ "$http_code" == "200" ]; then
   echo "✅ Create Quiz Successful (HTTP $http_code)"
else
   echo "❌ Create Quiz Failed (HTTP $http_code)"
   exit 1
fi

# 2. List Quizzes
echo ""
echo "2️⃣  Listing quizzes..."
list_response=$(curl -s "$QUIZ_ENDPOINT")
# counts items (approx check)
item_count=$(echo "$list_response" | grep -o "id" | wc -l)
echo "Found approx $item_count items"

if [ "$item_count" -gt 0 ]; then
  echo "✅ List Quizzes Successful"
else
  echo "❌ List Quizzes Failed or Empty"
  exit 1
fi

echo ""
echo "=============================================="
echo "🎉 Integration Test Passed!"
echo "=============================================="
