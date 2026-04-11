# x-amazon-apigateway-security-policy

Specifies a security policy for a REST API. If you create a security policy that starts with
`"SecurityPolicy_"`, you must also set the
[endpoint access mode](openapi-extensions-endpoint-access-mode.md). To learn more about security
policies, see [Security policies for REST APIs in API Gateway](apigateway-security-policies.md).

## `x-amazon-apigateway-security-policy` example

The following example specifies
`SecurityPolicy_TLS13_1_3_2025_0` for a REST API.

```nohighlight

"x-amazon-apigateway-security-policy": "SecurityPolicy_TLS13_1_3_2025_09"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-request-validators.requestValidator

x-amazon-apigateway-tag-value

All content copied from https://docs.aws.amazon.com/.
