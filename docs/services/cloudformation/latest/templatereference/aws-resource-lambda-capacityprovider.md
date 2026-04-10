This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider

Creates a capacity provider that manages compute resources for Lambda functions

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::CapacityProvider",
  "Properties" : {
      "CapacityProviderName" : String,
      "CapacityProviderScalingConfig" : CapacityProviderScalingConfig,
      "InstanceRequirements" : InstanceRequirements,
      "KmsKeyArn" : String,
      "PermissionsConfig" : CapacityProviderPermissionsConfig,
      "Tags" : [ Tag, ... ],
      "VpcConfig" : CapacityProviderVpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::CapacityProvider
Properties:
  CapacityProviderName: String
  CapacityProviderScalingConfig:
    CapacityProviderScalingConfig
  InstanceRequirements:
    InstanceRequirements
  KmsKeyArn: String
  PermissionsConfig:
    CapacityProviderPermissionsConfig
  Tags:
    - Tag
  VpcConfig:
    CapacityProviderVpcConfig

```

## Properties

`CapacityProviderName`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[a-zA-Z-]*:lambda:(eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:\d{12}:capacity-provider:[a-zA-Z0-9-_]+)|[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityProviderScalingConfig`

The scaling configuration for the capacity provider.

_Required_: No

_Type_: [CapacityProviderScalingConfig](aws-properties-lambda-capacityprovider-capacityproviderscalingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceRequirements`

The instance requirements for compute resources managed by the capacity provider.

_Required_: No

_Type_: [InstanceRequirements](aws-properties-lambda-capacityprovider-instancerequirements.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

The ARN of the KMS key used to encrypt the capacity provider's resources.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionsConfig`

The permissions configuration for the capacity provider.

_Required_: Yes

_Type_: [CapacityProviderPermissionsConfig](aws-properties-lambda-capacityprovider-capacityproviderpermissionsconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A key-value pair that provides metadata for the capacity provider.

_Required_: No

_Type_: Array of [Tag](aws-properties-lambda-capacityprovider-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

The VPC configuration for the capacity provider.

_Required_: Yes

_Type_: [CapacityProviderVpcConfig](aws-properties-lambda-capacityprovider-capacityprovidervpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

`State`

The current state of the capacity provider.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VersionWeight

CapacityProviderPermissionsConfig

All content copied from https://docs.aws.amazon.com/.
