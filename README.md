# Projeto de Geração de Carteiras Bitcoin em Go

Este projeto foi desenvolvido com o objetivo de compreender melhor o funcionamento da criptografia envolvida na criação de carteiras de Bitcoin. Utilizando a linguagem Go, o projeto gera pares de chaves públicas e privadas, e os converte em endereços de Bitcoin.

## Objetivos

- Estudar e entender os conceitos de criptografia aplicados às carteiras de Bitcoin.
- Implementar a geração de chaves privadas e públicas utilizando a linguagem Go.
- Converter as chaves públicas em endereços Bitcoin válidos.

## Pré-requisitos

- [Go](https://golang.org/dl/) instalado no sistema.

## Conceitos de criptografia utilizados

- ECDSA (Elliptic Curve Digital Signature Algorithm): Algoritmo utilizado para gerar as chaves pública e privada.
- SHA-256 e RIPEMD-160: Funções hash usadas para converter a chave pública em um endereço Bitcoin.
- WIF (Wallet Import Format): Formato de importação de carteira utilizado para simplificar o manuseio de chaves privadas.

## Como usar

1. Clone este repositório:
   ```sh
   git clone https://github.com/lpopedv/bit-go-wallet
   cd bit-go-waller
   ```

2. Execute o programa:
    ```sh
    go run cmd/main.go
    ```

3. O programa irá gerar e exibir a chave privada, a chave WIF (Wallet Import Format), a chave pública e o endereço Bitcoin correspondente no terminal.

## Estrutura do projeto
- main.go: Contém a lógica principal para geração de chaves e endereços de Bitcoin.
- pkg/: Contém pacotes auxiliares para geração de chaves e endereços.

## Exemplo
```sh
    Private Key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    WIF: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    Public Key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    Bitcoin Wallet: 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
```

