This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::HookDefaultVersion

The `AWS::CloudFormation::HookDefaultVersion` resource specifies the
default version of a Hook. The default version of the Hook is used in CloudFormation operations for this AWS account and AWS Region.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

This resource type is not compatible with Guard and Lambda Hooks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::HookDefaultVersion",
  "Properties" : {
      "TypeName" : String,
      "TypeVersionArn" : String,
      "VersionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::HookDefaultVersion
Properties:
  TypeName: String
  TypeVersionArn: String
  VersionId: String

```

## Properties

`TypeName`

The name of the Hook.

You must specify either `TypeVersionArn`, or `TypeName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeVersionArn`

The version ID of the type configuration.

You must specify either `TypeVersionArn`, or `TypeName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

The version ID of the type specified.

You must specify either `TypeVersionArn`, or `TypeName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the Hook version. For
example:

`arn:aws:cloudformation:us-west-2:123456789012:type/hook/Sample-CloudFormation-Hook/00000001`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Number (ARN) of the activated Hook in this account and
Region.

## Examples

- [Specifying a Hook default version](#aws-resource-cloudformation-hookdefaultversion--examples--Specifying_a_Hook_default_version)

- [Specifying a Hook default version using TypeVersionArn](#aws-resource-cloudformation-hookdefaultversion--examples--Specifying_a_Hook_default_version_using_TypeVersionArn)

- [Specifying a Hook default version using TypeName and VersionId](#aws-resource-cloudformation-hookdefaultversion--examples--Specifying_a_Hook_default_version_using_TypeName_and_VersionId)

### Specifying a Hook default version

The following example demonstrates how to specify a new Hook version and use
the `Ref` return value to set that version as the default
version.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "HookVersion": {
            "Type": "AWS::CloudFormation::HookVersion",
            "Properties": {
                "TypeName": "My::Sample::Hook",
                "SchemaHandlerPackage": "s3://amzn-s3-demo-bucket/my-sample-hook.zip"
            }
        },
        "HookDefaultVersion": {
            "Type": "AWS::CloudFormation::HookDefaultVersion",
            "Properties": {
                "TypeVersionArn": {
                    "Ref": "HookVersion"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  HookVersion:
    Type: AWS::CloudFormation::HookVersion
    Properties:
      TypeName: 'My::Sample::Hook'
      SchemaHandlerPackage: 's3://amzn-s3-demo-bucket/my-sample-hook.zip'
  HookDefaultVersion:
    Type: AWS::CloudFormation::HookDefaultVersion
    Properties:
      TypeVersionArn: !Ref HookVersion
```

### Specifying a Hook default version using TypeVersionArn

The following example demonstrates how to set a Hook version as the default
version through the `TypeVersionArn` property type.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "HookDefaultVersion": {
            "Type": "AWS::CloudFormation::HookDefaultVersion",
            "Properties": {
                "TypeVersionArn": "arn:aws:cloudformation:us-west-2:123456789012:type/hook/My-Sample-Hook/00000001"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  HookDefaultVersion:
    Type: AWS::CloudFormation::HookDefaultVersion
    Properties:
      TypeVersionArn: >-
        arn:aws:cloudformation:us-west-2:123456789012:type/hook/My-Sample-Hook/00000001
```

### Specifying a Hook default version using TypeName and VersionId

The following example demonstrates how to set a Hook version as the default
version through the `VersionId` property type.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "HookDefaultVersion": {
            "Type": "AWS::CloudFormation::HookDefaultVersion",
            "Properties": {
                "TypeName": "My::Sample::Hook",
                "VersionId": 1
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  HookDefaultVersion:
    Type: AWS::CloudFormation::HookDefaultVersion
    Properties:
      TypeName: 'My::Sample::Hook'
      VersionId: 1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetFiltersItems

AWS::CloudFormation::HookTypeConfig

All content copied from https://docs.aws.amazon.com/.
