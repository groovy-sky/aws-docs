# Understand error codes in the SharePoint (Online) connector

The following table provides information about error codes you may see for the
Microsoft SharePoint connector and suggested resolutions. If you used the Amazon Q section
of the AWS [console](https://console.aws.amazon.com/console/home) to
configure your connectors, be sure to make the changes associated with an error through
the console as well. You can refer to the Microsoft SharePoint [documentation](https://learn.microsoft.com/en-us/sharepoint) for more
information regarding SharePoint settings and details.

**Authentication issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5001Authentication failed. Configuration might contain wrong
credentials.Provide valid credentials like username, password or client Id,
client secret and tenant Id.Occurs during initial connector setup or when authentication
credentials expireConsole: During configuration or Sync run historySPE-5124There was a problem while retrieving authType. Auth-Type might be
empty or null.Ensure AUTH Type in configuration must be not null.Authentication type configuration errorConsole: Configuration pageSPE-5129There was a problem while retrieving password. Password might be
empty or null.Provide password.Basic authentication configuration errorConsole: During initial setupSPE-5130There was a problem while retrieving username.Username might be empty
or null.Provide username.Basic authentication configuration errorConsole: During initial setupSPE-5136The provided authType was not a valid Sharepoint Connector
authentication method.Provide valid authType. The value of authType should be one of
\[Basic, OAuth2Certificate, OAuth2\].Authentication method configuration errorConsole: Configuration pageSPE-5125There was a problem while retrieving clientId. Client ID might be
empty or null.Provide Client Id.OAuth configuration errorConsole: Authentication ConfigurationSPE-5126There was a problem while retrieving clientSecret. Client Secret
might be empty or null.Provide Client Secret.OAuth configuration errorConsole: Authentication ConfigurationSPE-5127There was a problem while retrieving tenantId. Tenant ID might be
empty or null.Provide Tenant Id.SharePoint tenant configuration errorConsole: Authentication Configuration

**SharePoint Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5135The provided version was not a valid Sharepoint Connector version.
Version should be one of \[Online, Server\].Version should be one of \[Online, Server\].SharePoint version selection errorConsole: During initial setupSPE-5138There was a problem while retrieving onPremVersion. On prem Version
might be empty or nullEnsure onPremVersion is not be null or non-empty.SharePoint on-premises version errorConsole: Version ConfigurationSPE-5139The provided onPremVersion was not valid Sharepoint on-prem version.
On prem version should be one of \[2013, 2016, 2019,
SubscriptionEdition\].Provide a valid onPremVersion. On prem version should be one of
\[2013, 2016, 2019, SubscriptionEdition\].SharePoint on-premises version selection errorConsole: Version ConfigurationSPE-5121There was a problem while retrieving values for crawl entities.
Values might be empty or incorrect. It should be either true or
false.In the connector's settings, ensure all crawl options are set to
either 'true' or 'false'. These settings determine what content types
are indexed.Crawler configuration errorConsole: Crawl ConfigurationSPE-5122There was a problem while retrieving domain. Domain might be empty or
null.Provide a valid SharePoint domain name in the connector
configuration.Domain configuration error during SharePoint connection
setup.Console: Domain ConfigurationSPE-5123There was a problem while retrieving version. Version might be empty
or null.Provide valid version and it should not be null.SharePoint version configuration errorConsole: Version Configuration

**Network and Connectivity issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5004Inet Address validation Failed.Check your network configuration and ensure the SharePoint server
address is accessible from your network. Verify DNS resolution and
network routing are properly configured.Network address validation errorCloudWatch Logs during syncSPE-5005Failed : HTTP protocol violation has occurred.Try running the connector again.Occurs when the connector is configured to use a proxy but cannot
establish connection.CloudWatch Logs during syncSPE-5200There was a problem while connecting to the URL.Verify that the SharePoint site URL is accessible and that you have
proper network connectivity. Check SharePoint site status and your
network configuration.Site connectivity errorConsole: Sync run historySPE-5002SPE-5002 Connection failed due to wrong credentials or invalid sites. Update your configuration and try again.Provide valid Host URL or Domain.Happens during connector configuration when trying to establish
initial connectionConsole: Basic ConfigurationSPE-5003Provided URL is incorrectProvide correct URL.Generic URL validation errorConsole: Site Configuration

**LDAP Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5009There was a problem while connecting to LDAP. Check LDAP
configuration.Provide valid LDAP configuration details.LDAP connection failureConsole: LDAP ConfigurationSPE-5140There was a problem while retrieving ldapUrl. LDAP Url might be empty
or null.Ensure ldapUrl is not null or empty.LDAP configuration errorConsole: LDAP ConfigurationSPE-5141There was a problem while retrieving baseDn. Base DN might be empty
or null.Ensure baseDn is not be null or empty.LDAP configuration errorConsole: LDAP ConfigurationSPE-5146There was a problem while retrieving ldapUsername. LDAP Username
might be empty or null.Ensure ldapUser is not null or empty.LDAP authentication errorConsole: LDAP ConfigurationSPE-5147There was a problem while retrieving ldapPassword. LDAP Password
might be empty or null.Ensure ldapPassword is not null or empty.LDAP authentication errorConsole: LDAP Configuration

**Microsoft Entra ID (formerly Azure AD) Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5152There was a problem while retrieving AD Client ID. AD Client ID
should not be empty.Ensure AD Client Id must be non-empty.Microsoft Entra ID (formerly Azure AD) configuration errorConsole: Microsoft Entra ID (formerly Azure AD) ConfigurationSPE-5153Invalid AD Client Id pattern.Provide valid AD Client Id pattern.Microsoft Entra ID (formerly Azure AD) Client ID format errorConsole: Microsoft Entra ID (formerly Azure AD) ConfigurationSPE-5154There was a problem while retrieving AD Client Secret. AD Client
Secret should not be empty.Ensure AD Client Secret is non-empty.Microsoft Entra ID (formerly Azure AD) configuration errorConsole: Microsoft Entra ID (formerly Azure AD) Configuration

**Document Access and Permissions issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5144There was a problem while retrieving aclConfiguration. ACL
Configuration might be empty, null or invalidProvide valid aclConfiguration. aclConfiguration should be one of \[
ACLWithLDAPEmailFmt, ACLWithManualEmailFmt, ACLWithUsernameFmt
\].ACL configuration errorConsole: Access Control ConfigurationSPE-5145There was a problem while retrieving emailDomain. Email Domain might
be empty or null.Ensure emailDomain is not null or empty.Email domain configuration errorConsole: Domain ConfigurationSPE-5155There can't be more than one site for SharePoint on-prem app-only
authentication.Ensure that their must be only single site present for SharePoint
on-prem app-only authentication.SharePoint on-premises app authentication configuration errorConsole: Site Configuration

**Security Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5008Valid SSL Certificate could not be found for connector.Upload a valid SSL certificate. For SharePoint on-premises
installations, ensure you have exported your SharePoint SSL certificate
and uploaded it to the connector.SSL certificate validation failureConsole: Security Configuration

**IAM Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5101There was a problem while retrieving dataSourceIamRoleArn. Data
Source IAM Role ARN might be empty or null.Verify that the IAM role exists and has the correct permissions.
Check the IAM console to ensure the role is properly configured.This error occurs when the connector cannot access the required IAM
role during synchronization.Console: IAM Configuration

**Site Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5128There was a problem while retrieving siteUrls. Site URLs might be
empty or null.Provide at least one Site Url.Site configuration errorConsole: Site ConfigurationSPE-5131There was a problem while retrieving username. Email was
invalid.Provide valid email address.User email validation errorConsole: User ConfigurationSPE-5132There was a problem while retrieving url. This URL was
invalid.Provide a valid URL.Provide a valid SharePoint site URL in the format
https://your-sharepoint-site.comConsole: Site ConfigurationSPE-5149The provided siteUrls contain duplicate sites. Remove
duplicates.Ensure SiteUrls must not be the same.Duplicate site URL configuration errorConsole: Site Configuration

**Other Configuration issues**

Error codeError messageSuggested resolutionContextWhere to find errorSPE-5133There was a problem while retrieving s3CertificateName. S3
Certificate Name might be empty or null.When using certificate-based authentication, upload your
authentication certificate to an S3 bucket and provide the certificate
name and bucket details in the connector configuration.S3 certificate configuration errorConsole: Security ConfigurationSPE-5134There was a problem while retrieving s3BucketName. S3 Bucket Name
might be empty or nullWhen using certificate-based authentication, upload your
authentication certificate to an S3 bucket and provide the certificate
name and bucket details in the connector configuration.S3 bucket configuration errorConsole: S3 ConfigurationSPE-5151Error parsing the field value. Size is over maximum allowed
limit.Reduce the field value size to within the maximum allowed
limit.Field size limit exceededConsole: Field Configuration

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM role

Microsoft SharePoint Server 2016
