This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate MetadataOptions

The metadata options for the instance. For more information, see [Instance metadata and user data](../../../ec2/latest/userguide/ec2-instance-metadata.md) in the
_Amazon EC2 User Guide_.

`MetadataOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpEndpoint" : String,
  "HttpProtocolIpv6" : String,
  "HttpPutResponseHopLimit" : Integer,
  "HttpTokens" : String,
  "InstanceMetadataTags" : String
}

```

### YAML

```yaml

  HttpEndpoint: String
  HttpProtocolIpv6: String
  HttpPutResponseHopLimit: Integer
  HttpTokens: String
  InstanceMetadataTags: String

```

## Properties

`HttpEndpoint`

Enables or disables the HTTP metadata endpoint on your instances. If the parameter is
not specified, the default state is `enabled`.

###### Note

If you specify a value of `disabled`, you will not be able to access
your instance metadata.

_Required_: No

_Type_: String

_Allowed values_: `disabled | enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpProtocolIpv6`

Enables or disables the IPv6 endpoint for the instance metadata service.

Default: `disabled`

_Required_: No

_Type_: String

_Allowed values_: `disabled | enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpPutResponseHopLimit`

The desired HTTP PUT response hop limit for instance metadata requests. The larger the
number, the further instance metadata requests can travel.

Default: `1`

Possible values: Integers from 1 to 64

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpTokens`

Indicates whether IMDSv2 is required.

- `optional` \- IMDSv2 is optional. You can choose whether to send a
session token in your instance metadata retrieval requests. If you retrieve IAM
role credentials without a session token, you receive the IMDSv1 role
credentials. If you retrieve IAM role credentials using a valid session token,
you receive the IMDSv2 role credentials.

- `required` \- IMDSv2 is required. You must send a session token in
your instance metadata retrieval requests. With this option, retrieving the IAM
role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not
available.

Default: If the value of `ImdsSupport` for the Amazon Machine Image (AMI)
for your instance is `v2.0`, the default is `required`.

_Required_: No

_Type_: String

_Allowed values_: `optional | required`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMetadataTags`

Set to `enabled` to allow access to instance tags from the instance
metadata. Set to `disabled` to turn off access to instance tags from the
instance metadata. For more information, see [View tags for your EC2\
instances using instance metadata](../../../ec2/latest/userguide/work-with-tags-in-imds.md).

Default: `disabled`

_Required_: No

_Type_: String

_Allowed values_: `disabled | enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MemoryMiB

Monitoring

All content copied from https://docs.aws.amazon.com/.
