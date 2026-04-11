This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::InstanceProfile

Provides information that defines an instance profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::InstanceProfile",
  "Properties" : {
      "AvailabilityZone" : String,
      "Description" : String,
      "InstanceProfileIdentifier" : String,
      "InstanceProfileName" : String,
      "KmsKeyArn" : String,
      "NetworkType" : String,
      "PubliclyAccessible" : Boolean,
      "SubnetGroupIdentifier" : String,
      "Tags" : [ Tag, ... ],
      "VpcSecurityGroups" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DMS::InstanceProfile
Properties:
  AvailabilityZone: String
  Description: String
  InstanceProfileIdentifier: String
  InstanceProfileName: String
  KmsKeyArn: String
  NetworkType: String
  PubliclyAccessible: Boolean
  SubnetGroupIdentifier: String
  Tags:
    - Tag
  VpcSecurityGroups:
    - String

```

## Properties

`AvailabilityZone`

The Availability Zone where the instance profile runs.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the instance profile. Descriptions can have up to 31 characters.
A description can contain only ASCII letters, digits, and hyphens ('-'). Also, it can't
end with a hyphen or contain two consecutive hyphens, and can only begin with a letter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileIdentifier`

The identifier of the instance profile. Identifiers must begin with a letter
and must contain only ASCII letters, digits, and hyphens. They can't end with
a hyphen, or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileName`

The user-friendly name for the instance profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt
the connection parameters for the instance profile.

If you don't specify a value for the `KmsKeyArn` parameter, then
AWS DMS uses an AWS owned encryption key to encrypt your resources.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkType`

Specifies the network type for the instance profile. A value of `IPV4`
represents an instance profile with IPv4 network type and only supports IPv4 addressing.
A value of `IPV6` represents an instance profile with IPv6 network type
and only supports IPv6 addressing. A value of `DUAL` represents an instance
profile with dual network type that supports IPv4 and IPv6 addressing.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Specifies the accessibility options for the instance profile. A value of
`true` represents an instance profile with a public IP address. A value of
`false` represents an instance profile with a private IP address. The default value
is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetGroupIdentifier`

The identifier of the subnet group that is associated with the instance profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-dms-instanceprofile-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroups`

The VPC security groups that are used with the instance profile.
The VPC security group must work with the VPC containing the instance profile.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`InstanceProfileArn`

The Amazon Resource Name (ARN) string that uniquely identifies the instance profile.

`InstanceProfileCreationTime`

The time the instance profile was created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
