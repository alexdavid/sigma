# sigma
A stable Golang API for MacOS Messages app.

---
I built sigma mostly for myself to be able to send and receive messages from Linux.
It is purely a Go library to abstract any future breakages Apple might make and provide
a simple and consistent API.

## Usage
Sigma by itself is probably not very useful to you unless you want to build your own frontend.
Instead, see one of the pre-built frontends:
* [Sigma-Web](https://github.com/alexdavid/sigma-web)
* Sigma-weechat (I'll get around to building this one day)

To write your own frontend see [GoDoc here](https://godoc.org/github.com/alexdavid/sigma/sigma).
Note: Sigma is still in very early development and the API *may* change, but probably won't.

## Requirements
This library requires an Apple computer running MacOS 10.12 or later.


## Alternatives
* [bboyairwreck/PieMessage](https://github.com/bboyairwreck/PieMessage)
  
## Todo
  - [ ] Look into better display name handling if a contact exists in address book
  - [ ] Test sigma on newer versions of MacOS (only has been tested on 10.12)
  - [ ] Create [weechat](https://weechat.org/) integration
