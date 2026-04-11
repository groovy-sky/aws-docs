This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable StreamSpecification

Represents the DynamoDB Streams configuration for a table in DynamoDB.

You can only modify this value for a `AWS::DynamoDB::GlobalTable` resource
configured for multi-Region eventual consistency (MREC, the default) if that resource
contains only one entry in `Replicas`. You must specify a value for this
property for a `AWS::DynamoDB::GlobalTable` resource configured for MREC with
more than one entry in `Replicas`. For Multi-Region Strong Consistency (MRSC),
Streams are not required and can be changed for existing tables.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StreamViewType" : String
}

```

### YAML

```yaml

  StreamViewType: String

```

## Properties

`StreamViewType`

When an item in the table is modified, `StreamViewType` determines what
information is written to the stream for this table. Valid values for
`StreamViewType` are:

- `KEYS_ONLY` \- Only the key attributes of the modified item are
written to the stream.

- `NEW_IMAGE` \- The entire item, as it appears after it was modified,
is written to the stream.

- `OLD_IMAGE` \- The entire item, as it appeared before it was modified,
is written to the stream.

- `NEW_AND_OLD_IMAGES` \- Both the new and the old item images of the
item are written to the stream.

_Required_: Yes

_Type_: String

_Allowed values_: `NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSESpecification

Tag

All content copied from https://docs.aws.amazon.com/.
