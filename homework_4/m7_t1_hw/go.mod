module example.com/m7_t1_hw

go 1.21.3

replace example.com/order => ../order
replace example.com/stdin_wrapper => ../stdin_wrapper

require example.com/order v0.0.0-00010101000000-000000000000
require example.com/stdin_wrapper v0.0.0-00010101000000-000000000000