## CLI tool

### Get a Repo
```bash
>>> git clone -b cli https://github.com/antik9/gootus
```

### Build the tool
```bash
>>> cd gootus
>>> go install gootus -mod vendor
```

### Run gootus

1. Unpack
```bash
>>> gootus -unpack
r4\32\\2ju
rrrr33\\ju
```

2. Top N words
```bash
>>> gootus -top 3
a b c d e a b c d c b a c b c
c
b
a
```
