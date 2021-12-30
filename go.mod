module learn.go

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.2.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

replace learn.go.tools => ../learn.go.tools

// 不管别的库用cobra用的哪个版本，强制设置为v1.2.0
replace github.com/spf13/cobra => github.com/spf13/cobra v1.2.0
