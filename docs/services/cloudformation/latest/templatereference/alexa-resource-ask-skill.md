This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# Alexa::ASK::Skill

The `Alexa::ASK::Skill` resource creates an Alexa skill that enables
customers to access new abilities. For more information about developing a skill, see
the [Build Skills with the Alexa Skills Kit developer documentation](https://developer.amazon.com/docs/ask-overviews/build-skills-with-the-alexa-skills-kit.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "Alexa::ASK::Skill",
  "Properties" : {
      "AuthenticationConfiguration" : AuthenticationConfiguration,
      "SkillPackage" : SkillPackage,
      "VendorId" : String
    }
}

```

### YAML

```yaml

Type: Alexa::ASK::Skill
Properties:
  AuthenticationConfiguration:
    AuthenticationConfiguration
  SkillPackage:
    SkillPackage
  VendorId: String

```

## Properties

`AuthenticationConfiguration`

Login with Amazon (LWA) configuration used to authenticate with the Alexa service.
Only Login with Amazon clients created through the [Amazon Developer Console](https://developer.amazon.com/lwa/sp/overview.html) are supported. The client ID, client secret, and refresh token are
required.

_Required_: Yes

_Type_: [AuthenticationConfiguration](alexa-properties-ask-skill-authenticationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkillPackage`

Configuration for the skill package that contains the components of the Alexa
skill. Skill packages are retrieved from an Amazon S3 bucket and key and used to create
and update the skill. For more information about the skill package format, see the
[Skill Package API Reference](https://developer.amazon.com/docs/smapi/skill-package-api-reference.html).

_Required_: Yes

_Type_: [SkillPackage](alexa-properties-ask-skill-skillpackage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VendorId`

The vendor ID associated with the Amazon developer account that will host the
skill. Details for retrieving the vendor ID are in [How to get your vendor ID](https://github.com/alexa/alexa-smarthome/wiki/How-to-get-your-vendor-ID). The provided LWA credentials must be linked to the
developer account associated with this vendor ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the skill ID, such as
amzn1.ask.skill.a3103cee-c48c-40a0-a2c9-251141888863.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Alexa Skill Resource Configuration

The following example retrieves the skill package from an S3 bucket and
provides credentials for authentication with the Alexa service through Login
with Amazon (LWA).

#### JSON

```json

"MySkill": {
  "Type": "Alexa::ASK::Skill",
  "Properties": {
    "SkillPackage": {
      "S3Bucket": "my-skill-packages",
      "S3Key": "skillpackage.zip",
      "S3BucketRole": { "Fn::GetAtt": [ "S3BucketReadRole", "Arn" ] },
      "Overrides": {
        "Manifest": {
          "apis": {
            "custom": {
              "endpoint": {
                "uri": { "Fn::GetAtt" : [ "SkillFunction", "Arn" ] }
              }
            }
          }
        }
      }
    },
    "AuthenticationConfiguration": {
      "ClientId": "amzn1.application-oa2-client.1234",
      "ClientSecret": "1234",
      "RefreshToken": "Atzr|1234"
    },
    "VendorId": "1234"
  }
}
```

#### YAML

```yaml

MySkill: Type: "Alexa::ASK::Skill"
  Properties:
    SkillPackage:
      S3Bucket: "my-skill-packages"
      S3Key: "skillpackage.zip"
      S3BucketRole: !GetAtt S3BucketReadRole.Arn
      Overrides:
        Manifest:
          apis:
            custom:
              endpoint:
                uri: !GetAtt SkillFunction.Arn
    AuthenticationConfiguration:
      ClientId: "amzn1.application-oa2-client.1234"
      ClientSecret: "1234"
      RefreshToken: "Atzr|1234"
    VendorId: "1234"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Alexa Skills Kit

AuthenticationConfiguration

All content copied from https://docs.aws.amazon.com/.
