build:
	vgo build github.com/mhoc/msgoraph
	vgo build github.com/mhoc/msgoraph/client
	vgo build github.com/mhoc/msgoraph/common
	vgo build github.com/mhoc/msgoraph/internal
	vgo build github.com/mhoc/msgoraph/scopes
	vgo build github.com/mhoc/msgoraph/users

docs:
	@echo "http://localhost:6060/pkg/github.com/mhoc/msgoraph/"
	godoc -http=:6060
