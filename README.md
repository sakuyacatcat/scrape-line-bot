# scrape-line-bot

## Requirements

- install
  - docker
  - docker-compose
  - GNU Make
- registration
  - fly.io

## Usage

### build

```:sh
make build
```

### local run

```:sh
make run
```

### login to fly.io

```:sh
make auth
```

### 1st deploy application on fly.io

```:sh
make app_start
```

### deploy application to fly.io

```:sh
make deploy
```

### destroy application

```:sh
make destroy
```

## Architecture

```mermaid
graph TD

  A[main.go]

  subgraph interface
    B[handler]
  end

  subgraph infrastructure
    C[scraper]
  end


  subgraph view
    D[view]
  end

  subgraph application service
    E[usecase]
  end

  subgraph domain
    F[service]
    G[model]
    H[repository]
    F --> G
  end

  A -->|"エントリーポイント"| B
  B -->|"UseCaseへ注入"| C
  B -->|"表示結果の整形"| D
  B -->|"コート探索"| E
  E -->|"フィルター"| F
  E -->|"クエリ"| H
  C -.->|"実装(抽象に依存)"| H
```
