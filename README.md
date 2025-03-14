# NimbleDB
A schema based nosql embedded database (that's a lot of words) using protobuf as a serializer and for the schema

## Overview
Why do I want to make a database that does all of that? Just because I'm bored and want to learn more about making databases. If this takes off then nice.

## Working on Version 0.1.0 - Core Engine
- [ ] block
  - [ ] interface
  - [ ] implementation
- [ ] record
  - [ ] interface
  - [ ] implementation
- [ ] serializer (protobuf)
  - [ ] interface
  - [ ] implementation
- [ ] index
  - [ ] interface
  - [ ] implementation
- [ ] minimal api
  - [ ] definition
  - [ ] implementation

## Future versions
- 0.2.0 - Query Engine and Transactions
- 0.3.0 - Performance and Features

## Issues and Edgecases
- the nature of protobuf forces us to be wary of schema changes and type changes for a message
  - Metadata layer necessary to get around renaming
  - how should type changes be handled? 
