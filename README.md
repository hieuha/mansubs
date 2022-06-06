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
mansubs -dump -domain mcafee.com |httpx -t 100 -silent|~/osmedeus-base/binaries/webanalyze-mod -c 50 -t 10 -a ~/osmedeus-base/data/technologies.json > tech.txt 

mansubs -update-tech -domain mcafee.com -tech-file tech.txt

```

```
mansubs -search -tech-search nginx
```