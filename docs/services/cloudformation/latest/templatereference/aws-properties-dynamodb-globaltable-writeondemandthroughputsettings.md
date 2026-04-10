This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable WriteOnDemandThroughputSettings

Sets the write request settings for a global table or a global secondary index. You can
only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxWriteRequestUnits" : Integer
}

```

### YAML

```yaml

  MaxWriteRequestUnits: Integer

```

## Properties

`MaxWriteRequestUnits`

Maximum number of write request settings for the specified replica of a global
table.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WarmThroughput

WriteProvisionedThroughputSettings

All content copied from https://docs.aws.amazon.com/.
