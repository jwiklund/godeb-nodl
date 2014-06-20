# deb package Go

See [In-flight deb packages of Go](http://blog.labix.org/2013/06/15/in-flight-deb-packages-of-go) for original, this contains only the conversion for when no list/get functionality has been implemented.

[Full clone](https://github.com/niemeyer/godeb) but didn't quite work on my machine, due to missing contains() expression in xmlpath library.

Usage
```
wget go1.3.linux-amd64.tar.gz
godeb-nodl go1.3.linux-amd64.tar.gz
sudo dpkg -i go_1.3-godeb1_amd64.deb
```

Bootstrap

```
wget https://github.com/jwiklund/godeb-nodl/releases/download/1.0/godeb-nodl
chmod +x godeb-nodl
```
