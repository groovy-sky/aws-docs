# Querying tables in DynamoDB

You can use the `Query` API operation in Amazon DynamoDB to find items based on
primary key values.

You must provide the name of the partition key attribute and a single value for that
attribute. `Query` returns all items with that partition key value. Optionally,
you can provide a sort key attribute and use a comparison operator to refine the search
results.

For more information on how to use `Query`, such as the request syntax,
response parameters, and additional examples, see [Query](../../../../reference/amazondynamodb/latest/apireference/api-query.md) in the
_Amazon DynamoDB API Reference_.

###### Topics

- [Key condition expressions for the Query operation in DynamoDB](query-keyconditionexpressions.md)

- [Filter expressions for the Query operation in DynamoDB](query-filterexpression.md)

- [Paginating table query results in DynamoDB](query-pagination.md)

- [Other aspects of working with the Query operation in DynamoDB](query-other.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Working with expired items

Key condition expressions for queries

All content copied from https://docs.aws.amazon.com/.
