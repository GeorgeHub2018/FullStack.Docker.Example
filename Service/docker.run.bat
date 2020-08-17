set /p version= < ver
docker run -d -p 38080:38080 fullstack:2.%version%
pause