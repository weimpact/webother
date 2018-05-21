rm $GOPATH/bin/server
echo "building ..."
go install ./cmd/...

echo "running ..."
$GOPATH/bin/server
