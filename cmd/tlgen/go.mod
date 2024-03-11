module github.com/xelaj/tl/cmd/tlgen

go 1.21.5

replace github.com/xelaj/tl => ../..

replace github.com/xelaj/tl/schema => ../../schema

require (
	dario.cat/mergo v1.0.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/dave/jennifer v1.7.0
	github.com/gabriel-vasile/mimetype v1.4.3
	github.com/iancoleman/strcase v0.3.0
	github.com/joho/godotenv v1.5.1
	github.com/quenbyako/ext v0.0.0-20240210142841-0ae6adb966e4
	github.com/rs/zerolog v1.31.0
	github.com/spf13/cobra v1.8.0
	github.com/xelaj/tl/schema v0.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/participle/v2 v2.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
