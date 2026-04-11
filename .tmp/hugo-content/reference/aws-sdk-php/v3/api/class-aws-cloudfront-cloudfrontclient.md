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

- [**2020-05-31 (latest)**](api-cloudfront-2020-05-31.md)

  - [AssociateAlias](api-cloudfront-2020-05-31-associatealias.md)
  - [AssociateDistributionTenantWebACL](api-cloudfront-2020-05-31-associatedistributiontenantwebacl.md)
  - [AssociateDistributionWebACL](api-cloudfront-2020-05-31-associatedistributionwebacl.md)
  - [CopyDistribution](api-cloudfront-2020-05-31-copydistribution.md)
  - [CreateAnycastIpList](api-cloudfront-2020-05-31-createanycastiplist.md)
  - [CreateCachePolicy](api-cloudfront-2020-05-31-createcachepolicy.md)
  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2020-05-31-createcloudfrontoriginaccessidentity.md)
  - [CreateConnectionFunction](api-cloudfront-2020-05-31-createconnectionfunction.md)
  - [CreateConnectionGroup](api-cloudfront-2020-05-31-createconnectiongroup.md)
  - [CreateContinuousDeploymentPolicy](api-cloudfront-2020-05-31-createcontinuousdeploymentpolicy.md)
  - [CreateDistribution](api-cloudfront-2020-05-31-createdistribution.md)
  - [CreateDistributionTenant](api-cloudfront-2020-05-31-createdistributiontenant.md)
  - [CreateDistributionWithTags](api-cloudfront-2020-05-31-createdistributionwithtags.md)
  - [CreateFieldLevelEncryptionConfig](api-cloudfront-2020-05-31-createfieldlevelencryptionconfig.md)
  - [CreateFieldLevelEncryptionProfile](api-cloudfront-2020-05-31-createfieldlevelencryptionprofile.md)
  - [CreateFunction](api-cloudfront-2020-05-31-createfunction.md)
  - [CreateInvalidation](api-cloudfront-2020-05-31-createinvalidation.md)
  - [CreateInvalidationForDistributionTenant](api-cloudfront-2020-05-31-createinvalidationfordistributiontenant.md)
  - [CreateKeyGroup](api-cloudfront-2020-05-31-createkeygroup.md)
  - [CreateKeyValueStore](api-cloudfront-2020-05-31-createkeyvaluestore.md)
  - [CreateMonitoringSubscription](api-cloudfront-2020-05-31-createmonitoringsubscription.md)
  - [CreateOriginAccessControl](api-cloudfront-2020-05-31-createoriginaccesscontrol.md)
  - [CreateOriginRequestPolicy](api-cloudfront-2020-05-31-createoriginrequestpolicy.md)
  - [CreatePublicKey](api-cloudfront-2020-05-31-createpublickey.md)
  - [CreateRealtimeLogConfig](api-cloudfront-2020-05-31-createrealtimelogconfig.md)
  - [CreateResponseHeadersPolicy](api-cloudfront-2020-05-31-createresponseheaderspolicy.md)
  - [CreateStreamingDistribution](api-cloudfront-2020-05-31-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2020-05-31-createstreamingdistributionwithtags.md)
  - [CreateTrustStore](api-cloudfront-2020-05-31-createtruststore.md)
  - [CreateVpcOrigin](api-cloudfront-2020-05-31-createvpcorigin.md)
  - [DeleteAnycastIpList](api-cloudfront-2020-05-31-deleteanycastiplist.md)
  - [DeleteCachePolicy](api-cloudfront-2020-05-31-deletecachepolicy.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2020-05-31-deletecloudfrontoriginaccessidentity.md)
  - [DeleteConnectionFunction](api-cloudfront-2020-05-31-deleteconnectionfunction.md)
  - [DeleteConnectionGroup](api-cloudfront-2020-05-31-deleteconnectiongroup.md)
  - [DeleteContinuousDeploymentPolicy](api-cloudfront-2020-05-31-deletecontinuousdeploymentpolicy.md)
  - [DeleteDistribution](api-cloudfront-2020-05-31-deletedistribution.md)
  - [DeleteDistributionTenant](api-cloudfront-2020-05-31-deletedistributiontenant.md)
  - [DeleteFieldLevelEncryptionConfig](api-cloudfront-2020-05-31-deletefieldlevelencryptionconfig.md)
  - [DeleteFieldLevelEncryptionProfile](api-cloudfront-2020-05-31-deletefieldlevelencryptionprofile.md)
  - [DeleteFunction](api-cloudfront-2020-05-31-deletefunction.md)
  - [DeleteKeyGroup](api-cloudfront-2020-05-31-deletekeygroup.md)
  - [DeleteKeyValueStore](api-cloudfront-2020-05-31-deletekeyvaluestore.md)
  - [DeleteMonitoringSubscription](api-cloudfront-2020-05-31-deletemonitoringsubscription.md)
  - [DeleteOriginAccessControl](api-cloudfront-2020-05-31-deleteoriginaccesscontrol.md)
  - [DeleteOriginRequestPolicy](api-cloudfront-2020-05-31-deleteoriginrequestpolicy.md)
  - [DeletePublicKey](api-cloudfront-2020-05-31-deletepublickey.md)
  - [DeleteRealtimeLogConfig](api-cloudfront-2020-05-31-deleterealtimelogconfig.md)
  - [DeleteResourcePolicy](api-cloudfront-2020-05-31-deleteresourcepolicy.md)
  - [DeleteResponseHeadersPolicy](api-cloudfront-2020-05-31-deleteresponseheaderspolicy.md)
  - [DeleteStreamingDistribution](api-cloudfront-2020-05-31-deletestreamingdistribution.md)
  - [DeleteTrustStore](api-cloudfront-2020-05-31-deletetruststore.md)
  - [DeleteVpcOrigin](api-cloudfront-2020-05-31-deletevpcorigin.md)
  - [DescribeConnectionFunction](api-cloudfront-2020-05-31-describeconnectionfunction.md)
  - [DescribeFunction](api-cloudfront-2020-05-31-describefunction.md)
  - [DescribeKeyValueStore](api-cloudfront-2020-05-31-describekeyvaluestore.md)
  - [DisassociateDistributionTenantWebACL](api-cloudfront-2020-05-31-disassociatedistributiontenantwebacl.md)
  - [DisassociateDistributionWebACL](api-cloudfront-2020-05-31-disassociatedistributionwebacl.md)
  - [GetAnycastIpList](api-cloudfront-2020-05-31-getanycastiplist.md)
  - [GetCachePolicy](api-cloudfront-2020-05-31-getcachepolicy.md)
  - [GetCachePolicyConfig](api-cloudfront-2020-05-31-getcachepolicyconfig.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2020-05-31-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2020-05-31-getcloudfrontoriginaccessidentityconfig.md)
  - [GetConnectionFunction](api-cloudfront-2020-05-31-getconnectionfunction.md)
  - [GetConnectionGroup](api-cloudfront-2020-05-31-getconnectiongroup.md)
  - [GetConnectionGroupByRoutingEndpoint](api-cloudfront-2020-05-31-getconnectiongroupbyroutingendpoint.md)
  - [GetContinuousDeploymentPolicy](api-cloudfront-2020-05-31-getcontinuousdeploymentpolicy.md)
  - [GetContinuousDeploymentPolicyConfig](api-cloudfront-2020-05-31-getcontinuousdeploymentpolicyconfig.md)
  - [GetDistribution](api-cloudfront-2020-05-31-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2020-05-31-getdistributionconfig.md)
  - [GetDistributionTenant](api-cloudfront-2020-05-31-getdistributiontenant.md)
  - [GetDistributionTenantByDomain](api-cloudfront-2020-05-31-getdistributiontenantbydomain.md)
  - [GetFieldLevelEncryption](api-cloudfront-2020-05-31-getfieldlevelencryption.md)
  - [GetFieldLevelEncryptionConfig](api-cloudfront-2020-05-31-getfieldlevelencryptionconfig.md)
  - [GetFieldLevelEncryptionProfile](api-cloudfront-2020-05-31-getfieldlevelencryptionprofile.md)
  - [GetFieldLevelEncryptionProfileConfig](api-cloudfront-2020-05-31-getfieldlevelencryptionprofileconfig.md)
  - [GetFunction](api-cloudfront-2020-05-31-getfunction.md)
  - [GetInvalidation](api-cloudfront-2020-05-31-getinvalidation.md)
  - [GetInvalidationForDistributionTenant](api-cloudfront-2020-05-31-getinvalidationfordistributiontenant.md)
  - [GetKeyGroup](api-cloudfront-2020-05-31-getkeygroup.md)
  - [GetKeyGroupConfig](api-cloudfront-2020-05-31-getkeygroupconfig.md)
  - [GetManagedCertificateDetails](api-cloudfront-2020-05-31-getmanagedcertificatedetails.md)
  - [GetMonitoringSubscription](api-cloudfront-2020-05-31-getmonitoringsubscription.md)
  - [GetOriginAccessControl](api-cloudfront-2020-05-31-getoriginaccesscontrol.md)
  - [GetOriginAccessControlConfig](api-cloudfront-2020-05-31-getoriginaccesscontrolconfig.md)
  - [GetOriginRequestPolicy](api-cloudfront-2020-05-31-getoriginrequestpolicy.md)
  - [GetOriginRequestPolicyConfig](api-cloudfront-2020-05-31-getoriginrequestpolicyconfig.md)
  - [GetPublicKey](api-cloudfront-2020-05-31-getpublickey.md)
  - [GetPublicKeyConfig](api-cloudfront-2020-05-31-getpublickeyconfig.md)
  - [GetRealtimeLogConfig](api-cloudfront-2020-05-31-getrealtimelogconfig.md)
  - [GetResourcePolicy](api-cloudfront-2020-05-31-getresourcepolicy.md)
  - [GetResponseHeadersPolicy](api-cloudfront-2020-05-31-getresponseheaderspolicy.md)
  - [GetResponseHeadersPolicyConfig](api-cloudfront-2020-05-31-getresponseheaderspolicyconfig.md)
  - [GetStreamingDistribution](api-cloudfront-2020-05-31-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2020-05-31-getstreamingdistributionconfig.md)
  - [GetTrustStore](api-cloudfront-2020-05-31-gettruststore.md)
  - [GetVpcOrigin](api-cloudfront-2020-05-31-getvpcorigin.md)
  - [ListAnycastIpLists](api-cloudfront-2020-05-31-listanycastiplists.md)
  - [ListCachePolicies](api-cloudfront-2020-05-31-listcachepolicies.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2020-05-31-listcloudfrontoriginaccessidentities.md)
  - [ListConflictingAliases](api-cloudfront-2020-05-31-listconflictingaliases.md)
  - [ListConnectionFunctions](api-cloudfront-2020-05-31-listconnectionfunctions.md)
  - [ListConnectionGroups](api-cloudfront-2020-05-31-listconnectiongroups.md)
  - [ListContinuousDeploymentPolicies](api-cloudfront-2020-05-31-listcontinuousdeploymentpolicies.md)
  - [ListDistributionTenants](api-cloudfront-2020-05-31-listdistributiontenants.md)
  - [ListDistributionTenantsByCustomization](api-cloudfront-2020-05-31-listdistributiontenantsbycustomization.md)
  - [ListDistributions](api-cloudfront-2020-05-31-listdistributions.md)
  - [ListDistributionsByAnycastIpListId](api-cloudfront-2020-05-31-listdistributionsbyanycastiplistid.md)
  - [ListDistributionsByCachePolicyId](api-cloudfront-2020-05-31-listdistributionsbycachepolicyid.md)
  - [ListDistributionsByConnectionFunction](api-cloudfront-2020-05-31-listdistributionsbyconnectionfunction.md)
  - [ListDistributionsByConnectionMode](api-cloudfront-2020-05-31-listdistributionsbyconnectionmode.md)
  - [ListDistributionsByKeyGroup](api-cloudfront-2020-05-31-listdistributionsbykeygroup.md)
  - [ListDistributionsByOriginRequestPolicyId](api-cloudfront-2020-05-31-listdistributionsbyoriginrequestpolicyid.md)
  - [ListDistributionsByOwnedResource](api-cloudfront-2020-05-31-listdistributionsbyownedresource.md)
  - [ListDistributionsByRealtimeLogConfig](api-cloudfront-2020-05-31-listdistributionsbyrealtimelogconfig.md)
  - [ListDistributionsByResponseHeadersPolicyId](api-cloudfront-2020-05-31-listdistributionsbyresponseheaderspolicyid.md)
  - [ListDistributionsByTrustStore](api-cloudfront-2020-05-31-listdistributionsbytruststore.md)
  - [ListDistributionsByVpcOriginId](api-cloudfront-2020-05-31-listdistributionsbyvpcoriginid.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2020-05-31-listdistributionsbywebaclid.md)
  - [ListDomainConflicts](api-cloudfront-2020-05-31-listdomainconflicts.md)
  - [ListFieldLevelEncryptionConfigs](api-cloudfront-2020-05-31-listfieldlevelencryptionconfigs.md)
  - [ListFieldLevelEncryptionProfiles](api-cloudfront-2020-05-31-listfieldlevelencryptionprofiles.md)
  - [ListFunctions](api-cloudfront-2020-05-31-listfunctions.md)
  - [ListInvalidations](api-cloudfront-2020-05-31-listinvalidations.md)
  - [ListInvalidationsForDistributionTenant](api-cloudfront-2020-05-31-listinvalidationsfordistributiontenant.md)
  - [ListKeyGroups](api-cloudfront-2020-05-31-listkeygroups.md)
  - [ListKeyValueStores](api-cloudfront-2020-05-31-listkeyvaluestores.md)
  - [ListOriginAccessControls](api-cloudfront-2020-05-31-listoriginaccesscontrols.md)
  - [ListOriginRequestPolicies](api-cloudfront-2020-05-31-listoriginrequestpolicies.md)
  - [ListPublicKeys](api-cloudfront-2020-05-31-listpublickeys.md)
  - [ListRealtimeLogConfigs](api-cloudfront-2020-05-31-listrealtimelogconfigs.md)
  - [ListResponseHeadersPolicies](api-cloudfront-2020-05-31-listresponseheaderspolicies.md)
  - [ListStreamingDistributions](api-cloudfront-2020-05-31-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2020-05-31-listtagsforresource.md)
  - [ListTrustStores](api-cloudfront-2020-05-31-listtruststores.md)
  - [ListVpcOrigins](api-cloudfront-2020-05-31-listvpcorigins.md)
  - [PublishConnectionFunction](api-cloudfront-2020-05-31-publishconnectionfunction.md)
  - [PublishFunction](api-cloudfront-2020-05-31-publishfunction.md)
  - [PutResourcePolicy](api-cloudfront-2020-05-31-putresourcepolicy.md)
  - [TagResource](api-cloudfront-2020-05-31-tagresource.md)
  - [TestConnectionFunction](api-cloudfront-2020-05-31-testconnectionfunction.md)
  - [TestFunction](api-cloudfront-2020-05-31-testfunction.md)
  - [UntagResource](api-cloudfront-2020-05-31-untagresource.md)
  - [UpdateAnycastIpList](api-cloudfront-2020-05-31-updateanycastiplist.md)
  - [UpdateCachePolicy](api-cloudfront-2020-05-31-updatecachepolicy.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2020-05-31-updatecloudfrontoriginaccessidentity.md)
  - [UpdateConnectionFunction](api-cloudfront-2020-05-31-updateconnectionfunction.md)
  - [UpdateConnectionGroup](api-cloudfront-2020-05-31-updateconnectiongroup.md)
  - [UpdateContinuousDeploymentPolicy](api-cloudfront-2020-05-31-updatecontinuousdeploymentpolicy.md)
  - [UpdateDistribution](api-cloudfront-2020-05-31-updatedistribution.md)
  - [UpdateDistributionTenant](api-cloudfront-2020-05-31-updatedistributiontenant.md)
  - [UpdateDistributionWithStagingConfig](api-cloudfront-2020-05-31-updatedistributionwithstagingconfig.md)
  - [UpdateDomainAssociation](api-cloudfront-2020-05-31-updatedomainassociation.md)
  - [UpdateFieldLevelEncryptionConfig](api-cloudfront-2020-05-31-updatefieldlevelencryptionconfig.md)
  - [UpdateFieldLevelEncryptionProfile](api-cloudfront-2020-05-31-updatefieldlevelencryptionprofile.md)
  - [UpdateFunction](api-cloudfront-2020-05-31-updatefunction.md)
  - [UpdateKeyGroup](api-cloudfront-2020-05-31-updatekeygroup.md)
  - [UpdateKeyValueStore](api-cloudfront-2020-05-31-updatekeyvaluestore.md)
  - [UpdateOriginAccessControl](api-cloudfront-2020-05-31-updateoriginaccesscontrol.md)
  - [UpdateOriginRequestPolicy](api-cloudfront-2020-05-31-updateoriginrequestpolicy.md)
  - [UpdatePublicKey](api-cloudfront-2020-05-31-updatepublickey.md)
  - [UpdateRealtimeLogConfig](api-cloudfront-2020-05-31-updaterealtimelogconfig.md)
  - [UpdateResponseHeadersPolicy](api-cloudfront-2020-05-31-updateresponseheaderspolicy.md)
  - [UpdateStreamingDistribution](api-cloudfront-2020-05-31-updatestreamingdistribution.md)
  - [UpdateTrustStore](api-cloudfront-2020-05-31-updatetruststore.md)
  - [UpdateVpcOrigin](api-cloudfront-2020-05-31-updatevpcorigin.md)
  - [VerifyDnsConfiguration](api-cloudfront-2020-05-31-verifydnsconfiguration.md)
- [**2019-03-26**](api-cloudfront-2019-03-26.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2019-03-26-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2019-03-26-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2019-03-26-createdistributionwithtags.md)
  - [CreateFieldLevelEncryptionConfig](api-cloudfront-2019-03-26-createfieldlevelencryptionconfig.md)
  - [CreateFieldLevelEncryptionProfile](api-cloudfront-2019-03-26-createfieldlevelencryptionprofile.md)
  - [CreateInvalidation](api-cloudfront-2019-03-26-createinvalidation.md)
  - [CreatePublicKey](api-cloudfront-2019-03-26-createpublickey.md)
  - [CreateStreamingDistribution](api-cloudfront-2019-03-26-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2019-03-26-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2019-03-26-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2019-03-26-deletedistribution.md)
  - [DeleteFieldLevelEncryptionConfig](api-cloudfront-2019-03-26-deletefieldlevelencryptionconfig.md)
  - [DeleteFieldLevelEncryptionProfile](api-cloudfront-2019-03-26-deletefieldlevelencryptionprofile.md)
  - [DeletePublicKey](api-cloudfront-2019-03-26-deletepublickey.md)
  - [DeleteStreamingDistribution](api-cloudfront-2019-03-26-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2019-03-26-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2019-03-26-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2019-03-26-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2019-03-26-getdistributionconfig.md)
  - [GetFieldLevelEncryption](api-cloudfront-2019-03-26-getfieldlevelencryption.md)
  - [GetFieldLevelEncryptionConfig](api-cloudfront-2019-03-26-getfieldlevelencryptionconfig.md)
  - [GetFieldLevelEncryptionProfile](api-cloudfront-2019-03-26-getfieldlevelencryptionprofile.md)
  - [GetFieldLevelEncryptionProfileConfig](api-cloudfront-2019-03-26-getfieldlevelencryptionprofileconfig.md)
  - [GetInvalidation](api-cloudfront-2019-03-26-getinvalidation.md)
  - [GetPublicKey](api-cloudfront-2019-03-26-getpublickey.md)
  - [GetPublicKeyConfig](api-cloudfront-2019-03-26-getpublickeyconfig.md)
  - [GetStreamingDistribution](api-cloudfront-2019-03-26-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2019-03-26-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2019-03-26-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2019-03-26-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2019-03-26-listdistributionsbywebaclid.md)
  - [ListFieldLevelEncryptionConfigs](api-cloudfront-2019-03-26-listfieldlevelencryptionconfigs.md)
  - [ListFieldLevelEncryptionProfiles](api-cloudfront-2019-03-26-listfieldlevelencryptionprofiles.md)
  - [ListInvalidations](api-cloudfront-2019-03-26-listinvalidations.md)
  - [ListPublicKeys](api-cloudfront-2019-03-26-listpublickeys.md)
  - [ListStreamingDistributions](api-cloudfront-2019-03-26-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2019-03-26-listtagsforresource.md)
  - [TagResource](api-cloudfront-2019-03-26-tagresource.md)
  - [UntagResource](api-cloudfront-2019-03-26-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2019-03-26-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2019-03-26-updatedistribution.md)
  - [UpdateFieldLevelEncryptionConfig](api-cloudfront-2019-03-26-updatefieldlevelencryptionconfig.md)
  - [UpdateFieldLevelEncryptionProfile](api-cloudfront-2019-03-26-updatefieldlevelencryptionprofile.md)
  - [UpdatePublicKey](api-cloudfront-2019-03-26-updatepublickey.md)
  - [UpdateStreamingDistribution](api-cloudfront-2019-03-26-updatestreamingdistribution.md)
- [**2018-11-05**](api-cloudfront-2018-11-05.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2018-11-05-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2018-11-05-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2018-11-05-createdistributionwithtags.md)
  - [CreateFieldLevelEncryptionConfig](api-cloudfront-2018-11-05-createfieldlevelencryptionconfig.md)
  - [CreateFieldLevelEncryptionProfile](api-cloudfront-2018-11-05-createfieldlevelencryptionprofile.md)
  - [CreateInvalidation](api-cloudfront-2018-11-05-createinvalidation.md)
  - [CreatePublicKey](api-cloudfront-2018-11-05-createpublickey.md)
  - [CreateStreamingDistribution](api-cloudfront-2018-11-05-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2018-11-05-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2018-11-05-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2018-11-05-deletedistribution.md)
  - [DeleteFieldLevelEncryptionConfig](api-cloudfront-2018-11-05-deletefieldlevelencryptionconfig.md)
  - [DeleteFieldLevelEncryptionProfile](api-cloudfront-2018-11-05-deletefieldlevelencryptionprofile.md)
  - [DeletePublicKey](api-cloudfront-2018-11-05-deletepublickey.md)
  - [DeleteStreamingDistribution](api-cloudfront-2018-11-05-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2018-11-05-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2018-11-05-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2018-11-05-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2018-11-05-getdistributionconfig.md)
  - [GetFieldLevelEncryption](api-cloudfront-2018-11-05-getfieldlevelencryption.md)
  - [GetFieldLevelEncryptionConfig](api-cloudfront-2018-11-05-getfieldlevelencryptionconfig.md)
  - [GetFieldLevelEncryptionProfile](api-cloudfront-2018-11-05-getfieldlevelencryptionprofile.md)
  - [GetFieldLevelEncryptionProfileConfig](api-cloudfront-2018-11-05-getfieldlevelencryptionprofileconfig.md)
  - [GetInvalidation](api-cloudfront-2018-11-05-getinvalidation.md)
  - [GetPublicKey](api-cloudfront-2018-11-05-getpublickey.md)
  - [GetPublicKeyConfig](api-cloudfront-2018-11-05-getpublickeyconfig.md)
  - [GetStreamingDistribution](api-cloudfront-2018-11-05-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2018-11-05-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2018-11-05-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2018-11-05-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2018-11-05-listdistributionsbywebaclid.md)
  - [ListFieldLevelEncryptionConfigs](api-cloudfront-2018-11-05-listfieldlevelencryptionconfigs.md)
  - [ListFieldLevelEncryptionProfiles](api-cloudfront-2018-11-05-listfieldlevelencryptionprofiles.md)
  - [ListInvalidations](api-cloudfront-2018-11-05-listinvalidations.md)
  - [ListPublicKeys](api-cloudfront-2018-11-05-listpublickeys.md)
  - [ListStreamingDistributions](api-cloudfront-2018-11-05-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2018-11-05-listtagsforresource.md)
  - [TagResource](api-cloudfront-2018-11-05-tagresource.md)
  - [UntagResource](api-cloudfront-2018-11-05-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2018-11-05-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2018-11-05-updatedistribution.md)
  - [UpdateFieldLevelEncryptionConfig](api-cloudfront-2018-11-05-updatefieldlevelencryptionconfig.md)
  - [UpdateFieldLevelEncryptionProfile](api-cloudfront-2018-11-05-updatefieldlevelencryptionprofile.md)
  - [UpdatePublicKey](api-cloudfront-2018-11-05-updatepublickey.md)
  - [UpdateStreamingDistribution](api-cloudfront-2018-11-05-updatestreamingdistribution.md)
- [**2018-06-18**](api-cloudfront-2018-06-18.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2018-06-18-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2018-06-18-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2018-06-18-createdistributionwithtags.md)
  - [CreateFieldLevelEncryptionConfig](api-cloudfront-2018-06-18-createfieldlevelencryptionconfig.md)
  - [CreateFieldLevelEncryptionProfile](api-cloudfront-2018-06-18-createfieldlevelencryptionprofile.md)
  - [CreateInvalidation](api-cloudfront-2018-06-18-createinvalidation.md)
  - [CreatePublicKey](api-cloudfront-2018-06-18-createpublickey.md)
  - [CreateStreamingDistribution](api-cloudfront-2018-06-18-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2018-06-18-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2018-06-18-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2018-06-18-deletedistribution.md)
  - [DeleteFieldLevelEncryptionConfig](api-cloudfront-2018-06-18-deletefieldlevelencryptionconfig.md)
  - [DeleteFieldLevelEncryptionProfile](api-cloudfront-2018-06-18-deletefieldlevelencryptionprofile.md)
  - [DeletePublicKey](api-cloudfront-2018-06-18-deletepublickey.md)
  - [DeleteStreamingDistribution](api-cloudfront-2018-06-18-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2018-06-18-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2018-06-18-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2018-06-18-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2018-06-18-getdistributionconfig.md)
  - [GetFieldLevelEncryption](api-cloudfront-2018-06-18-getfieldlevelencryption.md)
  - [GetFieldLevelEncryptionConfig](api-cloudfront-2018-06-18-getfieldlevelencryptionconfig.md)
  - [GetFieldLevelEncryptionProfile](api-cloudfront-2018-06-18-getfieldlevelencryptionprofile.md)
  - [GetFieldLevelEncryptionProfileConfig](api-cloudfront-2018-06-18-getfieldlevelencryptionprofileconfig.md)
  - [GetInvalidation](api-cloudfront-2018-06-18-getinvalidation.md)
  - [GetPublicKey](api-cloudfront-2018-06-18-getpublickey.md)
  - [GetPublicKeyConfig](api-cloudfront-2018-06-18-getpublickeyconfig.md)
  - [GetStreamingDistribution](api-cloudfront-2018-06-18-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2018-06-18-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2018-06-18-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2018-06-18-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2018-06-18-listdistributionsbywebaclid.md)
  - [ListFieldLevelEncryptionConfigs](api-cloudfront-2018-06-18-listfieldlevelencryptionconfigs.md)
  - [ListFieldLevelEncryptionProfiles](api-cloudfront-2018-06-18-listfieldlevelencryptionprofiles.md)
  - [ListInvalidations](api-cloudfront-2018-06-18-listinvalidations.md)
  - [ListPublicKeys](api-cloudfront-2018-06-18-listpublickeys.md)
  - [ListStreamingDistributions](api-cloudfront-2018-06-18-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2018-06-18-listtagsforresource.md)
  - [TagResource](api-cloudfront-2018-06-18-tagresource.md)
  - [UntagResource](api-cloudfront-2018-06-18-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2018-06-18-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2018-06-18-updatedistribution.md)
  - [UpdateFieldLevelEncryptionConfig](api-cloudfront-2018-06-18-updatefieldlevelencryptionconfig.md)
  - [UpdateFieldLevelEncryptionProfile](api-cloudfront-2018-06-18-updatefieldlevelencryptionprofile.md)
  - [UpdatePublicKey](api-cloudfront-2018-06-18-updatepublickey.md)
  - [UpdateStreamingDistribution](api-cloudfront-2018-06-18-updatestreamingdistribution.md)
- [**2017-10-30**](api-cloudfront-2017-10-30.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2017-10-30-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2017-10-30-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2017-10-30-createdistributionwithtags.md)
  - [CreateFieldLevelEncryptionConfig](api-cloudfront-2017-10-30-createfieldlevelencryptionconfig.md)
  - [CreateFieldLevelEncryptionProfile](api-cloudfront-2017-10-30-createfieldlevelencryptionprofile.md)
  - [CreateInvalidation](api-cloudfront-2017-10-30-createinvalidation.md)
  - [CreatePublicKey](api-cloudfront-2017-10-30-createpublickey.md)
  - [CreateStreamingDistribution](api-cloudfront-2017-10-30-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2017-10-30-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2017-10-30-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2017-10-30-deletedistribution.md)
  - [DeleteFieldLevelEncryptionConfig](api-cloudfront-2017-10-30-deletefieldlevelencryptionconfig.md)
  - [DeleteFieldLevelEncryptionProfile](api-cloudfront-2017-10-30-deletefieldlevelencryptionprofile.md)
  - [DeletePublicKey](api-cloudfront-2017-10-30-deletepublickey.md)
  - [DeleteStreamingDistribution](api-cloudfront-2017-10-30-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2017-10-30-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2017-10-30-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2017-10-30-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2017-10-30-getdistributionconfig.md)
  - [GetFieldLevelEncryption](api-cloudfront-2017-10-30-getfieldlevelencryption.md)
  - [GetFieldLevelEncryptionConfig](api-cloudfront-2017-10-30-getfieldlevelencryptionconfig.md)
  - [GetFieldLevelEncryptionProfile](api-cloudfront-2017-10-30-getfieldlevelencryptionprofile.md)
  - [GetFieldLevelEncryptionProfileConfig](api-cloudfront-2017-10-30-getfieldlevelencryptionprofileconfig.md)
  - [GetInvalidation](api-cloudfront-2017-10-30-getinvalidation.md)
  - [GetPublicKey](api-cloudfront-2017-10-30-getpublickey.md)
  - [GetPublicKeyConfig](api-cloudfront-2017-10-30-getpublickeyconfig.md)
  - [GetStreamingDistribution](api-cloudfront-2017-10-30-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2017-10-30-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2017-10-30-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2017-10-30-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2017-10-30-listdistributionsbywebaclid.md)
  - [ListFieldLevelEncryptionConfigs](api-cloudfront-2017-10-30-listfieldlevelencryptionconfigs.md)
  - [ListFieldLevelEncryptionProfiles](api-cloudfront-2017-10-30-listfieldlevelencryptionprofiles.md)
  - [ListInvalidations](api-cloudfront-2017-10-30-listinvalidations.md)
  - [ListPublicKeys](api-cloudfront-2017-10-30-listpublickeys.md)
  - [ListStreamingDistributions](api-cloudfront-2017-10-30-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2017-10-30-listtagsforresource.md)
  - [TagResource](api-cloudfront-2017-10-30-tagresource.md)
  - [UntagResource](api-cloudfront-2017-10-30-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2017-10-30-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2017-10-30-updatedistribution.md)
  - [UpdateFieldLevelEncryptionConfig](api-cloudfront-2017-10-30-updatefieldlevelencryptionconfig.md)
  - [UpdateFieldLevelEncryptionProfile](api-cloudfront-2017-10-30-updatefieldlevelencryptionprofile.md)
  - [UpdatePublicKey](api-cloudfront-2017-10-30-updatepublickey.md)
  - [UpdateStreamingDistribution](api-cloudfront-2017-10-30-updatestreamingdistribution.md)
- [**2017-03-25**](api-cloudfront-2017-03-25.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2017-03-25-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2017-03-25-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2017-03-25-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2017-03-25-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2017-03-25-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2017-03-25-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2017-03-25-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2017-03-25-deletedistribution.md)
  - [DeleteServiceLinkedRole](api-cloudfront-2017-03-25-deleteservicelinkedrole.md)
  - [DeleteStreamingDistribution](api-cloudfront-2017-03-25-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2017-03-25-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2017-03-25-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2017-03-25-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2017-03-25-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2017-03-25-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2017-03-25-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2017-03-25-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2017-03-25-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2017-03-25-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2017-03-25-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2017-03-25-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2017-03-25-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2017-03-25-listtagsforresource.md)
  - [TagResource](api-cloudfront-2017-03-25-tagresource.md)
  - [UntagResource](api-cloudfront-2017-03-25-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2017-03-25-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2017-03-25-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2017-03-25-updatestreamingdistribution.md)
- [**2016-11-25**](api-cloudfront-2016-11-25.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-11-25-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-11-25-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2016-11-25-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2016-11-25-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-11-25-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2016-11-25-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-11-25-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-11-25-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-11-25-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-11-25-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-11-25-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-11-25-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-11-25-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-11-25-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-11-25-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-11-25-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-11-25-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-11-25-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-11-25-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-11-25-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-11-25-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2016-11-25-listtagsforresource.md)
  - [TagResource](api-cloudfront-2016-11-25-tagresource.md)
  - [UntagResource](api-cloudfront-2016-11-25-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-11-25-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-11-25-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-11-25-updatestreamingdistribution.md)
- [**2016-09-29**](api-cloudfront-2016-09-29.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-29-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-09-29-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2016-09-29-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2016-09-29-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-09-29-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2016-09-29-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-29-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-09-29-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-09-29-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-29-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-09-29-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-09-29-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-09-29-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-09-29-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-09-29-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-09-29-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-09-29-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-09-29-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-09-29-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-09-29-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-09-29-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2016-09-29-listtagsforresource.md)
  - [TagResource](api-cloudfront-2016-09-29-tagresource.md)
  - [UntagResource](api-cloudfront-2016-09-29-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-29-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-09-29-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-09-29-updatestreamingdistribution.md)
- [**2016-09-07**](api-cloudfront-2016-09-07.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-07-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-09-07-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2016-09-07-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2016-09-07-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-09-07-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2016-09-07-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-07-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-09-07-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-09-07-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-07-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-09-07-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-09-07-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-09-07-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-09-07-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-09-07-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-09-07-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-09-07-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-09-07-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-09-07-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-09-07-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-09-07-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2016-09-07-listtagsforresource.md)
  - [TagResource](api-cloudfront-2016-09-07-tagresource.md)
  - [UntagResource](api-cloudfront-2016-09-07-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-09-07-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-09-07-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-09-07-updatestreamingdistribution.md)
- [**2016-08-20**](api-cloudfront-2016-08-20.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-20-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-08-20-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2016-08-20-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2016-08-20-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-08-20-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2016-08-20-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-20-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-08-20-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-08-20-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-20-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-08-20-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-08-20-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-08-20-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-08-20-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-08-20-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-08-20-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-08-20-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-08-20-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-08-20-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-08-20-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-08-20-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2016-08-20-listtagsforresource.md)
  - [TagResource](api-cloudfront-2016-08-20-tagresource.md)
  - [UntagResource](api-cloudfront-2016-08-20-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-20-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-08-20-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-08-20-updatestreamingdistribution.md)
- [**2016-08-01**](api-cloudfront-2016-08-01.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-01-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-08-01-createdistribution.md)
  - [CreateDistributionWithTags](api-cloudfront-2016-08-01-createdistributionwithtags.md)
  - [CreateInvalidation](api-cloudfront-2016-08-01-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-08-01-createstreamingdistribution.md)
  - [CreateStreamingDistributionWithTags](api-cloudfront-2016-08-01-createstreamingdistributionwithtags.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-01-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-08-01-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-08-01-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-01-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-08-01-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-08-01-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-08-01-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-08-01-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-08-01-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-08-01-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-08-01-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-08-01-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-08-01-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-08-01-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-08-01-liststreamingdistributions.md)
  - [ListTagsForResource](api-cloudfront-2016-08-01-listtagsforresource.md)
  - [TagResource](api-cloudfront-2016-08-01-tagresource.md)
  - [UntagResource](api-cloudfront-2016-08-01-untagresource.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-08-01-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-08-01-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-08-01-updatestreamingdistribution.md)
- [**2016-01-28**](api-cloudfront-2016-01-28.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2016-01-28-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2016-01-28-createdistribution.md)
  - [CreateInvalidation](api-cloudfront-2016-01-28-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2016-01-28-createstreamingdistribution.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2016-01-28-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2016-01-28-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2016-01-28-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2016-01-28-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2016-01-28-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2016-01-28-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2016-01-28-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2016-01-28-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2016-01-28-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2016-01-28-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2016-01-28-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2016-01-28-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2016-01-28-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2016-01-28-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2016-01-28-liststreamingdistributions.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2016-01-28-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2016-01-28-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2016-01-28-updatestreamingdistribution.md)
- [**2015-07-27**](api-cloudfront-2015-07-27.md)

  - [CreateCloudFrontOriginAccessIdentity](api-cloudfront-2015-07-27-createcloudfrontoriginaccessidentity.md)
  - [CreateDistribution](api-cloudfront-2015-07-27-createdistribution.md)
  - [CreateInvalidation](api-cloudfront-2015-07-27-createinvalidation.md)
  - [CreateStreamingDistribution](api-cloudfront-2015-07-27-createstreamingdistribution.md)
  - [DeleteCloudFrontOriginAccessIdentity](api-cloudfront-2015-07-27-deletecloudfrontoriginaccessidentity.md)
  - [DeleteDistribution](api-cloudfront-2015-07-27-deletedistribution.md)
  - [DeleteStreamingDistribution](api-cloudfront-2015-07-27-deletestreamingdistribution.md)
  - [GetCloudFrontOriginAccessIdentity](api-cloudfront-2015-07-27-getcloudfrontoriginaccessidentity.md)
  - [GetCloudFrontOriginAccessIdentityConfig](api-cloudfront-2015-07-27-getcloudfrontoriginaccessidentityconfig.md)
  - [GetDistribution](api-cloudfront-2015-07-27-getdistribution.md)
  - [GetDistributionConfig](api-cloudfront-2015-07-27-getdistributionconfig.md)
  - [GetInvalidation](api-cloudfront-2015-07-27-getinvalidation.md)
  - [GetStreamingDistribution](api-cloudfront-2015-07-27-getstreamingdistribution.md)
  - [GetStreamingDistributionConfig](api-cloudfront-2015-07-27-getstreamingdistributionconfig.md)
  - [ListCloudFrontOriginAccessIdentities](api-cloudfront-2015-07-27-listcloudfrontoriginaccessidentities.md)
  - [ListDistributions](api-cloudfront-2015-07-27-listdistributions.md)
  - [ListDistributionsByWebACLId](api-cloudfront-2015-07-27-listdistributionsbywebaclid.md)
  - [ListInvalidations](api-cloudfront-2015-07-27-listinvalidations.md)
  - [ListStreamingDistributions](api-cloudfront-2015-07-27-liststreamingdistributions.md)
  - [UpdateCloudFrontOriginAccessIdentity](api-cloudfront-2015-07-27-updatecloudfrontoriginaccessidentity.md)
  - [UpdateDistribution](api-cloudfront-2015-07-27-updatedistribution.md)
  - [UpdateStreamingDistribution](api-cloudfront-2015-07-27-updatestreamingdistribution.md)

## Examples

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](../../../sdk-for-php/v3/developer-guide/cf-examples.md).

- [Managing CloudFront Distributions](../../../sdk-for-php/v3/developer-guide/cloudfront-example-distribution.md)
- [Managing CloudFront Invalidations](../../../sdk-for-php/v3/developer-guide/cloudfront-example-invalidation.md)
- [Signing CloudFront URLs](../../../sdk-for-php/v3/developer-guide/cloudfront-example-signed-url.md)

### Table of Contents  [header link](class-aws-cloudfront-cloudfrontclient-toc.md)

#### Methods  [header link](class-aws-cloudfront-cloudfrontclient-toc-methods.md)

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
: callable Get the signature\_provider function of the client.[getSignedCookie()](class-aws-cloudfront-cloudfrontclient-method-getsignedcookie.md)
: array<string\|int, mixed> Create a signed Amazon CloudFront cookie.[getSignedUrl()](class-aws-cloudfront-cloudfrontclient-method-getsignedurl.md)
: string Create a signed Amazon CloudFront URL.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-cloudfront-cloudfrontclient-methods.md)

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

#### getSignedCookie()  [header link](class-aws-cloudfront-cloudfrontclient-method-getsignedcookie.md)

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

##### Tags  [header link](class-aws-cloudfront-cloudfrontclient-method-getsignedcookie-tags.md)

throwsInvalidArgumentException

if url, key\_pair\_id, or private\_key
were not specified.

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](../../../../services/amazoncloudfront/latest/developerguide/workingwithstreamingdistributions.md)

##### Return values

array<string\|int, mixed>
—

Key => value pairs of signed cookies to set

#### getSignedUrl()  [header link](class-aws-cloudfront-cloudfrontclient-method-getsignedurl.md)

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

##### Tags  [header link](class-aws-cloudfront-cloudfrontclient-method-getsignedurl-tags.md)

throwsInvalidArgumentException

if url, key\_pair\_id, or private\_key
were not specified.

link[http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/WorkingWithStreamingDistributions.html](../../../../services/amazoncloudfront/latest/developerguide/workingwithstreamingdistributions.md)

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
  - [Methods](class-aws-cloudfront-cloudfrontclient-toc-methods.md)
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
  - [getSignedCookie()](class-aws-cloudfront-cloudfrontclient-method-getsignedcookie.md)
  - [getSignedUrl()](class-aws-cloudfront-cloudfrontclient-method-getsignedurl.md)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-cloudfront-cloudfrontclient-top.md)
