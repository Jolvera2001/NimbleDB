# Architecture overview
The main features of NimbleDB are:
- Embedded Database
- NoSQL
- Schema based

And a lot of NimbleDB is inspired by how LiteDB is made, which uses an already defined serializer (BSON) and builds around that.

## Design Choices
I was originally going to also choose BSON but ended up choose Protobuf instead. This choice would help NimbleDB be usable across different languages and also provide a strong schema system

Another reason I wanted to do an embedded database was also because I was kind of frustrated with how MongoDB doesn't have an embedded mode (i could be wrong), but to me, a nosql db would naturally fit an embedded mode. Also I wanted something to do to keep learning and making my own database was on my bucket list

## The Plan
- 0.1.0 - Core
    - This is the base core engine version. I only plan to get the blocks, records, serializer, and index implemented here
    - I might update this with minor versions regardin how testing is done or implemented stuff to work with it, but it could just roll over to the next version
- 0.2.0 - Query and Transactions
    - Query Engine
    - Transactions
- 0.3.0 - Performance and features
    - Performance related changes (Looking into things like LSM engines)
    - Metadata layer related to schema