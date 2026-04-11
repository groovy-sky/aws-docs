# Using the EXISTS function with PartiQL for DynamoDB

You can use EXISTS to perform the same function as `ConditionCheck` does in
the [TransactWriteItems](transaction-apis.md#transaction-apis-txwriteitems) API. The EXISTS function can only be used in
transactions.

Given a value, returns `TRUE` if the value is a non-empty collection.
Otherwise, returns `FALSE`.

###### Note

This function can only be used in transactional operations.

## Syntax

```nohighlight

EXISTS ( statement )
```

## Arguments

`statement`

(Required) The SELECT statement that the function evaluates.

###### Note

The SELECT statement must specify a full primary key and one other condition.

## Return type

`bool`

## Examples

```nohighlight

EXISTS(
    SELECT * FROM "Music"
    WHERE "Artist" = 'Acme Band' AND "SongTitle" = 'PartiQL Rocks')
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Functions

Begins\_with

All content copied from https://docs.aws.amazon.com/.
