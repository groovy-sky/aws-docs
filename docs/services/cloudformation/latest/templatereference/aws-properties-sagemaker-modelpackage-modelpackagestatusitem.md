This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelPackageStatusItem

Represents the overall status of a model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureReason" : String,
  "Name" : String,
  "Status" : String
}

```

### YAML

```yaml

  FailureReason: String
  Name: String
  Status: String

```

## Properties

`FailureReason`

if the overall status is `Failed`, the reason for the failure.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the model package for which the overall status is being reported.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status.

_Required_: Yes

_Type_: String

_Allowed values_: `NotStarted | Failed | InProgress | Completed`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelPackageStatusDetails

ModelQuality

All content copied from https://docs.aws.amazon.com/.
