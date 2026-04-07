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

_Type_: [ContainerProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-containerprovider.html)

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

_Type_: [SecurityConfigurationData](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-securityconfigurationdata.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to assign to the security configuration.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrcontainers-securityconfiguration-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN (Amazon Resource Name) of the security configuration.

`Id`

The ID of the security configuration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AtRestEncryptionConfiguration
