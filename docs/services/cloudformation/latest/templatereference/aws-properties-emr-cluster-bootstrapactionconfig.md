This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster BootstrapActionConfig

`BootstrapActionConfig` is a property of `AWS::EMR::Cluster` that can be used to run bootstrap actions on EMR clusters. You can use a bootstrap action to install software and configure EC2 instances for all cluster nodes before EMR installs and configures open-source big data applications on cluster instances. For more information, see [Create Bootstrap Actions to Install Additional Software](../../../emr/latest/managementguide/emr-plan-bootstrap.md) in the _Amazon EMR Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "ScriptBootstrapAction" : ScriptBootstrapActionConfig
}

```

### YAML

```yaml

  Name: String
  ScriptBootstrapAction:
    ScriptBootstrapActionConfig

```

## Properties

`Name`

The name of the bootstrap action.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScriptBootstrapAction`

The script run by the bootstrap action.

_Required_: Yes

_Type_: [ScriptBootstrapActionConfig](aws-properties-emr-cluster-scriptbootstrapactionconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoTerminationPolicy

CloudWatchAlarmDefinition

All content copied from https://docs.aws.amazon.com/.
