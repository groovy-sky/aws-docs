This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset Metadata

Contains additional resource information needed for specific datasets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceArn" : String
}

```

### YAML

```yaml

  SourceArn: String

```

## Properties

`SourceArn`

The Amazon Resource Name (ARN) associated with the dataset. Currently, DataBrew
only supports ARNs from Amazon AppFlow.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JsonOptions

PathOptions

All content copied from https://docs.aws.amazon.com/.
