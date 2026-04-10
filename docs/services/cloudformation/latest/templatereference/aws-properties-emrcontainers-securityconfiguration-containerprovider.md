This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration ContainerProvider

The information about the container provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Info" : ContainerInfo,
  "Type" : String
}

```

### YAML

```yaml

  Id: String
  Info:
    ContainerInfo
  Type: String

```

## Properties

`Id`

The ID of the container cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z][A-Za-z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Info`

The information about the container cluster.

_Required_: No

_Type_: [ContainerInfo](aws-properties-emrcontainers-securityconfiguration-containerinfo.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the container provider. Amazon EKS is the only supported type as of
now.

_Required_: Yes

_Type_: String

_Allowed values_: `EKS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerInfo

EksInfo

All content copied from https://docs.aws.amazon.com/.
