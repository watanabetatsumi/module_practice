module github.com/watanabetatsumi/module_practice/greeting

go 1.22.0

require greeting/hello v0.0.0-00010101000000-000000000000

replace greeting/hello => ./hello
