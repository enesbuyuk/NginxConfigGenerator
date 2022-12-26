# Nginx Config Generator with Go
How to create Nginx config file with Go?

## Usage:
Notice: Change variables "domain, domainWWWPath, domainFilename" according to you.
### Example:

```
domainWWWPath   := "/var/www/ex.enesbuyuk.com"
domain 		      := "ex.enesbuyuk.com, www.ex.enesbuyuk.com"
domainFilename  := "ex.enesbuyuk.com"
```

```
go build main.go;./main 
```
