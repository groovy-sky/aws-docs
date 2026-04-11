This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationAggregator

The details about the configuration aggregator, including
information about source accounts, regions, and metadata of the
aggregator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::ConfigurationAggregator",
  "Properties" : {
      "AccountAggregationSources" : [ AccountAggregationSource, ... ],
      "ConfigurationAggregatorName" : String,
      "OrganizationAggregationSource" : OrganizationAggregationSource,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Config::ConfigurationAggregator
Properties:
  AccountAggregationSources:
    - AccountAggregationSource
  ConfigurationAggregatorName: String
  OrganizationAggregationSource:
    OrganizationAggregationSource
  Tags:
    - Tag

```

## Properties

`AccountAggregationSources`

Provides a list of source accounts and regions to be
aggregated.

_Required_: No

_Type_: Array of [AccountAggregationSource](aws-properties-config-configurationaggregator-accountaggregationsource.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationAggregatorName`

The name of the aggregator.

_Required_: No

_Type_: String

_Pattern_: `[\w\-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationAggregationSource`

Provides an organization and list of regions to be
aggregated.

_Required_: No

_Type_: [OrganizationAggregationSource](aws-properties-config-configurationaggregator-organizationaggregationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tag object.

_Required_: No

_Type_: Array of [Tag](aws-properties-config-configurationaggregator-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ConfigurationAggregatorName, such as `myConfigurationAggregator`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`ConfigurationAggregatorArn`

The Amazon Resource Name (ARN) of the aggregator.

## Examples

- [Configuration Aggregator With Multiple Accounts Multiple Regions](#aws-resource-config-configurationaggregator--examples--Configuration_Aggregator_With_Multiple_Accounts_Multiple_Regions)

- [Configuration Aggregator for an Organization](#aws-resource-config-configurationaggregator--examples--Configuration_Aggregator_for_an_Organization)

### Configuration Aggregator With Multiple Accounts Multiple Regions

The following example creates a `ConfigurationAggregator`.

#### JSON

```json

"ConfigurationAggregator": {
    "Type": "AWS::Config::ConfigurationAggregator",
    "Properties": {
      "AccountAggregationSources": [
        {
          "AccountIds": [
            "123456789012",
            "987654321012"
          ],
          "AwsRegions": [
            "us-west-2",
            "us-east-1"
          ],
          "AllAwsRegions": false
        }
      ],
      "ConfigurationAggregatorName": "MyConfigurationAggregator"
    }
  }
```

#### YAML

```yaml

ConfigurationAggregator:
  Type: 'AWS::Config::ConfigurationAggregator'
  Properties:
    AccountAggregationSources:
      - AccountIds:
          - '123456789012'
          - '987654321012'
        AwsRegions:
          - us-west-2
          - us-east-1
        AllAwsRegions: false
    ConfigurationAggregatorName: MyConfigurationAggregator
```

### Configuration Aggregator for an Organization

The following example creates a `ConfigurationAggregator` for an organization.

**Considerations**

- The aggregator account must be the management account or a delegated administrator account in the organization

- AWS Config must be enabled with proper service access in the organization

- The role must have proper permissions to call AWS Organizations APIs

#### JSON

```json

"ConfigurationAggregator": {
    "Type": "AWS::Config::ConfigurationAggregator",
    "Properties": {
        "OrganizationAggregationSource": {
            "RoleArn": { "Fn::GetAtt" : [ "MyRole", "Arn" ] },
            "AwsRegions": [
                "us-west-2",
                "us-east-1"
            ],
            "AllAwsRegions": false
        },
        "ConfigurationAggregatorName": "MyConfigurationAggregator"
    }
}

"MyRole": {
    "Type": "AWS::IAM::Role",
    "Properties": {
        "ManagedPolicyArns": [
            "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations"
        ],
        "Path": "/service-role/",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "config.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }
            ]
        },
        "Policies": [
            {
                "PolicyName": "OrganizationAccess",
                "PolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Action": [
                                "organizations:DescribeOrganization",
                                "organizations:ListAWSServiceAccessForOrganization",
                                "organizations:ListAccounts"
                            ],
                            "Resource": "*"
                        }
                    ]
                }
            }
        ]
    }
}
```

#### YAML

```yaml

ConfigurationAggregator:
    Type: 'AWS::Config::ConfigurationAggregator'
    Properties:
        OrganizationAggregationSource:
            RoleArn: !GetAtt MyRole.Arn
            AwsRegions:
                - us-west-2
                - us-east-1
            AllAwsRegions: false
        ConfigurationAggregatorName: MyConfigurationAggregator

MyRole:
    Type: AWS::IAM::Role
    Properties:
        ManagedPolicyArns:
            - arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations
        Path: "/service-role/"
        AssumeRolePolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Principal:
                    Service:
                        - config.amazonaws.com
                Action:
                    - 'sts:AssumeRole'
        Policies:
            - PolicyName: OrganizationAccess
              PolicyDocument:
                Version: "2012-10-17"
                Statement:
                    - Effect: Allow
                      Action:
                        - organizations:DescribeOrganization
                        - organizations:ListAWSServiceAccessForOrganization
                        - organizations:ListAccounts
                      Resource: "*"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceDetail

AccountAggregationSource

All content copied from https://docs.aws.amazon.com/.
