Essa é uma aplicação Golang

> Entre no diretório antes de configurar
> ```
> cd urbansoul-API
> ```

> Crie o binario
> ```
> go build .
> ```

> Rode o binario gerado
> ```
> ./urbanAPI
> ```

Todos os endpoints rodam em Localhost:3000/api/endpoint

> Para testes dentro do cmd (exemplo cliente)
> ```
> curl -X GET http://localhost:3000/api/customers      -H "Content-Type: application/json"      -d '{"customer_id": 1}'
> ```
