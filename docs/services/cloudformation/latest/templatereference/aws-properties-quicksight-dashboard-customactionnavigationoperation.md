This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard CustomActionNavigationOperation

The navigation operation that navigates between different sheets in the same analysis.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LocalNavigationConfiguration" : LocalNavigationConfiguration
}

```

### YAML

```yaml

  LocalNavigationConfiguration:
    LocalNavigationConfiguration

```

## Properties

`LocalNavigationConfiguration`

The configuration that chooses the navigation target.

_Required_: No

_Type_: [LocalNavigationConfiguration](aws-properties-quicksight-dashboard-localnavigationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionFilterOperation

CustomActionSetParametersOperation

All content copied from https://docs.aws.amazon.com/.
