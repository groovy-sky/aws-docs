This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::LogDeliveryConfiguration S3Configuration

Configuration for the Amazon S3 bucket destination of user activity log export with threat
protection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketArn" : String
}

```

### YAML

```yaml

  BucketArn: String

```

## Properties

`BucketArn`

The ARN of an Amazon S3 bucket that's the destination for threat protection log
export.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:::[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `3`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

AWS::Cognito::ManagedLoginBranding

All content copied from https://docs.aws.amazon.com/.
