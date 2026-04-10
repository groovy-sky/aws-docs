This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule PutItemInput

The input for the DynamoActionVS action that specifies the DynamoDB table to which
the message data will be written.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableName" : String
}

```

### YAML

```yaml

  TableName: String

```

## Properties

`TableName`

The table where the message data will be written.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutAssetPropertyValueEntry

RepublishAction

All content copied from https://docs.aws.amazon.com/.
