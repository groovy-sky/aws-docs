This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::Index MetadataConfiguration

The metadata configuration for the vector index. This configuration allows you to specify which metadata keys should be treated as non-filterable.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NonFilterableMetadataKeys" : [ String, ... ]
}

```

### YAML

```yaml

  NonFilterableMetadataKeys:
    - String

```

## Properties

`NonFilterableMetadataKeys`

Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval. Unlike default metadata keys, these keys can't be used as query filters. Non-filterable metadata keys can be retrieved but can't be searched, queried, or filtered. You can access non-filterable metadata keys of your vectors after finding the vectors.

You can specify 1 to 10 non-filterable metadata keys. Each key must be 1 to 63 characters long.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `63 | 10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
