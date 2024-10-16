# Script to extract unique Book IDs from multiple Google Books URLs in a text file
cls
$filePath = Read-Host "Enter the path to the text file with Google Books URLs"
if (-not (Test-Path $filePath)) {
    Write-Host "File not found. Please check the path."
    exit
}
$urls = Get-Content $filePath
$bookIDs = @()
foreach ($url in $urls) {
    if ($url -match "id=([^&]+)") {
        $bookIDs += $matches[1]
    } else {
        Write-Host "No Book ID found in URL: $url"
    }
}
$uniqueBookIDs = $bookIDs | Sort-Object -Unique
if ($uniqueBookIDs.Count -gt 0) {
    Write-Host "Extracted Unique Book IDs:"
    $uniqueBookIDs | ForEach-Object { Write-Host $_ }
} else {
    Write-Host "No Book IDs found in the file."
}
