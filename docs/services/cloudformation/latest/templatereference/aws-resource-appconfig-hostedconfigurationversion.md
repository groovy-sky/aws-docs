This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::HostedConfigurationVersion

Create a new configuration in the AWS AppConfig hosted configuration store.
Configurations must be 1 MB or smaller. The AWS AppConfig hosted configuration store
provides the following benefits over other configuration store options.

- You don't need to set up and configure other services such as Amazon Simple Storage Service
(Amazon S3) or Parameter Store.

- You don't need to configure AWS Identity and Access Management (IAM)
permissions to use the configuration store.

- You can store configurations in any content type.

- There is no cost to use the store.

- You can create a configuration and add it to the store when you create a configuration
profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::HostedConfigurationVersion",
  "Properties" : {
      "ApplicationId" : String,
      "ConfigurationProfileId" : String,
      "Content" : String,
      "ContentType" : String,
      "Description" : String,
      "LatestVersionNumber" : Integer,
      "VersionLabel" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::HostedConfigurationVersion
Properties:
  ApplicationId: String
  ConfigurationProfileId: String
  Content: String
  ContentType: String
  Description: String
  LatestVersionNumber: Integer
  VersionLabel: String

```

## Properties

`ApplicationId`

The application ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationProfileId`

The configuration profile ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Content`

The configuration data, as bytes.

###### Note

AWS AppConfig accepts any type of data, including text formats like JSON or
TOML, or binary formats like protocol buffers or compressed data.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContentType`

A standard MIME type describing the format of the configuration content. For more
information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the configuration.

###### Note

Due to HTTP limitations, this field only supports ASCII characters.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LatestVersionNumber`

An optional locking token used to prevent race conditions from overwriting configuration
updates when creating a new version. To ensure your data is not overwritten when creating
multiple hosted configuration versions in rapid succession, specify the version number of
the latest hosted configuration version.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VersionLabel`

A user-defined label for an AWS AppConfig hosted configuration version.

_Required_: No

_Type_: String

_Pattern_: `^$|.*[^0-9].*`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the version number.

### Fn::GetAtt

`VersionNumber`

The configuration version.

## Examples

- [AWS AppConfig feature flag](#aws-resource-appconfig-hostedconfigurationversion--examples--feature_flag)

- [AWS AppConfig hosted configuration](#aws-resource-appconfig-hostedconfigurationversion--examples--hosted_configuration)

### AWS AppConfig feature flag

The following example creates an AWS AppConfig configuration profile of type
`HostedConfigurationVersion`. The feature flag created by this example
enables cryptocurrency at checkout. AWS AppConfig stores the configuration data
for this profile in the AWS AppConfig hosted configuration store.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Transform": "AWS::LanguageExtensions",
  "Resources": {
    "MySuperCoolApp": {
      "Type": "AWS::AppConfig::Application",
      "Properties": {
        "Name": "MySuperCoolApp"
      }
    },
    "MyFeatureFlags": {
      "Type": "AWS::AppConfig::ConfigurationProfile",
      "Properties": {
        "Name": "MyFeatureFlags",
        "ApplicationId": "MySuperCoolApp",
        "LocationUri": "hosted",
        "Type": "AWS.AppConfig.FeatureFlags"
      }
    },
    "MyFeatureFlagsVersion": {
      "Type": "AWS::AppConfig::HostedConfigurationVersion",
      "Properties": {
        "ApplicationId": "MySuperCoolApp",
        "ConfigurationProfileId": "MyFeatureFlags",
        "ContentType": "application/json",
        "VersionLabel": "v1.0.0",
        "Content": {
          "Fn::ToJsonString": {
            "flags": {
              "allow-cryptocurrency-at-checkout": {
                "attributes": {
                  "allowed-currency": {
                    "constraints": {
                      "elements": {
                        "enum": [
                          "BTC",
                          "ETH",
                          "XRP"
                        ],
                        "type": "string"
                      },
                      "type": "array"
                    }
                  },
                  "bitcoin-discount-percentage": {
                    "constraints": {
                      "maximum": 25,
                      "minimum": 0,
                      "type": "number"
                    }
                  }
                },
                "name": "Allow Cryptocurrency at Checkout"
              }
            },
            "values": {
              "allow-cryptocurrency-at-checkout": {
                "allowed-currency": [
                  "BTC",
                  "ETH"
                ],
                "bitcoin-discount-percentage": 5,
                "enabled": true
              }
            },
            "version": "1"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Transform: 'AWS::LanguageExtensions'
Resources:
  MySuperCoolApp:
    Type: 'AWS::AppConfig::Application'
    Properties:
      Name: MySuperCoolApp

  MyFeatureFlags:
    Type: 'AWS::AppConfig::ConfigurationProfile'
    Properties:
      Name: MyFeatureFlags
      ApplicationId: !Ref MySuperCoolApp
      LocationUri: hosted
      Type: AWS.AppConfig.FeatureFlags

  MyFeatureFlagsVersion:
    Type: 'AWS::AppConfig::HostedConfigurationVersion'
    Properties:
      ApplicationId: !Ref MySuperCoolApp
      ConfigurationProfileId: !Ref MyFeatureFlags
      ContentType: application/json
      VersionLabel: "v1.0.0"
      Content:
        Fn::ToJsonString:
            flags:
              allow-cryptocurrency-at-checkout:
                attributes:
                  allowed-currency:
                    constraints:
                      elements:
                        enum:
                          - BTC
                          - ETH
                          - XRP
                        type: string
                      type: array
                  bitcoin-discount-percentage:
                    constraints:
                      maximum: 25
                      minimum: 0
                      type: number
                name: Allow Cryptocurrency at Checkout
            values:
              allow-cryptocurrency-at-checkout:
                allowed-currency:
                  - BTC
                  - ETH
                bitcoin-discount-percentage: 5
                enabled: true
            version: '1'
```

### AWS AppConfig hosted configuration

The following example creates an AWS AppConfig configuration profile named
`MyTestProfile` for an application called `MyApplication`. AWS AppConfig stores the configuration data for this profile in the AWS AppConfig hosted configuration store.

#### JSON

```json

{
  "Resources": {
    "DependentApplication": {
      "Type": "AWS::AppConfig::Application",
      "Properties": {
        "Name": "MyApplication"
      }
    },
    "DependentConfigurationProfile": {
      "Type": "AWS::AppConfig::ConfigurationProfile",
      "Properties": {
        "ApplicationId": "DependentApplication",
        "Name": "MyTestProfile",
        "LocationUri": "hosted"
      }
    },
    "BasicHostedConfigurationVersion": {
      "Type": "AWS::AppConfig::HostedConfigurationVersion",
      "Properties": {
        "ApplicationId": "DependentApplication",
        "ConfigurationProfileId": "DependentConfigurationProfile",
        "Description": "A sample hosted configuration version",
        "Content": "My hosted configuration content",
        "ContentType": "text/plain"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  DependentApplication:
    Type: AWS::AppConfig::Application
    Properties:
      Name: "MyApplication"
  DependentConfigurationProfile:
    Type: AWS::AppConfig::ConfigurationProfile
    Properties:
      ApplicationId: !Ref DependentApplication
      Name: "MyTestProfile"
      LocationUri: "hosted"
  BasicHostedConfigurationVersion:
    Type: AWS::AppConfig::HostedConfigurationVersion
    Properties:
      ApplicationId: !Ref DependentApplication
      ConfigurationProfileId: !Ref DependentConfigurationProfile
      Description: 'A sample hosted configuration version'
      Content: 'My hosted configuration content'
      ContentType: 'text/plain'
```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Creating a configuration and a configuration profile](https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-creating-configuration-and-profile.html)

- [About the AWS AppConfig hosted configuration store](https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-creating-configuration-and-profile.html#appconfig-creating-configuration-and-profile-about-hosted-store)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
