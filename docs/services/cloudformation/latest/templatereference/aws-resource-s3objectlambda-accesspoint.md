This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3ObjectLambda::AccessPoint

The `AWS::S3ObjectLambda::AccessPoint` resource specifies an Object Lambda
Access Point used to access a bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3ObjectLambda::AccessPoint",
  "Properties" : {
      "Name" : String,
      "ObjectLambdaConfiguration" : ObjectLambdaConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::S3ObjectLambda::AccessPoint
Properties:
  Name: String
  ObjectLambdaConfiguration:
    ObjectLambdaConfiguration

```

## Properties

`Name`

The name of this access point.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

_Minimum_: `3`

_Maximum_: `45`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectLambdaConfiguration`

A configuration used when creating an Object Lambda Access Point.

_Required_: Yes

_Type_: [ObjectLambdaConfiguration](aws-properties-s3objectlambda-accesspoint-objectlambdaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Alias.Status`

The status of the Object Lambda Access Point alias. Valid Values: `PROVISIONING` \| `READY`.

`Alias.Value`

The alias name value of the Object Lambda Access Point. For example: `myolap-1a4n8yjrb3kda96f67zwrwiiuse1a--ol-s3`.

`Arn`

Specifies the ARN for the Object Lambda Access
Point.

`CreationDate`

The date and time when the specified Object
Lambda Access Point was created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Object Lambda

Alias

All content copied from https://docs.aws.amazon.com/.
