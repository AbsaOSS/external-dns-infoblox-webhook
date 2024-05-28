package infoblox

import (
	"fmt"
	"reflect"
	"sort"

	"sigs.k8s.io/external-dns/endpoint"
)

/** test utility functions for endpoints verifications */

type byNames endpoint.ProviderSpecific

func (p byNames) Len() int           { return len(p) }
func (p byNames) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byNames) Less(i, j int) bool { return p[i].Name < p[j].Name }

// SameEndpoint returns true if two endpoints are same
// considers example.org. and example.org DNSName/Target as different endpoints
func SameEndpoint(a, b *endpoint.Endpoint) bool {

	dif := a.DNSName == b.DNSName && a.Targets.Same(b.Targets) && a.RecordType == b.RecordType && a.SetIdentifier == b.SetIdentifier &&
		a.Labels[endpoint.OwnerLabelKey] == b.Labels[endpoint.OwnerLabelKey] && a.RecordTTL == b.RecordTTL &&
		a.Labels[endpoint.ResourceLabelKey] == b.Labels[endpoint.ResourceLabelKey] &&
		a.Labels[endpoint.OwnedRecordLabelKey] == b.Labels[endpoint.OwnedRecordLabelKey] &&
		SameProviderSpecific(a.ProviderSpecific, b.ProviderSpecific)

	if !dif {
		fmt.Println("Different endpoints:")
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}

	return dif
}

// SameEndpoints compares two slices of endpoints regardless of order
// [x,y,z] == [z,x,y]
// [x,x,z] == [x,z,x]
// [x,y,y] != [x,x,y]
// [x,x,x] != [x,x,z]
func SameEndpoints(a, b []*endpoint.Endpoint) bool {
	if len(a) != len(b) {
		return false
	}

	sa := a
	sb := b

	for i := range sa {
		for j := range sb {
			if sa[i].DNSName == sb[j].DNSName && sa[i].RecordType == sb[j].RecordType {
				// fmt.Println("found:" + sa[i].DNSName + sa[i].RecordType)
				if !SameEndpoint(sa[i], sb[j]) {
					return false
				}
			}
		}
	}

	return true
}

// SameProviderSpecific verifies that two maps contain the same string/string key/value pairs
func SameProviderSpecific(a, b endpoint.ProviderSpecific) bool {
	sa := a
	sb := b
	sort.Sort(byNames(sa))
	sort.Sort(byNames(sb))
	return reflect.DeepEqual(sa, sb)
}

func AsString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func AsInt64(i *uint32) int64 {
	if i == nil {
		return 0
	}
	return int64(*i)
}
