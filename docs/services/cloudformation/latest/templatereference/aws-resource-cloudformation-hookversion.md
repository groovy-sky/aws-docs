This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::HookVersion

The `AWS::CloudFormation::HookVersion` resource publishes new or first
version of a Hook to the CloudFormation registry.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

This resource type is not compatible with Guard and Lambda Hooks.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::HookVersion",
  "Properties" : {
      "ExecutionRoleArn" : String,
      "LoggingConfig" : LoggingConfig,
      "SchemaHandlerPackage" : String,
      "TypeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::HookVersion
Properties:
  ExecutionRoleArn: String
  LoggingConfig:
    LoggingConfig
  SchemaHandlerPackage: String
  TypeName: String

```

## Properties

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the task execution role that grants the Hook
permission.

_Required_: No

_Type_: String

_Pattern_: `arn:.+:iam::[0-9]{12}:role/.+`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfig`

Contains logging configuration information for an extension.

_Required_: No

_Type_: [LoggingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-hookversion-loggingconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaHandlerPackage`

A URL to the Amazon S3 bucket for the Hook project package that contains the
necessary files for the Hook you want to register.

For information on generating a schema handler package, see [Modeling custom\
CloudFormation Hooks](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-model.html) in the _CloudFormation Hooks_
_User Guide_.

###### Note

To register the Hook, you must have `s3:GetObject` permissions to
access the S3 objects.

_Required_: Yes

_Type_: String

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TypeName`

The unique name for your Hook. Specifies a three-part namespace for your Hook, with a
recommended pattern of `Organization::Service::Hook`.

###### Note

The following organization namespaces are reserved and can't be used in your Hook type
names:

- `Alexa`

- `AMZN`

- `Amazon`

- `ASK`

- `AWS`

- `Custom`

- `Dev`

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the Hook version. For example:

`arn:aws:cloudformation:us-west-2:123456789012:type/hook/Sample-CloudFormation-Hook/00000001`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the Hook.

`IsDefaultVersion`

Whether the specified Hook version is set as the default version.

`TypeArn`

The Amazon Resource Number (ARN) assigned to this version of the Hook.

`VersionId`

The ID of this version of the Hook.

`Visibility`

The visibility level that determines who can see and use this Hook in CloudFormation operations:

- `PRIVATE`: The Hook is only visible and usable within the account
where it was registered. CloudFormation automatically marks any Hooks
you register as `PRIVATE`.

- `PUBLIC`: The Hook is publicly visible and usable within any
AWS account.

## Examples

- [Specifying a Hook version](#aws-resource-cloudformation-hookversion--examples--Specifying_a_Hook_version)

- [Specifying the default Hook version](#aws-resource-cloudformation-hookversion--examples--Specifying_the_default_Hook_version)

### Specifying a Hook version

The following example demonstrates how to specify a new Hook version.

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
```

### Specifying the default Hook version

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
      TypeName: My::Sample::Hook
      SchemaHandlerPackage: s3://amzn-s3-demo-bucket/my-sample-hook.zip
  HookDefaultVersion:
    Type: AWS::CloudFormation::HookDefaultVersion
    Properties:
      TypeVersionArn: !Ref HookVersion
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFormation::HookTypeConfig

LoggingConfig
