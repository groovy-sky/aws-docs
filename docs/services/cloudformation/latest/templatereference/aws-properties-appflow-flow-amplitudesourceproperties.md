This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow AmplitudeSourceProperties

The properties that are applied when Amplitude is being used as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Object" : String
}

```

### YAML

```yaml

  Object: String

```

## Properties

`Object`

The object specified in the Amplitude flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AmplitudeSourceProperties](../../../../reference/appflow/1-0/apireference/api-amplitudesourceproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AggregationConfig

ConnectorOperator

All content copied from https://docs.aws.amazon.com/.
