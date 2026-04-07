This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain

The AWS::OpenSearchService::Domain resource creates an Amazon OpenSearch Service
domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OpenSearchService::Domain",
  "Properties" : {
      "AccessPolicies" : Json,
      "AdvancedOptions" : {Key: Value, ...},
      "AdvancedSecurityOptions" : AdvancedSecurityOptionsInput,
      "AIMLOptions" : AIMLOptions,
      "ClusterConfig" : ClusterConfig,
      "CognitoOptions" : CognitoOptions,
      "DeploymentStrategyOptions" : DeploymentStrategyOptions,
      "DomainEndpointOptions" : DomainEndpointOptions,
      "DomainName" : String,
      "EBSOptions" : EBSOptions,
      "EncryptionAtRestOptions" : EncryptionAtRestOptions,
      "EngineVersion" : String,
      "IdentityCenterOptions" : IdentityCenterOptions,
      "IPAddressType" : String,
      "LogPublishingOptions" : {Key: Value, ...},
      "NodeToNodeEncryptionOptions" : NodeToNodeEncryptionOptions,
      "OffPeakWindowOptions" : OffPeakWindowOptions,
      "SkipShardMigrationWait" : Boolean,
      "SnapshotOptions" : SnapshotOptions,
      "SoftwareUpdateOptions" : SoftwareUpdateOptions,
      "Tags" : [ Tag, ... ],
      "VPCOptions" : VPCOptions
    }
}

```

### YAML

```yaml

Type: AWS::OpenSearchService::Domain
Properties:
  AccessPolicies: Json
  AdvancedOptions:
    Key: Value
  AdvancedSecurityOptions:
    AdvancedSecurityOptionsInput
  AIMLOptions:
    AIMLOptions
  ClusterConfig:
    ClusterConfig
  CognitoOptions:
    CognitoOptions
  DeploymentStrategyOptions:
    DeploymentStrategyOptions
  DomainEndpointOptions:
    DomainEndpointOptions
  DomainName: String
  EBSOptions:
    EBSOptions
  EncryptionAtRestOptions:
    EncryptionAtRestOptions
  EngineVersion: String
  IdentityCenterOptions:
    IdentityCenterOptions
  IPAddressType: String
  LogPublishingOptions:
    Key: Value
  NodeToNodeEncryptionOptions:
    NodeToNodeEncryptionOptions
  OffPeakWindowOptions:
    OffPeakWindowOptions
  SkipShardMigrationWait: Boolean
  SnapshotOptions:
    SnapshotOptions
  SoftwareUpdateOptions:
    SoftwareUpdateOptions
  Tags:
    - Tag
  VPCOptions:
    VPCOptions

```

## Properties

`AccessPolicies`

An AWS Identity and Access Management (IAM) policy document that specifies who can
access the OpenSearch Service domain and their permissions. For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the _Amazon OpenSearch Service Developer_
_Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedOptions`

Additional options to specify for the OpenSearch Service domain. For more information, see
[AdvancedOptions](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_CreateDomain.html#API_CreateDomain_RequestBody) in the OpenSearch Service API reference.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedSecurityOptions`

Specifies options for fine-grained access control and SAML authentication.

If you specify advanced security options, you must also enable node-to-node encryption ( [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.html))
and encryption at rest ( [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-encryptionatrestoptions.html)).
You must also enable `EnforceHTTPS` within [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-domainendpointoptions.html),
which requires HTTPS for all traffic to the domain.

_Required_: No

_Type_: [AdvancedSecurityOptionsInput](aws-properties-opensearchservice-domain-advancedsecurityoptionsinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AIMLOptions`

Container for parameters required to enable all machine learning features.

_Required_: No

_Type_: [AIMLOptions](aws-properties-opensearchservice-domain-aimloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterConfig`

Container for the cluster configuration of a domain.

_Required_: No

_Type_: [ClusterConfig](aws-properties-opensearchservice-domain-clusterconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoOptions`

Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch
Dashboards.

_Required_: No

_Type_: [CognitoOptions](aws-properties-opensearchservice-domain-cognitooptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentStrategyOptions`

The current status of the domain's deployment strategy options.

_Required_: No

_Type_: [DeploymentStrategyOptions](aws-properties-opensearchservice-domain-deploymentstrategyoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainEndpointOptions`

Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.

_Required_: No

_Type_: [DomainEndpointOptions](aws-properties-opensearchservice-domain-domainendpointoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

A name for the OpenSearch Service domain. The name must have a minimum length of 3 and a maximum length of 28. If you don't specify a name, CloudFormation generates a unique
physical ID and uses that ID for the domain name. For more information, see [Name\
Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

Required when creating a new domain.

###### Important

If you specify a name, you can't perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must replace
the resource, specify a new name.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EBSOptions`

The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to
data nodes in the OpenSearch Service domain. For more information, see [EBS volume size limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#ebsresource) in the _Amazon OpenSearch Service Developer_
_Guide_.

_Required_: No

_Type_: [EBSOptions](aws-properties-opensearchservice-domain-ebsoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionAtRestOptions`

Whether the domain should encrypt data at rest, and if so, the AWS KMS key to
use. See [Encryption of data at rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html).

If no encryption at rest options were initially specified in the template, updating this property by adding it causes no interruption. However, if you change this property after it's already been set within a template,
the domain is deleted and recreated in order to modify the property.

_Required_: No

_Type_: [EncryptionAtRestOptions](aws-properties-opensearchservice-domain-encryptionatrestoptions.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineVersion`

The version of OpenSearch to use. The value must be in the format
`OpenSearch_X.Y` or `Elasticsearch_X.Y`. If not specified, the latest version of OpenSearch is used. For
information about the versions that OpenSearch Service supports, see [Supported\
versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the _Amazon OpenSearch Service_
_Developer Guide_.

If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true`, you can update
`EngineVersion` without interruption. When `EnableVersionUpgrade` is
set to `false`, or is not specified, updating `EngineVersion` results in
[replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: String

_Pattern_: `^Elasticsearch_[0-9]{1}\.[0-9]{1,2}$|^OpenSearch_[0-9]{1,2}\.[0-9]{1,2}$`

_Minimum_: `14`

_Maximum_: `18`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityCenterOptions`

Configuration options for controlling IAM Identity Center integration within a
domain.

_Required_: No

_Type_: [IdentityCenterOptions](aws-properties-opensearchservice-domain-identitycenteroptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IPAddressType`

Choose either dual stack or IPv4 as your IP address type. Dual stack allows you to
share domain resources across IPv4 and IPv6 address types, and is the recommended
option. If you set your IP address type to dual stack, you can't change your address
type later.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPublishingOptions`

An object with one or more of the following keys: `SEARCH_SLOW_LOGS`,
`ES_APPLICATION_LOGS`, `INDEX_SLOW_LOGS`, `AUDIT_LOGS`,
depending on the types of logs you want to publish. Each key needs a valid
`LogPublishingOption` value. For the full syntax, see the [examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--examples).

_Required_: No

_Type_: Object of [LogPublishingOption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-opensearchservice-domain-logpublishingoption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeToNodeEncryptionOptions`

Specifies whether node-to-node encryption is enabled. See [Node-to-node encryption for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html).

_Required_: No

_Type_: [NodeToNodeEncryptionOptions](aws-properties-opensearchservice-domain-nodetonodeencryptionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OffPeakWindowOptions`

Options for a domain's off-peak window, during which OpenSearch Service can perform mandatory configuration changes on the domain.

_Required_: No

_Type_: [OffPeakWindowOptions](aws-properties-opensearchservice-domain-offpeakwindowoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkipShardMigrationWait`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotOptions`

**DEPRECATED**. The automated snapshot configuration for the
OpenSearch Service domain indexes.

_Required_: No

_Type_: [SnapshotOptions](aws-properties-opensearchservice-domain-snapshotoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftwareUpdateOptions`

Service software update options for the domain.

_Required_: No

_Type_: [SoftwareUpdateOptions](aws-properties-opensearchservice-domain-softwareupdateoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service
domain.

_Required_: No

_Type_: Array of [Tag](aws-properties-opensearchservice-domain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCOptions`

The virtual private cloud (VPC) configuration for the OpenSearch Service domain. For more
information, see [Launching your Amazon OpenSearch Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the _Amazon OpenSearch Service Developer_
_Guide_.

If you remove this entity altogether, along with its associated properties, it causes a replacement. You might encounter this scenario if you're updating your security configuration from a VPC to a public endpoint.

_Required_: No

_Type_: [VPCOptions](aws-properties-opensearchservice-domain-vpcoptions.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the resource name, such as `mystack-abc1d2efg3h4.` For more
information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

GetAtt returns a value for a specified attribute of this type. For more information, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`AdvancedSecurityOptions.AnonymousAuthDisableDate`

Date and time when the migration period will be disabled. Only necessary when [enabling\
fine-grained access control on an existing domain](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/fgac.html#fgac-enabling-existing).

`Arn`

The Amazon Resource Name (ARN) of the CloudFormation stack.

`DomainArn`

The Amazon Resource Name (ARN) of the domain. See [Identifiers for IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in
_Using AWS Identity and Access Management_ for
more information.

`DomainEndpoint`

The domain-specific endpoint used for requests to the OpenSearch APIs, such as
`search-mystack-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com`.

`DomainEndpointV2`

If `IPAddressType` to set to `dualstack`, a version 2 domain endpoint
is provisioned. This endpoint functions like a normal endpoint, except that it works with both
IPv4 and IPv6 IP addresses. Normal endpoints work only with IPv4 IP addresses.

`Id`

The resource ID. For example, `123456789012/my-domain`.

`IdentityCenterOptions.IdentityCenterApplicationARN`

The Amazon Resource Name (ARN) of the domain. See [Identifiers for IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/index.html) in
_Using AWS Identity and Access Management_ for
more information.

`IdentityCenterOptions.IdentityStoreId`

Property description not available.

## Remarks

_Migrating stacks from Elasticsearch to OpenSearch_

###### Important

You can't directly update CloudFormation templates to use the
`AWS::OpenSearchService::Domain` resource in place of
`AWS::Elasticsearch::Domain`, otherwise the corresponding domain will be
deleted along with all of its data.

Perform the following steps to migrate an Elasticsearch domain to an OpenSearch domain
if the domain is defined within CloudFormation.

**Step 1: Prepare your existing stack for deprecation**

Make a copy of your original CloudFormation template, which contains the Elasticsearch
domain resource, for use in step 3. Then add the following attributes to the Elasticsearch
domain resource at the same level as `Type` and `Properties`.

`DeletionPolicy: Retain`

`UpdateReplacePolicy: Retain`

These settings ensure that CloudFormation doesn't delete or modify the corresponding
domain when you delete this resource from your stack. If you have other custom resources
defined in the stack that aren't critically important during the short migration period, you
can delete them from the template and they'll be recreated when you create the new
stack.

**Step 2: Upgrade your domain to OpenSearch**

After you add the two policy attributes to your template, upgrade your domain to an
OpenSearch version using the normal upgrade process. For instructions, see [Starting an upgrade](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/version-migration.html#starting-upgrades). Make sure to take a snapshot of your domain before upgrading
it to prevent accidental loss of data.

**Step 3: Create a new CloudFormation template**

While you wait for the upgrade to complete, prepare your new OpenSearch template. Using
the copy of your original template that you made in step 1, make the following
changes:

- Change the domain resource type from `AWS::Elasticsearch::Domain` to
`AWS::OpenSearchService::Domain`.

- Add the `DeletionPolicy` and `UpdateReplacePolicy` attributes
to the resource, as you did in step 1.

- Change `ElasticsearchVersion` to `EngineVersion` and set its value to
`OpenSearch_1.0` (or whichever version of OpenSearch you want to upgrade
to).

- If your resource contains `ElasticsearchClusterConfig`, change it to
`ClusterConfig`.

- Change the suffixes of all instance types from `.elasticsearch` to
`.search`.

- If there are any `Fn::GetAtt` or `!GetAtt` references to your
domain ARN, change them to `!GetAtt MyDomain.Arn`.

- Comment out any resources not currently within the stack (most likely everything
except `AWS::OpenSearchService::Domain`).

See the next section for examples that demonstrate the new format.

**Step 4: Import the OpenSearch stack**

Once your domain upgrade finishes, you can import the new stack. Within CloudFormation,
choose **Create stack** and **With existing**
**resources (import resources)**, then upload the template you created in the
previous step.

CloudFormation prompts you for the name (identifier value) of the existing domain. Copy
the domain name directly from the OpenSearch console. Give the stack a name that's different
from the current one, then choose **Import resources**.

After the stack is created, uncomment any related resources from the stack and update it
to ensure they're recreated. You can remove the `DeletionPolicy` and
`UpdateReplacePolicy` attributes if you want, but they can help prevent
accidental deletions in the future.

**Step 5: Delete the Elasticsearch stack**

Now that your new stack is created, delete the old stack which contains the legacy
Elasticsearch resource.

###### Important

Before deleting the old stack, make absolutely sure that the template contains the
`DeletionPolicy: Retain` attribute.

\*The above steps were partially derived from this [blog\
post](https://onecloudplease.com/blog/migrating-to-opensearch-with-cloudformation).

## Examples

- [Create an OpenSearch Service domain that contains two data nodes and three master nodes](#aws-resource-opensearchservice-domain--examples--Create_an_OpenSearch_Service_domain_that_contains_two_data_nodes_and_three_master_nodes)

- [Create a domain with VPC options](#aws-resource-opensearchservice-domain--examples--Create_a_domain_with_VPC_options)

- [Create a domain with fine-grained access control](#aws-resource-opensearchservice-domain--examples--Create_a_domain_with_fine-grained_access_control)

### Create an OpenSearch Service domain that contains two data nodes and three master nodes

The following example creates an OpenSearch Service domain running OpenSearch 1.0 that
contains two data nodes and three dedicated master nodes. The domain has 40 GiB of storage
and enables log publishing for application logs, search slow logs, and index slow logs.
The access policy permits the root user for the AWS account to make all
HTTP requests to the domain, such as indexing documents or searching indexes.

#### JSON

```json

"OpenSearchServiceDomain": {
   "Type":"AWS::OpenSearchService::Domain",
   "Properties": {
      "DomainName": "test",
      "EngineVersion": "OpenSearch_1.0",
      "ClusterConfig": {
         "DedicatedMasterEnabled": true,
         "InstanceCount": "2",
         "ZoneAwarenessEnabled": true,
         "InstanceType": "m3.medium.search",
         "DedicatedMasterType": "m3.medium.search",
         "DedicatedMasterCount": "3"
      },
      "EBSOptions":{
         "EBSEnabled": true,
         "Iops": "0",
         "VolumeSize": "20",
         "VolumeType": "gp2"
      },
      "AccessPolicies": {
         "Version":"2012-10-17",
         "Statement":[
            {
               "Effect": "Allow",
               "Principal": {
                  "AWS": "arn:aws:iam::123456789012:user/opensearch-user"
               },
               "Action":"es:*",
               "Resource": "arn:aws:es:us-east-1:123456789012:domain/test/*"
            }
         ]
      },
      "LogPublishingOptions": {
         "ES_APPLICATION_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-application-logs",
           "Enabled": true
         },
         "SEARCH_SLOW_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-slow-logs",
           "Enabled": true
         },
         "INDEX_SLOW_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-index-slow-logs",
           "Enabled": true
         }
      },
      "AdvancedOptions": {
         "rest.action.multi.allow_explicit_index": "true",
         "override_main_response_version": "true"
      }
   }
}
```

#### YAML

```yaml

OpenSearchServiceDomain:
  Type: AWS::OpenSearchService::Domain
  Properties:
    DomainName: 'test'
    EngineVersion: 'OpenSearch_1.0'
    ClusterConfig:
      DedicatedMasterEnabled: true
      InstanceCount: '2'
      ZoneAwarenessEnabled: true
      InstanceType: 'm3.medium.search'
      DedicatedMasterType: 'm3.medium.search'
      DedicatedMasterCount: '3'
    EBSOptions:
      EBSEnabled: true
      Iops: '0'
      VolumeSize: '20'
      VolumeType: 'gp2'
    AccessPolicies:
      Version: '2012-10-17'
      Statement:
        -
          Effect: 'Allow'
          Principal:
            AWS: 'arn:aws:iam::123456789012:user/opensearch-user'
          Action: 'es:*'
          Resource: 'arn:aws:es:us-east-1:846973539254:domain/test/*'
    LogPublishingOptions:
      ES_APPLICATION_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-application-logs'
          Enabled: true
      SEARCH_SLOW_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-slow-logs'
          Enabled: true
      INDEX_SLOW_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearch/domains/opensearch-index-slow-logs'
          Enabled: true
    AdvancedOptions:
      rest.action.multi.allow_explicit_index: 'true'
      override_main_response_version: 'true'
```

### Create a domain with VPC options

The following example creates a domain with VPC options.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "OpenSearchServiceDomain resource",
  "Parameters": {
    "DomainName": {
      "Description": "User-defined OpenSearch domain name",
      "Type": "String"
    },
    "EngineVersion": {
      "Description": "User-defined OpenSearch version",
      "Type": "String"
    },
    "InstanceType": {
      "Type": "String"
    },
    "AvailabilityZone": {
      "Type": "String"
    },
    "CidrBlock": {
      "Type": "String"
    },
    "GroupDescription": {
      "Type": "String"
    },
    "SGName": {
      "Type": "String"
    }
  },
  "Resources": {
    "OpenSearchServiceDomain": {
      "Type": "AWS::OpenSearchService::Domain",
      "Properties": {
        "DomainName": {
          "Ref": "DomainName"
        },
        "EngineVersion": {
          "Ref": "EngineVersion"
        },
        "ClusterConfig": {
          "InstanceCount": "1",
          "InstanceType": {
            "Ref": "InstanceType"
          }
        },
        "EBSOptions": {
          "EBSEnabled": true,
          "Iops": "0",
          "VolumeSize": "10",
          "VolumeType": "standard"
        },
        "AccessPolicies": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Deny",
              "Principal": {
                "AWS": "*"
              },
              "Action": "es:*",
              "Resource": "*"
            }
          ]
        },
        "AdvancedOptions": {
          "rest.action.multi.allow_explicit_index": "true",
          "override_main_response_version": "true"
        },
        "Tags": [
          {
            "Key": "foo",
            "Value": "bar"
          }
        ],
        "VPCOptions": {
          "SubnetIds": [
            {
              "Ref": "subnet"
            }
          ],
          "SecurityGroupIds": [
            {
              "Ref": "mySecurityGroup"
            }
          ]
        }
      }
    },
    "vpc": {
      "Type": "AWS::EC2::VPC",
      "Properties": {
        "CidrBlock": "10.0.0.0/16"
      }
    },
    "subnet": {
      "Type": "AWS::EC2::Subnet",
      "Properties": {
        "VpcId": {
          "Ref": "vpc"
        },
        "CidrBlock": {
          "Ref": "CidrBlock"
        },
        "AvailabilityZone": {
          "Ref": "AvailabilityZone"
        }
      }
    },
    "mySecurityGroup": {
      "Type": "AWS::EC2::SecurityGroup",
      "Properties": {
        "GroupDescription": {
          "Ref": "GroupDescription"
        },
        "VpcId": {
          "Ref": "vpc"
        },
        "GroupName": {
          "Ref": "SGName"
        },
        "SecurityGroupIngress": [
          {
            "FromPort": 443,
            "IpProtocol": "tcp",
            "ToPort": 443,
            "CidrIp": "0.0.0.0/0"
          }
        ]
      }
    }
  },
  "Outputs": {
    "Arn": {
      "Value": {
        "Fn::GetAtt": [
          "OpenSearchServiceDomain",
          "Arn"
        ]
      }
    },
    "DomainEndpoint": {
      "Value": {
        "Fn::GetAtt": [
          "OpenSearchServiceDomain",
          "DomainEndpoint"
        ]
      }
    },
    "SecurityGroupId": {
      "Value": {
        "Ref": "mySecurityGroup"
      }
    },
    "SubnetId": {
      "Value": {
        "Ref": "subnet"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: OpenSearchServiceDomain resource
Parameters:
  DomainName:
    Description: User-defined OpenSearch domain name
    Type: String
  EngineVersion:
    Description: User-defined OpenSearch version
    Type: String
  InstanceType:
    Type: String
  AvailabilityZone:
    Type: String
  CidrBlock:
    Type: String
  GroupDescription:
    Type: String
  SGName:
    Type: String
Resources:
  OpenSearchServiceDomain:
    Type: 'AWS::OpenSearchService::Domain'
    Properties:
      DomainName:
        Ref: DomainName
      EngineVersion:
        Ref: EngineVersion
      ClusterConfig:
        InstanceCount: '1'
        InstanceType:
          Ref: InstanceType
      EBSOptions:
        EBSEnabled: true
        Iops: '0'
        VolumeSize: '10'
        VolumeType: 'standard'
      AccessPolicies:
        Version: '2012-10-17'
        Statement:
          - Effect: Deny
            Principal:
              AWS: '*'
            Action: 'es:*'
            Resource: '*'
      AdvancedOptions:
        rest.action.multi.allow_explicit_index: 'true'
        override_main_response_version: 'true'
      Tags:
        - Key: foo
          Value: bar
      VPCOptions:
        SubnetIds:
          - Ref: subnet
        SecurityGroupIds:
          - Ref: mySecurityGroup
  vpc:
    Type: 'AWS::EC2::VPC'
    Properties:
      CidrBlock: 10.0.0.0/16
  subnet:
    Type: 'AWS::EC2::Subnet'
    Properties:
      VpcId:
        Ref: vpc
      CidrBlock:
        Ref: CidrBlock
      AvailabilityZone:
        Ref: AvailabilityZone
  mySecurityGroup:
    Type: 'AWS::EC2::SecurityGroup'
    Properties:
      GroupDescription:
        Ref: GroupDescription
      VpcId:
        Ref: vpc
      GroupName:
        Ref: SGName
      SecurityGroupIngress:
        - FromPort: 443
          IpProtocol: tcp
          ToPort: 443
          CidrIp: 0.0.0.0/0
Outputs:
  Arn:
    Value:
      'Fn::GetAtt':
        - OpenSearchServiceDomain
        - Arn
  DomainEndpoint:
    Value:
      'Fn::GetAtt':
        - OpenSearchServiceDomain
        - DomainEndpoint
  SecurityGroupId:
    Value:
      Ref: mySecurityGroup
  SubnetId:
    Value:
      Ref: subnet
```

### Create a domain with fine-grained access control

The following example creates a domain with fine-grained access control.

#### JSON

```json

{
  "OpenSearchServiceDomain": {
    "Type": "AWS::OpenSearchService::Domain",
    "Properties": {
      "DomainName": "my-domain-logs",
      "EngineVersion": "OpenSearch_1.0",
      "ClusterConfig": {
        "InstanceCount": 2,
        "InstanceType": "r6g.xlarge.search",
        "DedicatedMasterEnabled": true,
        "DedicatedMasterCount": 3,
        "DedicatedMasterType": "r6g.xlarge.search"
      },
      "EBSOptions": {
        "EBSEnabled": true,
        "VolumeSize": 10,
        "VolumeType": "gp2"
      },
      "AccessPolicies": {
        "Version": "2012-10-17",
        "Statement": {
          "Effect": "Allow",
          "Principal": {
            "AWS": "arn:aws:iam::478253424788:role/Admin"
          },
          "Action": "es:*",
          "Resource": "arn:aws:es:us-east-1:478253424788:domain/my-domain-logs/*"
        }
      }
    },
    "AdvancedSecurityOptions": {
      "Enabled": true,
      "InternalUserDatabaseEnabled": true,
      "MasterUserOptions": {
        "MasterUserName": "<username>",
        "MasterUserPassword": "<password>"
      }
    }
  }
}
```

#### YAML

```yaml

OpenSearchServiceDomain:
  Type: 'AWS::OpenSearchService::Domain'
  Properties:
    DomainName: my-domain-logs
    EngineVersion: OpenSearch_1.0
    ClusterConfig:
      InstanceCount: 2
      InstanceType: r6g.xlarge.search
      DedicatedMasterEnabled: true
      DedicatedMasterCount: 3
      DedicatedMasterType: r6g.xlarge.search
    EBSOptions:
      EBSEnabled: true
      VolumeSize: 10
      VolumeType: gp2
    AccessPolicies:
      Version: '2012-10-17'
      Statement:
        Effect: Allow
        Principal:
          AWS: 'arn:aws:iam::478253424788:role/Admin'
        Action: 'es:*'
        Resource: 'arn:aws:es:us-east-1:478253424788:domain/my-domain-logs/*'
  AdvancedSecurityOptions:
    Enabled: true
    InternalUserDatabaseEnabled: true
    MasterUserOptions:
      MasterUserName: <username>
      MasterUserPassword: <password>
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AdvancedSecurityOptionsInput
