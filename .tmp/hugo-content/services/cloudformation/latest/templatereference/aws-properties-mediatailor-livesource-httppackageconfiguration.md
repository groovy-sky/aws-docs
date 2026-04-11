This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::LiveSource HttpPackageConfiguration

The HTTP package configuration properties for the requested VOD source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Path" : String,
  "SourceGroup" : String,
  "Type" : String
}

```

### YAML

```yaml

  Path: String
  SourceGroup: String
  Type: String

```

## Properties

`Path`

The relative path to the URL for this VOD source. This is combined with `SourceLocation::HttpConfiguration::BaseUrl` to form a valid URL.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceGroup`

The name of the source group. This has to match one of the `Channel::Outputs::SourceGroup`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The streaming protocol for this package configuration. Supported values are `HLS` and `DASH`.

_Required_: Yes

_Type_: String

_Allowed values_: `DASH | HLS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaTailor::LiveSource

Tag

All content copied from https://docs.aws.amazon.com/.
