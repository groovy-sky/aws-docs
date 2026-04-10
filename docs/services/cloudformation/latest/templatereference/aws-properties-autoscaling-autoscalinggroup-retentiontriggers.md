This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup RetentionTriggers

Defines the specific triggers that cause instances to be retained in a Retained state rather than terminated. Each trigger corresponds to a different failure scenario during the instance lifecycle. This allows fine-grained control over when to preserve instances for manual intervention.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TerminateHookAbandon" : String
}

```

### YAML

```yaml

  TerminateHookAbandon: String

```

## Properties

`TerminateHookAbandon`

Specifies the action when a termination lifecycle hook is abandoned due to failure, timeout, or explicit abandonment (calling CompleteLifecycleAction).

Set to `retain` to move instances to a retained state. Set to `terminate` for default termination behavior.

Retained instances don't count toward desired capacity and remain until you call `TerminateInstanceInAutoScalingGroup`.

_Required_: No

_Type_: String

_Allowed values_: `retain | terminate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PerformanceFactorReferenceRequest

TagProperty

All content copied from https://docs.aws.amazon.com/.
