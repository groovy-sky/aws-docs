This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage MetricsSource

Details about the metrics source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentDigest" : String,
  "ContentType" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  ContentDigest: String
  ContentType: String
  S3Uri: String

```

## Properties

`ContentDigest`

The hash key used for the metrics source.

_Required_: No

_Type_: String

_Pattern_: `^[Ss][Hh][Aa]256:[0-9a-fA-F]{64}$`

_Maximum_: `72`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContentType`

The metric source content type.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The S3 URI for the metrics source.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataProperties

ModelAccessConfig

All content copied from https://docs.aws.amazon.com/.
