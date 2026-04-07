# AWS condition keys that can be used in API Gateway resource policies

The following table contains AWS condition keys that can be used
in resource policies for APIs in API Gateway for each authorization type.

For more information about AWS condition keys, see [AWS Global Condition\
Context Keys](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html).

**Condition keys**

**Criteria**

**Needs `AuthN`?**

**Authorization type**

`aws:CurrentTime`NoneNoAll`aws:EpochTime`NoneNoAll`aws:TokenIssueTime`Key is present only in requests that are signed using temporary
security credentials.YesIAM`aws:MultiFactorAuthPresent`Key is present only in requests that are signed using temporary
security credentials.YesIAM`aws:MultiFactorAuthAge`Key is present only if MFA is present in the requests.YesIAM`aws:PrincipalAccount`NoneYesIAM`aws:PrincipalArn`NoneYesIAM`aws:PrincipalOrgID`This key is included in the request context only if the principal is a member of an organization.YesIAM`aws:PrincipalOrgPaths`This key is included in the request context only if the principal is a member of an organization.YesIAM`aws:PrincipalTag`This key is included in the request context if the principal is using an IAM user with attached tags. It is included for a
principal using an IAM role with attached tags or session tags.YesIAM`aws:PrincipalType`NoneYesIAM`aws:Referer`Key is present only if the value is provided by the caller in the
HTTP header.NoAll`aws:SecureTransport`NoneNoAll`aws:SourceArn`NoneNoAll`aws:SourceIp`NoneNoAll`aws:SourceVpc`This key can be used only for private APIs.NoAll`aws:SourceVpce`This key can be used only for private APIs.NoAll`aws:VpcSourceIp`This key can be used only for private APIs.NoAll`aws:UserAgent`Key is present only if the value is provided by the caller in the
HTTP header.NoAll`aws:userid`NoneYesIAM`aws:username`NoneYesIAM

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create and attach an API Gateway resource policy to an API

Use IAM permissions
