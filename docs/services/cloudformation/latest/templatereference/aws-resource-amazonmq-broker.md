This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker

Creates a broker. Note: This API is asynchronous.

To create a broker, you must either use the `AmazonMQFullAccess` IAM
policy or include the following EC2 permissions in your IAM policy.

- `ec2:CreateNetworkInterface`

This permission is required to allow Amazon MQ to create an elastic network
interface (ENI) on behalf of your account.

- `ec2:CreateNetworkInterfacePermission`

This permission is required to attach the ENI to the broker instance.

- `ec2:DeleteNetworkInterface`

- `ec2:DeleteNetworkInterfacePermission`

- `ec2:DetachNetworkInterface`

- `ec2:DescribeInternetGateways`

- `ec2:DescribeNetworkInterfaces`

- `ec2:DescribeNetworkInterfacePermissions`

- `ec2:DescribeRouteTables`

- `ec2:DescribeSecurityGroups`

- `ec2:DescribeSubnets`

- `ec2:DescribeVpcs`

For more information, see [Create an IAM User and Get Your AWS Credentials](../../../amazon-mq/latest/developer-guide/amazon-mq-setting-up.md#create-iam-user) and [Never Modify or Delete the Amazon MQ Elastic Network Interface](../../../amazon-mq/latest/developer-guide/connecting-to-amazon-mq.md#never-modify-delete-elastic-network-interface) in the
_Amazon MQ Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmazonMQ::Broker",
  "Properties" : {
      "AuthenticationStrategy" : String,
      "AutoMinorVersionUpgrade" : Boolean,
      "BrokerName" : String,
      "Configuration" : ConfigurationId,
      "DataReplicationMode" : String,
      "DataReplicationPrimaryBrokerArn" : String,
      "DeploymentMode" : String,
      "EncryptionOptions" : EncryptionOptions,
      "EngineType" : String,
      "EngineVersion" : String,
      "HostInstanceType" : String,
      "LdapServerMetadata" : LdapServerMetadata,
      "Logs" : LogList,
      "MaintenanceWindowStartTime" : MaintenanceWindow,
      "PubliclyAccessible" : Boolean,
      "SecurityGroups" : [ String, ... ],
      "StorageType" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ TagsEntry, ... ],
      "Users" : [ User, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AmazonMQ::Broker
Properties:
  AuthenticationStrategy: String
  AutoMinorVersionUpgrade: Boolean
  BrokerName: String
  Configuration:
    ConfigurationId
  DataReplicationMode: String
  DataReplicationPrimaryBrokerArn: String
  DeploymentMode: String
  EncryptionOptions:
    EncryptionOptions
  EngineType: String
  EngineVersion: String
  HostInstanceType: String
  LdapServerMetadata:
    LdapServerMetadata
  Logs:
    LogList
  MaintenanceWindowStartTime:
    MaintenanceWindow
  PubliclyAccessible: Boolean
  SecurityGroups:
    - String
  StorageType: String
  SubnetIds:
    - String
  Tags:
    - TagsEntry
  Users:
    - User

```

## Properties

`AuthenticationStrategy`

Optional. The authentication strategy used to secure the broker. The default is
`SIMPLE`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AutoMinorVersionUpgrade`

Enables automatic upgrades to new patch versions for brokers
as new versions are released and supported by Amazon MQ.
Automatic upgrades occur during the scheduled maintenance window
or after a manual broker reboot. Set to `true` by default, if no value is specified.

###### Note

Must be set to `true` for ActiveMQ brokers version 5.18 and above and for RabbitMQ brokers version 3.13 and above.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrokerName`

Required. The broker's name. This value must be unique in your AWS account, 1-50
characters long, must contain only letters, numbers, dashes, and underscores, and
must not contain white spaces, brackets, wildcard characters, or special
characters.

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names.
Broker names are accessible to other AWS services, including CloudWatch Logs. Broker names are not intended to be
used for private or sensitive data.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,50}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

A list of information about the configuration.

_Required_: No

_Type_: [ConfigurationId](aws-properties-amazonmq-broker-configurationid.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DataReplicationMode`

Defines whether this broker is a part of a data replication pair.

_Required_: No

_Type_: String

_Allowed values_: `NONE | CRDR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataReplicationPrimaryBrokerArn`

The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker. Must be set when dataReplicationMode is set to CRDR.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeploymentMode`

Required. The broker's deployment mode.

_Required_: Yes

_Type_: String

_Allowed values_: `SINGLE_INSTANCE | ACTIVE_STANDBY_MULTI_AZ | CLUSTER_MULTI_AZ`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionOptions`

Encryption options for the broker.

_Required_: No

_Type_: [EncryptionOptions](aws-properties-amazonmq-broker-encryptionoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineType`

Required. The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ`.

_Required_: Yes

_Type_: String

_Allowed values_: `ACTIVEMQ | RABBITMQ`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The broker engine version. Defaults to the latest available version for the specified broker engine type. For more information, see the
[ActiveMQ version management](../../../amazon-mq/latest/developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../../../amazon-mq/latest/developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostInstanceType`

Required. The broker's instance type.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LdapServerMetadata`

Optional. The metadata of the LDAP server used to authenticate and authorize
connections to the broker. Does not apply to RabbitMQ brokers.

_Required_: No

_Type_: [LdapServerMetadata](aws-properties-amazonmq-broker-ldapservermetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logs`

Enables Amazon CloudWatch logging for brokers.

_Required_: No

_Type_: [LogList](aws-properties-amazonmq-broker-loglist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceWindowStartTime`

The parameters that determine the WeeklyStartTime.

_Required_: No

_Type_: [MaintenanceWindow](aws-properties-amazonmq-broker-maintenancewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Enables connections from applications outside of the VPC that hosts the
broker's subnets. Set to `false` by default, if no value is provided.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

The list of rules (1 minimum, 125 maximum) that authorize connections to
brokers.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageType`

The broker's storage type.

_Required_: No

_Type_: String

_Allowed values_: `EBS | EFS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones. If you
specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create
VPC endpoints for your broker with multiple subnets in the same Availability Zone.
A SINGLE\_INSTANCE deployment requires one subnet (for example, the default subnet).
An ACTIVE\_STANDBY\_MULTI\_AZ Amazon MQ for ActiveMQ deployment requires two subnets.
A CLUSTER\_MULTI\_AZ Amazon MQ for RabbitMQ deployment has no subnet requirements when deployed with public accessibility. Deployment without public accessibility requires at least one subnet.

###### Important

If you specify subnets in a [shared VPC](../../../vpc/latest/userguide/vpc-sharing.md) for a RabbitMQ broker, the associated VPC to which
the specified subnets belong must be owned by your AWS account. Amazon MQ will not be able to create VPC endpoints in VPCs that are not owned by your AWS account.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Create tags when creating the broker.

_Required_: No

_Type_: Array of [TagsEntry](aws-properties-amazonmq-broker-tagsentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Users`

The list of broker users (persons or applications) who can access queues and topics.
For Amazon MQ for RabbitMQ brokers, one and only one administrative user is
accepted and created when a broker is first provisioned. All subsequent broker users are created by making
RabbitMQ API calls directly to brokers or via the RabbitMQ web console.

When OAuth 2.0 is enabled, the broker accepts one or no users.

_Required_: No

_Type_: Array of [User](aws-properties-amazonmq-broker-user.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon MQ broker ID. For example:

`b-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AmqpEndpoints`

The AMQP endpoints of each broker instance as a list of strings.

`amqp+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:5671`

`Arn`

The Amazon Resource Name (ARN) of the Amazon MQ broker.

`arn:aws:mq:us-east-2:123456789012:broker:MyBroker:b-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

`ConfigurationId`

The unique ID that Amazon MQ generates for the configuration.

`c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

`ConfigurationRevision`

The revision number of the configuration.

`1`

`ConsoleURLs`

Property description not available.

`EngineVersionCurrent`

Property description not available.

`IpAddresses`

The IP addresses of each broker instance as a list of strings. Does not apply to RabbitMQ brokers.

`['198.51.100.2', '203.0.113.9']`

`MqttEndpoints`

The MQTT endpoints of each broker instance as a list of strings.

`mqtt+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:8883`

`OpenWireEndpoints`

The OpenWire endpoints of each broker instance as a list of strings.

`ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61617`

`StompEndpoints`

The STOMP endpoints of each broker instance as a list of strings.

`stomp+ssl://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61614`

`WssEndpoints`

The WSS endpoints of each broker instance as a list of strings.

`wss://b-4aada85d-a80c-4be0-9d30-e344a01b921e-1.mq.eu-central-amazonaws.com:61619`

## Examples

- [Basic Amazon MQ Broker](#aws-resource-amazonmq-broker--examples--Basic_Amazon_MQ_Broker)

- [Complex Amazon MQ Broker](#aws-resource-amazonmq-broker--examples--Complex_Amazon_MQ_Broker)

### Basic Amazon MQ Broker

The following examples creates a basic Amazon MQ broker. The RabbitMQ example creates a broker with one administrative user,
while the ActiveMQ example creates a broker with one user that belongs to a group.

#### JSON

```json

{
  "Description": "Create a basic Amazon MQ for ActiveMQ broker",
  "Resources": {
    "BasicBroker": {
      "Type": "AWS::AmazonMQ::Broker",
      "Properties": {
        "AutoMinorVersionUpgrade": "false",
        "BrokerName": "MyBasicActiveBroker",
        "DeploymentMode": "SINGLE_INSTANCE",
        "EngineType": "ActiveMQ",
        "EngineVersion": "5.15.0",
        "HostInstanceType": "mq.t2.micro",
        "PubliclyAccessible": "true",
        "Users": [
          {
            "ConsoleAccess": "true",
            "Groups": [
              "MyGroup"
            ],
            "Password" : "AmazonMqPassword",
            "Username" : "AmazonMqUsername"
          }
        ]
      }
    }
  }
}
```

#### JSON

```json

{
"Description": "Create a basic Amazon MQ for RabbitMQ broker",
"Resources": {
  "BasicBroker": {
    "Type": "AWS::AmazonMQ::Broker",
    "Properties": {
      "AutoMinorVersionUpgrade": "false",
      "BrokerName": "MyBasicRabbitBroker",
      "DeploymentMode": "SINGLE_INSTANCE",
      "EngineType": "RabbitMQ",
      "EngineVersion": "3.8.6",
      "HostInstanceType": "mq.t3.micro",
      "PubliclyAccessible": "true",
      "Users": [
          {
            "Password" : "AmazonMqPassword",
            "Username" : "AmazonMqUsername"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

---
Description: "Create a basic Amazon MQ for ActiveMQ broker"
Resources:
  BasicBroker:
    Type: "AWS::AmazonMQ::Broker"
    Properties:
      AutoMinorVersionUpgrade: "false"
      BrokerName: MyBasicActiveBroker
      DeploymentMode: SINGLE_INSTANCE
      EngineType: ActiveMQ
      EngineVersion: "5.15.0"
      HostInstanceType: mq.t2.micro
      PubliclyAccessible: "true"
      Users:
        -
          ConsoleAccess: "true"
          Groups:
            - MyGroup
          Password: AmazonMqPassword
          Username: AmazonMqUsername
```

#### YAML

```yaml

---
Description: "Create a basic Amazon MQ for RabbitMQ broker"
Resources:
  BasicBroker:
    Type: "AWS::AmazonMQ::Broker"
    Properties:
      AutoMinorVersionUpgrade: "false"
      BrokerName: MyBasicRabbitBroker
      DeploymentMode: SINGLE_INSTANCE
      EngineType: RabbitMQ
      EngineVersion: "3.8.6"
      HostInstanceType: mq.t3.micro
      PubliclyAccessible: "true"
      Users:
        -
          Password: AmazonMqPassword
          Username: AmazonMqUsername
```

### Complex Amazon MQ Broker

The following example creates a complex Amazon MQ broker. The ActiveMQ example
creates a broker with two users that don't belong to a group and one user that belongs in a group.
The RabbitMQ example creates one administrator user, which can then create and manage other
users via the RabbitMQ web console or the management API.

#### JSON

```json

{
  "Description": "Create a complex, single-instance Amazon MQ for ActiveMQ broker",
  "Resources": {
    "ComplexBroker": {
      "Type": "AWS::AmazonMQ::Broker",
      "Properties": {
        "AutoMinorVersionUpgrade": "false",
        "BrokerName": "MyComplexActiveBroker",
        "Configuration": {
          "Id": { "Ref": "Configuration1" },
          "Revision" : { "Fn::GetAtt": ["Configuration1", "Revision"] }
        },
        "DeploymentMode": "SINGLE_INSTANCE",
        "EngineType": "ActiveMQ",
        "EngineVersion": "5.15.0",
        "HostInstanceType": "mq.t2.micro",
        "Logs": {
            "General": true,
            "Audit": false
        },
        "MaintenanceWindowStartTime": {
          "DayOfWeek": "Monday",
          "TimeOfDay": "22:45",
          "TimeZone": "America/Los_Angeles"
        },
        "PubliclyAccessible": "true",
        "SecurityGroups": [
          "sg-a1b234cd",
          "sg-e5f678gh"
        ],
        "SubnetIds": [
          "subnet-12a3b45c",
          "subnet-67d8e90f"
        ],
        "Users": [{
          "ConsoleAccess": "true",
          "Password" : "AmazonMqPassword",
          "Username" : "AmazonMqUsername"
        }, {
          "Password" : "AmazonMqPassword2",
          "Username" : "AmazonMqUsername2"
        }, {
          "Groups": [
            "MyGroup1",
            "MyGroup2"
          ],
          "Password" : "AmazonMqPassword3",
          "Username" : "AmazonMqUsername3"
        }]
      }
    }
  }
}
```

#### JSON

```json

{
  "Description": "Create a complex, single-instance Amazon MQ RabbitMQ broker without public accessibility",
  "Resources": {
    "ComplexBroker": {
      "Type": "AWS::AmazonMQ::Broker",
      "Properties": {
        "AutoMinorVersionUpgrade": "true",
        "BrokerName": "MyComplexRabbitBroker",
        "DeploymentMode": "SINGLE_INSTANCE",
        "EngineType": "RabbitMQ",
        "EngineVersion": "3.8.6",
        "HostInstanceType": "mq.t3.micro",
        "Logs": {
          "General": true
        },
        "MaintenanceWindowStartTime": {
          "DayOfWeek": "Monday",
          "TimeOfDay": "22:45",
          "TimeZone": "America/Los_Angeles"
        },
        "PubliclyAccessible": "false",
        "SecurityGroups": [
          "sg-1a234b5cd6efgh7i8"
        ],
        "SubnetIds": [
          "subnet-123456b7891abcd1f"
        ],
        "Users": [
          {
            "Password" : "AmazonMqPassword",
            "Username" : "AmazonMqUsername"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Description: Create a complex, single-instance Amazon MQ for ActiveMQ broker
Resources:
  ComplexBroker:
    Type: 'AWS::AmazonMQ::Broker'
    Properties:
      AutoMinorVersionUpgrade: 'false'
      BrokerName: MyComplexActiveBroker
      Configuration:
        Id: !Ref Configuration1
        Revision: !GetAtt
          - Configuration1
          - Revision
      DeploymentMode: SINGLE_INSTANCE
      EngineType: ActiveMQ
      EngineVersion: 5.15.0
      HostInstanceType: mq.t2.micro
      Logs:
        General: true
        Audit: false
      MaintenanceWindowStartTime:
        DayOfWeek: Monday
        TimeOfDay: '22:45'
        TimeZone: America/Los_Angeles
      PubliclyAccessible: 'true'
      SecurityGroups:
        - sg-a1b234cd
        - sg-e5f678gh
      SubnetIds:
        - subnet-12a3b45c
        - subnet-67d8e90f
      Users:
        - ConsoleAccess: 'true'
          Password: AmazonMqPassword
          Username: AmazonMqUsername
        - Password: AmazonMqPassword2
          Username: AmazonMqUsername2
        - Groups:
            - MyGroup1
            - MyGroup2
          Password: AmazonMqPassword3
          Username: AmazonMqUsername3

```

#### YAML

```yaml

Description: Create a single-instance Amazon MQ for RabbitMQ broker without public accessibility
Resources:
  ComplexBroker:
    Type: 'AWS::AmazonMQ::Broker'
    Properties:
      AutoMinorVersionUpgrade: false
      BrokerName: MyComplexRabbitBroker
      DeploymentMode: SINGLE_INSTANCE
      EngineType: RabbitMQ
      EngineVersion: 3.8.6
      HostInstanceType: mq.t3.micro
      Logs:
        General: true
      MaintenanceWindowStartTime:
        DayOfWeek: Monday
        TimeOfDay: '22:45'
        TimeZone: America/Los_Angeles
      PubliclyAccessible: false
      SecurityGroups:
        - 'sg-1a234b5cd6efgh7i8'
      SubnetIds:
        - 'subnet-123456b7891abcd1f'
      Users:
        - Password: AmazonMqPassword
          Username: AmazonMqUsername

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon MQ

ConfigurationId

All content copied from https://docs.aws.amazon.com/.
