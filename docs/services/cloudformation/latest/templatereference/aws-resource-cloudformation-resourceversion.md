This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::ResourceVersion

The `AWS::CloudFormation::ResourceVersion` resource registers a resource
version with the CloudFormation registry. Registering a resource version makes
it available for use in CloudFormation templates in your AWS account, and includes:

- Validating the resource schema.

- Determining which handlers, if any, have been specified for the
resource.

- Making the resource available for use in your account.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

You can have a maximum of 50 resource versions registered at a time. This maximum is
per account and per Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::ResourceVersion",
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

Type: AWS::CloudFormation::ResourceVersion
Properties:
  ExecutionRoleArn: String
  LoggingConfig:
    LoggingConfig
  SchemaHandlerPackage: String
  TypeName: String

```

## Properties

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role for CloudFormation to assume when invoking the resource. If your resource calls AWS APIs in any of its handlers, you must create an IAM
execution role that includes the necessary permissions to call those AWS
APIs, and provision that execution role in your account. When CloudFormation
needs to invoke the resource type handler, CloudFormation assumes this
execution role to create a temporary session token, which it then passes to the resource
type handler, thereby supplying your resource type with the appropriate
credentials.

_Required_: No

_Type_: String

_Pattern_: `arn:.+:iam::[0-9]{12}:role/.+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfig`

Logging configuration information for a resource.

_Required_: No

_Type_: [LoggingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-resourceversion-loggingconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaHandlerPackage`

A URL to the S3 bucket for the resource project package that contains the necessary
files for the resource you want to register.

For information on generating a schema handler package, see [Modeling resource\
types to use with CloudFormation](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-model.html) in the _CloudFormation_
_Command Line Interface (CLI) User Guide_.

###### Note

To register the resource version, you must have `s3:GetObject`
permissions to access the S3 objects.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TypeName`

The name of the resource being registered.

We recommend that resource names adhere to the following pattern:
_company\_or\_organization_:: _service_:: _type_.

###### Note

The following organization namespaces are reserved and can't be used in your
resource names:

- `Alexa`

- `AMZN`

- `Amazon`

- `AWS`

- `Custom`

- `Dev`

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the resource version. For example:

`arn:aws:cloudformation:us-west-2:123456789012:type/resource/Sample-CloudFormation-Resource/00000001`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resource.

`IsDefaultVersion`

Whether the specified resource version is set as the default version.

This applies only to private extensions you have registered in your account, and
extensions published by AWS. For public third-party extensions, whether
they are activated in your account, CloudFormation returns
`null`.

`ProvisioningType`

For resource type extensions, the provisioning behavior of the resource type. CloudFormation determines the provisioning type during registration, based on the
types of handlers in the schema handler package submitted.

Possible values:

- `FULLY_MUTABLE`: The resource type includes an update handler to
process updates to the type during stack update operations.

- `IMMUTABLE`: The resource type doesn't include an update handler,
so the type can't be updated and must instead be replaced during stack update
operations.

- `NON_PROVISIONABLE`: The resource type doesn't include all the
following handlers, and therefore can't actually be provisioned.

- create

- read

- delete

`TypeArn`

The Amazon Resource Name (ARN) for the extension.

`VersionId`

The ID of a specific version of the resource. The version ID is the value at the end
of the Amazon Resource Name (ARN) assigned to the resource version when it is
registered.

`Visibility`

The visibility level that determines who can see and use this resource in CloudFormation operations:

- `PRIVATE`: The resource is only visible and usable within the
account where it was registered. CloudFormation automatically marks any
resources you register as `PRIVATE`.

- `PUBLIC`: The resource is publicly visible and usable within any
AWS account.

## Examples

- [Specifying a resource version](#aws-resource-cloudformation-resourceversion--examples--Specifying_a_resource_version)

- [Specifying a resource version and setting it as the default version](#aws-resource-cloudformation-resourceversion--examples--Specifying_a_resource_version_and_setting_it_as_the_default_version)

### Specifying a resource version

The following example demonstrates how to specify a new resource
version.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "ResourceVersion": {
            "Type": "AWS::CloudFormation::ResourceVersion",
            "Properties": {
                "TypeName": "My::Sample::Resource",
                "SchemaHandlerPackage": "s3://amzn-s3-demo-bucket/my-sample-resource.zip"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  ResourceVersion:
    Type: AWS::CloudFormation::ResourceVersion
    Properties:
      TypeName: My::Sample::Resource
      SchemaHandlerPackage: s3://amzn-s3-demo-bucket/my-sample-resource.zip
```

### Specifying a resource version and setting it as the default version

The following example demonstrates how to specify and new resource version,
and use the `Ref` return value to set that version as the default
version.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "ResourceVersion": {
            "Type": "AWS::CloudFormation::ResourceVersion",
            "Properties": {
                "TypeName": "My::Sample::Resource",
                "SchemaHandlerPackage": "s3://amzn-s3-demo-bucket/my-sample-resource.zip"
            }
        },
        "ResourceDefaultVersion": {
            "Type": "AWS::CloudFormation::ResourceDefaultVersion",
            "Properties": {
                "TypeVersionArn": {
                    "Ref": "ResourceVersion"
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
  ResourceVersion:
    Type: AWS::CloudFormation::ResourceVersion
    Properties:
      TypeName: My::Sample::Resource
      SchemaHandlerPackage: s3://amzn-s3-demo-bucket/my-sample-resource.zip
  ResourceDefaultVersion:
    Type: AWS::CloudFormation::ResourceDefaultVersion
    Properties:
      TypeVersionArn: !Ref ResourceVersion
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFormation::ResourceDefaultVersion

LoggingConfig
