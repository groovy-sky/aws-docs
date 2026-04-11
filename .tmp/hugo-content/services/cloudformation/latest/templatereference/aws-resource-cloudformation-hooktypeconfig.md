This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::HookTypeConfig

The `AWS::CloudFormation::HookTypeConfig` resource specifies the
configuration of an activated Hook.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::HookTypeConfig",
  "Properties" : {
      "Configuration" : String,
      "ConfigurationAlias" : String,
      "TypeArn" : String,
      "TypeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::HookTypeConfig
Properties:
  Configuration: String
  ConfigurationAlias: String
  TypeArn: String
  TypeName: String

```

## Properties

`Configuration`

Specifies the activated Hook type configuration, in this AWS account
and AWS Region.

You must specify either `TypeName` and `Configuration` or
`TypeArn` and `Configuration`.

_Required_: Conditional

_Type_: String

_Pattern_: `[\s\S]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationAlias`

An alias by which to refer to this configuration data.

Defaults to `default` alias. Hook types currently support default
configuration alias.

_Required_: No

_Type_: String

_Allowed values_: `default`

_Pattern_: `^[a-zA-Z0-9]{1,256}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TypeArn`

The Amazon Resource Number (ARN) for the Hook to set `Configuration`
for.

You must specify either `TypeName` and `Configuration` or
`TypeArn` and `Configuration`.

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The unique name for your Hook. Specifies a three-part namespace for your Hook, with a
recommended pattern of `Organization::Service::Hook`.

You must specify either `TypeName` and `Configuration` or
`TypeArn` and `Configuration`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the resource version. For example:

`arn:aws:cloudformation:us-west-2:123456789012:type-configuration/hook/My-Sample-Hook/default`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfigurationArn`

The Amazon Resource Number (ARN) of the activated Hook type configuration in this
account and Region.

## Examples

- [Specifying the configuration of a Hook using TypeName](#aws-resource-cloudformation-hooktypeconfig--examples--Specifying_the_configuration_of_a_Hook_using_TypeName)

- [Specifying the configuration of a Hook using TypeArn](#aws-resource-cloudformation-hooktypeconfig--examples--Specifying_the_configuration_of_a_Hook_using_TypeArn)

### Specifying the configuration of a Hook using TypeName

The following example demonstrates how to specify a new Hook configuration
with the `TypeName` property type.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "HookTypeConfig": {
            "Type": "AWS::CloudFormation::HookTypeConfig",
            "Properties": {
                "TypeName": "My::Sample::Hook",
                "Configuration": "{\"CloudFormationConfiguration\":{\"HookConfiguration\":{\"HookInvocationStatus\":\"ENABLED\",\"FailureMode\":\"WARN\",\"Properties\":{\"limitSize\": \"1\",\"encryptionAlgorithm\": \"aws:kms\"}}}}"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  HookTypeConfig:
    Type: AWS::CloudFormation::HookTypeConfig
    Properties:
      TypeName: My::Sample::Hook
      Configuration: '{"CloudFormationConfiguration":{"HookConfiguration":{"HookInvocationStatus":"ENABLED","FailureMode":"WARN","Properties":{"limitSize": "1","encryptionAlgorithm": "aws:kms"}}}}'
```

### Specifying the configuration of a Hook using TypeArn

The following example demonstrates how to specify a new Hook configuration
with the `TypeArn` property type.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "HookTypeConfig": {
            "Type": "AWS::CloudFormation::HookTypeConfig",
            "Properties": {
                "TypeArn": "arn:aws:cloudformation:us-west-2:123456789012:type/hook/My-Sample-Hook",
                "Configuration": "{\"CloudFormationConfiguration\":{\"HookConfiguration\":{\"HookInvocationStatus\":\"ENABLED\",\"FailureMode\":\"WARN\",\"Properties\":{\"limitSize\": \"1\",\"encryptionAlgorithm\": \"aws:kms\"}}}}"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  HookTypeConfig:
    Type: AWS::CloudFormation::HookTypeConfig
    Properties:
      TypeArn: arn:aws:cloudformation:us-west-2:123456789012:type/hook/My-Sample-Hook
      Configuration: '{"CloudFormationConfiguration":{"HookConfiguration":{"HookInvocationStatus":"ENABLED","FailureMode":"WARN","Properties":{"limitSize": "1","encryptionAlgorithm": "aws:kms"}}}}'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::HookDefaultVersion

AWS::CloudFormation::HookVersion

All content copied from https://docs.aws.amazon.com/.
