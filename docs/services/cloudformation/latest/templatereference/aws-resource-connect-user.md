This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::User

Specifies a user account for an Amazon Connect instance.

For information about how to create user accounts using the Amazon Connect
console, see [Add Users](../../../connect/latest/adminguide/user-management.md) in the
_Amazon Connect Administrator Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::User",
  "Properties" : {
      "AfterContactWorkConfigs" : [ AfterContactWorkConfigPerChannel, ... ],
      "AutoAcceptConfigs" : [ AutoAcceptConfig, ... ],
      "DirectoryUserId" : String,
      "HierarchyGroupArn" : String,
      "IdentityInfo" : UserIdentityInfo,
      "InstanceArn" : String,
      "Password" : String,
      "PersistentConnectionConfigs" : [ PersistentConnectionConfig, ... ],
      "PhoneConfig" : UserPhoneConfig,
      "PhoneNumberConfigs" : [ PhoneNumberConfig, ... ],
      "RoutingProfileArn" : String,
      "SecurityProfileArns" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "Username" : String,
      "UserProficiencies" : [ UserProficiency, ... ],
      "VoiceEnhancementConfigs" : [ VoiceEnhancementConfig, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::User
Properties:
  AfterContactWorkConfigs:
    - AfterContactWorkConfigPerChannel
  AutoAcceptConfigs:
    - AutoAcceptConfig
  DirectoryUserId: String
  HierarchyGroupArn: String
  IdentityInfo:
    UserIdentityInfo
  InstanceArn: String
  Password: String
  PersistentConnectionConfigs:
    - PersistentConnectionConfig
  PhoneConfig:
    UserPhoneConfig
  PhoneNumberConfigs:
    - PhoneNumberConfig
  RoutingProfileArn: String
  SecurityProfileArns:
    - String
  Tags:
    - Tag
  Username: String
  UserProficiencies:
    - UserProficiency
  VoiceEnhancementConfigs:
    - VoiceEnhancementConfig

```

## Properties

`AfterContactWorkConfigs`

The list of after contact work (ACW) timeout configuration settings for each channel.

_Required_: No

_Type_: Array of [AfterContactWorkConfigPerChannel](aws-properties-connect-user-aftercontactworkconfigperchannel.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoAcceptConfigs`

The list of auto-accept configuration settings for each channel.

_Required_: No

_Type_: Array of [AutoAcceptConfig](aws-properties-connect-user-autoacceptconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectoryUserId`

The identifier of the user account in the directory used for identity management.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyGroupArn`

The Amazon Resource Name (ARN) of the user's hierarchy group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityInfo`

Information about the user identity.

_Required_: No

_Type_: [UserIdentityInfo](aws-properties-connect-user-useridentityinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

The user's password.

_Required_: No

_Type_: String

_Pattern_: `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d\S]{8,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PersistentConnectionConfigs`

The list of persistent connection configuration settings for each channel.

_Required_: No

_Type_: Array of [PersistentConnectionConfig](aws-properties-connect-user-persistentconnectionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneConfig`

Information about the phone configuration for the user.

_Required_: No

_Type_: [UserPhoneConfig](aws-properties-connect-user-userphoneconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumberConfigs`

The list of phone number configuration settings for each channel.

_Required_: No

_Type_: Array of [PhoneNumberConfig](aws-properties-connect-user-phonenumberconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingProfileArn`

The Amazon Resource Name (ARN) of the user's routing profile.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/routing-profile/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityProfileArns`

The Amazon Resource Name (ARN) of the user's security profile.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The
tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-connect-user-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name assigned to the user account.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9\_\-\.\@]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProficiencies`

One or more predefined attributes assigned to a user, with a numeric value that
indicates how their level of skill in a specified area.

_Required_: No

_Type_: Array of [UserProficiency](aws-properties-connect-user-userproficiency.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceEnhancementConfigs`

The list of voice enhancement configuration settings for each channel.

_Required_: No

_Type_: Array of [VoiceEnhancementConfig](aws-properties-connect-user-voiceenhancementconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the user. For example:

`{ "Ref": "myUser" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`UserArn`

The Amazon Resource Name (ARN) of the user.

## Examples

### Specify a user resource

The following example specifies a user resource for an Amazon Connect
instance. This example specifies a user under an Amazon Connect instance.
We recommend using a [dynamic reference](../userguide/security-best-practices.md#creds) to specify a password value or mask the parameter
with `NoEcho`.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a user for an Amazon Connect instance
Resources:
    User:
      Type: 'AWS::Connect::User'
      Properties:
        IdentityInfo:
          FirstName: 'firstname'
          LastName: 'lastname'
          Email: 'example@email.com'
        PhoneConfig:
          PhoneType: 'DESK_PHONE'
          AutoAccept: true
          DeskPhoneNumber: '+12345678902'
          AfterContactWorkTimeLimit: 10
        Username: 'exampleuser'
        InstanceArn: 'arn:aws:connect:region-name:aws-account-id:instance/instance-arn'
        RoutingProfileArn: 'arn:aws:connect:region-name:aws-account-id:instance/instance-arn/routing-profile/routing-arn'
        SecurityProfileArns: [arn:aws:connect:region-name:aws-account-id:instance/instance-arn/security-profile/security-arn]
        Password: !Ref password
        Tags:
          - Key: 'tagKey'
            Value: 'tagValue'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AfterContactWorkConfig

All content copied from https://docs.aws.amazon.com/.
