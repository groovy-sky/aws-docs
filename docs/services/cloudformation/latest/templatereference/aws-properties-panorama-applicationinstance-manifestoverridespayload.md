This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Panorama::ApplicationInstance ManifestOverridesPayload

Parameter overrides for an application instance. This is a JSON document that has a single key
( `PayloadData`) where the value is an escaped string representation of the overrides document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PayloadData" : String
}

```

### YAML

```yaml

  PayloadData: String

```

## Properties

`PayloadData`

The overrides document.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Panorama::ApplicationInstance

ManifestPayload
