This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Studio

The `AWS::EMR::Studio` resource specifies an Amazon EMR Studio. An EMR Studio is a web-based, integrated development environment for fully managed Jupyter notebooks that run on Amazon EMR clusters. For more information, see the [_Amazon EMR Management Guide_](../../../emr/latest/managementguide/emr-studio.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::Studio",
  "Properties" : {
      "AuthMode" : String,
      "DefaultS3Location" : String,
      "Description" : String,
      "EncryptionKeyArn" : String,
      "EngineSecurityGroupId" : String,
      "IdcInstanceArn" : String,
      "IdcUserAssignment" : String,
      "IdpAuthUrl" : String,
      "IdpRelayStateParameterName" : String,
      "Name" : String,
      "ServiceRole" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TrustedIdentityPropagationEnabled" : Boolean,
      "UserRole" : String,
      "VpcId" : String,
      "WorkspaceSecurityGroupId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMR::Studio
Properties:
  AuthMode: String
  DefaultS3Location: String
  Description: String
  EncryptionKeyArn: String
  EngineSecurityGroupId: String
  IdcInstanceArn: String
  IdcUserAssignment: String
  IdpAuthUrl: String
  IdpRelayStateParameterName: String
  Name: String
  ServiceRole: String
  SubnetIds:
    - String
  Tags:
    - Tag
  TrustedIdentityPropagationEnabled: Boolean
  UserRole: String
  VpcId: String
  WorkspaceSecurityGroupId: String

```

## Properties

`AuthMode`

Specifies whether the Studio authenticates users using IAM Identity Center or IAM.

_Required_: Yes

_Type_: String

_Allowed values_: `SSO | IAM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultS3Location`

The Amazon S3 location to back up EMR Studio Workspaces and notebook files.

_Required_: Yes

_Type_: String

_Pattern_: `^s3://.*`

_Minimum_: `6`

_Maximum_: `10280`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A detailed description of the Amazon EMR Studio.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The AWS KMS key identifier (ARN) used to encrypt Amazon EMR Studio workspace and notebook files when backed up to Amazon S3.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso-f|iso-e))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineSecurityGroupId`

The ID of the Amazon EMR Studio Engine security group. The Engine security group
allows inbound network traffic from the Workspace security group, and it must be in the
same VPC specified by `VpcId`.

_Required_: Yes

_Type_: String

_Pattern_: `^sg-[a-zA-Z0-9\-._]+$`

_Minimum_: `4`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdcInstanceArn`

The ARN of the IAM Identity Center instance the Studio application belongs to.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdcUserAssignment`

Indicates whether the Studio has `REQUIRED` or `OPTIONAL` IAM Identity Center user assignment. If the value is set to `REQUIRED`, users must be explicitly assigned to the Studio application to access the Studio.

_Required_: No

_Type_: String

_Allowed values_: `REQUIRED | OPTIONAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdpAuthUrl`

Your identity provider's authentication endpoint. Amazon EMR Studio redirects
federated users to this endpoint for authentication when logging in to a Studio with the
Studio URL.

_Required_: No

_Type_: String

_Pattern_: `^https://[0-9a-zA-Z]([-.\w]*[0-9a-zA-Z])(:[0-9]*)*([?/#].*)?$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdpRelayStateParameterName`

The name of your identity provider's `RelayState` parameter.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive name for the Amazon EMR Studio.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRole`

The Amazon Resource Name (ARN) of the IAM role that will be assumed by the Amazon EMR Studio. The service role provides a
way for Amazon EMR Studio to interoperate with other AWS services.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso-f|iso-e))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of subnet IDs to associate with the Amazon EMR Studio. A Studio can have
a maximum of 5 subnets. The subnets must belong to the VPC specified by `VpcId`.
Studio users can create a Workspace in any of the specified subnets.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-emr-studio-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedIdentityPropagationEnabled`

Indicates whether the Studio has Trusted identity propagation enabled. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserRole`

The Amazon Resource Name (ARN) of the IAM user role that will be assumed by users and groups logged in to a Studio. The
permissions attached to this IAM role can be scoped down for each user or group using
session policies. You only need to specify `UserRole` when you set `AuthMode` to `SSO`.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso-f|iso-e))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the
Studio.

_Required_: Yes

_Type_: String

_Pattern_: `^(vpc-[0-9a-f]{8}|vpc-[0-9a-f]{17})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkspaceSecurityGroupId`

The ID of the Workspace security group associated with the Amazon EMR Studio.
The Workspace security group allows outbound network traffic to resources in the Engine
security group and to the internet.

_Required_: Yes

_Type_: String

_Pattern_: `^sg-[a-zA-Z0-9\-._]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ID. For example:

`{ "Ref": "es-EXAMPLE12345678XXXXXXXXXXX" }`

Ref returns the ID of the Amazon EMR Studio.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the Amazon EMR Studio. For example: `arn:aws:elasticmapreduce:us-east-1:653XXXXXXXXX:studio/es-EXAMPLE12345678XXXXXXXXXXX`.

`StudioId`

The ID of the Amazon EMR Studio. For example: `es-EXAMPLE12345678XXXXXXXXXXX`.

`Url`

The unique access URL of the Amazon EMR Studio. For example: `https://es-EXAMPLE12345678XXXXXXXXXXX.emrstudio-prod.us-east-1.amazonaws.com`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyValue

Tag

All content copied from https://docs.aws.amazon.com/.
