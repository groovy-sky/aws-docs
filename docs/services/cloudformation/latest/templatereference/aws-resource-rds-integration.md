This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::Integration

A zero-ETL integration with Amazon Redshift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::Integration",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "DataFilter" : String,
      "Description" : String,
      "IntegrationName" : String,
      "KMSKeyId" : String,
      "SourceArn" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::RDS::Integration
Properties:
  AdditionalEncryptionContext:
    Key: Value
  DataFilter: String
  Description: String
  IntegrationName: String
  KMSKeyId: String
  SourceArn: String
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`AdditionalEncryptionContext`

An optional set of non-secret key–value pairs that contains additional contextual
information about the data. For more information, see [Encryption\
context](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the _AWS Key Management Service Developer_
_Guide_.

You can only include this parameter if you specify the `KMSKeyId` parameter.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataFilter`

Data filters for the integration. These filters determine which tables
from the source database are sent to the target Amazon Redshift data warehouse.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_ "\\\-$,*.:?+\/]*`

_Minimum_: `1`

_Maximum_: `25600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the integration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationName`

The name of the integration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSKeyId`

The AWS Key Management System (AWS KMS) key identifier for the key to use to
encrypt the integration. If you don't specify an encryption key, RDS uses a default
AWS owned key.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceArn`

The Amazon Resource Name (ARN) of the database to use as the source for
replication.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:rds(-[a-z]*)?:[a-z0-9\-]*:[0-9]*:(cluster|db):[a-z][a-z0-9]*(-[a-z0-9]+)*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional array of key-value pairs to apply to this integration.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-integration-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN of the Redshift data warehouse to use as the target for replication.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the integration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreateTime`

The time when the integration was created, in Universal Coordinated Time
(UTC).

`IntegrationArn`

The ARN of the integration.

## Examples

- [Create an Aurora zero-ETL integration using an AWS owned KMS key](#aws-resource-rds-integration--examples--Create_an_Aurora_zero-ETL_integration_using_an_owned_KMS_key)

- [Create an Aurora zero-ETL integration using a customer managed KMS key](#aws-resource-rds-integration--examples--Create_an_Aurora_zero-ETL_integration_using_a_customer_managed_KMS_key)

- [Create an RDS zero-ETL integration using an AWS owned KMS key](#aws-resource-rds-integration--examples--Create_an_RDS_zero-ETL_integration_using_an_owned_KMS_key)

- [Create an RDS zero-ETL integration using a customer managed KMS key](#aws-resource-rds-integration--examples--Create_an_RDS_zero-ETL_integration_using_a_customer_managed_KMS_key)

### Create an Aurora zero-ETL integration using an AWS owned KMS key

The following example creates an Amazon Aurora MySQL DB cluster, Redshift provisioned cluster, and
a zero-ETL integration between the two clusters encrypted with an AWS owned KMS key.
For more information,
see [Working with Aurora zero-ETL integrations with Amazon Redshift](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/zero-etl.html)
in the _Amazon Aurora User Guide_.

#### JSON

```json

{
    "Parameters": {
        "Username": {
            "Description": "Username for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "bevelvoerder",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "Password": {
            "NoEcho": "true",
            "Description": "Password for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Passw0rd",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        },
        "RSUsername": {
            "Description": "Username for Redshift cluster access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "bevelvoerder",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "RSPassword": {
            "NoEcho": "true",
            "Description": "Password for Redshift cluster access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Passw0rd",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "RDSDBClusterParameterGroup": {
            "Type": "AWS::RDS::DBClusterParameterGroup",
            "Properties": {
                "Description": "Enables enhanced binlog",
                "Family": "aurora-mysql8.0",
                "Parameters": {
                    "aurora_enhanced_binlog": "1",
                    "binlog_format": "ROW",
                    "binlog_backup": "0",
                    "binlog_replication_globaldb": "0",
                    "binlog_row_metadata": "full",
                    "binlog_row_image": "full"
                }
            }
        },
        "RDSDBCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "DBClusterParameterGroupName": {
                    "Ref": "RDSDBClusterParameterGroup"
                },
                "Engine": "aurora-mysql",
                "EngineVersion": "8.0.mysql_aurora.3.05.1",
                "MasterUsername": {
                    "Ref": "Username"
                },
                "MasterUserPassword": {
                    "Ref": "Password"
                }
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBClusterIdentifier": {
                    "Ref": "RDSDBCluster"
                },
                "DBInstanceClass": "db.r5.large",
                "Engine": "aurora-mysql",
                "PubliclyAccessible": true
            }
        },
        "RSClusterParameterGroup": {
            "Type": "AWS::Redshift::ClusterParameterGroup",
            "Properties": {
                "Description": "Enables case sensitivity",
                "ParameterGroupFamily": "redshift-1.0",
                "Parameters": [
                    {
                        "ParameterName": "enable_case_sensitive_identifier",
                        "ParameterValue": "true"
                    }
                ]
            }
        },
        "RSCluster": {
            "Type": "AWS::Redshift::Cluster",
            "Properties": {
                "ClusterParameterGroupName": {
                    "Ref": "RSClusterParameterGroup"
                },
                "ClusterType": "multi-node",
                "DBName": "dev",
                "Encrypted": true,
                "MasterUsername": {
                    "Ref": "RSUsername"
                },
                "MasterUserPassword": {
                    "Ref": "RSPassword"
                },
                "NodeType": "ra3.xlplus",
                "NumberOfNodes": 2,
                "NamespaceResourcePolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "redshift:AuthorizeInboundIntegration",
                            "Condition": {
                                "StringEquals": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "RDSDBCluster",
                                            "DBClusterArn"
                                        ]
                                    }
                                }
                            }
                        },
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "redshift:CreateInboundIntegration"
                        }
                    ]
                }
            }
        },
        "Integration": {
            "Type": "AWS::RDS::Integration",
            "Properties": {
                "SourceArn": {
                    "Fn::GetAtt": [
                        "RDSDBCluster",
                        "DBClusterArn"
                    ]
                },
                "TargetArn": {
                    "Fn::GetAtt": [
                        "RSCluster",
                        "ClusterNamespaceArn"
                    ]
                }
            }
        }
    },
    "Outputs": {
        "IntegrationARN": {
            "Description": "Integration ARN",
            "Value": {
                "Ref": "Integration"
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Username:
    Description: Username for Aurora MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: "bevelvoerder"
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  Password:
    NoEcho: 'true'
    Description: Password for Aurora MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: "Passw0rd"
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
  RSUsername:
    Description: Username for Redshift cluster access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: "bevelvoerder"
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  RSPassword:
    NoEcho: 'true'
    Description: Password for Redshift cluster access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: "Passw0rd"
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSDBClusterParameterGroup:
    Type: 'AWS::RDS::DBClusterParameterGroup'
    Properties:
      Description: Enables enhanced binlog
      Family: aurora-mysql8.0
      Parameters:
        aurora_enhanced_binlog: '1'
        binlog_format: ROW
        binlog_backup: '0'
        binlog_replication_globaldb: '0'
        binlog_row_metadata: full
        binlog_row_image: full
  RDSDBCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      DBClusterParameterGroupName: !Ref RDSDBClusterParameterGroup
      Engine: aurora-mysql
      EngineVersion: 8.0.mysql_aurora.3.05.1
      MasterUsername: !Ref Username
      MasterUserPassword: !Ref Password
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBClusterIdentifier: !Ref RDSDBCluster
      DBInstanceClass: db.r5.large
      Engine: aurora-mysql
      PubliclyAccessible: true
  RSClusterParameterGroup:
    Type: 'AWS::Redshift::ClusterParameterGroup'
    Properties:
      Description: Enables case sensitivity
      ParameterGroupFamily: redshift-1.0
      Parameters:
        - ParameterName: enable_case_sensitive_identifier
          ParameterValue: 'true'
  RSCluster:
    Type: 'AWS::Redshift::Cluster'
    Properties:
      ClusterParameterGroupName: !Ref RSClusterParameterGroup
      ClusterType: multi-node
      DBName: dev
      Encrypted: true
      MasterUsername: !Ref RSUsername
      MasterUserPassword: !Ref RSPassword
      NodeType: ra3.xlplus
      NumberOfNodes: 2
      NamespaceResourcePolicy:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: redshift:AuthorizeInboundIntegration
            Condition:
              StringEquals:
                aws:SourceArn: !GetAtt RDSDBCluster.DBClusterArn
          - Effect: Allow
            Principal:
              AWS: !Join [":", ["arn", !Ref AWS::Partition, "iam", "", !Ref AWS::AccountId, "root"]]
            Action: redshift:CreateInboundIntegration
  Integration:
    Type: 'AWS::RDS::Integration'
    Properties:
      SourceArn: !GetAtt RDSDBCluster.DBClusterArn
      TargetArn: !GetAtt RSCluster.ClusterNamespaceArn
Outputs:
  IntegrationARN:
    Description: Integration ARN
    Value: !Ref Integration
```

### Create an Aurora zero-ETL integration using a customer managed KMS key

The following example creates an Amazon Aurora MySQL DB cluster, Redshift provisioned cluster, and
a zero-ETL integration between the two clusters encrypted with a customer managed KMS key.
For more information,
see [Working with Aurora zero-ETL integrations with Amazon Redshift](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/zero-etl.html)
in the _Amazon Aurora User Guide_.

#### JSON

```json

{
    "Parameters": {
        "Username": {
            "Description": "Username for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "bevelvoerder",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "Password": {
            "NoEcho": "true",
            "Description": "Password for Aurora MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Passw0rd",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        },
        "RSUsername": {
            "Description": "Username for Redshift cluster access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "bevelvoerder",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "RSPassword": {
            "NoEcho": "true",
            "Description": "Password for Redshift cluster access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Passw0rd",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "kmsKey": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "Description": "Key used to encrypt the zero-ETL integration.",
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Sid": "statement1",
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "kms:CreateGrant",
                            "Resource": "*"
                        },
                        {
                            "Sid": "statement2",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "kms:*",
                            "Resource": "*"
                        }
                    ]
                }
            }
        },
        "RDSDBClusterParameterGroup": {
            "Type": "AWS::RDS::DBClusterParameterGroup",
            "Properties": {
                "Description": "Enables enhanced binlog",
                "Family": "aurora-mysql8.0",
                "Parameters": {
                    "aurora_enhanced_binlog": "1",
                    "binlog_format": "ROW",
                    "binlog_backup": "0",
                    "binlog_replication_globaldb": "0",
                    "binlog_row_metadata": "full",
                    "binlog_row_image": "full"
                }
            }
        },
        "RDSDBCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "DBClusterParameterGroupName": {
                    "Ref": "RDSDBClusterParameterGroup"
                },
                "Engine": "aurora-mysql",
                "EngineVersion": "8.0.mysql_aurora.3.05.1",
                "MasterUsername": {
                    "Ref": "Username"
                },
                "MasterUserPassword": {
                    "Ref": "Password"
                }
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBClusterIdentifier": {
                    "Ref": "RDSDBCluster"
                },
                "DBInstanceClass": "db.r5.large",
                "Engine": "aurora-mysql",
                "PubliclyAccessible": true
            }
        },
        "RSClusterParameterGroup": {
            "Type": "AWS::Redshift::ClusterParameterGroup",
            "Properties": {
                "Description": "Enables case sensitivity",
                "ParameterGroupFamily": "redshift-1.0",
                "Parameters": [
                    {
                        "ParameterName": "enable_case_sensitive_identifier",
                        "ParameterValue": "true"
                    }
                ]
            }
        },
        "RSCluster": {
            "Type": "AWS::Redshift::Cluster",
            "Properties": {
                "ClusterParameterGroupName": {
                    "Ref": "RSClusterParameterGroup"
                },
                "ClusterType": "multi-node",
                "DBName": "dev",
                "Encrypted": true,
                "MasterUsername": {
                    "Ref": "RSUsername"
                },
                "MasterUserPassword": {
                    "Ref": "RSPassword"
                },
                "NodeType": "ra3.xlplus",
                "NumberOfNodes": 2,
                "NamespaceResourcePolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "redshift:AuthorizeInboundIntegration",
                            "Condition": {
                                "StringEquals": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "RDSDBCluster",
                                            "DBClusterArn"
                                        ]
                                    }
                                }
                            }
                        },
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "redshift:CreateInboundIntegration"
                        }
                    ]
                }
            }
        },
        "Integration": {
            "Type": "AWS::RDS::Integration",
            "Properties": {
                "SourceArn": {
                    "Fn::GetAtt": [
                        "RDSDBCluster",
                        "DBClusterArn"
                    ]
                },
                "TargetArn": {
                    "Fn::GetAtt": [
                        "RSCluster",
                        "ClusterNamespaceArn"
                    ]
                },
                "KMSKeyId": {
                    "Fn::GetAtt": [
                        "kmsKey",
                        "Arn"
                    ]
                },
                "AdditionalEncryptionContext": {
                    "custom-context-1": "custom-context-value-1",
                    "custom-context-2": "custom-context-value-2"
                }
            }
        }
    },
    "Outputs": {
        "IntegrationARN": {
            "Description": "Integration ARN",
            "Value": {
                "Ref": "Integration"
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
  Username:
    Description: Username for Aurora MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: "bevelvoerder"
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  Password:
    NoEcho: 'true'
    Description: Password for Aurora MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: "Passw0rd"
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
  RSUsername:
    Description: Username for Redshift cluster access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: "bevelvoerder"
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  RSPassword:
    NoEcho: 'true'
    Description: Password for Redshift cluster access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: "Passw0rd"
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  kmsKey:
    Type: 'AWS::KMS::Key'
    Properties:
      Description: Key used to encrypt the zero-ETL integration.
      KeyPolicy:
        Version: 2012-10-17
        Statement:
          - Sid: statement1
            Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: "kms:CreateGrant"
            Resource: "*"
          - Sid: statement2
            Effect: Allow
            Principal:
              AWS: !Join [":", ["arn", !Ref AWS::Partition, "iam", "", !Ref AWS::AccountId, "root"]]
            Action: "kms:*"
            Resource: "*"
  RDSDBClusterParameterGroup:
    Type: 'AWS::RDS::DBClusterParameterGroup'
    Properties:
      Description: Enables enhanced binlog
      Family: aurora-mysql8.0
      Parameters:
        aurora_enhanced_binlog: '1'
        binlog_format: ROW
        binlog_backup: '0'
        binlog_replication_globaldb: '0'
        binlog_row_metadata: full
        binlog_row_image: full
  RDSDBCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      DBClusterParameterGroupName: !Ref RDSDBClusterParameterGroup
      Engine: aurora-mysql
      EngineVersion: 8.0.mysql_aurora.3.05.1
      MasterUsername: !Ref Username
      MasterUserPassword: !Ref Password
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBClusterIdentifier: !Ref RDSDBCluster
      DBInstanceClass: db.r5.large
      Engine: aurora-mysql
      PubliclyAccessible: true
  RSClusterParameterGroup:
    Type: 'AWS::Redshift::ClusterParameterGroup'
    Properties:
      Description: Enables case sensitivity
      ParameterGroupFamily: redshift-1.0
      Parameters:
        - ParameterName: enable_case_sensitive_identifier
          ParameterValue: 'true'
  RSCluster:
    Type: 'AWS::Redshift::Cluster'
    Properties:
      ClusterParameterGroupName: !Ref RSClusterParameterGroup
      ClusterType: multi-node
      DBName: dev
      Encrypted: true
      MasterUsername: !Ref RSUsername
      MasterUserPassword: !Ref RSPassword
      NodeType: ra3.xlplus
      NumberOfNodes: 2
      NamespaceResourcePolicy:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: redshift:AuthorizeInboundIntegration
            Condition:
              StringEquals:
                aws:SourceArn: !GetAtt RDSDBCluster.DBClusterArn
          - Effect: Allow
            Principal:
              AWS: !Join [":", ["arn", !Ref AWS::Partition, "iam", "", !Ref AWS::AccountId, "root"]]
            Action: redshift:CreateInboundIntegration
  Integration:
    Type: 'AWS::RDS::Integration'
    Properties:
      SourceArn: !GetAtt RDSDBCluster.DBClusterArn
      TargetArn: !GetAtt RSCluster.ClusterNamespaceArn
      KMSKeyId: !GetAtt kmsKey.Arn
      AdditionalEncryptionContext:
        custom-context-1: custom-context-value-1
        custom-context-2: custom-context-value-2
Outputs:
  IntegrationARN:
    Description: Integration ARN
    Value: !Ref Integration
```

### Create an RDS zero-ETL integration using an AWS owned KMS key

The following example creates an Amazon RDS MySQL DB instance, Redshift
provisioned cluster, and a zero-ETL integration between them encrypted with an
AWS owned key. For more information, see [Working with RDS zero-ETL\
integrations with Amazon Redshift](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/zero-etl.html) in the _Amazon RDS User_
_Guide_.

#### JSON

```json

{
    "Parameters": {
        "Username": {
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "Password": {
            "NoEcho": "true",
            "Description": "Password for MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        },
        "RSUsername": {
            "Description": "Username for Redshift cluster access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "RSPassword": {
            "NoEcho": "true",
            "Description": "Password for Redshift cluster access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Sqladmin2024",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "RDSDBParameterGroup": {
            "Type": "AWS::RDS::DBParameterGroup",
            "Properties": {
            	"DBParameterGroupName": "zeroetl-mysql-parameter-group",
                "Description": "Enables enhanced binlog",
                "Family": "mysql8.0",
                "Parameters": {
                    "binlog_checksum": "NONE",
                    "binlog_format": "ROW",
                    "binlog_row_image": "FULL"
                }
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
            	"DBInstanceIdentifier": "zeroetl-mysql-source",
                "DBParameterGroupName": {
                    "Ref": "RDSDBParameterGroup"
                },
                "Engine": "mysql",
                "EngineVersion": "8.0.33",
                "MasterUsername": {
                    "Ref": "Username"
                },
                "MasterUserPassword": {
                    "Ref": "Password"
                },
                "DBInstanceClass": "db.r5.large",
                "BackupRetentionPeriod": 1,
                "StorageType": "gp2",
                "AllocatedStorage": "100"
            }
        },
        "RSClusterParameterGroup": {
            "Type": "AWS::Redshift::ClusterParameterGroup",
            "Properties": {
            	"ParameterGroupName": "zeroetl-redshift-cluster-parameter-group",
                "Description": "Enables case sensitivity",
                "ParameterGroupFamily": "redshift-1.0",
                "Parameters": [
                    {
                        "ParameterName": "enable_case_sensitive_identifier",
                        "ParameterValue": "true"
                    }
                ]
            }
        },
        "RSCluster": {
            "Type": "AWS::Redshift::Cluster",
            "Properties": {
            	"ClusterIdentifier": "zeroetl-redshift-cluster",
                "ClusterParameterGroupName": {
                    "Ref": "RSClusterParameterGroup"
                },
                "ClusterType": "multi-node",
                "DBName": "sqladmin",
                "MaintenanceTrackName": "PREVIEW_2023",
                "Encrypted": true,
                "MasterUsername": {
                    "Ref": "RSUsername"
                },
                "MasterUserPassword": {
                    "Ref": "RSPassword"
                },
                "NodeType": "ra3.xlplus",
                "NumberOfNodes": 2,
                "NamespaceResourcePolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "redshift:AuthorizeInboundIntegration",
                            "Condition": {
                                "StringEquals": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "RDSDBInstance",
                                            "DBInstanceArn"
                                        ]
                                    }
                                }
                            }
                        },
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "redshift:CreateInboundIntegration"
                        }
                    ]
                }
            }
        },
        "Integration": {
            "Type": "AWS::RDS::Integration",
            "Properties": {
            	"IntegrationName": "zeroetl-integration-from-mysql-to-redshift",
                "SourceArn": {
                    "Fn::GetAtt": [
                        "RDSDBInstance",
                        "DBInstanceArn"
                    ]
                },
                "TargetArn": {
                    "Fn::GetAtt": [
                        "RSCluster",
                        "ClusterNamespaceArn"
                    ]
                }
            }
        }
    },
    "Outputs": {
    	"IntegrationCreateTime": {
    		"Description": "Integration Creation Time",
    		"Value": {
    			"Fn::GetAtt": [
    				"Integration",
    				"CreateTime"
    			]
    		}
    	},
        "IntegrationARN": {
            "Description": "Integration ARN",
            "Value": {
                "Ref": "Integration"
            }
        }
    }
}

```

#### YAML

```yaml

Parameters:
  Username:
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  Password:
    NoEcho: 'true'
    Description: Password for MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
  RSUsername:
    Description: Username for Redshift cluster access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  RSPassword:
    NoEcho: 'true'
    Description: Password for Redshift cluster access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: Sqladmin2024
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSDBParameterGroup:
    Type: 'AWS::RDS::DBParameterGroup'
    Properties:
      DBParameterGroupName: zeroetl-mysql-parameter-group
      Description: Enables enhanced binlog
      Family: mysql8.0
      Parameters:
        binlog_checksum: NONE
        binlog_format: ROW
        binlog_row_image: FULL
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBInstanceIdentifier: zeroetl-mysql-source
      DBParameterGroupName: !Ref RDSDBParameterGroup
      Engine: mysql
      EngineVersion: 8.0.33
      MasterUsername: !Ref Username
      MasterUserPassword: !Ref Password
      DBInstanceClass: db.r5.large
      BackupRetentionPeriod: 1
      StorageType: gp2
      AllocatedStorage: '100'
  RSClusterParameterGroup:
    Type: 'AWS::Redshift::ClusterParameterGroup'
    Properties:
      ParameterGroupName: zeroetl-redshift-cluster-parameter-group
      Description: Enables case sensitivity
      ParameterGroupFamily: redshift-1.0
      Parameters:
        - ParameterName: enable_case_sensitive_identifier
          ParameterValue: 'true'
  RSCluster:
    Type: 'AWS::Redshift::Cluster'
    Properties:
      ClusterIdentifier: zeroetl-redshift-cluster
      ClusterParameterGroupName: !Ref RSClusterParameterGroup
      ClusterType: multi-node
      DBName: sqladmin
      MaintenanceTrackName: PREVIEW_2023
      Encrypted: true
      MasterUsername: !Ref RSUsername
      MasterUserPassword: !Ref RSPassword
      NodeType: ra3.xlplus
      NumberOfNodes: 2
      NamespaceResourcePolicy:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: 'redshift:AuthorizeInboundIntegration'
            Condition:
              StringEquals:
                'aws:SourceArn': !GetAtt
                  - RDSDBInstance
                  - DBInstanceArn
          - Effect: Allow
            Principal:
              AWS: !Join
                - ':'
                - - arn
                  - !Ref 'AWS::Partition'
                  - iam
                  - ''
                  - !Ref 'AWS::AccountId'
                  - root
            Action: 'redshift:CreateInboundIntegration'
  Integration:
    Type: 'AWS::RDS::Integration'
    Properties:
      IntegrationName: zeroetl-integration-from-mysql-to-redshift
      SourceArn: !GetAtt
        - RDSDBInstance
        - DBInstanceArn
      TargetArn: !GetAtt
        - RSCluster
        - ClusterNamespaceArn
Outputs:
  IntegrationCreateTime:
    Description: Integration Creation Time
    Value: !GetAtt
      - Integration
      - CreateTime
  IntegrationARN:
    Description: Integration ARN
    Value: !Ref Integration

```

### Create an RDS zero-ETL integration using a customer managed KMS key

The following example creates an Amazon RDS MySQL DB instance, Redshift
provisioned cluster, and a zero-ETL integration between them encrypted with a
customer managed KMS key. For more information, see [Working with RDS zero-ETL\
integrations with Amazon Redshift](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/zero-etl.html) in the _Amazon RDS User_
_Guide_.

#### JSON

```json

{
    "Parameters": {
        "Username": {
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "Password": {
            "NoEcho": "true",
            "Description": "Password for MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        },
        "RSUsername": {
            "Description": "Username for Redshift cluster access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "Default": "sqladmin",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "RSPassword": {
            "NoEcho": "true",
            "Description": "Password for Redshift cluster access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "Default": "Sqladmin2024",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
    	"kmsKey": {
            "Type": "AWS::KMS::Key",
            "Properties": {
                "Description": "Key used to encrypt the zero-ETL integration.",
                "KeyPolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Sid": "statement1",
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "kms:CreateGrant",
                            "Resource": "*"
                        },
                        {
                            "Sid": "statement2",
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "kms:*",
                            "Resource": "*"
                        }
                    ]
                }
            }
        },
        "RDSDBParameterGroup": {
            "Type": "AWS::RDS::DBParameterGroup",
            "Properties": {
            	"DBParameterGroupName": "zeroetl-mysql-parameter-group",
                "Description": "Enables enhanced binlog",
                "Family": "mysql8.0",
                "Parameters": {
                    "binlog_checksum": "NONE",
                    "binlog_format": "ROW",
                    "binlog_row_image": "FULL"
                }
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
            	"DBInstanceIdentifier": "zeroetl-mysql-source",
                "DBParameterGroupName": {
                    "Ref": "RDSDBParameterGroup"
                },
                "Engine": "mysql",
                "EngineVersion": "8.0.33",
                "MasterUsername": {
                    "Ref": "Username"
                },
                "MasterUserPassword": {
                    "Ref": "Password"
                },
                "DBInstanceClass": "db.r5.large",
                "BackupRetentionPeriod": 1,
                "StorageType": "gp2",
                "AllocatedStorage": "100"
            }
        },
        "RSClusterParameterGroup": {
            "Type": "AWS::Redshift::ClusterParameterGroup",
            "Properties": {
            	"ParameterGroupName": "zeroetl-redshift-cluster-parameter-group",
                "Description": "Enables case sensitivity",
                "ParameterGroupFamily": "redshift-1.0",
                "Parameters": [
                    {
                        "ParameterName": "enable_case_sensitive_identifier",
                        "ParameterValue": "true"
                    }
                ]
            }
        },
        "RSCluster": {
            "Type": "AWS::Redshift::Cluster",
            "Properties": {
            	"ClusterIdentifier": "zeroetl-redshift-cluster",
                "ClusterParameterGroupName": {
                    "Ref": "RSClusterParameterGroup"
                },
                "ClusterType": "multi-node",
                "DBName": "sqladmin",
                "MaintenanceTrackName": "PREVIEW_2023",
                "Encrypted": true,
                "MasterUsername": {
                    "Ref": "RSUsername"
                },
                "MasterUserPassword": {
                    "Ref": "RSPassword"
                },
                "NodeType": "ra3.xlplus",
                "NumberOfNodes": 2,
                "NamespaceResourcePolicy": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": "redshift.amazonaws.com"
                            },
                            "Action": "redshift:AuthorizeInboundIntegration",
                            "Condition": {
                                "StringEquals": {
                                    "aws:SourceArn": {
                                        "Fn::GetAtt": [
                                            "RDSDBInstance",
                                            "DBInstanceArn"
                                        ]
                                    }
                                }
                            }
                        },
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "AWS": {
                                    "Fn::Join": [
                                        ":",
                                        [
                                            "arn",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "iam",
                                            "",
                                            {
                                                "Ref": "AWS::AccountId"
                                            },
                                            "root"
                                        ]
                                    ]
                                }
                            },
                            "Action": "redshift:CreateInboundIntegration"
                        }
                    ]
                }
            }
        },
        "Integration": {
            "Type": "AWS::RDS::Integration",
            "Properties": {
            	"IntegrationName": "zeroetl-integration-from-mysql-to-redshift",
                "SourceArn": {
                    "Fn::GetAtt": [
                        "RDSDBInstance",
                        "DBInstanceArn"
                    ]
                },
                "TargetArn": {
                    "Fn::GetAtt": [
                        "RSCluster",
                        "ClusterNamespaceArn"
                    ]
                },
                "KMSKeyId": {
                    "Fn::GetAtt": [
                        "kmsKey",
                        "Arn"
                    ]
                },
                "AdditionalEncryptionContext": {
                    "integrationName": "zeroetl-integration-from-mysql-to-redshift"
                }
            }
        }
    },
    "Outputs": {
    	"IntegrationCreateTime": {
    		"Description": "Integration Creation Time",
    		"Value": {
    			"Fn::GetAtt": [
    				"Integration",
    				"CreateTime"
    			]
    		}
    	},
        "IntegrationARN": {
            "Description": "Integration ARN",
            "Value": {
                "Ref": "Integration"
            }
        }
    }
}

```

#### YAML

```yaml

Parameters:
  Username:
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  Password:
    NoEcho: 'true'
    Description: Password for MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
  RSUsername:
    Description: Username for Redshift cluster access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    Default: sqladmin
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  RSPassword:
    NoEcho: 'true'
    Description: Password for Redshift cluster access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    Default: Sqladmin2024
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  kmsKey:
    Type: 'AWS::KMS::Key'
    Properties:
      Description: Key used to encrypt the zero-ETL integration.
      KeyPolicy:
        Version: 2012-10-17
        Statement:
          - Sid: statement1
            Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: 'kms:CreateGrant'
            Resource: '*'
          - Sid: statement2
            Effect: Allow
            Principal:
              AWS: !Join
                - ':'
                - - arn
                  - !Ref 'AWS::Partition'
                  - iam
                  - ''
                  - !Ref 'AWS::AccountId'
                  - root
            Action: 'kms:*'
            Resource: '*'
  RDSDBParameterGroup:
    Type: 'AWS::RDS::DBParameterGroup'
    Properties:
      DBParameterGroupName: zeroetl-mysql-parameter-group
      Description: Enables enhanced binlog
      Family: mysql8.0
      Parameters:
        binlog_checksum: NONE
        binlog_format: ROW
        binlog_row_image: FULL
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      DBInstanceIdentifier: zeroetl-mysql-source
      DBParameterGroupName: !Ref RDSDBParameterGroup
      Engine: mysql
      EngineVersion: 8.0.33
      MasterUsername: !Ref Username
      MasterUserPassword: !Ref Password
      DBInstanceClass: db.r5.large
      BackupRetentionPeriod: 1
      StorageType: gp2
      AllocatedStorage: '100'
  RSClusterParameterGroup:
    Type: 'AWS::Redshift::ClusterParameterGroup'
    Properties:
      ParameterGroupName: zeroetl-redshift-cluster-parameter-group
      Description: Enables case sensitivity
      ParameterGroupFamily: redshift-1.0
      Parameters:
        - ParameterName: enable_case_sensitive_identifier
          ParameterValue: 'true'
  RSCluster:
    Type: 'AWS::Redshift::Cluster'
    Properties:
      ClusterIdentifier: zeroetl-redshift-cluster
      ClusterParameterGroupName: !Ref RSClusterParameterGroup
      ClusterType: multi-node
      DBName: sqladmin
      MaintenanceTrackName: PREVIEW_2023
      Encrypted: true
      MasterUsername: !Ref RSUsername
      MasterUserPassword: !Ref RSPassword
      NodeType: ra3.xlplus
      NumberOfNodes: 2
      NamespaceResourcePolicy:
        Version: 2012-10-17
        Statement:
          - Effect: Allow
            Principal:
              Service: redshift.amazonaws.com
            Action: 'redshift:AuthorizeInboundIntegration'
            Condition:
              StringEquals:
                'aws:SourceArn': !GetAtt
                  - RDSDBInstance
                  - DBInstanceArn
          - Effect: Allow
            Principal:
              AWS: !Join
                - ':'
                - - arn
                  - !Ref 'AWS::Partition'
                  - iam
                  - ''
                  - !Ref 'AWS::AccountId'
                  - root
            Action: 'redshift:CreateInboundIntegration'
  Integration:
    Type: 'AWS::RDS::Integration'
    Properties:
      IntegrationName: zeroetl-integration-from-mysql-to-redshift
      SourceArn: !GetAtt
        - RDSDBInstance
        - DBInstanceArn
      TargetArn: !GetAtt
        - RSCluster
        - ClusterNamespaceArn
      KMSKeyId: !GetAtt
        - kmsKey
        - Arn
      AdditionalEncryptionContext:
        integrationName: zeroetl-integration-from-mysql-to-redshift
Outputs:
  IntegrationCreateTime:
    Description: Integration Creation Time
    Value: !GetAtt
      - Integration
      - CreateTime
  IntegrationARN:
    Description: Integration ARN
    Value: !Ref Integration

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
