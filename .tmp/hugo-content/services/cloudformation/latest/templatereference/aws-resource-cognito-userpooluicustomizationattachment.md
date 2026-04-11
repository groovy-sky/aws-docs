This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolUICustomizationAttachment

A container for the UI customization information for the hosted UI in a user
pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolUICustomizationAttachment",
  "Properties" : {
      "ClientId" : String,
      "CSS" : String,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolUICustomizationAttachment
Properties:
  ClientId: String
  CSS: String
  UserPoolId: String

```

## Properties

`ClientId`

The app client ID for your UI customization. When this value isn't present, the
customization applies to all user pool app clients that don't have client-level
settings..

_Required_: Yes

_Type_: String

_Pattern_: `[\w+]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CSS`

A plaintext CSS file that contains the custom fields that you want to apply to your
user pool or app client. To download a template, go to the Amazon Cognito console. Navigate to
your user pool _App clients_ tab, select _Login_
_pages_, edit _Hosted UI (classic) style_, and select
the link to `CSS template.css`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool where you want to apply branding to the classic hosted
UI.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the physicalResourceId, which is
“UserPoolUICustomizationAttachment-UserPoolId-ClientId". For example:

`{ "Ref":
                "UserPoolUICustomizationAttachment-us-east-1_FAKEPOOLID-2asc123fakeclientidajjulj6bh"
                }`

For the Amazon Cognito user pool domain
`UserPoolUICustomizationAttachment-us-east-1_FAKEPOOLID-2asc123fakeclientidajjulj6bh`,
Ref returns the name of the UI customization attachment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new UI customization attachment for a user pool

The following example sets UI customization settings in the referenced user
pool and client.

#### JSON

```json

{
   "UserPoolUICustomization":{
      "Type":"AWS::Cognito::UserPoolUICustomizationAttachment",
      "Properties":{
         "UserPoolId":{
            "Ref":"UserPool"
         },
         "ClientId":{
            "Ref":"Client"
         },
         "CSS":".banner-customizable {\nbackground:
        linear-gradient(#9940B8, #C27BDB)\n}"
      }
   }
}
```

#### YAML

```yaml

UserPoolUICustomization:
  Type: AWS::Cognito::UserPoolUICustomizationAttachment
  Properties:
    UserPoolId: !Ref UserPool
    ClientId: !Ref Client
    CSS: ".banner-customizable {
      background: linear-gradient(#9940B8, #C27BDB)
    }"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RiskExceptionConfigurationType

AWS::Cognito::UserPoolUser

All content copied from https://docs.aws.amazon.com/.
