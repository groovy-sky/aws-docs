This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster AutoTerminationPolicy

An auto-termination policy for an Amazon EMR cluster. An auto-termination policy
defines the amount of idle time in seconds after which a cluster automatically terminates.
For alternative cluster termination options, see [Control cluster\
termination](../../../emr/latest/managementguide/emr-plan-termination.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTimeout" : Integer
}

```

### YAML

```yaml

  IdleTimeout: Integer

```

## Properties

`IdleTimeout`

Specifies the amount of idle time in seconds after which the cluster automatically
terminates. You can specify a minimum of 60 seconds and a maximum of 604800 seconds (seven
days).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingPolicy

BootstrapActionConfig

All content copied from https://docs.aws.amazon.com/.
