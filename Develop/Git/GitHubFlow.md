

### branch を切って開発

```
$ git checkout -b testb
```

```
$ git branch -r
```


### pull

```
$ git checkout -b testc
```

```
$ git pull origin testb
```

### レビューア

```
$ git checkout master
```

```
$ git merge testc
```

### master を push

```
$ git push origin master
```

