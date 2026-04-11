This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelInput

Input object for the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataInputConfig" : String
}

```

### YAML

```yaml

  DataInputConfig: String

```

## Properties

`DataInputConfig`

The input configuration object for the model.

_Required_: Yes

_Type_: String

_Pattern_: `[\S\s]+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelDataSource

ModelMetrics

All content copied from https://docs.aws.amazon.com/.
