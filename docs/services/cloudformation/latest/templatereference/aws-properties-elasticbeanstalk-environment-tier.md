This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::Environment Tier

Describes the environment tier for an [AWS::ElasticBeanstalk::Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html) resource. For more information, see [Environment Tiers](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features-managing-env-tiers.html) in the _AWS Elastic Beanstalk Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String,
  "Version" : String
}

```

### YAML

```yaml

  Name: String
  Type: String
  Version: String

```

## Properties

`Name`

The name of this environment tier.

Valid values:

- For _Web server tier_ – `WebServer`

- For _Worker tier_ – `Worker`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of this environment tier.

Valid values:

- For _Web server tier_ – `Standard`

- For _Worker tier_ – `SQS/HTTP`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the
latest compatible worker tier version.

###### Note

This member is deprecated. Any specific version that you set may become out of date.
We recommend leaving it unspecified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
