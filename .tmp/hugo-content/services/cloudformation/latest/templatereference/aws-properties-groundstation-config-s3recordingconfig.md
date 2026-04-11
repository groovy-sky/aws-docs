This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config S3RecordingConfig

Provides information about how AWS Ground Station should save downlink data to S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketArn" : String,
  "Prefix" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  BucketArn: String
  Prefix: String
  RoleArn: String

```

## Properties

`BucketArn`

S3 Bucket where the data is written. The name of the S3 Bucket provided must begin with `aws-groundstation`.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:s3:::[A-Za-z0-9-]{1,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

The prefix of the S3 data object.
If you choose to use any optional keys for substitution, these values will be replaced with the corresponding information from your contact details.
For example, a prefix of `{satellite_id}/{year}/{month}/{day}/` will replaced with `fake_satellite_id/2021/01/10/`

_Optional keys for substitution_: `{satellite_id}` \| `{config-name}` \| `{config-id}` \| `{year}` \| `{month}` \| `{day}`

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_\-=/]|\{satellite_id\}|\{config\-name}|\{s3\-config-id}|\{year\}|\{month\}|\{day\}){1,900}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Defines the ARN of the role assumed for putting archives to S3.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:\n]+:iam::[^:\n]+:role\/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an S3RecordingConfig

The following example creates a Ground Station `S3RecordingConfig`

#### JSON

```json

{
    "S3RecordingConfig": {
      "BucketArn": "arn:aws:s3:us-west-2:123456789012:bucket-name",
      "Prefix": "{satellite_id}/{config-name}_{config-id}/{year}/{month}/{day}",
      "RoleArn": "arn:aws:iam::123456789012:role/BucketRole"
    }
}
```

#### YAML

```yaml

S3RecordingConfig:
    BucketArn: arn:aws:s3:us-west-2:123456789012:bucket-name
    Prefix: {satellite_id}/{config-name}_{config-id}/{year}/{month}/{day}
    RoleArn: arn:aws:iam::123456789012:role/BucketRole

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisDataStreamData

SpectrumConfig

All content copied from https://docs.aws.amazon.com/.
