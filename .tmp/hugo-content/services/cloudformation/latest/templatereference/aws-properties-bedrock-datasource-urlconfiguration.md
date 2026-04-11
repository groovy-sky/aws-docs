This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource UrlConfiguration

The configuration of web URLs that you want to crawl.
You should be authorized to crawl the URLs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SeedUrls" : [ SeedUrl, ... ]
}

```

### YAML

```yaml

  SeedUrls:
    - SeedUrl

```

## Properties

`SeedUrls`

One or more seed or starting point URLs.

_Required_: Yes

_Type_: Array of [SeedUrl](aws-properties-bedrock-datasource-seedurl.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformationLambdaConfiguration

VectorIngestionConfiguration

All content copied from https://docs.aws.amazon.com/.
