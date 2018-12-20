# totp-tool


## Usage

### passcode

```
./totp-tool passcode -s {secret}
```

### check

```
./totp-tool check -s {secret}
```

### qrcode

```
// issuer default: syhlion
// account default: syhlion[@]gmail dot com

./totp-tool qrcode -issuer {issuer} -account {account}
```
