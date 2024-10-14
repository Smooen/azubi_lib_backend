# This script extracts the ID of a book from a Google Books URL
cls
$url = Read-Host "Enter the Google Books URL"
if ($url -match "id=([^&]+)") {
    $bookID = $matches[1]
    Write-Host "Extracted Book ID: $bookID"
} else {
    Write-Host "No Book ID found in the URL."
}
