This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application DeployAsApplicationConfiguration

The information required to deploy a Kinesis Data Analytics Studio notebook as an
application with durable state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3ContentLocation" : S3ContentBaseLocation
}

```

### YAML

```yaml

  S3ContentLocation:
    S3ContentBaseLocation

```

## Properties

`S3ContentLocation`

The description of an Amazon S3 object that contains the Amazon Data Analytics
application, including the Amazon Resource Name (ARN) of the S3 bucket, the name of the
Amazon S3 object that contains the data, and the version number of the Amazon S3 object
that contains the data.

_Required_: Yes

_Type_: [S3ContentBaseLocation](aws-properties-kinesisanalyticsv2-application-s3contentbaselocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomArtifactConfiguration

EnvironmentProperties

All content copied from https://docs.aws.amazon.com/.
