# User

###### Note

Does not apply to RabbitMQ brokers.

An ActiveMQ user is a person or an application that can access the queues and
topics of an ActiveMQ broker. For more information, see [User](../developer-guide/user.md) in the Amazon MQ Developer Guide.

A user can belong to a group. You can configure which users belong to which groups
and which groups have permission to send to, receive from, and administer specific
queues and topics.

###### Important

Making changes to a user does not apply the changes to the user immediately. To
apply your changes, you must wait for the next maintenance window or reboot the
broker.

## URI

`/v1/brokers/broker-id/users/username`

## HTTP methods

### GET

**Operation ID:** `DescribeUser`

Returns information about an ActiveMQ user.

Path parametersNameTypeRequiredDescription`username`StringTrue

The username of the ActiveMQ user. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
2-100 characters long.

`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200``

         DescribeUserOutput`

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`404``Error`

HTTP Status Code 404: Resource not found due to incorrect input. Correct your
request and then retry it.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### POST

**Operation ID:** `CreateUser`

Creates an ActiveMQ user.

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames.
Broker usernames are accessible to other AWS services, including CloudWatch Logs. Broker usernames are not intended to be
used for private or sensitive data.

Path parametersNameTypeRequiredDescription`username`StringTrue

The username of the ActiveMQ user. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
2-100 characters long.

`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`404``Error`

HTTP Status Code 404: Resource not found due to incorrect input. Correct your
request and then retry it.

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### PUT

**Operation ID:** `UpdateUser`

Updates the information for an ActiveMQ user.

Path parametersNameTypeRequiredDescription`username`StringTrue

The username of the ActiveMQ user. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
2-100 characters long.

`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`404``Error`

HTTP Status Code 404: Resource not found due to incorrect input. Correct your
request and then retry it.

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### DELETE

**Operation ID:** `DeleteUser`

Deletes an ActiveMQ user.

Path parametersNameTypeRequiredDescription`username`StringTrue

The username of the ActiveMQ user. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
2-100 characters long.

`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`404``Error`

HTTP Status Code 404: Resource not found due to incorrect input. Correct your
request and then retry it.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

Path parametersNameTypeRequiredDescription`username`StringTrue

The username of the ActiveMQ user. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
2-100 characters long.

`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Request bodies

```json

{
  "password": "string",
  "replicationUser": boolean,
  "groups": [
    "string"
  ],
  "consoleAccess": boolean
}
```

```json

{
  "password": "string",
  "replicationUser": boolean,
  "groups": [
    "string"
  ],
  "consoleAccess": boolean
}
```

### Response bodies

```json

{
  "brokerId": "string",
  "replicationUser": boolean,
  "pending": {
    "pendingChange": enum,
    "groups": [
      "string"
    ],
    "consoleAccess": boolean
  },
  "groups": [
    "string"
  ],
  "consoleAccess": boolean,
  "username": "string"
}
```

```json

{
  "errorAttribute": "string",
  "message": "string"
}
```

## Properties

### ChangeType

The type of change pending for the ActiveMQ user.

- `CREATE`

- `UPDATE`

- `DELETE`

### CreateUserInput

Creates a new ActiveMQ user.

PropertyTypeRequiredDescription`consoleAccess`

boolean

False

Enables access to the ActiveMQ Web Console for the ActiveMQ user.

`groups`

Array of type string

False

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long.

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

### DescribeUserOutput

Returns information about an ActiveMQ user.

PropertyTypeRequiredDescription`brokerId`

string

True

Required. The unique ID that Amazon MQ generates for the broker.

`consoleAccess`

boolean

False

Enables access to the the ActiveMQ Web Console for the ActiveMQ user.

`groups`

Array of type string

False

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long.

`pending`

[UserPendingChanges](#brokers-broker-id-users-username-model-userpendingchanges)

False

The status of the changes pending for the ActiveMQ user.

`replicationUser`

boolean

False

Describes whether the user is intended for data replication replication

`username`

string

True

Required. The username of the ActiveMQ user. This value can contain only
alphanumeric characters, dashes, periods, underscores, and tildes (- . \_ ~). This
value must be 2-100 characters long.

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

### UpdateUserInput

Updates the information for an ActiveMQ user.

PropertyTypeRequiredDescription`consoleAccess`

boolean

False

Enables access to the the ActiveMQ Web Console for the ActiveMQ user.

`groups`

Array of type string

False

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long.

`password`

string

Format: password

False

The password of the user. This value must be at least 12 characters long,
must contain at least 4 unique characters, and must not contain commas, colons, or equal signs (,:=).

`replicationUser`

boolean

False

Defines whether the user is intended for data replication replication.

### UserPendingChanges

Returns information about the status of the changes pending for the ActiveMQ
user.

PropertyTypeRequiredDescription`consoleAccess`

boolean

False

Enables access to the the ActiveMQ Web Console for the ActiveMQ user.

`groups`

Array of type string

False

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long.

`pendingChange`

[ChangeType](#brokers-broker-id-users-username-model-changetype)

True

Required. The type of change pending for the ActiveMQ user.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### DescribeUser

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/describeuser.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/describeuser.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/describeuser.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/describeuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/describeuser.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/describeuser.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/describeuser.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/describeuser.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/describeuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/describeuser.md)

### CreateUser

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/createuser.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/createuser.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/createuser.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/createuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/createuser.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/createuser.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/createuser.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/createuser.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/createuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/createuser.md)

### UpdateUser

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/updateuser.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/updateuser.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/updateuser.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/updateuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/updateuser.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/updateuser.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/updateuser.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/updateuser.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/updateuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/updateuser.md)

### DeleteUser

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/deleteuser.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/deleteuser.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/deleteuser.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/deleteuser.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/deleteuser.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/deleteuser.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/deleteuser.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/deleteuser.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/deleteuser.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/deleteuser.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Users

All content copied from https://docs.aws.amazon.com/.
