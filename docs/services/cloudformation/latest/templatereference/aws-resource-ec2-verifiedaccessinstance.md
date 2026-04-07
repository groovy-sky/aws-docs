This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessInstance

An AWS Verified Access instance is a regional entity that evaluates application requests and grants
access only when your security requirements are met.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VerifiedAccessInstance",
  "Properties" : {
      "CidrEndpointsCustomSubDomain" : String,
      "Description" : String,
      "FipsEnabled" : Boolean,
      "LoggingConfigurations" : VerifiedAccessLogs,
      "Tags" : [ Tag, ... ],
      "VerifiedAccessTrustProviderIds" : [ String, ... ],
      "VerifiedAccessTrustProviders" : [ VerifiedAccessTrustProvider, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VerifiedAccessInstance
Properties:
  CidrEndpointsCustomSubDomain: String
  Description: String
  FipsEnabled: Boolean
  LoggingConfigurations:
    VerifiedAccessLogs
  Tags:
    - Tag
  VerifiedAccessTrustProviderIds:
    - String
  VerifiedAccessTrustProviders:
    - VerifiedAccessTrustProvider

```

## Properties

`CidrEndpointsCustomSubDomain`

The custom subdomain.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the AWS Verified Access instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FipsEnabled`

Indicates whether support for Federal Information Processing Standards (FIPS) is enabled on the instance.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfigurations`

The logging configuration for the Verified Access instances.

_Required_: No

_Type_: [VerifiedAccessLogs](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessinstance-verifiedaccesslogs.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessinstance-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifiedAccessTrustProviderIds`

The IDs of the AWS Verified Access trust providers.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifiedAccessTrustProviders`

The IDs of the AWS Verified Access trust providers.

_Required_: No

_Type_: Array of [VerifiedAccessTrustProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessinstance-verifiedaccesstrustprovider.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Verified Access instance.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CidrEndpointsCustomSubDomainNameServers`

The name servers.

`CreationTime`

The creation time.

`LastUpdatedTime`

The last updated time.

`VerifiedAccessInstanceId`

The ID of the Verified Access instance.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CloudWatchLogs
