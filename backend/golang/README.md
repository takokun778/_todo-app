# Golang

# 構成

```bash
.
├── README.md
├── domain
│   ├── task
│   │   ├── creator_id.go
│   │   ├── deadline.go
│   │   ├── detail.go
│   │   ├── id.go
│   │   ├── status.go
│   │   ├── task.go
│   │   ├── task_list.go
│   │   ├── task_repository.go
│   │   ├── task_usecase.go
│   │   ├── time.go
│   │   ├── title.go
│   │   └── user_id.go
│   └── user
│       ├── age.go
│       ├── birthday.go
│       ├── id.go
│       ├── name.go
│       ├── user.go
│       ├── user_repository.go
│       └── user_usecase.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── database
│   │   └── gorm.go
│   ├── logger
│   │   └── logger.go
│   └── proto
│       ├── task
│       └── user
├── main.go
├── task
│   ├── controller
│   │   ├── task_controller.go
│   │   └── task_translator.go
│   ├── gateway
│   │   ├── task_entity.go
│   │   ├── task_gateway.go
│   │   ├── task_gateway_gorm.go
│   │   └── task_gateway_in_memory.go
│   └── usecase
│       ├── task_create_usecase.go
│       ├── task_delete_usecase.go
│       ├── task_done_usecase.go
│       ├── task_find_all_by_user_id_usecase.go
│       ├── task_get_usecase.go
│       ├── task_undone_usecase.go
│       ├── task_update_usecase.go
│       └── task_usecase.go
└── user
    ├── controller
    │   ├── user_controller.go
    │   └── user_translator.go
    ├── gateway
    │   ├── user_entity.go
    │   ├── user_gateway.go
    │   ├── user_gateway_gorm.go
    │   └── user_gateway_in_memory.go
    └── usecase
        ├── user_get_usecase.go
        ├── user_register_usecase.go
        ├── user_terminate_usecase.go
        ├── user_update_usecase.go
        └── user_usecase.go
```

# Rule
- Service(gRPC)-Controller-Usecase の命名は貫通
- ファイル名は 
    - ${DomainModel}_controller.go
    - ${DomainModel}_usecase.go
    - ${DomainModel}_repository.go
    - ${DomainModel}_gateway.go
- 複数形はなるべく避ける

# 導入ライブラリ
logger
- [zap](https://pkg.go.dev/go.uber.org/zap)
- [【Golang】zapを用いたロギング](https://qiita.com/kwashi/items/8a6a948d45841f53ccf4)
- [golangの高速な構造化ログライブラリ「zap」の使い方](https://qiita.com/emonuh/items/28dbee9bf2fe51d28153)

uuid
- [google/uuid](https://github.com/google/uuid)

orm
gorm
- [Golangのための素晴らしいORMライブラリ](https://gorm.io/ja_JP/)
- [GORM バージョン2の変更点](https://qiita.com/Syoitu/items/a6fb856d3033d536fa53)
- [GOのORMを分かりやすくまとめてみた【GORM公式ドキュメントの焼き回し】](https://qiita.com/gold-kou/items/45a95d61d253184b0f33)
- [GolangのGORMで単数形のテーブルを扱う](https://qiita.com/cobot00/items/02fd9adddf851e72587a)
- [【GORM】Go言語でORM触ってみた](https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81)
- [GORMというORMが超絶便利だった件](https://medium.com/@taka.abc.hiko/gorm%E3%81%A8%E3%81%84%E3%81%86orm%E3%81%8C%E8%B6%85%E7%B5%B6%E4%BE%BF%E5%88%A9%E3%81%A0%E3%81%A3%E3%81%9F%E4%BB%B6-8d279489c38f)
- []()
TimezoneはアプリとDBともにUTCで統一  
~~TimezoneをPostgreSQLに設定するとparseでバグる~~

prisma
- [Go と prisma と lit-html と ky で作るモダンな TODO アプリ](https://zenn.dev/mattn/articles/1c4eb193d81a3a)

# 参考
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
- [Trying Clean Architecture on Golang](https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047)