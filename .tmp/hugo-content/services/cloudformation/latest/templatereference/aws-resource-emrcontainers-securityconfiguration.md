This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration

Creates a security configuration. Security configurations in Amazon EMR on EKS are
templates for different security setups. You can use security configurations to configure
the AWS Lake Formation integration setup. You can also create a security configuration
to re-use a security setup each time you create a virtual cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMRContainers::SecurityConfiguration",
  "Properties" : {
      "ContainerProvider" : ContainerProvider,
      "Name" : String,
      "SecurityConfigurationData" : SecurityConfigurationData,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EMRContainers::SecurityConfiguration
Properties:
  ContainerProvider:
    ContainerProvider
  Name: String
  SecurityConfigurationData:
    SecurityConfigurationData
  Tags:
    - Tag

```

## Properties

`ContainerProvider`

The information about the container provider.

_Required_: No

_Type_: [ContainerProvider](aws-properties-emrcontainers-securityconfiguration-containerprovider.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the security configuration.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfigurationData`

Security configuration inputs for the request.

_Required_: Yes

_Type_: [SecurityConfigurationData](aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to assign to the security configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-emrcontainers-securityconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN (Amazon Resource Name) of the security configuration.

`Id`

The ID of the security configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AtRestEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
