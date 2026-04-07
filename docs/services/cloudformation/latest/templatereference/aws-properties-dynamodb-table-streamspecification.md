This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table StreamSpecification

Represents the DynamoDB Streams configuration for a table in DynamoDB.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourcePolicy" : ResourcePolicy,
  "StreamViewType" : String
}

```

### YAML

```yaml

  ResourcePolicy:
    ResourcePolicy
  StreamViewType: String

```

## Properties

`ResourcePolicy`

Creates or updates a resource-based policy document that contains the permissions for
DynamoDB resources, such as a table's streams. Resource-based policies let you
define access permissions by specifying who has access to each resource, and the actions
they are allowed to perform on each resource.

###### Note

When you remove the `StreamSpecification` property from the template, DynamoDB disables the stream but retains any attached resource policy until the stream is deleted after 24 hours. When you modify the `StreamViewType` property, DynamoDB creates a new stream and retains the old stream's resource policy. The old stream and its resource policy are deleted after the 24-hour retention period.

In a CloudFormation template, you can provide the policy in JSON or YAML format
because CloudFormation converts YAML to JSON before submitting it to DynamoDB. For more information about resource-based policies, see [Using\
resource-based policies for DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy\
examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html).

_Required_: No

_Type_: [ResourcePolicy](aws-properties-dynamodb-table-resourcepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SSESpecification

Tag
