# Go RnD Summit Workshop

This is a workshop to teach Go to devs

# Initial Setup

`brew install go` to install Golang.

Following the steps on https://github.com/golang/go/wiki/SettingGOPATH to set $GOPATH.

`go get github.com/gotoronto/rnd-workshop` to download the repo.

`cd $GOPATH/src/github.com/gotoronto/rnd-workshop` to change your working directory to the repo.

If everything is set up correctly, you should see `Hello world!` in the console when you run `go run main.go`.

`brew install dep` to install go dep. It's a tool that manages libraries (similar to bundle for rails).

`cd $GOPATH/src/github.com/gotoronto/rnd-workshop && dep ensure` to make sure dep works.

## Slides

`slides.key` is the presentation for this workshop.
