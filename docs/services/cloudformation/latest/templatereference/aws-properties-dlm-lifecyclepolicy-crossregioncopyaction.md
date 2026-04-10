This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy CrossRegionCopyAction

**\[Event-based policies only\]** Specifies a cross-Region copy action for event-based policies.

###### Note

To specify a cross-Region copy rule for snapshot and AMI policies, use
[CrossRegionCopyRule](../../../../reference/dlm/latest/apireference/api-crossregioncopyrule.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConfiguration" : EncryptionConfiguration,
  "RetainRule" : CrossRegionCopyRetainRule,
  "Target" : String
}

```

### YAML

```yaml

  EncryptionConfiguration:
    EncryptionConfiguration
  RetainRule:
    CrossRegionCopyRetainRule
  Target: String

```

## Properties

`EncryptionConfiguration`

The encryption settings for the copied snapshot.

_Required_: Yes

_Type_: [EncryptionConfiguration](aws-properties-dlm-lifecyclepolicy-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetainRule`

Specifies a retention rule for cross-Region snapshot copies created by snapshot or
event-based policies, or cross-Region AMI copies created by AMI policies. After the
retention period expires, the cross-Region copy is deleted.

_Required_: No

_Type_: [CrossRegionCopyRetainRule](aws-properties-dlm-lifecyclepolicy-crossregioncopyretainrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The target Region.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w:\-\/\*]+$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateRule

CrossRegionCopyDeprecateRule

All content copied from https://docs.aws.amazon.com/.
