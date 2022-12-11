# Github Action for Feishu Robot Messages

[![build-test](https://github.com/x-actions/feishu/actions/workflows/workflow.yaml/badge.svg)](https://github.com/x-actions/feishu/actions/workflows/workflow.yaml)
[![GoDoc](https://godoc.org/github.com/x-actions/feishu?status.svg)](https://pkg.go.dev/github.com/x-actions/feishu)
[![Go Report Card](https://goreportcard.com/badge/github.com/x-actions/feishu)](https://goreportcard.com/report/github.com/x-actions/feishu)

Github Action for Sending feishu robot messages.

## Environment Variables

- FEISHU_CUSTOMERBOT_WEBHOOK: feishu access token
- FEISHU_CUSTOMERBOT_SECRET: feishu secret

## How to Use

```
    - name: Sending Feishu Message
      uses: x-actions/feishu@main
      env:
        FEISHU_CUSTOMERBOT_WEBHOOK: ${{ secrets.FEISHU_CUSTOMERBOT_WEBHOOK }}
        FEISHU_CUSTOMERBOT_SECRET: ${{ secrets.FEISHU_CUSTOMERBOT_SECRET }}
        MSGTYPE: interactive
        TITLE: "x-actions feishu message"
        CONTENT: |
          > Build Github Action Done.
          > ^_^
```

## Options

- release

```
git tag v0.1.0
git push origin --tags
```

- download

```
curl -Lfs -o feishu https://github.com/x-actions/feishu/releases/latest/download/feishu-{linux|darwin|windows}
chmod +x feishu
./feishu
```

- help

```
Usage of ./feishu:
```

## Docs

- https://www.feishu.cn/hc/zh-CN/articles/360024984973
- https://open.feishu.cn/document/ukTMukTMukTM/uETO1YjLxkTN24SM5UjN
- https://github.com/larksuite/oapi-sdk-go/blob/v3_main/README.md
