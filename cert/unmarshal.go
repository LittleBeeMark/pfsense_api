package cert

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"pfsense/cli"
	"strings"
)

// EditCertResp doc
type EditCertResp struct {
	CSRF  string `json:"csrf"`
	Descr string `json:"descr"`
}

var html = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<link rel="apple-touch-icon-precomposed" href="/apple-touch/apple-touch-icon-iphone-60x60-precomposed.png">
	<link rel="apple-touch-icon-precomposed" sizes="60x60" href="/apple-touch/apple-touch-icon-ipad-76x76-precomposed.png">
	<link rel="apple-touch-icon-precomposed" sizes="114x114" href="/apple-touch/apple-touch-icon-iphone-retina-120x120-precomposed.png">
	<link rel="apple-touch-icon-precomposed" sizes="144x144" href="/apple-touch/apple-touch-icon-ipad-retina-152x152-precomposed.png">
	<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
	<link rel="manifest" href="/manifest.json">
	<link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5">
	<meta name="theme-color" content="#ffffff">

	<link rel="stylesheet" href="/vendor/font-awesome/css/all.min.css?v=1622201721">
	<link rel="stylesheet" href="/vendor/font-awesome/css/v4-shims.css?v=1622201721">
	<link rel="stylesheet" href="/vendor/sortable/sortable-theme-bootstrap.css?v=1622201721">
	<link rel="stylesheet" href="/css/pfSense.css?v=1622201721" />

	<title>pfSense.home.arpa - System: Certificate Manager: Certificates</title>
	<script type="text/javascript">
	//<![CDATA[
	var events = events || [];
	var newSeperator = false;
	//]]>
	</script>
<script type="text/javascript">if (top != self) {top.location.href = self.location.href;}</script><script type="text/javascript">var csrfMagicToken = "sid:e40d8c3404e6febd323d9a91d8f4e62e7f411129,1631021452";var csrfMagicName = "__csrf_magic";</script><script src="/csrf/csrf-magic.js" type="text/javascript"></script></head>

<body id="2">
<nav id="topmenu" class="navbar navbar-static-top navbar-inverse">
	<div class="container">
		<div class="navbar-header">
			<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#pf-navbar">
				<span class="sr-only">Toggle navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<a class="navbar-brand" href="/">
				<svg id="logo" role="img" aria-labelledby="pfsense-logo" x="0px" y="0px" viewBox="0 0 282.8 84.2">
	<title id="pfsense-logo-svg">pfSense Logo</title>
	<path class="logo-st0" d="M27.8,57.7c2.9,0,5.4-0.9,7.5-2.6c2.1-1.7,3.6-4,4.4-6.8c0.8-2.8,0.6-5.1-0.5-6.8c-1.1-1.7-3.2-2.6-6.1-2.6 c-2.9,0-5.4,0.9-7.5,2.6c-2.1,1.7-3.5,4-4.3,6.8c-0.8,2.8-0.7,5.1,0.5,6.8C22.8,56.9,24.8,57.7,27.8,57.7"/>
	<path class="logo-st0" d="M115.1,46.6c-1.5-0.8-3-1.4-4.7-1.8c-1.7-0.4-3.2-0.7-4.7-1.1c-1.5-0.3-2.7-0.7-3.6-1.1c-0.9-0.4-1.4-1.1-1.4-2 c0-1.1,0.5-1.9,1.4-2.4c0.9-0.5,1.9-0.7,2.8-0.7c2.8,0,5,1,6.7,3.1l7-7c-1.7-1.8-3.9-3.1-6.4-3.8c-2.5-0.7-5-1.1-7.4-1.1 c-1.9,0-3.9,0.2-5.7,0.7c-1.9,0.5-3.6,1.2-5,2.3c-1.5,1-2.6,2.3-3.5,3.9c-0.9,1.6-1.3,3.5-1.3,5.7c0,2.3,0.5,4.2,1.4,5.6 c0.9,1.4,2.1,2.5,3.6,3.3c1.5,0.8,3,1.3,4.7,1.7c1.7,0.4,3.2,0.7,4.7,1.1c1.5,0.3,2.7,0.7,3.6,1.2c0.9,0.5,1.4,1.2,1.4,2.2 c0,1-0.5,1.7-1.6,2.1c-1.1,0.4-2.3,0.6-3.6,0.6c-1.7,0-3.3-0.3-4.6-1c-1.3-0.7-2.5-1.7-3.6-3l-7,7.7c1.8,1.9,4.1,3.2,6.7,3.9 c2.7,0.7,5.3,1.1,7.9,1.1c2,0,4-0.2,6.1-0.6c2-0.4,3.9-1,5.5-2c1.6-0.9,3-2.2,4-3.8c1-1.6,1.6-3.5,1.6-5.9c0-2.3-0.5-4.2-1.4-5.6 C117.7,48.6,116.5,47.4,115.1,46.6"/>
	<path class="logo-st0" d="M156.3,34.1c-1.5-1.7-3.3-3-5.5-3.9c-2.2-0.9-4.6-1.4-7.2-1.4c-2.9,0-5.6,0.5-8.1,1.4c-2.5,0.9-4.7,2.2-6.6,3.9 c-1.9,1.7-3.3,3.8-4.4,6.2c-1.1,2.4-1.6,5.1-1.6,8c0,3,0.5,5.6,1.6,8c1.1,2.4,2.5,4.5,4.4,6.2c1.9,1.7,4.1,3,6.6,3.9 c2.5,0.9,5.2,1.4,8.1,1.4c3,0,5.9-0.6,8.7-1.9c2.8-1.3,5.1-3.1,7-5.4l-8-5.9c-1,1.3-2.1,2.4-3.4,3.3c-1.3,0.8-2.9,1.3-4.8,1.3 c-2.2,0-4.1-0.7-5.7-2c-1.5-1.3-2.5-3.1-3-5.2H161v-3.6c0-3-0.4-5.6-1.2-8C159,37.9,157.8,35.8,156.3,34.1 M134.3,44.1 c0.1-0.9,0.3-1.8,0.7-2.6c0.4-0.8,0.9-1.6,1.6-2.2c0.7-0.6,1.5-1.2,2.5-1.6c1-0.4,2.1-0.6,3.4-0.6c2.1,0,3.8,0.7,5.1,2.1 c1.3,1.4,2,3,1.9,5H134.3z"/>
	<path class="logo-st0" d="M198.3,33.8c-1-1.6-2.4-2.8-4.2-3.7c-1.8-0.9-4.1-1.3-7-1.3c-1.4,0-2.7,0.2-3.8,0.5c-1.2,0.4-2.2,0.8-3.1,1.4 c-0.9,0.6-1.7,1.2-2.4,1.9c-0.7,0.7-1.2,1.4-1.5,2.1H176v-5.1h-11v37.2h11.5V48.4c0-1.2,0.1-2.4,0.2-3.5c0.2-1.1,0.5-2.1,1-3 c0.5-0.9,1.2-1.6,2.1-2.1c0.9-0.5,2.1-0.8,3.6-0.8c1.5,0,2.6,0.3,3.4,0.9c0.8,0.6,1.4,1.4,1.8,2.4c0.4,1,0.6,2,0.7,3.2 c0.1,1.1,0.1,2.3,0.1,3.3v18.2h11.5V46.4c0-2.5-0.2-4.8-0.5-7C199.9,37.3,199.3,35.4,198.3,33.8"/>
	<path class="logo-st0" d="M231.5,46.6c-1.5-0.8-3-1.4-4.7-1.8c-1.7-0.4-3.2-0.7-4.7-1.1c-1.5-0.3-2.7-0.7-3.6-1.1c-0.9-0.4-1.4-1.1-1.4-2 c0-1.1,0.5-1.9,1.4-2.4c0.9-0.5,1.9-0.7,2.8-0.7c2.8,0,5,1,6.7,3.1l7-7c-1.7-1.8-3.9-3.1-6.4-3.8c-2.5-0.7-5-1.1-7.4-1.1 c-1.9,0-3.9,0.2-5.7,0.7c-1.9,0.5-3.6,1.2-5,2.3c-1.5,1-2.6,2.3-3.5,3.9c-0.9,1.6-1.3,3.5-1.3,5.7c0,2.3,0.5,4.2,1.4,5.6 c0.9,1.4,2.1,2.5,3.6,3.3c1.5,0.8,3,1.3,4.7,1.7c1.7,0.4,3.2,0.7,4.7,1.1c1.5,0.3,2.7,0.7,3.6,1.2c0.9,0.5,1.4,1.2,1.4,2.2 c0,1-0.5,1.7-1.6,2.1c-1.1,0.4-2.3,0.6-3.6,0.6c-1.7,0-3.3-0.3-4.6-1c-1.3-0.7-2.5-1.7-3.6-3l-7,7.7c1.8,1.9,4.1,3.2,6.7,3.9 c2.7,0.7,5.3,1.1,7.9,1.1c2,0,4-0.2,6.1-0.6c2-0.4,3.9-1,5.5-2c1.6-0.9,3-2.2,4-3.8c1-1.6,1.6-3.5,1.6-5.9c0-2.3-0.5-4.2-1.4-5.6 C234.1,48.6,232.9,47.4,231.5,46.6"/>
	<path class="logo-st0" d="M277.4,51.9v-4.2c-0.1-2.7-0.5-5.2-1.2-7.4c-0.8-2.4-2-4.5-3.5-6.2c-1.5-1.7-3.3-3-5.5-3.9 c-2.2-0.9-4.6-1.4-7.2-1.4c-2.9,0-5.6,0.5-8.1,1.4c-2.5,0.9-4.7,2.2-6.6,3.9c-1.9,1.7-3.3,3.8-4.4,6.2c-1.1,2.4-1.6,5.1-1.6,8 c0,3,0.5,5.6,1.6,8c1.1,2.4,2.5,4.5,4.4,6.2c1.9,1.7,4.1,3,6.6,3.9c2.5,0.9,5.2,1.4,8.1,1.4c3,0,5.9-0.6,8.7-1.9 c2.8-1.3,5.1-3.1,7-5.4l-8-5.9c-1,1.3-2.1,2.4-3.4,3.3c-1.3,0.8-2.9,1.3-4.8,1.3c-2.2,0-4.1-0.7-5.7-2c-1.5-1.3-2.5-3.1-3-5.2H277.4 z M250.7,44.1c0.1-0.9,0.3-1.8,0.7-2.6c0.4-0.8,0.9-1.6,1.6-2.2c0.7-0.6,1.5-1.2,2.5-1.6c1-0.4,2.1-0.6,3.4-0.6 c2.1,0,3.8,0.7,5.1,2.1c1.3,1.4,2,3,1.9,5H250.7z"/>
	<path class="logo-st1" d="M52.6,38.9l2.6-9.2h4.6l1.8-6.6c0.6-2,1.3-4,2.2-5.8c0.8-1.8,2-3.4,3.4-4.8c1.4-1.4,3.2-2.5,5.3-3.3 c2.1-0.8,4.8-1.2,7.9-1.2c0.8,0,1.5,0,2.3,0.1c-0.7-2.9-3.3-5-6.3-5.1H11.9c-3.6,0-6.5,3-6.5,6.6V67l10.5-37.3h10.6l-1.4,4.9h0.2 c0.6-0.7,1.4-1.3,2.4-2c1-0.7,2-1.3,3.1-1.9c1.1-0.6,2.3-1,3.6-1.4c1.3-0.4,2.6-0.5,3.9-0.5c2.8,0,5.1,0.5,7.1,1.4 c2,0.9,3.5,2.3,4.7,4c1,1.5,1.6,3.3,1.9,5.4l0.8-0.6H52.6z"/>
	<path class="logo-st2" d="M82.1,17.9c-0.5-0.1-1.1-0.2-1.8-0.2c-1.8,0-3.3,0.4-4.5,1.2c-1.1,0.8-2.1,2.4-2.8,4.9l-1.7,5.9h6.5l1.6,5.1 l-4.2,4.1h-6.5l-7.9,28H49.4l7.9-28h-4.4L52,39.5c0,0.2,0.1,0.5,0.1,0.7c0.2,2.3-0.1,4.9-0.9,7.7c-0.7,2.6-1.8,5.1-3.3,7.5 c-1.5,2.4-3.2,4.5-5.1,6.3c-2,1.8-4.2,3.3-6.6,4.4c-2.4,1.1-4.9,1.6-7.6,1.6c-2.4,0-4.5-0.4-6.4-1.1c-1.9-0.7-3.2-2-4-3.8h-0.2 l-5,17.7h63.3c3.6,0,6.6-2.9,6.6-6.6V18.2C82.6,18.1,82.3,18,82.1,17.9"/>
	<path class="logo-st0" d="M277.6,68.5h0.8c0.4,0,0.6-0.1,0.7-0.2c0.1-0.1,0.2-0.2,0.2-0.4c0-0.1,0-0.2-0.1-0.3c-0.1-0.1-0.1-0.2-0.3-0.2 c-0.1,0-0.3-0.1-0.6-0.1h-0.7V68.5z M277,70.6v-3.8h1.3c0.5,0,0.8,0,1,0.1c0.2,0.1,0.4,0.2,0.5,0.4c0.1,0.2,0.2,0.4,0.2,0.6 c0,0.3-0.1,0.5-0.3,0.7c-0.2,0.2-0.5,0.3-0.8,0.3c0.1,0.1,0.2,0.1,0.3,0.2c0.2,0.2,0.3,0.4,0.6,0.8l0.5,0.7h-0.8l-0.3-0.6 c-0.3-0.5-0.5-0.8-0.6-0.9c-0.1-0.1-0.3-0.1-0.5-0.1h-0.4v1.6H277z M278.6,65.7c-0.5,0-1,0.1-1.5,0.4c-0.5,0.3-0.8,0.6-1.1,1.1 c-0.3,0.5-0.4,1-0.4,1.5c0,0.5,0.1,1,0.4,1.5c0.3,0.5,0.6,0.8,1.1,1.1c0.5,0.3,1,0.4,1.5,0.4c0.5,0,1-0.1,1.5-0.4 c0.5-0.3,0.8-0.6,1.1-1.1c0.3-0.5,0.4-1,0.4-1.5c0-0.5-0.1-1-0.4-1.5c-0.3-0.5-0.6-0.8-1.1-1.1C279.6,65.8,279.1,65.7,278.6,65.7z M278.6,65.1c0.6,0,1.2,0.2,1.8,0.5c0.6,0.3,1,0.7,1.3,1.3c0.3,0.6,0.5,1.2,0.5,1.8c0,0.6-0.2,1.2-0.5,1.8c-0.3,0.6-0.8,1-1.3,1.3 c-0.6,0.3-1.2,0.5-1.8,0.5c-0.6,0-1.2-0.2-1.8-0.5c-0.6-0.3-1-0.8-1.3-1.3c-0.3-0.6-0.5-1.2-0.5-1.8c0-0.6,0.2-1.2,0.5-1.8 c0.3-0.6,0.8-1,1.3-1.3C277.4,65.2,278,65.1,278.6,65.1z"/>
</svg>
				<span style="color:white;font-size:.5em;text-transform:uppercase;letter-spacing:1px;">Community Edition</span>
			</a>
		</div>
		<div class="collapse navbar-collapse" id="pf-navbar">
			<ul class="nav navbar-nav">
							<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						System						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/system_advanced_admin.php" class="navlnk" >Advanced</a></li>
<li><a href="/api/" class="navlnk" >API</a></li>
<li><a href="/system_camanager.php" class="navlnk" >Cert. Manager</a></li>
<li><a href="/system.php" class="navlnk" >General Setup</a></li>
<li><a href="/system_hasync.php" class="navlnk" >High Avail. Sync</a></li>
<li><a href="/index.php?logout" class="navlnk" usepost>Logout (admin)</a></li>
<li><a href="/pkg_mgr_installed.php" class="navlnk" >Package Manager</a></li>
<li><a href="/system_gateways.php" class="navlnk" >Routing</a></li>
<li><a href="/wizard.php?xml=setup_wizard.xml" class="navlnk" >Setup Wizard</a></li>
<li><a href="/pkg_mgr_install.php?id=firmware" class="navlnk" >Update</a></li>
<li><a href="/system_usermanager.php" class="navlnk" >User Manager</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Interfaces						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/interfaces_assign.php" class="navlnk" >Assignments</a></li>
 <li class="divider"></li><li><a href="/interfaces.php?if=wan" class="navlnk" >WAN</a></li>
<li><a href="/interfaces.php?if=lan" class="navlnk" >LAN</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Firewall						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/firewall_aliases.php" class="navlnk" >Aliases</a></li>
<li><a href="/firewall_nat.php" class="navlnk" >NAT</a></li>
<li><a href="/firewall_rules.php" class="navlnk" >Rules</a></li>
<li><a href="/firewall_schedule.php" class="navlnk" >Schedules</a></li>
<li><a href="/firewall_shaper.php" class="navlnk" >Traffic Shaper</a></li>
<li><a href="/firewall_virtual_ip.php" class="navlnk" >Virtual IPs</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Services						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/acme/acme_certificates.php" class="navlnk" >Acme Certificates</a></li>
<li><a href="/services_acb.php" class="navlnk" >Auto Config Backup</a></li>
<li><a href="/services_captiveportal.php" class="navlnk" >Captive Portal</a></li>
<li><a href="/services_dhcp_relay.php" class="navlnk" >DHCP Relay</a></li>
<li><a href="/services_dhcp.php" class="navlnk" >DHCP Server</a></li>
<li><a href="/services_dhcpv6_relay.php" class="navlnk" >DHCPv6 Relay</a></li>
<li><a href="/services_dhcpv6.php" class="navlnk" >DHCPv6 Server &amp; RA</a></li>
<li><a href="/services_dnsmasq.php" class="navlnk" >DNS Forwarder</a></li>
<li><a href="/services_unbound.php" class="navlnk" >DNS Resolver</a></li>
<li><a href="/services_dyndns.php" class="navlnk" >Dynamic DNS</a></li>
<li><a href="/haproxy/haproxy_listeners.php" class="navlnk" >HAProxy</a></li>
<li><a href="/services_igmpproxy.php" class="navlnk" >IGMP Proxy</a></li>
<li><a href="/services_ntpd.php" class="navlnk" >NTP</a></li>
<li><a href="/services_pppoe.php" class="navlnk" >PPPoE Server</a></li>
<li><a href="/services_snmp.php" class="navlnk" >SNMP</a></li>
<li><a href="/pkg_edit.php?xml=miniupnpd.xml" class="navlnk" >UPnP &amp; NAT-PMP</a></li>
<li><a href="/services_wol.php" class="navlnk" >Wake-on-LAN</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						VPN						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/vpn_ipsec.php" class="navlnk" >IPsec</a></li>
<li><a href="/vpn_l2tp.php" class="navlnk" >L2TP</a></li>
<li><a href="/vpn_openvpn_server.php" class="navlnk" >OpenVPN</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Status						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/status_captiveportal.php" class="navlnk" >Captive Portal</a></li>
<li><a href="/status_carp.php" class="navlnk" >CARP (failover)</a></li>
<li><a href="/index.php" class="navlnk" >Dashboard</a></li>
<li><a href="/status_dhcp_leases.php" class="navlnk" >DHCP Leases</a></li>
<li><a href="/status_dhcpv6_leases.php" class="navlnk" >DHCPv6 Leases</a></li>
<li><a href="/status_unbound.php" class="navlnk" >DNS Resolver</a></li>
<li><a href="/status_filter_reload.php?user=true" class="navlnk" >Filter Reload</a></li>
<li><a href="/status_gateways.php" class="navlnk" >Gateways</a></li>
<li><a href="/haproxy/haproxy_stats.php?haproxystats=1" class="navlnk" >HAProxy Stats</a></li>
<li><a href="/status_interfaces.php" class="navlnk" >Interfaces</a></li>
<li><a href="/status_ipsec.php" class="navlnk" >IPsec</a></li>
<li><a href="/status_monitoring.php" class="navlnk" >Monitoring</a></li>
<li><a href="/status_ntpd.php" class="navlnk" >NTP</a></li>
<li><a href="/status_openvpn.php" class="navlnk" >OpenVPN</a></li>
<li><a href="/status_queues.php" class="navlnk" >Queues</a></li>
<li><a href="/status_services.php" class="navlnk" >Services</a></li>
<li><a href="/status_logs.php" class="navlnk" >System Logs</a></li>
<li><a href="/status_graph.php" class="navlnk" >Traffic Graph</a></li>
<li><a href="/status_upnp.php" class="navlnk" >UPnP &amp; NAT-PMP</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Diagnostics						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/diag_arp.php" class="navlnk" >ARP Table</a></li>
<li><a href="/diag_authentication.php" class="navlnk" >Authentication</a></li>
<li><a href="/diag_backup.php" class="navlnk" >Backup &amp; Restore</a></li>
<li><a href="/diag_command.php" class="navlnk" >Command Prompt</a></li>
<li><a href="/diag_dns.php" class="navlnk" >DNS Lookup</a></li>
<li><a href="/diag_edit.php" class="navlnk" >Edit File</a></li>
<li><a href="/diag_defaults.php" class="navlnk" >Factory Defaults</a></li>
<li><a href="/diag_halt.php" class="navlnk" >Halt System</a></li>
<li><a href="/diag_limiter_info.php" class="navlnk" >Limiter Info</a></li>
<li><a href="/diag_ndp.php" class="navlnk" >NDP Table</a></li>
<li><a href="/diag_packet_capture.php" class="navlnk" >Packet Capture</a></li>
<li><a href="/diag_pf_info.php" class="navlnk" >pfInfo</a></li>
<li><a href="/diag_pftop.php" class="navlnk" >pfTop</a></li>
<li><a href="/diag_ping.php" class="navlnk" >Ping</a></li>
<li><a href="/diag_reboot.php" class="navlnk" >Reboot</a></li>
<li><a href="/diag_routes.php" class="navlnk" >Routes</a></li>
<li><a href="/diag_smart.php" class="navlnk" >S.M.A.R.T. Status</a></li>
<li><a href="/diag_sockets.php" class="navlnk" >Sockets</a></li>
<li><a href="/diag_dump_states.php" class="navlnk" >States</a></li>
<li><a href="/diag_states_summary.php" class="navlnk" >States Summary</a></li>
<li><a href="/diag_system_activity.php" class="navlnk" >System Activity</a></li>
<li><a href="/diag_tables.php" class="navlnk" >Tables</a></li>
<li><a href="/diag_testport.php" class="navlnk" >Test Port</a></li>
<li><a href="/diag_traceroute.php" class="navlnk" >Traceroute</a></li>
</ul>
				</li>

				<li class="dropdown">
					<a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
						Help						<span class="caret"></span>
					</a>
					<ul class="dropdown-menu" role="menu"><li><a href="/help.php?page=system_certmanager.php" target="_blank" class="navlnk" >About this Page</a></li>
<li><a href="https://redirects.netgate.com/issues" target="_blank" class="navlnk" >Bug Database</a></li>
<li><a href="https://redirects.netgate.com/docs" target="_blank" class="navlnk" >Documentation</a></li>
<li><a href="https://redirects.netgate.com/fbsdhandbook" target="_blank" class="navlnk" >FreeBSD Handbook</a></li>
<li><a href="https://redirects.netgate.com/support" target="_blank" class="navlnk" >Paid Support</a></li>
<li><a href="https://redirects.netgate.com/book" target="_blank" class="navlnk" >pfSense Book</a></li>
<li><a href="https://redirects.netgate.com/forum" target="_blank" class="navlnk" >User Forum</a></li>
<li><a href="https://redirects.netgate.com/survey_1" target="_blank" class="navlnk" >User survey</a></li>
</ul>
				</li>

			</ul>
			<ul class="nav navbar-nav navbar-right">
									<li class="dropdown">
						<a href="/index.php?logout" usepost>
							<i class="fa fa-sign-out" title="Logout (admin@pfSense.home.arpa)"></i>
						</a>
					</li>
			</ul>
		</div>
	</div>
</nav>

<div class="container static" >

<div class="alert alert-danger"><strong>WARNING:</strong> The 'admin' account password is set to the default value.  <a href="/system_usermanager.php?act=edit&userid=0"> Change the password in the User Manager.</a></div>
	<header class="header">

<ol class="breadcrumb"><li>System</li><li><a href="/system_camanager.php">Certificate Manager</a></li><li><a href="/system_certmanager.php">Certificates</a></li></ol>		<ul class="context-links">

	
	
	
	
	
	
			<li>
			<a href="/help.php?page=system_certmanager.php" target="_blank" title="Help for items on this page">
				<i class="fa fa-question-circle"></i>
			</a>
		</li>
			</ul>
	</header>
<ul class="nav nav-pills"><li role="presentation"><a href="system_camanager.php" >CAs</a></li><li role="presentation" class="active"><a href="system_certmanager.php" >Certificates</a></li><li role="presentation"><a href="system_crlmanager.php" >Certificate Revocation</a></li></ul>	<form class="form-horizontal" method="post" action="system_certmanager.php" enctype="multipart/form-data"><input type='hidden' name='__csrf_magic' value="sid:e40d8c3404e6febd323d9a91d8f4e62e7f411129,1631021452" />
			<div class="panel panel-default">
		<div class="panel-heading">
			<h2 class="panel-title">Edit an Existing Certificate</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Method</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="method" id="method" data-toggle="collapse">
		<option value="edit">Edit an existing certificate</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Descriptive name</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control toggle-internal toggle-import toggle-edit toggle-external toggle-sign toggle-existing collapse" name="descr" id="descr" type="text" value="testcert">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Subject</span>
		</label>
			<div class="col-sm-10">
		CN=mmm.com, C=CN

		
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-sign collapse">
		<div class="panel-heading">
			<h2 class="panel-title">Sign CSR</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">CA to sign with</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="catosignwith" id="catosignwith">
		
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">CSR to sign</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csrtosign" id="csrtosign">
		<option value="new" selected>New CSR (Paste below)</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>CSR data</span>
		</label>
			<div class="col-sm-10">
			<textarea rows="5" class="form-control" name="csrpaste" id="csrpaste"></textarea>

		<span class="help-block">Paste a Certificate Signing Request in X.509 PEM format here.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Key data</span>
		</label>
			<div class="col-sm-10">
			<textarea rows="5" class="form-control" name="keypaste" id="keypaste"></textarea>

		<span class="help-block">Optionally paste a private key here. The key will be associated with the newly signed certificate in pfSense</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Certificate Lifetime (days)</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csrsign_lifetime" id="csrsign_lifetime" type="number" value="3650" max="12000" min="1" step="1">

		<span class="help-block">The length of time the signed certificate will be valid, in days. <br/>Server certificates should not have a lifetime over 398 days or some platforms may consider the certificate invalid.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Digest Algorithm</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csrsign_digest_alg" id="csrsign_digest_alg">
		<option value="sha1">sha1</option><option value="sha224">sha224</option><option value="sha256">sha256</option><option value="sha384">sha384</option><option value="sha512">sha512</option>
	</select>

		<span class="help-block">The digest method used when the certificate is signed. <br/>The best practice is to use an algorithm stronger than SHA1. Some platforms may consider weaker digest algorithms invalid</span>
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-import toggle-edit collapse">
		<div class="panel-heading">
			<h2 class="panel-title">Edit Certificate</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Certificate Type</span>
		</label>
			<div class="checkbox col-sm-5">
		<label class="chkboxlbl"><input class="import_type_toggle" name="import_type" id="import_type_x509:9f7e" type="radio" value="x509" checked="checked"> X.509 (PEM)</label>

		
	</div>	<div class="checkbox col-sm-5">
		<label class="chkboxlbl"><input class="import_type_toggle" name="import_type" id="import_type_pkcs12:9f8e" type="radio" value="pkcs12"> PKCS #12 (PFX)</label>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Certificate data</span>
		</label>
			<div class="col-sm-10">
			<textarea rows="5" class="form-control" name="cert" id="cert">-----BEGIN CERTIFICATE-----
MIIDojCCAoqgAwIBAgIIMqvJIdvMN/swDQYJKoZIhvcNAQELBQAwejELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MR8wHQYDVQQDExZLZXlN
YW5hZ2VyIFRlc3QgUlNBIENBMB4XDTIxMDgyMDA2NTY0NFoXDTIyMDgyMDA2NTY0
NFowHzELMAkGA1UEBhMCQ04xEDAOBgNVBAMTB21tbS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDZ2nyclm06xcg1uKqW2JMpAP8hjU9Zl59MQfzF
LBXJ3zEJn5x21kiRU4gRJiSSEocYfhAccf3L3mF7kro8aJjz6xem4OXvvP5CVHtU
9S86ZFzhnP/X+MT+3jkQhB8XhS5b+iw7B1FVhUsw5xRu1ZW+Bgc4uhFAFlbIYn5r
baLjMf778qsYtNrKHMBPSfORpoZqkSVZyOmEEbyy3dc/4NmIHTQQCSC1Y7bHb+LE
CaNllICmKwUfNdEexXB8MTJ0jumkUvArHD3FzB0kdtkBkfMIQcHrDpQQUHf8ePLH
4edm1nL//xhP3fUkXcwBEd237xtloLD1MT3H75jOAW4Dss4fAgMBAAGjgYYwgYMw
DgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAd
BgNVHQ4EFgQUmiN70WHD2SXhRQsh+p1nuiaVc4UwHwYDVR0jBBgwFoAUECM+bgg+
elE+HXbkMaIApFqFmZIwEgYDVR0RBAswCYIHbW1tLmNvbTANBgkqhkiG9w0BAQsF
AAOCAQEAQiv38MYOj0XQ1Qst8I6zNLaLGc7Vi3wOEjyTdXFP6F+cUQLNQp5JgZf9
3sa6L73Fh71HXTxrMyX2q3xZHWyb9uKqMJy/5z3NCZO8uVtHgUhu21JhkxV1z3VT
WV9fKSn5WyR6P/RUk0sGVrQ0F+E6rJTC9689E1Dbe8ocdp8b9HopeV2OcmMKQIZN
gYQlNi8lSCN0DWyB98ktTbMrbnomPutU6kNzFSoYkg2kmHMaOK84nVPMgX32+Vb8
IdkWsJSNOR2wHFUEos5hGLsJgg+m/+4RHL6jzfOiW8PWISnyuv+eityKrGI5Tjxy
a5KLqnLjWk1Rvu/LeJUsa7/laaH2Sw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIID2jCCAsKgAwIBAgIIeSI3M4rCCEswDQYJKoZIhvcNAQELBQAwezELMAkGA1UE
BhMCQ04xFzAVBgNVBAoTDktleU1hbmFnZXIub3JnMTEwLwYDVQQLEyhLZXlNYW5h
Z2VyIFRlc3QgUm9vdCAtIEZvciBUZXN0IFVzZSBPbmx5MSAwHgYDVQQDExdLZXlN
YW5hZ2VyIFRlc3QgUm9vdCBDQTAeFw0yMTA3MTQwNTEyMDBaFw0zMTA3MTQwNTEy
MDBaMHoxCzAJBgNVBAYTAkNOMRcwFQYDVQQKEw5LZXlNYW5hZ2VyLm9yZzExMC8G
A1UECxMoS2V5TWFuYWdlciBUZXN0IFJvb3QgLSBGb3IgVGVzdCBVc2UgT25seTEf
MB0GA1UEAxMWS2V5TWFuYWdlciBUZXN0IFJTQSBDQTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBANgs0N+IrziLtph3gDHpapH7Wn4moycKDi9ymb5FHBpj
Gs2TRIn7uJMhFbAJklcdN9usbXjgWkmP2oFdfTJsQQyvF8KcNpSL8OMSS8zy79sU
jV+VvW0w0Uv43lBrPVkXF2c2AhkguWFb5DjunFxBCY5PnEVUP3wYBOsJW6HCvUx8
tuICkfBudOvf612YtEixA2GAig6kviTDrxBbN2QsrrnZEyaAnR6+rDAbLzmVK+Px
a6JxbTQlC+qrjDV/o09XnbLWtVE8U67D4IhYfT0HIsugTRGUIkff49CQrAEZ0/pi
RWln7KSoM+rDeDp08ORWZsQ7Mx5G//Zep6evhQuCnWcCAwEAAaNjMGEwDgYDVR0P
AQH/BAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFBAjPm4IPnpRPh12
5DGiAKRahZmSMB8GA1UdIwQYMBaAFPKcEGuzWGA4xJynrNCqlL6+wFpQMA0GCSqG
SIb3DQEBCwUAA4IBAQA7r6WhiJerQweBsL+sUaPti1O5cR30lpw7YxhGERQsik7h
pRhc9tE+PWGzepBhx1tN9pq8N9lS+Mbcx2oOJOc8e1RPd3q3meU76868OuTSKtD+
mV3pcJ+rnbWr2pD1FWu5GDn1/5cmpNXotha+pWIZpGJ8lVtrhJCwmH9hUFFkxKr/
dXsis03TzIcbH5fyiJJOIXf77IunzFIgmwA4wBLOD+aRPP1wMa2Q/dagbvzRbx5g
sFw0xuJyZInrpC7czwz8AK2NLCbzeMLQL5HvQu71AhqysXQDBFZjdAQ2IEC8ZMsq
ew6BXkqiZDA5jg05jQSGydiIedO1ScXfyd8nMZls
-----END CERTIFICATE-----
</textarea>

		<span class="help-block">Paste a certificate in X.509 PEM format here.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Private key data</span>
		</label>
			<div class="col-sm-10">
			<textarea rows="5" class="form-control" name="key" id="key">-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA2dp8nJZtOsXINbiqltiTKQD/IY1PWZefTEH8xSwVyd8xCZ+c
dtZIkVOIESYkkhKHGH4QHHH9y95he5K6PGiY8+sXpuDl77z+QlR7VPUvOmRc4Zz/
1/jE/t45EIQfF4UuW/osOwdRVYVLMOcUbtWVvgYHOLoRQBZWyGJ+a22i4zH++/Kr
GLTayhzAT0nzkaaGapElWcjphBG8st3XP+DZiB00EAkgtWO2x2/ixAmjZZSApisF
HzXRHsVwfDEydI7ppFLwKxw9xcwdJHbZAZHzCEHB6w6UEFB3/Hjyx+HnZtZy//8Y
T931JF3MARHdt+8bZaCw9TE9x++YzgFuA7LOHwIDAQABAoIBAD/MzbuqDjktHRIm
j8b3jDlw8kboNHnffqZ9mMJBw+vH8nuIA/GFleEBnpKvIfypcmkI2j0KYTJoYRIo
iWQRmeGtUtLrgEtyhN/2D9x0pa0rIUxthzu/vimJ+RpOJzDjLw1+uZ9b6ETscXXT
5tcCtATfjRPe4hhrsmSi+7UIebChNCKbQ8pv3fDDbDjnRBZiNAAlfTRJkBnba8W7
c6r61Jpe/bDKxpSphOPm+gGIVN88oNXv1xJwA7Qx0YoNPzFvm2kIjZ9LLDs9p10h
P/TolNGlKv6ac/q7qKEPRJw4UqAXvh2Y6jon62jHcdIwIcutxmlH4i18LECPiKa9
BOj9GXECgYEA+rRgR05eScB+xBevbm5msH/MFXIAD5lq0v9AkL12WHrfxk50XZnF
NS/E8vDOUFvVw/rzTbYyqV49jqdHb4ZuD4lCvBayCmmEMe56aaaTNri/ajEr4E5x
wwZyzRxXlDS+2ZOkrlHM0HLtsNahhmH1aOCyD1IrnZ864caLsGuE3CkCgYEA3nR5
6K7MoZhe+cxzi8V9xegVD1kiu27oFpiPvsLHjpLvIfiwNQlzgoU/2+eHKPLPXmZ2
rDBnY9MTxtY/8xlA3Z5TMUR2NhpFkC46eENkgnkf3yAM19WrLTVokvhXsfwUtymz
qYjgKG6MfsDvMAfb+ogMot9RvKO9g+1IDHMjoQcCgYAvoEiSAz9CP4FVezJmhi6X
5Q8+G7QLQpfakYcQeA2dbWpJX+oXRfkCy5pclIZ9GZUYb/n8j1o8dpy3FuwpMZ6C
8Q5ucNlNxRHJ8oXqwCxDPwGOCN1O9VgDNpxkfrfcfdCrwLKOMxf3mX2yFHQG9WEL
lXP+GRwUC4XCElfDIgnRUQKBgDK4vh821AO4eVddra7d7eqVG1Avk8LG6/ZS/NuT
D+tLR2koigzdxc+p0EC0ztWgX3X3yPFD7B8Pvr+klFo6lNazRebC5G07mkbgs4Y+
X4l8Uq8OYL9JwckCF4EDTQORJawJvyRVyD6PzksMdL0v3ZGHOdJdNwbbEtgk3zuv
eR07AoGANmcoQNUTkgczg9gKfysjCt2tCCR0Mk+rMiYE3q9LFUZT8uKue3EIXHPW
ATXb2bGle5s2ArSVtwGAjh3Cr8nlSpP+SlCxlBTKP/hrWsW5U+iIYRZiyAVF3KhW
3oLrXY39wqXrdry8TwbfX8PoYMrQVNtUM9zz5in+6pe5Myv7C7E=
-----END RSA PRIVATE KEY-----
</textarea>

		<span class="help-block">Paste a private key in X.509 PEM format here. This field may remain empty in certain cases, such as when the private key is stored on a PKCS#11 token.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>PKCS #12 certificate</span>
		</label>
			<div class="col-sm-10">
		<input name="pkcs12_cert" id="pkcs12_cert" type="file">

		<span class="help-block">Select a PKCS #12 certificate store.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>PKCS #12 certificate password</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="pkcs12_pass" id="pkcs12_pass" type="password">

		<span class="help-block">Enter the password to unlock the PKCS #12 certificate store.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Intermediates</span>
		</label>
			<div class="checkbox col-sm-10">
		<label class="chkboxlbl"><input name="pkcs12_intermediate" id="pkcs12_intermediate" type="checkbox" value="yes"> Import intermediate CAs</label>

		<span class="help-block">Import any intermediate certificate authorities found in the PKCS #12 certificate store.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Export Password</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control toggle-edit collapse" name="exportpass" id="exportpass" type="password" placeholder="Export Password" autocomplete="new-password">

		<span class="help-block">Enter the password to use when using the export buttons below (not stored)</span>
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-internal collapse">
		<div class="panel-heading">
			<h2 class="panel-title">Internal Certificate</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Certificate authority</span>
		</label>
			<div class="col-sm-10">
		No internal Certificate Authorities have been defined. An internal CA must be defined in order to create an internal certificate. <a href="system_camanager.php?act=new&amp;method=internal"> Create</a> an internal CA.

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Key type</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="keytype" id="keytype">
		<option value="RSA">RSA</option><option value="ECDSA">ECDSA</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group rsakeys">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="keylen" id="keylen">
		<option value="1024">1024</option><option value="2048">2048</option><option value="3072">3072</option><option value="4096">4096</option><option value="6144">6144</option><option value="7680">7680</option><option value="8192">8192</option><option value="15360">15360</option><option value="16384">16384</option>
	</select>

		<span class="help-block">The length to use when generating a new RSA key, in bits. <br/>The Key Length should not be lower than 2048 or some platforms may consider the certificate invalid.</span>
	</div>
		
	</div>	<div class="form-group ecnames">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="ecname" id="ecname">
		<option value="secp112r1">secp112r1</option><option value="secp112r2">secp112r2</option><option value="secp128r1">secp128r1</option><option value="secp128r2">secp128r2</option><option value="secp160k1">secp160k1</option><option value="secp160r1">secp160r1</option><option value="secp160r2">secp160r2</option><option value="secp192k1">secp192k1</option><option value="secp224k1">secp224k1</option><option value="secp224r1">secp224r1</option><option value="secp256k1">secp256k1</option><option value="secp384r1">secp384r1 [HTTPS] [IPsec] [OpenVPN]</option><option value="secp521r1">secp521r1 [IPsec] [OpenVPN]</option><option value="prime192v1">prime192v1</option><option value="prime192v2">prime192v2</option><option value="prime192v3">prime192v3</option><option value="prime239v1">prime239v1</option><option value="prime239v2">prime239v2</option><option value="prime239v3">prime239v3</option><option value="prime256v1">prime256v1 [HTTPS] [IPsec] [OpenVPN]</option><option value="sect113r1">sect113r1</option><option value="sect113r2">sect113r2</option><option value="sect131r1">sect131r1</option><option value="sect131r2">sect131r2</option><option value="sect163k1">sect163k1</option><option value="sect163r1">sect163r1</option><option value="sect163r2">sect163r2</option><option value="sect193r1">sect193r1</option><option value="sect193r2">sect193r2</option><option value="sect233k1">sect233k1</option><option value="sect233r1">sect233r1</option><option value="sect239k1">sect239k1</option><option value="sect283k1">sect283k1</option><option value="sect283r1">sect283r1</option><option value="sect409k1">sect409k1</option><option value="sect409r1">sect409r1</option><option value="sect571k1">sect571k1</option><option value="sect571r1">sect571r1</option><option value="c2pnb163v1">c2pnb163v1</option><option value="c2pnb163v2">c2pnb163v2</option><option value="c2pnb163v3">c2pnb163v3</option><option value="c2pnb176v1">c2pnb176v1</option><option value="c2tnb191v1">c2tnb191v1</option><option value="c2tnb191v2">c2tnb191v2</option><option value="c2tnb191v3">c2tnb191v3</option><option value="c2pnb208w1">c2pnb208w1</option><option value="c2tnb239v1">c2tnb239v1</option><option value="c2tnb239v2">c2tnb239v2</option><option value="c2tnb239v3">c2tnb239v3</option><option value="c2pnb272w1">c2pnb272w1</option><option value="c2pnb304w1">c2pnb304w1</option><option value="c2tnb359v1">c2tnb359v1</option><option value="c2pnb368w1">c2pnb368w1</option><option value="c2tnb431r1">c2tnb431r1</option><option value="wap-wsg-idm-ecid-wtls1">wap-wsg-idm-ecid-wtls1</option><option value="wap-wsg-idm-ecid-wtls3">wap-wsg-idm-ecid-wtls3</option><option value="wap-wsg-idm-ecid-wtls4">wap-wsg-idm-ecid-wtls4</option><option value="wap-wsg-idm-ecid-wtls5">wap-wsg-idm-ecid-wtls5</option><option value="wap-wsg-idm-ecid-wtls6">wap-wsg-idm-ecid-wtls6</option><option value="wap-wsg-idm-ecid-wtls7">wap-wsg-idm-ecid-wtls7</option><option value="wap-wsg-idm-ecid-wtls8">wap-wsg-idm-ecid-wtls8</option><option value="wap-wsg-idm-ecid-wtls9">wap-wsg-idm-ecid-wtls9</option><option value="wap-wsg-idm-ecid-wtls10">wap-wsg-idm-ecid-wtls10</option><option value="wap-wsg-idm-ecid-wtls11">wap-wsg-idm-ecid-wtls11</option><option value="wap-wsg-idm-ecid-wtls12">wap-wsg-idm-ecid-wtls12</option><option value="Oakley-EC2N-3">Oakley-EC2N-3</option><option value="Oakley-EC2N-4">Oakley-EC2N-4</option><option value="brainpoolP160r1">brainpoolP160r1</option><option value="brainpoolP160t1">brainpoolP160t1</option><option value="brainpoolP192r1">brainpoolP192r1</option><option value="brainpoolP192t1">brainpoolP192t1</option><option value="brainpoolP224r1">brainpoolP224r1</option><option value="brainpoolP224t1">brainpoolP224t1</option><option value="brainpoolP256r1">brainpoolP256r1</option><option value="brainpoolP256t1">brainpoolP256t1</option><option value="brainpoolP320r1">brainpoolP320r1</option><option value="brainpoolP320t1">brainpoolP320t1</option><option value="brainpoolP384r1">brainpoolP384r1</option><option value="brainpoolP384t1">brainpoolP384t1</option><option value="brainpoolP512r1">brainpoolP512r1</option><option value="brainpoolP512t1">brainpoolP512t1</option><option value="SM2">SM2</option>
	</select>

		<span class="help-block">Curves may not be compatible with all uses. Known compatible curve uses are denoted in brackets.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Digest Algorithm</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="digest_alg" id="digest_alg">
		<option value="sha1">sha1</option><option value="sha224">sha224</option><option value="sha256">sha256</option><option value="sha384">sha384</option><option value="sha512">sha512</option>
	</select>

		<span class="help-block">The digest method used when the certificate is signed. <br/>The best practice is to use an algorithm stronger than SHA1. Some platforms may consider weaker digest algorithms invalid</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Lifetime (days)</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="lifetime" id="lifetime" type="number" max="12000" min="1" step="1">

		<span class="help-block">The length of time the signed certificate will be valid, in days. <br/>Server certificates should not have a lifetime over 398 days or some platforms may consider the certificate invalid.</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Common Name</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="dn_commonname" id="dn_commonname" type="text" placeholder="e.g. www.example.com">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
		The following certificate subject components are optional and may be left blank.

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Country Code</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="dn_country" id="dn_country">
		<option value="" selected>None</option><option value="US">US</option><option value="CA">CA</option><option value="AD">AD</option><option value="AE">AE</option><option value="AF">AF</option><option value="AG">AG</option><option value="AI">AI</option><option value="AL">AL</option><option value="AM">AM</option><option value="AN">AN</option><option value="AO">AO</option><option value="AQ">AQ</option><option value="AR">AR</option><option value="AS">AS</option><option value="AT">AT</option><option value="AU">AU</option><option value="AW">AW</option><option value="AX">AX</option><option value="AZ">AZ</option><option value="BA">BA</option><option value="BB">BB</option><option value="BD">BD</option><option value="BE">BE</option><option value="BF">BF</option><option value="BG">BG</option><option value="BH">BH</option><option value="BI">BI</option><option value="BJ">BJ</option><option value="BL">BL</option><option value="BM">BM</option><option value="BN">BN</option><option value="BO">BO</option><option value="BR">BR</option><option value="BS">BS</option><option value="BT">BT</option><option value="BV">BV</option><option value="BW">BW</option><option value="BY">BY</option><option value="BZ">BZ</option><option value="CC">CC</option><option value="CD">CD</option><option value="CF">CF</option><option value="CG">CG</option><option value="CH">CH</option><option value="CI">CI</option><option value="CK">CK</option><option value="CL">CL</option><option value="CM">CM</option><option value="CN">CN</option><option value="CO">CO</option><option value="CR">CR</option><option value="CU">CU</option><option value="CV">CV</option><option value="CX">CX</option><option value="CY">CY</option><option value="CZ">CZ</option><option value="DE">DE</option><option value="DJ">DJ</option><option value="DK">DK</option><option value="DM">DM</option><option value="DO">DO</option><option value="DZ">DZ</option><option value="EC">EC</option><option value="EE">EE</option><option value="EG">EG</option><option value="EH">EH</option><option value="ER">ER</option><option value="ES">ES</option><option value="ET">ET</option><option value="FI">FI</option><option value="FJ">FJ</option><option value="FK">FK</option><option value="FM">FM</option><option value="FO">FO</option><option value="FR">FR</option><option value="GA">GA</option><option value="GB">GB</option><option value="GD">GD</option><option value="GE">GE</option><option value="GF">GF</option><option value="GG">GG</option><option value="GH">GH</option><option value="GI">GI</option><option value="GL">GL</option><option value="GM">GM</option><option value="GN">GN</option><option value="GP">GP</option><option value="GQ">GQ</option><option value="GR">GR</option><option value="GS">GS</option><option value="GT">GT</option><option value="GU">GU</option><option value="GW">GW</option><option value="GY">GY</option><option value="HK">HK</option><option value="HM">HM</option><option value="HN">HN</option><option value="HR">HR</option><option value="HT">HT</option><option value="HU">HU</option><option value="ID">ID</option><option value="IE">IE</option><option value="IL">IL</option><option value="IM">IM</option><option value="IN">IN</option><option value="IO">IO</option><option value="IQ">IQ</option><option value="IR">IR</option><option value="IS">IS</option><option value="IT">IT</option><option value="JE">JE</option><option value="JM">JM</option><option value="JO">JO</option><option value="JP">JP</option><option value="KE">KE</option><option value="KG">KG</option><option value="KH">KH</option><option value="KI">KI</option><option value="KM">KM</option><option value="KN">KN</option><option value="KP">KP</option><option value="KR">KR</option><option value="KW">KW</option><option value="KY">KY</option><option value="KZ">KZ</option><option value="LA">LA</option><option value="LB">LB</option><option value="LC">LC</option><option value="LI">LI</option><option value="LK">LK</option><option value="LR">LR</option><option value="LS">LS</option><option value="LT">LT</option><option value="LU">LU</option><option value="LV">LV</option><option value="LY">LY</option><option value="MA">MA</option><option value="MC">MC</option><option value="MD">MD</option><option value="ME">ME</option><option value="MF">MF</option><option value="MG">MG</option><option value="MH">MH</option><option value="MK">MK</option><option value="ML">ML</option><option value="MM">MM</option><option value="MN">MN</option><option value="MO">MO</option><option value="MP">MP</option><option value="MQ">MQ</option><option value="MR">MR</option><option value="MS">MS</option><option value="MT">MT</option><option value="MU">MU</option><option value="MV">MV</option><option value="MW">MW</option><option value="MX">MX</option><option value="MY">MY</option><option value="MZ">MZ</option><option value="NA">NA</option><option value="NC">NC</option><option value="NE">NE</option><option value="NF">NF</option><option value="NG">NG</option><option value="NI">NI</option><option value="NL">NL</option><option value="NO">NO</option><option value="NP">NP</option><option value="NR">NR</option><option value="NU">NU</option><option value="NZ">NZ</option><option value="OM">OM</option><option value="PA">PA</option><option value="PE">PE</option><option value="PF">PF</option><option value="PG">PG</option><option value="PH">PH</option><option value="PK">PK</option><option value="PL">PL</option><option value="PM">PM</option><option value="PN">PN</option><option value="PR">PR</option><option value="PS">PS</option><option value="PT">PT</option><option value="PW">PW</option><option value="PY">PY</option><option value="QA">QA</option><option value="RE">RE</option><option value="RO">RO</option><option value="RS">RS</option><option value="RU">RU</option><option value="RW">RW</option><option value="SA">SA</option><option value="SB">SB</option><option value="SC">SC</option><option value="SD">SD</option><option value="SE">SE</option><option value="SG">SG</option><option value="SH">SH</option><option value="SI">SI</option><option value="SJ">SJ</option><option value="SK">SK</option><option value="SL">SL</option><option value="SM">SM</option><option value="SN">SN</option><option value="SO">SO</option><option value="SR">SR</option><option value="ST">ST</option><option value="SV">SV</option><option value="SY">SY</option><option value="SZ">SZ</option><option value="TC">TC</option><option value="TD">TD</option><option value="TF">TF</option><option value="TG">TG</option><option value="TH">TH</option><option value="TJ">TJ</option><option value="TK">TK</option><option value="TL">TL</option><option value="TM">TM</option><option value="TN">TN</option><option value="TO">TO</option><option value="TR">TR</option><option value="TT">TT</option><option value="TV">TV</option><option value="TW">TW</option><option value="TZ">TZ</option><option value="UA">UA</option><option value="UG">UG</option><option value="UM">UM</option><option value="UY">UY</option><option value="UZ">UZ</option><option value="VA">VA</option><option value="VC">VC</option><option value="VE">VE</option><option value="VG">VG</option><option value="VI">VI</option><option value="VN">VN</option><option value="VU">VU</option><option value="WF">WF</option><option value="WS">WS</option><option value="YE">YE</option><option value="YT">YT</option><option value="ZA">ZA</option><option value="ZM">ZM</option><option value="ZW">ZW</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>State or Province</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="dn_state" id="dn_state" type="text" placeholder="e.g. Texas">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>City</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="dn_city" id="dn_city" type="text" placeholder="e.g. Austin">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Organization</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="dn_organization" id="dn_organization" type="text" placeholder="e.g. My Company Inc">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Organizational Unit</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="dn_organizationalunit" id="dn_organizationalunit" type="text" placeholder="e.g. My Department Name (optional)">

		
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-external collapse">
		<div class="panel-heading">
			<h2 class="panel-title">External Signing Request</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Key type</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csr_keytype" id="csr_keytype">
		<option value="RSA">RSA</option><option value="ECDSA">ECDSA</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group csr_rsakeys">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csr_keylen" id="csr_keylen">
		<option value="1024">1024</option><option value="2048">2048</option><option value="3072">3072</option><option value="4096">4096</option><option value="6144">6144</option><option value="7680">7680</option><option value="8192">8192</option><option value="15360">15360</option><option value="16384">16384</option>
	</select>

		<span class="help-block">The length to use when generating a new RSA key, in bits. <br/>The Key Length should not be lower than 2048 or some platforms may consider the certificate invalid.</span>
	</div>
		
	</div>	<div class="form-group csr_ecnames">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csr_ecname" id="csr_ecname">
		<option value="secp112r1">secp112r1</option><option value="secp112r2">secp112r2</option><option value="secp128r1">secp128r1</option><option value="secp128r2">secp128r2</option><option value="secp160k1">secp160k1</option><option value="secp160r1">secp160r1</option><option value="secp160r2">secp160r2</option><option value="secp192k1">secp192k1</option><option value="secp224k1">secp224k1</option><option value="secp224r1">secp224r1</option><option value="secp256k1">secp256k1</option><option value="secp384r1">secp384r1 [HTTPS] [IPsec] [OpenVPN]</option><option value="secp521r1">secp521r1 [IPsec] [OpenVPN]</option><option value="prime192v1">prime192v1</option><option value="prime192v2">prime192v2</option><option value="prime192v3">prime192v3</option><option value="prime239v1">prime239v1</option><option value="prime239v2">prime239v2</option><option value="prime239v3">prime239v3</option><option value="prime256v1">prime256v1 [HTTPS] [IPsec] [OpenVPN]</option><option value="sect113r1">sect113r1</option><option value="sect113r2">sect113r2</option><option value="sect131r1">sect131r1</option><option value="sect131r2">sect131r2</option><option value="sect163k1">sect163k1</option><option value="sect163r1">sect163r1</option><option value="sect163r2">sect163r2</option><option value="sect193r1">sect193r1</option><option value="sect193r2">sect193r2</option><option value="sect233k1">sect233k1</option><option value="sect233r1">sect233r1</option><option value="sect239k1">sect239k1</option><option value="sect283k1">sect283k1</option><option value="sect283r1">sect283r1</option><option value="sect409k1">sect409k1</option><option value="sect409r1">sect409r1</option><option value="sect571k1">sect571k1</option><option value="sect571r1">sect571r1</option><option value="c2pnb163v1">c2pnb163v1</option><option value="c2pnb163v2">c2pnb163v2</option><option value="c2pnb163v3">c2pnb163v3</option><option value="c2pnb176v1">c2pnb176v1</option><option value="c2tnb191v1">c2tnb191v1</option><option value="c2tnb191v2">c2tnb191v2</option><option value="c2tnb191v3">c2tnb191v3</option><option value="c2pnb208w1">c2pnb208w1</option><option value="c2tnb239v1">c2tnb239v1</option><option value="c2tnb239v2">c2tnb239v2</option><option value="c2tnb239v3">c2tnb239v3</option><option value="c2pnb272w1">c2pnb272w1</option><option value="c2pnb304w1">c2pnb304w1</option><option value="c2tnb359v1">c2tnb359v1</option><option value="c2pnb368w1">c2pnb368w1</option><option value="c2tnb431r1">c2tnb431r1</option><option value="wap-wsg-idm-ecid-wtls1">wap-wsg-idm-ecid-wtls1</option><option value="wap-wsg-idm-ecid-wtls3">wap-wsg-idm-ecid-wtls3</option><option value="wap-wsg-idm-ecid-wtls4">wap-wsg-idm-ecid-wtls4</option><option value="wap-wsg-idm-ecid-wtls5">wap-wsg-idm-ecid-wtls5</option><option value="wap-wsg-idm-ecid-wtls6">wap-wsg-idm-ecid-wtls6</option><option value="wap-wsg-idm-ecid-wtls7">wap-wsg-idm-ecid-wtls7</option><option value="wap-wsg-idm-ecid-wtls8">wap-wsg-idm-ecid-wtls8</option><option value="wap-wsg-idm-ecid-wtls9">wap-wsg-idm-ecid-wtls9</option><option value="wap-wsg-idm-ecid-wtls10">wap-wsg-idm-ecid-wtls10</option><option value="wap-wsg-idm-ecid-wtls11">wap-wsg-idm-ecid-wtls11</option><option value="wap-wsg-idm-ecid-wtls12">wap-wsg-idm-ecid-wtls12</option><option value="Oakley-EC2N-3">Oakley-EC2N-3</option><option value="Oakley-EC2N-4">Oakley-EC2N-4</option><option value="brainpoolP160r1">brainpoolP160r1</option><option value="brainpoolP160t1">brainpoolP160t1</option><option value="brainpoolP192r1">brainpoolP192r1</option><option value="brainpoolP192t1">brainpoolP192t1</option><option value="brainpoolP224r1">brainpoolP224r1</option><option value="brainpoolP224t1">brainpoolP224t1</option><option value="brainpoolP256r1">brainpoolP256r1</option><option value="brainpoolP256t1">brainpoolP256t1</option><option value="brainpoolP320r1">brainpoolP320r1</option><option value="brainpoolP320t1">brainpoolP320t1</option><option value="brainpoolP384r1">brainpoolP384r1</option><option value="brainpoolP384t1">brainpoolP384t1</option><option value="brainpoolP512r1">brainpoolP512r1</option><option value="brainpoolP512t1">brainpoolP512t1</option><option value="SM2">SM2</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Digest Algorithm</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csr_digest_alg" id="csr_digest_alg">
		<option value="sha1">sha1</option><option value="sha224">sha224</option><option value="sha256">sha256</option><option value="sha384">sha384</option><option value="sha512">sha512</option>
	</select>

		<span class="help-block">The digest method used when the certificate is signed. <br/>The best practice is to use an algorithm stronger than SHA1. Some platforms may consider weaker digest algorithms invalid</span>
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Common Name</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csr_dn_commonname" id="csr_dn_commonname" type="text" placeholder="e.g. internal-ca">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			
		</label>
			<div class="col-sm-10">
		The following certificate subject components are optional and may be left blank.

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Country Code</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="csr_dn_country" id="csr_dn_country">
		<option value="" selected>None</option><option value="US">US</option><option value="CA">CA</option><option value="AD">AD</option><option value="AE">AE</option><option value="AF">AF</option><option value="AG">AG</option><option value="AI">AI</option><option value="AL">AL</option><option value="AM">AM</option><option value="AN">AN</option><option value="AO">AO</option><option value="AQ">AQ</option><option value="AR">AR</option><option value="AS">AS</option><option value="AT">AT</option><option value="AU">AU</option><option value="AW">AW</option><option value="AX">AX</option><option value="AZ">AZ</option><option value="BA">BA</option><option value="BB">BB</option><option value="BD">BD</option><option value="BE">BE</option><option value="BF">BF</option><option value="BG">BG</option><option value="BH">BH</option><option value="BI">BI</option><option value="BJ">BJ</option><option value="BL">BL</option><option value="BM">BM</option><option value="BN">BN</option><option value="BO">BO</option><option value="BR">BR</option><option value="BS">BS</option><option value="BT">BT</option><option value="BV">BV</option><option value="BW">BW</option><option value="BY">BY</option><option value="BZ">BZ</option><option value="CC">CC</option><option value="CD">CD</option><option value="CF">CF</option><option value="CG">CG</option><option value="CH">CH</option><option value="CI">CI</option><option value="CK">CK</option><option value="CL">CL</option><option value="CM">CM</option><option value="CN">CN</option><option value="CO">CO</option><option value="CR">CR</option><option value="CU">CU</option><option value="CV">CV</option><option value="CX">CX</option><option value="CY">CY</option><option value="CZ">CZ</option><option value="DE">DE</option><option value="DJ">DJ</option><option value="DK">DK</option><option value="DM">DM</option><option value="DO">DO</option><option value="DZ">DZ</option><option value="EC">EC</option><option value="EE">EE</option><option value="EG">EG</option><option value="EH">EH</option><option value="ER">ER</option><option value="ES">ES</option><option value="ET">ET</option><option value="FI">FI</option><option value="FJ">FJ</option><option value="FK">FK</option><option value="FM">FM</option><option value="FO">FO</option><option value="FR">FR</option><option value="GA">GA</option><option value="GB">GB</option><option value="GD">GD</option><option value="GE">GE</option><option value="GF">GF</option><option value="GG">GG</option><option value="GH">GH</option><option value="GI">GI</option><option value="GL">GL</option><option value="GM">GM</option><option value="GN">GN</option><option value="GP">GP</option><option value="GQ">GQ</option><option value="GR">GR</option><option value="GS">GS</option><option value="GT">GT</option><option value="GU">GU</option><option value="GW">GW</option><option value="GY">GY</option><option value="HK">HK</option><option value="HM">HM</option><option value="HN">HN</option><option value="HR">HR</option><option value="HT">HT</option><option value="HU">HU</option><option value="ID">ID</option><option value="IE">IE</option><option value="IL">IL</option><option value="IM">IM</option><option value="IN">IN</option><option value="IO">IO</option><option value="IQ">IQ</option><option value="IR">IR</option><option value="IS">IS</option><option value="IT">IT</option><option value="JE">JE</option><option value="JM">JM</option><option value="JO">JO</option><option value="JP">JP</option><option value="KE">KE</option><option value="KG">KG</option><option value="KH">KH</option><option value="KI">KI</option><option value="KM">KM</option><option value="KN">KN</option><option value="KP">KP</option><option value="KR">KR</option><option value="KW">KW</option><option value="KY">KY</option><option value="KZ">KZ</option><option value="LA">LA</option><option value="LB">LB</option><option value="LC">LC</option><option value="LI">LI</option><option value="LK">LK</option><option value="LR">LR</option><option value="LS">LS</option><option value="LT">LT</option><option value="LU">LU</option><option value="LV">LV</option><option value="LY">LY</option><option value="MA">MA</option><option value="MC">MC</option><option value="MD">MD</option><option value="ME">ME</option><option value="MF">MF</option><option value="MG">MG</option><option value="MH">MH</option><option value="MK">MK</option><option value="ML">ML</option><option value="MM">MM</option><option value="MN">MN</option><option value="MO">MO</option><option value="MP">MP</option><option value="MQ">MQ</option><option value="MR">MR</option><option value="MS">MS</option><option value="MT">MT</option><option value="MU">MU</option><option value="MV">MV</option><option value="MW">MW</option><option value="MX">MX</option><option value="MY">MY</option><option value="MZ">MZ</option><option value="NA">NA</option><option value="NC">NC</option><option value="NE">NE</option><option value="NF">NF</option><option value="NG">NG</option><option value="NI">NI</option><option value="NL">NL</option><option value="NO">NO</option><option value="NP">NP</option><option value="NR">NR</option><option value="NU">NU</option><option value="NZ">NZ</option><option value="OM">OM</option><option value="PA">PA</option><option value="PE">PE</option><option value="PF">PF</option><option value="PG">PG</option><option value="PH">PH</option><option value="PK">PK</option><option value="PL">PL</option><option value="PM">PM</option><option value="PN">PN</option><option value="PR">PR</option><option value="PS">PS</option><option value="PT">PT</option><option value="PW">PW</option><option value="PY">PY</option><option value="QA">QA</option><option value="RE">RE</option><option value="RO">RO</option><option value="RS">RS</option><option value="RU">RU</option><option value="RW">RW</option><option value="SA">SA</option><option value="SB">SB</option><option value="SC">SC</option><option value="SD">SD</option><option value="SE">SE</option><option value="SG">SG</option><option value="SH">SH</option><option value="SI">SI</option><option value="SJ">SJ</option><option value="SK">SK</option><option value="SL">SL</option><option value="SM">SM</option><option value="SN">SN</option><option value="SO">SO</option><option value="SR">SR</option><option value="ST">ST</option><option value="SV">SV</option><option value="SY">SY</option><option value="SZ">SZ</option><option value="TC">TC</option><option value="TD">TD</option><option value="TF">TF</option><option value="TG">TG</option><option value="TH">TH</option><option value="TJ">TJ</option><option value="TK">TK</option><option value="TL">TL</option><option value="TM">TM</option><option value="TN">TN</option><option value="TO">TO</option><option value="TR">TR</option><option value="TT">TT</option><option value="TV">TV</option><option value="TW">TW</option><option value="TZ">TZ</option><option value="UA">UA</option><option value="UG">UG</option><option value="UM">UM</option><option value="UY">UY</option><option value="UZ">UZ</option><option value="VA">VA</option><option value="VC">VC</option><option value="VE">VE</option><option value="VG">VG</option><option value="VI">VI</option><option value="VN">VN</option><option value="VU">VU</option><option value="WF">WF</option><option value="WS">WS</option><option value="YE">YE</option><option value="YT">YT</option><option value="ZA">ZA</option><option value="ZM">ZM</option><option value="ZW">ZW</option>
	</select>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>State or Province</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csr_dn_state" id="csr_dn_state" type="text" placeholder="e.g. Texas">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>City</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csr_dn_city" id="csr_dn_city" type="text" placeholder="e.g. Austin">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Organization</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csr_dn_organization" id="csr_dn_organization" type="text" placeholder="e.g. My Company Inc">

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Organizational Unit</span>
		</label>
			<div class="col-sm-10">
		<input class="form-control" name="csr_dn_organizationalunit" id="csr_dn_organizationalunit" type="text" placeholder="e.g. My Department Name (optional)">

		
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-existing collapse">
		<div class="panel-heading">
			<h2 class="panel-title">Choose an Existing Certificate</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Existing Certificates</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="certref" id="certref">
		<option value="61334dfdecccc">web</option><option value="613356ef7ac58">test.abc</option><option value="6136cfbfdca6d">testcert (In Use)</option>
	</select>

		
	</div>
		
	</div>
		</div>
	</div>	<div class="panel panel-default toggle-external toggle-internal toggle-sign collapse">
		<div class="panel-heading">
			<h2 class="panel-title">Certificate Attributes</h2>
		</div>
		<div class="panel-body">
				<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Attribute Notes</span>
		</label>
			<div class="col-sm-10">
		<span class="help-block">The following attributes are added to certificates and requests when they are created or signed. These attributes behave differently depending on the selected mode.<br/><br/><span class="toggle-internal collapse">For Internal Certificates, these attributes are added directly to the certificate as shown.</span><span class="toggle-external collapse">For Certificate Signing Requests, These attributes are added to the request but they may be ignored or changed by the CA that signs the request. <br/><br/>If this CSR will be signed using the Certificate Manager on this firewall, set the attributes when signing instead as they cannot be carried over.</span><span class="toggle-sign collapse">When Signing a Certificate Request, existing attributes in the request cannot be copied. The attributes below will be applied to the resulting certificate.</span></span>

		
	</div>
		
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span class="element-required">Certificate Type</span>
		</label>
			<div class="col-sm-10">
			<select class="form-control" name="type" id="type">
		<option value="server">Server Certificate</option><option value="user">User Certificate</option>
	</select>

		<span class="help-block">Add type-specific usage attributes to the signed certificate. Used for placing usage restrictions on, or granting abilities to, the signed certificate.</span>
	</div>
		
	</div>	<div class="form-group repeatable">
		<label class="col-sm-2 control-label">
			<span>Alternative Names</span>
		</label>
			<div class="col-sm-3">
			<select class="form-control" name="altname_type0" id="altname_type0">
		<option value="DNS">FQDN or Hostname</option><option value="IP">IP address</option><option value="URI">URI</option><option value="email">email address</option>
	</select>

		<span class="help-block">Type</span>
	</div>	<div class="col-sm-3">
		<input class="form-control" name="altname_value0" id="altname_value0" type="text">

		<span class="help-block">Value</span>
	</div>	<div class="col-sm-3">
		<button class="btn btn-warning" type="submit" value="Delete" name="deleterow0" id="deleterow0"><i class="fa fa-trash icon-embed-btn"> </i>Delete</button>

		
	</div>
			<div class="col-sm-10 col-sm-offset-2">
		<span class="help-block">
			Enter additional identifiers for the certificate in this list. The Common Name field is automatically added to the certificate as an Alternative Name. The signing CA may ignore or change these values.
		</span>
	</div>
	</div>	<div class="form-group">
		<label class="col-sm-2 control-label">
			<span>Add</span>
		</label>
			<div class="col-sm-10">
		<button class="btn btn-success" type="submit" value="Add" name="addrow" id="addrow"><i class="fa fa-plus icon-embed-btn"> </i>Add</button>

		
	</div>
		
	</div>
		</div>
	</div><input class="form-control" name="id" id="id" type="hidden" value="6136cfbfdca6d"><div class="col-sm-10 col-sm-offset-2"><button class="btn btn-primary" type="submit" value="Save" name="save" id="save"><i class="fa fa-save icon-embed-btn"> </i>Save</button><button class="btn btn-primary" type="submit" value="Export Private Key" name="exportpkey" id="exportpkey"><i class="fa fa-key icon-embed-btn"> </i>Export Private Key</button><button class="btn btn-primary" type="submit" value="Export PKCS#12" name="exportp12" id="exportp12"><i class="fa fa-archive icon-embed-btn"> </i>Export PKCS#12</button></div>
	</form><script type="text/javascript">
//<![CDATA[
events.push(function() {

	$('.import_type_toggle').click(function() {
		var x509 = (this.value === 'x509');
		hideInput('cert', !x509);
		setRequired('cert', x509);
		hideInput('key', !x509);
		setRequired('key', x509);
		hideInput('pkcs12_cert', x509);
		setRequired('pkcs12_cert', !x509);
		hideInput('pkcs12_pass', x509);
		hideCheckbox('pkcs12_intermediate', x509);
	});
	if ($('input[name=import_type]:checked').val() == 'x509') {
		hideInput('pkcs12_cert', true);
		setRequired('pkcs12_cert', false);
		hideInput('pkcs12_pass', true);
		hideCheckbox('pkcs12_intermediate', true);
		hideInput('cert', false);
		setRequired('cert', true);
		hideInput('key', false);
		setRequired('key', true);
	} else if ($('input[name=import_type]:checked').val() == 'pkcs12') {
		hideInput('cert', true);
		setRequired('cert', false);
		hideInput('key', true);
		setRequired('key', false);
		setRequired('pkcs12_cert', false);
	}



});
//]]>
</script>
	</div>
	<footer class="footer">
		<div class="container">
			<p class="text-muted">
				<a id="tpl" style="display: none;" href="#" title="Top of page"><i class="fa fa-caret-square-o-up pull-left"></i></a>
				<a target="_blank" href="https://pfsense.org">pfSense</a> is developed and maintained by <a target="_blank" href="https://netgate.com">Netgate. </a> &copy; ESF 2004 - 2021<a target="_blank" href="https://pfsense.org/license"> View license.</a>				<a id="tpr" style="display: none;" href="#" title="Top of page"><i class="fa fa-caret-square-o-up pull-right"></i></a>
			</p>
		</div>
	</footer>

	<!-- This use of filemtime() is intended to fool the browser into reloading the file (not using the cached version) if the file is changed -->

	<script src="/vendor/jquery/jquery-3.5.1.min.js?v=1622201721"></script>
 	<script src="/vendor/jquery-ui/jquery-ui-1.12.1.min.js?v=1622201721"></script>
	<script src="/vendor/bootstrap/js/bootstrap.min.js?v=1622201721"></script>
	<script src="/js/pfSense.js?v=1622201721"></script>
	<script src="/js/pfSenseHelpers.js?v=1622201721"></script>
	<script src="/js/polyfills.js?v=1622201721"></script>
	<script src="/vendor/sortable/sortable.js?v=1622201721"></script>

	<script type="text/javascript">
	//<![CDATA[
		// Un-hide the "Top of page" icons if the page is larger than the window
		if ($(document).height() > $(window).height()) {
		    $('[id^=tp]').show();
		}
	//]]>
	</script>
<script type="text/javascript">CsrfMagic.end();</script></body>
</html>
`

func getDescriptionName(html string) (string, error) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	if val, ok := reader.Find("input#descr").Attr("value"); ok {
		return val, nil
	}

	return "", errors.New("获取证书名称失败！")

}

// unmarshalEditCertRespAction dco
var unmarshalEditCertRespAction = cli.NamedAction{
	Name: "pfsense.unmarshalEditCertResp",
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		bodyBuf, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		//fmt.Println("string resp :", string(bodyBuf))

		csrf, err := cli.GetCsrfInfo(string(bodyBuf))
		if err != nil {
			r.Error = err
			return
		}

		descr, err := getDescriptionName(string(bodyBuf))
		if err != nil {
			r.Error = fmt.Errorf("未找到证书编辑页的证书名称: %v", err)
			return
		}

		out := EditCertResp{
			CSRF:  csrf,
			Descr: descr,
		}

		outRaw, err := json.Marshal(out)
		if err != nil {
			r.Error = fmt.Errorf("marshal index output info err : %v", err)
			return
		}

		err = json.Unmarshal(outRaw, &r.Data)
		if err != nil {
			r.Error = fmt.Errorf("unmarshal index output info to data err : %v", err)
			return
		}

	},
}

// unmarshalEditRespAction dco
var unmarshalEditRespAction = cli.NamedAction{
	Name: "pfsense.unmarshalEditCertResp",
	Fn: func(r *cli.Request) {
		defer r.HTTPResponse.Body.Close()

		bodyBuf, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			r.Error = fmt.Errorf("unmashal read body err : %v", err)
			return
		}
		fmt.Println("string resp :", string(bodyBuf))

	},
}
