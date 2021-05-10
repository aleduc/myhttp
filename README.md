# Myhttp
Myhttp is package for doing http get requests concurrently. It returns response as a MD5 hash

## Build
```bash
go build cmd\myhttp\myhttp.go
```
Or use prebuilt myhttp.exe for Windows or myhttp for Linux.

## Documentation
Module does not return errors, only MD5 hash in any case.

### Input data format
List of urls without a schema. 
````
myhttp adjust.com google.com facebook.com
````

### Output data format 
URL and MD5 hash of response. 
````
http://google.com f54a964ccc757e57d41b051ed8d5c5a3
http://adjust.com 787d1722dc1a1bd33029b8b9e13a0e1e
http://facebook.com 804564070a1b4930381a3367ee5d33ca
````

### Flags
- `-parallel`: Max concurrent requests. Default value is 10.
