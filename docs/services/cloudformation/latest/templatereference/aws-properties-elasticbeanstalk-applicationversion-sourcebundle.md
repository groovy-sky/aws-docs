This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticBeanstalk::ApplicationVersion SourceBundle

The `SourceBundle` property is an embedded property of the [AWS::ElasticBeanstalk::ApplicationVersion](../userguide/aws-properties-beanstalk-sourcebundle.md) resource. It specifies the Amazon S3
location of the source bundle for an AWS Elastic Beanstalk application version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3Key" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3Key: String

```

## Properties

`S3Bucket`

The Amazon S3 bucket where the data is located.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Key`

The Amazon S3 key where the data is located.

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElasticBeanstalk::ApplicationVersion

AWS::ElasticBeanstalk::ConfigurationTemplate

All content copied from https://docs.aws.amazon.com/.
