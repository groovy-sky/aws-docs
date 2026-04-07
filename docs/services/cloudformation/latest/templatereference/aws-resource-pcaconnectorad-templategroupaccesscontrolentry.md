This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::TemplateGroupAccessControlEntry

Create a group access control entry. Allow or deny Active Directory groups from enrolling and/or
autoenrolling with the template based on the group security identifiers (SIDs).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorAD::TemplateGroupAccessControlEntry",
  "Properties" : {
      "AccessRights" : AccessRights,
      "GroupDisplayName" : String,
      "GroupSecurityIdentifier" : String,
      "TemplateArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorAD::TemplateGroupAccessControlEntry
Properties:
  AccessRights:
    AccessRights
  GroupDisplayName: String
  GroupSecurityIdentifier: String
  TemplateArn: String

```

## Properties

`AccessRights`

Permissions to allow or deny an Active Directory group to enroll or autoenroll certificates issued
against a template.

_Required_: Yes

_Type_: [AccessRights](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pcaconnectorad-templategroupaccesscontrolentry-accessrights.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupDisplayName`

Name of the Active Directory group. This name does not need to match the group name in Active Directory.

_Required_: Yes

_Type_: String

_Pattern_: `^[\x20-\x7E]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupSecurityIdentifier`

Security identifier (SID) of the group object from Active Directory. The SID starts with
"S-".

_Required_: No

_Type_: String

_Pattern_: `^S-[0-9]-([0-9]+-){1,14}[0-9]+$`

_Minimum_: `7`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TemplateArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w-]+:pca-connector-ad:[\w-]+:[0-9]+:connector(\/[\w-]+)\/template(\/[\w-]+)$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ValidityPeriod

AccessRights
