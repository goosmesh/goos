package redis

import (
	"context"

	// "fmt"
	"time"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

// ServeDNS implements the plugin.Handler interface.
func (redis *Redis) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Info("serveDNS")
	state := request.Request{W: w, Req: r}

	qname := state.Name()
	qtype := state.Type()

	log.Info("name : ", qname)
	log.Info("type : ", qtype)

	if time.Since(redis.LastZoneUpdate) > zoneUpdateTime {
		redis.LoadZones()
	}

	zone := plugin.Zones(redis.Zones).Matches(qname)
	log.Info("zone : ", zone)
	if zone == "" {
		return plugin.NextOrFailure(qname, redis.Next, ctx, w, r)
	}

	z := redis.load(zone)
	if z == nil {
		return redis.errorResponse(state, zone, dns.RcodeServerFailure, nil)
	}

	location := redis.findLocation(qname, z)
	if len(location) == 0 { // empty, no results
		return redis.errorResponse(state, zone, dns.RcodeNameError, nil)
	}
	log.Info("location : ", location)

	answers := make([]dns.RR, 0, 10)
	extras := make([]dns.RR, 0, 10)

	record := redis.get(location, z)

	switch qtype {
	case "A":
		answers, extras = redis.A(qname, z, record, state.Proto())
	case "AAAA":
		answers, extras = redis.AAAA(qname, z, record, state.Proto())
	case "CNAME":
		answers, extras = redis.CNAME(qname, z, record)
	case "TXT":
		answers, extras = redis.TXT(qname, z, record)
	case "NS":
		answers, extras = redis.NS(qname, z, record, state.Proto())
	case "MX":
		answers, extras = redis.MX(qname, z, record, state.Proto())
	case "SRV":
		answers, extras = redis.SRV(qname, z, record, state.Proto())
	case "SOA":
		answers, extras = redis.SOA(qname, z, record)
	case "CAA":
		answers, extras = redis.CAA(qname, z, record)
	default:
		return redis.errorResponse(state, zone, dns.RcodeNotImplemented, nil)
	}

	//log.Info(answers)

	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative, m.RecursionAvailable, m.Compress = true, false, true

	m.Answer = append(m.Answer, answers...)
	m.Extra = append(m.Extra, extras...)

	state.SizeAndDo(m)
	m = state.Scrub(m)
	_ = w.WriteMsg(m)
	return dns.RcodeSuccess, nil
}

// Name implements the Handler interface.
func (redis *Redis) Name() string { return "redis" }

func (redis *Redis) errorResponse(state request.Request, zone string, rcode int, err error) (int, error) {
	m := new(dns.Msg)
	m.SetRcode(state.Req, rcode)
	m.Authoritative, m.RecursionAvailable, m.Compress = true, false, true

	// m.Ns, _ = redis.SOA(state.Name(), zone, nil)

	state.SizeAndDo(m)
	_ = state.W.WriteMsg(m)
	// Return success as the rcode to signal we have written to the client.
	return dns.RcodeSuccess, err
}
