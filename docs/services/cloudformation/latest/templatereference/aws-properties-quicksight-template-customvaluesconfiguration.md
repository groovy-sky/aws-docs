This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template CustomValuesConfiguration

The configuration of custom values for the destination parameter in `DestinationParameterValueConfiguration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomValues" : CustomParameterValues,
  "IncludeNullValue" : Boolean
}

```

### YAML

```yaml

  CustomValues:
    CustomParameterValues
  IncludeNullValue: Boolean

```

## Properties

`CustomValues`

Property description not available.

_Required_: Yes

_Type_: [CustomParameterValues](aws-properties-quicksight-template-customparametervalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeNullValue`

Includes the null value in custom action parameter values.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomParameterValues

DataBarsOptions

All content copied from https://docs.aws.amazon.com/.
