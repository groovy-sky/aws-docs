This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset DatasetContentVersionValue

The dataset whose latest contents are used as input to the notebook or application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatasetName" : String
}

```

### YAML

```yaml

  DatasetName: String

```

## Properties

`DatasetName`

The name of the dataset whose latest contents are used as input to the notebook or
application.

_Required_: Yes

_Type_: String

_Pattern_: `(^(?!_{2}))(^[a-zA-Z0-9_]+$)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetContentDeliveryRuleDestination

DeltaTime

All content copied from https://docs.aws.amazon.com/.
