# DeleteDomain

## Description

The `DeleteDomain` operation deletes a domain. Any items (and their
attributes) in the domain are deleted as well. The `DeleteDomain` operation
might take 10 or more seconds to complete.

###### Note

Running `DeleteDomain` on a domain that does not exist or running the
function multiple times using the same domain name will _not_ result in an
error response.

###### Important

Domain deletion is blocked while any export for that domain is in PENDING or IN\_PROGRESS
status. Ensure all exports have completed (SUCCEEDED or FAILED status) before attempting to
delete a domain. You can use the `ListExports` operation to check for active exports.

## Request Parameters

Name  Description  Required`DomainName`

The name of the domain to delete.

Type: String

Yes

## Response Elements

See [Common Response Elements](sdb-api-commonresponseelements.md).

## Special Errors

Error  Description `MissingParameter` The request must contain the parameter
`DomainName`.

## Examples

### Sample Request

```

https://sdb.amazonaws.com/
?Action=DeleteDomain
&AWSAccessKeyId=[valid access key id]
&DomainName=MyDomain
&SignatureVersion=2
&SignatureMethod=HmacSHA256
&Timestamp=2010-01-25T15%3A02%3A20-07%3A00
&Version=2009-04-15
&Signature=[valid signature]

```

### Sample Response

```nohighlight

<DeleteDomainResponse>
  <ResponseMetadata>
    <RequestId>c522638b-31a2-4d69-b376-8c5428744704</RequestId>
    <BoxUsage>0.0000219907</BoxUsage>
  </ResponseMetadata>
</DeleteDomainResponse>

```

## Related Actions

- [CreateDomain](sdb-api-createdomain.md)

- [ListDomains](sdb-api-listdomains.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAttributes

DomainMetadata

All content copied from https://docs.aws.amazon.com/.
