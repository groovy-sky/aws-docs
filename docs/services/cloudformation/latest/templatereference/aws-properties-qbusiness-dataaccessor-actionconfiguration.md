This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor ActionConfiguration

Specifies an allowed action and its associated filter configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "FilterConfiguration" : ActionFilterConfiguration
}

```

### YAML

```yaml

  Action: String
  FilterConfiguration:
    ActionFilterConfiguration

```

## Properties

`Action`

The Amazon Q Business action that is allowed.

_Required_: Yes

_Type_: String

_Pattern_: `^qbusiness:[a-zA-Z]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

The filter configuration for the action, if any.

_Required_: No

_Type_: [ActionFilterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-dataaccessor-actionfilterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::QBusiness::DataAccessor

ActionFilterConfiguration
