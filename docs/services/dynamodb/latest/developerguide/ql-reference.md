# PartiQL - a SQL-compatible query language for Amazon DynamoDB

Amazon DynamoDB supports [PartiQL](https://partiql.org/), a SQL-compatible
query language, to select, insert, update, and delete data in Amazon DynamoDB. Using PartiQL, you
can easily interact with DynamoDB tables and run ad hoc queries using the AWS Management Console, NoSQL
Workbench, AWS Command Line Interface, and DynamoDB APIs for PartiQL.

PartiQL operations provide the same availability, latency, and performance as the other
DynamoDB data plane operations.

The following sections describe the DynamoDB implementation of PartiQL.

###### Topics

- [What is PartiQL?](#ql-reference.what-is)

- [PartiQL in Amazon DynamoDB](#ql-reference.what-is)

- [Getting started](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-gettingstarted.html)

- [Data types](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.data-types.html)

- [Statements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.statements.html)

- [Functions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-functions.html)

- [Operators](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-operators.html)

- [Transactions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.multiplestatements.transactions.html)

- [Batch operations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-reference.multiplestatements.batching.html)

- [IAM policies](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ql-iam.html)

## What is PartiQL?

_PartiQL_ provides SQL-compatible query access across multiple data
stores containing structured data, semistructured data, and nested data. It is widely
used within Amazon and is now available as part of many AWS services, including
DynamoDB.

For the PartiQL specification and a tutorial on the core query language, see the
[PartiQL documentation](https://partiql.org/docs.html).

###### Note

- Amazon DynamoDB supports a _subset_ of the [PartiQL](https://partiql.org/) query language.

- Amazon DynamoDB does not support the [Amazon ion](http://amzn.github.io/ion-docs) data format or
Amazon Ion literals.

## PartiQL in Amazon DynamoDB

To run PartiQL queries in DynamoDB, you can use:

- The DynamoDB console

- The NoSQL Workbench

- The AWS Command Line Interface (AWS CLI)

- The DynamoDB APIs

For information about using these methods to access DynamoDB, see [Accessing\
DynamoDB](accessingdynamodb.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Scanning tables

Getting started
