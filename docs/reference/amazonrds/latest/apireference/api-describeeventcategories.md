# DescribeEventCategories

Displays a list of categories for all event source types, or, if specified, for a specified source type.
You can also see this list in the "Amazon RDS event categories and event messages" section of the [_Amazon RDS User Guide_](../../../../services/amazonrds/latest/userguide/user-events-messages.md) or the
[_Amazon Aurora User Guide_](../../../../services/amazonrds/latest/aurorauserguide/user-events-messages.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Filters.Filter.N**

This parameter isn't currently supported.

Type: Array of [Filter](api-filter.md) objects

Required: No

**SourceType**

The type of source that is generating the events. For RDS Proxy events, specify `db-proxy`.

Valid Values: `db-instance` \| `db-cluster` \| `db-parameter-group` \| `db-security-group` \| `db-snapshot` \| `db-cluster-snapshot` \| `db-proxy`

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**EventCategoriesMapList.EventCategoriesMap.N**

A list of `EventCategoriesMap` data types.

Type: Array of [EventCategoriesMap](api-eventcategoriesmap.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example

This example illustrates one usage of DescribeEventCategories.

#### Sample Request

```

https://rds.us-west-2.amazonaws.com/
   ?Action=DescribeEventCategories
   &SignatureMethod=HmacSHA256
   &SignatureVersion=4
   &Version=2014-10-31
   &X-Amz-Algorithm=AWS4-HMAC-SHA256
   &X-Amz-Credential=AWS_ACCESS_KEY_ID_REDACTED/20140421/us-west-2/rds/aws4_request
   &X-Amz-Date=20140421T194732Z
   &X-Amz-SignedHeaders=content-type;host;user-agent;x-amz-content-sha256;x-amz-date
   &X-Amz-Signature=6e25c542bf96fe24b28c12976ec92d2f856ab1d2a158e21c35441a736e4fde2b

```

#### Sample Response

```

<DescribeEventCategoriesResponse xmlns="http://rds.amazonaws.com/doc/2014-10-31/">
  <DescribeEventCategoriesResult>
    <EventCategoriesMapList>
      <EventCategoriesMap>
        <SourceType>db-instance</SourceType>
        <EventCategories>
          <EventCategory>backup</EventCategory>
          <EventCategory>recovery</EventCategory>
          <EventCategory>restoration</EventCategory>
          <EventCategory>failover</EventCategory>
          <EventCategory>low storage</EventCategory>
          <EventCategory>maintenance</EventCategory>
          <EventCategory>deletion</EventCategory>
          <EventCategory>availability</EventCategory>
          <EventCategory>configuration change</EventCategory>
          <EventCategory>notification</EventCategory>
          <EventCategory>failure</EventCategory>
          <EventCategory>creation</EventCategory>
        </EventCategories>
      </EventCategoriesMap>
      <EventCategoriesMap>
        <SourceType>db-security-group</SourceType>
        <EventCategories>
          <EventCategory>configuration change</EventCategory>
          <EventCategory>failure</EventCategory>
        </EventCategories>
      </EventCategoriesMap>
      <EventCategoriesMap>
        <SourceType>db-parameter-group</SourceType>
        <EventCategories>
          <EventCategory>configuration change</EventCategory>
        </EventCategories>
      </EventCategoriesMap>
      <EventCategoriesMap>
        <SourceType>db-snapshot</SourceType>
        <EventCategories>
          <EventCategory>deletion</EventCategory>
          <EventCategory>restoration</EventCategory>
          <EventCategory>notification</EventCategory>
          <EventCategory>failure</EventCategory>
          <EventCategory>creation</EventCategory>
        </EventCategories>
      </EventCategoriesMap>
    </EventCategoriesMap>
  </DescribeEventCategoriesResult>
  <ResponseMetadata>
    <RequestId>b79456f2-b98c-11d3-f272-7cd6cce12cc5</RequestId>
  </ResponseMetadata>
</DescribeEventCategoriesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/describeeventcategories.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/describeeventcategories.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeEngineDefaultParameters

DescribeEvents

All content copied from https://docs.aws.amazon.com/.
