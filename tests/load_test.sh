#!/bin/bash

# Load Test for Rate Limiting
# Usage: sh tests/load_test.sh [PORT] (Default: 3000)

PORT=${1:-3000}
URL="http://localhost:$PORT/api/v1/quizzes"

echo "=============================================="
echo "🚀 Starting Rate Limit Load Test (Parallel)"
echo "Target: $URL"
echo "Sending 50 requests concurrently..."
echo "=============================================="

# Create simplified results file
rm -f results.txt

# Run requests in background
for i in {1..50}; do
  curl -s -o /dev/null -w "%{http_code}\n" "$URL" >> results.txt &
done

# Wait for all background jobs
wait

echo "Analysis:"
# Count status codes
ALLOWED=$(grep -c "200" results.txt)
BLOCKED=$(grep -c "429" results.txt)
ERRORS=$(grep -v "200" results.txt | grep -v "429" | grep -v "^$" -c)

rm -f results.txt

echo "=============================================="
echo "📊 Results:"
echo "✅ Allowed: $ALLOWED"
echo "⛔ Blocked: $BLOCKED"
echo "❌ Errors:  $ERRORS"
echo "=============================================="

if [ "$BLOCKED" -gt 0 ]; then
  echo "✅ Rate Limiting is WORKING (Requests were blocked)."
  exit 0
else
  echo "⚠️  Rate Limiting might NOT be working (No requests blocked)."
  echo "   (Try increasing request count or lowering limit)"
  exit 1
fi
