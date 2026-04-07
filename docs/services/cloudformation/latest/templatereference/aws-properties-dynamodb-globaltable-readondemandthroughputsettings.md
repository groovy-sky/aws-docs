This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReadOnDemandThroughputSettings

Sets the read request settings for a replica table or a replica global secondary index.
You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxReadRequestUnits" : Integer
}

```

### YAML

```yaml

  MaxReadRequestUnits: Integer

```

## Properties

`MaxReadRequestUnits`

Maximum number of read request units for the specified replica of a global table.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Projection

ReadProvisionedThroughputSettings
