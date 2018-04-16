# gook
external memory file content sort

installation:
```
go install ./...
```

check pattern file (unsorted)
```
gookcheck -in ./assets/bible.txt
```
generate new file:

```
gookgen -max 104857600 -output ./generated.txt -limit=1000 -pattern=./assets/bible.txt
```

check generated file (unsorted):
```
gookcheck -in ./generated.txt
```

sort generated file with chunks 10mb and check status (sorted):
```
gooksort -in ./generated.txt -out ./sorted.txt -chunk 10240
gookcheck -in ./sorted.txt
```