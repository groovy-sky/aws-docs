This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::ConfigurationProfile

The `AWS::AppConfig::ConfigurationProfile` resource creates a configuration
profile that enables AWS AppConfig to access the configuration source. Valid
configuration sources include AWS Systems Manager (SSM) documents, SSM Parameter Store
parameters, and Amazon S3. A configuration profile includes the following
information.

- The Uri location of the configuration data.

- The AWS Identity and Access Management (IAM) role that provides access to the
configuration data.

- A validator for the configuration data. Available validators include either a JSON
Schema or the Amazon Resource Name (ARN) of an AWS Lambda function.

AWS AppConfig requires that you create resources and deploy a configuration in the
following order:

1. Create an application

2. Create an environment

3. Create a configuration profile

4. Choose a pre-defined deployment strategy or create your own

5. Deploy the configuration

For more information, see [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::ConfigurationProfile",
  "Properties" : {
      "ApplicationId" : String,
      "DeletionProtectionCheck" : String,
      "Description" : String,
      "KmsKeyIdentifier" : String,
      "LocationUri" : String,
      "Name" : String,
      "RetrievalRoleArn" : String,
      "Tags" : [ Tags, ... ],
      "Type" : String,
      "Validators" : [ Validators, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::ConfigurationProfile
Properties:
  ApplicationId: String
  DeletionProtectionCheck: String
  Description: String
  KmsKeyIdentifier: String
  LocationUri: String
  Name: String
  RetrievalRoleArn: String
  Tags:
    - Tags
  Type: String
  Validators:
    - Validators

```

## Properties

`ApplicationId`

The application ID.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9]{4,7}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeletionProtectionCheck`

A parameter to configure deletion protection. Deletion protection prevents a user from
deleting a configuration profile if your application has called either [GetLatestConfiguration](../../../../reference/appconfig/2019-10-09/apireference/api-appconfigdata-getlatestconfiguration.md) or [GetConfiguration](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfiguration.html) for the
configuration profile during the specified interval.

This parameter supports the following values:

- `BYPASS`: Instructs AWS AppConfig to bypass the deletion
protection check and delete a configuration profile even if deletion protection would
have otherwise prevented it.

- `APPLY`: Instructs the deletion protection check to run, even if
deletion protection is disabled at the account level. `APPLY` also forces
the deletion protection check to run against resources created in the past hour,
which are normally excluded from deletion protection checks.

- `ACCOUNT_DEFAULT`: The default setting, which instructs AWS AppConfig to implement the deletion protection value specified in the
`UpdateAccountSettings` API.

_Required_: No

_Type_: String

_Allowed values_: `ACCOUNT_DEFAULT | APPLY | BYPASS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the configuration profile.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when
the resource was created or updated.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationUri`

A URI to locate the configuration. You can specify the following:

- For the AWS AppConfig hosted configuration store and for feature flags,
specify `hosted`.

- For an AWS Systems Manager Parameter Store parameter, specify either the parameter name in
the format `ssm-parameter://<parameter name>` or the ARN.

- For an AWS
CodePipeline pipeline, specify the URI in the following format:
`codepipeline`://<pipeline name>.

- For an AWS Secrets Manager secret, specify the URI in the following format:
`secretsmanager`://<secret name>.

- For an Amazon S3 object, specify the URI in the following format:
`s3://<bucket>/<objectKey> `. Here is an example:
`s3://amzn-s3-demo-bucket/my-app/us-east-1/my-config.json`

- For an SSM document, specify either the document name in the format
`ssm-document://<document name>` or the Amazon Resource Name
(ARN).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A name for the configuration profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetrievalRoleArn`

The ARN of an IAM role with permission to access the configuration at the specified
`LocationUri`.

###### Important

A retrieval role ARN is not required for configurations stored in AWS CodePipeline or the AWS AppConfig hosted configuration store. It is required for all other sources that
store your configuration.

_Required_: No

_Type_: String

_Pattern_: `^((arn):(aws|aws-cn|aws-iso|aws-iso-[a-z]{1}|aws-us-gov|aws-eusc):(iam)::\d{12}:role[/].*)$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata to assign to the configuration profile. Tags help organize and categorize your
AWS AppConfig resources. Each tag consists of a key and an optional value, both of
which you define.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-configurationprofile-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-configurationprofile-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of configurations contained in the profile. AWS AppConfig supports
`feature flags` and `freeform` configurations. We recommend you
create feature flag configurations to enable or disable new features and freeform
configurations to distribute configurations to an application. When calling this API, enter
one of the following values for `Type`:

`AWS.AppConfig.FeatureFlags`

`AWS.Freeform`

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z\.]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Validators`

A list of methods for validating the configuration.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-configurationprofile-validators.html) of [Validators](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-configurationprofile-validators.html)

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the configuration profile ID.

### Fn::GetAtt

`ConfigurationProfileId`

The configuration profile ID.

`KmsKeyArn`

The Amazon Resource Name of the AWS Key Management Service key to encrypt new configuration
data versions in the AWS AppConfig hosted configuration store. This attribute is only
used for `hosted` configuration types. To encrypt data managed in other
configuration stores, see the documentation for how to specify an AWS KMS key
for that particular service.

## Examples

- [AWS AppConfig feature flag](#aws-resource-appconfig-configurationprofile--examples--feature_flag)

- [AWS AppConfig configuration profile example - AWS CodePipeline](#aws-resource-appconfig-configurationprofile--examples--configuration_profile_example_-)

- [AWS AppConfig configuration profile example - Parameter Store](#aws-resource-appconfig-configurationprofile--examples--configuration_profile_example_-_Parameter_Store)

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

### AWS AppConfig configuration profile example - AWS CodePipeline

The following examples creates an AWS AppConfig configuration profile named
MyTestConfigurationProfile. A configuration profile includes source information for
accessing your configuration data. A configuration profile can also include optional
validators to ensure your configuration data is syntactically and semantically correct.
The following configuration profile example uses the specified `LocationUri` to
retrieve configuration data from AWS CodePipeline.

#### JSON

```json

{
  "Resources": {
    "CodePipelineConfigurationProfile": {
      "Type": "AWS::AppConfig::ConfigurationProfile",
      "DependsOn": "MyTestApplication",
      "Properties": {
        "ApplicationId": "MyTestApplication",
        "Name": "MyTestConfigurationProfile",
        "Description": "My test configuration profile",
        "LocationUri": "codepipeline://YourPipelineName",
        "Validators": [
          {
            "Type": "LAMBDA",
            "Content": "MyLambdaValidator"
          }
        ],
        "Tags": [
          {
            "Key": "Env",
            "Value": "test"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  CodePipelineConfigurationProfile:
    Type: AWS::AppConfig::ConfigurationProfile
    Properties:
      ApplicationId: !Ref MyTestApplication
      Name: "MyTestConfigurationProfile"
      Description: "My test configuration profile"
      LocationUri: "codepipeline://YourPipelineName"
      Validators:
        - Type: LAMBDA
          Content: !ImportValue MyLambdaValidator
      Tags:
        - Key: Env
          Value: test
```

### AWS AppConfig configuration profile example - Parameter Store

The following examples creates an AWS AppConfig configuration profile named
MyTestConfigurationProfile. A configuration profile includes source information for
accessing your configuration data. A configuration profile can also include optional
validators to ensure your configuration data is syntactically and semantically correct.
The following configuration profile example uses the specified
`RetrievalRoleArn` and `LocationUri` to retrieve configuration
data from an SSM parameter.

#### JSON

```json

{
  "Type": "AWS::AppConfig::ConfigurationProfile",
  "DependsOn": "MyTestApplication",
  "Properties": {
    "ApplicationId": {
      "Ref": "MyTestApplication"
    },
    "Name": "MyTestConfigurationProfile",
    "Description": "My test configuration profile",
    "RetrievalRoleArn": {
      "Fn::ImportValue": "ConfigurationRetrievalAndMonitoringRole"
    },
    "LocationUri": {
      "Fn::Sub": [
        "ssm-parameter://${ParameterName}",
        {
          "ParameterName": {
            "Fn::ImportValue": "SSMParameter"
          }
        }
      ]
    },
    "Validators": [
      {
        "Type": "LAMBDA",
        "Content": {
          "Fn::ImportValue": "MyLambdaValidator"
        }
      }
    ],
    "Tags": [
      {
        "Key": "Env",
        "Value": "Test"
      }
    ]
  }
}

```

#### YAML

```yaml

Resources:
  BasicConfigurationProfile:
    Type: AWS::AppConfig::ConfigurationProfile
    Properties:
      ApplicationId: !Ref MyTestApplication
      Name: "MyTestConfigurationProfile"
      Description: "My test configuration profile"
      RetrievalRoleArn: !ImportValue ConfigurationRetrievalAndMonitoringRole
      LocationUri:
        Fn::Sub:
          - "ssm-parameter://${ParameterName}"
          - ParameterName: !ImportValue SSMParameter
      Validators:
        - Type: LAMBDA
          Content: !ImportValue MyLambdaValidator
      Tags:
        - Key: Env
          Value: test

```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Creating a configuration and a configuration profile](https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig-creating-configuration-and-profile.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

Tags
