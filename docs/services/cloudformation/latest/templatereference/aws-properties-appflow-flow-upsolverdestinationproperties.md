This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow UpsolverDestinationProperties

The properties that are applied when Upsolver is used as a destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "S3OutputFormatConfig" : UpsolverS3OutputFormatConfig
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  S3OutputFormatConfig:
    UpsolverS3OutputFormatConfig

```

## Properties

`BucketName`

The Upsolver Amazon S3 bucket name in which Amazon AppFlow places the
transferred data.

_Required_: Yes

_Type_: String

_Pattern_: `^(upsolver-appflow)\S*`

_Minimum_: `16`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The object key for the destination Upsolver Amazon S3 bucket in which Amazon AppFlow places the files.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputFormatConfig`

The configuration that determines how data is formatted when Upsolver is used as the flow
destination.

_Required_: Yes

_Type_: [UpsolverS3OutputFormatConfig](aws-properties-appflow-flow-upsolvers3outputformatconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [UpsolverDestinationProperties](../../../../reference/appflow/1-0/apireference/api-upsolverdestinationproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TriggerConfig

UpsolverS3OutputFormatConfig

All content copied from https://docs.aws.amazon.com/.
