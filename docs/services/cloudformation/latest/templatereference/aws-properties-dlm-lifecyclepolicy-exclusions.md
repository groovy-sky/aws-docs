This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy Exclusions

**\[Default policies only\]** Specifies exclusion parameters for volumes or instances for which you
do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs
for target resources that match any of the specified exclusion parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludeBootVolumes" : Boolean,
  "ExcludeTags" : [ Tag, ... ],
  "ExcludeVolumeTypes" : [ Json, ... ]
}

```

### YAML

```yaml

  ExcludeBootVolumes: Boolean
  ExcludeTags:
    - Tag
  ExcludeVolumeTypes:
    - Json

```

## Properties

`ExcludeBootVolumes`

**\[Default policies for EBS snapshots only\]** Indicates whether to exclude volumes that are attached to
instances as the boot volume. If you exclude boot volumes, only volumes attached as data
(non-boot) volumes will be backed up by the policy. To exclude boot volumes, specify
`true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeTags`

**\[Default policies for EBS-backed AMIs only\]** Specifies whether to exclude volumes that have specific tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dlm-lifecyclepolicy-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludeVolumeTypes`

**\[Default policies for EBS snapshots only\]** Specifies the volume types to exclude. Volumes of the specified
types will not be targeted by the policy.

_Required_: No

_Type_: Array of Json

_Minimum_: `0`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventSource

FastRestoreRule
