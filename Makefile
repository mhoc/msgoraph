build:
	go build github.com/mhoc/msgoraph
	go build github.com/mhoc/msgoraph/auth
	go build github.com/mhoc/msgoraph/common
	go build github.com/mhoc/msgoraph/internal
	go build github.com/mhoc/msgoraph/users

docs:
	@echo "http://localhost:6060/pkg/github.com/mhoc/msgoraph/"
	godoc -http=:6060
