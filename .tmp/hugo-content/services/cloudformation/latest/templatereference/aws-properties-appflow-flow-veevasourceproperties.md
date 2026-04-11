This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow VeevaSourceProperties

The properties that are applied when using Veeva as a flow source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentType" : String,
  "IncludeAllVersions" : Boolean,
  "IncludeRenditions" : Boolean,
  "IncludeSourceFiles" : Boolean,
  "Object" : String
}

```

### YAML

```yaml

  DocumentType: String
  IncludeAllVersions: Boolean
  IncludeRenditions: Boolean
  IncludeSourceFiles: Boolean
  Object: String

```

## Properties

`DocumentType`

The document type specified in the Veeva document extract flow.

_Required_: No

_Type_: String

_Pattern_: `[\s\w_-]+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeAllVersions`

Boolean value to include All Versions of files in Veeva document extract flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeRenditions`

Boolean value to include file renditions in Veeva document extract flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeSourceFiles`

Boolean value to include source files in Veeva document extract flow.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Object`

The object specified in the Veeva flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [VeevaSourceProperties](../../../../reference/appflow/1-0/apireference/api-veevasourceproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpsolverS3OutputFormatConfig

ZendeskDestinationProperties

All content copied from https://docs.aws.amazon.com/.
