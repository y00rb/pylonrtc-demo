# pylonrtc-demo
The demo is built top of [signaler server](https://github.com/y00rb/pylon-signaler). the `./assets/index.js` is one "library" to help build WebRTC features in front-end.

It is a front-end application, but for the learning go web programming,
I built up it by go with http library, and it only launch http server
with html and js file, no more special purpose.

## Usage

Reqiure go 1.12 or above
1. Clone code
```
git clone https://github.com/y00rb/pylonrtc-demo
```
2. Install dependency
```
go get
```
3. Running signaler server in local
```
go run main.go
```
4. Open `http://localhost:8282?roomID=${roomID}`
   in browser's private mode, Once passed `roomID` as query string, the signaler server
will build meeting session by the `roomID`, other *another* page in
browser's private mode with URL
`http://localhost:8282?roomID=${roomID}`, you would running one `peer to
peer` meeting in local.
> If only open `http://localhost:8282` without roomID, signaler server
> will create uniq roomID and show on html page, please copy the
> generated roomID in *another* page in private mode with `http://localhost:8282?roomID=${roomID}`
