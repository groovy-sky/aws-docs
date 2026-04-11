This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::SiteToSiteVpnAttachment ProposedSegmentChange

Describes a proposed segment change. In some cases, the segment change must first be evaluated and accepted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachmentPolicyRuleNumber" : Integer,
  "SegmentName" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  AttachmentPolicyRuleNumber: Integer
  SegmentName: String
  Tags:
    - Tag

```

## Properties

`AttachmentPolicyRuleNumber`

The rule number in the policy document that applies to this change.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentName`

The name of the segment to change.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value tags that changed for the segment.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-sitetositevpnattachment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProposedNetworkFunctionGroupChange

Tag

All content copied from https://docs.aws.amazon.com/.
