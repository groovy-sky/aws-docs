# ListPrices

Lists the following prices for either all the TLDs supported by Route 53, or
the specified TLD:

- Registration

- Transfer

- Owner change

- Domain renewal

- Domain restoration

## Request Syntax

```nohighlight

{
   "Marker": "string",
   "MaxItems": number,
   "Tld": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Marker](#API_domains_ListPrices_RequestSyntax)**

For an initial request for a list of prices, omit this element. If the number of
prices that are not yet complete is greater than the value that you specified for
`MaxItems`, you can use `Marker` to return additional prices.
Get the value of `NextPageMarker` from the previous response, and submit
another request that includes the value of `NextPageMarker` in the
`Marker` element.

Used only for all TLDs. If you specify a TLD, don't specify a
`Marker`.

Type: String

Length Constraints: Maximum length of 4096.

Required: No

**[MaxItems](#API_domains_ListPrices_RequestSyntax)**

Number of `Prices` to be returned.

Used only for all TLDs. If you specify a TLD, don't specify a
`MaxItems`.

Type: Integer

Valid Range: Maximum value of 1000.

Required: No

**[Tld](#API_domains_ListPrices_RequestSyntax)**

The TLD for which you want to receive the pricing information. For example.
`.net`.

If a `Tld` value is not provided, a list of prices for all TLDs supported
by Route 53 is returned.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 255.

Required: No

## Response Syntax

```nohighlight

{
   "NextPageMarker": "string",
   "Prices": [
      {
         "ChangeOwnershipPrice": {
            "Currency": "string",
            "Price": number
         },
         "Name": "string",
         "RegistrationPrice": {
            "Currency": "string",
            "Price": number
         },
         "RenewalPrice": {
            "Currency": "string",
            "Price": number
         },
         "RestorationPrice": {
            "Currency": "string",
            "Price": number
         },
         "TransferPrice": {
            "Currency": "string",
            "Price": number
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextPageMarker](#API_domains_ListPrices_ResponseSyntax)**

If there are more prices than you specified for `MaxItems` in the request,
submit another request and include the value of `NextPageMarker` in the value
of `Marker`.

Used only for all TLDs. If you specify a TLD, don't specify a
`NextPageMarker`.

Type: String

Length Constraints: Maximum length of 4096.

**[Prices](#API_domains_ListPrices_ResponseSyntax)**

A complex type that includes all the pricing information. If you specify a TLD, this
array contains only the pricing for that TLD.

Type: Array of [DomainPrice](api-domains-domainprice.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The requested item is not acceptable. For example, for APIs that accept a domain name,
the request might specify a domain name that doesn't belong to the account that
submitted the request. For `AcceptDomainTransferFromAnotherAwsAccount`, the
password might be invalid.

**message**

The requested item is not acceptable. For example, for an OperationId it might refer
to the ID of an operation that is already completed. For a domain name, it might not be
a valid domain name or belong to the requester account.

HTTP Status Code: 400

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/listprices.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/listprices.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/listprices.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/listprices.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/listprices.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/listprices.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/listprices.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/listprices.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/listprices.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/listprices.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListOperations

ListTagsForDomain
