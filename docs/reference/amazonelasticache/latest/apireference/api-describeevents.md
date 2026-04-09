# DescribeEvents

Returns events related to clusters, cache security groups, and cache parameter groups.
You can obtain events specific to a particular cluster, cache security group, or cache
parameter group by providing the name as a parameter.

By default, only the events occurring within the last hour are returned; however, you
can retrieve up to 14 days' worth of events if necessary.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Duration**

The number of minutes worth of events to retrieve.

Type: Integer

Required: No

**EndTime**

The end of the time interval for which to retrieve events, specified in ISO 8601
format.

**Example:** 2017-03-30T07:03:49.555Z

Type: Timestamp

Required: No

**Marker**

An optional marker returned from a prior request. Use this marker for pagination of
results from this operation. If this parameter is specified, the response includes only
records beyond the marker, up to the value specified by `MaxRecords`.

Type: String

Required: No

**MaxRecords**

The maximum number of records to include in the response. If more records exist than
the specified `MaxRecords` value, a marker is included in the response so
that the remaining results can be retrieved.

Default: 100

Constraints: minimum 20; maximum 100.

Type: Integer

Required: No

**SourceIdentifier**

The identifier of the event source for which events are returned. If not specified,
all sources are included in the response.

Type: String

Required: No

**SourceType**

The event source to retrieve events for. If no value is specified, all events are
returned.

Type: String

Valid Values: `cache-cluster | cache-parameter-group | cache-security-group | cache-subnet-group | replication-group | serverless-cache | serverless-cache-snapshot | user | user-group`

Required: No

**StartTime**

The beginning of the time interval to retrieve events for, specified in ISO 8601
format.

**Example:** 2017-03-30T07:03:49.555Z

Type: Timestamp

Required: No

## Response Elements

The following elements are returned by the service.

**Events.Event.N**

A list of events. Each element in the list contains detailed information about one
event.

Type: Array of [Event](api-event.md) objects

**Marker**

Provides an identifier to allow retrieval of paginated results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

## Examples

### DescribeEvents

Some of the output has been omitted for brevity.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
   ?Action=DescribeEvents
   &MaxRecords=100
   &Version=2015-02-02
   &SignatureVersion=4
   &SignatureMethod=HmacSHA256
   &Timestamp=20150202T192317Z
   &X-Amz-Credential=<credential>

```

#### Sample Response

```

<DescribeEventsResponse xmlns="http://elasticache.amazonaws.com/doc/2015-02-02/">
    <DescribeEventsResult>
        <Events>
            <Event>
                <Message>Cache cluster created</Message>
                <SourceType>cache-cluster</SourceType>
                <Date>2015-02-02T18:22:18.202Z</Date>
                <SourceIdentifier>my-redis-primary</SourceIdentifier>
            </Event>

 (...output omitted...)

        </Events>
    </DescribeEventsResult>
    <ResponseMetadata>
        <RequestId>e21c81b4-b9cd-11e3-8a16-7978bb24ffdf</RequestId>
    </ResponseMetadata>
</DescribeEventsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/describeevents.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/describeevents.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeEngineDefaultParameters

DescribeGlobalReplicationGroups

All content copied from https://docs.aws.amazon.com/.
