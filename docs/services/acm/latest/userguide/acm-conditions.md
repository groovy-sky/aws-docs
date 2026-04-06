# Use condition keys with ACM

AWS Certificate Manager uses AWS Identity and Access Management (IAM) [condition keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html) to limit access to certificate requests. With condition keys from
IAM policies or Service Control Policies (SCP) you can create certificate requests that conform to your organization's
guidelines.

###### Note

Combine ACM condition keys with AWS [global condition keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html) such as `aws:PrincipalArn` to further restrict actions to specific users or roles.

## Supported conditions for ACM

Use the scroll bars to see the rest of the table.

ACM API operations and supported conditionsCondition KeySupported ACM API OperationsTypeDescription

`acm:ValidationMethod`

[RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)

String ( `DNS`, `EMAIL`, `HTTP`)

Filter requests based on ACM [validation method](domain-ownership-validation.md)

`acm:DomainNames`

[RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)

ArrayOfString

Filter based on [domain names](acm-concepts.md#concept-dn) in the ACM request

`acm:KeyAlgorithm`

[RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)

String

Filter requests based on ACM [key algorithm and size](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms)

`acm:CertificateTransparencyLogging`

[RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)

String ( `ENABLED`, `DISABLED`)

Filter requests based on ACM [certificate transparency logging preference](acm-concepts.md#concept-transparency)

`acm:CertificateAuthority`

[RequestCertificate](../../../../reference/acm/latest/apireference/api-requestcertificate.md)

ARN

Filter requests based on [certificate authorities](acm-concepts.md#concept-ca) in the ACM request

## Example 1: Restricting validation method

The following policy denies new certificate requests using the [Email Validation](domain-ownership-validation.md)
method except for a request made using the `arn:aws:iam::123456789012:role/AllowedEmailValidation` role.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition":{
            "StringLike" : {
                "acm:ValidationMethod":"EMAIL"
            },
            "ArnNotLike": {
                "aws:PrincipalArn": [ "arn:aws:iam::123456789012:role/AllowedEmailValidation"]
            }
        }
    }
}

```

## Example 2: Preventing wildcard domains

The following policy denies any new ACM certificate request that uses wildcard domains.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition": {
            "ForAnyValue:StringLike": {
                "acm:DomainNames": [
                    "${*}.*"
                ]
            }
        }
    }
}

```

## Example 3: Restricting certificate domains

The following policy denies any new ACM certificate request for domains that don't end with `*.amazonaws.com`

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition": {
            "ForAnyValue:StringNotLike": {
                "acm:DomainNames": ["*.amazonaws.com"]
            }
        }
    }
}

```

The policy could be further restricted to specific subdomains. This policy
would only allow requests where every domain matches at least one of the conditional domain names.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition": {
            "ForAllValues:StringNotLike": {
                "acm:DomainNames": ["support.amazonaws.com", "developer.amazonaws.com"]
            }
        }
    }
}

```

## Example 4: Restricting key algorithm

The following policy uses the condition key `StringNotLike` to allow only certificates requested with the ECDSA 384 bit ( `EC_secp384r1`) key algorithm.

JSON

```json

{
    "Version":"2012-10-17",
        "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition":{
            "StringNotLike" : {
                "acm:KeyAlgorithm":"EC_secp384r1"
            }
        }
    }
}

```

The following policy uses the condition key `StringLike` and wildcard `*` matching to prevent requests for new certificates in ACM with any `RSA` key algorithm.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition":{
            "StringLike" : {
                "acm:KeyAlgorithm":"RSA*"
            }
        }
    }
}

```

## Example 5: Restricting certificate authority

The following policy would only allow requests for private certificates using the provided Private Certificate Authority (PCA) ARN.

JSON

```json

{
    "Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition":{
            "StringNotLike": {
                "acm:CertificateAuthority":" arn:aws:acm-pca:region:account:certificate-authority/CA_ID"
            }
        }
    }
}

```

This policy uses the `acm:CertificateAuthority` condition to allow only requests for publicly trusted certificates issued by
Amazon Trust Services. Setting the Certificate Authority ARN to `false` prevents requests for private certificates from PCA.

JSON

```json

{
"Version":"2012-10-17",
    "Statement":{
        "Effect":"Deny",
        "Action":"acm:RequestCertificate",
        "Resource":"*",
        "Condition":{
            "Null" : {
                "acm:CertificateAuthority":"false"
            }
        }
    }
}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS managed policies

Use service-linked roles
