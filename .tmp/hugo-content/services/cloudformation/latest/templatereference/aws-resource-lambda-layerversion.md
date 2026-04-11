This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::LayerVersion

The `AWS::Lambda::LayerVersion` resource creates a [Lambda layer](../../../lambda/latest/dg/configuration-layers.md) from a ZIP archive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::LayerVersion",
  "Properties" : {
      "CompatibleArchitectures" : [ String, ... ],
      "CompatibleRuntimes" : [ String, ... ],
      "Content" : Content,
      "Description" : String,
      "LayerName" : String,
      "LicenseInfo" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::LayerVersion
Properties:
  CompatibleArchitectures:
    - String
  CompatibleRuntimes:
    - String
  Content:
    Content
  Description: String
  LayerName: String
  LicenseInfo: String

```

## Properties

`CompatibleArchitectures`

A list of compatible
[instruction set architectures](../../../lambda/latest/dg/foundation-arch.md).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CompatibleRuntimes`

A list of compatible [function\
runtimes](../../../lambda/latest/dg/lambda-runtimes.md). Used for filtering with [ListLayers](../../../lambda/latest/dg/api-listlayers.md) and [ListLayerVersions](../../../lambda/latest/dg/api-listlayerversions.md).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Content`

The function layer archive.

_Required_: Yes

_Type_: [Content](aws-properties-lambda-layerversion-content.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the version.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LayerName`

The name or Amazon Resource Name (ARN) of the layer.

_Required_: No

_Type_: String

_Pattern_: `(arn:[a-zA-Z0-9-]+:lambda:[a-zA-Z0-9-]+:\d{12}:layer:[a-zA-Z0-9-_]+)|[a-zA-Z0-9-_]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LicenseInfo`

The layer's software license. It can be any of the following:

- An [SPDX license identifier](https://spdx.org/licenses). For example,
`MIT`.

- The URL of a license hosted on the internet. For example,
`https://opensource.org/licenses/MIT`.

- The full text of the license.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the layer version, such as
`arn:aws:lambda:us-east-2:123456789012:layer:my-layer:1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`LayerVersionArn`

The ARN of the layer version.

## Examples

### Layer Version

Create a layer named `my-layer`.

#### JSON

```json

"MyLayer": {
    "Type": "AWS::Lambda::LayerVersion",
    "Properties": {
        "CompatibleRuntimes": [
            "python3.12",
            "python3.11"
        ],
        "Content": {
            "S3Bucket": "amzn-s3-demo-bucket",
            "S3Key": "layer.zip"
        },
        "Description": "My layer",
        "LayerName": "my-layer",
        "LicenseInfo": "MIT"
    }
}
```

#### YAML

```yaml

MyLayer:
  Type: AWS::Lambda::LayerVersion
  Properties:
    CompatibleRuntimes:
      - python3.12
      - python3.11
    Content:
      S3Bucket: amzn-s3-demo-bucket
      S3Key: layer.zip
    Description: My layer
    LayerName: my-layer
    LicenseInfo: MIT
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

Content

All content copied from https://docs.aws.amazon.com/.
