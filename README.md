# Multiversion-Cosmos-Client

## Description

Multiversion-Cosmos-Client is a versatile gRPC client crafted in Golang to effectively manage various versions of the Cosmos SDK. 
This innovative implementation addresses the limitations posed by Golang's go.mod system, which restricts the usage of multiple versions of the same package within a single module. 
Often, developers are forced to maintain separate branches for different SDK versions. However, with Multiversion-Cosmos-Client, you have a seamless and efficient alternative for handling multiple SDK versions within a single codebase.

## Why Multiversion-Cosmos-Client?

Multiversion-Cosmos-Client offers several compelling advantages:

- **Golang gRPC Client for Cosmos SDK:** It provides a reliable and feature-rich gRPC client tailored specifically for the Cosmos SDK ecosystem.

- **Elegant Version Management:** Overcoming the limitations of go.mod, this implementation allows you to manage multiple Cosmos SDK versions seamlessly, eliminating the need for maintaining separate branches.

## Build

To harness the power of Multiversion-Cosmos-Client, follow these straightforward steps:

1. **Configure SDK Version**: In the `.env` file, set the desired Cosmos SDK version by modifying the `SDK_VERSION` variable. (Currently supporting versions 45 and 46).

2. **Build and Run**:
    - Execute `make build` to utilize version-specific predefined `go.mod` and `go.sum` files.
    - Launch the client with `make run`.

## Important Notes

When managing packages and dependencies, keep the following in mind:

- Running `go get ./...` and `go mod tidy` may attempt to include all files, including those with different tags. To avoid conflicts, consider commenting out or removing code related to other tag files before running these commands. You can uncomment these files once the commands have been executed.

- Be vigilant about updating the predefined `go.mod` and `go.sum` files whenever new dependencies are added or existing ones are updated.

- Different packages may maintain version-specific files using tags, so it's important to stay aware of these nuances when working with various versions and go build tags.
