This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::BrowserSettings

This resource specifies browser settings that can be associated with a web portal. Once
associated with a web portal, browser settings control how the browser will behave once a
user starts a streaming session for the web portal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::BrowserSettings",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "BrowserPolicy" : String,
      "CustomerManagedKey" : String,
      "Tags" : [ Tag, ... ],
      "WebContentFilteringPolicy" : WebContentFilteringPolicy
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::BrowserSettings
Properties:
  AdditionalEncryptionContext:
    Key: Value
  BrowserPolicy: String
  CustomerManagedKey: String
  Tags:
    - Tag
  WebContentFilteringPolicy:
    WebContentFilteringPolicy

```

## Properties

`AdditionalEncryptionContext`

Additional encryption context of the browser settings.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BrowserPolicy`

A JSON string containing Chrome Enterprise policies that will be applied to all
streaming sessions.

_Required_: No

_Type_: String

_Pattern_: `^\{[\S\s]*\}\s*$`

_Minimum_: `2`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerManagedKey`

The custom managed key of the browser settings.

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to add to the browser settings resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesweb-browsersettings-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebContentFilteringPolicy`

The policy that specifies which URLs end users are allowed to access or which URLs or domain categories they are restricted from accessing for enhanced security.

_Required_: No

_Type_: [WebContentFilteringPolicy](aws-properties-workspacesweb-browsersettings-webcontentfilteringpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that the browser settings resource is associated with.

`BrowserSettingsArn`

The ARN of the browser settings.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon WorkSpaces Secure Browser

Tag

All content copied from https://docs.aws.amazon.com/.
