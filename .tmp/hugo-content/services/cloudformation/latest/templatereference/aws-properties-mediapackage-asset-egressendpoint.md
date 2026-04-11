This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::Asset EgressEndpoint

The playback endpoint for a packaging configuration on an asset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PackagingConfigurationId" : String,
  "Url" : String
}

```

### YAML

```yaml

  PackagingConfigurationId: String
  Url: String

```

## Properties

`PackagingConfigurationId`

The ID of a packaging configuration that's applied to this asset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL that's used to request content from this endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaPackage::Asset

Tag

All content copied from https://docs.aws.amazon.com/.
