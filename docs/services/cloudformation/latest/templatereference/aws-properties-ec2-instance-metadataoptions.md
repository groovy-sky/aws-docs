This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance MetadataOptions

Specifies the metadata options for the instance.

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

Enables or disables the HTTP metadata endpoint on your instances.

If you specify a value of `disabled`, you cannot access your instance
metadata.

Default: `enabled`

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

The maximum number of hops that the metadata token can travel.

Possible values: Integers from 1 to 64

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpTokens`

Indicates whether IMDSv2 is required.

- `optional` \- IMDSv2 is optional, which means that you can use
either IMDSv2 or IMDSv1.

- `required` \- IMDSv2 is required, which means that IMDSv1 is
disabled, and you must use IMDSv2.

Default:

- If the value of `ImdsSupport` for the Amazon Machine Image (AMI)
for your instance is `v2.0` and the account level default is set to
`no-preference`, the default is `required`.

- If the value of `ImdsSupport` for the Amazon Machine Image (AMI)
for your instance is `v2.0`, but the account level default is set to
`V1 or V2`, the default is `optional`.

The default value can also be affected by other combinations of parameters. For more
information, see [Order of precedence for instance metadata options](../../../ec2/latest/userguide/configuring-instance-metadata-options.md#instance-metadata-options-order-of-precedence) in the
_Amazon EC2 User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `optional | required`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMetadataTags`

Set to `enabled` to allow access to instance tags from the instance
metadata. Set to `disabled` to turn off access to instance tags from the
instance metadata. For more information, see [View tags for your EC2\
instances using instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/work-with-tags-in-IMDS.html).

Default: `disabled`

_Required_: No

_Type_: String

_Allowed values_: `disabled | enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LicenseSpecification

NetworkInterface
