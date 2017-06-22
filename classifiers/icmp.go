package classifiers

import (
	"github.com/google/gopacket/layers"
	"github.com/mushorg/go-dpi"
)

type IcmpClassifier struct{}

func (_ IcmpClassifier) HeuristicClassify(flow *godpi.Flow) bool {
	if len(flow.Packets) == 0 {
		return false
	}
	for _, packet := range flow.Packets {
		if layer := (*packet).Layer(layers.LayerTypeIPv4); layer != nil {
			ipLayer := layer.(*layers.IPv4)
			if ipLayer.Protocol == layers.IPProtocolICMPv4 {
				return true
			}
		} else if layer := (*packet).Layer(layers.LayerTypeIPv6); layer != nil {
			ipLayer := layer.(*layers.IPv6)
			if ipLayer.NextHeader == layers.IPProtocolICMPv6 {
				return true
			}
		}
	}
	return false
}

func (_ IcmpClassifier) GetProtocol() godpi.Protocol {
	return godpi.Icmp
}
