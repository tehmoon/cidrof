## Cidrof

`cidrof` is a simple command that will list all the address of one of your network interface.

## Example

```
$> cidrof eth0
172.23.0.123/16
```

## Caveats

For now it only displays `IPv4` addresses, I have a flags ready to be added but I don't need it now.

## How to build

Assuming you have the latest version of [go](https://golang.org)

```
$> git clone https://tehmoon/cidrof
$> cd src
$> go build .
$> ./cidrof eth0
```
