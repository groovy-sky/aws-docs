This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution StatusCodes

A complex data type for the status codes that you specify that, when returned by a
primary origin, trigger CloudFront to failover to a second origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ Integer, ... ],
  "Quantity" : Integer
}

```

### YAML

```yaml

  Items:
    - Integer
  Quantity: Integer

```

## Properties

`Items`

The items (status codes) for an origin group.

_Required_: Yes

_Type_: Array of Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Quantity`

The number of status codes.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3OriginConfig

StringSchema

All content copied from https://docs.aws.amazon.com/.
