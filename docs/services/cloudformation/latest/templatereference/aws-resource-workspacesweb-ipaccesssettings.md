This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::IpAccessSettings

This resource specifies IP access settings that can be associated with a web portal. For
more information, see [Set up IP access controls (optional)](https://docs.aws.amazon.com/workspaces-web/latest/adminguide/ip-access-controls.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::IpAccessSettings",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "CustomerManagedKey" : String,
      "Description" : String,
      "DisplayName" : String,
      "IpRules" : [ IpRule, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::IpAccessSettings
Properties:
  AdditionalEncryptionContext:
    Key: Value
  CustomerManagedKey: String
  Description: String
  DisplayName: String
  IpRules:
    - IpRule
  Tags:
    - Tag

```

## Properties

`AdditionalEncryptionContext`

Additional encryption context of the IP access settings.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerManagedKey`

The custom managed key of the IP access settings.

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the IP access settings.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the IP access settings.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpRules`

The IP rules of the IP access settings.

_Required_: Yes

_Type_: Array of [IpRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-ipaccesssettings-iprule.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the IP access settings resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-ipaccesssettings-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this IP access settings resource is associated with.

`CreationDate`

The creation date timestamp of the IP access settings.

`IpAccessSettingsArn`

The ARN of the IP access settings resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IpRule
