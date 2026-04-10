This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint ContainerLogRotationConfiguration

The settings for container log rotation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxFilesToKeep" : Integer,
  "RotationSize" : String
}

```

### YAML

```yaml

  MaxFilesToKeep: Integer
  RotationSize: String

```

## Properties

`MaxFilesToKeep`

The number of files to keep in container after rotation.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RotationSize`

The file size at which to rotate logs. Minimum of 2KB, Maximum of 2GB.

_Required_: Yes

_Type_: String

_Pattern_: `^\d+(\.\d+)?[KMG][Bb]?$`

_Minimum_: `3`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationOverrides

EMREKSConfiguration

All content copied from https://docs.aws.amazon.com/.
