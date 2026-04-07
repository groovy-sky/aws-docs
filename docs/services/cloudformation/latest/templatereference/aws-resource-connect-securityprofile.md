This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityProfile

Creates a security profile.

For information about security profiles, see [Security Profiles](https://docs.aws.amazon.com/connect/latest/adminguide/connect-security-profiles.html) in the _Amazon Connect Administrator Guide_. For a mapping of the API name and user interface name of the security
profile permissions, see [List\
of security profile permissions](https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::SecurityProfile",
  "Properties" : {
      "AllowedAccessControlHierarchyGroupId" : String,
      "AllowedAccessControlTags" : [ Tag, ... ],
      "AllowedFlowModules" : [ FlowModule, ... ],
      "Applications" : [ Application, ... ],
      "Description" : String,
      "GranularAccessControlConfiguration" : GranularAccessControlConfiguration,
      "HierarchyRestrictedResources" : [ String, ... ],
      "InstanceArn" : String,
      "Permissions" : [ String, ... ],
      "SecurityProfileName" : String,
      "TagRestrictedResources" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::SecurityProfile
Properties:
  AllowedAccessControlHierarchyGroupId: String
  AllowedAccessControlTags:
    - Tag
  AllowedFlowModules:
    - FlowModule
  Applications:
    - Application
  Description: String
  GranularAccessControlConfiguration:
    GranularAccessControlConfiguration
  HierarchyRestrictedResources:
    - String
  InstanceArn: String
  Permissions:
    - String
  SecurityProfileName: String
  TagRestrictedResources:
    - String
  Tags:
    - Tag

```

## Properties

`AllowedAccessControlHierarchyGroupId`

The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `0`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedAccessControlTags`

The list of tags that a security profile uses to restrict access to resources in Amazon Connect.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-securityprofile-tag.html)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedFlowModules`

Property description not available.

_Required_: No

_Type_: Array of [FlowModule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-securityprofile-flowmodule.html)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Applications`

Property description not available.

_Required_: No

_Type_: Array of [Application](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-securityprofile-application.html)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the security profile.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GranularAccessControlConfiguration`

The granular access control configuration for the security profile, including data table permissions.

_Required_: No

_Type_: [GranularAccessControlConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-securityprofile-granularaccesscontrolconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyRestrictedResources`

The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect. Following
are acceptable ResourceNames: `User`.

_Required_: No

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The identifier of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

Permissions assigned to the security profile. For a list of valid permissions, see
[List of security profile\
permissions](https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-list.html).

_Required_: No

_Type_: Array of String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityProfileName`

The name for the security profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[ a-zA-Z0-9_@-]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagRestrictedResources`

The list of resources that a security profile applies tag restrictions to in Amazon Connect.

_Required_: No

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, { "Tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-securityprofile-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the security profile. For example:

`{ "Ref": "mySecurityProfileName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LastModifiedRegion`

The AWS Region where this resource was last modified.

`LastModifiedTime`

The timestamp when this resource was last modified.

`SecurityProfileArn`

The Amazon Resource Name (ARN) of the security profile.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::SecurityKey

Application
