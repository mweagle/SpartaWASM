# SpartaWASM
[Sparta](https://github.com/mweagle/Sparta) application that demonstrates provisioning an S3 backed site running a Go app compiled to [WebAssembly](https://github.com/golang/go/wiki/WebAssembly).

The WASM compilation step is handled as part of a standard one-command `provision` command.

```shell
$ NOOP=1 mage provision
INFO[0000] ════════════════════════════════════════════════
INFO[0000] ╔═╗╔═╗╔═╗╦═╗╔╦╗╔═╗   Version : 1.5.0
INFO[0000] ╚═╗╠═╝╠═╣╠╦╝ ║ ╠═╣   SHA     : 8f199e1
INFO[0000] ╚═╝╩  ╩ ╩╩╚═ ╩ ╩ ╩   Go      : go1.11.1
INFO[0000] ════════════════════════════════════════════════
INFO[0000] Service: SpartaWASM-mweagle                   LinkFlags= Option=provision UTC="2018-10-20T14:35:46Z"
INFO[0000] ════════════════════════════════════════════════
INFO[0000] Provisioning service                          BuildID=9af7799eb1cee20d5222d629d19e509b9ae17e71 CodePipelineTrigger= InPlaceUpdates=false NOOP=true Tags=
INFO[0000] Verifying IAM Lambda execution roles
INFO[0000] IAM roles verified                            Count=1
INFO[0000] Skipping S3 preconditions check due to -n/-noop flag  Bucket=weagle Region=us-west-2 VersioningEnabled=false
INFO[0000] Running `go generate`
INFO[0000] Running target: BuildWASM
INFO[0000] exec: go build -o ./resources/sparta.wasm ./wasm/main.go

...

```

## Instructions

1. `git clone https://github.com/mweagle/SpartaWASM`
1. `cd SpartaWASM`
1. `go get -u -v github.com/magefile/mage`
1. `S3_BUCKET=<MY_S3_BUCKET_NAME> mage provision`
1. In the _Stack output_ section of the log, look for the **S3SiteURL** key and open the provided URL in your browser (eg: _http://spartahtml-site09b75dfd6a3e4d7e2167f6eca73957e-zp9okcokn7o.s3-website-us-west-2.amazonaws.com_).

## Result


<div align="center"><img src="https://raw.githubusercontent.com/mweagle/SpartaWASM/master/site/wasm-demo.gif" />
</div>

## Credits

  - [danieljoos](https://github.com/danieljoos/go-wasm-examples/) for the WASM example code
