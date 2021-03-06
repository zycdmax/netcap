/*
 * NETCAP - Traffic Analysis Framework
 * Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package encoder

import (
	"github.com/dreadl0ck/netcap/types"
	"github.com/golang/protobuf/proto"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var ip6FragmentEncoder = CreateLayerEncoder(types.Type_NC_IPv6Fragment, layers.LayerTypeIPv6Fragment, func(layer gopacket.Layer, timestamp string) proto.Message {
	if ip6f, ok := layer.(*layers.IPv6Fragment); ok {
		return &types.IPv6Fragment{
			Timestamp:      timestamp,
			NextHeader:     int32(ip6f.NextHeader),
			Reserved1:      int32(ip6f.Reserved1),
			FragmentOffset: int32(ip6f.FragmentOffset),
			Reserved2:      int32(ip6f.Reserved2),
			MoreFragments:  bool(ip6f.MoreFragments),
			Identification: uint32(ip6f.Identification),
		}
	}
	return nil
})
