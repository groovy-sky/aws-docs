This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::Profile

Creates a Profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RolesAnywhere::Profile",
  "Properties" : {
      "AcceptRoleSessionName" : Boolean,
      "AttributeMappings" : [ AttributeMapping, ... ],
      "DurationSeconds" : Number,
      "Enabled" : Boolean,
      "ManagedPolicyArns" : [ String, ... ],
      "Name" : String,
      "RequireInstanceProperties" : Boolean,
      "RoleArns" : [ String, ... ],
      "SessionPolicy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RolesAnywhere::Profile
Properties:
  AcceptRoleSessionName: Boolean
  AttributeMappings:
    - AttributeMapping
  DurationSeconds: Number
  Enabled: Boolean
  ManagedPolicyArns:
    - String
  Name: String
  RequireInstanceProperties: Boolean
  RoleArns:
    - String
  SessionPolicy: String
  Tags:
    - Tag

```

## Properties

`AcceptRoleSessionName`

Used to determine if a custom role session name will be accepted in a temporary credential request.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeMappings`

A mapping applied to the authenticating end-entity certificate.

_Required_: No

_Type_: Array of [AttributeMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rolesanywhere-profile-attributemapping.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DurationSeconds`

The number of seconds vended session credentials will be valid for

_Required_: No

_Type_: Number

_Minimum_: `900`

_Maximum_: `43200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

The enabled status of the resource.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedPolicyArns`

A list of managed policy ARNs. Managed policies identified by this list will be applied to the vended session credentials.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The customer specified name of the resource.

_Required_: Yes

_Type_: String

_Pattern_: `[ a-zA-Z0-9-_]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireInstanceProperties`

Specifies whether instance properties are required in CreateSession requests with this profile.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArns`

A list of IAM role ARNs that can be assumed when this profile is specified in a CreateSession request.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionPolicy`

A session policy that will applied to the trust boundary of the vended session credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of Tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rolesanywhere-profile-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The name of the Profile

### Fn::GetAtt

`ProfileArn`

The ARN of the profile.

`ProfileId`

The unique primary identifier of the Profile

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AttributeMapping
