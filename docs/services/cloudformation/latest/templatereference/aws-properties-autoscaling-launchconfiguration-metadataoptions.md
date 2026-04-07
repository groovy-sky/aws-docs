This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::LaunchConfiguration MetadataOptions

`MetadataOptions` is a property of [AWS::AutoScaling::LaunchConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-launchconfiguration.html) that describes metadata options for the
instances.

For more information, see [Configure the instance metadata options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-config.html#launch-configurations-imds) in the _Amazon EC2 Auto Scaling_
_User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpEndpoint" : String,
  "HttpPutResponseHopLimit" : Integer,
  "HttpTokens" : String
}

```

### YAML

```yaml

  HttpEndpoint: String
  HttpPutResponseHopLimit: Integer
  HttpTokens: String

```

## Properties

`HttpEndpoint`

This parameter enables or disables the HTTP metadata endpoint on your instances. If
the parameter is not specified, the default state is `enabled`.

###### Note

If you specify a value of `disabled`, you will not be able to access
your instance metadata.

_Required_: No

_Type_: String

_Allowed values_: `disabled | enabled`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HttpPutResponseHopLimit`

The desired HTTP PUT response hop limit for instance metadata requests. The larger the
number, the further instance metadata requests can travel.

Default: 1

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HttpTokens`

The state of token usage for your instance metadata requests. If the parameter is not
specified in the request, the default state is `optional`.

If the state is `optional`, you can choose to retrieve instance metadata
with or without a signed token header on your request. If you retrieve the IAM role
credentials without a token, the version 1.0 role credentials are returned. If you
retrieve the IAM role credentials using a valid signed token, the version 2.0 role
credentials are returned.

If the state is `required`, you must send a signed token header with any
instance metadata retrieval requests. In this state, retrieving the IAM role credentials
always returns the version 2.0 credentials; the version 1.0 credentials are not
available.

_Required_: No

_Type_: String

_Allowed values_: `optional | required`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlockDeviceMapping

AWS::AutoScaling::LifecycleHook
