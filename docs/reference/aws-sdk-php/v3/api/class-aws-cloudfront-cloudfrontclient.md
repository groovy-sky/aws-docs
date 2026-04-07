Menu

- [Aws](namespace-aws.md)
- [CloudFront](namespace-aws-cloudfront.md)

## CloudFrontClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon CloudFront** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2020-05-31 (latest)**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html)

  - [AssociateAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#associatealias)
  - [AssociateDistributionTenantWebACL](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#associatedistributiontenantwebacl)
  - [AssociateDistributionWebACL](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#associatedistributionwebacl)
  - [CopyDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#copydistribution)
  - [CreateAnycastIpList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createanycastiplist)
  - [CreateCachePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createcachepolicy)
  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createcloudfrontoriginaccessidentity)
  - [CreateConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createconnectionfunction)
  - [CreateConnectionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createconnectiongroup)
  - [CreateContinuousDeploymentPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createcontinuousdeploymentpolicy)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createdistribution)
  - [CreateDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createdistributiontenant)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createdistributionwithtags)
  - [CreateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createfieldlevelencryptionconfig)
  - [CreateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createfieldlevelencryptionprofile)
  - [CreateFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createfunction)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createinvalidation)
  - [CreateInvalidationForDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createinvalidationfordistributiontenant)
  - [CreateKeyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createkeygroup)
  - [CreateKeyValueStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createkeyvaluestore)
  - [CreateMonitoringSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createmonitoringsubscription)
  - [CreateOriginAccessControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createoriginaccesscontrol)
  - [CreateOriginRequestPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createoriginrequestpolicy)
  - [CreatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createpublickey)
  - [CreateRealtimeLogConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createrealtimelogconfig)
  - [CreateResponseHeadersPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createresponseheaderspolicy)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createstreamingdistributionwithtags)
  - [CreateTrustStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createtruststore)
  - [CreateVpcOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#createvpcorigin)
  - [DeleteAnycastIpList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteanycastiplist)
  - [DeleteCachePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletecachepolicy)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletecloudfrontoriginaccessidentity)
  - [DeleteConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteconnectionfunction)
  - [DeleteConnectionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteconnectiongroup)
  - [DeleteContinuousDeploymentPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletecontinuousdeploymentpolicy)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletedistribution)
  - [DeleteDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletedistributiontenant)
  - [DeleteFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletefieldlevelencryptionconfig)
  - [DeleteFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletefieldlevelencryptionprofile)
  - [DeleteFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletefunction)
  - [DeleteKeyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletekeygroup)
  - [DeleteKeyValueStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletekeyvaluestore)
  - [DeleteMonitoringSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletemonitoringsubscription)
  - [DeleteOriginAccessControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteoriginaccesscontrol)
  - [DeleteOriginRequestPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteoriginrequestpolicy)
  - [DeletePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletepublickey)
  - [DeleteRealtimeLogConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleterealtimelogconfig)
  - [DeleteResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteresourcepolicy)
  - [DeleteResponseHeadersPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deleteresponseheaderspolicy)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletestreamingdistribution)
  - [DeleteTrustStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletetruststore)
  - [DeleteVpcOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#deletevpcorigin)
  - [DescribeConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#describeconnectionfunction)
  - [DescribeFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#describefunction)
  - [DescribeKeyValueStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#describekeyvaluestore)
  - [DisassociateDistributionTenantWebACL](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#disassociatedistributiontenantwebacl)
  - [DisassociateDistributionWebACL](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#disassociatedistributionwebacl)
  - [GetAnycastIpList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getanycastiplist)
  - [GetCachePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcachepolicy)
  - [GetCachePolicyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcachepolicyconfig)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcloudfrontoriginaccessidentityconfig)
  - [GetConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getconnectionfunction)
  - [GetConnectionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getconnectiongroup)
  - [GetConnectionGroupByRoutingEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getconnectiongroupbyroutingendpoint)
  - [GetContinuousDeploymentPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcontinuousdeploymentpolicy)
  - [GetContinuousDeploymentPolicyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getcontinuousdeploymentpolicyconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getdistributionconfig)
  - [GetDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getdistributiontenant)
  - [GetDistributionTenantByDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getdistributiontenantbydomain)
  - [GetFieldLevelEncryption](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getfieldlevelencryption)
  - [GetFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getfieldlevelencryptionconfig)
  - [GetFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getfieldlevelencryptionprofile)
  - [GetFieldLevelEncryptionProfileConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getfieldlevelencryptionprofileconfig)
  - [GetFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getfunction)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getinvalidation)
  - [GetInvalidationForDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getinvalidationfordistributiontenant)
  - [GetKeyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getkeygroup)
  - [GetKeyGroupConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getkeygroupconfig)
  - [GetManagedCertificateDetails](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getmanagedcertificatedetails)
  - [GetMonitoringSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getmonitoringsubscription)
  - [GetOriginAccessControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getoriginaccesscontrol)
  - [GetOriginAccessControlConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getoriginaccesscontrolconfig)
  - [GetOriginRequestPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getoriginrequestpolicy)
  - [GetOriginRequestPolicyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getoriginrequestpolicyconfig)
  - [GetPublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getpublickey)
  - [GetPublicKeyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getpublickeyconfig)
  - [GetRealtimeLogConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getrealtimelogconfig)
  - [GetResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getresourcepolicy)
  - [GetResponseHeadersPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getresponseheaderspolicy)
  - [GetResponseHeadersPolicyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getresponseheaderspolicyconfig)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getstreamingdistributionconfig)
  - [GetTrustStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#gettruststore)
  - [GetVpcOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#getvpcorigin)
  - [ListAnycastIpLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listanycastiplists)
  - [ListCachePolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listcachepolicies)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listcloudfrontoriginaccessidentities)
  - [ListConflictingAliases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listconflictingaliases)
  - [ListConnectionFunctions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listconnectionfunctions)
  - [ListConnectionGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listconnectiongroups)
  - [ListContinuousDeploymentPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listcontinuousdeploymentpolicies)
  - [ListDistributionTenants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributiontenants)
  - [ListDistributionTenantsByCustomization](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributiontenantsbycustomization)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributions)
  - [ListDistributionsByAnycastIpListId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyanycastiplistid)
  - [ListDistributionsByCachePolicyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbycachepolicyid)
  - [ListDistributionsByConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyconnectionfunction)
  - [ListDistributionsByConnectionMode](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyconnectionmode)
  - [ListDistributionsByKeyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbykeygroup)
  - [ListDistributionsByOriginRequestPolicyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyoriginrequestpolicyid)
  - [ListDistributionsByOwnedResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyownedresource)
  - [ListDistributionsByRealtimeLogConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyrealtimelogconfig)
  - [ListDistributionsByResponseHeadersPolicyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyresponseheaderspolicyid)
  - [ListDistributionsByTrustStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbytruststore)
  - [ListDistributionsByVpcOriginId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbyvpcoriginid)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdistributionsbywebaclid)
  - [ListDomainConflicts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listdomainconflicts)
  - [ListFieldLevelEncryptionConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listfieldlevelencryptionconfigs)
  - [ListFieldLevelEncryptionProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listfieldlevelencryptionprofiles)
  - [ListFunctions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listfunctions)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listinvalidations)
  - [ListInvalidationsForDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listinvalidationsfordistributiontenant)
  - [ListKeyGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listkeygroups)
  - [ListKeyValueStores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listkeyvaluestores)
  - [ListOriginAccessControls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listoriginaccesscontrols)
  - [ListOriginRequestPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listoriginrequestpolicies)
  - [ListPublicKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listpublickeys)
  - [ListRealtimeLogConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listrealtimelogconfigs)
  - [ListResponseHeadersPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listresponseheaderspolicies)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listtagsforresource)
  - [ListTrustStores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listtruststores)
  - [ListVpcOrigins](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#listvpcorigins)
  - [PublishConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#publishconnectionfunction)
  - [PublishFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#publishfunction)
  - [PutResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#putresourcepolicy)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#tagresource)
  - [TestConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#testconnectionfunction)
  - [TestFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#testfunction)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#untagresource)
  - [UpdateAnycastIpList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateanycastiplist)
  - [UpdateCachePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatecachepolicy)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatecloudfrontoriginaccessidentity)
  - [UpdateConnectionFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateconnectionfunction)
  - [UpdateConnectionGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateconnectiongroup)
  - [UpdateContinuousDeploymentPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatecontinuousdeploymentpolicy)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatedistribution)
  - [UpdateDistributionTenant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatedistributiontenant)
  - [UpdateDistributionWithStagingConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatedistributionwithstagingconfig)
  - [UpdateDomainAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatedomainassociation)
  - [UpdateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatefieldlevelencryptionconfig)
  - [UpdateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatefieldlevelencryptionprofile)
  - [UpdateFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatefunction)
  - [UpdateKeyGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatekeygroup)
  - [UpdateKeyValueStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatekeyvaluestore)
  - [UpdateOriginAccessControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateoriginaccesscontrol)
  - [UpdateOriginRequestPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateoriginrequestpolicy)
  - [UpdatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatepublickey)
  - [UpdateRealtimeLogConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updaterealtimelogconfig)
  - [UpdateResponseHeadersPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updateresponseheaderspolicy)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatestreamingdistribution)
  - [UpdateTrustStore](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatetruststore)
  - [UpdateVpcOrigin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#updatevpcorigin)
  - [VerifyDnsConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2020-05-31.html#verifydnsconfiguration)
- [**2019-03-26**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createdistributionwithtags)
  - [CreateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createfieldlevelencryptionconfig)
  - [CreateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createfieldlevelencryptionprofile)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createinvalidation)
  - [CreatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createpublickey)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletedistribution)
  - [DeleteFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletefieldlevelencryptionconfig)
  - [DeleteFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletefieldlevelencryptionprofile)
  - [DeletePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletepublickey)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getdistributionconfig)
  - [GetFieldLevelEncryption](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getfieldlevelencryption)
  - [GetFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getfieldlevelencryptionconfig)
  - [GetFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getfieldlevelencryptionprofile)
  - [GetFieldLevelEncryptionProfileConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getfieldlevelencryptionprofileconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getinvalidation)
  - [GetPublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getpublickey)
  - [GetPublicKeyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getpublickeyconfig)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listdistributionsbywebaclid)
  - [ListFieldLevelEncryptionConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listfieldlevelencryptionconfigs)
  - [ListFieldLevelEncryptionProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listfieldlevelencryptionprofiles)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listinvalidations)
  - [ListPublicKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listpublickeys)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatedistribution)
  - [UpdateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatefieldlevelencryptionconfig)
  - [UpdateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatefieldlevelencryptionprofile)
  - [UpdatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatepublickey)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2019-03-26.html#updatestreamingdistribution)
- [**2018-11-05**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createdistributionwithtags)
  - [CreateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createfieldlevelencryptionconfig)
  - [CreateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createfieldlevelencryptionprofile)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createinvalidation)
  - [CreatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createpublickey)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletedistribution)
  - [DeleteFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletefieldlevelencryptionconfig)
  - [DeleteFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletefieldlevelencryptionprofile)
  - [DeletePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletepublickey)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getdistributionconfig)
  - [GetFieldLevelEncryption](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getfieldlevelencryption)
  - [GetFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getfieldlevelencryptionconfig)
  - [GetFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getfieldlevelencryptionprofile)
  - [GetFieldLevelEncryptionProfileConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getfieldlevelencryptionprofileconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getinvalidation)
  - [GetPublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getpublickey)
  - [GetPublicKeyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getpublickeyconfig)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listdistributionsbywebaclid)
  - [ListFieldLevelEncryptionConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listfieldlevelencryptionconfigs)
  - [ListFieldLevelEncryptionProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listfieldlevelencryptionprofiles)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listinvalidations)
  - [ListPublicKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listpublickeys)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatedistribution)
  - [UpdateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatefieldlevelencryptionconfig)
  - [UpdateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatefieldlevelencryptionprofile)
  - [UpdatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatepublickey)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-11-05.html#updatestreamingdistribution)
- [**2018-06-18**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createdistributionwithtags)
  - [CreateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createfieldlevelencryptionconfig)
  - [CreateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createfieldlevelencryptionprofile)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createinvalidation)
  - [CreatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createpublickey)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletedistribution)
  - [DeleteFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletefieldlevelencryptionconfig)
  - [DeleteFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletefieldlevelencryptionprofile)
  - [DeletePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletepublickey)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getdistributionconfig)
  - [GetFieldLevelEncryption](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getfieldlevelencryption)
  - [GetFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getfieldlevelencryptionconfig)
  - [GetFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getfieldlevelencryptionprofile)
  - [GetFieldLevelEncryptionProfileConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getfieldlevelencryptionprofileconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getinvalidation)
  - [GetPublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getpublickey)
  - [GetPublicKeyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getpublickeyconfig)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listdistributionsbywebaclid)
  - [ListFieldLevelEncryptionConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listfieldlevelencryptionconfigs)
  - [ListFieldLevelEncryptionProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listfieldlevelencryptionprofiles)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listinvalidations)
  - [ListPublicKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listpublickeys)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatedistribution)
  - [UpdateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatefieldlevelencryptionconfig)
  - [UpdateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatefieldlevelencryptionprofile)
  - [UpdatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatepublickey)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2018-06-18.html#updatestreamingdistribution)
- [**2017-10-30**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createdistributionwithtags)
  - [CreateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createfieldlevelencryptionconfig)
  - [CreateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createfieldlevelencryptionprofile)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createinvalidation)
  - [CreatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createpublickey)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletedistribution)
  - [DeleteFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletefieldlevelencryptionconfig)
  - [DeleteFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletefieldlevelencryptionprofile)
  - [DeletePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletepublickey)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getdistributionconfig)
  - [GetFieldLevelEncryption](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getfieldlevelencryption)
  - [GetFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getfieldlevelencryptionconfig)
  - [GetFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getfieldlevelencryptionprofile)
  - [GetFieldLevelEncryptionProfileConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getfieldlevelencryptionprofileconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getinvalidation)
  - [GetPublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getpublickey)
  - [GetPublicKeyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getpublickeyconfig)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listdistributionsbywebaclid)
  - [ListFieldLevelEncryptionConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listfieldlevelencryptionconfigs)
  - [ListFieldLevelEncryptionProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listfieldlevelencryptionprofiles)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listinvalidations)
  - [ListPublicKeys](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listpublickeys)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatedistribution)
  - [UpdateFieldLevelEncryptionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatefieldlevelencryptionconfig)
  - [UpdateFieldLevelEncryptionProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatefieldlevelencryptionprofile)
  - [UpdatePublicKey](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatepublickey)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-10-30.html#updatestreamingdistribution)
- [**2017-03-25**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#deletedistribution)
  - [DeleteServiceLinkedRole](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#deleteservicelinkedrole)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2017-03-25.html#updatestreamingdistribution)
- [**2016-11-25**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-11-25.html#updatestreamingdistribution)
- [**2016-09-29**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-29.html#updatestreamingdistribution)
- [**2016-09-07**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-09-07.html#updatestreamingdistribution)
- [**2016-08-20**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-20.html#updatestreamingdistribution)
- [**2016-08-01**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createdistribution)
  - [CreateDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createdistributionwithtags)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createstreamingdistribution)
  - [CreateStreamingDistributionWithTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#createstreamingdistributionwithtags)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#liststreamingdistributions)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#listtagsforresource)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#untagresource)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-08-01.html#updatestreamingdistribution)
- [**2016-01-28**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#createdistribution)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#createstreamingdistribution)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#liststreamingdistributions)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2016-01-28.html#updatestreamingdistribution)
- [**2015-07-27**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html)

  - [CreateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#createcloudfrontoriginaccessidentity)
  - [CreateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#createdistribution)
  - [CreateInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#createinvalidation)
  - [CreateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#createstreamingdistribution)
  - [DeleteCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#deletecloudfrontoriginaccessidentity)
  - [DeleteDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#deletedistribution)
  - [DeleteStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#deletestreamingdistribution)
  - [GetCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getcloudfrontoriginaccessidentity)
  - [GetCloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getcloudfrontoriginaccessidentityconfig)
  - [GetDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getdistribution)
  - [GetDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getdistributionconfig)
  - [GetInvalidation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getinvalidation)
  - [GetStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getstreamingdistribution)
  - [GetStreamingDistributionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#getstreamingdistributionconfig)
  - [ListCloudFrontOriginAccessIdentities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#listcloudfrontoriginaccessidentities)
  - [ListDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#listdistributions)
  - [ListDistributionsByWebACLId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#listdistributionsbywebaclid)
  - [ListInvalidations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#listinvalidations)
  - [ListStreamingDistributions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#liststreamingdistributions)
  - [UpdateCloudFrontOriginAccessIdentity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#updatecloudfrontoriginaccessidentity)
  - [UpdateDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#updatedistribution)
  - [UpdateStreamingDistribution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-cloudfront-2015-07-27.html#updatestreamingdistribution)

## Examples

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/cf-examples.html).

- [Managing CloudFront Distributions](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/cloudfront-example-distribution.html)
- [Managing CloudFront Invalidations](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/cloudfront-example-invalidation.html)
- [Signing CloudFront URLs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/cloudfront-example-signed-url.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-awsclient.md#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getSignedCookie()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#method_getSignedCookie)
: array<string\|int, mixed> Create a signed Amazon CloudFront cookie.[getSignedUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#method_getSignedUrl)
: string Create a signed Amazon CloudFront URL.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-awsclient.md\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

##### Tags  [header link](class-aws-awsclient.md\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing or
the service is not supported.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getSignedCookie()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#method_getSignedCookie)

Create a signed Amazon CloudFront cookie.

`
    public
                    getSignedCookie(array<string|int, mixed> $options) : array<string|int, mixed>`

This method accepts an array of configuration options:

- url: (string) URL of the resource being signed (can include query
string and wildcards). For example: http://d111111abcdef8.cloudfront.net/images/horizon.jpg?size=large&license=yes
- policy: (string) JSON policy. Use this option when creating a signed
URL for a custom policy.
- expires: (int) UTC Unix timestamp used when signing with a canned
policy. Not required when passing a custom 'policy' option.
- key\_pair\_id: (string) The ID of the key pair used to sign CloudFront
URLs for private distributions.
- private\_key: (string) The filepath ot the private key used to sign
CloudFront URLs for private distributions.

##### Parameters

$options
: array<string\|int, mixed>

Array of configuration options used when signing

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#method_getSignedCookie\#tags)

throwsInvalidArgumentException

if url, key\_pair\_id, or private\_key
were not specified.

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html)

##### Return values

array<string\|int, mixed>
—

Key => value pairs of signed cookies to set

#### getSignedUrl()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#method_getSignedUrl)

Create a signed Amazon CloudFront URL.

`
    public
                    getSignedUrl(array<string|int, mixed> $options) : string`

This method accepts an array of configuration options:

- url: (string) URL of the resource being signed (can include query
string and wildcards). For example: rtmp://s5c39gqb8ow64r.cloudfront.net/videos/mp3\_name.mp3
http://d111111abcdef8.cloudfront.net/images/horizon.jpg?size=large&license=yes
- policy: (string) JSON policy. Use this option when creating a signed
URL for a custom policy.
- expires: (int) UTC Unix timestamp used when signing with a canned
policy. Not required when passing a custom 'policy' option.
- key\_pair\_id: (string) The ID of the key pair used to sign CloudFront
URLs for private distributions.
- private\_key: (string) The filepath to the private key used to sign
CloudFront URLs for private distributions.

##### Parameters

$options
: array<string\|int, mixed>

Array of configuration options used when signing

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html\#method_getSignedUrl\#tags)

throwsInvalidArgumentException

if url, key\_pair\_id, or private\_key
were not specified.

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html)

##### Return values

string
—

Signed URL with authentication parameters

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-awsclient.md#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getSignedCookie()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#method_getSignedCookie)
  - [getSignedUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#method_getSignedUrl)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CloudFront.CloudFrontClient.html#top)
