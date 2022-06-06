## mansubs
Take a list of domains and add to a SQLite database for easy managing

## Install
```
go install -v github.com/hieuha/mansubs@latest
```

## Basic Usage
```
echo -n 'mcafee.com' | assetfinder -subs-only | mansubs -domain mcafee.com -create 
```

```
mansubs -dump -domain mcafee.com
```

```
mansubs -update-tech -domain mcafee.com
```