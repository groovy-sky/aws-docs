This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::AccessGrant

The `AWS::S3::AccessGrant` resource creates an access grant that gives a grantee access to your S3 data. The grantee can be an IAM user or role or a directory user, or group. Before you can create a grant, you must have an S3 Access Grants instance in the same Region as the S3 data. You can create an S3 Access Grants instance using the [AWS::S3::AccessGrantsInstance](../userguide/aws-resource-s3-accessgrantsinstance.md). You must also have registered at least one S3 data location in your S3 Access Grants instance using [AWS::S3::AccessGrantsLocation](../userguide/aws-resource-s3-accessgrantslocation.md).

Permissions

You must have the `s3:CreateAccessGrant` permission to use this resource.

Additional Permissions

For any directory identity - `sso:DescribeInstance` and `sso:DescribeApplication`

For directory users - `identitystore:DescribeUser`

For directory groups - `identitystore:DescribeGroup`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::AccessGrant",
  "Properties" : {
      "AccessGrantsLocationConfiguration" : AccessGrantsLocationConfiguration,
      "AccessGrantsLocationId" : String,
      "ApplicationArn" : String,
      "Grantee" : Grantee,
      "Permission" : String,
      "S3PrefixType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3::AccessGrant
Properties:
  AccessGrantsLocationConfiguration:
    AccessGrantsLocationConfiguration
  AccessGrantsLocationId: String
  ApplicationArn: String
  Grantee:
    Grantee
  Permission: String
  S3PrefixType: String
  Tags:
    - Tag

```

## Properties

`AccessGrantsLocationConfiguration`

The configuration options of the grant location. The grant location is the S3 path to the data to which you are granting access. It contains the `S3SubPrefix` field. The grant scope is the result of appending the subprefix to the location scope of the registered location.

_Required_: No

_Type_: [AccessGrantsLocationConfiguration](aws-properties-s3-accessgrant-accessgrantslocationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccessGrantsLocationId`

The ID of the registered location to which you are granting access. S3 Access Grants assigns this ID when you register the location. S3 Access Grants assigns the ID `default` to the default location `s3://` and assigns an auto-generated ID to other locations that you register.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationArn`

The Amazon Resource Name (ARN) of an AWSIAM Identity Center application associated with your Identity Center instance. If the grant includes an application ARN, the grantee can only access the S3 data through this application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Grantee`

The user, group, or role to which you are granting access. You can grant access to an IAM user or role. If you have added your corporate directory to AWSIAM Identity Center and associated your Identity Center instance with your S3 Access Grants instance, the grantee can also be a corporate directory user or group.

_Required_: Yes

_Type_: [Grantee](aws-properties-s3-accessgrant-grantee.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permission`

The type of access that you are granting to your S3 data, which can be set to one of the following values:

- `READ` – Grant read-only access to the S3 data.

- `WRITE` – Grant write-only access to the S3 data.

- `READWRITE` – Grant both read and write access to the S3 data.

_Required_: Yes

_Type_: String

_Allowed values_: `READ | WRITE | READWRITE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3PrefixType`

The type of `S3SubPrefix`. The only possible value is `Object`. Pass this value if the access grant scope is an object. Do not pass this value if the access grant scope is a bucket or a bucket and a prefix.

_Required_: No

_Type_: String

_Allowed values_: `Object`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The AWS resource tags that you are adding to the access grant. Each tag is a label consisting of a user-defined key and value. Tags can help you manage, identify, organize, search for, and filter resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-s3-accessgrant-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns some information about your access grant.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessGrantArn`

The ARN of the access grant.

`AccessGrantId`

The ID of the access grant. S3 Access Grants auto-generates this ID when you create the access grant.

`GrantScope`

The S3 path of the data to which you are granting access. It is the result of appending the `Subprefix` to the location scope.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3

AccessGrantsLocationConfiguration

All content copied from https://docs.aws.amazon.com/.
