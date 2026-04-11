# Brokers

This is a collection of brokers. A broker is a message broker environment running
on Amazon MQ. It is the basic building block of Amazon MQ. For more information, see
[Broker instance types](../developer-guide/broker-instance-types.md) in the _Amazon MQ Developer Guide_.

## URI

`/v1/brokers`

## HTTP methods

### GET

**Operation ID:** `ListBrokers`

Returns a list of all brokers.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

ResponsesStatus codeResponse modelDescription`200``

         ListBrokersOutput`

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### POST

**Operation ID:** `CreateBroker`

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

For more information, see [Create an IAM User and Get Your AWS Credentials](../developer-guide/amazon-mq-setting-up.md#create-iam-user) and [Never Modify or Delete the Amazon MQ Elastic Network Interface](../developer-guide/connecting-to-amazon-mq.md#never-modify-delete-elastic-network-interface) in the
_Amazon MQ Developer Guide_.

ResponsesStatus codeResponse modelDescription`200``

         CreateBrokerOutput`

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`401``Error`

HTTP Status Code 401: Unauthorized request. The provided credentials couldn't be
validated.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Request bodies

```json

{
  "engineVersion": "string",
  "deploymentMode": enum,
  "maintenanceWindowStartTime": {
    "dayOfWeek": enum,
    "timeZone": "string",
    "timeOfDay": "string"
  },
  "configuration": {
    "id": "string",
    "revision": integer
  },
  "authenticationStrategy": enum,
  "engineType": enum,
  "hostInstanceType": "string",
  "users": [
    {
      "password": "string",
      "replicationUser": boolean,
      "groups": [
        "string"
      ],
      "consoleAccess": boolean,
      "username": "string"
    }
  ],
  "tags": {
  },
  "dataReplicationMode": enum,
  "creatorRequestId": "string",
  "storageConfiguration": {
    "efs": {
      "throughputMode": enum
    }
  },
  "encryptionOptions": {
    "useAwsOwnedKey": boolean,
    "kmsKeyId": "string"
  },
  "publiclyAccessible": boolean,
  "storageType": enum,
  "securityGroups": [
    "string"
  ],
  "dataReplicationPrimaryBrokerArn": "string",
  "brokerName": "string",
  "logs": {
    "general": boolean,
    "audit": boolean
  },
  "ldapServerMetadata": {
    "roleSearchMatching": "string",
    "serviceAccountPassword": "string",
    "roleBase": "string",
    "hosts": [
      "string"
    ],
    "roleName": "string",
    "userBase": "string",
    "roleSearchSubtree": boolean,
    "serviceAccountUsername": "string",
    "userRoleName": "string",
    "userSearchMatching": "string",
    "userSearchSubtree": boolean
  },
  "autoMinorVersionUpgrade": boolean,
  "subnetIds": [
    "string"
  ]
}
```

### Response bodies

```json

{
  "brokerSummaries": [
    {
      "brokerArn": "string",
      "brokerId": "string",
      "deploymentMode": enum,
      "created": "string",
      "brokerState": enum,
      "engineType": enum,
      "brokerName": "string",
      "hostInstanceType": "string"
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "brokerArn": "string",
  "brokerId": "string"
}
```

```json

{
  "errorAttribute": "string",
  "message": "string"
}
```

## Properties

### AuthenticationStrategy

Optional. The authentication strategy used to secure the broker. The
default is `SIMPLE`.

- `SIMPLE`

- `LDAP`

- `CONFIG_MANAGED`

### BrokerState

The broker's status.

- `CREATION_IN_PROGRESS`

- `CREATION_FAILED`

- `DELETION_IN_PROGRESS`

- `RUNNING`

- `REBOOT_IN_PROGRESS`

- `CRITICAL_ACTION_REQUIRED`

- `REPLICA`

### BrokerStorageConfiguration

The storage configurations of a broker.

PropertyTypeRequiredDescription`efs`

[EfsBrokerStorageConfiguration](#brokers-model-efsbrokerstorageconfiguration)

False

### BrokerStorageType

The broker's storage type.

###### Important

`EFS` is not supported for RabbitMQ engine type.

- `EBS`

- `EFS`

### BrokerSummary

Returns information about all brokers.

PropertyTypeRequiredDescription`brokerArn`

string

False

The broker's Amazon Resource Name (ARN).

`brokerId`

string

False

The unique ID that Amazon MQ generates for the broker.

`brokerName`

string

False

The broker's name. This value is unique in your AWS account, 1-50 characters
long, and containing only letters, numbers, dashes, and underscores, and must not
contain white spaces, brackets, wildcard characters, or special characters.

`brokerState`

[BrokerState](#brokers-model-brokerstate)

False

The broker's status.

`created`

string

Format: date-time

False

The time when the broker was created.

`deploymentMode`

[DeploymentMode](#brokers-model-deploymentmode)

True

The broker's deployment mode.

`engineType`

[EngineType](#brokers-model-enginetype)

True

The type of broker engine.

`hostInstanceType`

string

False

The broker's instance type.

### ConfigurationId

A list of information about the configuration.

PropertyTypeRequiredDescription`id`

string

True

Required. The unique ID that Amazon MQ generates for the configuration.

`revision`

integer

False

The revision number of the configuration.

### CreateBrokerInput

Creates a broker.

PropertyTypeRequiredDescription`authenticationStrategy`

[AuthenticationStrategy](#brokers-model-authenticationstrategy)

False

Optional. The authentication strategy used to secure the broker. The default is
`SIMPLE`.

`autoMinorVersionUpgrade`

boolean

False

Enables automatic upgrades to new patch versions for brokers
as new versions are released and supported by Amazon MQ.
Automatic upgrades occur during the scheduled maintenance window
or after a manual broker reboot. Set to `true` by default, if no value is specified.

###### Note

Must be set to `true` for ActiveMQ brokers version 5.18 and above and for RabbitMQ brokers version 3.13 and above.

`brokerName`

string

True

Required. The broker's name. This value must be unique in your AWS account, 1-50
characters long, must contain only letters, numbers, dashes, and underscores, and
must not contain white spaces, brackets, wildcard characters, or special
characters.

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names.
Broker names are accessible to other AWS services, including CloudWatch Logs. Broker names are not intended to be
used for private or sensitive data.

`configuration`

[ConfigurationId](#brokers-model-configurationid)

False

A list of information about the configuration.

`creatorRequestId`

string

False

The unique ID that the requester receives for the created broker. Amazon MQ passes
your ID with the API action.

###### Note

We recommend using a Universally Unique Identifier
(UUID) for the creatorRequestId. You may omit the `creatorRequestId` if your
application doesn't require idempotency.

`dataReplicationMode`

[DataReplicationMode](#brokers-model-datareplicationmode)

False

Defines whether this broker is a part of a data replication pair.

`dataReplicationPrimaryBrokerArn`

string

False

The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker. Must be set when dataReplicationMode is set to CRDR.

`deploymentMode`

[DeploymentMode](#brokers-model-deploymentmode)

True

Required. The broker's deployment mode.

`encryptionOptions`

[EncryptionOptions](#brokers-model-encryptionoptions)

False

Encryption options for the broker.

`engineType`

[EngineType](#brokers-model-enginetype)

True

Required. The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ`.

`engineVersion`

string

False

The broker engine version. Defaults to the latest available version for the specified broker engine type. For more information, see the
[ActiveMQ version management](../developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

`hostInstanceType`

string

True

Required. The broker's instance type.

`ldapServerMetadata`

[LdapServerMetadataInput](#brokers-model-ldapservermetadatainput)

False

Optional. The metadata of the LDAP server used to authenticate and authorize
connections to the broker. Does not apply to RabbitMQ brokers.

`logs`

[Logs](#brokers-model-logs)

False

Enables Amazon CloudWatch logging for brokers.

`maintenanceWindowStartTime`

[WeeklyStartTime](#brokers-model-weeklystarttime)

False

The parameters that determine the WeeklyStartTime.

`publiclyAccessible`

boolean

True

Enables connections from applications outside of the VPC that hosts the
broker's subnets. Set to `false` by default, if no value is provided.

`securityGroups`

Array of type string

False

The list of rules (1 minimum, 125 maximum) that authorize connections to
brokers.

`storageConfiguration`

[BrokerStorageConfiguration](#brokers-model-brokerstorageconfiguration)

False

The broker's storage configuration.

`storageType`

[BrokerStorageType](#brokers-model-brokerstoragetype)

False

The broker's storage type.

`subnetIds`

Array of type string

False

The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones. If you
specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create
VPC endpoints for your broker with multiple subnets in the same Availability Zone.
A SINGLE\_INSTANCE deployment requires one subnet (for example, the default subnet).
An ACTIVE\_STANDBY\_MULTI\_AZ Amazon MQ for ActiveMQ deployment requires two subnets.
A CLUSTER\_MULTI\_AZ Amazon MQ for RabbitMQ deployment has no subnet requirements when deployed with public accessibility. Deployment without public accessibility requires at least one subnet.

###### Important

If you specify subnets in a [shared VPC](../../../vpc/latest/userguide/vpc-sharing.md) for a RabbitMQ broker, the associated VPC to which
the specified subnets belong must be owned by your AWS account. Amazon MQ will not be able to create VPC endpoints in VPCs that are not owned by your AWS account.

`tags`

object

False

Create tags when creating the broker.

`users`

Array of type User

False

The list of broker users (persons or applications) who can access queues and topics.
For Amazon MQ for RabbitMQ brokers, one and only one administrative user is
accepted and created when a broker is first provisioned. All subsequent broker users are created by making
RabbitMQ API calls directly to brokers or via the RabbitMQ web console.

When OAuth 2.0 is enabled, the broker accepts one or no users.

### CreateBrokerOutput

Returns information about the created broker.

PropertyTypeRequiredDescription`brokerArn`

string

False

The broker's Amazon Resource Name (ARN).

`brokerId`

string

False

The unique ID that Amazon MQ generates for the broker.

### DataReplicationMode

Specifies whether a broker is a part of a data replication pair.

- `NONE`

- `CRDR`

### DeploymentMode

The broker's deployment mode.

- `SINGLE_INSTANCE`

- `ACTIVE_STANDBY_MULTI_AZ`

- `CLUSTER_MULTI_AZ`

### EfsBrokerStorageConfiguration

The storage configuration of an EFS broker storage

PropertyTypeRequiredDescription`throughputMode`

string

Values: `ELASTIC | STANDARD`

False

Throughput mode of EFS broker storage

### EncryptionOptions

Encryption options for the broker.

PropertyTypeRequiredDescription`kmsKeyId`

string

False

The customer master key (CMK) to use for the A AWS KMS (KMS).
This key is used to encrypt your data at rest. If not provided, Amazon MQ will use a
default CMK to encrypt your data.

`useAwsOwnedKey`

boolean

True

Enables the use of an AWS owned CMK using AWS KMS (KMS). Set to `true` by default, if no value is provided, for example,
for RabbitMQ brokers.

### EngineType

The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.

- `ACTIVEMQ`

- `RABBITMQ`

### Error

Returns information about an error.

PropertyTypeRequiredDescription`errorAttribute`

string

False

The attribute which caused the error.

`message`

string

False

The explanation of the error.

### LdapServerMetadataInput

Optional. The metadata of the LDAP server used to authenticate and authorize connections to
the broker.

###### Important

Does not apply to RabbitMQ brokers.

PropertyTypeRequiredDescription`hosts`

Array of type string

True

Specifies the location of the LDAP server such as AWS Directory Service for Microsoft Active Directory. Optional failover server.

`roleBase`

string

True

The distinguished name of the node in the directory information tree (DIT) to
search for roles or groups. For example, `ou=group, ou=corp, dc=corp,
                  dc=example, dc=com`.

`roleName`

string

False

Specifies the LDAP attribute that identifies the group name attribute in the
object returned from the group membership query.

`roleSearchMatching`

string

True

The LDAP search filter used to find roles within the `roleBase`. The
distinguished name of the user matched by `userSearchMatching` is
substituted into the `{0}` placeholder in the search filter. The client's
username is substituted into the `{1}` placeholder. For example, if you
set this option to `(member=uid={1})` for the user `janedoe`,
the search filter becomes `(member=uid=janedoe)` after string
substitution. It matches all role entries that have a member attribute equal to
`uid=janedoe` under the subtree selected by the
`roleBase`.

`roleSearchSubtree`

boolean

False

The directory search scope for the role. If set to true, scope is to search the
entire subtree.

`serviceAccountPassword`

string

True

Service account password. A service account is an account in your LDAP server that
has access to initiate a connection. For example, `cn=admin,dc=corp, dc=example,
                  dc=com`.

`serviceAccountUsername`

string

True

Service account username. A service account is an account in your LDAP server that
has access to initiate a connection. For example, `cn=admin,dc=corp, dc=example,
                  dc=com`.

`userBase`

string

True

Select a particular subtree of the directory information tree (DIT) to search for
user entries. The subtree is specified by a DN, which specifies the base node of the
subtree. For example, by setting this option to `ou=Users,ou=corp, dc=corp,
                  dc=example, dc=com`, the search for user entries is restricted to the
subtree beneath `ou=Users, ou=corp, dc=corp, dc=example, dc=com`.

`userRoleName`

string

False

Specifies the name of the LDAP attribute for the user group membership.

`userSearchMatching`

string

True

The LDAP search filter used to find users within the `userBase`. The
client's username is substituted into the `{0}` placeholder in the search
filter. For example, if this option is set to `(uid={0})` and the received
username is `janedoe`, the search filter becomes
`(uid=janedoe)` after string substitution. It will result in matching
an entry like `uid=janedoe, ou=Users,ou=corp, dc=corp, dc=example,
                  dc=com`.

`userSearchSubtree`

boolean

False

The directory search scope for the user. If set to true, scope is to search the
entire subtree.

### ListBrokersOutput

A list of information about all brokers.

PropertyTypeRequiredDescription`brokerSummaries`

Array of type BrokerSummary

False

A list of information about all brokers.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

### Logs

The list of information about logs to be enabled for the specified broker.

PropertyTypeRequiredDescription`audit`

boolean

False

Enables audit logging. Every user management action made using JMX or the ActiveMQ
Web Console is logged. Does not apply to RabbitMQ brokers.

`general`

boolean

False

Enables general logging.

### User

A user associated with the broker. For Amazon MQ for RabbitMQ brokers, one and only one administrative
user is accepted and created when a broker is first provisioned.
All subsequent broker users are created by making RabbitMQ API calls directly to brokers or via the RabbitMQ web console.

PropertyTypeRequiredDescription`consoleAccess`

boolean

False

Enables access to the ActiveMQ Web Console for the ActiveMQ user. Does not apply to RabbitMQ brokers.

`groups`

Array of type string

False

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long. Does not apply to RabbitMQ brokers.

`password`

string

Format: password

True

Required. The password of the user. This value must be at least 12 characters
long, must contain at least 4 unique characters, and must not contain commas, colons, or equal signs (,:=).

`replicationUser`

boolean

False

Defines if this user is intended for CRDR replication purposes.

`username`

string

True

The username of the broker user. The following restrictions apply to broker usernames:

- For Amazon MQ for ActiveMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_ ~).
This value must be 2-100 characters long.

- For Amazon MQ for RabbitMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores (- . \_).
This value must not contain a tilde (~) character. Amazon MQ prohibts using `guest` as a valid usename.
This value must be 2-100 characters long.

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames.
Broker usernames are accessible to other AWS services, including CloudWatch Logs. Broker usernames are not intended to be
used for private or sensitive data.

### WeeklyStartTime

The scheduled time period relative to UTC during which Amazon MQ begins to apply
pending updates or patches to the broker.

PropertyTypeRequiredDescription`dayOfWeek`

string

Values: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

True

Required. The day of the week.

`timeOfDay`

string

True

Required. The time, in 24-hour format.

`timeZone`

string

False

The time zone, UTC by default, in either the Country/City format, or the UTC
offset format.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### ListBrokers

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/listbrokers.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/listbrokers.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/listbrokers.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/listbrokers.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/listbrokers.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/listbrokers.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/listbrokers.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/listbrokers.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/listbrokers.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/listbrokers.md)

### CreateBroker

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/createbroker.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/createbroker.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/createbroker.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/createbroker.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/createbroker.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/createbroker.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/createbroker.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/createbroker.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/createbroker.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/createbroker.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Broker Reboot

Brokers broker-id Promote

All content copied from https://docs.aws.amazon.com/.
