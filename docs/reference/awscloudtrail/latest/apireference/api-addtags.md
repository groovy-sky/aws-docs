# AddTags

Adds one or more tags to a trail, event data store, dashboard, or channel, up to a limit of 50. Overwrites an
existing tag's value when a new value is specified for an existing tag key. Tag key names
must be unique; you cannot have two keys with the same name but different
values. If you specify a key without a value, the tag will be created with the specified
key and a value of null. You can tag a trail or event data store that applies to all
AWS Regions only from the Region in which the trail or event data store
was created (also known as its home Region).

## Request Syntax

```nohighlight

{
   "ResourceId": "string",
   "TagsList": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceId](#API_AddTags_RequestSyntax)**

Specifies the ARN of the trail, event data store, dashboard, or channel to which one or more tags will be
added.

The format of a trail ARN is:
`arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail`

The format of an event data store ARN is:
`arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE`

The format of a dashboard ARN is: `arn:aws:cloudtrail:us-east-1:123456789012:dashboard/exampleDash`

The format of a channel ARN is:
`arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890`

Type: String

Required: Yes

**[TagsList](#API_AddTags_RequestSyntax)**

Contains a list of tags, up to a limit of 50

Type: Array of [Tag](api-tag.md) objects

Array Members: Maximum number of 200 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ChannelARNInvalidException**

This exception is thrown when the specified value of `ChannelARN` is not
valid.

HTTP Status Code: 400

**ChannelNotFoundException**

This exception is thrown when CloudTrail cannot find the specified channel.

HTTP Status Code: 400

**CloudTrailARNInvalidException**

This exception is thrown when an operation is called with an ARN that is not valid.

The following is the format of a trail ARN: `arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail`

The following is the format of an event data store ARN:
`arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE`

The following is the format of a dashboard ARN: `arn:aws:cloudtrail:us-east-1:123456789012:dashboard/exampleDash`

The following is the format of a channel ARN:
`arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890`

HTTP Status Code: 400

**ConflictException**

This exception is thrown when the specified resource is not ready for an operation. This
can occur when you try to run an operation on a resource before CloudTrail has time
to fully load the resource, or because another operation is modifying the resource. If this exception occurs, wait a few minutes, and then try the
operation again.

HTTP Status Code: 400

**EventDataStoreARNInvalidException**

The specified event data store ARN is not valid or does not map to an event data store
in your account.

HTTP Status Code: 400

**EventDataStoreNotFoundException**

The specified event data store was not found.

HTTP Status Code: 400

**InactiveEventDataStoreException**

The event data store is inactive.

HTTP Status Code: 400

**InvalidTagParameterException**

This exception is thrown when the specified tag key or values are not valid. It can also
occur if there are duplicate tags or too many tags on the resource.

HTTP Status Code: 400

**InvalidTrailNameException**

This exception is thrown when the provided trail name is not valid. Trail names must
meet the following requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
(\_), or dashes (-)

- Start with a letter or number, and end with a letter or number

- Be between 3 and 128 characters

- Have no adjacent periods, underscores or dashes. Names like
`my-_namespace` and `my--namespace` are not valid.

- Not be in IP address format (for example, 192.168.5.4)

HTTP Status Code: 400

**NoManagementAccountSLRExistsException**

This exception is thrown when the management account does not have a service-linked
role.

HTTP Status Code: 400

**NotOrganizationMasterAccountException**

This exception is thrown when the AWS account making the request to
create or update an organization trail or event data store is not the management account
for an organization in AWS Organizations. For more information, see [Prepare For Creating a Trail For Your Organization](../../../../services/awscloudtrail/latest/userguide/creating-an-organizational-trail-prepare.md) or [Organization event data stores](../../../../services/awscloudtrail/latest/userguide/cloudtrail-lake-organizations.md).

HTTP Status Code: 400

**OperationNotPermittedException**

This exception is thrown when the requested operation is not permitted.

HTTP Status Code: 400

**ResourceNotFoundException**

This exception is thrown when the specified resource is not found.

HTTP Status Code: 400

**ResourceTypeNotSupportedException**

This exception is thrown when the specified resource type is not supported by CloudTrail.

HTTP Status Code: 400

**TagsLimitExceededException**

The number of tags per trail, event data store, dashboard, or channel has exceeded the permitted amount. Currently, the limit is
50.

HTTP Status Code: 400

**UnsupportedOperationException**

This exception is thrown when the requested operation is not supported.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudtrail-2013-11-01/addtags.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudtrail-2013-11-01/addtags.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

CancelQuery

All content copied from https://docs.aws.amazon.com/.
