set /p version= < ver
docker run -d -p 38081:38081 fullstack.web:2.%version%
pause
