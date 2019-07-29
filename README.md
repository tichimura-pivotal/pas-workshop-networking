

# pas-workshop-networking - route service


## 1. ルートサービスの作成

1. clone this repository
2. ```route-service-header``` ディレクトリに移動
```
cd route-service-header
```
3. cf push
4. create route service
```
cf create-user-provided-service SERVICE_NAME_TO_BE_USED -r URL_TO_APP0
```

  * SERVICE_NAME_TO_BE_USED: 任意のサービス名
  * URL_TO_APP0: ルートサービス用にプッシュしたアプリのURL

  ```
  $ cf service route-header
  Showing info of service route-header in org pivot-tichimura / space myground as tichimura@pivotal.io...

  name:                route-header
  service:             user-provided
  tags:                
  route service url:   https://route-service-header.apps.pcfone.io

  There are no bound apps for this service.
  ```

### 2. アプリ(1)のデプロイ

1. ```route-service-oldapp``` ディレクトリに移動
2. cf push

### 3. アプリ(2)のデプロイ

1. ```route-service-newapp``` ディレクトリに移動
2. cf push

### 4. アプリ(1)へのルートサービスのバインド

1. cf bind-route-serviceの実行
```
cf bind-route-service apps.pcfone.io SERVICE_NAME_TO_BE_USED --hostname APP_1_NAME
```
  * SERVICE_NAME_TO_BE_USED: ルートサービスとして作成したサービス名
  * APP_1_NAME: アプリ(1)の名前（URLではない)

2. cf routesで確認 ( アプリ(1)に2つのルートがあることを確認)
```
myground   route-service-olddemo   apps.pcfone.io                             route-service-olddemo   route-header
```

### 5. テスト

1. アプリ(1)へのアクセスを確認
```
$ curl -k https://route-service-olddemo.apps.pcfone.io
this is old path....
```

2. アプリ(1)にHeader情報をつけて確認
```
$ curl -k https://route-service-olddemo.apps.pcfone.io -H "X-Cf-Canary-Url: https://route-service-newdemo.apps.pcfone.io"
this is new path....
```

###  その他

ご要望はtichimura@pivotal.ioまで
