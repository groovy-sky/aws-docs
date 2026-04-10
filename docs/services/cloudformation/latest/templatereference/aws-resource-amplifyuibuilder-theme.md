This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Theme

The AWS::AmplifyUIBuilder::Theme resource specifies a theme within an Amplify app. A theme
is a collection of style settings that apply globally to the components associated with the
app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmplifyUIBuilder::Theme",
  "Properties" : {
      "AppId" : String,
      "EnvironmentName" : String,
      "Name" : String,
      "Overrides" : [ ThemeValues, ... ],
      "Tags" : {Key: Value, ...},
      "Values" : [ ThemeValues, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AmplifyUIBuilder::Theme
Properties:
  AppId: String
  EnvironmentName: String
  Name: String
  Overrides:
    - ThemeValues
  Tags:
    Key: Value
  Values:
    - ThemeValues

```

## Properties

`AppId`

The unique ID for the Amplify app associated with the theme.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentName`

The name of the backend environment that is a part of the Amplify
app.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the theme.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overrides`

Describes the properties that can be overriden to customize a theme.

_Required_: No

_Type_: Array of [ThemeValues](aws-properties-amplifyuibuilder-theme-themevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more key-value pairs to use when tagging the theme.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

A list of key-value pairs that defines the properties of the theme.

_Required_: No

_Type_: Array of [ThemeValues](aws-properties-amplifyuibuilder-theme-themevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time that the theme was created.

`Id`

The ID for the theme.

`ModifiedAt`

The time that the theme was modified.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValueMappings

ThemeValue

All content copied from https://docs.aws.amazon.com/.
