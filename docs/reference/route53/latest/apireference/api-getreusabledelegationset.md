# GetReusableDelegationSet

Retrieves information about a specified reusable delegation set, including the four
name servers that are assigned to the delegation set.

## Request Syntax

```nohighlight

GET /2013-04-01/delegationset/Id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_GetReusableDelegationSet_RequestSyntax)**

The ID of the reusable delegation set that you want to get a list of name servers
for.

Length Constraints: Maximum length of 32.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetReusableDelegationSetResponse>
   <DelegationSet>
      <CallerReference>string</CallerReference>
      <Id>string</Id>
      <NameServers>
         <NameServer>string</NameServer>
      </NameServers>
   </DelegationSet>
</GetReusableDelegationSetResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetReusableDelegationSetResponse](#API_GetReusableDelegationSet_ResponseSyntax)**

Root level tag for the GetReusableDelegationSetResponse parameters.

Required: Yes

**[DelegationSet](#API_GetReusableDelegationSet_ResponseSyntax)**

A complex type that contains information about the reusable delegation set.

Type: [DelegationSet](api-delegationset.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DelegationSetNotReusable**

A reusable delegation set with the specified ID does not exist.

**message**

HTTP Status Code: 400

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchDelegationSet**

A reusable delegation set with the specified ID does not exist.

**message**

HTTP Status Code: 400

## Examples

### Example Request

This example illustrates one usage of GetReusableDelegationSet.

```

GET /2013-04-01/delegationset/N1PA6795SAMPLE
```

### Example Response

This example illustrates one usage of GetReusableDelegationSet.

```

<?xml version="1.0" encoding="UTF-8"?>
<GetReusableDelegationSetResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <DelegationSet>
      <Id>/delegationset/N1PA6795SAMPLE</Id>
      <CallerReference>2014-10-13T16:30:01Z</CallerReference>
      <NameServers>
         <NameServer>ns-2048.awsdns-64.com</NameServer>
         <NameServer>ns-2049.awsdns-65.net</NameServer>
         <NameServer>ns-2050.awsdns-66.org</NameServer>
         <NameServer>ns-2051.awsdns-67.co.uk</NameServer>
      </NameServers>
   </DelegationSet>
</GetReusableDelegationSetResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getreusabledelegationset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getreusabledelegationset.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetQueryLoggingConfig

GetReusableDelegationSetLimit
