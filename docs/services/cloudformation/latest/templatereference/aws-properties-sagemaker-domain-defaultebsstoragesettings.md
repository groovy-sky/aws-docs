This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain DefaultEbsStorageSettings

A collection of default EBS storage settings that apply to spaces created within a domain or user profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultEbsVolumeSizeInGb" : Integer,
  "MaximumEbsVolumeSizeInGb" : Integer
}

```

### YAML

```yaml

  DefaultEbsVolumeSizeInGb: Integer
  MaximumEbsVolumeSizeInGb: Integer

```

## Properties

`DefaultEbsVolumeSizeInGb`

The default size of the EBS storage volume for a space.

_Required_: Yes

_Type_: Integer

_Minimum_: `5`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumEbsVolumeSizeInGb`

The maximum size of the EBS storage volume for a space.

_Required_: Yes

_Type_: Integer

_Minimum_: `5`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPosixUserConfig

DefaultSpaceSettings

All content copied from https://docs.aws.amazon.com/.
