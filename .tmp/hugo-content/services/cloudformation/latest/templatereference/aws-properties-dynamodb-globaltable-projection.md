This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable Projection

Represents attributes that are copied (projected) from the table into an index. These
are in addition to the primary key attributes and index key attributes, which are
automatically projected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NonKeyAttributes" : [ String, ... ],
  "ProjectionType" : String
}

```

### YAML

```yaml

  NonKeyAttributes:
    - String
  ProjectionType: String

```

## Properties

`NonKeyAttributes`

Represents the non-key attribute names which will be projected into the index.

For global and local secondary indexes, the total count of
`NonKeyAttributes` summed across all of the secondary indexes, must not
exceed 100. If you project the same attribute into two different indexes, this counts as
two distinct attributes when determining the total. This limit only applies when you
specify the ProjectionType of `INCLUDE`. You still can specify the
ProjectionType of `ALL` to project all attributes from the source table, even
if the table has more than 100 attributes.

_Required_: No

_Type_: Array of String

_Maximum_: `20`

_Update requires_: Updates are not supported.

`ProjectionType`

The set of attributes that are projected into the index:

- `KEYS_ONLY` \- Only the index and primary keys are projected into the
index.

- `INCLUDE` \- In addition to the attributes described in
`KEYS_ONLY`, the secondary index will include other non-key
attributes that you specify.

- `ALL` \- All of the table attributes are projected into the
index.

When using the DynamoDB console, `ALL` is selected by default.

_Required_: No

_Type_: String

_Allowed values_: `ALL | KEYS_ONLY | INCLUDE`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PointInTimeRecoverySpecification

ReadOnDemandThroughputSettings

All content copied from https://docs.aws.amazon.com/.
