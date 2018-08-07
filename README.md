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

* DOCKER COMPOSER

```bash
cardcreditvalidate$ sudo docker-composer up .
```

Validando CPF com retorno da api é json

```bash
$ curl GET http://172.20.0.2:8000/validate-cpf/74321824001
```

Validando CNPJ com retorno da api é json

```bash
$ curl GET http://172.20.0.2:8000/validate-cnpj/99586878000118
```

EM DESENVOLVIMENTO

```bash
$ curl GET http://localhost:3000/v1/card-credit-validate/numbercard
```