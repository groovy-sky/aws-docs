# Creating an Amazon Q Business application using Identity Federation through IAM

This section walks you through creating and configuring an Amazon Q Business
application using IAM federation to manage end user access.

Amazon Q Business supports identity federation through AWS Identity and Access Management. When you use
identity federation, you can manage users with your enterprise identity provider (IdP) and
use AWS Identity and Access Management to authenticate users when they sign in to Amazon Q Business.

You can use any third-party identity provider that supports Security Assertion Markup
Language 2.0 (SAML 2.0) or OpenID Connect (OIDC) to provide an onboarding flow for your
Amazon Q Business users. Such identity providers include, but aren't limited to
Okta and Ping Identity.

###### Important

Amazon Q Business supports Microsoft Entra ID with SAML 2.0, but
doesn't support OIDC for Google or Microsoft Entra
ID.

With identity federation, your users get one-click access to their Amazon Q Business
applications using their existing identity credentials. You also have the security benefit
of identity authentication by your identity provider. You can control which users have
access to Amazon Q Business using your existing identity provider.

###### Note

Federated groups aren't supported through IAM Federation. If you want to ingest
federated groups, use the [PutGroup](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_PutGroup.html) API.

###### Topics

- [Creating an Amazon Q Business application using IAM Federation through Okta](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-application-iam-okta.html)

- [Creating an Amazon Q Business application using IAM federation through Microsoft Entra ID](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/create-application-iam-entraid.html)

- [Connecting multiple Amazon Q Business applications to an Identity Provider](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/multiple-qbusiness-apps-idp.html)

- [Making authenticated Amazon Q Business API calls using IAM federation](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/making-sigv4-authenticated-api-calls-iam.html)

- [Managing Amazon Q Business resources](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/managing-resources-iam.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tagging resources

Using Okta for IAM federation
