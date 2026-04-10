This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::GlobalCluster

The `AWS::RDS::GlobalCluster` resource creates or updates an Amazon Aurora global database
spread across multiple AWS Regions.

The global database contains a single primary cluster with read-write capability,
and a read-only secondary cluster that receives
data from the primary cluster through high-speed replication
performed by the Aurora storage subsystem.

You can create a global database that is initially empty, and then
add a primary cluster and a secondary cluster to it.

For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](../../../amazonrds/latest/aurorauserguide/aurora-global-database.md) in the _Amazon Aurora User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::GlobalCluster",
  "Properties" : {
      "DeletionProtection" : Boolean,
      "Engine" : String,
      "EngineLifecycleSupport" : String,
      "EngineVersion" : String,
      "GlobalClusterIdentifier" : String,
      "SourceDBClusterIdentifier" : String,
      "StorageEncrypted" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::GlobalCluster
Properties:
  DeletionProtection: Boolean
  Engine: String
  EngineLifecycleSupport: String
  EngineVersion: String
  GlobalClusterIdentifier: String
  SourceDBClusterIdentifier: String
  StorageEncrypted: Boolean
  Tags:
    - Tag

```

## Properties

`DeletionProtection`

Specifies whether to enable deletion protection for the new global database cluster.
The global database can't be deleted when deletion protection is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The database engine to use for this global database cluster.

Valid Values: `aurora-mysql | aurora-postgresql`

Constraints:

- Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the engine of the source DB cluster.

_Required_: Conditional

_Type_: String

_Allowed values_: `aurora | aurora-mysql | aurora-postgresql`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineLifecycleSupport`

The life cycle type for this global database cluster.

###### Note

By default, this value is set to `open-source-rds-extended-support`, which enrolls your global cluster into Amazon RDS Extended Support.
At the end of standard support, you can avoid charges for Extended Support by setting the value to `open-source-rds-extended-support-disabled`. In this case,
creating the global cluster will fail if the DB major version is past its end of standard support date.

This setting only applies to Aurora PostgreSQL-based global databases.

You can use this setting to enroll your global cluster into Amazon RDS Extended Support. With RDS Extended Support,
you can run the selected major engine version on your global cluster past the end of standard support for that engine version. For more information, see [Amazon RDS Extended Support with Amazon Aurora](../../../amazonrds/latest/aurorauserguide/extended-support.md) in the _Amazon Aurora User Guide_.

Valid Values: `open-source-rds-extended-support | open-source-rds-extended-support-disabled`

Default: `open-source-rds-extended-support`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EngineVersion`

The engine version to use for this global database cluster.

Constraints:

- Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the engine version of the source DB cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalClusterIdentifier`

The cluster identifier for this global database cluster. This parameter is stored as a lowercase string.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceDBClusterIdentifier`

The Amazon Resource Name (ARN) to use as the primary cluster of the global database.

If you provide a value for this parameter, don't specify values for the following settings because Amazon Aurora uses the values from the specified source DB cluster:

- `DatabaseName`

- `Engine`

- `EngineVersion`

- `StorageEncrypted`

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageEncrypted`

Specifies whether to enable storage encryption for the new global database cluster.

Constraints:

- Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the setting from the source DB cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata assigned to an Amazon RDS resource consisting of a key-value pair.

For more information, see
[Tagging Amazon RDS resources](../../../amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](../../../amazonrds/latest/aurorauserguide/user-tagging.md) in the _Amazon Aurora User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-globalcluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the global database cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Creating a Global Database cluster for Aurora MySQL](#aws-resource-rds-globalcluster--examples--Creating_a_Global_Database_cluster_for_Aurora_MySQL)

- [Creating a Global Database cluster for Aurora PostgreSQL](#aws-resource-rds-globalcluster--examples--Creating_a_Global_Database_cluster_for_Aurora_PostgreSQL)

- [Adding a Region to an Aurora database cluster](#aws-resource-rds-globalcluster--examples--Adding_a_Region_to_an_Aurora_database_cluster)

- [Adding a DB cluster to a Global Database cluster](#aws-resource-rds-globalcluster--examples--Adding_a_DB_cluster_to_a_Global_Database_cluster)

### Creating a Global Database cluster for Aurora MySQL

The following example creates a global database cluster with an Aurora MySQL DB cluster and DB instance.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "GlobalClusterIdentifier": {
          "Type": "String",
          "Description": "Identifier used for global database cluster",
          "AllowedPattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$"
        },
        "username": {
            "NoEcho": "true",
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "password": {
            "NoEcho": "true",
            "Description": "Password for MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "GlobalCluster": {
            "Type": "AWS::RDS::GlobalCluster",
            "Properties": {
                "GlobalClusterIdentifier": {
                    "Ref": "GlobalClusterIdentifier"
                },
                "SourceDBClusterIdentifier": {
                    "Ref": "RDSCluster"
                }
            }
        },
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "MasterUsername": {
                    "Ref": "username"
                },
                "MasterUserPassword": {
                    "Ref": "password"
                },
                "DBClusterParameterGroupName": "default.aurora-mysql5.7",
                "Engine": "aurora-mysql"
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "DBParameterGroupName": "default.aurora-mysql5.7",
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r5.xlarge"
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  GlobalClusterIdentifier:
    Type: String
    Description: Identifier used for global database cluster
    AllowedPattern: '^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$'
  username:
    NoEcho: 'true'
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  password:
    NoEcho: 'true'
    Description: Password for MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  GlobalCluster:
    Type: 'AWS::RDS::GlobalCluster'
    Properties:
      GlobalClusterIdentifier: !Ref GlobalClusterIdentifier
      SourceDBClusterIdentifier: !Ref RDSCluster
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref username
      MasterUserPassword: !Ref password
      DBClusterParameterGroupName: default.aurora-mysql5.7
      Engine: aurora-mysql
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-mysql
      DBClusterIdentifier: !Ref RDSCluster
      DBParameterGroupName: default.aurora-mysql5.7
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r5.xlarge

```

### Creating a Global Database cluster for Aurora PostgreSQL

The following example creates a global database cluster with an Aurora PostgreSQL DB cluster and DB instance.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "GlobalClusterIdentifier": {
          "Type": "String",
          "Description": "Identifier used for global database cluster",
          "AllowedPattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$"
        },
        "username": {
            "NoEcho": "true",
            "Description": "Username for PostgreSQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "password": {
            "NoEcho": "true",
            "Description": "Password for PostgreSQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "GlobalCluster": {
            "Type": "AWS::RDS::GlobalCluster",
            "Properties": {
                "GlobalClusterIdentifier": {
                    "Ref": "GlobalClusterIdentifier"
                },
                "SourceDBClusterIdentifier": {
                    "Ref": "RDSCluster"
                }
            }
        },
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "MasterUsername": {
                    "Ref": "username"
                },
                "MasterUserPassword": {
                    "Ref": "password"
                },
                "DBClusterParameterGroupName": "default.aurora-postgresql11",
                "Engine": "aurora-postgresql"
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "Engine": "aurora-postgresql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "DBParameterGroupName": "default.aurora-postgresql11",
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r5.xlarge"
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  GlobalClusterIdentifier:
    Type: String
    Description: Identifier used for global database cluster
    AllowedPattern: '^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$'
  username:
    NoEcho: 'true'
    Description: Username for PostgreSQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  password:
    NoEcho: 'true'
    Description: Password for PostgreSQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  GlobalCluster:
    Type: 'AWS::RDS::GlobalCluster'
    Properties:
      GlobalClusterIdentifier: !Ref GlobalClusterIdentifier
      SourceDBClusterIdentifier: !Ref RDSCluster
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref username
      MasterUserPassword: !Ref password
      DBClusterParameterGroupName: default.aurora-postgresql11
      Engine: aurora-postgresql
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-postgresql
      DBClusterIdentifier: !Ref RDSCluster
      DBParameterGroupName: default.aurora-postgresql11
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r5.xlarge

```

### Adding a Region to an Aurora database cluster

The following example creates a new Aurora DB cluster, attaches it to a global database cluster as a read-only
secondary cluster, and then adds a DB instance to the new DB cluster.

Specify the `GlobalClusterIdentifier` of a global database cluster with the primary DB cluster in a separate AWS Region.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "GlobalClusterIdentifier": {
          "Type": "String",
          "Description": "Identifier used for global database cluster",
          "AllowedPattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$"
        }
    },
    "Resources": {
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "GlobalClusterIdentifier": {
                    "Ref": "GlobalClusterIdentifier"
                },
                "DBClusterParameterGroupName": "default.aurora-mysql5.7",
                "Engine": "aurora-mysql"
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "DBParameterGroupName": "default.aurora-mysql5.7",
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r5.xlarge"
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  GlobalClusterIdentifier:
    Type: String
    Description: Identifier used for global database cluster
    AllowedPattern: '^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$'
Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      GlobalClusterIdentifier: !Ref GlobalClusterIdentifier
      DBClusterParameterGroupName: default.aurora-mysql5.7
      Engine: aurora-mysql
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-mysql
      DBClusterIdentifier: !Ref RDSCluster
      DBParameterGroupName: default.aurora-mysql5.7
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r5.xlarge

```

### Adding a DB cluster to a Global Database cluster

The following example adds a DB cluster to a global database cluster.

The example includes the template that was used to create the DB cluster. After the DB cluster created by the first template exists,
the second template in the example adds the DB cluster to a global database cluster.

#### JSON

```json

The following template was used to create DB cluster that you want to add to the global database cluster.
{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "username": {
            "NoEcho": "true",
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
         },
         "password": {
             "NoEcho": "true",
             "Description": "Password MySQL database access",
             "Type": "String",
             "MinLength": "8",
             "MaxLength": "41",
             "AllowedPattern": "[a-zA-Z0-9]*",
             "ConstraintDescription": "must contain only alphanumeric characters."
         }
    },
    "Resources": {
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "MasterUsername": {
                    "Ref": "username"
                },
                "MasterUserPassword": {
                    "Ref": "password"
                },
                "DBClusterParameterGroupName": "default.aurora-mysql8.0",
                "Engine": "aurora-mysql",
                "EngineVersion": "8.0.mysql_aurora.8.0.30"
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
                "Properties": {
                    "Engine": "aurora-mysql",
                    "DBClusterIdentifier": {
                        "Ref": "RDSCluster"
                },
                "DBParameterGroupName": "default.aurora-mysql8.0",
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r5.xlarge"
            }
        }
    }
}

The following template adds the DB cluster created by the previous template to a global database cluster.
{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "GlobalClusterIdentifier": {
            "Description": "Global cluster identifier",
            "Type": "String"
        },
        "username": {
            "NoEcho": "true",
            "Description": "Username for MySQL database access",
            "Type": "String",
            "MinLength": "1",
            "MaxLength": "16",
            "AllowedPattern": "[a-zA-Z][a-zA-Z0-9]*",
            "ConstraintDescription": "must begin with a letter and contain only alphanumeric characters."
        },
        "password": {
            "NoEcho": "true",
            "Description": "Password MySQL database access",
            "Type": "String",
            "MinLength": "8",
            "MaxLength": "41",
            "AllowedPattern": "[a-zA-Z0-9]*",
            "ConstraintDescription": "must contain only alphanumeric characters."
        }
    },
    "Resources": {
        "GlobalCluster": {
            "Type": "AWS::RDS::GlobalCluster",
            "Properties": {
                "GlobalClusterIdentifier": {
                    "Ref": "GlobalClusterIdentifier"
                },
                "SourceDBClusterIdentifier": {
                    "Ref": "RDSCluster"
                }
            }
        },
        "RDSCluster": {
            "Type": "AWS::RDS::DBCluster",
            "Properties": {
                "MasterUsername": {
                    "Ref": "username"
                },
                "MasterUserPassword": {
                    "Ref": "password"
                },
                "DBClusterParameterGroupName": "default.aurora-mysql8.0",
                "Engine": "aurora-mysql",
                "EngineVersion": "8.0.mysql_aurora.8.0.30"
            }
        },
        "RDSDBInstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "Engine": "aurora-mysql",
                "DBClusterIdentifier": {
                    "Ref": "RDSCluster"
                },
                "DBParameterGroupName": "default.aurora-mysql8.0",
                "PubliclyAccessible": "true",
                "DBInstanceClass": "db.r5.xlarge"
            }
        }
    }
}

```

#### YAML

```yaml

The following template created the DB cluster that you want to add to the global database cluster.
AWSTemplateFormatVersion: 2010-09-09
Parameters:
  username:
    NoEcho: 'true'
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  password:
    NoEcho: 'true'
    Description: Password MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref username
      MasterUserPassword: !Ref password
      DBClusterParameterGroupName: default.aurora-mysql8.0
      Engine: aurora-mysql
      EngineVersion: 8.0.mysql_aurora.8.0.30
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-mysql
      DBClusterIdentifier: !Ref RDSCluster
      DBParameterGroupName: default.aurora-mysql8.0
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r5.xlarge

The following template adds the DB cluster created by the previous template to a global database cluster.
AWSTemplateFormatVersion: 2010-09-09
Parameters:
  GlobalClusterIdentifier:
    Description: Global cluster identifier
    Type: String
  username:
    NoEcho: 'true'
    Description: Username for MySQL database access
    Type: String
    MinLength: '1'
    MaxLength: '16'
    AllowedPattern: '[a-zA-Z][a-zA-Z0-9]*'
    ConstraintDescription: must begin with a letter and contain only alphanumeric characters.
  password:
    NoEcho: 'true'
    Description: Password MySQL database access
    Type: String
    MinLength: '8'
    MaxLength: '41'
    AllowedPattern: '[a-zA-Z0-9]*'
    ConstraintDescription: must contain only alphanumeric characters.
Resources:
  GlobalCluster:
    Type: 'AWS::RDS::GlobalCluster'
    Properties:
      GlobalClusterIdentifier: !Ref GlobalClusterIdentifier
      SourceDBClusterIdentifier: !Ref RDSCluster
  RDSCluster:
    Type: 'AWS::RDS::DBCluster'
    Properties:
      MasterUsername: !Ref username
      MasterUserPassword: !Ref password
      DBClusterParameterGroupName: default.aurora-mysql8.0
      Engine: aurora-mysql
      EngineVersion: 8.0.mysql_aurora.8.0.30
  RDSDBInstance:
    Type: 'AWS::RDS::DBInstance'
    Properties:
      Engine: aurora-mysql
      DBClusterIdentifier: !Ref RDSCluster
      DBParameterGroupName: default.aurora-mysql8.0
      PubliclyAccessible: 'true'
      DBInstanceClass: db.r5.xlarge

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

GlobalEndpoint

All content copied from https://docs.aws.amazon.com/.
