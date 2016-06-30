// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/gtm"

type GTM struct {
	c f5.Client

	datacenter                      DatacenterResource
	distributedApp                  DistributedAppResource
	globalSettings                  GlobalSettingsResource
	globalSettingsGeneral           GlobalSettingsGeneralResource
	globalSettingsLoadBalancing     GlobalSettingsLoadBalancingResource
	globalSettingsMetrics           GlobalSettingsMetricsResource
	globalSettingsMetricsExclusions GlobalSettingsMetricsExclusionsResource
	link                            LinkResource
	listener                        ListenerResource
	listenerProfiles                ListenerProfilesResource
	monitor                         MonitorResource
	monitorBigIP                    MonitorBigIPResource
	monitorBigIPLink                MonitorBigIPLinkResource
	monitorExternal                 MonitorExternalResource
	monitorFTP                      MonitorFTPResource
	monitorFirepass                 MonitorFirepassResource
	monitorGTP                      MonitorGTPResource
	monitorHTTP                     MonitorHTTPResource
	monitorHTTPS                    MonitorHTTPSResource
	monitorICMP                     MonitorICMPResource
	monitorIMAP                     MonitorIMAPResource
	monitorLDAP                     MonitorLDAPResource
	monitorMSSQL                    MonitorMSSQLResource
	monitorMySQL                    MonitorMySQLResource
	monitorNNTP                     MonitorNNTPResource
	monitorNone                     MonitorNoneResource
	monitorOracle                   MonitorOracleResource
	monitorPOP3                     MonitorPOP3Resource
	monitorPostgreSQL               MonitorPostgreSQLResource
	monitorRadius                   MonitorRadiusResource
	monitorRadiusAccounting         MonitorRadiusAccountingResource
	monitorRealServer               MonitorRealServerResource
	monitorSIP                      MonitorSIPResource
	monitorSMTP                     MonitorSMTPResource
	monitorSNMP                     MonitorSNMPResource
	monitorSNMPLink                 MonitorSNMPLinkResource
	monitorSOAP                     MonitorSOAPResource
	monitorTCP                      MonitorTCPResource
	monitorTCPHalf                  MonitorTCPHalfResource
	monitorUDP                      MonitorUDPResource
	monitorWAP                      MonitorWAPResource
	monitorWMI                      MonitorWMIResource
	monitorscripted                 MonitorscriptedResource
	persist                         PersistResource
	pool                            PoolResource
	poolA                           PoolAResource
	poolAAAA                        PoolAAAAResource
	poolAAAAMembers                 PoolAAAAMembersResource
	poolAMembers                    PoolAMembersResource
	poolCNAME                       PoolCNAMEResource
	poolCNAMEMembers                PoolCNAMEMembersResource
	poolMX                          PoolMXResource
	poolMXMembers                   PoolMXMembersResource
	poolNAPTR                       PoolNAPTRResource
	poolNAPTRMembers                PoolNAPTRMembersResource
	poolSRV                         PoolSRVResource
	poolSRVMembers                  PoolSRVMembersResource
	proberPool                      ProberPoolResource
	proberPoolMembers               ProberPoolMembersResource
	region                          RegionResource
	rule                            RuleResource
	server                          ServerResource
	serverVirtualServers            ServerVirtualServersResource
	topology                        TopologyResource
	wideIP                          WideIPResource
	wideipA                         WideipAResource
	wideipAAAA                      WideipAAAAResource
	wideipCname                     WideipCnameResource
	wideipMx                        WideipMxResource
	wideipNaptr                     WideipNaptrResource
	wideipSrv                       WideipSrvResource
}

// New creates a new GTM client.
func New(c f5.Client) GTM {
	return GTM{
		c: c,

		datacenter:                      DatacenterResource{c: c},
		distributedApp:                  DistributedAppResource{c: c},
		globalSettings:                  GlobalSettingsResource{c: c},
		globalSettingsGeneral:           GlobalSettingsGeneralResource{c: c},
		globalSettingsLoadBalancing:     GlobalSettingsLoadBalancingResource{c: c},
		globalSettingsMetrics:           GlobalSettingsMetricsResource{c: c},
		globalSettingsMetricsExclusions: GlobalSettingsMetricsExclusionsResource{c: c},
		link:                    LinkResource{c: c},
		listener:                ListenerResource{c: c},
		listenerProfiles:        ListenerProfilesResource{c: c},
		monitor:                 MonitorResource{c: c},
		monitorBigIP:            MonitorBigIPResource{c: c},
		monitorBigIPLink:        MonitorBigIPLinkResource{c: c},
		monitorExternal:         MonitorExternalResource{c: c},
		monitorFTP:              MonitorFTPResource{c: c},
		monitorFirepass:         MonitorFirepassResource{c: c},
		monitorGTP:              MonitorGTPResource{c: c},
		monitorHTTP:             MonitorHTTPResource{c: c},
		monitorHTTPS:            MonitorHTTPSResource{c: c},
		monitorICMP:             MonitorICMPResource{c: c},
		monitorIMAP:             MonitorIMAPResource{c: c},
		monitorLDAP:             MonitorLDAPResource{c: c},
		monitorMSSQL:            MonitorMSSQLResource{c: c},
		monitorMySQL:            MonitorMySQLResource{c: c},
		monitorNNTP:             MonitorNNTPResource{c: c},
		monitorNone:             MonitorNoneResource{c: c},
		monitorOracle:           MonitorOracleResource{c: c},
		monitorPOP3:             MonitorPOP3Resource{c: c},
		monitorPostgreSQL:       MonitorPostgreSQLResource{c: c},
		monitorRadius:           MonitorRadiusResource{c: c},
		monitorRadiusAccounting: MonitorRadiusAccountingResource{c: c},
		monitorRealServer:       MonitorRealServerResource{c: c},
		monitorSIP:              MonitorSIPResource{c: c},
		monitorSMTP:             MonitorSMTPResource{c: c},
		monitorSNMP:             MonitorSNMPResource{c: c},
		monitorSNMPLink:         MonitorSNMPLinkResource{c: c},
		monitorSOAP:             MonitorSOAPResource{c: c},
		monitorTCP:              MonitorTCPResource{c: c},
		monitorTCPHalf:          MonitorTCPHalfResource{c: c},
		monitorUDP:              MonitorUDPResource{c: c},
		monitorWAP:              MonitorWAPResource{c: c},
		monitorWMI:              MonitorWMIResource{c: c},
		monitorscripted:         MonitorscriptedResource{c: c},
		persist:                 PersistResource{c: c},
		pool:                    PoolResource{c: c},
		poolA:                   PoolAResource{c: c},
		poolAAAA:                PoolAAAAResource{c: c},
		poolAAAAMembers:         PoolAAAAMembersResource{c: c},
		poolAMembers:            PoolAMembersResource{c: c},
		poolCNAME:               PoolCNAMEResource{c: c},
		poolCNAMEMembers:        PoolCNAMEMembersResource{c: c},
		poolMX:                  PoolMXResource{c: c},
		poolMXMembers:           PoolMXMembersResource{c: c},
		poolNAPTR:               PoolNAPTRResource{c: c},
		poolNAPTRMembers:        PoolNAPTRMembersResource{c: c},
		poolSRV:                 PoolSRVResource{c: c},
		poolSRVMembers:          PoolSRVMembersResource{c: c},
		proberPool:              ProberPoolResource{c: c},
		proberPoolMembers:       ProberPoolMembersResource{c: c},
		region:                  RegionResource{c: c},
		rule:                    RuleResource{c: c},
		server:                  ServerResource{c: c},
		serverVirtualServers:    ServerVirtualServersResource{c: c},
		topology:                TopologyResource{c: c},
		wideIP:                  WideIPResource{c: c},
		wideipA:                 WideipAResource{c: c},
		wideipAAAA:              WideipAAAAResource{c: c},
		wideipCname:             WideipCnameResource{c: c},
		wideipMx:                WideipMxResource{c: c},
		wideipNaptr:             WideipNaptrResource{c: c},
		wideipSrv:               WideipSrvResource{c: c},
	}
}

// Datacenter returns a configured DatacenterResource.
func (gtm GTM) Datacenter() *DatacenterResource {
	return &gtm.datacenter
}

// DistributedApp returns a configured DistributedAppResource.
func (gtm GTM) DistributedApp() *DistributedAppResource {
	return &gtm.distributedApp
}

// GlobalSettings returns a configured GlobalSettingsResource.
func (gtm GTM) GlobalSettings() *GlobalSettingsResource {
	return &gtm.globalSettings
}

// GlobalSettingsGeneral returns a configured GlobalSettingsGeneralResource.
func (gtm GTM) GlobalSettingsGeneral() *GlobalSettingsGeneralResource {
	return &gtm.globalSettingsGeneral
}

// GlobalSettingsLoadBalancing returns a configured GlobalSettingsLoadBalancingResource.
func (gtm GTM) GlobalSettingsLoadBalancing() *GlobalSettingsLoadBalancingResource {
	return &gtm.globalSettingsLoadBalancing
}

// GlobalSettingsMetrics returns a configured GlobalSettingsMetricsResource.
func (gtm GTM) GlobalSettingsMetrics() *GlobalSettingsMetricsResource {
	return &gtm.globalSettingsMetrics
}

// GlobalSettingsMetricsExclusions returns a configured GlobalSettingsMetricsExclusionsResource.
func (gtm GTM) GlobalSettingsMetricsExclusions() *GlobalSettingsMetricsExclusionsResource {
	return &gtm.globalSettingsMetricsExclusions
}

// Link returns a configured LinkResource.
func (gtm GTM) Link() *LinkResource {
	return &gtm.link
}

// Listener returns a configured ListenerResource.
func (gtm GTM) Listener() *ListenerResource {
	return &gtm.listener
}

// ListenerProfiles returns a configured ListenerProfilesResource.
func (gtm GTM) ListenerProfiles() *ListenerProfilesResource {
	return &gtm.listenerProfiles
}

// Monitor returns a configured MonitorResource.
func (gtm GTM) Monitor() *MonitorResource {
	return &gtm.monitor
}

// MonitorBigIP returns a configured MonitorBigIPResource.
func (gtm GTM) MonitorBigIP() *MonitorBigIPResource {
	return &gtm.monitorBigIP
}

// MonitorBigIPLink returns a configured MonitorBigIPLinkResource.
func (gtm GTM) MonitorBigIPLink() *MonitorBigIPLinkResource {
	return &gtm.monitorBigIPLink
}

// MonitorExternal returns a configured MonitorExternalResource.
func (gtm GTM) MonitorExternal() *MonitorExternalResource {
	return &gtm.monitorExternal
}

// MonitorFTP returns a configured MonitorFTPResource.
func (gtm GTM) MonitorFTP() *MonitorFTPResource {
	return &gtm.monitorFTP
}

// MonitorFirepass returns a configured MonitorFirepassResource.
func (gtm GTM) MonitorFirepass() *MonitorFirepassResource {
	return &gtm.monitorFirepass
}

// MonitorGTP returns a configured MonitorGTPResource.
func (gtm GTM) MonitorGTP() *MonitorGTPResource {
	return &gtm.monitorGTP
}

// MonitorHTTP returns a configured MonitorHTTPResource.
func (gtm GTM) MonitorHTTP() *MonitorHTTPResource {
	return &gtm.monitorHTTP
}

// MonitorHTTPS returns a configured MonitorHTTPSResource.
func (gtm GTM) MonitorHTTPS() *MonitorHTTPSResource {
	return &gtm.monitorHTTPS
}

// MonitorICMP returns a configured MonitorICMPResource.
func (gtm GTM) MonitorICMP() *MonitorICMPResource {
	return &gtm.monitorICMP
}

// MonitorIMAP returns a configured MonitorIMAPResource.
func (gtm GTM) MonitorIMAP() *MonitorIMAPResource {
	return &gtm.monitorIMAP
}

// MonitorLDAP returns a configured MonitorLDAPResource.
func (gtm GTM) MonitorLDAP() *MonitorLDAPResource {
	return &gtm.monitorLDAP
}

// MonitorMSSQL returns a configured MonitorMSSQLResource.
func (gtm GTM) MonitorMSSQL() *MonitorMSSQLResource {
	return &gtm.monitorMSSQL
}

// MonitorMySQL returns a configured MonitorMySQLResource.
func (gtm GTM) MonitorMySQL() *MonitorMySQLResource {
	return &gtm.monitorMySQL
}

// MonitorNNTP returns a configured MonitorNNTPResource.
func (gtm GTM) MonitorNNTP() *MonitorNNTPResource {
	return &gtm.monitorNNTP
}

// MonitorNone returns a configured MonitorNoneResource.
func (gtm GTM) MonitorNone() *MonitorNoneResource {
	return &gtm.monitorNone
}

// MonitorOracle returns a configured MonitorOracleResource.
func (gtm GTM) MonitorOracle() *MonitorOracleResource {
	return &gtm.monitorOracle
}

// MonitorPOP3 returns a configured MonitorPOP3Resource.
func (gtm GTM) MonitorPOP3() *MonitorPOP3Resource {
	return &gtm.monitorPOP3
}

// MonitorPostgreSQL returns a configured MonitorPostgreSQLResource.
func (gtm GTM) MonitorPostgreSQL() *MonitorPostgreSQLResource {
	return &gtm.monitorPostgreSQL
}

// MonitorRadius returns a configured MonitorRadiusResource.
func (gtm GTM) MonitorRadius() *MonitorRadiusResource {
	return &gtm.monitorRadius
}

// MonitorRadiusAccounting returns a configured MonitorRadiusAccountingResource.
func (gtm GTM) MonitorRadiusAccounting() *MonitorRadiusAccountingResource {
	return &gtm.monitorRadiusAccounting
}

// MonitorRealServer returns a configured MonitorRealServerResource.
func (gtm GTM) MonitorRealServer() *MonitorRealServerResource {
	return &gtm.monitorRealServer
}

// MonitorSIP returns a configured MonitorSIPResource.
func (gtm GTM) MonitorSIP() *MonitorSIPResource {
	return &gtm.monitorSIP
}

// MonitorSMTP returns a configured MonitorSMTPResource.
func (gtm GTM) MonitorSMTP() *MonitorSMTPResource {
	return &gtm.monitorSMTP
}

// MonitorSNMP returns a configured MonitorSNMPResource.
func (gtm GTM) MonitorSNMP() *MonitorSNMPResource {
	return &gtm.monitorSNMP
}

// MonitorSNMPLink returns a configured MonitorSNMPLinkResource.
func (gtm GTM) MonitorSNMPLink() *MonitorSNMPLinkResource {
	return &gtm.monitorSNMPLink
}

// MonitorSOAP returns a configured MonitorSOAPResource.
func (gtm GTM) MonitorSOAP() *MonitorSOAPResource {
	return &gtm.monitorSOAP
}

// MonitorTCP returns a configured MonitorTCPResource.
func (gtm GTM) MonitorTCP() *MonitorTCPResource {
	return &gtm.monitorTCP
}

// MonitorTCPHalf returns a configured MonitorTCPHalfResource.
func (gtm GTM) MonitorTCPHalf() *MonitorTCPHalfResource {
	return &gtm.monitorTCPHalf
}

// MonitorUDP returns a configured MonitorUDPResource.
func (gtm GTM) MonitorUDP() *MonitorUDPResource {
	return &gtm.monitorUDP
}

// MonitorWAP returns a configured MonitorWAPResource.
func (gtm GTM) MonitorWAP() *MonitorWAPResource {
	return &gtm.monitorWAP
}

// MonitorWMI returns a configured MonitorWMIResource.
func (gtm GTM) MonitorWMI() *MonitorWMIResource {
	return &gtm.monitorWMI
}

// Monitorscripted returns a configured MonitorscriptedResource.
func (gtm GTM) Monitorscripted() *MonitorscriptedResource {
	return &gtm.monitorscripted
}

// Persist returns a configured PersistResource.
func (gtm GTM) Persist() *PersistResource {
	return &gtm.persist
}

// Pool returns a configured PoolResource.
func (gtm GTM) Pool() *PoolResource {
	return &gtm.pool
}

// PoolA returns a configured PoolAResource.
func (gtm GTM) PoolA() *PoolAResource {
	return &gtm.poolA
}

// PoolAAAA returns a configured PoolAAAAResource.
func (gtm GTM) PoolAAAA() *PoolAAAAResource {
	return &gtm.poolAAAA
}

// PoolAAAAMembers returns a configured PoolAAAAMembersResource.
func (gtm GTM) PoolAAAAMembers() *PoolAAAAMembersResource {
	return &gtm.poolAAAAMembers
}

// PoolAMembers returns a configured PoolAMembersResource.
func (gtm GTM) PoolAMembers() *PoolAMembersResource {
	return &gtm.poolAMembers
}

// PoolCNAME returns a configured PoolCNAMEResource.
func (gtm GTM) PoolCNAME() *PoolCNAMEResource {
	return &gtm.poolCNAME
}

// PoolCNAMEMembers returns a configured PoolCNAMEMembersResource.
func (gtm GTM) PoolCNAMEMembers() *PoolCNAMEMembersResource {
	return &gtm.poolCNAMEMembers
}

// PoolMX returns a configured PoolMXResource.
func (gtm GTM) PoolMX() *PoolMXResource {
	return &gtm.poolMX
}

// PoolMXMembers returns a configured PoolMXMembersResource.
func (gtm GTM) PoolMXMembers() *PoolMXMembersResource {
	return &gtm.poolMXMembers
}

// PoolNAPTR returns a configured PoolNAPTRResource.
func (gtm GTM) PoolNAPTR() *PoolNAPTRResource {
	return &gtm.poolNAPTR
}

// PoolNAPTRMembers returns a configured PoolNAPTRMembersResource.
func (gtm GTM) PoolNAPTRMembers() *PoolNAPTRMembersResource {
	return &gtm.poolNAPTRMembers
}

// PoolSRV returns a configured PoolSRVResource.
func (gtm GTM) PoolSRV() *PoolSRVResource {
	return &gtm.poolSRV
}

// PoolSRVMembers returns a configured PoolSRVMembersResource.
func (gtm GTM) PoolSRVMembers() *PoolSRVMembersResource {
	return &gtm.poolSRVMembers
}

// ProberPool returns a configured ProberPoolResource.
func (gtm GTM) ProberPool() *ProberPoolResource {
	return &gtm.proberPool
}

// ProberPoolMembers returns a configured ProberPoolMembersResource.
func (gtm GTM) ProberPoolMembers() *ProberPoolMembersResource {
	return &gtm.proberPoolMembers
}

// Region returns a configured RegionResource.
func (gtm GTM) Region() *RegionResource {
	return &gtm.region
}

// Rule returns a configured RuleResource.
func (gtm GTM) Rule() *RuleResource {
	return &gtm.rule
}

// Server returns a configured ServerResource.
func (gtm GTM) Server() *ServerResource {
	return &gtm.server
}

// ServerVirtualServers returns a configured ServerVirtualServersResource.
func (gtm GTM) ServerVirtualServers() *ServerVirtualServersResource {
	return &gtm.serverVirtualServers
}

// Topology returns a configured TopologyResource.
func (gtm GTM) Topology() *TopologyResource {
	return &gtm.topology
}

// WideIP returns a configured WideIPResource.
func (gtm GTM) WideIP() *WideIPResource {
	return &gtm.wideIP
}

// WideipA returns a configured WideipAResource.
func (gtm GTM) WideipA() *WideipAResource {
	return &gtm.wideipA
}

// WideipAAAA returns a configured WideipAAAAResource.
func (gtm GTM) WideipAAAA() *WideipAAAAResource {
	return &gtm.wideipAAAA
}

// WideipCname returns a configured WideipCnameResource.
func (gtm GTM) WideipCname() *WideipCnameResource {
	return &gtm.wideipCname
}

// WideipMx returns a configured WideipMxResource.
func (gtm GTM) WideipMx() *WideipMxResource {
	return &gtm.wideipMx
}

// WideipNaptr returns a configured WideipNaptrResource.
func (gtm GTM) WideipNaptr() *WideipNaptrResource {
	return &gtm.wideipNaptr
}

// WideipSrv returns a configured WideipSrvResource.
func (gtm GTM) WideipSrv() *WideipSrvResource {
	return &gtm.wideipSrv
}
