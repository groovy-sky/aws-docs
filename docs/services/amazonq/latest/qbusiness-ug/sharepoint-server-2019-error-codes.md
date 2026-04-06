# Understand error codes in the SharePoint Server 2019 connector

The following table provides information about error codes you may see for the
Microsoft SharePoint connector and suggested resolutions.

Error codeError messageSuggested resolutionSPE-5001Authentication failed. Configuration might contain wrong
credentials.Provide valid credentials like username, password or client Id,
client secret and tenant Id. SPE-5002There was a problem while connecting to Host Url and/or Domain.
hostUrl and/or domain values might be incorrect.Provide valid Host URL or Domain.SPE-5003Provided URL is incorrectProvide correct URL.SPE-5004Inet Address validation Failed.Provide valid Inet AddressSPE-5005Failed : HTTP protocol violation has occurred.Try running the connector again.SPE-5006Cannot connect to proxy. Check the proxy configurations. Provide valid Proxy configuration details.SPE-5007Proxy port is invalid. Check the proxy port.Provide valid Proxy port.SPE-5008 Valid SSL Certificate could not be found for connector. Provide valid SSL certificate.SPE-5009 There was a problem while connecting to LDAP. Check LDAP
configuration. Provide valid LDAP configuration details.SPE-5100There was a problem while retrieving repositoryId. Repository ID
might be empty or null.Ensure that repository Id must not be null or empty.SPE-5101There was a problem while retrieving dataSourceIamRoleArn. Data
Source IAM Role ARN might be empty or null.Ensure that dataSourceIamRoleArn must not be null or empty.SPE-5102There was a problem while retrieving repository configurations.
Repository configurations might be empty or incorrect.Provide valid repository configurations.SPE-5115There was a problem while retrieving field mapping values for event
entity. Field mapping values might be empty or incorrect.Field mapping values for event entity should be correct or
non-empty.SPE-5116There was a problem while retrieving field mapping values for file
entity. Field mapping values might be empty or incorrect.Field mapping values for file entity should be correct or
non-empty.SPE-5117There was a problem while retrieving field mapping values for page
entity. Field mapping values might be empty or incorrect.Field mapping values for page entity should be correct or
non-empty.SPE-5118There was a problem while retrieving field mapping values for link
entity. Field mapping values might be empty or incorrect.Field mapping values for link entity should be correct or
non-empty.SPE-5119There was a problem while retrieving field mapping values for comment
entity. Field mapping values might be empty or incorrect.Field mapping values for comment entity should be correct or
non-empty.SPE-5120There was a problem while retrieving field mapping values for
attachment entity. Field mapping values might be empty or
incorrect.Field mapping values for attachment entity should be correct or
non-empty.SPE-5121There was a problem while retrieving values for crawl entities.
Values might be empty or incorrect. It should be either true or
false.There might be some incorrect value given in any one of the crawling
entities like – null, TRUE or any dummy string. Ensure the value must be
non-empty and either true or false.SPE-5122There was a problem while retrieving domain. Domain might be empty or
null.Provide Client Id.SPE-5123There was a problem while retrieving version. Version might be empty
or null.Provide valid version and it should not be null.SPE-5124There was a problem while retrieving authType. Auth-Type might be
empty or null.Ensure AUTH Type in configuration must be not null.SPE-5125There was a problem while retrieving clientId. Client ID might be
empty or null.Provide Client Id.SPE-5126There was a problem while retrieving clientSecret. Client Secret
might be empty or null.Provide Client Secret.SPE-5127There was a problem while retrieving tenantId. Tenant ID might be
empty or null.Provide Tenant Id.SPE-5128There was a problem while retrieving siteUrls. Site URLs might be
empty or null.Provide at least one Site Url.SPE-5129There was a problem while retrieving password. Password might be
empty or null.Provide password.SPE-5130There was a problem while retrieving username.Username might be empty
or null.Provide username.SPE-5131There was a problem while retrieving username. Email was
invalid.Provide valid email address.SPE-5132There was a problem while retrieving url. This URL was
invalid.Provide a valid URL.SPE-5133There was a problem while retrieving s3CertificateName. S3
Certificate Name might be empty or null.Ensure s3CertificateName is not null or non-empty.SPE-5134There was a problem while retrieving s3BucketName. S3 Bucket Name
might be empty or nullEnsure s3BucketName is not null or non-empty.SPE-5135The provided version was not a valid Sharepoint Connector version.
Version should be one of \[Online, Server\].Version should be one of \[Online, Server\].SPE-5136The provided authType was not a valid Sharepoint Connector
authentication method.Provide valid authType. The value of authType should be one of
\[Basic, OAuth2Certificate, OAuth2\].SPE-5138There was a problem while retrieving onPremVersion. On prem Version
might be empty or nullEnsure onPremVersion is not be null or non-empty.SPE-5139The provided onPremVersion was not valid Sharepoint on-prem version.
On prem version should be one of \[2013, 2016, 2019,
SubscriptionEdition\].Provide a valid onPremVersion. On prem version should be one of
\[2013, 2016, 2019, SubscriptionEdition\].SPE-5140There was a problem while retrieving ldapUrl. LDAP Url might be empty
or null.Ensure ldapUrl is not null or empty.SPE-5141There was a problem while retrieving baseDn. Base DN might be empty
or null.Ensure baseDn is not be null or empty.SPE-5142There was a problem while retrieving privateKey. Private Key might be
empty or null.Please ensure privateKey is not be null or empty.SPE-5144There was a problem while retrieving aclConfiguration. ACL
Configuration might be empty, null or invalidProvide valid aclConfiguration. aclConfiguration should be one of \[
ACLWithLDAPEmailFmt, ACLWithManualEmailFmt, ACLWithUsernameFmt \].
SPE-5145There was a problem while retrieving emailDomain. Email Domain might
be empty or null.Ensure emailDomain is not null or empty.SPE-5146There was a problem while retrieving ldapUsername. LDAP Username
might be empty or null.Ensure ldapUser is not null or empty.SPE-5147There was a problem while retrieving ldapPassword. LDAP Password
might be empty or null.Ensure ldapPassword is not null or empty.SPE-5149The provided siteUrls contain duplicate sites. Remove
duplicates.Ensure SiteUrls must not be the same.SPE-5150Invalid Client Id pattern.Provide the correct client ID.SPE-5151Error parsing the field value. Size is over maximum allowed
limit.Ensure the size limit.SPE-5152There was a problem while retrieving AD Client ID. AD Client ID
should not be empty.Ensure AD Client Id must be non-empty.SPE-5153Invalid AD Client Id pattern.Provide valid AD Client Id pattern.SPE-5154 There was a problem while retrieving AD Client Secret. AD Client
Secret should not be empty.Ensure AD Client Secret is non-empty.SPE-5155There can't be more than one site for SharePoint on-prem app-only
authentication.Ensure that their must be only single site present for SharePoint
on-prem app-only authentication.SPE-5200There was a problem while connecting to the URL.Ensure the siteUrl exists.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM role

Microsoft SharePoint Server (Subscription Edition)
