# rp Simple Cmd to Repeat Input 

rp takes input from Stdin and will repeat 1 or more times. That is it.

### Install
```
git clone https://github.com/sysoftheworld/rp.git && cd rp && make install
```

### Use

```
rp -h
  -c uint
        number of times to repeat the input (default 1)
  -h    help
```

Simple case:
```
echo "hello" | rp -c 2
hello
hello
```

Repeat Request:
```
echo "https://vimeo.com" | rp -c 3 | xargs -I {} curl -o /dev/null -s -w "%{http_code}\n" {}
200
200
200
```

