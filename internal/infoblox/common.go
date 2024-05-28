package infoblox

import (
	"sort"

	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"sigs.k8s.io/external-dns/endpoint"
)

type ResponseDetail struct {
	Target string
	TTL    int64
}

type ResponseDetails []ResponseDetail

type ResponseMap struct {
	RecordType string
	Map        map[string]ResponseDetails
}

func ToAResponseMap(res []ibclient.RecordA) *ResponseMap {
	rm := &ResponseMap{
		Map:        make(map[string]ResponseDetails),
		RecordType: ibclient.ARecord,
	}
	for _, record := range res {
		if _, ok := rm.Map[AsString(record.Name)]; !ok {
			rm.Map[AsString(record.Name)] = ResponseDetails{{Target: AsString(record.Ipv4Addr), TTL: AsInt64(record.Ttl)}}
			continue
		}
		rm.Map[AsString(record.Name)] = append(rm.Map[AsString(record.Name)], ResponseDetail{Target: AsString(record.Ipv4Addr), TTL: AsInt64(record.Ttl)})
	}
	return rm
}

func ToCNAMEResponseMap(res []ibclient.RecordCNAME) *ResponseMap {
	rm := &ResponseMap{
		Map:        make(map[string]ResponseDetails),
		RecordType: ibclient.CnameRecord,
	}
	for _, record := range res {
		if _, ok := rm.Map[AsString(record.Name)]; !ok {
			rm.Map[AsString(record.Name)] = ResponseDetails{{Target: AsString(record.Canonical), TTL: AsInt64(record.Ttl)}}
			continue
		}
		rm.Map[AsString(record.Name)] = append(rm.Map[AsString(record.Name)], ResponseDetail{Target: AsString(record.Canonical), TTL: AsInt64(record.Ttl)})
	}
	return rm
}

func ToTXTResponseMap(res []ibclient.RecordTXT) *ResponseMap {
	rm := &ResponseMap{
		Map:        make(map[string]ResponseDetails),
		RecordType: ibclient.TxtRecord,
	}
	for _, record := range res {
		if _, ok := rm.Map[AsString(record.Name)]; !ok {
			rm.Map[AsString(record.Name)] = ResponseDetails{{Target: AsString(record.Text), TTL: AsInt64(record.Ttl)}}
			continue
		}
		rm.Map[AsString(record.Name)] = append(rm.Map[AsString(record.Name)], ResponseDetail{Target: AsString(record.Text), TTL: AsInt64(record.Ttl)})
	}
	return rm
}

func ToHostResponseMap(res []ibclient.HostRecord) *ResponseMap {
	rm := &ResponseMap{
		Map:        make(map[string]ResponseDetails),
		RecordType: ibclient.ARecord, //.HostRecordConst,
	}
	for _, record := range res {
		rds := ResponseDetails{}
		for _, ip := range record.Ipv4Addrs {
			rds = append(rds, ResponseDetail{Target: AsString(ip.Ipv4Addr), TTL: AsInt64(record.Ttl)})
		}
		if _, ok := rm.Map[AsString(record.Name)]; !ok {
			rm.Map[AsString(record.Name)] = rds
			continue
		}
		rm.Map[AsString(record.Name)] = append(rm.Map[AsString(record.Name)], rds...)
	}
	return rm
}

// TODO: ToPTRResponseMap
//if p.createPTR {
//	// infoblox doesn't accept reverse zone's fqdn, and instead expects .in-addr.arpa zone
//	// so convert our zone fqdn (if it is a correct cidr block) into in-addr.arpa address and pass that into infoblox
//	// example: 10.196.38.0/24 becomes 38.196.10.in-addr.arpa
//	arpaZone, err := transform.ReverseDomainName(zone.Fqdn)
//	if err == nil {
//		var resP []ibclient.RecordPTR
//		objP := ibclient.NewEmptyRecordPTR()
//		objP.Zone = arpaZone
//		objP.View = p.view
//		err = p.client.GetObject(objP, "", searchParams, &resP)
//		if err != nil && !isNotFoundError(err) {
//			return nil, fmt.Errorf("could not fetch PTR records from zone '%s': %w", zone.Fqdn, err)
//		}
//		for _, res := range resP {
//			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(res.PtrdName,
//				endpoint.RecordTypePTR,
//				endpoint.TTL(int(res.Ttl)),
//				res.Ipv4Addr,
//			),
//			)
//		}
//	}
//}

func (rd ResponseDetails) ToEndpointDetail() (targets []string, ttl endpoint.TTL) {
	for _, v := range rd {
		targets = append(targets, v.Target)
		ttl = endpoint.TTL(v.TTL)
	}
	return
}

func (rm *ResponseMap) ToEndpoints() []*endpoint.Endpoint {
	// TODO: PTR provider specific label records
	//		if p.createPTR {
	//			newEndpoint.WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
	//		}
	var endpoints []*endpoint.Endpoint
	for k, v := range rm.Map {
		targets, ttl := v.ToEndpointDetail()
		ep := endpoint.NewEndpointWithTTL(k, rm.RecordType, ttl, targets...)
		sort.Sort(ep.Targets)
		endpoints = append(endpoints, ep)
	}
	return endpoints
}

// TODO: update A records that have PTR record created for them already
//if p.createPTR {
//	// save all ptr records into map for a quick look up
//	ptrRecordsMap := make(map[string]bool)
//	for _, ptrRecord := range endpoints {
//		if ptrRecord.RecordType != endpoint.RecordTypePTR {
//			continue
//		}
//		ptrRecordsMap[ptrRecord.DNSName] = true
//	}
//
//	for i := range endpoints {
//		if endpoints[i].RecordType != endpoint.RecordTypeA {
//			continue
//		}
//		// if PTR record already exists for A record, then mark it as such
//		if ptrRecordsMap[endpoints[i].DNSName] {
//			found := false
//			for j := range endpoints[i].ProviderSpecific {
//				if endpoints[i].ProviderSpecific[j].Name == providerSpecificInfobloxPtrRecord {
//					endpoints[i].ProviderSpecific[j].Value = "true"
//					found = true
//				}
//			}
//			if !found {
//				endpoints[i].WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
//			}
//		}
//	}
//}
