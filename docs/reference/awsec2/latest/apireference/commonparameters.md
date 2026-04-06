# Common query parameters

Most Amazon EC2 API actions support the parameters described in the following tables. The common
parameters vary depending on whether you're using Signature Version 2 or Signature Version 4 to
sign your requests. For more information, see [Signing AWS API requests](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html) in the _IAM User Guide_.

###### Contents

- [Parameters for Signature Version 4](#common-parameters-sigv4)

- [Parameters for Signature Version 2](#common-parameters-sigv2)

## Parameters for Signature Version 4

NameDescriptionRequired

`Action`

The action to perform.

Example: `RunInstances`

Yes

`Version`

The API version to use.

Yes

`X-Amz-Algorithm`

The hash algorithm you use to create the request signature.

Example: `AWS4-HMAC-SHA256`

Yes

`X-Amz-Credential`

The credential scope for the request, in the format
`access-key-ID`/ `YYYYMMDD`/ `region`/ `service`/ `aws4_request`

Example: `AKIDEXAMPLE/20140707/us-east-1/ec2/aws4_request`

Yes

`X-Amz-Date`

The date and time at which the request is signed, in the format YYYYMMDDThhmmssZ. The
date must match the date that's included in the credential scope for the
`X-Amz-Credential` parameter, or the date used in an
`Authorization` header (see the note below the table).

Example: `20140707T150456Z`

Yes`X-Amz-SignedHeaders`

The headers you are including as part of the request. At a minimum, you must include the
`host` header. If you include an `x-amz-date` header in your
request, you must include it in the list of signed headers.

Example: `content-type;host;user-agent`

Yes

`X-Amz-Signature`

A signature derived from your secret access key.

Example: `ced6826de92d2bdeed8f846f0bf508e8559example`

Yes

`X-Amz-Security-Token`

The temporary security token obtained through a call to AWS Security Token Service.

Example: `AQoEXAMPLEH4aoAH0gNCAPyJxz4BlCFFxWNE1OPTgk5TthT+FvwqnKwRcOIfrRh3c/L`

No

`DryRun`

Checks whether you have the required permissions for the action, without actually making the request.
If you have the required permissions, the request returns `DryRunOperation`; otherwise, it
returns `UnauthorizedOperation`.

No

The `X-Amz-Algorithm`, `X-Amz-Credential`,
`X-Amz-SignedHeaders`, and `X-Amz-Signature` parameters can either
be specified as separate parameters in the query string, or their values can be included in
a single `Authorization` header. For more information, see [Signing AWS API requests](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html) in the _IAM User Guide_.

## Parameters for Signature Version 2

NameDescriptionRequired

`Action`

The action to perform.

Example: `RunInstances`

Yes

`Version`

The API version to use.

Yes

`AWSAccessKeyId`

The access key ID for the request sender. This identifies the account which will be
charged for usage of the service. The account that's associated with the access key
ID must be signed up for Amazon EC2, or the request isn't accepted.

Example: `AKIAIOSFODNN7EXAMPLE`

Yes

`Expires`

The date and time at which the signature included in the request expires, in the
format YYYY-MM-DDThh:mm:ssZ. For more information, see [ISO 8601](http://www.w3.org/TR/NOTE-datetime).

Example: `2006-07-07T15:04:56Z`

Conditional. Requests must include either `Timestamp` or
`Expires`, but cannot contain both.

`Timestamp`

The date and time at which the request is signed,
in the format YYYY-MM-DDThh:mm:ssZ. For more information, see
[ISO 8601](http://www.w3.org/TR/NOTE-datetime).

Example: `2006-07-07T15:04:56Z`

Conditional. Requests must include either `Timestamp` or
`Expires`, but cannot contain both.

`Signature`

The request signature.

Example: `Qnpl4Qk/7tINHzfXCiT7VEXAMPLE`

Yes

`SignatureMethod`

The hash algorithm you use to create the request signature. Valid values:
`HmacSHA256` \| `HmacSHA1`.

Example: `HmacSHA256`

Yes

`SignatureVersion`

The signature version you use to sign the request. Set this value to `2`.

Example: `2`

Yes

`DryRun`

Checks whether you have the required permissions for the action, without
actually making the request. If you have the required permissions, the request
returns `DryRunOperation`; otherwise, it returns
`UnauthorizedOperation`.

No

`SecurityToken`

The temporary security token obtained through a call to AWS Security Token Service.

Example: `AQoEXAMPLEH4aoAH0gNCAPyJxz4BlCFFxWNE1OPTgk5TthT+FvwqnKwRcOIfrRh3c/L`

No

Parameter values must be URL-encoded. This is true for any Query parameter passed to Amazon EC2
and is typically necessary in the `Signature` parameter. Some clients do
this automatically, but this is not the norm.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VM Import Manifest

Permissions
