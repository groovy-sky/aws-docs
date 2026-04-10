This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster ScriptBootstrapActionConfig

`ScriptBootstrapActionConfig` is a subproperty of the `BootstrapActionConfig` property type. `ScriptBootstrapActionConfig` specifies the arguments and location of the bootstrap script for EMR to run on all cluster nodes before it installs open-source big data applications on them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Args" : [ String, ... ],
  "Path" : String
}

```

### YAML

```yaml

  Args:
    - String
  Path: String

```

## Properties

`Args`

A list of command line arguments to pass to the bootstrap action script.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

Location in Amazon S3 of the script to run during a bootstrap action.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalingTrigger

SimpleScalingPolicyConfiguration

All content copied from https://docs.aws.amazon.com/.
