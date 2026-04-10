This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::LayerVersion Content

A ZIP archive that contains the contents of an [Lambda layer](../../../lambda/latest/dg/configuration-layers.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3Key: String
  S3ObjectVersion: String

```

## Properties

`S3Bucket`

The Amazon S3 bucket of the layer archive.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z\.\-_]*(?<!\.)`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Key`

The Amazon S3 key of the layer archive.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3ObjectVersion`

For versioned objects, the version of the layer archive object to use.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Layer Content

The location of a layer archive in Amazon S3.

#### JSON

```json

        "Content": {
            "S3Bucket": "amzn-s3-demo-bucket",
            "S3Key": "layer.zip"
        }
```

#### YAML

```yaml

    Content:
      S3Bucket: amzn-s3-demo-bucket
      S3Key: layer.zip
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::LayerVersion

AWS::Lambda::LayerVersionPermission

All content copied from https://docs.aws.amazon.com/.
