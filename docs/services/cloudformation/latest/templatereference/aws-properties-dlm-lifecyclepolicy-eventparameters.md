This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy EventParameters

**\[Event-based policies only\]** Specifies an event that activates an event-based policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DescriptionRegex" : String,
  "EventType" : String,
  "SnapshotOwner" : [ String, ... ]
}

```

### YAML

```yaml

  DescriptionRegex: String
  EventType: String
  SnapshotOwner:
    - String

```

## Properties

`DescriptionRegex`

The snapshot description that can trigger the policy. The description pattern is specified using
a regular expression. The policy runs only if a snapshot with a description that matches the
specified pattern is shared with your account.

For example, specifying `^.*Created for policy: policy-1234567890abcdef0.*$`
configures the policy to run only if snapshots created by policy `policy-1234567890abcdef0`
are shared with your account.

_Required_: No

_Type_: String

_Pattern_: `[\p{all}]*`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventType`

The type of event. Currently, only snapshot sharing events are supported.

_Required_: Yes

_Type_: String

_Allowed values_: `shareSnapshot`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotOwner`

The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
The policy only runs if one of the specified AWS accounts shares a snapshot with your account.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

EventSource

All content copied from https://docs.aws.amazon.com/.
