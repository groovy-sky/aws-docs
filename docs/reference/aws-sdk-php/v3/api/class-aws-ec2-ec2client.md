Menu

- [Aws](namespace-aws.md)
- [Ec2](namespace-aws-ec2.md)

## Ec2Client     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

Client used to interact with **Amazon EC2**.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2016-11-15 (latest)**](api-ec2-2016-11-15.md)

  - [AcceptAddressTransfer](api-ec2-2016-11-15-acceptaddresstransfer.md)
  - [AcceptCapacityReservationBillingOwnership](api-ec2-2016-11-15-acceptcapacityreservationbillingownership.md)
  - [AcceptReservedInstancesExchangeQuote](api-ec2-2016-11-15-acceptreservedinstancesexchangequote.md)
  - [AcceptTransitGatewayMulticastDomainAssociations](api-ec2-2016-11-15-accepttransitgatewaymulticastdomainassociations.md)
  - [AcceptTransitGatewayPeeringAttachment](api-ec2-2016-11-15-accepttransitgatewaypeeringattachment.md)
  - [AcceptTransitGatewayVpcAttachment](api-ec2-2016-11-15-accepttransitgatewayvpcattachment.md)
  - [AcceptVpcEndpointConnections](api-ec2-2016-11-15-acceptvpcendpointconnections.md)
  - [AcceptVpcPeeringConnection](api-ec2-2016-11-15-acceptvpcpeeringconnection.md)
  - [AdvertiseByoipCidr](api-ec2-2016-11-15-advertisebyoipcidr.md)
  - [AllocateAddress](api-ec2-2016-11-15-allocateaddress.md)
  - [AllocateHosts](api-ec2-2016-11-15-allocatehosts.md)
  - [AllocateIpamPoolCidr](api-ec2-2016-11-15-allocateipampoolcidr.md)
  - [ApplySecurityGroupsToClientVpnTargetNetwork](api-ec2-2016-11-15-applysecuritygroupstoclientvpntargetnetwork.md)
  - [AssignIpv6Addresses](api-ec2-2016-11-15-assignipv6addresses.md)
  - [AssignPrivateIpAddresses](api-ec2-2016-11-15-assignprivateipaddresses.md)
  - [AssignPrivateNatGatewayAddress](api-ec2-2016-11-15-assignprivatenatgatewayaddress.md)
  - [AssociateAddress](api-ec2-2016-11-15-associateaddress.md)
  - [AssociateCapacityReservationBillingOwner](api-ec2-2016-11-15-associatecapacityreservationbillingowner.md)
  - [AssociateClientVpnTargetNetwork](api-ec2-2016-11-15-associateclientvpntargetnetwork.md)
  - [AssociateDhcpOptions](api-ec2-2016-11-15-associatedhcpoptions.md)
  - [AssociateEnclaveCertificateIamRole](api-ec2-2016-11-15-associateenclavecertificateiamrole.md)
  - [AssociateIamInstanceProfile](api-ec2-2016-11-15-associateiaminstanceprofile.md)
  - [AssociateInstanceEventWindow](api-ec2-2016-11-15-associateinstanceeventwindow.md)
  - [AssociateIpamByoasn](api-ec2-2016-11-15-associateipambyoasn.md)
  - [AssociateIpamResourceDiscovery](api-ec2-2016-11-15-associateipamresourcediscovery.md)
  - [AssociateNatGatewayAddress](api-ec2-2016-11-15-associatenatgatewayaddress.md)
  - [AssociateRouteServer](api-ec2-2016-11-15-associaterouteserver.md)
  - [AssociateRouteTable](api-ec2-2016-11-15-associateroutetable.md)
  - [AssociateSecurityGroupVpc](api-ec2-2016-11-15-associatesecuritygroupvpc.md)
  - [AssociateSubnetCidrBlock](api-ec2-2016-11-15-associatesubnetcidrblock.md)
  - [AssociateTransitGatewayMulticastDomain](api-ec2-2016-11-15-associatetransitgatewaymulticastdomain.md)
  - [AssociateTransitGatewayPolicyTable](api-ec2-2016-11-15-associatetransitgatewaypolicytable.md)
  - [AssociateTransitGatewayRouteTable](api-ec2-2016-11-15-associatetransitgatewayroutetable.md)
  - [AssociateTrunkInterface](api-ec2-2016-11-15-associatetrunkinterface.md)
  - [AssociateVpcCidrBlock](api-ec2-2016-11-15-associatevpccidrblock.md)
  - [AttachClassicLinkVpc](api-ec2-2016-11-15-attachclassiclinkvpc.md)
  - [AttachInternetGateway](api-ec2-2016-11-15-attachinternetgateway.md)
  - [AttachNetworkInterface](api-ec2-2016-11-15-attachnetworkinterface.md)
  - [AttachVerifiedAccessTrustProvider](api-ec2-2016-11-15-attachverifiedaccesstrustprovider.md)
  - [AttachVolume](api-ec2-2016-11-15-attachvolume.md)
  - [AttachVpnGateway](api-ec2-2016-11-15-attachvpngateway.md)
  - [AuthorizeClientVpnIngress](api-ec2-2016-11-15-authorizeclientvpningress.md)
  - [AuthorizeSecurityGroupEgress](api-ec2-2016-11-15-authorizesecuritygroupegress.md)
  - [AuthorizeSecurityGroupIngress](api-ec2-2016-11-15-authorizesecuritygroupingress.md)
  - [BundleInstance](api-ec2-2016-11-15-bundleinstance.md)
  - [CancelBundleTask](api-ec2-2016-11-15-cancelbundletask.md)
  - [CancelCapacityReservation](api-ec2-2016-11-15-cancelcapacityreservation.md)
  - [CancelCapacityReservationFleets](api-ec2-2016-11-15-cancelcapacityreservationfleets.md)
  - [CancelConversionTask](api-ec2-2016-11-15-cancelconversiontask.md)
  - [CancelDeclarativePoliciesReport](api-ec2-2016-11-15-canceldeclarativepoliciesreport.md)
  - [CancelExportTask](api-ec2-2016-11-15-cancelexporttask.md)
  - [CancelImageLaunchPermission](api-ec2-2016-11-15-cancelimagelaunchpermission.md)
  - [CancelImportTask](api-ec2-2016-11-15-cancelimporttask.md)
  - [CancelReservedInstancesListing](api-ec2-2016-11-15-cancelreservedinstanceslisting.md)
  - [CancelSpotFleetRequests](api-ec2-2016-11-15-cancelspotfleetrequests.md)
  - [CancelSpotInstanceRequests](api-ec2-2016-11-15-cancelspotinstancerequests.md)
  - [ConfirmProductInstance](api-ec2-2016-11-15-confirmproductinstance.md)
  - [CopyFpgaImage](api-ec2-2016-11-15-copyfpgaimage.md)
  - [CopyImage](api-ec2-2016-11-15-copyimage.md)
  - [CopySnapshot](api-ec2-2016-11-15-copysnapshot.md)
  - [CopyVolumes](api-ec2-2016-11-15-copyvolumes.md)
  - [CreateCapacityManagerDataExport](api-ec2-2016-11-15-createcapacitymanagerdataexport.md)
  - [CreateCapacityReservation](api-ec2-2016-11-15-createcapacityreservation.md)
  - [CreateCapacityReservationBySplitting](api-ec2-2016-11-15-createcapacityreservationbysplitting.md)
  - [CreateCapacityReservationFleet](api-ec2-2016-11-15-createcapacityreservationfleet.md)
  - [CreateCarrierGateway](api-ec2-2016-11-15-createcarriergateway.md)
  - [CreateClientVpnEndpoint](api-ec2-2016-11-15-createclientvpnendpoint.md)
  - [CreateClientVpnRoute](api-ec2-2016-11-15-createclientvpnroute.md)
  - [CreateCoipCidr](api-ec2-2016-11-15-createcoipcidr.md)
  - [CreateCoipPool](api-ec2-2016-11-15-createcoippool.md)
  - [CreateCustomerGateway](api-ec2-2016-11-15-createcustomergateway.md)
  - [CreateDefaultSubnet](api-ec2-2016-11-15-createdefaultsubnet.md)
  - [CreateDefaultVpc](api-ec2-2016-11-15-createdefaultvpc.md)
  - [CreateDelegateMacVolumeOwnershipTask](api-ec2-2016-11-15-createdelegatemacvolumeownershiptask.md)
  - [CreateDhcpOptions](api-ec2-2016-11-15-createdhcpoptions.md)
  - [CreateEgressOnlyInternetGateway](api-ec2-2016-11-15-createegressonlyinternetgateway.md)
  - [CreateFleet](api-ec2-2016-11-15-createfleet.md)
  - [CreateFlowLogs](api-ec2-2016-11-15-createflowlogs.md)
  - [CreateFpgaImage](api-ec2-2016-11-15-createfpgaimage.md)
  - [CreateImage](api-ec2-2016-11-15-createimage.md)
  - [CreateImageUsageReport](api-ec2-2016-11-15-createimageusagereport.md)
  - [CreateInstanceConnectEndpoint](api-ec2-2016-11-15-createinstanceconnectendpoint.md)
  - [CreateInstanceEventWindow](api-ec2-2016-11-15-createinstanceeventwindow.md)
  - [CreateInstanceExportTask](api-ec2-2016-11-15-createinstanceexporttask.md)
  - [CreateInternetGateway](api-ec2-2016-11-15-createinternetgateway.md)
  - [CreateInterruptibleCapacityReservationAllocation](api-ec2-2016-11-15-createinterruptiblecapacityreservationallocation.md)
  - [CreateIpam](api-ec2-2016-11-15-createipam.md)
  - [CreateIpamExternalResourceVerificationToken](api-ec2-2016-11-15-createipamexternalresourceverificationtoken.md)
  - [CreateIpamPolicy](api-ec2-2016-11-15-createipampolicy.md)
  - [CreateIpamPool](api-ec2-2016-11-15-createipampool.md)
  - [CreateIpamPrefixListResolver](api-ec2-2016-11-15-createipamprefixlistresolver.md)
  - [CreateIpamPrefixListResolverTarget](api-ec2-2016-11-15-createipamprefixlistresolvertarget.md)
  - [CreateIpamResourceDiscovery](api-ec2-2016-11-15-createipamresourcediscovery.md)
  - [CreateIpamScope](api-ec2-2016-11-15-createipamscope.md)
  - [CreateKeyPair](api-ec2-2016-11-15-createkeypair.md)
  - [CreateLaunchTemplate](api-ec2-2016-11-15-createlaunchtemplate.md)
  - [CreateLaunchTemplateVersion](api-ec2-2016-11-15-createlaunchtemplateversion.md)
  - [CreateLocalGatewayRoute](api-ec2-2016-11-15-createlocalgatewayroute.md)
  - [CreateLocalGatewayRouteTable](api-ec2-2016-11-15-createlocalgatewayroutetable.md)
  - [CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation](api-ec2-2016-11-15-createlocalgatewayroutetablevirtualinterfacegroupassociation.md)
  - [CreateLocalGatewayRouteTableVpcAssociation](api-ec2-2016-11-15-createlocalgatewayroutetablevpcassociation.md)
  - [CreateLocalGatewayVirtualInterface](api-ec2-2016-11-15-createlocalgatewayvirtualinterface.md)
  - [CreateLocalGatewayVirtualInterfaceGroup](api-ec2-2016-11-15-createlocalgatewayvirtualinterfacegroup.md)
  - [CreateMacSystemIntegrityProtectionModificationTask](api-ec2-2016-11-15-createmacsystemintegrityprotectionmodificationtask.md)
  - [CreateManagedPrefixList](api-ec2-2016-11-15-createmanagedprefixlist.md)
  - [CreateNatGateway](api-ec2-2016-11-15-createnatgateway.md)
  - [CreateNetworkAcl](api-ec2-2016-11-15-createnetworkacl.md)
  - [CreateNetworkAclEntry](api-ec2-2016-11-15-createnetworkaclentry.md)
  - [CreateNetworkInsightsAccessScope](api-ec2-2016-11-15-createnetworkinsightsaccessscope.md)
  - [CreateNetworkInsightsPath](api-ec2-2016-11-15-createnetworkinsightspath.md)
  - [CreateNetworkInterface](api-ec2-2016-11-15-createnetworkinterface.md)
  - [CreateNetworkInterfacePermission](api-ec2-2016-11-15-createnetworkinterfacepermission.md)
  - [CreatePlacementGroup](api-ec2-2016-11-15-createplacementgroup.md)
  - [CreatePublicIpv4Pool](api-ec2-2016-11-15-createpublicipv4pool.md)
  - [CreateReplaceRootVolumeTask](api-ec2-2016-11-15-createreplacerootvolumetask.md)
  - [CreateReservedInstancesListing](api-ec2-2016-11-15-createreservedinstanceslisting.md)
  - [CreateRestoreImageTask](api-ec2-2016-11-15-createrestoreimagetask.md)
  - [CreateRoute](api-ec2-2016-11-15-createroute.md)
  - [CreateRouteServer](api-ec2-2016-11-15-createrouteserver.md)
  - [CreateRouteServerEndpoint](api-ec2-2016-11-15-createrouteserverendpoint.md)
  - [CreateRouteServerPeer](api-ec2-2016-11-15-createrouteserverpeer.md)
  - [CreateRouteTable](api-ec2-2016-11-15-createroutetable.md)
  - [CreateSecondaryNetwork](api-ec2-2016-11-15-createsecondarynetwork.md)
  - [CreateSecondarySubnet](api-ec2-2016-11-15-createsecondarysubnet.md)
  - [CreateSecurityGroup](api-ec2-2016-11-15-createsecuritygroup.md)
  - [CreateSnapshot](api-ec2-2016-11-15-createsnapshot.md)
  - [CreateSnapshots](api-ec2-2016-11-15-createsnapshots.md)
  - [CreateSpotDatafeedSubscription](api-ec2-2016-11-15-createspotdatafeedsubscription.md)
  - [CreateStoreImageTask](api-ec2-2016-11-15-createstoreimagetask.md)
  - [CreateSubnet](api-ec2-2016-11-15-createsubnet.md)
  - [CreateSubnetCidrReservation](api-ec2-2016-11-15-createsubnetcidrreservation.md)
  - [CreateTags](api-ec2-2016-11-15-createtags.md)
  - [CreateTrafficMirrorFilter](api-ec2-2016-11-15-createtrafficmirrorfilter.md)
  - [CreateTrafficMirrorFilterRule](api-ec2-2016-11-15-createtrafficmirrorfilterrule.md)
  - [CreateTrafficMirrorSession](api-ec2-2016-11-15-createtrafficmirrorsession.md)
  - [CreateTrafficMirrorTarget](api-ec2-2016-11-15-createtrafficmirrortarget.md)
  - [CreateTransitGateway](api-ec2-2016-11-15-createtransitgateway.md)
  - [CreateTransitGatewayConnect](api-ec2-2016-11-15-createtransitgatewayconnect.md)
  - [CreateTransitGatewayConnectPeer](api-ec2-2016-11-15-createtransitgatewayconnectpeer.md)
  - [CreateTransitGatewayMeteringPolicy](api-ec2-2016-11-15-createtransitgatewaymeteringpolicy.md)
  - [CreateTransitGatewayMeteringPolicyEntry](api-ec2-2016-11-15-createtransitgatewaymeteringpolicyentry.md)
  - [CreateTransitGatewayMulticastDomain](api-ec2-2016-11-15-createtransitgatewaymulticastdomain.md)
  - [CreateTransitGatewayPeeringAttachment](api-ec2-2016-11-15-createtransitgatewaypeeringattachment.md)
  - [CreateTransitGatewayPolicyTable](api-ec2-2016-11-15-createtransitgatewaypolicytable.md)
  - [CreateTransitGatewayPrefixListReference](api-ec2-2016-11-15-createtransitgatewayprefixlistreference.md)
  - [CreateTransitGatewayRoute](api-ec2-2016-11-15-createtransitgatewayroute.md)
  - [CreateTransitGatewayRouteTable](api-ec2-2016-11-15-createtransitgatewayroutetable.md)
  - [CreateTransitGatewayRouteTableAnnouncement](api-ec2-2016-11-15-createtransitgatewayroutetableannouncement.md)
  - [CreateTransitGatewayVpcAttachment](api-ec2-2016-11-15-createtransitgatewayvpcattachment.md)
  - [CreateVerifiedAccessEndpoint](api-ec2-2016-11-15-createverifiedaccessendpoint.md)
  - [CreateVerifiedAccessGroup](api-ec2-2016-11-15-createverifiedaccessgroup.md)
  - [CreateVerifiedAccessInstance](api-ec2-2016-11-15-createverifiedaccessinstance.md)
  - [CreateVerifiedAccessTrustProvider](api-ec2-2016-11-15-createverifiedaccesstrustprovider.md)
  - [CreateVolume](api-ec2-2016-11-15-createvolume.md)
  - [CreateVpc](api-ec2-2016-11-15-createvpc.md)
  - [CreateVpcBlockPublicAccessExclusion](api-ec2-2016-11-15-createvpcblockpublicaccessexclusion.md)
  - [CreateVpcEncryptionControl](api-ec2-2016-11-15-createvpcencryptioncontrol.md)
  - [CreateVpcEndpoint](api-ec2-2016-11-15-createvpcendpoint.md)
  - [CreateVpcEndpointConnectionNotification](api-ec2-2016-11-15-createvpcendpointconnectionnotification.md)
  - [CreateVpcEndpointServiceConfiguration](api-ec2-2016-11-15-createvpcendpointserviceconfiguration.md)
  - [CreateVpcPeeringConnection](api-ec2-2016-11-15-createvpcpeeringconnection.md)
  - [CreateVpnConcentrator](api-ec2-2016-11-15-createvpnconcentrator.md)
  - [CreateVpnConnection](api-ec2-2016-11-15-createvpnconnection.md)
  - [CreateVpnConnectionRoute](api-ec2-2016-11-15-createvpnconnectionroute.md)
  - [CreateVpnGateway](api-ec2-2016-11-15-createvpngateway.md)
  - [DeleteCapacityManagerDataExport](api-ec2-2016-11-15-deletecapacitymanagerdataexport.md)
  - [DeleteCarrierGateway](api-ec2-2016-11-15-deletecarriergateway.md)
  - [DeleteClientVpnEndpoint](api-ec2-2016-11-15-deleteclientvpnendpoint.md)
  - [DeleteClientVpnRoute](api-ec2-2016-11-15-deleteclientvpnroute.md)
  - [DeleteCoipCidr](api-ec2-2016-11-15-deletecoipcidr.md)
  - [DeleteCoipPool](api-ec2-2016-11-15-deletecoippool.md)
  - [DeleteCustomerGateway](api-ec2-2016-11-15-deletecustomergateway.md)
  - [DeleteDhcpOptions](api-ec2-2016-11-15-deletedhcpoptions.md)
  - [DeleteEgressOnlyInternetGateway](api-ec2-2016-11-15-deleteegressonlyinternetgateway.md)
  - [DeleteFleets](api-ec2-2016-11-15-deletefleets.md)
  - [DeleteFlowLogs](api-ec2-2016-11-15-deleteflowlogs.md)
  - [DeleteFpgaImage](api-ec2-2016-11-15-deletefpgaimage.md)
  - [DeleteImageUsageReport](api-ec2-2016-11-15-deleteimageusagereport.md)
  - [DeleteInstanceConnectEndpoint](api-ec2-2016-11-15-deleteinstanceconnectendpoint.md)
  - [DeleteInstanceEventWindow](api-ec2-2016-11-15-deleteinstanceeventwindow.md)
  - [DeleteInternetGateway](api-ec2-2016-11-15-deleteinternetgateway.md)
  - [DeleteIpam](api-ec2-2016-11-15-deleteipam.md)
  - [DeleteIpamExternalResourceVerificationToken](api-ec2-2016-11-15-deleteipamexternalresourceverificationtoken.md)
  - [DeleteIpamPolicy](api-ec2-2016-11-15-deleteipampolicy.md)
  - [DeleteIpamPool](api-ec2-2016-11-15-deleteipampool.md)
  - [DeleteIpamPrefixListResolver](api-ec2-2016-11-15-deleteipamprefixlistresolver.md)
  - [DeleteIpamPrefixListResolverTarget](api-ec2-2016-11-15-deleteipamprefixlistresolvertarget.md)
  - [DeleteIpamResourceDiscovery](api-ec2-2016-11-15-deleteipamresourcediscovery.md)
  - [DeleteIpamScope](api-ec2-2016-11-15-deleteipamscope.md)
  - [DeleteKeyPair](api-ec2-2016-11-15-deletekeypair.md)
  - [DeleteLaunchTemplate](api-ec2-2016-11-15-deletelaunchtemplate.md)
  - [DeleteLaunchTemplateVersions](api-ec2-2016-11-15-deletelaunchtemplateversions.md)
  - [DeleteLocalGatewayRoute](api-ec2-2016-11-15-deletelocalgatewayroute.md)
  - [DeleteLocalGatewayRouteTable](api-ec2-2016-11-15-deletelocalgatewayroutetable.md)
  - [DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation](api-ec2-2016-11-15-deletelocalgatewayroutetablevirtualinterfacegroupassociation.md)
  - [DeleteLocalGatewayRouteTableVpcAssociation](api-ec2-2016-11-15-deletelocalgatewayroutetablevpcassociation.md)
  - [DeleteLocalGatewayVirtualInterface](api-ec2-2016-11-15-deletelocalgatewayvirtualinterface.md)
  - [DeleteLocalGatewayVirtualInterfaceGroup](api-ec2-2016-11-15-deletelocalgatewayvirtualinterfacegroup.md)
  - [DeleteManagedPrefixList](api-ec2-2016-11-15-deletemanagedprefixlist.md)
  - [DeleteNatGateway](api-ec2-2016-11-15-deletenatgateway.md)
  - [DeleteNetworkAcl](api-ec2-2016-11-15-deletenetworkacl.md)
  - [DeleteNetworkAclEntry](api-ec2-2016-11-15-deletenetworkaclentry.md)
  - [DeleteNetworkInsightsAccessScope](api-ec2-2016-11-15-deletenetworkinsightsaccessscope.md)
  - [DeleteNetworkInsightsAccessScopeAnalysis](api-ec2-2016-11-15-deletenetworkinsightsaccessscopeanalysis.md)
  - [DeleteNetworkInsightsAnalysis](api-ec2-2016-11-15-deletenetworkinsightsanalysis.md)
  - [DeleteNetworkInsightsPath](api-ec2-2016-11-15-deletenetworkinsightspath.md)
  - [DeleteNetworkInterface](api-ec2-2016-11-15-deletenetworkinterface.md)
  - [DeleteNetworkInterfacePermission](api-ec2-2016-11-15-deletenetworkinterfacepermission.md)
  - [DeletePlacementGroup](api-ec2-2016-11-15-deleteplacementgroup.md)
  - [DeletePublicIpv4Pool](api-ec2-2016-11-15-deletepublicipv4pool.md)
  - [DeleteQueuedReservedInstances](api-ec2-2016-11-15-deletequeuedreservedinstances.md)
  - [DeleteRoute](api-ec2-2016-11-15-deleteroute.md)
  - [DeleteRouteServer](api-ec2-2016-11-15-deleterouteserver.md)
  - [DeleteRouteServerEndpoint](api-ec2-2016-11-15-deleterouteserverendpoint.md)
  - [DeleteRouteServerPeer](api-ec2-2016-11-15-deleterouteserverpeer.md)
  - [DeleteRouteTable](api-ec2-2016-11-15-deleteroutetable.md)
  - [DeleteSecondaryNetwork](api-ec2-2016-11-15-deletesecondarynetwork.md)
  - [DeleteSecondarySubnet](api-ec2-2016-11-15-deletesecondarysubnet.md)
  - [DeleteSecurityGroup](api-ec2-2016-11-15-deletesecuritygroup.md)
  - [DeleteSnapshot](api-ec2-2016-11-15-deletesnapshot.md)
  - [DeleteSpotDatafeedSubscription](api-ec2-2016-11-15-deletespotdatafeedsubscription.md)
  - [DeleteSubnet](api-ec2-2016-11-15-deletesubnet.md)
  - [DeleteSubnetCidrReservation](api-ec2-2016-11-15-deletesubnetcidrreservation.md)
  - [DeleteTags](api-ec2-2016-11-15-deletetags.md)
  - [DeleteTrafficMirrorFilter](api-ec2-2016-11-15-deletetrafficmirrorfilter.md)
  - [DeleteTrafficMirrorFilterRule](api-ec2-2016-11-15-deletetrafficmirrorfilterrule.md)
  - [DeleteTrafficMirrorSession](api-ec2-2016-11-15-deletetrafficmirrorsession.md)
  - [DeleteTrafficMirrorTarget](api-ec2-2016-11-15-deletetrafficmirrortarget.md)
  - [DeleteTransitGateway](api-ec2-2016-11-15-deletetransitgateway.md)
  - [DeleteTransitGatewayConnect](api-ec2-2016-11-15-deletetransitgatewayconnect.md)
  - [DeleteTransitGatewayConnectPeer](api-ec2-2016-11-15-deletetransitgatewayconnectpeer.md)
  - [DeleteTransitGatewayMeteringPolicy](api-ec2-2016-11-15-deletetransitgatewaymeteringpolicy.md)
  - [DeleteTransitGatewayMeteringPolicyEntry](api-ec2-2016-11-15-deletetransitgatewaymeteringpolicyentry.md)
  - [DeleteTransitGatewayMulticastDomain](api-ec2-2016-11-15-deletetransitgatewaymulticastdomain.md)
  - [DeleteTransitGatewayPeeringAttachment](api-ec2-2016-11-15-deletetransitgatewaypeeringattachment.md)
  - [DeleteTransitGatewayPolicyTable](api-ec2-2016-11-15-deletetransitgatewaypolicytable.md)
  - [DeleteTransitGatewayPrefixListReference](api-ec2-2016-11-15-deletetransitgatewayprefixlistreference.md)
  - [DeleteTransitGatewayRoute](api-ec2-2016-11-15-deletetransitgatewayroute.md)
  - [DeleteTransitGatewayRouteTable](api-ec2-2016-11-15-deletetransitgatewayroutetable.md)
  - [DeleteTransitGatewayRouteTableAnnouncement](api-ec2-2016-11-15-deletetransitgatewayroutetableannouncement.md)
  - [DeleteTransitGatewayVpcAttachment](api-ec2-2016-11-15-deletetransitgatewayvpcattachment.md)
  - [DeleteVerifiedAccessEndpoint](api-ec2-2016-11-15-deleteverifiedaccessendpoint.md)
  - [DeleteVerifiedAccessGroup](api-ec2-2016-11-15-deleteverifiedaccessgroup.md)
  - [DeleteVerifiedAccessInstance](api-ec2-2016-11-15-deleteverifiedaccessinstance.md)
  - [DeleteVerifiedAccessTrustProvider](api-ec2-2016-11-15-deleteverifiedaccesstrustprovider.md)
  - [DeleteVolume](api-ec2-2016-11-15-deletevolume.md)
  - [DeleteVpc](api-ec2-2016-11-15-deletevpc.md)
  - [DeleteVpcBlockPublicAccessExclusion](api-ec2-2016-11-15-deletevpcblockpublicaccessexclusion.md)
  - [DeleteVpcEncryptionControl](api-ec2-2016-11-15-deletevpcencryptioncontrol.md)
  - [DeleteVpcEndpointConnectionNotifications](api-ec2-2016-11-15-deletevpcendpointconnectionnotifications.md)
  - [DeleteVpcEndpointServiceConfigurations](api-ec2-2016-11-15-deletevpcendpointserviceconfigurations.md)
  - [DeleteVpcEndpoints](api-ec2-2016-11-15-deletevpcendpoints.md)
  - [DeleteVpcPeeringConnection](api-ec2-2016-11-15-deletevpcpeeringconnection.md)
  - [DeleteVpnConcentrator](api-ec2-2016-11-15-deletevpnconcentrator.md)
  - [DeleteVpnConnection](api-ec2-2016-11-15-deletevpnconnection.md)
  - [DeleteVpnConnectionRoute](api-ec2-2016-11-15-deletevpnconnectionroute.md)
  - [DeleteVpnGateway](api-ec2-2016-11-15-deletevpngateway.md)
  - [DeprovisionByoipCidr](api-ec2-2016-11-15-deprovisionbyoipcidr.md)
  - [DeprovisionIpamByoasn](api-ec2-2016-11-15-deprovisionipambyoasn.md)
  - [DeprovisionIpamPoolCidr](api-ec2-2016-11-15-deprovisionipampoolcidr.md)
  - [DeprovisionPublicIpv4PoolCidr](api-ec2-2016-11-15-deprovisionpublicipv4poolcidr.md)
  - [DeregisterImage](api-ec2-2016-11-15-deregisterimage.md)
  - [DeregisterInstanceEventNotificationAttributes](api-ec2-2016-11-15-deregisterinstanceeventnotificationattributes.md)
  - [DeregisterTransitGatewayMulticastGroupMembers](api-ec2-2016-11-15-deregistertransitgatewaymulticastgroupmembers.md)
  - [DeregisterTransitGatewayMulticastGroupSources](api-ec2-2016-11-15-deregistertransitgatewaymulticastgroupsources.md)
  - [DescribeAccountAttributes](api-ec2-2016-11-15-describeaccountattributes.md)
  - [DescribeAddressTransfers](api-ec2-2016-11-15-describeaddresstransfers.md)
  - [DescribeAddresses](api-ec2-2016-11-15-describeaddresses.md)
  - [DescribeAddressesAttribute](api-ec2-2016-11-15-describeaddressesattribute.md)
  - [DescribeAggregateIdFormat](api-ec2-2016-11-15-describeaggregateidformat.md)
  - [DescribeAvailabilityZones](api-ec2-2016-11-15-describeavailabilityzones.md)
  - [DescribeAwsNetworkPerformanceMetricSubscriptions](api-ec2-2016-11-15-describeawsnetworkperformancemetricsubscriptions.md)
  - [DescribeBundleTasks](api-ec2-2016-11-15-describebundletasks.md)
  - [DescribeByoipCidrs](api-ec2-2016-11-15-describebyoipcidrs.md)
  - [DescribeCapacityBlockExtensionHistory](api-ec2-2016-11-15-describecapacityblockextensionhistory.md)
  - [DescribeCapacityBlockExtensionOfferings](api-ec2-2016-11-15-describecapacityblockextensionofferings.md)
  - [DescribeCapacityBlockOfferings](api-ec2-2016-11-15-describecapacityblockofferings.md)
  - [DescribeCapacityBlockStatus](api-ec2-2016-11-15-describecapacityblockstatus.md)
  - [DescribeCapacityBlocks](api-ec2-2016-11-15-describecapacityblocks.md)
  - [DescribeCapacityManagerDataExports](api-ec2-2016-11-15-describecapacitymanagerdataexports.md)
  - [DescribeCapacityReservationBillingRequests](api-ec2-2016-11-15-describecapacityreservationbillingrequests.md)
  - [DescribeCapacityReservationFleets](api-ec2-2016-11-15-describecapacityreservationfleets.md)
  - [DescribeCapacityReservationTopology](api-ec2-2016-11-15-describecapacityreservationtopology.md)
  - [DescribeCapacityReservations](api-ec2-2016-11-15-describecapacityreservations.md)
  - [DescribeCarrierGateways](api-ec2-2016-11-15-describecarriergateways.md)
  - [DescribeClassicLinkInstances](api-ec2-2016-11-15-describeclassiclinkinstances.md)
  - [DescribeClientVpnAuthorizationRules](api-ec2-2016-11-15-describeclientvpnauthorizationrules.md)
  - [DescribeClientVpnConnections](api-ec2-2016-11-15-describeclientvpnconnections.md)
  - [DescribeClientVpnEndpoints](api-ec2-2016-11-15-describeclientvpnendpoints.md)
  - [DescribeClientVpnRoutes](api-ec2-2016-11-15-describeclientvpnroutes.md)
  - [DescribeClientVpnTargetNetworks](api-ec2-2016-11-15-describeclientvpntargetnetworks.md)
  - [DescribeCoipPools](api-ec2-2016-11-15-describecoippools.md)
  - [DescribeConversionTasks](api-ec2-2016-11-15-describeconversiontasks.md)
  - [DescribeCustomerGateways](api-ec2-2016-11-15-describecustomergateways.md)
  - [DescribeDeclarativePoliciesReports](api-ec2-2016-11-15-describedeclarativepoliciesreports.md)
  - [DescribeDhcpOptions](api-ec2-2016-11-15-describedhcpoptions.md)
  - [DescribeEgressOnlyInternetGateways](api-ec2-2016-11-15-describeegressonlyinternetgateways.md)
  - [DescribeElasticGpus](api-ec2-2016-11-15-describeelasticgpus.md)
  - [DescribeExportImageTasks](api-ec2-2016-11-15-describeexportimagetasks.md)
  - [DescribeExportTasks](api-ec2-2016-11-15-describeexporttasks.md)
  - [DescribeFastLaunchImages](api-ec2-2016-11-15-describefastlaunchimages.md)
  - [DescribeFastSnapshotRestores](api-ec2-2016-11-15-describefastsnapshotrestores.md)
  - [DescribeFleetHistory](api-ec2-2016-11-15-describefleethistory.md)
  - [DescribeFleetInstances](api-ec2-2016-11-15-describefleetinstances.md)
  - [DescribeFleets](api-ec2-2016-11-15-describefleets.md)
  - [DescribeFlowLogs](api-ec2-2016-11-15-describeflowlogs.md)
  - [DescribeFpgaImageAttribute](api-ec2-2016-11-15-describefpgaimageattribute.md)
  - [DescribeFpgaImages](api-ec2-2016-11-15-describefpgaimages.md)
  - [DescribeHostReservationOfferings](api-ec2-2016-11-15-describehostreservationofferings.md)
  - [DescribeHostReservations](api-ec2-2016-11-15-describehostreservations.md)
  - [DescribeHosts](api-ec2-2016-11-15-describehosts.md)
  - [DescribeIamInstanceProfileAssociations](api-ec2-2016-11-15-describeiaminstanceprofileassociations.md)
  - [DescribeIdFormat](api-ec2-2016-11-15-describeidformat.md)
  - [DescribeIdentityIdFormat](api-ec2-2016-11-15-describeidentityidformat.md)
  - [DescribeImageAttribute](api-ec2-2016-11-15-describeimageattribute.md)
  - [DescribeImageReferences](api-ec2-2016-11-15-describeimagereferences.md)
  - [DescribeImageUsageReportEntries](api-ec2-2016-11-15-describeimageusagereportentries.md)
  - [DescribeImageUsageReports](api-ec2-2016-11-15-describeimageusagereports.md)
  - [DescribeImages](api-ec2-2016-11-15-describeimages.md)
  - [DescribeImportImageTasks](api-ec2-2016-11-15-describeimportimagetasks.md)
  - [DescribeImportSnapshotTasks](api-ec2-2016-11-15-describeimportsnapshottasks.md)
  - [DescribeInstanceAttribute](api-ec2-2016-11-15-describeinstanceattribute.md)
  - [DescribeInstanceConnectEndpoints](api-ec2-2016-11-15-describeinstanceconnectendpoints.md)
  - [DescribeInstanceCreditSpecifications](api-ec2-2016-11-15-describeinstancecreditspecifications.md)
  - [DescribeInstanceEventNotificationAttributes](api-ec2-2016-11-15-describeinstanceeventnotificationattributes.md)
  - [DescribeInstanceEventWindows](api-ec2-2016-11-15-describeinstanceeventwindows.md)
  - [DescribeInstanceImageMetadata](api-ec2-2016-11-15-describeinstanceimagemetadata.md)
  - [DescribeInstanceSqlHaHistoryStates](api-ec2-2016-11-15-describeinstancesqlhahistorystates.md)
  - [DescribeInstanceSqlHaStates](api-ec2-2016-11-15-describeinstancesqlhastates.md)
  - [DescribeInstanceStatus](api-ec2-2016-11-15-describeinstancestatus.md)
  - [DescribeInstanceTopology](api-ec2-2016-11-15-describeinstancetopology.md)
  - [DescribeInstanceTypeOfferings](api-ec2-2016-11-15-describeinstancetypeofferings.md)
  - [DescribeInstanceTypes](api-ec2-2016-11-15-describeinstancetypes.md)
  - [DescribeInstances](api-ec2-2016-11-15-describeinstances.md)
  - [DescribeInternetGateways](api-ec2-2016-11-15-describeinternetgateways.md)
  - [DescribeIpamByoasn](api-ec2-2016-11-15-describeipambyoasn.md)
  - [DescribeIpamExternalResourceVerificationTokens](api-ec2-2016-11-15-describeipamexternalresourceverificationtokens.md)
  - [DescribeIpamPolicies](api-ec2-2016-11-15-describeipampolicies.md)
  - [DescribeIpamPools](api-ec2-2016-11-15-describeipampools.md)
  - [DescribeIpamPrefixListResolverTargets](api-ec2-2016-11-15-describeipamprefixlistresolvertargets.md)
  - [DescribeIpamPrefixListResolvers](api-ec2-2016-11-15-describeipamprefixlistresolvers.md)
  - [DescribeIpamResourceDiscoveries](api-ec2-2016-11-15-describeipamresourcediscoveries.md)
  - [DescribeIpamResourceDiscoveryAssociations](api-ec2-2016-11-15-describeipamresourcediscoveryassociations.md)
  - [DescribeIpamScopes](api-ec2-2016-11-15-describeipamscopes.md)
  - [DescribeIpams](api-ec2-2016-11-15-describeipams.md)
  - [DescribeIpv6Pools](api-ec2-2016-11-15-describeipv6pools.md)
  - [DescribeKeyPairs](api-ec2-2016-11-15-describekeypairs.md)
  - [DescribeLaunchTemplateVersions](api-ec2-2016-11-15-describelaunchtemplateversions.md)
  - [DescribeLaunchTemplates](api-ec2-2016-11-15-describelaunchtemplates.md)
  - [DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations](api-ec2-2016-11-15-describelocalgatewayroutetablevirtualinterfacegroupassociations.md)
  - [DescribeLocalGatewayRouteTableVpcAssociations](api-ec2-2016-11-15-describelocalgatewayroutetablevpcassociations.md)
  - [DescribeLocalGatewayRouteTables](api-ec2-2016-11-15-describelocalgatewayroutetables.md)
  - [DescribeLocalGatewayVirtualInterfaceGroups](api-ec2-2016-11-15-describelocalgatewayvirtualinterfacegroups.md)
  - [DescribeLocalGatewayVirtualInterfaces](api-ec2-2016-11-15-describelocalgatewayvirtualinterfaces.md)
  - [DescribeLocalGateways](api-ec2-2016-11-15-describelocalgateways.md)
  - [DescribeLockedSnapshots](api-ec2-2016-11-15-describelockedsnapshots.md)
  - [DescribeMacHosts](api-ec2-2016-11-15-describemachosts.md)
  - [DescribeMacModificationTasks](api-ec2-2016-11-15-describemacmodificationtasks.md)
  - [DescribeManagedPrefixLists](api-ec2-2016-11-15-describemanagedprefixlists.md)
  - [DescribeMovingAddresses](api-ec2-2016-11-15-describemovingaddresses.md)
  - [DescribeNatGateways](api-ec2-2016-11-15-describenatgateways.md)
  - [DescribeNetworkAcls](api-ec2-2016-11-15-describenetworkacls.md)
  - [DescribeNetworkInsightsAccessScopeAnalyses](api-ec2-2016-11-15-describenetworkinsightsaccessscopeanalyses.md)
  - [DescribeNetworkInsightsAccessScopes](api-ec2-2016-11-15-describenetworkinsightsaccessscopes.md)
  - [DescribeNetworkInsightsAnalyses](api-ec2-2016-11-15-describenetworkinsightsanalyses.md)
  - [DescribeNetworkInsightsPaths](api-ec2-2016-11-15-describenetworkinsightspaths.md)
  - [DescribeNetworkInterfaceAttribute](api-ec2-2016-11-15-describenetworkinterfaceattribute.md)
  - [DescribeNetworkInterfacePermissions](api-ec2-2016-11-15-describenetworkinterfacepermissions.md)
  - [DescribeNetworkInterfaces](api-ec2-2016-11-15-describenetworkinterfaces.md)
  - [DescribeOutpostLags](api-ec2-2016-11-15-describeoutpostlags.md)
  - [DescribePlacementGroups](api-ec2-2016-11-15-describeplacementgroups.md)
  - [DescribePrefixLists](api-ec2-2016-11-15-describeprefixlists.md)
  - [DescribePrincipalIdFormat](api-ec2-2016-11-15-describeprincipalidformat.md)
  - [DescribePublicIpv4Pools](api-ec2-2016-11-15-describepublicipv4pools.md)
  - [DescribeRegions](api-ec2-2016-11-15-describeregions.md)
  - [DescribeReplaceRootVolumeTasks](api-ec2-2016-11-15-describereplacerootvolumetasks.md)
  - [DescribeReservedInstances](api-ec2-2016-11-15-describereservedinstances.md)
  - [DescribeReservedInstancesListings](api-ec2-2016-11-15-describereservedinstanceslistings.md)
  - [DescribeReservedInstancesModifications](api-ec2-2016-11-15-describereservedinstancesmodifications.md)
  - [DescribeReservedInstancesOfferings](api-ec2-2016-11-15-describereservedinstancesofferings.md)
  - [DescribeRouteServerEndpoints](api-ec2-2016-11-15-describerouteserverendpoints.md)
  - [DescribeRouteServerPeers](api-ec2-2016-11-15-describerouteserverpeers.md)
  - [DescribeRouteServers](api-ec2-2016-11-15-describerouteservers.md)
  - [DescribeRouteTables](api-ec2-2016-11-15-describeroutetables.md)
  - [DescribeScheduledInstanceAvailability](api-ec2-2016-11-15-describescheduledinstanceavailability.md)
  - [DescribeScheduledInstances](api-ec2-2016-11-15-describescheduledinstances.md)
  - [DescribeSecondaryInterfaces](api-ec2-2016-11-15-describesecondaryinterfaces.md)
  - [DescribeSecondaryNetworks](api-ec2-2016-11-15-describesecondarynetworks.md)
  - [DescribeSecondarySubnets](api-ec2-2016-11-15-describesecondarysubnets.md)
  - [DescribeSecurityGroupReferences](api-ec2-2016-11-15-describesecuritygroupreferences.md)
  - [DescribeSecurityGroupRules](api-ec2-2016-11-15-describesecuritygrouprules.md)
  - [DescribeSecurityGroupVpcAssociations](api-ec2-2016-11-15-describesecuritygroupvpcassociations.md)
  - [DescribeSecurityGroups](api-ec2-2016-11-15-describesecuritygroups.md)
  - [DescribeServiceLinkVirtualInterfaces](api-ec2-2016-11-15-describeservicelinkvirtualinterfaces.md)
  - [DescribeSnapshotAttribute](api-ec2-2016-11-15-describesnapshotattribute.md)
  - [DescribeSnapshotTierStatus](api-ec2-2016-11-15-describesnapshottierstatus.md)
  - [DescribeSnapshots](api-ec2-2016-11-15-describesnapshots.md)
  - [DescribeSpotDatafeedSubscription](api-ec2-2016-11-15-describespotdatafeedsubscription.md)
  - [DescribeSpotFleetInstances](api-ec2-2016-11-15-describespotfleetinstances.md)
  - [DescribeSpotFleetRequestHistory](api-ec2-2016-11-15-describespotfleetrequesthistory.md)
  - [DescribeSpotFleetRequests](api-ec2-2016-11-15-describespotfleetrequests.md)
  - [DescribeSpotInstanceRequests](api-ec2-2016-11-15-describespotinstancerequests.md)
  - [DescribeSpotPriceHistory](api-ec2-2016-11-15-describespotpricehistory.md)
  - [DescribeStaleSecurityGroups](api-ec2-2016-11-15-describestalesecuritygroups.md)
  - [DescribeStoreImageTasks](api-ec2-2016-11-15-describestoreimagetasks.md)
  - [DescribeSubnets](api-ec2-2016-11-15-describesubnets.md)
  - [DescribeTags](api-ec2-2016-11-15-describetags.md)
  - [DescribeTrafficMirrorFilterRules](api-ec2-2016-11-15-describetrafficmirrorfilterrules.md)
  - [DescribeTrafficMirrorFilters](api-ec2-2016-11-15-describetrafficmirrorfilters.md)
  - [DescribeTrafficMirrorSessions](api-ec2-2016-11-15-describetrafficmirrorsessions.md)
  - [DescribeTrafficMirrorTargets](api-ec2-2016-11-15-describetrafficmirrortargets.md)
  - [DescribeTransitGatewayAttachments](api-ec2-2016-11-15-describetransitgatewayattachments.md)
  - [DescribeTransitGatewayConnectPeers](api-ec2-2016-11-15-describetransitgatewayconnectpeers.md)
  - [DescribeTransitGatewayConnects](api-ec2-2016-11-15-describetransitgatewayconnects.md)
  - [DescribeTransitGatewayMeteringPolicies](api-ec2-2016-11-15-describetransitgatewaymeteringpolicies.md)
  - [DescribeTransitGatewayMulticastDomains](api-ec2-2016-11-15-describetransitgatewaymulticastdomains.md)
  - [DescribeTransitGatewayPeeringAttachments](api-ec2-2016-11-15-describetransitgatewaypeeringattachments.md)
  - [DescribeTransitGatewayPolicyTables](api-ec2-2016-11-15-describetransitgatewaypolicytables.md)
  - [DescribeTransitGatewayRouteTableAnnouncements](api-ec2-2016-11-15-describetransitgatewayroutetableannouncements.md)
  - [DescribeTransitGatewayRouteTables](api-ec2-2016-11-15-describetransitgatewayroutetables.md)
  - [DescribeTransitGatewayVpcAttachments](api-ec2-2016-11-15-describetransitgatewayvpcattachments.md)
  - [DescribeTransitGateways](api-ec2-2016-11-15-describetransitgateways.md)
  - [DescribeTrunkInterfaceAssociations](api-ec2-2016-11-15-describetrunkinterfaceassociations.md)
  - [DescribeVerifiedAccessEndpoints](api-ec2-2016-11-15-describeverifiedaccessendpoints.md)
  - [DescribeVerifiedAccessGroups](api-ec2-2016-11-15-describeverifiedaccessgroups.md)
  - [DescribeVerifiedAccessInstanceLoggingConfigurations](api-ec2-2016-11-15-describeverifiedaccessinstanceloggingconfigurations.md)
  - [DescribeVerifiedAccessInstances](api-ec2-2016-11-15-describeverifiedaccessinstances.md)
  - [DescribeVerifiedAccessTrustProviders](api-ec2-2016-11-15-describeverifiedaccesstrustproviders.md)
  - [DescribeVolumeAttribute](api-ec2-2016-11-15-describevolumeattribute.md)
  - [DescribeVolumeStatus](api-ec2-2016-11-15-describevolumestatus.md)
  - [DescribeVolumes](api-ec2-2016-11-15-describevolumes.md)
  - [DescribeVolumesModifications](api-ec2-2016-11-15-describevolumesmodifications.md)
  - [DescribeVpcAttribute](api-ec2-2016-11-15-describevpcattribute.md)
  - [DescribeVpcBlockPublicAccessExclusions](api-ec2-2016-11-15-describevpcblockpublicaccessexclusions.md)
  - [DescribeVpcBlockPublicAccessOptions](api-ec2-2016-11-15-describevpcblockpublicaccessoptions.md)
  - [DescribeVpcClassicLink](api-ec2-2016-11-15-describevpcclassiclink.md)
  - [DescribeVpcClassicLinkDnsSupport](api-ec2-2016-11-15-describevpcclassiclinkdnssupport.md)
  - [DescribeVpcEncryptionControls](api-ec2-2016-11-15-describevpcencryptioncontrols.md)
  - [DescribeVpcEndpointAssociations](api-ec2-2016-11-15-describevpcendpointassociations.md)
  - [DescribeVpcEndpointConnectionNotifications](api-ec2-2016-11-15-describevpcendpointconnectionnotifications.md)
  - [DescribeVpcEndpointConnections](api-ec2-2016-11-15-describevpcendpointconnections.md)
  - [DescribeVpcEndpointServiceConfigurations](api-ec2-2016-11-15-describevpcendpointserviceconfigurations.md)
  - [DescribeVpcEndpointServicePermissions](api-ec2-2016-11-15-describevpcendpointservicepermissions.md)
  - [DescribeVpcEndpointServices](api-ec2-2016-11-15-describevpcendpointservices.md)
  - [DescribeVpcEndpoints](api-ec2-2016-11-15-describevpcendpoints.md)
  - [DescribeVpcPeeringConnections](api-ec2-2016-11-15-describevpcpeeringconnections.md)
  - [DescribeVpcs](api-ec2-2016-11-15-describevpcs.md)
  - [DescribeVpnConcentrators](api-ec2-2016-11-15-describevpnconcentrators.md)
  - [DescribeVpnConnections](api-ec2-2016-11-15-describevpnconnections.md)
  - [DescribeVpnGateways](api-ec2-2016-11-15-describevpngateways.md)
  - [DetachClassicLinkVpc](api-ec2-2016-11-15-detachclassiclinkvpc.md)
  - [DetachInternetGateway](api-ec2-2016-11-15-detachinternetgateway.md)
  - [DetachNetworkInterface](api-ec2-2016-11-15-detachnetworkinterface.md)
  - [DetachVerifiedAccessTrustProvider](api-ec2-2016-11-15-detachverifiedaccesstrustprovider.md)
  - [DetachVolume](api-ec2-2016-11-15-detachvolume.md)
  - [DetachVpnGateway](api-ec2-2016-11-15-detachvpngateway.md)
  - [DisableAddressTransfer](api-ec2-2016-11-15-disableaddresstransfer.md)
  - [DisableAllowedImagesSettings](api-ec2-2016-11-15-disableallowedimagessettings.md)
  - [DisableAwsNetworkPerformanceMetricSubscription](api-ec2-2016-11-15-disableawsnetworkperformancemetricsubscription.md)
  - [DisableCapacityManager](api-ec2-2016-11-15-disablecapacitymanager.md)
  - [DisableEbsEncryptionByDefault](api-ec2-2016-11-15-disableebsencryptionbydefault.md)
  - [DisableFastLaunch](api-ec2-2016-11-15-disablefastlaunch.md)
  - [DisableFastSnapshotRestores](api-ec2-2016-11-15-disablefastsnapshotrestores.md)
  - [DisableImage](api-ec2-2016-11-15-disableimage.md)
  - [DisableImageBlockPublicAccess](api-ec2-2016-11-15-disableimageblockpublicaccess.md)
  - [DisableImageDeprecation](api-ec2-2016-11-15-disableimagedeprecation.md)
  - [DisableImageDeregistrationProtection](api-ec2-2016-11-15-disableimagederegistrationprotection.md)
  - [DisableInstanceSqlHaStandbyDetections](api-ec2-2016-11-15-disableinstancesqlhastandbydetections.md)
  - [DisableIpamOrganizationAdminAccount](api-ec2-2016-11-15-disableipamorganizationadminaccount.md)
  - [DisableIpamPolicy](api-ec2-2016-11-15-disableipampolicy.md)
  - [DisableRouteServerPropagation](api-ec2-2016-11-15-disablerouteserverpropagation.md)
  - [DisableSerialConsoleAccess](api-ec2-2016-11-15-disableserialconsoleaccess.md)
  - [DisableSnapshotBlockPublicAccess](api-ec2-2016-11-15-disablesnapshotblockpublicaccess.md)
  - [DisableTransitGatewayRouteTablePropagation](api-ec2-2016-11-15-disabletransitgatewayroutetablepropagation.md)
  - [DisableVgwRoutePropagation](api-ec2-2016-11-15-disablevgwroutepropagation.md)
  - [DisableVpcClassicLink](api-ec2-2016-11-15-disablevpcclassiclink.md)
  - [DisableVpcClassicLinkDnsSupport](api-ec2-2016-11-15-disablevpcclassiclinkdnssupport.md)
  - [DisassociateAddress](api-ec2-2016-11-15-disassociateaddress.md)
  - [DisassociateCapacityReservationBillingOwner](api-ec2-2016-11-15-disassociatecapacityreservationbillingowner.md)
  - [DisassociateClientVpnTargetNetwork](api-ec2-2016-11-15-disassociateclientvpntargetnetwork.md)
  - [DisassociateEnclaveCertificateIamRole](api-ec2-2016-11-15-disassociateenclavecertificateiamrole.md)
  - [DisassociateIamInstanceProfile](api-ec2-2016-11-15-disassociateiaminstanceprofile.md)
  - [DisassociateInstanceEventWindow](api-ec2-2016-11-15-disassociateinstanceeventwindow.md)
  - [DisassociateIpamByoasn](api-ec2-2016-11-15-disassociateipambyoasn.md)
  - [DisassociateIpamResourceDiscovery](api-ec2-2016-11-15-disassociateipamresourcediscovery.md)
  - [DisassociateNatGatewayAddress](api-ec2-2016-11-15-disassociatenatgatewayaddress.md)
  - [DisassociateRouteServer](api-ec2-2016-11-15-disassociaterouteserver.md)
  - [DisassociateRouteTable](api-ec2-2016-11-15-disassociateroutetable.md)
  - [DisassociateSecurityGroupVpc](api-ec2-2016-11-15-disassociatesecuritygroupvpc.md)
  - [DisassociateSubnetCidrBlock](api-ec2-2016-11-15-disassociatesubnetcidrblock.md)
  - [DisassociateTransitGatewayMulticastDomain](api-ec2-2016-11-15-disassociatetransitgatewaymulticastdomain.md)
  - [DisassociateTransitGatewayPolicyTable](api-ec2-2016-11-15-disassociatetransitgatewaypolicytable.md)
  - [DisassociateTransitGatewayRouteTable](api-ec2-2016-11-15-disassociatetransitgatewayroutetable.md)
  - [DisassociateTrunkInterface](api-ec2-2016-11-15-disassociatetrunkinterface.md)
  - [DisassociateVpcCidrBlock](api-ec2-2016-11-15-disassociatevpccidrblock.md)
  - [EnableAddressTransfer](api-ec2-2016-11-15-enableaddresstransfer.md)
  - [EnableAllowedImagesSettings](api-ec2-2016-11-15-enableallowedimagessettings.md)
  - [EnableAwsNetworkPerformanceMetricSubscription](api-ec2-2016-11-15-enableawsnetworkperformancemetricsubscription.md)
  - [EnableCapacityManager](api-ec2-2016-11-15-enablecapacitymanager.md)
  - [EnableEbsEncryptionByDefault](api-ec2-2016-11-15-enableebsencryptionbydefault.md)
  - [EnableFastLaunch](api-ec2-2016-11-15-enablefastlaunch.md)
  - [EnableFastSnapshotRestores](api-ec2-2016-11-15-enablefastsnapshotrestores.md)
  - [EnableImage](api-ec2-2016-11-15-enableimage.md)
  - [EnableImageBlockPublicAccess](api-ec2-2016-11-15-enableimageblockpublicaccess.md)
  - [EnableImageDeprecation](api-ec2-2016-11-15-enableimagedeprecation.md)
  - [EnableImageDeregistrationProtection](api-ec2-2016-11-15-enableimagederegistrationprotection.md)
  - [EnableInstanceSqlHaStandbyDetections](api-ec2-2016-11-15-enableinstancesqlhastandbydetections.md)
  - [EnableIpamOrganizationAdminAccount](api-ec2-2016-11-15-enableipamorganizationadminaccount.md)
  - [EnableIpamPolicy](api-ec2-2016-11-15-enableipampolicy.md)
  - [EnableReachabilityAnalyzerOrganizationSharing](api-ec2-2016-11-15-enablereachabilityanalyzerorganizationsharing.md)
  - [EnableRouteServerPropagation](api-ec2-2016-11-15-enablerouteserverpropagation.md)
  - [EnableSerialConsoleAccess](api-ec2-2016-11-15-enableserialconsoleaccess.md)
  - [EnableSnapshotBlockPublicAccess](api-ec2-2016-11-15-enablesnapshotblockpublicaccess.md)
  - [EnableTransitGatewayRouteTablePropagation](api-ec2-2016-11-15-enabletransitgatewayroutetablepropagation.md)
  - [EnableVgwRoutePropagation](api-ec2-2016-11-15-enablevgwroutepropagation.md)
  - [EnableVolumeIO](api-ec2-2016-11-15-enablevolumeio.md)
  - [EnableVpcClassicLink](api-ec2-2016-11-15-enablevpcclassiclink.md)
  - [EnableVpcClassicLinkDnsSupport](api-ec2-2016-11-15-enablevpcclassiclinkdnssupport.md)
  - [ExportClientVpnClientCertificateRevocationList](api-ec2-2016-11-15-exportclientvpnclientcertificaterevocationlist.md)
  - [ExportClientVpnClientConfiguration](api-ec2-2016-11-15-exportclientvpnclientconfiguration.md)
  - [ExportImage](api-ec2-2016-11-15-exportimage.md)
  - [ExportTransitGatewayRoutes](api-ec2-2016-11-15-exporttransitgatewayroutes.md)
  - [ExportVerifiedAccessInstanceClientConfiguration](api-ec2-2016-11-15-exportverifiedaccessinstanceclientconfiguration.md)
  - [GetActiveVpnTunnelStatus](api-ec2-2016-11-15-getactivevpntunnelstatus.md)
  - [GetAllowedImagesSettings](api-ec2-2016-11-15-getallowedimagessettings.md)
  - [GetAssociatedEnclaveCertificateIamRoles](api-ec2-2016-11-15-getassociatedenclavecertificateiamroles.md)
  - [GetAssociatedIpv6PoolCidrs](api-ec2-2016-11-15-getassociatedipv6poolcidrs.md)
  - [GetAwsNetworkPerformanceData](api-ec2-2016-11-15-getawsnetworkperformancedata.md)
  - [GetCapacityManagerAttributes](api-ec2-2016-11-15-getcapacitymanagerattributes.md)
  - [GetCapacityManagerMetricData](api-ec2-2016-11-15-getcapacitymanagermetricdata.md)
  - [GetCapacityManagerMetricDimensions](api-ec2-2016-11-15-getcapacitymanagermetricdimensions.md)
  - [GetCapacityReservationUsage](api-ec2-2016-11-15-getcapacityreservationusage.md)
  - [GetCoipPoolUsage](api-ec2-2016-11-15-getcoippoolusage.md)
  - [GetConsoleOutput](api-ec2-2016-11-15-getconsoleoutput.md)
  - [GetConsoleScreenshot](api-ec2-2016-11-15-getconsolescreenshot.md)
  - [GetDeclarativePoliciesReportSummary](api-ec2-2016-11-15-getdeclarativepoliciesreportsummary.md)
  - [GetDefaultCreditSpecification](api-ec2-2016-11-15-getdefaultcreditspecification.md)
  - [GetEbsDefaultKmsKeyId](api-ec2-2016-11-15-getebsdefaultkmskeyid.md)
  - [GetEbsEncryptionByDefault](api-ec2-2016-11-15-getebsencryptionbydefault.md)
  - [GetEnabledIpamPolicy](api-ec2-2016-11-15-getenabledipampolicy.md)
  - [GetFlowLogsIntegrationTemplate](api-ec2-2016-11-15-getflowlogsintegrationtemplate.md)
  - [GetGroupsForCapacityReservation](api-ec2-2016-11-15-getgroupsforcapacityreservation.md)
  - [GetHostReservationPurchasePreview](api-ec2-2016-11-15-gethostreservationpurchasepreview.md)
  - [GetImageAncestry](api-ec2-2016-11-15-getimageancestry.md)
  - [GetImageBlockPublicAccessState](api-ec2-2016-11-15-getimageblockpublicaccessstate.md)
  - [GetInstanceMetadataDefaults](api-ec2-2016-11-15-getinstancemetadatadefaults.md)
  - [GetInstanceTpmEkPub](api-ec2-2016-11-15-getinstancetpmekpub.md)
  - [GetInstanceTypesFromInstanceRequirements](api-ec2-2016-11-15-getinstancetypesfrominstancerequirements.md)
  - [GetInstanceUefiData](api-ec2-2016-11-15-getinstanceuefidata.md)
  - [GetIpamAddressHistory](api-ec2-2016-11-15-getipamaddresshistory.md)
  - [GetIpamDiscoveredAccounts](api-ec2-2016-11-15-getipamdiscoveredaccounts.md)
  - [GetIpamDiscoveredPublicAddresses](api-ec2-2016-11-15-getipamdiscoveredpublicaddresses.md)
  - [GetIpamDiscoveredResourceCidrs](api-ec2-2016-11-15-getipamdiscoveredresourcecidrs.md)
  - [GetIpamPolicyAllocationRules](api-ec2-2016-11-15-getipampolicyallocationrules.md)
  - [GetIpamPolicyOrganizationTargets](api-ec2-2016-11-15-getipampolicyorganizationtargets.md)
  - [GetIpamPoolAllocations](api-ec2-2016-11-15-getipampoolallocations.md)
  - [GetIpamPoolCidrs](api-ec2-2016-11-15-getipampoolcidrs.md)
  - [GetIpamPrefixListResolverRules](api-ec2-2016-11-15-getipamprefixlistresolverrules.md)
  - [GetIpamPrefixListResolverVersionEntries](api-ec2-2016-11-15-getipamprefixlistresolverversionentries.md)
  - [GetIpamPrefixListResolverVersions](api-ec2-2016-11-15-getipamprefixlistresolverversions.md)
  - [GetIpamResourceCidrs](api-ec2-2016-11-15-getipamresourcecidrs.md)
  - [GetLaunchTemplateData](api-ec2-2016-11-15-getlaunchtemplatedata.md)
  - [GetManagedPrefixListAssociations](api-ec2-2016-11-15-getmanagedprefixlistassociations.md)
  - [GetManagedPrefixListEntries](api-ec2-2016-11-15-getmanagedprefixlistentries.md)
  - [GetNetworkInsightsAccessScopeAnalysisFindings](api-ec2-2016-11-15-getnetworkinsightsaccessscopeanalysisfindings.md)
  - [GetNetworkInsightsAccessScopeContent](api-ec2-2016-11-15-getnetworkinsightsaccessscopecontent.md)
  - [GetPasswordData](api-ec2-2016-11-15-getpassworddata.md)
  - [GetReservedInstancesExchangeQuote](api-ec2-2016-11-15-getreservedinstancesexchangequote.md)
  - [GetRouteServerAssociations](api-ec2-2016-11-15-getrouteserverassociations.md)
  - [GetRouteServerPropagations](api-ec2-2016-11-15-getrouteserverpropagations.md)
  - [GetRouteServerRoutingDatabase](api-ec2-2016-11-15-getrouteserverroutingdatabase.md)
  - [GetSecurityGroupsForVpc](api-ec2-2016-11-15-getsecuritygroupsforvpc.md)
  - [GetSerialConsoleAccessStatus](api-ec2-2016-11-15-getserialconsoleaccessstatus.md)
  - [GetSnapshotBlockPublicAccessState](api-ec2-2016-11-15-getsnapshotblockpublicaccessstate.md)
  - [GetSpotPlacementScores](api-ec2-2016-11-15-getspotplacementscores.md)
  - [GetSubnetCidrReservations](api-ec2-2016-11-15-getsubnetcidrreservations.md)
  - [GetTransitGatewayAttachmentPropagations](api-ec2-2016-11-15-gettransitgatewayattachmentpropagations.md)
  - [GetTransitGatewayMeteringPolicyEntries](api-ec2-2016-11-15-gettransitgatewaymeteringpolicyentries.md)
  - [GetTransitGatewayMulticastDomainAssociations](api-ec2-2016-11-15-gettransitgatewaymulticastdomainassociations.md)
  - [GetTransitGatewayPolicyTableAssociations](api-ec2-2016-11-15-gettransitgatewaypolicytableassociations.md)
  - [GetTransitGatewayPolicyTableEntries](api-ec2-2016-11-15-gettransitgatewaypolicytableentries.md)
  - [GetTransitGatewayPrefixListReferences](api-ec2-2016-11-15-gettransitgatewayprefixlistreferences.md)
  - [GetTransitGatewayRouteTableAssociations](api-ec2-2016-11-15-gettransitgatewayroutetableassociations.md)
  - [GetTransitGatewayRouteTablePropagations](api-ec2-2016-11-15-gettransitgatewayroutetablepropagations.md)
  - [GetVerifiedAccessEndpointPolicy](api-ec2-2016-11-15-getverifiedaccessendpointpolicy.md)
  - [GetVerifiedAccessEndpointTargets](api-ec2-2016-11-15-getverifiedaccessendpointtargets.md)
  - [GetVerifiedAccessGroupPolicy](api-ec2-2016-11-15-getverifiedaccessgrouppolicy.md)
  - [GetVpcResourcesBlockingEncryptionEnforcement](api-ec2-2016-11-15-getvpcresourcesblockingencryptionenforcement.md)
  - [GetVpnConnectionDeviceSampleConfiguration](api-ec2-2016-11-15-getvpnconnectiondevicesampleconfiguration.md)
  - [GetVpnConnectionDeviceTypes](api-ec2-2016-11-15-getvpnconnectiondevicetypes.md)
  - [GetVpnTunnelReplacementStatus](api-ec2-2016-11-15-getvpntunnelreplacementstatus.md)
  - [ImportClientVpnClientCertificateRevocationList](api-ec2-2016-11-15-importclientvpnclientcertificaterevocationlist.md)
  - [ImportImage](api-ec2-2016-11-15-importimage.md)
  - [ImportInstance](api-ec2-2016-11-15-importinstance.md)
  - [ImportKeyPair](api-ec2-2016-11-15-importkeypair.md)
  - [ImportSnapshot](api-ec2-2016-11-15-importsnapshot.md)
  - [ImportVolume](api-ec2-2016-11-15-importvolume.md)
  - [ListImagesInRecycleBin](api-ec2-2016-11-15-listimagesinrecyclebin.md)
  - [ListSnapshotsInRecycleBin](api-ec2-2016-11-15-listsnapshotsinrecyclebin.md)
  - [ListVolumesInRecycleBin](api-ec2-2016-11-15-listvolumesinrecyclebin.md)
  - [LockSnapshot](api-ec2-2016-11-15-locksnapshot.md)
  - [ModifyAddressAttribute](api-ec2-2016-11-15-modifyaddressattribute.md)
  - [ModifyAvailabilityZoneGroup](api-ec2-2016-11-15-modifyavailabilityzonegroup.md)
  - [ModifyCapacityReservation](api-ec2-2016-11-15-modifycapacityreservation.md)
  - [ModifyCapacityReservationFleet](api-ec2-2016-11-15-modifycapacityreservationfleet.md)
  - [ModifyClientVpnEndpoint](api-ec2-2016-11-15-modifyclientvpnendpoint.md)
  - [ModifyDefaultCreditSpecification](api-ec2-2016-11-15-modifydefaultcreditspecification.md)
  - [ModifyEbsDefaultKmsKeyId](api-ec2-2016-11-15-modifyebsdefaultkmskeyid.md)
  - [ModifyFleet](api-ec2-2016-11-15-modifyfleet.md)
  - [ModifyFpgaImageAttribute](api-ec2-2016-11-15-modifyfpgaimageattribute.md)
  - [ModifyHosts](api-ec2-2016-11-15-modifyhosts.md)
  - [ModifyIdFormat](api-ec2-2016-11-15-modifyidformat.md)
  - [ModifyIdentityIdFormat](api-ec2-2016-11-15-modifyidentityidformat.md)
  - [ModifyImageAttribute](api-ec2-2016-11-15-modifyimageattribute.md)
  - [ModifyInstanceAttribute](api-ec2-2016-11-15-modifyinstanceattribute.md)
  - [ModifyInstanceCapacityReservationAttributes](api-ec2-2016-11-15-modifyinstancecapacityreservationattributes.md)
  - [ModifyInstanceConnectEndpoint](api-ec2-2016-11-15-modifyinstanceconnectendpoint.md)
  - [ModifyInstanceCpuOptions](api-ec2-2016-11-15-modifyinstancecpuoptions.md)
  - [ModifyInstanceCreditSpecification](api-ec2-2016-11-15-modifyinstancecreditspecification.md)
  - [ModifyInstanceEventStartTime](api-ec2-2016-11-15-modifyinstanceeventstarttime.md)
  - [ModifyInstanceEventWindow](api-ec2-2016-11-15-modifyinstanceeventwindow.md)
  - [ModifyInstanceMaintenanceOptions](api-ec2-2016-11-15-modifyinstancemaintenanceoptions.md)
  - [ModifyInstanceMetadataDefaults](api-ec2-2016-11-15-modifyinstancemetadatadefaults.md)
  - [ModifyInstanceMetadataOptions](api-ec2-2016-11-15-modifyinstancemetadataoptions.md)
  - [ModifyInstanceNetworkPerformanceOptions](api-ec2-2016-11-15-modifyinstancenetworkperformanceoptions.md)
  - [ModifyInstancePlacement](api-ec2-2016-11-15-modifyinstanceplacement.md)
  - [ModifyIpam](api-ec2-2016-11-15-modifyipam.md)
  - [ModifyIpamPolicyAllocationRules](api-ec2-2016-11-15-modifyipampolicyallocationrules.md)
  - [ModifyIpamPool](api-ec2-2016-11-15-modifyipampool.md)
  - [ModifyIpamPrefixListResolver](api-ec2-2016-11-15-modifyipamprefixlistresolver.md)
  - [ModifyIpamPrefixListResolverTarget](api-ec2-2016-11-15-modifyipamprefixlistresolvertarget.md)
  - [ModifyIpamResourceCidr](api-ec2-2016-11-15-modifyipamresourcecidr.md)
  - [ModifyIpamResourceDiscovery](api-ec2-2016-11-15-modifyipamresourcediscovery.md)
  - [ModifyIpamScope](api-ec2-2016-11-15-modifyipamscope.md)
  - [ModifyLaunchTemplate](api-ec2-2016-11-15-modifylaunchtemplate.md)
  - [ModifyLocalGatewayRoute](api-ec2-2016-11-15-modifylocalgatewayroute.md)
  - [ModifyManagedPrefixList](api-ec2-2016-11-15-modifymanagedprefixlist.md)
  - [ModifyNetworkInterfaceAttribute](api-ec2-2016-11-15-modifynetworkinterfaceattribute.md)
  - [ModifyPrivateDnsNameOptions](api-ec2-2016-11-15-modifyprivatednsnameoptions.md)
  - [ModifyPublicIpDnsNameOptions](api-ec2-2016-11-15-modifypublicipdnsnameoptions.md)
  - [ModifyReservedInstances](api-ec2-2016-11-15-modifyreservedinstances.md)
  - [ModifyRouteServer](api-ec2-2016-11-15-modifyrouteserver.md)
  - [ModifySecurityGroupRules](api-ec2-2016-11-15-modifysecuritygrouprules.md)
  - [ModifySnapshotAttribute](api-ec2-2016-11-15-modifysnapshotattribute.md)
  - [ModifySnapshotTier](api-ec2-2016-11-15-modifysnapshottier.md)
  - [ModifySpotFleetRequest](api-ec2-2016-11-15-modifyspotfleetrequest.md)
  - [ModifySubnetAttribute](api-ec2-2016-11-15-modifysubnetattribute.md)
  - [ModifyTrafficMirrorFilterNetworkServices](api-ec2-2016-11-15-modifytrafficmirrorfilternetworkservices.md)
  - [ModifyTrafficMirrorFilterRule](api-ec2-2016-11-15-modifytrafficmirrorfilterrule.md)
  - [ModifyTrafficMirrorSession](api-ec2-2016-11-15-modifytrafficmirrorsession.md)
  - [ModifyTransitGateway](api-ec2-2016-11-15-modifytransitgateway.md)
  - [ModifyTransitGatewayMeteringPolicy](api-ec2-2016-11-15-modifytransitgatewaymeteringpolicy.md)
  - [ModifyTransitGatewayPrefixListReference](api-ec2-2016-11-15-modifytransitgatewayprefixlistreference.md)
  - [ModifyTransitGatewayVpcAttachment](api-ec2-2016-11-15-modifytransitgatewayvpcattachment.md)
  - [ModifyVerifiedAccessEndpoint](api-ec2-2016-11-15-modifyverifiedaccessendpoint.md)
  - [ModifyVerifiedAccessEndpointPolicy](api-ec2-2016-11-15-modifyverifiedaccessendpointpolicy.md)
  - [ModifyVerifiedAccessGroup](api-ec2-2016-11-15-modifyverifiedaccessgroup.md)
  - [ModifyVerifiedAccessGroupPolicy](api-ec2-2016-11-15-modifyverifiedaccessgrouppolicy.md)
  - [ModifyVerifiedAccessInstance](api-ec2-2016-11-15-modifyverifiedaccessinstance.md)
  - [ModifyVerifiedAccessInstanceLoggingConfiguration](api-ec2-2016-11-15-modifyverifiedaccessinstanceloggingconfiguration.md)
  - [ModifyVerifiedAccessTrustProvider](api-ec2-2016-11-15-modifyverifiedaccesstrustprovider.md)
  - [ModifyVolume](api-ec2-2016-11-15-modifyvolume.md)
  - [ModifyVolumeAttribute](api-ec2-2016-11-15-modifyvolumeattribute.md)
  - [ModifyVpcAttribute](api-ec2-2016-11-15-modifyvpcattribute.md)
  - [ModifyVpcBlockPublicAccessExclusion](api-ec2-2016-11-15-modifyvpcblockpublicaccessexclusion.md)
  - [ModifyVpcBlockPublicAccessOptions](api-ec2-2016-11-15-modifyvpcblockpublicaccessoptions.md)
  - [ModifyVpcEncryptionControl](api-ec2-2016-11-15-modifyvpcencryptioncontrol.md)
  - [ModifyVpcEndpoint](api-ec2-2016-11-15-modifyvpcendpoint.md)
  - [ModifyVpcEndpointConnectionNotification](api-ec2-2016-11-15-modifyvpcendpointconnectionnotification.md)
  - [ModifyVpcEndpointServiceConfiguration](api-ec2-2016-11-15-modifyvpcendpointserviceconfiguration.md)
  - [ModifyVpcEndpointServicePayerResponsibility](api-ec2-2016-11-15-modifyvpcendpointservicepayerresponsibility.md)
  - [ModifyVpcEndpointServicePermissions](api-ec2-2016-11-15-modifyvpcendpointservicepermissions.md)
  - [ModifyVpcPeeringConnectionOptions](api-ec2-2016-11-15-modifyvpcpeeringconnectionoptions.md)
  - [ModifyVpcTenancy](api-ec2-2016-11-15-modifyvpctenancy.md)
  - [ModifyVpnConnection](api-ec2-2016-11-15-modifyvpnconnection.md)
  - [ModifyVpnConnectionOptions](api-ec2-2016-11-15-modifyvpnconnectionoptions.md)
  - [ModifyVpnTunnelCertificate](api-ec2-2016-11-15-modifyvpntunnelcertificate.md)
  - [ModifyVpnTunnelOptions](api-ec2-2016-11-15-modifyvpntunneloptions.md)
  - [MonitorInstances](api-ec2-2016-11-15-monitorinstances.md)
  - [MoveAddressToVpc](api-ec2-2016-11-15-moveaddresstovpc.md)
  - [MoveByoipCidrToIpam](api-ec2-2016-11-15-movebyoipcidrtoipam.md)
  - [MoveCapacityReservationInstances](api-ec2-2016-11-15-movecapacityreservationinstances.md)
  - [ProvisionByoipCidr](api-ec2-2016-11-15-provisionbyoipcidr.md)
  - [ProvisionIpamByoasn](api-ec2-2016-11-15-provisionipambyoasn.md)
  - [ProvisionIpamPoolCidr](api-ec2-2016-11-15-provisionipampoolcidr.md)
  - [ProvisionPublicIpv4PoolCidr](api-ec2-2016-11-15-provisionpublicipv4poolcidr.md)
  - [PurchaseCapacityBlock](api-ec2-2016-11-15-purchasecapacityblock.md)
  - [PurchaseCapacityBlockExtension](api-ec2-2016-11-15-purchasecapacityblockextension.md)
  - [PurchaseHostReservation](api-ec2-2016-11-15-purchasehostreservation.md)
  - [PurchaseReservedInstancesOffering](api-ec2-2016-11-15-purchasereservedinstancesoffering.md)
  - [PurchaseScheduledInstances](api-ec2-2016-11-15-purchasescheduledinstances.md)
  - [RebootInstances](api-ec2-2016-11-15-rebootinstances.md)
  - [RegisterImage](api-ec2-2016-11-15-registerimage.md)
  - [RegisterInstanceEventNotificationAttributes](api-ec2-2016-11-15-registerinstanceeventnotificationattributes.md)
  - [RegisterTransitGatewayMulticastGroupMembers](api-ec2-2016-11-15-registertransitgatewaymulticastgroupmembers.md)
  - [RegisterTransitGatewayMulticastGroupSources](api-ec2-2016-11-15-registertransitgatewaymulticastgroupsources.md)
  - [RejectCapacityReservationBillingOwnership](api-ec2-2016-11-15-rejectcapacityreservationbillingownership.md)
  - [RejectTransitGatewayMulticastDomainAssociations](api-ec2-2016-11-15-rejecttransitgatewaymulticastdomainassociations.md)
  - [RejectTransitGatewayPeeringAttachment](api-ec2-2016-11-15-rejecttransitgatewaypeeringattachment.md)
  - [RejectTransitGatewayVpcAttachment](api-ec2-2016-11-15-rejecttransitgatewayvpcattachment.md)
  - [RejectVpcEndpointConnections](api-ec2-2016-11-15-rejectvpcendpointconnections.md)
  - [RejectVpcPeeringConnection](api-ec2-2016-11-15-rejectvpcpeeringconnection.md)
  - [ReleaseAddress](api-ec2-2016-11-15-releaseaddress.md)
  - [ReleaseHosts](api-ec2-2016-11-15-releasehosts.md)
  - [ReleaseIpamPoolAllocation](api-ec2-2016-11-15-releaseipampoolallocation.md)
  - [ReplaceIamInstanceProfileAssociation](api-ec2-2016-11-15-replaceiaminstanceprofileassociation.md)
  - [ReplaceImageCriteriaInAllowedImagesSettings](api-ec2-2016-11-15-replaceimagecriteriainallowedimagessettings.md)
  - [ReplaceNetworkAclAssociation](api-ec2-2016-11-15-replacenetworkaclassociation.md)
  - [ReplaceNetworkAclEntry](api-ec2-2016-11-15-replacenetworkaclentry.md)
  - [ReplaceRoute](api-ec2-2016-11-15-replaceroute.md)
  - [ReplaceRouteTableAssociation](api-ec2-2016-11-15-replaceroutetableassociation.md)
  - [ReplaceTransitGatewayRoute](api-ec2-2016-11-15-replacetransitgatewayroute.md)
  - [ReplaceVpnTunnel](api-ec2-2016-11-15-replacevpntunnel.md)
  - [ReportInstanceStatus](api-ec2-2016-11-15-reportinstancestatus.md)
  - [RequestSpotFleet](api-ec2-2016-11-15-requestspotfleet.md)
  - [RequestSpotInstances](api-ec2-2016-11-15-requestspotinstances.md)
  - [ResetAddressAttribute](api-ec2-2016-11-15-resetaddressattribute.md)
  - [ResetEbsDefaultKmsKeyId](api-ec2-2016-11-15-resetebsdefaultkmskeyid.md)
  - [ResetFpgaImageAttribute](api-ec2-2016-11-15-resetfpgaimageattribute.md)
  - [ResetImageAttribute](api-ec2-2016-11-15-resetimageattribute.md)
  - [ResetInstanceAttribute](api-ec2-2016-11-15-resetinstanceattribute.md)
  - [ResetNetworkInterfaceAttribute](api-ec2-2016-11-15-resetnetworkinterfaceattribute.md)
  - [ResetSnapshotAttribute](api-ec2-2016-11-15-resetsnapshotattribute.md)
  - [RestoreAddressToClassic](api-ec2-2016-11-15-restoreaddresstoclassic.md)
  - [RestoreImageFromRecycleBin](api-ec2-2016-11-15-restoreimagefromrecyclebin.md)
  - [RestoreManagedPrefixListVersion](api-ec2-2016-11-15-restoremanagedprefixlistversion.md)
  - [RestoreSnapshotFromRecycleBin](api-ec2-2016-11-15-restoresnapshotfromrecyclebin.md)
  - [RestoreSnapshotTier](api-ec2-2016-11-15-restoresnapshottier.md)
  - [RestoreVolumeFromRecycleBin](api-ec2-2016-11-15-restorevolumefromrecyclebin.md)
  - [RevokeClientVpnIngress](api-ec2-2016-11-15-revokeclientvpningress.md)
  - [RevokeSecurityGroupEgress](api-ec2-2016-11-15-revokesecuritygroupegress.md)
  - [RevokeSecurityGroupIngress](api-ec2-2016-11-15-revokesecuritygroupingress.md)
  - [RunInstances](api-ec2-2016-11-15-runinstances.md)
  - [RunScheduledInstances](api-ec2-2016-11-15-runscheduledinstances.md)
  - [SearchLocalGatewayRoutes](api-ec2-2016-11-15-searchlocalgatewayroutes.md)
  - [SearchTransitGatewayMulticastGroups](api-ec2-2016-11-15-searchtransitgatewaymulticastgroups.md)
  - [SearchTransitGatewayRoutes](api-ec2-2016-11-15-searchtransitgatewayroutes.md)
  - [SendDiagnosticInterrupt](api-ec2-2016-11-15-senddiagnosticinterrupt.md)
  - [StartDeclarativePoliciesReport](api-ec2-2016-11-15-startdeclarativepoliciesreport.md)
  - [StartInstances](api-ec2-2016-11-15-startinstances.md)
  - [StartNetworkInsightsAccessScopeAnalysis](api-ec2-2016-11-15-startnetworkinsightsaccessscopeanalysis.md)
  - [StartNetworkInsightsAnalysis](api-ec2-2016-11-15-startnetworkinsightsanalysis.md)
  - [StartVpcEndpointServicePrivateDnsVerification](api-ec2-2016-11-15-startvpcendpointserviceprivatednsverification.md)
  - [StopInstances](api-ec2-2016-11-15-stopinstances.md)
  - [TerminateClientVpnConnections](api-ec2-2016-11-15-terminateclientvpnconnections.md)
  - [TerminateInstances](api-ec2-2016-11-15-terminateinstances.md)
  - [UnassignIpv6Addresses](api-ec2-2016-11-15-unassignipv6addresses.md)
  - [UnassignPrivateIpAddresses](api-ec2-2016-11-15-unassignprivateipaddresses.md)
  - [UnassignPrivateNatGatewayAddress](api-ec2-2016-11-15-unassignprivatenatgatewayaddress.md)
  - [UnlockSnapshot](api-ec2-2016-11-15-unlocksnapshot.md)
  - [UnmonitorInstances](api-ec2-2016-11-15-unmonitorinstances.md)
  - [UpdateCapacityManagerOrganizationsAccess](api-ec2-2016-11-15-updatecapacitymanagerorganizationsaccess.md)
  - [UpdateInterruptibleCapacityReservationAllocation](api-ec2-2016-11-15-updateinterruptiblecapacityreservationallocation.md)
  - [UpdateSecurityGroupRuleDescriptionsEgress](api-ec2-2016-11-15-updatesecuritygroupruledescriptionsegress.md)
  - [UpdateSecurityGroupRuleDescriptionsIngress](api-ec2-2016-11-15-updatesecuritygroupruledescriptionsingress.md)
  - [WithdrawByoipCidr](api-ec2-2016-11-15-withdrawbyoipcidr.md)
- [**2016-09-15**](api-ec2-2016-09-15.md)

  - [AcceptReservedInstancesExchangeQuote](api-ec2-2016-09-15-acceptreservedinstancesexchangequote.md)
  - [AcceptVpcPeeringConnection](api-ec2-2016-09-15-acceptvpcpeeringconnection.md)
  - [AllocateAddress](api-ec2-2016-09-15-allocateaddress.md)
  - [AllocateHosts](api-ec2-2016-09-15-allocatehosts.md)
  - [AssignPrivateIpAddresses](api-ec2-2016-09-15-assignprivateipaddresses.md)
  - [AssociateAddress](api-ec2-2016-09-15-associateaddress.md)
  - [AssociateDhcpOptions](api-ec2-2016-09-15-associatedhcpoptions.md)
  - [AssociateRouteTable](api-ec2-2016-09-15-associateroutetable.md)
  - [AttachClassicLinkVpc](api-ec2-2016-09-15-attachclassiclinkvpc.md)
  - [AttachInternetGateway](api-ec2-2016-09-15-attachinternetgateway.md)
  - [AttachNetworkInterface](api-ec2-2016-09-15-attachnetworkinterface.md)
  - [AttachVolume](api-ec2-2016-09-15-attachvolume.md)
  - [AttachVpnGateway](api-ec2-2016-09-15-attachvpngateway.md)
  - [AuthorizeSecurityGroupEgress](api-ec2-2016-09-15-authorizesecuritygroupegress.md)
  - [AuthorizeSecurityGroupIngress](api-ec2-2016-09-15-authorizesecuritygroupingress.md)
  - [BundleInstance](api-ec2-2016-09-15-bundleinstance.md)
  - [CancelBundleTask](api-ec2-2016-09-15-cancelbundletask.md)
  - [CancelConversionTask](api-ec2-2016-09-15-cancelconversiontask.md)
  - [CancelExportTask](api-ec2-2016-09-15-cancelexporttask.md)
  - [CancelImportTask](api-ec2-2016-09-15-cancelimporttask.md)
  - [CancelReservedInstancesListing](api-ec2-2016-09-15-cancelreservedinstanceslisting.md)
  - [CancelSpotFleetRequests](api-ec2-2016-09-15-cancelspotfleetrequests.md)
  - [CancelSpotInstanceRequests](api-ec2-2016-09-15-cancelspotinstancerequests.md)
  - [ConfirmProductInstance](api-ec2-2016-09-15-confirmproductinstance.md)
  - [CopyImage](api-ec2-2016-09-15-copyimage.md)
  - [CopySnapshot](api-ec2-2016-09-15-copysnapshot.md)
  - [CreateCustomerGateway](api-ec2-2016-09-15-createcustomergateway.md)
  - [CreateDhcpOptions](api-ec2-2016-09-15-createdhcpoptions.md)
  - [CreateFlowLogs](api-ec2-2016-09-15-createflowlogs.md)
  - [CreateImage](api-ec2-2016-09-15-createimage.md)
  - [CreateInstanceExportTask](api-ec2-2016-09-15-createinstanceexporttask.md)
  - [CreateInternetGateway](api-ec2-2016-09-15-createinternetgateway.md)
  - [CreateKeyPair](api-ec2-2016-09-15-createkeypair.md)
  - [CreateNatGateway](api-ec2-2016-09-15-createnatgateway.md)
  - [CreateNetworkAcl](api-ec2-2016-09-15-createnetworkacl.md)
  - [CreateNetworkAclEntry](api-ec2-2016-09-15-createnetworkaclentry.md)
  - [CreateNetworkInterface](api-ec2-2016-09-15-createnetworkinterface.md)
  - [CreatePlacementGroup](api-ec2-2016-09-15-createplacementgroup.md)
  - [CreateReservedInstancesListing](api-ec2-2016-09-15-createreservedinstanceslisting.md)
  - [CreateRoute](api-ec2-2016-09-15-createroute.md)
  - [CreateRouteTable](api-ec2-2016-09-15-createroutetable.md)
  - [CreateSecurityGroup](api-ec2-2016-09-15-createsecuritygroup.md)
  - [CreateSnapshot](api-ec2-2016-09-15-createsnapshot.md)
  - [CreateSpotDatafeedSubscription](api-ec2-2016-09-15-createspotdatafeedsubscription.md)
  - [CreateSubnet](api-ec2-2016-09-15-createsubnet.md)
  - [CreateTags](api-ec2-2016-09-15-createtags.md)
  - [CreateVolume](api-ec2-2016-09-15-createvolume.md)
  - [CreateVpc](api-ec2-2016-09-15-createvpc.md)
  - [CreateVpcEndpoint](api-ec2-2016-09-15-createvpcendpoint.md)
  - [CreateVpcPeeringConnection](api-ec2-2016-09-15-createvpcpeeringconnection.md)
  - [CreateVpnConnection](api-ec2-2016-09-15-createvpnconnection.md)
  - [CreateVpnConnectionRoute](api-ec2-2016-09-15-createvpnconnectionroute.md)
  - [CreateVpnGateway](api-ec2-2016-09-15-createvpngateway.md)
  - [DeleteCustomerGateway](api-ec2-2016-09-15-deletecustomergateway.md)
  - [DeleteDhcpOptions](api-ec2-2016-09-15-deletedhcpoptions.md)
  - [DeleteFlowLogs](api-ec2-2016-09-15-deleteflowlogs.md)
  - [DeleteInternetGateway](api-ec2-2016-09-15-deleteinternetgateway.md)
  - [DeleteKeyPair](api-ec2-2016-09-15-deletekeypair.md)
  - [DeleteNatGateway](api-ec2-2016-09-15-deletenatgateway.md)
  - [DeleteNetworkAcl](api-ec2-2016-09-15-deletenetworkacl.md)
  - [DeleteNetworkAclEntry](api-ec2-2016-09-15-deletenetworkaclentry.md)
  - [DeleteNetworkInterface](api-ec2-2016-09-15-deletenetworkinterface.md)
  - [DeletePlacementGroup](api-ec2-2016-09-15-deleteplacementgroup.md)
  - [DeleteRoute](api-ec2-2016-09-15-deleteroute.md)
  - [DeleteRouteTable](api-ec2-2016-09-15-deleteroutetable.md)
  - [DeleteSecurityGroup](api-ec2-2016-09-15-deletesecuritygroup.md)
  - [DeleteSnapshot](api-ec2-2016-09-15-deletesnapshot.md)
  - [DeleteSpotDatafeedSubscription](api-ec2-2016-09-15-deletespotdatafeedsubscription.md)
  - [DeleteSubnet](api-ec2-2016-09-15-deletesubnet.md)
  - [DeleteTags](api-ec2-2016-09-15-deletetags.md)
  - [DeleteVolume](api-ec2-2016-09-15-deletevolume.md)
  - [DeleteVpc](api-ec2-2016-09-15-deletevpc.md)
  - [DeleteVpcEndpoints](api-ec2-2016-09-15-deletevpcendpoints.md)
  - [DeleteVpcPeeringConnection](api-ec2-2016-09-15-deletevpcpeeringconnection.md)
  - [DeleteVpnConnection](api-ec2-2016-09-15-deletevpnconnection.md)
  - [DeleteVpnConnectionRoute](api-ec2-2016-09-15-deletevpnconnectionroute.md)
  - [DeleteVpnGateway](api-ec2-2016-09-15-deletevpngateway.md)
  - [DeregisterImage](api-ec2-2016-09-15-deregisterimage.md)
  - [DescribeAccountAttributes](api-ec2-2016-09-15-describeaccountattributes.md)
  - [DescribeAddresses](api-ec2-2016-09-15-describeaddresses.md)
  - [DescribeAvailabilityZones](api-ec2-2016-09-15-describeavailabilityzones.md)
  - [DescribeBundleTasks](api-ec2-2016-09-15-describebundletasks.md)
  - [DescribeClassicLinkInstances](api-ec2-2016-09-15-describeclassiclinkinstances.md)
  - [DescribeConversionTasks](api-ec2-2016-09-15-describeconversiontasks.md)
  - [DescribeCustomerGateways](api-ec2-2016-09-15-describecustomergateways.md)
  - [DescribeDhcpOptions](api-ec2-2016-09-15-describedhcpoptions.md)
  - [DescribeExportTasks](api-ec2-2016-09-15-describeexporttasks.md)
  - [DescribeFlowLogs](api-ec2-2016-09-15-describeflowlogs.md)
  - [DescribeHostReservationOfferings](api-ec2-2016-09-15-describehostreservationofferings.md)
  - [DescribeHostReservations](api-ec2-2016-09-15-describehostreservations.md)
  - [DescribeHosts](api-ec2-2016-09-15-describehosts.md)
  - [DescribeIdFormat](api-ec2-2016-09-15-describeidformat.md)
  - [DescribeIdentityIdFormat](api-ec2-2016-09-15-describeidentityidformat.md)
  - [DescribeImageAttribute](api-ec2-2016-09-15-describeimageattribute.md)
  - [DescribeImages](api-ec2-2016-09-15-describeimages.md)
  - [DescribeImportImageTasks](api-ec2-2016-09-15-describeimportimagetasks.md)
  - [DescribeImportSnapshotTasks](api-ec2-2016-09-15-describeimportsnapshottasks.md)
  - [DescribeInstanceAttribute](api-ec2-2016-09-15-describeinstanceattribute.md)
  - [DescribeInstanceStatus](api-ec2-2016-09-15-describeinstancestatus.md)
  - [DescribeInstances](api-ec2-2016-09-15-describeinstances.md)
  - [DescribeInternetGateways](api-ec2-2016-09-15-describeinternetgateways.md)
  - [DescribeKeyPairs](api-ec2-2016-09-15-describekeypairs.md)
  - [DescribeMovingAddresses](api-ec2-2016-09-15-describemovingaddresses.md)
  - [DescribeNatGateways](api-ec2-2016-09-15-describenatgateways.md)
  - [DescribeNetworkAcls](api-ec2-2016-09-15-describenetworkacls.md)
  - [DescribeNetworkInterfaceAttribute](api-ec2-2016-09-15-describenetworkinterfaceattribute.md)
  - [DescribeNetworkInterfaces](api-ec2-2016-09-15-describenetworkinterfaces.md)
  - [DescribePlacementGroups](api-ec2-2016-09-15-describeplacementgroups.md)
  - [DescribePrefixLists](api-ec2-2016-09-15-describeprefixlists.md)
  - [DescribeRegions](api-ec2-2016-09-15-describeregions.md)
  - [DescribeReservedInstances](api-ec2-2016-09-15-describereservedinstances.md)
  - [DescribeReservedInstancesListings](api-ec2-2016-09-15-describereservedinstanceslistings.md)
  - [DescribeReservedInstancesModifications](api-ec2-2016-09-15-describereservedinstancesmodifications.md)
  - [DescribeReservedInstancesOfferings](api-ec2-2016-09-15-describereservedinstancesofferings.md)
  - [DescribeRouteTables](api-ec2-2016-09-15-describeroutetables.md)
  - [DescribeScheduledInstanceAvailability](api-ec2-2016-09-15-describescheduledinstanceavailability.md)
  - [DescribeScheduledInstances](api-ec2-2016-09-15-describescheduledinstances.md)
  - [DescribeSecurityGroupReferences](api-ec2-2016-09-15-describesecuritygroupreferences.md)
  - [DescribeSecurityGroups](api-ec2-2016-09-15-describesecuritygroups.md)
  - [DescribeSnapshotAttribute](api-ec2-2016-09-15-describesnapshotattribute.md)
  - [DescribeSnapshots](api-ec2-2016-09-15-describesnapshots.md)
  - [DescribeSpotDatafeedSubscription](api-ec2-2016-09-15-describespotdatafeedsubscription.md)
  - [DescribeSpotFleetInstances](api-ec2-2016-09-15-describespotfleetinstances.md)
  - [DescribeSpotFleetRequestHistory](api-ec2-2016-09-15-describespotfleetrequesthistory.md)
  - [DescribeSpotFleetRequests](api-ec2-2016-09-15-describespotfleetrequests.md)
  - [DescribeSpotInstanceRequests](api-ec2-2016-09-15-describespotinstancerequests.md)
  - [DescribeSpotPriceHistory](api-ec2-2016-09-15-describespotpricehistory.md)
  - [DescribeStaleSecurityGroups](api-ec2-2016-09-15-describestalesecuritygroups.md)
  - [DescribeSubnets](api-ec2-2016-09-15-describesubnets.md)
  - [DescribeTags](api-ec2-2016-09-15-describetags.md)
  - [DescribeVolumeAttribute](api-ec2-2016-09-15-describevolumeattribute.md)
  - [DescribeVolumeStatus](api-ec2-2016-09-15-describevolumestatus.md)
  - [DescribeVolumes](api-ec2-2016-09-15-describevolumes.md)
  - [DescribeVpcAttribute](api-ec2-2016-09-15-describevpcattribute.md)
  - [DescribeVpcClassicLink](api-ec2-2016-09-15-describevpcclassiclink.md)
  - [DescribeVpcClassicLinkDnsSupport](api-ec2-2016-09-15-describevpcclassiclinkdnssupport.md)
  - [DescribeVpcEndpointServices](api-ec2-2016-09-15-describevpcendpointservices.md)
  - [DescribeVpcEndpoints](api-ec2-2016-09-15-describevpcendpoints.md)
  - [DescribeVpcPeeringConnections](api-ec2-2016-09-15-describevpcpeeringconnections.md)
  - [DescribeVpcs](api-ec2-2016-09-15-describevpcs.md)
  - [DescribeVpnConnections](api-ec2-2016-09-15-describevpnconnections.md)
  - [DescribeVpnGateways](api-ec2-2016-09-15-describevpngateways.md)
  - [DetachClassicLinkVpc](api-ec2-2016-09-15-detachclassiclinkvpc.md)
  - [DetachInternetGateway](api-ec2-2016-09-15-detachinternetgateway.md)
  - [DetachNetworkInterface](api-ec2-2016-09-15-detachnetworkinterface.md)
  - [DetachVolume](api-ec2-2016-09-15-detachvolume.md)
  - [DetachVpnGateway](api-ec2-2016-09-15-detachvpngateway.md)
  - [DisableVgwRoutePropagation](api-ec2-2016-09-15-disablevgwroutepropagation.md)
  - [DisableVpcClassicLink](api-ec2-2016-09-15-disablevpcclassiclink.md)
  - [DisableVpcClassicLinkDnsSupport](api-ec2-2016-09-15-disablevpcclassiclinkdnssupport.md)
  - [DisassociateAddress](api-ec2-2016-09-15-disassociateaddress.md)
  - [DisassociateRouteTable](api-ec2-2016-09-15-disassociateroutetable.md)
  - [EnableVgwRoutePropagation](api-ec2-2016-09-15-enablevgwroutepropagation.md)
  - [EnableVolumeIO](api-ec2-2016-09-15-enablevolumeio.md)
  - [EnableVpcClassicLink](api-ec2-2016-09-15-enablevpcclassiclink.md)
  - [EnableVpcClassicLinkDnsSupport](api-ec2-2016-09-15-enablevpcclassiclinkdnssupport.md)
  - [GetConsoleOutput](api-ec2-2016-09-15-getconsoleoutput.md)
  - [GetConsoleScreenshot](api-ec2-2016-09-15-getconsolescreenshot.md)
  - [GetHostReservationPurchasePreview](api-ec2-2016-09-15-gethostreservationpurchasepreview.md)
  - [GetPasswordData](api-ec2-2016-09-15-getpassworddata.md)
  - [GetReservedInstancesExchangeQuote](api-ec2-2016-09-15-getreservedinstancesexchangequote.md)
  - [ImportImage](api-ec2-2016-09-15-importimage.md)
  - [ImportInstance](api-ec2-2016-09-15-importinstance.md)
  - [ImportKeyPair](api-ec2-2016-09-15-importkeypair.md)
  - [ImportSnapshot](api-ec2-2016-09-15-importsnapshot.md)
  - [ImportVolume](api-ec2-2016-09-15-importvolume.md)
  - [ModifyHosts](api-ec2-2016-09-15-modifyhosts.md)
  - [ModifyIdFormat](api-ec2-2016-09-15-modifyidformat.md)
  - [ModifyIdentityIdFormat](api-ec2-2016-09-15-modifyidentityidformat.md)
  - [ModifyImageAttribute](api-ec2-2016-09-15-modifyimageattribute.md)
  - [ModifyInstanceAttribute](api-ec2-2016-09-15-modifyinstanceattribute.md)
  - [ModifyInstancePlacement](api-ec2-2016-09-15-modifyinstanceplacement.md)
  - [ModifyNetworkInterfaceAttribute](api-ec2-2016-09-15-modifynetworkinterfaceattribute.md)
  - [ModifyReservedInstances](api-ec2-2016-09-15-modifyreservedinstances.md)
  - [ModifySnapshotAttribute](api-ec2-2016-09-15-modifysnapshotattribute.md)
  - [ModifySpotFleetRequest](api-ec2-2016-09-15-modifyspotfleetrequest.md)
  - [ModifySubnetAttribute](api-ec2-2016-09-15-modifysubnetattribute.md)
  - [ModifyVolumeAttribute](api-ec2-2016-09-15-modifyvolumeattribute.md)
  - [ModifyVpcAttribute](api-ec2-2016-09-15-modifyvpcattribute.md)
  - [ModifyVpcEndpoint](api-ec2-2016-09-15-modifyvpcendpoint.md)
  - [ModifyVpcPeeringConnectionOptions](api-ec2-2016-09-15-modifyvpcpeeringconnectionoptions.md)
  - [MonitorInstances](api-ec2-2016-09-15-monitorinstances.md)
  - [MoveAddressToVpc](api-ec2-2016-09-15-moveaddresstovpc.md)
  - [PurchaseHostReservation](api-ec2-2016-09-15-purchasehostreservation.md)
  - [PurchaseReservedInstancesOffering](api-ec2-2016-09-15-purchasereservedinstancesoffering.md)
  - [PurchaseScheduledInstances](api-ec2-2016-09-15-purchasescheduledinstances.md)
  - [RebootInstances](api-ec2-2016-09-15-rebootinstances.md)
  - [RegisterImage](api-ec2-2016-09-15-registerimage.md)
  - [RejectVpcPeeringConnection](api-ec2-2016-09-15-rejectvpcpeeringconnection.md)
  - [ReleaseAddress](api-ec2-2016-09-15-releaseaddress.md)
  - [ReleaseHosts](api-ec2-2016-09-15-releasehosts.md)
  - [ReplaceNetworkAclAssociation](api-ec2-2016-09-15-replacenetworkaclassociation.md)
  - [ReplaceNetworkAclEntry](api-ec2-2016-09-15-replacenetworkaclentry.md)
  - [ReplaceRoute](api-ec2-2016-09-15-replaceroute.md)
  - [ReplaceRouteTableAssociation](api-ec2-2016-09-15-replaceroutetableassociation.md)
  - [ReportInstanceStatus](api-ec2-2016-09-15-reportinstancestatus.md)
  - [RequestSpotFleet](api-ec2-2016-09-15-requestspotfleet.md)
  - [RequestSpotInstances](api-ec2-2016-09-15-requestspotinstances.md)
  - [ResetImageAttribute](api-ec2-2016-09-15-resetimageattribute.md)
  - [ResetInstanceAttribute](api-ec2-2016-09-15-resetinstanceattribute.md)
  - [ResetNetworkInterfaceAttribute](api-ec2-2016-09-15-resetnetworkinterfaceattribute.md)
  - [ResetSnapshotAttribute](api-ec2-2016-09-15-resetsnapshotattribute.md)
  - [RestoreAddressToClassic](api-ec2-2016-09-15-restoreaddresstoclassic.md)
  - [RevokeSecurityGroupEgress](api-ec2-2016-09-15-revokesecuritygroupegress.md)
  - [RevokeSecurityGroupIngress](api-ec2-2016-09-15-revokesecuritygroupingress.md)
  - [RunInstances](api-ec2-2016-09-15-runinstances.md)
  - [RunScheduledInstances](api-ec2-2016-09-15-runscheduledinstances.md)
  - [StartInstances](api-ec2-2016-09-15-startinstances.md)
  - [StopInstances](api-ec2-2016-09-15-stopinstances.md)
  - [TerminateInstances](api-ec2-2016-09-15-terminateinstances.md)
  - [UnassignPrivateIpAddresses](api-ec2-2016-09-15-unassignprivateipaddresses.md)
  - [UnmonitorInstances](api-ec2-2016-09-15-unmonitorinstances.md)
- [**2016-04-01**](api-ec2-2016-04-01.md)

  - [AcceptVpcPeeringConnection](api-ec2-2016-04-01-acceptvpcpeeringconnection.md)
  - [AllocateAddress](api-ec2-2016-04-01-allocateaddress.md)
  - [AllocateHosts](api-ec2-2016-04-01-allocatehosts.md)
  - [AssignPrivateIpAddresses](api-ec2-2016-04-01-assignprivateipaddresses.md)
  - [AssociateAddress](api-ec2-2016-04-01-associateaddress.md)
  - [AssociateDhcpOptions](api-ec2-2016-04-01-associatedhcpoptions.md)
  - [AssociateRouteTable](api-ec2-2016-04-01-associateroutetable.md)
  - [AttachClassicLinkVpc](api-ec2-2016-04-01-attachclassiclinkvpc.md)
  - [AttachInternetGateway](api-ec2-2016-04-01-attachinternetgateway.md)
  - [AttachNetworkInterface](api-ec2-2016-04-01-attachnetworkinterface.md)
  - [AttachVolume](api-ec2-2016-04-01-attachvolume.md)
  - [AttachVpnGateway](api-ec2-2016-04-01-attachvpngateway.md)
  - [AuthorizeSecurityGroupEgress](api-ec2-2016-04-01-authorizesecuritygroupegress.md)
  - [AuthorizeSecurityGroupIngress](api-ec2-2016-04-01-authorizesecuritygroupingress.md)
  - [BundleInstance](api-ec2-2016-04-01-bundleinstance.md)
  - [CancelBundleTask](api-ec2-2016-04-01-cancelbundletask.md)
  - [CancelConversionTask](api-ec2-2016-04-01-cancelconversiontask.md)
  - [CancelExportTask](api-ec2-2016-04-01-cancelexporttask.md)
  - [CancelImportTask](api-ec2-2016-04-01-cancelimporttask.md)
  - [CancelReservedInstancesListing](api-ec2-2016-04-01-cancelreservedinstanceslisting.md)
  - [CancelSpotFleetRequests](api-ec2-2016-04-01-cancelspotfleetrequests.md)
  - [CancelSpotInstanceRequests](api-ec2-2016-04-01-cancelspotinstancerequests.md)
  - [ConfirmProductInstance](api-ec2-2016-04-01-confirmproductinstance.md)
  - [CopyImage](api-ec2-2016-04-01-copyimage.md)
  - [CopySnapshot](api-ec2-2016-04-01-copysnapshot.md)
  - [CreateCustomerGateway](api-ec2-2016-04-01-createcustomergateway.md)
  - [CreateDhcpOptions](api-ec2-2016-04-01-createdhcpoptions.md)
  - [CreateFlowLogs](api-ec2-2016-04-01-createflowlogs.md)
  - [CreateImage](api-ec2-2016-04-01-createimage.md)
  - [CreateInstanceExportTask](api-ec2-2016-04-01-createinstanceexporttask.md)
  - [CreateInternetGateway](api-ec2-2016-04-01-createinternetgateway.md)
  - [CreateKeyPair](api-ec2-2016-04-01-createkeypair.md)
  - [CreateNatGateway](api-ec2-2016-04-01-createnatgateway.md)
  - [CreateNetworkAcl](api-ec2-2016-04-01-createnetworkacl.md)
  - [CreateNetworkAclEntry](api-ec2-2016-04-01-createnetworkaclentry.md)
  - [CreateNetworkInterface](api-ec2-2016-04-01-createnetworkinterface.md)
  - [CreatePlacementGroup](api-ec2-2016-04-01-createplacementgroup.md)
  - [CreateReservedInstancesListing](api-ec2-2016-04-01-createreservedinstanceslisting.md)
  - [CreateRoute](api-ec2-2016-04-01-createroute.md)
  - [CreateRouteTable](api-ec2-2016-04-01-createroutetable.md)
  - [CreateSecurityGroup](api-ec2-2016-04-01-createsecuritygroup.md)
  - [CreateSnapshot](api-ec2-2016-04-01-createsnapshot.md)
  - [CreateSpotDatafeedSubscription](api-ec2-2016-04-01-createspotdatafeedsubscription.md)
  - [CreateSubnet](api-ec2-2016-04-01-createsubnet.md)
  - [CreateTags](api-ec2-2016-04-01-createtags.md)
  - [CreateVolume](api-ec2-2016-04-01-createvolume.md)
  - [CreateVpc](api-ec2-2016-04-01-createvpc.md)
  - [CreateVpcEndpoint](api-ec2-2016-04-01-createvpcendpoint.md)
  - [CreateVpcPeeringConnection](api-ec2-2016-04-01-createvpcpeeringconnection.md)
  - [CreateVpnConnection](api-ec2-2016-04-01-createvpnconnection.md)
  - [CreateVpnConnectionRoute](api-ec2-2016-04-01-createvpnconnectionroute.md)
  - [CreateVpnGateway](api-ec2-2016-04-01-createvpngateway.md)
  - [DeleteCustomerGateway](api-ec2-2016-04-01-deletecustomergateway.md)
  - [DeleteDhcpOptions](api-ec2-2016-04-01-deletedhcpoptions.md)
  - [DeleteFlowLogs](api-ec2-2016-04-01-deleteflowlogs.md)
  - [DeleteInternetGateway](api-ec2-2016-04-01-deleteinternetgateway.md)
  - [DeleteKeyPair](api-ec2-2016-04-01-deletekeypair.md)
  - [DeleteNatGateway](api-ec2-2016-04-01-deletenatgateway.md)
  - [DeleteNetworkAcl](api-ec2-2016-04-01-deletenetworkacl.md)
  - [DeleteNetworkAclEntry](api-ec2-2016-04-01-deletenetworkaclentry.md)
  - [DeleteNetworkInterface](api-ec2-2016-04-01-deletenetworkinterface.md)
  - [DeletePlacementGroup](api-ec2-2016-04-01-deleteplacementgroup.md)
  - [DeleteRoute](api-ec2-2016-04-01-deleteroute.md)
  - [DeleteRouteTable](api-ec2-2016-04-01-deleteroutetable.md)
  - [DeleteSecurityGroup](api-ec2-2016-04-01-deletesecuritygroup.md)
  - [DeleteSnapshot](api-ec2-2016-04-01-deletesnapshot.md)
  - [DeleteSpotDatafeedSubscription](api-ec2-2016-04-01-deletespotdatafeedsubscription.md)
  - [DeleteSubnet](api-ec2-2016-04-01-deletesubnet.md)
  - [DeleteTags](api-ec2-2016-04-01-deletetags.md)
  - [DeleteVolume](api-ec2-2016-04-01-deletevolume.md)
  - [DeleteVpc](api-ec2-2016-04-01-deletevpc.md)
  - [DeleteVpcEndpoints](api-ec2-2016-04-01-deletevpcendpoints.md)
  - [DeleteVpcPeeringConnection](api-ec2-2016-04-01-deletevpcpeeringconnection.md)
  - [DeleteVpnConnection](api-ec2-2016-04-01-deletevpnconnection.md)
  - [DeleteVpnConnectionRoute](api-ec2-2016-04-01-deletevpnconnectionroute.md)
  - [DeleteVpnGateway](api-ec2-2016-04-01-deletevpngateway.md)
  - [DeregisterImage](api-ec2-2016-04-01-deregisterimage.md)
  - [DescribeAccountAttributes](api-ec2-2016-04-01-describeaccountattributes.md)
  - [DescribeAddresses](api-ec2-2016-04-01-describeaddresses.md)
  - [DescribeAvailabilityZones](api-ec2-2016-04-01-describeavailabilityzones.md)
  - [DescribeBundleTasks](api-ec2-2016-04-01-describebundletasks.md)
  - [DescribeClassicLinkInstances](api-ec2-2016-04-01-describeclassiclinkinstances.md)
  - [DescribeConversionTasks](api-ec2-2016-04-01-describeconversiontasks.md)
  - [DescribeCustomerGateways](api-ec2-2016-04-01-describecustomergateways.md)
  - [DescribeDhcpOptions](api-ec2-2016-04-01-describedhcpoptions.md)
  - [DescribeExportTasks](api-ec2-2016-04-01-describeexporttasks.md)
  - [DescribeFlowLogs](api-ec2-2016-04-01-describeflowlogs.md)
  - [DescribeHostReservationOfferings](api-ec2-2016-04-01-describehostreservationofferings.md)
  - [DescribeHostReservations](api-ec2-2016-04-01-describehostreservations.md)
  - [DescribeHosts](api-ec2-2016-04-01-describehosts.md)
  - [DescribeIdFormat](api-ec2-2016-04-01-describeidformat.md)
  - [DescribeIdentityIdFormat](api-ec2-2016-04-01-describeidentityidformat.md)
  - [DescribeImageAttribute](api-ec2-2016-04-01-describeimageattribute.md)
  - [DescribeImages](api-ec2-2016-04-01-describeimages.md)
  - [DescribeImportImageTasks](api-ec2-2016-04-01-describeimportimagetasks.md)
  - [DescribeImportSnapshotTasks](api-ec2-2016-04-01-describeimportsnapshottasks.md)
  - [DescribeInstanceAttribute](api-ec2-2016-04-01-describeinstanceattribute.md)
  - [DescribeInstanceStatus](api-ec2-2016-04-01-describeinstancestatus.md)
  - [DescribeInstances](api-ec2-2016-04-01-describeinstances.md)
  - [DescribeInternetGateways](api-ec2-2016-04-01-describeinternetgateways.md)
  - [DescribeKeyPairs](api-ec2-2016-04-01-describekeypairs.md)
  - [DescribeMovingAddresses](api-ec2-2016-04-01-describemovingaddresses.md)
  - [DescribeNatGateways](api-ec2-2016-04-01-describenatgateways.md)
  - [DescribeNetworkAcls](api-ec2-2016-04-01-describenetworkacls.md)
  - [DescribeNetworkInterfaceAttribute](api-ec2-2016-04-01-describenetworkinterfaceattribute.md)
  - [DescribeNetworkInterfaces](api-ec2-2016-04-01-describenetworkinterfaces.md)
  - [DescribePlacementGroups](api-ec2-2016-04-01-describeplacementgroups.md)
  - [DescribePrefixLists](api-ec2-2016-04-01-describeprefixlists.md)
  - [DescribeRegions](api-ec2-2016-04-01-describeregions.md)
  - [DescribeReservedInstances](api-ec2-2016-04-01-describereservedinstances.md)
  - [DescribeReservedInstancesListings](api-ec2-2016-04-01-describereservedinstanceslistings.md)
  - [DescribeReservedInstancesModifications](api-ec2-2016-04-01-describereservedinstancesmodifications.md)
  - [DescribeReservedInstancesOfferings](api-ec2-2016-04-01-describereservedinstancesofferings.md)
  - [DescribeRouteTables](api-ec2-2016-04-01-describeroutetables.md)
  - [DescribeScheduledInstanceAvailability](api-ec2-2016-04-01-describescheduledinstanceavailability.md)
  - [DescribeScheduledInstances](api-ec2-2016-04-01-describescheduledinstances.md)
  - [DescribeSecurityGroupReferences](api-ec2-2016-04-01-describesecuritygroupreferences.md)
  - [DescribeSecurityGroups](api-ec2-2016-04-01-describesecuritygroups.md)
  - [DescribeSnapshotAttribute](api-ec2-2016-04-01-describesnapshotattribute.md)
  - [DescribeSnapshots](api-ec2-2016-04-01-describesnapshots.md)
  - [DescribeSpotDatafeedSubscription](api-ec2-2016-04-01-describespotdatafeedsubscription.md)
  - [DescribeSpotFleetInstances](api-ec2-2016-04-01-describespotfleetinstances.md)
  - [DescribeSpotFleetRequestHistory](api-ec2-2016-04-01-describespotfleetrequesthistory.md)
  - [DescribeSpotFleetRequests](api-ec2-2016-04-01-describespotfleetrequests.md)
  - [DescribeSpotInstanceRequests](api-ec2-2016-04-01-describespotinstancerequests.md)
  - [DescribeSpotPriceHistory](api-ec2-2016-04-01-describespotpricehistory.md)
  - [DescribeStaleSecurityGroups](api-ec2-2016-04-01-describestalesecuritygroups.md)
  - [DescribeSubnets](api-ec2-2016-04-01-describesubnets.md)
  - [DescribeTags](api-ec2-2016-04-01-describetags.md)
  - [DescribeVolumeAttribute](api-ec2-2016-04-01-describevolumeattribute.md)
  - [DescribeVolumeStatus](api-ec2-2016-04-01-describevolumestatus.md)
  - [DescribeVolumes](api-ec2-2016-04-01-describevolumes.md)
  - [DescribeVpcAttribute](api-ec2-2016-04-01-describevpcattribute.md)
  - [DescribeVpcClassicLink](api-ec2-2016-04-01-describevpcclassiclink.md)
  - [DescribeVpcClassicLinkDnsSupport](api-ec2-2016-04-01-describevpcclassiclinkdnssupport.md)
  - [DescribeVpcEndpointServices](api-ec2-2016-04-01-describevpcendpointservices.md)
  - [DescribeVpcEndpoints](api-ec2-2016-04-01-describevpcendpoints.md)
  - [DescribeVpcPeeringConnections](api-ec2-2016-04-01-describevpcpeeringconnections.md)
  - [DescribeVpcs](api-ec2-2016-04-01-describevpcs.md)
  - [DescribeVpnConnections](api-ec2-2016-04-01-describevpnconnections.md)
  - [DescribeVpnGateways](api-ec2-2016-04-01-describevpngateways.md)
  - [DetachClassicLinkVpc](api-ec2-2016-04-01-detachclassiclinkvpc.md)
  - [DetachInternetGateway](api-ec2-2016-04-01-detachinternetgateway.md)
  - [DetachNetworkInterface](api-ec2-2016-04-01-detachnetworkinterface.md)
  - [DetachVolume](api-ec2-2016-04-01-detachvolume.md)
  - [DetachVpnGateway](api-ec2-2016-04-01-detachvpngateway.md)
  - [DisableVgwRoutePropagation](api-ec2-2016-04-01-disablevgwroutepropagation.md)
  - [DisableVpcClassicLink](api-ec2-2016-04-01-disablevpcclassiclink.md)
  - [DisableVpcClassicLinkDnsSupport](api-ec2-2016-04-01-disablevpcclassiclinkdnssupport.md)
  - [DisassociateAddress](api-ec2-2016-04-01-disassociateaddress.md)
  - [DisassociateRouteTable](api-ec2-2016-04-01-disassociateroutetable.md)
  - [EnableVgwRoutePropagation](api-ec2-2016-04-01-enablevgwroutepropagation.md)
  - [EnableVolumeIO](api-ec2-2016-04-01-enablevolumeio.md)
  - [EnableVpcClassicLink](api-ec2-2016-04-01-enablevpcclassiclink.md)
  - [EnableVpcClassicLinkDnsSupport](api-ec2-2016-04-01-enablevpcclassiclinkdnssupport.md)
  - [GetConsoleOutput](api-ec2-2016-04-01-getconsoleoutput.md)
  - [GetConsoleScreenshot](api-ec2-2016-04-01-getconsolescreenshot.md)
  - [GetHostReservationPurchasePreview](api-ec2-2016-04-01-gethostreservationpurchasepreview.md)
  - [GetPasswordData](api-ec2-2016-04-01-getpassworddata.md)
  - [ImportImage](api-ec2-2016-04-01-importimage.md)
  - [ImportInstance](api-ec2-2016-04-01-importinstance.md)
  - [ImportKeyPair](api-ec2-2016-04-01-importkeypair.md)
  - [ImportSnapshot](api-ec2-2016-04-01-importsnapshot.md)
  - [ImportVolume](api-ec2-2016-04-01-importvolume.md)
  - [ModifyHosts](api-ec2-2016-04-01-modifyhosts.md)
  - [ModifyIdFormat](api-ec2-2016-04-01-modifyidformat.md)
  - [ModifyIdentityIdFormat](api-ec2-2016-04-01-modifyidentityidformat.md)
  - [ModifyImageAttribute](api-ec2-2016-04-01-modifyimageattribute.md)
  - [ModifyInstanceAttribute](api-ec2-2016-04-01-modifyinstanceattribute.md)
  - [ModifyInstancePlacement](api-ec2-2016-04-01-modifyinstanceplacement.md)
  - [ModifyNetworkInterfaceAttribute](api-ec2-2016-04-01-modifynetworkinterfaceattribute.md)
  - [ModifyReservedInstances](api-ec2-2016-04-01-modifyreservedinstances.md)
  - [ModifySnapshotAttribute](api-ec2-2016-04-01-modifysnapshotattribute.md)
  - [ModifySpotFleetRequest](api-ec2-2016-04-01-modifyspotfleetrequest.md)
  - [ModifySubnetAttribute](api-ec2-2016-04-01-modifysubnetattribute.md)
  - [ModifyVolumeAttribute](api-ec2-2016-04-01-modifyvolumeattribute.md)
  - [ModifyVpcAttribute](api-ec2-2016-04-01-modifyvpcattribute.md)
  - [ModifyVpcEndpoint](api-ec2-2016-04-01-modifyvpcendpoint.md)
  - [ModifyVpcPeeringConnectionOptions](api-ec2-2016-04-01-modifyvpcpeeringconnectionoptions.md)
  - [MonitorInstances](api-ec2-2016-04-01-monitorinstances.md)
  - [MoveAddressToVpc](api-ec2-2016-04-01-moveaddresstovpc.md)
  - [PurchaseHostReservation](api-ec2-2016-04-01-purchasehostreservation.md)
  - [PurchaseReservedInstancesOffering](api-ec2-2016-04-01-purchasereservedinstancesoffering.md)
  - [PurchaseScheduledInstances](api-ec2-2016-04-01-purchasescheduledinstances.md)
  - [RebootInstances](api-ec2-2016-04-01-rebootinstances.md)
  - [RegisterImage](api-ec2-2016-04-01-registerimage.md)
  - [RejectVpcPeeringConnection](api-ec2-2016-04-01-rejectvpcpeeringconnection.md)
  - [ReleaseAddress](api-ec2-2016-04-01-releaseaddress.md)
  - [ReleaseHosts](api-ec2-2016-04-01-releasehosts.md)
  - [ReplaceNetworkAclAssociation](api-ec2-2016-04-01-replacenetworkaclassociation.md)
  - [ReplaceNetworkAclEntry](api-ec2-2016-04-01-replacenetworkaclentry.md)
  - [ReplaceRoute](api-ec2-2016-04-01-replaceroute.md)
  - [ReplaceRouteTableAssociation](api-ec2-2016-04-01-replaceroutetableassociation.md)
  - [ReportInstanceStatus](api-ec2-2016-04-01-reportinstancestatus.md)
  - [RequestSpotFleet](api-ec2-2016-04-01-requestspotfleet.md)
  - [RequestSpotInstances](api-ec2-2016-04-01-requestspotinstances.md)
  - [ResetImageAttribute](api-ec2-2016-04-01-resetimageattribute.md)
  - [ResetInstanceAttribute](api-ec2-2016-04-01-resetinstanceattribute.md)
  - [ResetNetworkInterfaceAttribute](api-ec2-2016-04-01-resetnetworkinterfaceattribute.md)
  - [ResetSnapshotAttribute](api-ec2-2016-04-01-resetsnapshotattribute.md)
  - [RestoreAddressToClassic](api-ec2-2016-04-01-restoreaddresstoclassic.md)
  - [RevokeSecurityGroupEgress](api-ec2-2016-04-01-revokesecuritygroupegress.md)
  - [RevokeSecurityGroupIngress](api-ec2-2016-04-01-revokesecuritygroupingress.md)
  - [RunInstances](api-ec2-2016-04-01-runinstances.md)
  - [RunScheduledInstances](api-ec2-2016-04-01-runscheduledinstances.md)
  - [StartInstances](api-ec2-2016-04-01-startinstances.md)
  - [StopInstances](api-ec2-2016-04-01-stopinstances.md)
  - [TerminateInstances](api-ec2-2016-04-01-terminateinstances.md)
  - [UnassignPrivateIpAddresses](api-ec2-2016-04-01-unassignprivateipaddresses.md)
  - [UnmonitorInstances](api-ec2-2016-04-01-unmonitorinstances.md)
- [**2015-10-01**](api-ec2-2015-10-01.md)

  - [AcceptVpcPeeringConnection](api-ec2-2015-10-01-acceptvpcpeeringconnection.md)
  - [AllocateAddress](api-ec2-2015-10-01-allocateaddress.md)
  - [AllocateHosts](api-ec2-2015-10-01-allocatehosts.md)
  - [AssignPrivateIpAddresses](api-ec2-2015-10-01-assignprivateipaddresses.md)
  - [AssociateAddress](api-ec2-2015-10-01-associateaddress.md)
  - [AssociateDhcpOptions](api-ec2-2015-10-01-associatedhcpoptions.md)
  - [AssociateRouteTable](api-ec2-2015-10-01-associateroutetable.md)
  - [AttachClassicLinkVpc](api-ec2-2015-10-01-attachclassiclinkvpc.md)
  - [AttachInternetGateway](api-ec2-2015-10-01-attachinternetgateway.md)
  - [AttachNetworkInterface](api-ec2-2015-10-01-attachnetworkinterface.md)
  - [AttachVolume](api-ec2-2015-10-01-attachvolume.md)
  - [AttachVpnGateway](api-ec2-2015-10-01-attachvpngateway.md)
  - [AuthorizeSecurityGroupEgress](api-ec2-2015-10-01-authorizesecuritygroupegress.md)
  - [AuthorizeSecurityGroupIngress](api-ec2-2015-10-01-authorizesecuritygroupingress.md)
  - [BundleInstance](api-ec2-2015-10-01-bundleinstance.md)
  - [CancelBundleTask](api-ec2-2015-10-01-cancelbundletask.md)
  - [CancelConversionTask](api-ec2-2015-10-01-cancelconversiontask.md)
  - [CancelExportTask](api-ec2-2015-10-01-cancelexporttask.md)
  - [CancelImportTask](api-ec2-2015-10-01-cancelimporttask.md)
  - [CancelReservedInstancesListing](api-ec2-2015-10-01-cancelreservedinstanceslisting.md)
  - [CancelSpotFleetRequests](api-ec2-2015-10-01-cancelspotfleetrequests.md)
  - [CancelSpotInstanceRequests](api-ec2-2015-10-01-cancelspotinstancerequests.md)
  - [ConfirmProductInstance](api-ec2-2015-10-01-confirmproductinstance.md)
  - [CopyImage](api-ec2-2015-10-01-copyimage.md)
  - [CopySnapshot](api-ec2-2015-10-01-copysnapshot.md)
  - [CreateCustomerGateway](api-ec2-2015-10-01-createcustomergateway.md)
  - [CreateDhcpOptions](api-ec2-2015-10-01-createdhcpoptions.md)
  - [CreateFlowLogs](api-ec2-2015-10-01-createflowlogs.md)
  - [CreateImage](api-ec2-2015-10-01-createimage.md)
  - [CreateInstanceExportTask](api-ec2-2015-10-01-createinstanceexporttask.md)
  - [CreateInternetGateway](api-ec2-2015-10-01-createinternetgateway.md)
  - [CreateKeyPair](api-ec2-2015-10-01-createkeypair.md)
  - [CreateNatGateway](api-ec2-2015-10-01-createnatgateway.md)
  - [CreateNetworkAcl](api-ec2-2015-10-01-createnetworkacl.md)
  - [CreateNetworkAclEntry](api-ec2-2015-10-01-createnetworkaclentry.md)
  - [CreateNetworkInterface](api-ec2-2015-10-01-createnetworkinterface.md)
  - [CreatePlacementGroup](api-ec2-2015-10-01-createplacementgroup.md)
  - [CreateReservedInstancesListing](api-ec2-2015-10-01-createreservedinstanceslisting.md)
  - [CreateRoute](api-ec2-2015-10-01-createroute.md)
  - [CreateRouteTable](api-ec2-2015-10-01-createroutetable.md)
  - [CreateSecurityGroup](api-ec2-2015-10-01-createsecuritygroup.md)
  - [CreateSnapshot](api-ec2-2015-10-01-createsnapshot.md)
  - [CreateSpotDatafeedSubscription](api-ec2-2015-10-01-createspotdatafeedsubscription.md)
  - [CreateSubnet](api-ec2-2015-10-01-createsubnet.md)
  - [CreateTags](api-ec2-2015-10-01-createtags.md)
  - [CreateVolume](api-ec2-2015-10-01-createvolume.md)
  - [CreateVpc](api-ec2-2015-10-01-createvpc.md)
  - [CreateVpcEndpoint](api-ec2-2015-10-01-createvpcendpoint.md)
  - [CreateVpcPeeringConnection](api-ec2-2015-10-01-createvpcpeeringconnection.md)
  - [CreateVpnConnection](api-ec2-2015-10-01-createvpnconnection.md)
  - [CreateVpnConnectionRoute](api-ec2-2015-10-01-createvpnconnectionroute.md)
  - [CreateVpnGateway](api-ec2-2015-10-01-createvpngateway.md)
  - [DeleteCustomerGateway](api-ec2-2015-10-01-deletecustomergateway.md)
  - [DeleteDhcpOptions](api-ec2-2015-10-01-deletedhcpoptions.md)
  - [DeleteFlowLogs](api-ec2-2015-10-01-deleteflowlogs.md)
  - [DeleteInternetGateway](api-ec2-2015-10-01-deleteinternetgateway.md)
  - [DeleteKeyPair](api-ec2-2015-10-01-deletekeypair.md)
  - [DeleteNatGateway](api-ec2-2015-10-01-deletenatgateway.md)
  - [DeleteNetworkAcl](api-ec2-2015-10-01-deletenetworkacl.md)
  - [DeleteNetworkAclEntry](api-ec2-2015-10-01-deletenetworkaclentry.md)
  - [DeleteNetworkInterface](api-ec2-2015-10-01-deletenetworkinterface.md)
  - [DeletePlacementGroup](api-ec2-2015-10-01-deleteplacementgroup.md)
  - [DeleteRoute](api-ec2-2015-10-01-deleteroute.md)
  - [DeleteRouteTable](api-ec2-2015-10-01-deleteroutetable.md)
  - [DeleteSecurityGroup](api-ec2-2015-10-01-deletesecuritygroup.md)
  - [DeleteSnapshot](api-ec2-2015-10-01-deletesnapshot.md)
  - [DeleteSpotDatafeedSubscription](api-ec2-2015-10-01-deletespotdatafeedsubscription.md)
  - [DeleteSubnet](api-ec2-2015-10-01-deletesubnet.md)
  - [DeleteTags](api-ec2-2015-10-01-deletetags.md)
  - [DeleteVolume](api-ec2-2015-10-01-deletevolume.md)
  - [DeleteVpc](api-ec2-2015-10-01-deletevpc.md)
  - [DeleteVpcEndpoints](api-ec2-2015-10-01-deletevpcendpoints.md)
  - [DeleteVpcPeeringConnection](api-ec2-2015-10-01-deletevpcpeeringconnection.md)
  - [DeleteVpnConnection](api-ec2-2015-10-01-deletevpnconnection.md)
  - [DeleteVpnConnectionRoute](api-ec2-2015-10-01-deletevpnconnectionroute.md)
  - [DeleteVpnGateway](api-ec2-2015-10-01-deletevpngateway.md)
  - [DeregisterImage](api-ec2-2015-10-01-deregisterimage.md)
  - [DescribeAccountAttributes](api-ec2-2015-10-01-describeaccountattributes.md)
  - [DescribeAddresses](api-ec2-2015-10-01-describeaddresses.md)
  - [DescribeAvailabilityZones](api-ec2-2015-10-01-describeavailabilityzones.md)
  - [DescribeBundleTasks](api-ec2-2015-10-01-describebundletasks.md)
  - [DescribeClassicLinkInstances](api-ec2-2015-10-01-describeclassiclinkinstances.md)
  - [DescribeConversionTasks](api-ec2-2015-10-01-describeconversiontasks.md)
  - [DescribeCustomerGateways](api-ec2-2015-10-01-describecustomergateways.md)
  - [DescribeDhcpOptions](api-ec2-2015-10-01-describedhcpoptions.md)
  - [DescribeExportTasks](api-ec2-2015-10-01-describeexporttasks.md)
  - [DescribeFlowLogs](api-ec2-2015-10-01-describeflowlogs.md)
  - [DescribeHosts](api-ec2-2015-10-01-describehosts.md)
  - [DescribeIdFormat](api-ec2-2015-10-01-describeidformat.md)
  - [DescribeImageAttribute](api-ec2-2015-10-01-describeimageattribute.md)
  - [DescribeImages](api-ec2-2015-10-01-describeimages.md)
  - [DescribeImportImageTasks](api-ec2-2015-10-01-describeimportimagetasks.md)
  - [DescribeImportSnapshotTasks](api-ec2-2015-10-01-describeimportsnapshottasks.md)
  - [DescribeInstanceAttribute](api-ec2-2015-10-01-describeinstanceattribute.md)
  - [DescribeInstanceStatus](api-ec2-2015-10-01-describeinstancestatus.md)
  - [DescribeInstances](api-ec2-2015-10-01-describeinstances.md)
  - [DescribeInternetGateways](api-ec2-2015-10-01-describeinternetgateways.md)
  - [DescribeKeyPairs](api-ec2-2015-10-01-describekeypairs.md)
  - [DescribeMovingAddresses](api-ec2-2015-10-01-describemovingaddresses.md)
  - [DescribeNatGateways](api-ec2-2015-10-01-describenatgateways.md)
  - [DescribeNetworkAcls](api-ec2-2015-10-01-describenetworkacls.md)
  - [DescribeNetworkInterfaceAttribute](api-ec2-2015-10-01-describenetworkinterfaceattribute.md)
  - [DescribeNetworkInterfaces](api-ec2-2015-10-01-describenetworkinterfaces.md)
  - [DescribePlacementGroups](api-ec2-2015-10-01-describeplacementgroups.md)
  - [DescribePrefixLists](api-ec2-2015-10-01-describeprefixlists.md)
  - [DescribeRegions](api-ec2-2015-10-01-describeregions.md)
  - [DescribeReservedInstances](api-ec2-2015-10-01-describereservedinstances.md)
  - [DescribeReservedInstancesListings](api-ec2-2015-10-01-describereservedinstanceslistings.md)
  - [DescribeReservedInstancesModifications](api-ec2-2015-10-01-describereservedinstancesmodifications.md)
  - [DescribeReservedInstancesOfferings](api-ec2-2015-10-01-describereservedinstancesofferings.md)
  - [DescribeRouteTables](api-ec2-2015-10-01-describeroutetables.md)
  - [DescribeScheduledInstanceAvailability](api-ec2-2015-10-01-describescheduledinstanceavailability.md)
  - [DescribeScheduledInstances](api-ec2-2015-10-01-describescheduledinstances.md)
  - [DescribeSecurityGroupReferences](api-ec2-2015-10-01-describesecuritygroupreferences.md)
  - [DescribeSecurityGroups](api-ec2-2015-10-01-describesecuritygroups.md)
  - [DescribeSnapshotAttribute](api-ec2-2015-10-01-describesnapshotattribute.md)
  - [DescribeSnapshots](api-ec2-2015-10-01-describesnapshots.md)
  - [DescribeSpotDatafeedSubscription](api-ec2-2015-10-01-describespotdatafeedsubscription.md)
  - [DescribeSpotFleetInstances](api-ec2-2015-10-01-describespotfleetinstances.md)
  - [DescribeSpotFleetRequestHistory](api-ec2-2015-10-01-describespotfleetrequesthistory.md)
  - [DescribeSpotFleetRequests](api-ec2-2015-10-01-describespotfleetrequests.md)
  - [DescribeSpotInstanceRequests](api-ec2-2015-10-01-describespotinstancerequests.md)
  - [DescribeSpotPriceHistory](api-ec2-2015-10-01-describespotpricehistory.md)
  - [DescribeStaleSecurityGroups](api-ec2-2015-10-01-describestalesecuritygroups.md)
  - [DescribeSubnets](api-ec2-2015-10-01-describesubnets.md)
  - [DescribeTags](api-ec2-2015-10-01-describetags.md)
  - [DescribeVolumeAttribute](api-ec2-2015-10-01-describevolumeattribute.md)
  - [DescribeVolumeStatus](api-ec2-2015-10-01-describevolumestatus.md)
  - [DescribeVolumes](api-ec2-2015-10-01-describevolumes.md)
  - [DescribeVpcAttribute](api-ec2-2015-10-01-describevpcattribute.md)
  - [DescribeVpcClassicLink](api-ec2-2015-10-01-describevpcclassiclink.md)
  - [DescribeVpcClassicLinkDnsSupport](api-ec2-2015-10-01-describevpcclassiclinkdnssupport.md)
  - [DescribeVpcEndpointServices](api-ec2-2015-10-01-describevpcendpointservices.md)
  - [DescribeVpcEndpoints](api-ec2-2015-10-01-describevpcendpoints.md)
  - [DescribeVpcPeeringConnections](api-ec2-2015-10-01-describevpcpeeringconnections.md)
  - [DescribeVpcs](api-ec2-2015-10-01-describevpcs.md)
  - [DescribeVpnConnections](api-ec2-2015-10-01-describevpnconnections.md)
  - [DescribeVpnGateways](api-ec2-2015-10-01-describevpngateways.md)
  - [DetachClassicLinkVpc](api-ec2-2015-10-01-detachclassiclinkvpc.md)
  - [DetachInternetGateway](api-ec2-2015-10-01-detachinternetgateway.md)
  - [DetachNetworkInterface](api-ec2-2015-10-01-detachnetworkinterface.md)
  - [DetachVolume](api-ec2-2015-10-01-detachvolume.md)
  - [DetachVpnGateway](api-ec2-2015-10-01-detachvpngateway.md)
  - [DisableVgwRoutePropagation](api-ec2-2015-10-01-disablevgwroutepropagation.md)
  - [DisableVpcClassicLink](api-ec2-2015-10-01-disablevpcclassiclink.md)
  - [DisableVpcClassicLinkDnsSupport](api-ec2-2015-10-01-disablevpcclassiclinkdnssupport.md)
  - [DisassociateAddress](api-ec2-2015-10-01-disassociateaddress.md)
  - [DisassociateRouteTable](api-ec2-2015-10-01-disassociateroutetable.md)
  - [EnableVgwRoutePropagation](api-ec2-2015-10-01-enablevgwroutepropagation.md)
  - [EnableVolumeIO](api-ec2-2015-10-01-enablevolumeio.md)
  - [EnableVpcClassicLink](api-ec2-2015-10-01-enablevpcclassiclink.md)
  - [EnableVpcClassicLinkDnsSupport](api-ec2-2015-10-01-enablevpcclassiclinkdnssupport.md)
  - [GetConsoleOutput](api-ec2-2015-10-01-getconsoleoutput.md)
  - [GetConsoleScreenshot](api-ec2-2015-10-01-getconsolescreenshot.md)
  - [GetPasswordData](api-ec2-2015-10-01-getpassworddata.md)
  - [ImportImage](api-ec2-2015-10-01-importimage.md)
  - [ImportInstance](api-ec2-2015-10-01-importinstance.md)
  - [ImportKeyPair](api-ec2-2015-10-01-importkeypair.md)
  - [ImportSnapshot](api-ec2-2015-10-01-importsnapshot.md)
  - [ImportVolume](api-ec2-2015-10-01-importvolume.md)
  - [ModifyHosts](api-ec2-2015-10-01-modifyhosts.md)
  - [ModifyIdFormat](api-ec2-2015-10-01-modifyidformat.md)
  - [ModifyImageAttribute](api-ec2-2015-10-01-modifyimageattribute.md)
  - [ModifyInstanceAttribute](api-ec2-2015-10-01-modifyinstanceattribute.md)
  - [ModifyInstancePlacement](api-ec2-2015-10-01-modifyinstanceplacement.md)
  - [ModifyNetworkInterfaceAttribute](api-ec2-2015-10-01-modifynetworkinterfaceattribute.md)
  - [ModifyReservedInstances](api-ec2-2015-10-01-modifyreservedinstances.md)
  - [ModifySnapshotAttribute](api-ec2-2015-10-01-modifysnapshotattribute.md)
  - [ModifySpotFleetRequest](api-ec2-2015-10-01-modifyspotfleetrequest.md)
  - [ModifySubnetAttribute](api-ec2-2015-10-01-modifysubnetattribute.md)
  - [ModifyVolumeAttribute](api-ec2-2015-10-01-modifyvolumeattribute.md)
  - [ModifyVpcAttribute](api-ec2-2015-10-01-modifyvpcattribute.md)
  - [ModifyVpcEndpoint](api-ec2-2015-10-01-modifyvpcendpoint.md)
  - [ModifyVpcPeeringConnectionOptions](api-ec2-2015-10-01-modifyvpcpeeringconnectionoptions.md)
  - [MonitorInstances](api-ec2-2015-10-01-monitorinstances.md)
  - [MoveAddressToVpc](api-ec2-2015-10-01-moveaddresstovpc.md)
  - [PurchaseReservedInstancesOffering](api-ec2-2015-10-01-purchasereservedinstancesoffering.md)
  - [PurchaseScheduledInstances](api-ec2-2015-10-01-purchasescheduledinstances.md)
  - [RebootInstances](api-ec2-2015-10-01-rebootinstances.md)
  - [RegisterImage](api-ec2-2015-10-01-registerimage.md)
  - [RejectVpcPeeringConnection](api-ec2-2015-10-01-rejectvpcpeeringconnection.md)
  - [ReleaseAddress](api-ec2-2015-10-01-releaseaddress.md)
  - [ReleaseHosts](api-ec2-2015-10-01-releasehosts.md)
  - [ReplaceNetworkAclAssociation](api-ec2-2015-10-01-replacenetworkaclassociation.md)
  - [ReplaceNetworkAclEntry](api-ec2-2015-10-01-replacenetworkaclentry.md)
  - [ReplaceRoute](api-ec2-2015-10-01-replaceroute.md)
  - [ReplaceRouteTableAssociation](api-ec2-2015-10-01-replaceroutetableassociation.md)
  - [ReportInstanceStatus](api-ec2-2015-10-01-reportinstancestatus.md)
  - [RequestSpotFleet](api-ec2-2015-10-01-requestspotfleet.md)
  - [RequestSpotInstances](api-ec2-2015-10-01-requestspotinstances.md)
  - [ResetImageAttribute](api-ec2-2015-10-01-resetimageattribute.md)
  - [ResetInstanceAttribute](api-ec2-2015-10-01-resetinstanceattribute.md)
  - [ResetNetworkInterfaceAttribute](api-ec2-2015-10-01-resetnetworkinterfaceattribute.md)
  - [ResetSnapshotAttribute](api-ec2-2015-10-01-resetsnapshotattribute.md)
  - [RestoreAddressToClassic](api-ec2-2015-10-01-restoreaddresstoclassic.md)
  - [RevokeSecurityGroupEgress](api-ec2-2015-10-01-revokesecuritygroupegress.md)
  - [RevokeSecurityGroupIngress](api-ec2-2015-10-01-revokesecuritygroupingress.md)
  - [RunInstances](api-ec2-2015-10-01-runinstances.md)
  - [RunScheduledInstances](api-ec2-2015-10-01-runscheduledinstances.md)
  - [StartInstances](api-ec2-2015-10-01-startinstances.md)
  - [StopInstances](api-ec2-2015-10-01-stopinstances.md)
  - [TerminateInstances](api-ec2-2015-10-01-terminateinstances.md)
  - [UnassignPrivateIpAddresses](api-ec2-2015-10-01-unassignprivateipaddresses.md)
  - [UnmonitorInstances](api-ec2-2015-10-01-unmonitorinstances.md)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with Amazon Elastic Compute Cloud.

- [See examples on AWS Docs](../../../sdk-for-php/v3/developer-guide/php-ec2-code-examples.md)

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](../../../sdk-for-php/v3/developer-guide/ec2-examples.md).

- [Managing Amazon EC2 instances](../../../sdk-for-php/v3/developer-guide/ec2-examples-managing-instances.md)
- [Using Elastic IP addresses](../../../sdk-for-php/v3/developer-guide/ec2-examples-using-elastic-ip-addresses.md)
- [Using regions and availability zones](../../../sdk-for-php/v3/developer-guide/ec2-examples-using-regions-and-zones.md)
- [Working with key pairs](../../../sdk-for-php/v3/developer-guide/ec2-examples-working-with-key-pairs.md)
- [Working with security groups](../../../sdk-for-php/v3/developer-guide/ec2-examples-using-security-groups.md)

### Table of Contents  [header link](class-aws-ec2-ec2client-toc.md)

#### Methods  [header link](class-aws-ec2-ec2client-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-ec2-ec2client-method-construct.md)
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
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-ec2-ec2client-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-ec2-ec2client-method-construct.md)

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
  - [Methods](class-aws-ec2-ec2client-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-ec2-ec2client-method-construct.md)
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
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-ec2-ec2client-top.md)
