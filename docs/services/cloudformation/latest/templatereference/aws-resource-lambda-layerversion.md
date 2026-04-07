This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::LayerVersion

The `AWS::Lambda::LayerVersion` resource creates a [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a ZIP archive.

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
[instruction set architectures](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CompatibleRuntimes`

A list of compatible [function\
runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Used for filtering with [ListLayers](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayers.html) and [ListLayerVersions](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayerVersions.html).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `15`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Content`

The function layer archive.

_Required_: Yes

_Type_: [Content](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-layerversion-content.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

Content
