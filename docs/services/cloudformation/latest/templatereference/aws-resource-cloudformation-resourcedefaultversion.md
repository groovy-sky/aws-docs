This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::ResourceDefaultVersion

The `AWS::CloudFormation::ResourceDefaultVersion` resource specifies the
default version of a resource. The default version of a resource will be used in CloudFormation operations.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::ResourceDefaultVersion",
  "Properties" : {
      "TypeName" : String,
      "TypeVersionArn" : String,
      "VersionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::ResourceDefaultVersion
Properties:
  TypeName: String
  TypeVersionArn: String
  VersionId: String

```

## Properties

`TypeName`

The name of the resource.

Conditional: You must specify either `TypeVersionArn`, or
`TypeName` and `VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeVersionArn`

The Amazon Resource Name (ARN) of the resource version.

Conditional: You must specify either `TypeVersionArn`, or
`TypeName` and `VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/resource/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

The ID of a specific version of the resource. The version ID is the value at the end
of the Amazon Resource Name (ARN) assigned to the resource version when it's
registered.

Conditional: You must specify either `TypeVersionArn`, or
`TypeName` and `VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the resource type. For example:

`arn:aws:cloudformation:us-west-2:123456789012:type/resource/Sample-CloudFormation-Resource`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resource.

## Examples

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

AWS::CloudFormation::Publisher

AWS::CloudFormation::ResourceVersion
