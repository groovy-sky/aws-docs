This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group ConfigurationParameter

One parameter for a group configuration item. For details about service configurations
and how to construct them, see [Service configurations for resource\
groups](../../../../reference/arg/latest/apireference/about-slg.md) in the _AWS Resource Groups User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Name: String
  Values:
    - String

```

## Properties

`Name`

The name of the group configuration parameter. For the list of parameters that you can
use with each configuration item type, see [Supported resource\
types and parameters](../../../../reference/arg/latest/apireference/about-slg.md#about-slg-types) in the _AWS Resource Groups User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-z-]+`

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value or values to be used for the specified parameter. For the list of values you
can use with each parameter, see [Supported resource\
types and parameters](../../../../reference/arg/latest/apireference/about-slg.md#about-slg-types).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationItem

Query

All content copied from https://docs.aws.amazon.com/.
