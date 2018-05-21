echo "path:$GOPATH"
which go
go --version
rm $GOPATH/bin/server
echo "building ..."
go install ./cmd/...

echo "running ..."
$GOPATH/bin/server
