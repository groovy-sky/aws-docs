This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Elasticsearch::Domain

The AWS::Elasticsearch::Domain resource creates an Amazon OpenSearch Service
domain.

###### Important

The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource
and options are still supported, we recommend modifying your existing Cloudformation
templates to use the new OpenSearch Service resource, which supports both OpenSearch and
legacy Elasticsearch. For instructions to upgrade domains defined within CloudFormation from
Elasticsearch to OpenSearch, see [Remarks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html#aws-resource-opensearchservice-domain--remarks).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Elasticsearch::Domain",
  "Properties" : {
      "AccessPolicies" : Json,
      "AdvancedOptions" : {Key: Value, ...},
      "AdvancedSecurityOptions" : AdvancedSecurityOptionsInput,
      "CognitoOptions" : CognitoOptions,
      "DomainEndpointOptions" : DomainEndpointOptions,
      "DomainName" : String,
      "EBSOptions" : EBSOptions,
      "ElasticsearchClusterConfig" : ElasticsearchClusterConfig,
      "ElasticsearchVersion" : String,
      "EncryptionAtRestOptions" : EncryptionAtRestOptions,
      "LogPublishingOptions" : {Key: Value, ...},
      "NodeToNodeEncryptionOptions" : NodeToNodeEncryptionOptions,
      "SnapshotOptions" : SnapshotOptions,
      "Tags" : [ Tag, ... ],
      "VPCOptions" : VPCOptions
    }
}

```

### YAML

```yaml

Type: AWS::Elasticsearch::Domain
Properties:
  AccessPolicies: Json
  AdvancedOptions:
    Key: Value
  AdvancedSecurityOptions:
    AdvancedSecurityOptionsInput
  CognitoOptions:
    CognitoOptions
  DomainEndpointOptions:
    DomainEndpointOptions
  DomainName: String
  EBSOptions:
    EBSOptions
  ElasticsearchClusterConfig:
    ElasticsearchClusterConfig
  ElasticsearchVersion: String
  EncryptionAtRestOptions:
    EncryptionAtRestOptions
  LogPublishingOptions:
    Key: Value
  NodeToNodeEncryptionOptions:
    NodeToNodeEncryptionOptions
  SnapshotOptions:
    SnapshotOptions
  Tags:
    - Tag
  VPCOptions:
    VPCOptions

```

## Properties

`AccessPolicies`

An AWS Identity and Access Management (IAM) policy document that specifies who can
access the OpenSearch Service domain and their permissions. For more information, see [Configuring access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ac.html#ac-creating) in the _Amazon OpenSearch Service Developer_
_Guid_ e.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedOptions`

Additional options to specify for the OpenSearch Service domain. For more information, see [Advanced cluster parameters](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options) in the _Amazon OpenSearch Service_
_Developer Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdvancedSecurityOptions`

Specifies options for fine-grained access control.

_Required_: No

_Type_: [AdvancedSecurityOptionsInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-advancedsecurityoptionsinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoOptions`

Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.

_Required_: No

_Type_: [CognitoOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-cognitooptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainEndpointOptions`

Specifies additional options for the domain endpoint, such as whether to require HTTPS for all traffic or whether to use a custom endpoint rather than the default endpoint.

_Required_: No

_Type_: [DomainEndpointOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-domainendpointoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

A name for the OpenSearch Service domain. For valid values, see the [DomainName](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/configuration-api.html#configuration-api-datatypes-domainname) data type in the _Amazon OpenSearch Service Developer_
_Guide_. If you don't specify a name, CloudFormation generates a unique
physical ID and uses that ID for the domain name. For more information, see [Name\
Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).

###### Important

If you specify a name, you cannot perform updates that require replacement of this
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

_Type_: [EBSOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-ebsoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElasticsearchClusterConfig`

ElasticsearchClusterConfig is a property of the AWS::Elasticsearch::Domain resource that
configures the cluster of an Amazon OpenSearch Service domain.

_Required_: No

_Type_: [ElasticsearchClusterConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElasticsearchVersion`

The version of Elasticsearch to use, such as 2.3. If not specified, 1.5 is used as the
default. For information about the versions that OpenSearch Service supports, see [Supported\
versions of OpenSearch and Elasticsearch](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/what-is.html#choosing-version) in the _Amazon OpenSearch Service_
_Developer Guide_.

If you set the [EnableVersionUpgrade](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-upgradeopensearchdomain) update policy to `true`, you can update
`ElasticsearchVersion` without interruption. When
`EnableVersionUpgrade` is set to `false`, or is not specified,
updating `ElasticsearchVersion` results in [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionAtRestOptions`

Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service
key to use. See [Encryption of data at\
rest for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/encryption-at-rest.html).

_Required_: No

_Type_: [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-encryptionatrestoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPublishingOptions`

An object with one or more of the following keys: `SEARCH_SLOW_LOGS`,
`ES_APPLICATION_LOGS`, `INDEX_SLOW_LOGS`, `AUDIT_LOGS`,
depending on the types of logs you want to publish. Each key needs a valid
`LogPublishingOption` value.

_Required_: No

_Type_: Object of [LogPublishingOption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-logpublishingoption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeToNodeEncryptionOptions`

Specifies whether node-to-node encryption is enabled. See [Node-to-node encryption for Amazon\
OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ntn.html).

_Required_: No

_Type_: [NodeToNodeEncryptionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-nodetonodeencryptionoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotOptions`

**DEPRECATED**. The automated snapshot configuration for the
OpenSearch Service domain indices.

_Required_: No

_Type_: [SnapshotOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-snapshotoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key–value pairs) to associate with the OpenSearch Service domain.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCOptions`

The virtual private cloud (VPC) configuration for the OpenSearch Service domain. For more
information, see [Launching your Amazon OpenSearch\
Service domains within a VPC](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html) in the _Amazon OpenSearch Service Developer_
_Guide_.

_Required_: No

_Type_: [VPCOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticsearch-domain-vpcoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the resource name, such as `mystack-elasticsea-abc1d2efg3h4.` For more
information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

Fn::GetAtt returns a value for a specified attribute of this type. For more information,
see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`Arn`

The Amazon Resource Name (ARN) of the domain, such as
`arn:aws:es:us-west-2:123456789012:domain/mystack-elasti-1ab2cdefghij`. This
returned value is the same as the one returned by
`AWS::Elasticsearch::Domain.DomainArn`.

`DomainArn`

The Amazon Resource Name (ARN) of the domain, such as
`arn:aws:es:us-west-2:123456789012:domain/mystack-elasti-1ab2cdefghij`. This
returned value is the same as the one returned by
`AWS::Elasticsearch::Domain.Arn`.

`DomainEndpoint`

The domain-specific endpoint that's used for requests to the OpenSearch APIs, such as
`search-mystack-elasti-1ab2cdefghij-ab1c2deckoyb3hofw7wpqa3cm.us-west-1.es.amazonaws.com`.

## Remarks

_Migrating stacks from Elasticsearch to OpenSearch_

###### Important

You can't directly update a CloudFormation templates to use the
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

- If your template includes `ColdStorageOptions`, remove it, as it's not
currently supported for OpenSearch resources.

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

- [Create an OpenSearch Service domain that contains two data nodes and three master nodes](#aws-resource-elasticsearch-domain--examples--Create_an_OpenSearch_Service_domain_that_contains_two_data_nodes_and_three_master_nodes)

- [Create a domain with VPC options](#aws-resource-elasticsearch-domain--examples--Create_a_domain_with_VPC_options)

### Create an OpenSearch Service domain that contains two data nodes and three master nodes

The following example creates an OpenSearch Service domain running Elasticsearch 7.10 that
contains two data nodes and three dedicated master nodes. The domain has 40 GiB of storage
and enables log publishing for application logs, search slow logs, and index slow logs.
The access policy permits the root user for the AWS account to make all HTTP requests to
the domain, such as indexing documents or searching indices.

#### JSON

```json

"ElasticsearchDomain": {
   "Type":"AWS::Elasticsearch::Domain",
   "Properties": {
      "DomainName": "test",
      "ElasticsearchVersion": "7.10",
      "ElasticsearchClusterConfig": {
         "DedicatedMasterEnabled": true,
         "InstanceCount": "2",
         "ZoneAwarenessEnabled": true,
         "InstanceType": "m3.medium.elasticsearch",
         "DedicatedMasterType": "m3.medium.elasticsearch",
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
                  "AWS": "arn:aws:iam::123456789012:user/es-user"
               },
               "Action":"es:*",
               "Resource": "arn:aws:es:us-east-1:123456789012:domain/test/*"
            }
         ]
      },
      "LogPublishingOptions": {
         "ES_APPLICATION_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-application-logs",
           "Enabled": true
         },
         "SEARCH_SLOW_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-slow-logs",
           "Enabled": true
         },
         "INDEX_SLOW_LOGS": {
           "CloudWatchLogsLogGroupArn": "arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-index-slow-logs",
           "Enabled": true
         }
      },
      "AdvancedOptions": {
         "rest.action.multi.allow_explicit_index": true
      }
   }
}
```

#### YAML

```yaml

ElasticsearchDomain:
  Type: AWS::Elasticsearch::Domain
  Properties:
    DomainName: 'test'
    ElasticsearchVersion: '7.10'
    ElasticsearchClusterConfig:
      DedicatedMasterEnabled: true
      InstanceCount: '2'
      ZoneAwarenessEnabled: true
      InstanceType: 'm3.medium.elasticsearch'
      DedicatedMasterType: 'm3.medium.elasticsearch'
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
            AWS: 'arn:aws:iam::123456789012:user/es-user'
          Action: 'es:*'
          Resource: 'arn:aws:es:us-east-1:846973539254:domain/test/*'
    LogPublishingOptions:
      ES_APPLICATION_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-application-logs'
          Enabled: true
      SEARCH_SLOW_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-slow-logs'
          Enabled: true
      INDEX_SLOW_LOGS:
          CloudWatchLogsLogGroupArn: 'arn:aws:logs:us-east-1:123456789012:log-group:/aws/opensearchservice/domains/es-index-slow-logs'
          Enabled: true
    AdvancedOptions:
      rest.action.multi.allow_explicit_index: true
```

### Create a domain with VPC options

The following example creates a domain with VPC options.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "ElasticsearchDomain resource",
  "Parameters": {
    "DomainName": {
      "Description": "User defined Elasticsearch Domain name",
      "Type": "String"
    },
    "ElasticsearchVersion": {
      "Description": "User defined Elasticsearch Version",
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
    "ElasticsearchDomain": {
      "Type": "AWS::Elasticsearch::Domain",
      "Properties": {
        "DomainName": {
          "Ref": "DomainName"
        },
        "ElasticsearchVersion": {
          "Ref": "ElasticsearchVersion"
        },
        "ElasticsearchClusterConfig": {
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
          "rest.action.multi.allow_explicit_index": true
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
    "DomainArn": {
      "Value": {
        "Fn::GetAtt": [
          "ElasticsearchDomain",
          "DomainArn"
        ]
      }
    },
    "DomainEndpoint": {
      "Value": {
        "Fn::GetAtt": [
          "ElasticsearchDomain",
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
Description: ElasticsearchDomain resource
Parameters:
  DomainName:
    Description: User-defined Elasticsearch domain name
    Type: String
  ElasticsearchVersion:
    Description: User-defined Elasticsearch version
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
  ElasticsearchDomain:
    Type: 'AWS::Elasticsearch::Domain'
    Properties:
      DomainName:
        Ref: DomainName
      ElasticsearchVersion:
        Ref: ElasticsearchVersion
      ElasticsearchClusterConfig:
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
        rest.action.multi.allow_explicit_index: true
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
  DomainArn:
    Value:
      'Fn::GetAtt':
        - ElasticsearchDomain
        - DomainArn
  DomainEndpoint:
    Value:
      'Fn::GetAtt':
        - ElasticsearchDomain
        - DomainEndpoint
  SecurityGroupId:
    Value:
      Ref: mySecurityGroup
  SubnetId:
    Value:
      Ref: subnet
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon OpenSearch Service (legacy Elasticsearch resource)

AdvancedSecurityOptionsInput
