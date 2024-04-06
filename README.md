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
subscriber-1    | ST-0 - Temperature: 4 ºC [OK Refrigerator] 
subscriber-1    | ST-0 - Temperature: 3 ºC [OK Refrigerator] 
subscriber-1    | ST-1 - Temperature: -19 ºC [OK Freezer] 
subscriber-1    | ST-2 - Temperature: -20 ºC [OK Freezer] 
subscriber-1    | ST-4 - Temperature: -14 ºC [ALERT High Temperature - Freezer] 
subscriber-1    | ST-3 - Temperature: -20 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: 9 ºC [OK Refrigerator] 
subscriber-1    | ST-0 - Temperature: 3 ºC [OK Refrigerator] 
subscriber-1    | ST-0 - Temperature: 8 ºC [OK Refrigerator] 
subscriber-1    | ST-2 - Temperature: -18 ºC [OK Freezer] 
subscriber-1    | ST-3 - Temperature: -23 ºC [OK Freezer] 
subscriber-1    | ST-4 - Temperature: -21 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: -16 ºC [OK Freezer] 
subscriber-1    | ST-1 - Temperature: -26 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-0 - Temperature: 7 ºC [OK Refrigerator] 
subscriber-1    | ST-1 - Temperature: -19 ºC [OK Freezer] 
subscriber-1    | ST-3 - Temperature: -16 ºC [OK Freezer] 
subscriber-1    | ST-2 - Temperature: -21 ºC [OK Freezer] 
subscriber-1    | ST-4 - Temperature: -24 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: 9 ºC [OK Refrigerator] 
subscriber-1    | ST-0 - Temperature: 3 ºC [OK Refrigerator] 
subscriber-1    | ST-4 - Temperature: -10 ºC [ALERT High Temperature - Freezer] 
subscriber-1    | ST-1 - Temperature: -10 ºC [ALERT High Temperature - Freezer] 
subscriber-1    | ST-2 - Temperature: -28 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-3 - Temperature: -22 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: -22 ºC [OK Freezer] 
subscriber-1    | ST-3 - Temperature: -24 ºC [OK Freezer] 
subscriber-1    | ST-2 - Temperature: -29 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-1 - Temperature: -26 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-0 - Temperature: -27 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-0 - Temperature: 0 ºC [ALERT Low Temperature - Refrigerator] 
subscriber-1    | ST-0 - Temperature: 12 ºC [ALERT High Temperature - Refrigerator] 
subscriber-1    | ST-0 - Temperature: 11 ºC [ALERT High Temperature - Refrigerator] 
subscriber-1    | ST-1 - Temperature: -12 ºC [ALERT High Temperature - Freezer] 
subscriber-1    | ST-3 - Temperature: -19 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: -19 ºC [OK Freezer] 
subscriber-1    | ST-2 - Temperature: -29 ºC [ALERT Low Temperature - Freezer] 
subscriber-1    | ST-4 - Temperature: -18 ºC [OK Freezer] 
subscriber-1    | ST-0 - Temperature: 9 ºC [OK Refrigerator] 
subscriber-1    | ST-0 - Temperature: 7 ºC [OK Refrigerator] 
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
