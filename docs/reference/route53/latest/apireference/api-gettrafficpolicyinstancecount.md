# GetTrafficPolicyInstanceCount

Gets the number of traffic policy instances that are associated with the current
AWS account.

## Request Syntax

```

GET /2013-04-01/trafficpolicyinstancecount HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyInstanceCountResponse>
   <TrafficPolicyInstanceCount>integer</TrafficPolicyInstanceCount>
</GetTrafficPolicyInstanceCountResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetTrafficPolicyInstanceCountResponse](#API_GetTrafficPolicyInstanceCount_ResponseSyntax)**

Root level tag for the GetTrafficPolicyInstanceCountResponse parameters.

Required: Yes

**[TrafficPolicyInstanceCount](#API_GetTrafficPolicyInstanceCount_ResponseSyntax)**

The number of traffic policy instances that are associated with the current AWS account.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### Example Request

This example illustrates one usage of GetTrafficPolicyInstanceCount.

```

GET /2013-04-01/trafficpolicyinstancecount/
```

### Example Response

This example illustrates one usage of GetTrafficPolicyInstanceCount.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetTrafficPolicyInstanceCountResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <TrafficPolicyInstanceCount>42</TrafficPolicyInstanceCount>
</GetTrafficPolicyInstanceCountResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/gettrafficpolicyinstancecount.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/gettrafficpolicyinstancecount.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetTrafficPolicyInstance

ListCidrBlocks
