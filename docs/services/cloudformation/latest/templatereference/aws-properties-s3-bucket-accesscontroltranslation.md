This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket AccessControlTranslation

Specify this only in a cross-account scenario (where source and destination bucket owners are not
the same), and you want to change replica ownership to the AWS account that owns the destination
bucket. If this is not specified in the replication configuration, the replicas are owned by same
AWS account that owns the source object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Owner" : String
}

```

### YAML

```yaml

  Owner: String

```

## Properties

`Owner`

Specifies the replica ownership. For default and valid values, see [PUT bucket replication](../../../s3/latest/api/restbucketputreplication.md) in the
_Amazon S3 API Reference_.

_Required_: Yes

_Type_: String

_Allowed values_: `Destination`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccelerateConfiguration

AnalyticsConfiguration

All content copied from https://docs.aws.amazon.com/.
