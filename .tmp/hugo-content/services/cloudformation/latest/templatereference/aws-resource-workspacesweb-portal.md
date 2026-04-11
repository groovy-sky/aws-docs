This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::Portal

This resource specifies a web portal, which users use to start browsing sessions. A `Standard` web portal can't start browsing sessions unless you have at defined and associated an `IdentityProvider` and `NetworkSettings` resource. An `IAM Identity Center` web portal does not require an `IdentityProvider` resource.

For more information about web portals, see [What is Amazon WorkSpaces Secure Browser?](../../../workspaces-web/latest/adminguide/what-is-workspaces-web.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::Portal",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "AuthenticationType" : String,
      "BrowserSettingsArn" : String,
      "CustomerManagedKey" : String,
      "DataProtectionSettingsArn" : String,
      "DisplayName" : String,
      "InstanceType" : String,
      "IpAccessSettingsArn" : String,
      "MaxConcurrentSessions" : Number,
      "NetworkSettingsArn" : String,
      "PortalCustomDomain" : String,
      "SessionLoggerArn" : String,
      "Tags" : [ Tag, ... ],
      "TrustStoreArn" : String,
      "UserAccessLoggingSettingsArn" : String,
      "UserSettingsArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::Portal
Properties:
  AdditionalEncryptionContext:
    Key: Value
  AuthenticationType: String
  BrowserSettingsArn: String
  CustomerManagedKey: String
  DataProtectionSettingsArn: String
  DisplayName: String
  InstanceType: String
  IpAccessSettingsArn: String
  MaxConcurrentSessions: Number
  NetworkSettingsArn: String
  PortalCustomDomain: String
  SessionLoggerArn: String
  Tags:
    - Tag
  TrustStoreArn: String
  UserAccessLoggingSettingsArn: String
  UserSettingsArn: String

```

## Properties

`AdditionalEncryptionContext`

The additional encryption context of the portal.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthenticationType`

The type of authentication integration points used when signing into the web portal.
Defaults to `Standard`.

`Standard` web portals are authenticated directly through your identity
provider (IdP). User and group access to your web portal is controlled through your IdP.
You need to include an IdP resource in your template to integrate your IdP with your web
portal. Completing the configuration for your IdP requires exchanging WorkSpaces Secure Browser’s SP
metadata with your IdP’s IdP metadata. If your IdP requires the SP metadata first before
returning the IdP metadata, you should follow these steps:

1\. Create and deploy a CloudFormation template with a `Standard` portal with
no `IdentityProvider` resource.

2\. Retrieve the SP metadata using `Fn:GetAtt`, the WorkSpaces Secure Browser console, or
by the calling the `GetPortalServiceProviderMetadata` API.

3\. Submit the data to your IdP.

4\. Add an `IdentityProvider` resource to your CloudFormation template.

`
                            IAM Identity Center
                        ` web portals are authenticated through AWS IAM Identity Center. They provide additional features, such as IdP-initiated authentication. Identity
sources (including external identity provider integration) and other identity provider
information must be configured in IAM Identity Center. User and group assignment must be done through the WorkSpaces Secure Browser console. These cannot
be configured in CloudFormation.

_Required_: No

_Type_: String

_Allowed values_: `Standard | IAM_Identity_Center`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrowserSettingsArn`

The ARN of the browser settings that is associated with this web portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerManagedKey`

The customer managed key of the web portal.

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataProtectionSettingsArn`

The ARN of the data protection settings.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the web portal.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The type and resources of the underlying instance.

_Required_: No

_Type_: String

_Allowed values_: `standard.regular | standard.large | standard.xlarge`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAccessSettingsArn`

The ARN of the IP access settings that is associated with the web portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxConcurrentSessions`

The maximum number of concurrent sessions for the portal.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `5000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSettingsArn`

The ARN of the network settings that is associated with the web portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortalCustomDomain`

The custom domain of the web portal that users access in order to start streaming
sessions.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]?((?!-)([A-Za-z0-9-]*[A-Za-z0-9])\.)+[a-zA-Z0-9]+$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionLoggerArn`

The ARN of the session logger that is associated with the portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the web portal. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](aws-properties-workspacesweb-portal-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreArn`

The ARN of the trust store that is associated with the web portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAccessLoggingSettingsArn`

The ARN of the user access logging settings that is associated with the web
portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserSettingsArn`

The ARN of the user settings that is associated with the web portal.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:[a-zA-Z0-9\-]+:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:[a-zA-Z]+(\/[a-fA-F0-9\-]{36})+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`BrowserType`

The browser that users see when using a streaming session.

`CreationDate`

The creation date of the web portal.

`PortalArn`

The ARN of the web portal.

`PortalEndpoint`

The endpoint URL of the web portal that users access in order to start streaming sessions.

`PortalStatus`

The status of the web portal.

`RendererType`

The renderer that is used in streaming sessions.

`ServiceProviderSamlMetadata`

The SAML metadata of the service provider.

`StatusReason`

A message that explains why the web portal is in its current status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
