// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ebpfcommon

import (
	"encoding/hex"
	"log/slog"

	"go.opentelemetry.io/obi/pkg/internal/largebuf"
)

func hexHead(b []byte, max int) string {
	if len(b) == 0 {
		return ""
	}
	if len(b) > max {
		b = b[:max]
	}
	return hex.EncodeToString(b)
}

func debugLogHTTPLargeBufferChunk(event *TCPLargeBufferHeader, key largeBufferKey, chunk []byte, totalLen int) {
	if key.kind != KindLayerApp {
		return
	}
	slog.Debug("HTTP large buffer chunk",
		"packetType", event.PacketType,
		"direction", event.Direction,
		"action", event.Action,
		"source", event.Source,
		"chunkLen", len(chunk),
		"totalLen", totalLen,
		"chunkHead32Hex", hexHead(chunk, 32),
		"traceID", key.traceID,
		"spanID", key.spanID,
		"conn", key.connInfo,
	)
}

func debugLogHTTPLargeBufferExtract(key largeBufferKey, lb *largebuf.LargeBuffer) {
	if key.kind != KindLayerApp {
		return
	}
	p := largebuf.ProbeShape(lb)
	slog.Debug("HTTP large buffer extract",
		"packetType", key.packetType,
		"direction", key.direction,
		"totalLen", p.Len,
		"startsWithHTTP", p.StartsWithHTTP,
		"head64Hex", p.Head64Hex,
		"traceID", key.traceID,
		"spanID", key.spanID,
		"conn", key.connInfo,
	)
}
