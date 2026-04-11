# Users

###### Note

Does not apply to RabbitMQ brokers.

This is a collection of ActiveMQ users for the specified broker. An ActiveMQ user
is a person or an application that can access the queues and topics of an ActiveMQ
broker. For more information, see [User](../developer-guide/user.md) in the Amazon MQ Developer Guide.

You can configure to have specific permissions. For example, you can allow some
users to access the [ActiveMQ\
Web Console](http://activemq.apache.org/web-console.html).

A user can belong to a group. You can configure which users belong to which groups
and which groups have permission to send to, receive from, and administer specific
queues and topics.

###### Important

Making changes to a user does not apply the changes to the user immediately. To
apply your changes, you must wait for the next maintenance window or reboot the
broker.

## URI

`/v1/brokers/broker-id/users`

## HTTP methods

### GET

**Operation ID:** `ListUsers`

Returns a list of all ActiveMQ users.

Path parametersNameTypeRequiredDescription`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

ResponsesStatus codeResponse modelDescription`200``

         ListUsersOutput`

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

Path parametersNameTypeRequiredDescription`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Response bodies

```json

{
  "brokerId": "string",
  "nextToken": "string",
  "maxResults": integer,
  "users": [
    {
      "pendingChange": enum,
      "username": "string"
    }
  ]
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

### ListUsersOutput

Returns a list of all ActiveMQ users.

PropertyTypeRequiredDescription`brokerId`

string

True

Required. The unique ID that Amazon MQ generates for the broker.

`maxResults`

integer

Minimum: 5

Maximum: 100

True

Required. The maximum number of ActiveMQ users that can be returned per page (20
by default). This value must be an integer from 5 to 100.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`users`

Array of type UserSummary

True

Required. The list of all ActiveMQ usernames for the specified broker. Does not apply to RabbitMQ brokers.

### UserSummary

Returns a list of all broker users. Does not apply to RabbitMQ brokers.

PropertyTypeRequiredDescription`pendingChange`

[ChangeType](#brokers-broker-id-users-model-changetype)

False

The type of change pending for the broker user.

`username`

string

True

Required. The username of the broker user.
This value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_ ~).
This value must be 2-100 characters long.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### ListUsers

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/listusers.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/listusers.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/listusers.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/listusers.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/listusers.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/listusers.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/listusers.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/listusers.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/listusers.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/listusers.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User

Operations

All content copied from https://docs.aws.amazon.com/.
