This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream AuthenticationConfiguration

The authentication configuration of the Amazon MSK cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Connectivity" : String,
  "RoleARN" : String
}

```

### YAML

```yaml

  Connectivity: String
  RoleARN: String

```

## Properties

`Connectivity`

The type of connectivity used to access the Amazon MSK cluster.

_Required_: Yes

_Type_: String

_Allowed values_: `PUBLIC | PRIVATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleARN`

The ARN of the role used to access the Amazon MSK cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AmazonopensearchserviceRetryOptions

BufferingHints

All content copied from https://docs.aws.amazon.com/.
