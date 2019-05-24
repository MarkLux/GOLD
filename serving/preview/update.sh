
echo "$1" >> tmp.go

mv tmp.go "$2"

rm tmp.go

kill -9 $3

go build "$4"

./wrapper