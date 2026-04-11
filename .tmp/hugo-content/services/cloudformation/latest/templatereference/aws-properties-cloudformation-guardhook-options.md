This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::GuardHook Options

Specifies the input parameters for a Guard Hook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputParams" : S3Location
}

```

### YAML

```yaml

  InputParams:
    S3Location

```

## Properties

`InputParams`

Specifies the S3 location where your input parameters are located.

_Required_: No

_Type_: [S3Location](aws-properties-cloudformation-guardhook-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::GuardHook

S3Location

All content copied from https://docs.aws.amazon.com/.
