# testcase

Just a testcase for https://github.com/golang/go/issues/24499

Run with: go test -bench .

# results

    goos: linux
    goarch: amd64
    BenchmarkLoadImageOneSTB-8   	      20	  68040344 ns/op
    BenchmarkLoadImageOneGo-8    	       5	 215866115 ns/op
    BenchmarkLoadImageTwoSTB-8   	      10	 106599105 ns/op
    BenchmarkLoadImageTwoGo-8    	       5	 329030637 ns/op
    PASS
    ok  	_/home/burchr/imgtestcase	6.853s
