# Go Debugging & Performance Lab

A Go backend project demonstrating practical software engineering techniques including concurrent programming, debugging, bug prevention, code organization, automated testing, race detection, benchmarking, performance analysis, and containerized deployment.

## Project Overview

This project implements a lightweight task-management REST API using Go.

The goal is not only to build an API, but also to demonstrate engineering practices commonly required when working with production codebases:

- Feature implementation
- Concurrent programming
- Bug prevention and debugging
- Codebase organization and refactoring
- Unit and concurrency testing
- Race-condition detection
- Performance benchmarking
- Algorithm and data-structure analysis
- Reproducible Docker environments

## Architecture

The application separates HTTP handling, business logic, and domain models.

```text
HTTP Request
     |
     v
TaskHandler
     |
     v
TaskService
     |
     v
Task Model
