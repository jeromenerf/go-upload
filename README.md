# go-upload

> Simple one time upload and download share service

I sometimes have to exchange large files with people who can't use dedicated
services and protocols, such as Dropbox, syncthing or ssh.

`go-upload` provides :

1. an HTML file upload form
2. a 20 sloc or something daemon to handle upload and streaming to file on disk
   and serving from...
3. the `uploads` `http.Dir`

That's all, there is no authentication, no protection from overwriting existing
files, nothing.

1. Git clone the repo
2. `go run main.go`
3. point a browser to http://host:8080/
4. uploads `foobar.txt`
5. wait for the upload to finish
6. grab the file from `http://host:8080/uploads/foobar.txt` or the `uploads`
   dir on the server
