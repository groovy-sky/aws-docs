This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form

The AWS::AmplifyUIBuilder::Form resource specifies all of the information that is required to create a form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmplifyUIBuilder::Form",
  "Properties" : {
      "AppId" : String,
      "Cta" : FormCTA,
      "DataType" : FormDataTypeConfig,
      "EnvironmentName" : String,
      "Fields" : {Key: Value, ...},
      "FormActionType" : String,
      "LabelDecorator" : String,
      "Name" : String,
      "SchemaVersion" : String,
      "SectionalElements" : {Key: Value, ...},
      "Style" : FormStyle,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::AmplifyUIBuilder::Form
Properties:
  AppId: String
  Cta:
    FormCTA
  DataType:
    FormDataTypeConfig
  EnvironmentName: String
  Fields:
    Key: Value
  FormActionType: String
  LabelDecorator: String
  Name: String
  SchemaVersion: String
  SectionalElements:
    Key: Value
  Style:
    FormStyle
  Tags:
    Key: Value

```

## Properties

`AppId`

The unique ID of the Amplify app associated with the form.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cta`

The `FormCTA` object that stores the call to action configuration for the
form.

_Required_: No

_Type_: [FormCTA](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formcta.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The type of data source to use to create the form.

_Required_: No

_Type_: [FormDataTypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formdatatypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentName`

The name of the backend environment that is a part of the Amplify app.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Fields`

The configuration information for the form's fields.

_Required_: No

_Type_: Object of [FieldConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-fieldconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormActionType`

Specifies whether to perform a create or update action on the form.

_Required_: No

_Type_: String

_Allowed values_: `create | update`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelDecorator`

Specifies an icon or decoration to display on the form.

_Required_: No

_Type_: String

_Allowed values_: `required | optional | none`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the form.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaVersion`

The schema version of the form.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SectionalElements`

The configuration information for the visual helper elements for the form. These elements
are not associated with any data.

_Required_: No

_Type_: Object of [SectionalElement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-sectionalelement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Style`

The configuration for the form's style.

_Required_: No

_Type_: [FormStyle](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formstyle.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more key-value pairs to use when tagging the form data.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for the form.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SortProperty

FieldConfig
