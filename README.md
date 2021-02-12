# Go-database-env connection for postgres using GORM & GOLANG
* We will establish a postgres database connection in golang using environment variables.

## Steps to store your own environment variable before running the project.
- Open your terminal and type `vim .bashrc`
- Select the `-edit mode (e)` now press `i` for insert text.
``` 
export YOUR_DB_USER_NAME=some_value
export YOUR_DB_USER_PASSWORD=some_value
export YOUR_DB_NAME=some_value
export YOUR_DB_SSL_MODE=disable 
```
- You need to save your file and exit, ***Note:*** don't close te terminal directly or your work will be lost.
- `press [Esc button] and type :wq`
- It will get you back to your terminal.
- Source your file by this command 
`source ~/.bashrc` Now you need to close the terminal and open it again after that you will be able to use your environment variables values.

### Now you can run your project successfully 
- `go run main.go` 
