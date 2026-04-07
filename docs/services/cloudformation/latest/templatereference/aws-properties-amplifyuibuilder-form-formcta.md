This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormCTA

The `FormCTA` property specifies the call to action button configuration for the form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cancel" : FormButton,
  "Clear" : FormButton,
  "Position" : String,
  "Submit" : FormButton
}

```

### YAML

```yaml

  Cancel:
    FormButton
  Clear:
    FormButton
  Position: String
  Submit:
    FormButton

```

## Properties

`Cancel`

Displays a cancel button.

_Required_: No

_Type_: [FormButton](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formbutton.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Clear`

Displays a clear button.

_Required_: No

_Type_: [FormButton](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formbutton.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

The position of the button.

_Required_: No

_Type_: String

_Allowed values_: `top | bottom | top_and_bottom`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Submit`

Displays a submit button.

_Required_: No

_Type_: [FormButton](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formbutton.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FormButton

FormDataTypeConfig
