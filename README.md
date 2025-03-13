# NimbleDB
A schema based nosql embedded database (that's a lot of words) using protobuf as a serializer and for the schema

## Overview
Why do I want to make a database that does all of that? Just because I'm bored and want to learn more about making databases. If this takes off then nice.

## Version 0.1.0 - Core Engine
- [ ] block
- [ ] record
- [ ] serializer (protobuf)
- [ ] index (skip list)
- [ ] minimal api

## Issues and Edgecases
- the nature of protobuf forces us to be wary of schema changes and type changes for a message
  - Metadata layer necessary to get around renaming
  - how should type changes be handled? 
