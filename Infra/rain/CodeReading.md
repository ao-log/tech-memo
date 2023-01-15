

[.github/workflows/release.yml](https://github.com/aws-cloudformation/rain/blob/main/.github/workflows/release.yml)

GitHub Actions ワークフローのファイル。

次のコマンドでビルドしている。

```
GOOS="$os" GOARCH="$arch" go build -ldflags=-w -o "dist/${name}/" ./cmd/rain
```


[main.go](https://github.com/aws-cloudformation/rain/blob/main/cmd/rain/main.go)

```go
func main() {
	cmd.Execute(rain.Cmd)
}
```

オプションの解析に cobra を使用している。

```go
// Execute wraps a command with error trapping that deals with the debug flag
func Execute(cmd *cobra.Command) {
	os.Exit(execute(cmd))
}
```

コマンドオプションは次のディレクトリ下でオプションごとに実装している。

https://github.com/aws-cloudformation/rain/tree/main/internal/cmd


deploy はこちらで実装している。

https://github.com/aws-cloudformation/rain/blob/main/internal/cmd/deploy/deploy.go

```func init()``` でオプション類のチェックをしている。

パラメータの読み取りは getParameters 関数によって行なっている。

https://github.com/aws-cloudformation/rain/blob/1f2a27fb7481489428f945333d4139348583c163/internal/cmd/deploy/util.go#L85

変更セットの作成は CreateChangeSet 関数によって行なっている。

https://github.com/aws-cloudformation/rain/blob/1f2a27fb7481489428f945333d4139348583c163/internal/aws/cfn/cfn.go#L205

認証情報のセットは loadConfig 関数によって行なっている。

https://github.com/aws-cloudformation/rain/blob/1f2a27fb7481489428f945333d4139348583c163/internal/aws/aws.go#L41


