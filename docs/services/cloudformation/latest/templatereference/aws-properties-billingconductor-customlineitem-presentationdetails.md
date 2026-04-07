This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BillingConductor::CustomLineItem PresentationDetails

An object that defines how custom line item charges are presented in the bill, containing specifications for service presentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Service" : String
}

```

### YAML

```yaml

  Service: String

```

## Properties

`Service`

The service under which the custom line item charges will be presented. Must be a string between 1 and 128 characters matching the pattern `^[a-zA-Z0-9]+$`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LineItemFilter

Tag
