## slice-linkedlist_bench

### On Russian
Тема популярного задания на собеседовании: `Дан односвязный список. Нужно удалить N-й элемент с конца.`   
Здесь результаты `benchmark tests` двух модулей, которые выполняют это задание разными способами в цикле за 1000 итераций.  
Производительность при использовании среза (:+1:*benchmark test with slice*):

	> go test -bench=.

	goos: windows
	goarch: amd64
	pkg: slice/slice
	cpu: Intel(R) Pentium(R) CPU 5405U @ 2.30GHz
	BenchmarkRemove-4       [1 2 3 5]
	1000000000               0.04790 ns/op         0 B/op          0 allocs/op
	PASS
	ok      slice/slice     0.445s


При использовании односвязного списка (:-1:*benchmark test with linked list*):

	goos: windows
	goarch: amd64
	pkg: linkedlist/linkedlist
	cpu: Intel(R) Pentium(R) CPU 5405U @ 2.30GHz
	BenchmarkMain-4         
	1
	2
	3
	5
	1000000000               0.2738 ns/op          0 B/op          0 allocs/op
	PASS
	ok      linkedlist/linkedlist   2.829s

### On English

Topic of a popular question on job interview: `Given a singly linked list. Need to remove the Nth element from the end.`  
Here are the results of the `benchmark tests` of two modules that do this task in different ways in a loop of 1000 iterations.
