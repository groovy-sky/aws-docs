This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::DataProtectionSettings

The data protection settings resource that can be associated with a web portal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::DataProtectionSettings",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "CustomerManagedKey" : String,
      "Description" : String,
      "DisplayName" : String,
      "InlineRedactionConfiguration" : InlineRedactionConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::DataProtectionSettings
Properties:
  AdditionalEncryptionContext:
    Key: Value
  CustomerManagedKey: String
  Description: String
  DisplayName: String
  InlineRedactionConfiguration:
    InlineRedactionConfiguration
  Tags:
    - Tag

```

## Properties

`AdditionalEncryptionContext`

The additional encryption context of the data protection settings.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerManagedKey`

The customer managed key used to encrypt sensitive information in the data protection
settings.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the data protection settings.

_Required_: No

_Type_: String

_Pattern_: `^[ _\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the data protection settings.

_Required_: No

_Type_: String

_Pattern_: `^[ _\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InlineRedactionConfiguration`

The inline redaction configuration for the data protection settings.

_Required_: No

_Type_: [InlineRedactionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the data protection settings.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-dataprotectionsettings-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this data protection settings resource is associated
with.

`CreationDate`

The creation date timestamp of the data protection settings.

`DataProtectionSettingsArn`

The ARN of the data protection settings resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WebContentFilteringPolicy

CustomPattern
