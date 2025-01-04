# Bitcoin Wallet Generation Project in Go

This project was developed to better understand the cryptography involved in creating Bitcoin wallets. Using the Go programming language, the project generates public and private key pairs and converts them into Bitcoin addresses.

## Objectives

- Study and understand the cryptographic concepts applied to Bitcoin wallets.
- Implement the generation of private and public keys using Go.
- Convert the public keys into valid Bitcoin addresses.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your system.

## Cryptographic Concepts Used

- ECDSA (Elliptic Curve Digital Signature Algorithm): The algorithm used to generate public and private keys.
- SHA-256 and RIPEMD-160: Hash functions used to convert the public key into a Bitcoin address.
- WIF (Wallet Import Format): A wallet import format used to simplify handling private keys.

## How to Use

1. Clone this repository:
   ```sh
   git clone https://github.com/lpopedv/bit-go-wallet
   cd bit-go-wallet
   ```

2. Run the program:
    ```sh
    go run cmd/main.go
    ```

3. The program will generate and display the private key, WIF (Wallet Import Format) key, public key, and the corresponding Bitcoin address in the terminal.

## Project Structure
- main.go: Contains the main logic for generating keys and Bitcoin addresses.
- pkg/: Contains auxiliary packages for key and address generation.

## Example
```sh
    Private Key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    WIF: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    Public Key: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx 
    Bitcoin Wallet: 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa
```
