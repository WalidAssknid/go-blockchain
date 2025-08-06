# Blockchain in Go

A simple blockchain implementation written in Go with a command-line interface.

## Description

This project demonstrates the basic concepts of blockchain technology:
- Block structure with index, timestamp, data, previous hash, and current hash
- SHA256 hashing for block integrity
- Genesis block creation
- Adding new blocks to the chain
- Simple CLI for interaction

## Features

- Create and manage a blockchain
- Add new blocks with custom data
- View the entire blockchain
- SHA256 hashing for block security
- Interactive command-line interface

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to the project directory
3. Run the program:
   ```bash
   go run main.go
   ```

## Usage

The CLI provides three options:
1. **Add new block** - Enter data for a new block
2. **View blockchain** - Display all blocks in the chain
3. **Exit** - Close the program

## Project Structure

- `main.go` - Contains the blockchain implementation and CLI # go-blockchain
