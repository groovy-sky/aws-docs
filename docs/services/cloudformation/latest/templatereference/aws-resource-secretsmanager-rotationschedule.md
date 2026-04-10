This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::RotationSchedule

Configure the rotation schedule and Lambda rotation function for a
secret. For more information, see [How rotation\
works](../../../secretsmanager/latest/userguide/rotate-secrets-how.md).

For database credentials, refer to the following resources:

- Amazon RDS master user credentials: [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md)

- Amazon Redshift admin user credentials: [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md)

Choose one of the following options for the rotation function:

- Create a new rotation function using `HostedRotationLambda` based on a
[Secrets Manager rotation function template](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md).

- Use an existing rotation function by specifying its ARN with
`RotationLambdaARN`.

###### Important

For database secrets defined in the same CloudFormation template as the
database or service:

1. Use the [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md) resource to populate the
    secret with connection details.

2. Add a `DependsOn` attribute to the `RotationSchedule`
    resource that uses a `SecretTargetAttachment`. This ensures the
    rotation is configured after the secret is populated with connection
    details.

###### Note

You can define only one rotation schedule per secret.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecretsManager::RotationSchedule",
  "Properties" : {
      "ExternalSecretRotationMetadata" : [ ExternalSecretRotationMetadataItem, ... ],
      "ExternalSecretRotationRoleArn" : String,
      "HostedRotationLambda" : HostedRotationLambda,
      "RotateImmediatelyOnUpdate" : Boolean,
      "RotationLambdaARN" : String,
      "RotationRules" : RotationRules,
      "SecretId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecretsManager::RotationSchedule
Properties:
  ExternalSecretRotationMetadata:
    - ExternalSecretRotationMetadataItem
  ExternalSecretRotationRoleArn: String
  HostedRotationLambda:
    HostedRotationLambda
  RotateImmediatelyOnUpdate: Boolean
  RotationLambdaARN: String
  RotationRules:
    RotationRules
  SecretId: String

```

## Properties

`ExternalSecretRotationMetadata`

Property description not available.

_Required_: No

_Type_: Array of [ExternalSecretRotationMetadataItem](aws-properties-secretsmanager-rotationschedule-externalsecretrotationmetadataitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalSecretRotationRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostedRotationLambda`

Creates a new Lambda rotation
function based on one of the [Secrets Manager rotation function templates](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md). To use a rotation function that already
exists, specify `RotationLambdaARN` instead.

You must specify `Transform: AWS::SecretsManager-2024-09-16` at the beginning of the CloudFormation template.
Transforms are macros hosted by AWS CloudFormation that help you create and manage complex infrastructure.
The `Transform: AWS::SecretsManager-2024-09-16` transform automatically extends the CloudFormation stack to include
a nested stack (of type `AWS::CloudFormation::Stack`), which then creates and updates on your behalf during subsequent stack operations,
the appropriate rotation Lambda function for your database or service.
For general information on transforms, see the [AWS CloudFormation documentation.](../userguide/transform-reference.md)

For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md).

For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md).

_Required_: No

_Type_: [HostedRotationLambda](aws-properties-secretsmanager-rotationschedule-hostedrotationlambda.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotateImmediatelyOnUpdate`

Determines whether to rotate the secret immediately or wait until the next scheduled
rotation window when the rotation schedule is updated. The rotation schedule is defined in
`RotationRules`.

The default for `RotateImmediatelyOnUpdate` is `true`. If you
don't specify this value, Secrets Manager rotates the secret immediately.

If you set `RotateImmediatelyOnUpdate` to `false`, Secrets Manager
tests the rotation configuration by running the [`testSecret`\
step](../../../secretsmanager/latest/userguide/rotate-secrets-how.md) of the Lambda rotation function. This test creates an `AWSPENDING`
version of the secret and then removes it.

###### Important

When changing an existing rotation schedule and setting
`RotateImmediatelyOnUpdate` to `false`:

- If using `AutomaticallyAfterDays` or a
`ScheduleExpression` with `rate()`, the previously scheduled
rotation might still occur.

- To prevent unintended rotations, use a `ScheduleExpression` with
`cron()` for granular control over rotation windows.

Rotation is an asynchronous process. For more information, see [How\
rotation works](../../../secretsmanager/latest/userguide/rotate-secrets-how.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationLambdaARN`

The ARN of an existing Lambda rotation function. To specify a rotation function
that is also defined in this template, use the [Ref](../userguide/intrinsic-function-reference-ref.md)
function.

For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](../userguide/aws-properties-rds-dbcluster-masterusersecret.md).

For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](../userguide/aws-resource-redshift-cluster.md).

To create a new rotation function based on one of the [Secrets Manager rotation function templates](../../../secretsmanager/latest/userguide/reference-available-rotation-templates.md), specify `HostedRotationLambda`
instead.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationRules`

A structure that defines the rotation configuration for this secret.

_Required_: No

_Type_: [RotationRules](aws-properties-secretsmanager-rotationschedule-rotationrules.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretId`

The ARN or name of the secret to rotate. This is unique for each rotation schedule definition.

To reference a secret also created in this template, use the [Ref](../userguide/intrinsic-function-reference-ref.md)
function with the secret's logical ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of an `AWS::SecretsManager::RotationSchedule`
resource to the intrinsic `Ref` function, the function returns the ARN of the
secret being configured, such as:

_arn:aws:secretsmanager:_
_us-west-2_: _123456789012_:secret: _my-path/my-secret-name_- _1a2b3c_

You can use the ARN to reference a secret you create in one part of the stack template
from within the definition of another resource later, in the same template. You typically do
this when you define the [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md) resource type.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

- [Automatic rotation with a cron expression](#aws-resource-secretsmanager-rotationschedule--examples--Automatic_rotation_with_a_cron_expression)

- [Automatic rotation with a rate expression](#aws-resource-secretsmanager-rotationschedule--examples--Automatic_rotation_with_a_rate_expression)

- [DocumentDB secret rotation example](#aws-resource-secretsmanager-rotationschedule--examples--DocumentDB_secret_rotation_example)

### Automatic rotation with a cron expression

The following example rotates a secret every day between 1:00 AM and 3:00 AM
UTC.

#### JSON

```json

"MySecretRotationSchedule": {
  "Type": "AWS::SecretsManager::RotationSchedule",
  "DependsOn": "MyRotationLambda",
  "Properties": {
    "SecretId": {"Ref": "MySecret"},
    "RotationLambdaARN": {"Fn::GetAtt": "MyRotationLambda.Arn"},
    "RotationRules": {
      "Duration": "2h",
      "ScheduleExpression": "cron(0 1 * * ? *)"
    }
  }
}
```

#### YAML

```yaml

MySecretRotationSchedule:
  Type: AWS::SecretsManager::RotationSchedule
  DependsOn: MyRotationLambda
  Properties:
    SecretId: !Ref MySecret
    RotationLambdaARN: !GetAtt MyRotationLambda.Arn
    RotationRules:
      Duration: 2h
      ScheduleExpression: 'cron(0 1 * * ? *)'
```

### Automatic rotation with a rate expression

The following example rotates a secret between midnight and 6:00 AM UTC every 10
days.

#### JSON

```json

"MySecretRotationSchedule": {
  "Type": "AWS::SecretsManager::RotationSchedule",
  "DependsOn": "MyRotationLambda",
  "Properties": {
  "SecretId": {"Ref": "MySecret"},
    "RotationLambdaARN": {"Fn::GetAtt": "MyRotationLambda.Arn"},
    "RotationRules": {
      "Duration": "6h",
      "ScheduleExpression": "rate(10 days)"
      }
    }
  }
```

#### YAML

```yaml

MySecretRotationSchedule:
  Type: AWS::SecretsManager::RotationSchedule
  DependsOn: MyRotationLambda
  Properties:
    SecretId: !Ref MySecret
    RotationLambdaARN: !GetAtt MyRotationLambda.Arn
    RotationRules:
      Duration: 6h
      ScheduleExpression: 'rate(10 days)'
```

### DocumentDB secret rotation example

The following example creates a DocumentDB database instance and a secret with
credentials. The secret is configured to rotate on the first Sunday of every month between
4:00 AM and 6:00 AM UTC.

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Transform":"AWS::SecretsManager-2024-09-16",
   "Resources":{
      "TestVPC":{
         "Type":"AWS::EC2::VPC",
         "Properties":{
            "CidrBlock":"10.0.0.0/16",
            "EnableDnsHostnames":true,
            "EnableDnsSupport":true
         }
      },
      "TestSubnet01":{
         "Type":"AWS::EC2::Subnet",
         "Properties":{
            "CidrBlock":"10.0.96.0/19",
            "AvailabilityZone":{
               "Fn::Select":[
                  "0",
                  {
                     "Fn::GetAZs":{
                        "Ref":"AWS::Region"
                     }
                  }
               ]
            },
            "VpcId":{
               "Ref":"TestVPC"
            }
         }
      },
      "TestSubnet02":{
         "Type":"AWS::EC2::Subnet",
         "Properties":{
            "CidrBlock":"10.0.128.0/19",
            "AvailabilityZone":{
               "Fn::Select":[
                  "1",
                  {
                     "Fn::GetAZs":{
                        "Ref":"AWS::Region"
                     }
                  }
               ]
            },
            "VpcId":{
               "Ref":"TestVPC"
            }
         }
      },
      "SecretsManagerVPCEndpoint":{
         "Type":"AWS::EC2::VPCEndpoint",
         "Properties":{
            "SubnetIds":[
               {
                  "Ref":"TestSubnet01"
               },
               {
                  "Ref":"TestSubnet02"
               }
            ],
            "SecurityGroupIds":[
               {
                  "Fn::GetAtt":[
                     "TestVPC",
                     "DefaultSecurityGroup"
                  ]
               }
            ],
            "VpcEndpointType":"Interface",
            "ServiceName":{
               "Fn::Sub":"com.amazonaws.${AWS::Region}.secretsmanager"
            },
            "PrivateDnsEnabled":true,
            "VpcId":{
               "Ref":"TestVPC"
            }
         }
      },
      "MyDocDBClusterRotationSecret":{
         "Type":"AWS::SecretsManager::Secret",
         "Properties":{
            "GenerateSecretString":{
        "SecretStringTemplate":"{"username": "someadmin","ssl": true}",
               "GenerateStringKey":"password",
               "PasswordLength":16,
               "ExcludeCharacters":"\"@/\\"
            },
            "Tags":[
               {
                  "Key":"AppName",
                  "Value":"MyApp"
               }
            ]
         }
      },
      "MyDocDBCluster":{
         "Type":"AWS::DocDB::DBCluster",
         "Properties":{
            "DBSubnetGroupName":{
               "Ref":"MyDBSubnetGroup"
            },
            "MasterUsername":{
               "Fn::Sub":"{{resolve:secretsmanager:${MyDocDBClusterRotationSecret}::username}}"
            },
        "MasterUserPassword":{
               "Fn::Sub":"{{resolve:secretsmanager:${MyDocDBClusterRotationSecret}::password}}"
            },
            "VpcSecurityGroupIds":[
               {
                  "Fn::GetAtt":[
                     "TestVPC",
                     "DefaultSecurityGroup"
                  ]
               }
            ]
         }
      },
      "DocDBInstance":{
         "Type":"AWS::DocDB::DBInstance",
         "Properties":{
            "DBClusterIdentifier":{
               "Ref":"MyDocDBCluster"
            },
            "DBInstanceClass":"db.r5.large"
         }
      },
      "MyDBSubnetGroup":{
         "Type":"AWS::DocDB::DBSubnetGroup",
         "Properties":{
            "DBSubnetGroupDescription":"DescriptionS",
            "SubnetIds":[
               {
                  "Ref":"TestSubnet01"
               },
               {
                  "Ref":"TestSubnet02"
               }
            ]
         }
      },
      "SecretDocDBClusterAttachment":{
         "Type":"AWS::SecretsManager::SecretTargetAttachment",
         "Properties":{
            "SecretId":{
               "Ref":"MyDocDBClusterRotationSecret"
            },
            "TargetId":{
               "Ref":"MyDocDBCluster"
            },
            "TargetType":"AWS::DocDB::DBCluster"
         }
      },
      "MySecretRotationSchedule":{
         "Type":"AWS::SecretsManager::RotationSchedule",
         "DependsOn":"SecretDocDBClusterAttachment",
         "Properties":{
            "SecretId":{
               "Ref":"MyDocDBClusterRotationSecret"
            },
            "HostedRotationLambda":{
               "RotationType":"MongoDBSingleUser",
               "RotationLambdaName":"MongoDBSingleUser",
               "VpcSecurityGroupIds":{
                  "Fn::GetAtt":[
                     "TestVPC",
                     "DefaultSecurityGroup"
                  ]
               },
               "VpcSubnetIds":{
                  "Fn::Join":[
                     ",",
                     [
                        {
                           "Ref":"TestSubnet01"
                        },
                        {
                           "Ref":"TestSubnet02"
                        }
                     ]
                  ]
               }
            },
            "RotationRules":{
              "Duration": "2h",
              "ScheduleExpression": "cron(0 4 ? * SUN#1 *)"
            }
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::SecretsManager-2024-09-16
Resources:
  TestVPC:
    Type: AWS::EC2::VPC
    Properties:
      CidrBlock: 10.0.0.0/16
      EnableDnsHostnames: true
      EnableDnsSupport: true
  TestSubnet01:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 10.0.96.0/19
      AvailabilityZone:
        Fn::Select:
        - '0'
        - Fn::GetAZs:
            Ref: AWS::Region
      VpcId:
        Ref: TestVPC
  TestSubnet02:
    Type: AWS::EC2::Subnet
    Properties:
      CidrBlock: 10.0.128.0/19
      AvailabilityZone:
        Fn::Select:
        - '1'
        - Fn::GetAZs:
            Ref: AWS::Region
      VpcId:
        Ref: TestVPC
  SecretsManagerVPCEndpoint:
    Type: AWS::EC2::VPCEndpoint
    Properties:
      SubnetIds:
      - Ref: TestSubnet01
      - Ref: TestSubnet02
      SecurityGroupIds:
      - Fn::GetAtt:
        - TestVPC
        - DefaultSecurityGroup
      VpcEndpointType: Interface
      ServiceName:
        Fn::Sub: com.amazonaws.${AWS::Region}.secretsmanager
      PrivateDnsEnabled: true
      VpcId:
        Ref: TestVPC
  MyDocDBClusterRotationSecret:
    Type: AWS::SecretsManager::Secret
    Properties:
      GenerateSecretString:
        SecretStringTemplate: '{"username": "someadmin","ssl": true}'
        GenerateStringKey: password
        PasswordLength: 16
        ExcludeCharacters: "\"@/\\"
      Tags:
      - Key: AppName
        Value: MyApp
  MyDocDBCluster:
    Type: AWS::DocDB::DBCluster
    Properties:
      DBSubnetGroupName:
        Ref: MyDBSubnetGroup
      MasterUsername:
        Fn::Sub: "{{resolve:secretsmanager:${MyDocDBClusterRotationSecret}::username}}"
      MasterUserPassword:
        Fn::Sub: "{{resolve:secretsmanager:${MyDocDBClusterRotationSecret}::password}}"
      VpcSecurityGroupIds:
      - Fn::GetAtt:
        - TestVPC
        - DefaultSecurityGroup
  DocDBInstance:
    Type: AWS::DocDB::DBInstance
    Properties:
      DBClusterIdentifier:
        Ref: MyDocDBCluster
      DBInstanceClass: db.r5.large
  MyDBSubnetGroup:
    Type: AWS::DocDB::DBSubnetGroup
    Properties:
      DBSubnetGroupDescription: 'Description'
      SubnetIds:
      - Ref: TestSubnet01
      - Ref: TestSubnet02
  SecretDocDBClusterAttachment:
    Type: AWS::SecretsManager::SecretTargetAttachment
    Properties:
      SecretId:
        Ref: MyDocDBClusterRotationSecret
      TargetId:
        Ref: MyDocDBCluster
      TargetType: AWS::DocDB::DBCluster
  MySecretRotationSchedule:
    Type: AWS::SecretsManager::RotationSchedule
    DependsOn: SecretDocDBClusterAttachment
    Properties:
      SecretId:
        Ref: MyDocDBClusterRotationSecret
      HostedRotationLambda:
        RotationType: MongoDBSingleUser
        RotationLambdaName: MongoDBSingleUser
        VpcSecurityGroupIds:
          Fn::GetAtt:
          - TestVPC
          - DefaultSecurityGroup
        VpcSubnetIds:
          Fn::Join:
          - ","
          - - Ref: TestSubnet01
            - Ref: TestSubnet02
      RotationRules:
        Duration: 2h
        ScheduleExpression: 'cron(0 4 ? * SUN#1 *)'

```

## See also

- [RotateSecret](../../../../reference/secretsmanager/latest/apireference/api-rotatesecret.md) in
the AWS Secrets Manager API Reference

- [Rotate secrets](../../../secretsmanager/latest/userguide/rotating-secrets.md) in the AWS Secrets Manager User Guide

- [AWS::SecretsManager::RotationSchedule HostedRotationLambda](../userguide/aws-properties-secretsmanager-rotationschedule-hostedrotationlambda.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecretsManager::ResourcePolicy

ExternalSecretRotationMetadataItem

All content copied from https://docs.aws.amazon.com/.
