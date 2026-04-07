This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormStyle

The `FormStyle` property specifies the configuration for the form's style.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HorizontalGap" : FormStyleConfig,
  "OuterPadding" : FormStyleConfig,
  "VerticalGap" : FormStyleConfig
}

```

### YAML

```yaml

  HorizontalGap:
    FormStyleConfig
  OuterPadding:
    FormStyleConfig
  VerticalGap:
    FormStyleConfig

```

## Properties

`HorizontalGap`

The spacing for the horizontal gap.

_Required_: No

_Type_: [FormStyleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formstyleconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OuterPadding`

The size of the outer padding for the form.

_Required_: No

_Type_: [FormStyleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formstyleconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerticalGap`

The spacing for the vertical gap.

_Required_: No

_Type_: [FormStyleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-formstyleconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FormInputValuePropertyBindingProperties

FormStyleConfig
