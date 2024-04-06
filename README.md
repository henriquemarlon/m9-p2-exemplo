# M9-P2

O código contido neste repositório representa o sistema de simulação. Este projeto foi construído conforme as [golang-standards](https://github.com/golang-standards/project-layout) [^1].

## Dependências e Serviços

Antes de continuar, é necessário instalar as dependências e criar os serviços listados para a execução dos comandos posteriores. Para isso siga as seguintes instruções:

- Cluster Kafka - [Confluent Cloud](https://docs.confluent.io/cloud/current/clusters/create-cluster.html#create-ak-clusters)
- Cluster MQTT - [HiveMQ Cloud](https://www.hivemq.com/article/step-by-step-guide-using-hivemq-cloud-starter-iot/)
- Cluster MongoDB - [MongoDB Atlas](https://www.mongodb.com/basics/clusters/mongodb-cluster-setup)
- Docker engine - [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
- Build Essential - [What is Build Essential Package in Ubuntu?](https://itsfoss.com/build-essential-ubuntu/)

## Como rodar o sistema:
Siga as intruções abaixo para rodar o sistema junto a todos os seus recortes, simulação, mensageria, banco de dados.

### Definir as variáveis de ambiente:
Rode o comando abaixo e preecha com as respectivas variáveis de ambiente o arquivo `.env` criado dentro da pasta `/config`.

#### Comando:
```shell
make env
```

#### Output:
```shell
cp ./config/.env.develop.tmpl ./config/.env
```

> [!NOTE]
> Antes de preencher o arquivo `.env` é necessário criar os serviços de cloud presentes nas seção [#Dependências e Serviços]()

### Rodar as migrações:
As migrações, referem-se ao conjunto "queries" criadas com o objetivo de trazer agilidade ao processo de desevolvimento, que criam sensores no banco de dados que por sua vez servirão para contruir a simulação. 

#### Comando:
```shell
make migrations
```

#### Output:
```shell
migrations  | Connection established successfully
migrations  | Documents inserted. IDs: [ObjectID("65f0575382f1be93d94ae2c6") ObjectID("65f0575382f1be93d94ae2c7") ObjectID("65f0575382f1be93d94ae2c8") ObjectID("65f0575382f1be93d94ae2c9") ObjectID("65f0575382f1be93d94ae2ca")]
migrations  | Connection to MongoDB closed.
migrations exited with code 0
```

Abaixo estão as possíveis interações e as instruções de como realizá-las.

#### Rodar testes:

Aqui, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

###### Comando:

```shell
make test
```

> [!NOTE]
> - No meio do processo, é necessário subir a simulação e o consumer para realizar os testes de transmissão de mensagens, integradade e persistência.

#### Rodar a sistema:

Mais uma vez, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui]().

###### Comando:

```bash
make run
```

###### Output:

```shell
[+] Running 2/2
 ✔ Container simulation  Recreated                                                                                     0.3s 
 ✔ Container consumer    Recreated                                                                                     0.3s 
Attaching to consumer, simulation
simulation  | 2024/04/06 02:24:22 Selecting all Sensors from the MongoDB collection sensors
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8251: {"sensor_id":"660fbf475a93f408559f5cc6","unit":"PM2.5","level":15.792103948025987,"timestamp":"2024-04-06T02:24:23.826488282Z"}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8252: {"sensor_id":"660fbf475a93f408559f5cc7","unit":"PM2.5","level":84.25787106446776,"timestamp":"2024-04-06T02:24:23.823722097Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efc8")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8253: {"sensor_id":"660fbf475a93f408559f5cc3","unit":"PM2.5","level":12.793603198400799,"timestamp":"2024-04-06T02:24:23.836820516Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efc9")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8254: {"sensor_id":"660fbf475a93f408559f5cc5","unit":"PM2.5","level":61.76911544227887,"timestamp":"2024-04-06T02:24:23.842860643Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efca")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8255: {"sensor_id":"660fbf475a93f408559f5cc4","unit":"PM2.5","level":88.75562218890555,"timestamp":"2024-04-06T02:24:23.836500214Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efcb")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8256: {"sensor_id":"660fbf475a93f408559f5cc2","unit":"PM2.5","level":45.67716141929036,"timestamp":"2024-04-06T02:24:23.907274194Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efcc")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8257: {"sensor_id":"660fbf475a93f408559f5cc3","unit":"PM2.5","level":49.77511244377811,"timestamp":"2024-04-06T02:24:34.05660772Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efcd")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8258: {"sensor_id":"660fbf475a93f408559f5cc5","unit":"PM2.5","level":34.88255872063968,"timestamp":"2024-04-06T02:24:34.056739029Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efce")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8259: {"sensor_id":"660fbf475a93f408559f5cc6","unit":"PM2.5","level":76.9615192403798,"timestamp":"2024-04-06T02:24:34.056703831Z"}
consumer    | 2024/04/06 02:24:50 Inserting log into the MongoDB collection with id: &{ObjectID("6610b272f4539a64d787efcf")}
consumer    | 2024/04/06 02:24:50 Consuming message on qualidadeAr[4]@8260: {"sensor_id":"660fbf475a93f408559f5cc7","unit":"PM2.5","level":30.484757621189406,"timestamp":"2024-04-06T02:24:34.056794413Z"} 
```

> [!NOTE]
>  - Este comando está subindo todos os serviços presentes no arquivo compose.yml. São eles, o broker local, a simulação e a api-test que está sendo usada, por hora apenas para mostrar o log do que está sendo transmitido pela simulação.

#### Rodar a visualização da cobertura de testes:

Novamente, todos os comandos necessários estão sendo abstraídos por um arquivo Makefile. Se você tiver curiosidade para saber o que o comando abaixo faz, basta conferir [aqui](https://github.com/Inteli-College/2024-T0002-EC09-G04/blob/main/backend/Makefile#L21).

###### Comando:

```bash
make coverage 
```

###### Output:
![output_coverage](https://github.com/henriquemarlon/p1-m9/assets/89201795/4128b513-10bd-4200-8e06-285da5701830)

> [!NOTE]
>  - Este comando está criando, a partir do arquivo `coverage_sheet.md`, uma visualização da cobertura de testes nos principais arquivos Go.

[^1]: A estrutura de pastas escolhida para este projeto está de acordo com as convenções e padrões utilizados pela comunidade de desenvolvedores Golang.
