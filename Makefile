fmt:
	gofmt -w=true jsonfmt.go
	gofmt -w=true indent/*
	gofmt -w=true decode/*
	gofmt -w=true util/*
