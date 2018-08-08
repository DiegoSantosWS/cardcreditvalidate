# cardcreditvalidate
API PARA VALIDAR CPF, CNPJ estamos desenvolvendo a parte para validar cartão de credito

# USING #

## CLONE

```bash
$ git clone git@github.com:DiegoSantosWS/cardcreditvalidate.git
```

## DOCKER

Usando o docker para subir o container de execução

```bash
$ cd cardcreditvalidate/  
cardcreditvalidate$ sudo docker build .
```

* DOCKER COMPOSE

```bash
cardcreditvalidate$ sudo docker-compose up .
```

Validando CPF com retorno da api é json

```bash
$ curl GET http://172.20.0.2:8000/validate-cpf/74321824001
```

Validando CNPJ com retorno da api é json

```bash
$ curl GET http://172.20.0.2:8000/validate-cnpj/99586878000118
```

EM DESENVOLVIMENTO...

```bash
$ curl GET http://localhost:3000/v1/validate-credit-card/numbercard
```

* Retorno com sucesso

```json
 {
    msg: "Credit Card is validate",
    Comp: {
        Short: "visa",
        Long: "Visa"
    }
 }
```

* Retorno invalido (caso não encontre informação da bandeira)

```json
 {
    msg: "Credit Card is validate",
    Comp: {
        Short: "",
        Long: ""
    }
 }
```

* Retorno invalido para numero do cartão

```json
 {
    msg: "the number credit card don't valid",
    err: true,
    Comp: {
        Short: "",
        Long: ""
    }
 }
```