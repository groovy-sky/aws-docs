This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group ConfigurationItem

One of the items in the service configuration assigned to a resource group. A service
configuration can consist of one or more items. For details service configurations and
how to construct them, see [Service configurations for resource\
groups](https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html) in the _AWS Resource Groups User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : [ ConfigurationParameter, ... ],
  "Type" : String
}

```

### YAML

```yaml

  Parameters:
    - ConfigurationParameter
  Type: String

```

## Properties

`Parameters`

A collection of parameters for this configuration item. For the list of parameters
that you can use with each configuration item `Type`, see [Supported resource types and parameters](https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html#about-slg-types) in the _AWS Resource Groups User Guide_.

_Required_: No

_Type_: Array of [ConfigurationParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resourcegroups-group-configurationparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of configuration item. Each item must have a unique value for type.
For the list of the types that you can specify for a configuration item, see [Supported resource types and parameters](https://docs.aws.amazon.com/ARG/latest/APIReference/about-slg.html#about-slg-types) in the _AWS Resource Groups User Guide_.

_Required_: No

_Type_: String

_Pattern_: `AWS::[a-zA-Z0-9]+::[a-zA-Z0-9]+`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ResourceGroups::Group

ConfigurationParameter
