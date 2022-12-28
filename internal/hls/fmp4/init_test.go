//nolint:dupl
package fmp4

import (
	"testing"

	"github.com/aler9/gortsplib/v2/pkg/codecs/mpeg4audio"
	"github.com/aler9/gortsplib/v2/pkg/format"
	"github.com/stretchr/testify/require"
)

var testSPS = []byte{
	0x67, 0x42, 0xc0, 0x28, 0xd9, 0x00, 0x78, 0x02,
	0x27, 0xe5, 0x84, 0x00, 0x00, 0x03, 0x00, 0x04,
	0x00, 0x00, 0x03, 0x00, 0xf0, 0x3c, 0x60, 0xc9,
	0x20,
}

var testVideoTrack = &format.H264{
	PayloadTyp:        96,
	SPS:               testSPS,
	PPS:               []byte{0x08},
	PacketizationMode: 1,
}

var testAudioTrack = &format.MPEG4Audio{
	PayloadTyp: 97,
	Config: &mpeg4audio.Config{
		Type:         2,
		SampleRate:   44100,
		ChannelCount: 2,
	},
	SizeLength:       13,
	IndexLength:      3,
	IndexDeltaLength: 3,
}

func TestInitMarshal(t *testing.T) {
	t.Run("video + audio", func(t *testing.T) {
		init := Init{
			Tracks: []*InitTrack{
				{
					ID:        1,
					TimeScale: 90000,
					Format:    testVideoTrack,
				},
				{
					ID:        2,
					TimeScale: uint32(testAudioTrack.ClockRate()),
					Format:    testAudioTrack,
				},
			},
		}

		byts, err := init.Marshal()
		require.NoError(t, err)

		require.Equal(t, []byte{
			0x00, 0x00, 0x00, 0x20,
			'f', 't', 'y', 'p',
			0x6d, 0x70, 0x34, 0x32, 0x00, 0x00, 0x00, 0x01,
			0x6d, 0x70, 0x34, 0x31, 0x6d, 0x70, 0x34, 0x32,
			0x69, 0x73, 0x6f, 0x6d, 0x68, 0x6c, 0x73, 0x66,
			0x00, 0x00, 0x04, 0x64,
			'm', 'o', 'o', 'v',
			0x00, 0x00, 0x00, 0x6c,
			'm', 'v', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xe8,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x01, 0xec,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x03,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x07, 0x80, 0x00, 0x00, 0x04, 0x38, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x88, 0x6d, 0x64, 0x69, 0x61,
			0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x5f, 0x90,
			0x00, 0x00, 0x00, 0x00, 0x55, 0xc4, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x2d, 0x68, 0x64, 0x6c, 0x72,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x76, 0x69, 0x64, 0x65, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x56, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x61, 0x6e,
			0x64, 0x6c, 0x65, 0x72, 0x00, 0x00, 0x00, 0x01,
			0x33,
			'm', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x14,
			'v', 'm', 'h', 'd',
			0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x24, 0x64, 0x69, 0x6e,
			0x66, 0x00, 0x00, 0x00, 0x1c, 0x64, 0x72, 0x65,
			0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x0c, 0x75, 0x72, 0x6c,
			0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0xf3, 0x73, 0x74, 0x62, 0x6c, 0x00, 0x00, 0x00,
			0xa7, 0x73, 0x74, 0x73, 0x64, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x97, 0x61, 0x76, 0x63, 0x31, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x80, 0x04,
			0x38, 0x00, 0x48, 0x00, 0x00, 0x00, 0x48, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x18, 0xff, 0xff, 0x00, 0x00, 0x00, 0x2d, 0x61,
			0x76, 0x63, 0x43, 0x01, 0x42, 0xc0, 0x28, 0x03,
			0x01, 0x00, 0x19, 0x67, 0x42, 0xc0, 0x28, 0xd9,
			0x00, 0x78, 0x02, 0x27, 0xe5, 0x84, 0x00, 0x00,
			0x03, 0x00, 0x04, 0x00, 0x00, 0x03, 0x00, 0xf0,
			0x3c, 0x60, 0xc9, 0x20, 0x01, 0x00, 0x01, 0x08,
			0x00, 0x00, 0x00, 0x14, 0x62, 0x74, 0x72, 0x74,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0x40,
			0x00, 0x0f, 0x42, 0x40, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x74, 0x73, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x73, 0x63, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14,
			0x73, 0x74, 0x73, 0x7a, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x73, 0x74, 0x63, 0x6f,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x01, 0xbc,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x58,
			0x6d, 0x64, 0x69, 0x61, 0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0xac, 0x44, 0x00, 0x00, 0x00, 0x00,
			0x55, 0xc4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2d,
			0x68, 0x64, 0x6c, 0x72, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x73, 0x6f, 0x75, 0x6e,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x53, 0x6f, 0x75, 0x6e,
			0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
			0x00, 0x00, 0x00, 0x01, 0x03, 0x6d, 0x69, 0x6e,
			0x66, 0x00, 0x00, 0x00, 0x10,
			's', 'm', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x24, 0x64, 0x69, 0x6e,
			0x66, 0x00, 0x00, 0x00, 0x1c, 0x64, 0x72, 0x65,
			0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x0c, 0x75, 0x72, 0x6c,
			0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0xc7, 0x73, 0x74, 0x62, 0x6c, 0x00, 0x00, 0x00,
			0x7b, 0x73, 0x74, 0x73, 0x64, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x6b, 0x6d, 0x70, 0x34, 0x61, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00,
			0x10, 0x00, 0x00, 0x00, 0x00, 0xac, 0x44, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x33, 0x65, 0x73, 0x64,
			0x73, 0x00, 0x00, 0x00, 0x00, 0x03, 0x80, 0x80,
			0x80, 0x22, 0x00, 0x02, 0x00, 0x04, 0x80, 0x80,
			0x80, 0x14, 0x40, 0x15, 0x00, 0x00, 0x00, 0x00,
			0x01, 0xf7, 0x39, 0x00, 0x01, 0xf7, 0x39, 0x05,
			0x80, 0x80, 0x80, 0x02, 0x12, 0x10, 0x06, 0x80,
			0x80, 0x80, 0x01, 0x02, 0x00, 0x00, 0x00, 0x14,
			0x62, 0x74, 0x72, 0x74, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0xf7, 0x39, 0x00, 0x01, 0xf7, 0x39,
			0x00, 0x00, 0x00, 0x10, 0x73, 0x74, 0x74, 0x73,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x73, 0x74, 0x73, 0x63,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x14, 0x73, 0x74, 0x73, 0x7a,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x63, 0x6f, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x48,
			0x6d, 0x76, 0x65, 0x78, 0x00, 0x00, 0x00, 0x20,
			0x74, 0x72, 0x65, 0x78, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20,
			0x74, 0x72, 0x65, 0x78, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		}, byts)
	})

	t.Run("video only", func(t *testing.T) {
		init := Init{
			Tracks: []*InitTrack{
				{
					ID:        1,
					TimeScale: 90000,
					Format:    testVideoTrack,
				},
			},
		}

		byts, err := init.Marshal()
		require.NoError(t, err)

		require.Equal(t, []byte{
			0x00, 0x00, 0x00, 0x20,
			'f', 't', 'y', 'p',
			0x6d, 0x70, 0x34, 0x32, 0x00, 0x00, 0x00, 0x01,
			0x6d, 0x70, 0x34, 0x31, 0x6d, 0x70, 0x34, 0x32,
			0x69, 0x73, 0x6f, 0x6d, 0x68, 0x6c, 0x73, 0x66,
			0x00, 0x00, 0x02, 0x88,
			'm', 'o', 'o', 'v',
			0x00, 0x00, 0x00, 0x6c,
			'm', 'v', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xe8,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x01, 0xec,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x03,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x07, 0x80, 0x00, 0x00, 0x04, 0x38, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x88, 0x6d, 0x64, 0x69, 0x61,
			0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x5f, 0x90,
			0x00, 0x00, 0x00, 0x00, 0x55, 0xc4, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x2d, 0x68, 0x64, 0x6c, 0x72,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x76, 0x69, 0x64, 0x65, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x56, 0x69, 0x64, 0x65, 0x6f, 0x48, 0x61, 0x6e,
			0x64, 0x6c, 0x65, 0x72, 0x00, 0x00, 0x00, 0x01,
			0x33,
			'm', 'i', 'n', 'f',
			0x00, 0x00, 0x00,
			0x14,
			'v', 'm', 'h', 'd',
			0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x24,
			'd', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x1c, 0x64, 0x72, 0x65,
			0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x0c, 0x75, 0x72, 0x6c,
			0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0xf3, 0x73, 0x74, 0x62, 0x6c, 0x00, 0x00, 0x00,
			0xa7, 0x73, 0x74, 0x73, 0x64, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x97, 0x61, 0x76, 0x63, 0x31, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x80, 0x04,
			0x38, 0x00, 0x48, 0x00, 0x00, 0x00, 0x48, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x18, 0xff, 0xff, 0x00, 0x00, 0x00, 0x2d, 0x61,
			0x76, 0x63, 0x43, 0x01, 0x42, 0xc0, 0x28, 0x03,
			0x01, 0x00, 0x19, 0x67, 0x42, 0xc0, 0x28, 0xd9,
			0x00, 0x78, 0x02, 0x27, 0xe5, 0x84, 0x00, 0x00,
			0x03, 0x00, 0x04, 0x00, 0x00, 0x03, 0x00, 0xf0,
			0x3c, 0x60, 0xc9, 0x20, 0x01, 0x00, 0x01, 0x08,
			0x00, 0x00, 0x00, 0x14, 0x62, 0x74, 0x72, 0x74,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0x40,
			0x00, 0x0f, 0x42, 0x40, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x74, 0x73, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x73, 0x63, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14,
			0x73, 0x74, 0x73, 0x7a, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x73, 0x74, 0x63, 0x6f,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x28, 0x6d, 0x76, 0x65, 0x78,
			0x00, 0x00, 0x00, 0x20, 0x74, 0x72, 0x65, 0x78,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, byts)
	})

	t.Run("audio only", func(t *testing.T) {
		init := &Init{
			Tracks: []*InitTrack{
				{
					ID:        1,
					TimeScale: uint32(testAudioTrack.ClockRate()),
					Format:    testAudioTrack,
				},
			},
		}

		byts, err := init.Marshal()
		require.NoError(t, err)

		require.Equal(t, []byte{
			0x00, 0x00, 0x00, 0x20,
			'f', 't', 'y', 'p',
			0x6d, 0x70, 0x34, 0x32, 0x00, 0x00, 0x00, 0x01,
			0x6d, 0x70, 0x34, 0x31, 0x6d, 0x70, 0x34, 0x32,
			0x69, 0x73, 0x6f, 0x6d, 0x68, 0x6c, 0x73, 0x66,
			0x00, 0x00, 0x02, 0x58,
			'm', 'o', 'o', 'v',
			0x00, 0x00, 0x00, 0x6c,
			'm', 'v', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xe8,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x01, 0xbc,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x03,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x58,
			'm', 'd', 'i', 'a',
			0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xac, 0x44,
			0x00, 0x00, 0x00, 0x00, 0x55, 0xc4, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x2d,
			'h', 'd', 'l', 'r',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x73, 0x6f, 0x75, 0x6e, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x53, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x61, 0x6e,
			0x64, 0x6c, 0x65, 0x72, 0x00, 0x00, 0x00, 0x01,
			0x03,
			'm', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x10,
			's', 'm', 'h', 'd',
			0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x24,
			'd', 'i', 'n', 'f',
			0x00, 0x00, 0x00,
			0x1c, 0x64, 0x72, 0x65, 0x66, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x0c, 0x75, 0x72, 0x6c, 0x20, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0xc7, 0x73, 0x74, 0x62,
			0x6c, 0x00, 0x00, 0x00, 0x7b, 0x73, 0x74, 0x73,
			0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x6b,
			'm', 'p', '4', 'a',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x02, 0x00, 0x10, 0x00, 0x00, 0x00,
			0x00, 0xac, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x33,
			'e', 's', 'd', 's',
			0x00, 0x00, 0x00,
			0x00, 0x03, 0x80, 0x80, 0x80, 0x22, 0x00, 0x01,
			0x00, 0x04, 0x80, 0x80, 0x80, 0x14, 0x40, 0x15,
			0x00, 0x00, 0x00, 0x00, 0x01, 0xf7, 0x39, 0x00,
			0x01, 0xf7, 0x39, 0x05, 0x80, 0x80, 0x80, 0x02,
			0x12, 0x10, 0x06, 0x80, 0x80, 0x80, 0x01, 0x02,
			0x00, 0x00, 0x00, 0x14,
			'b', 't', 'r', 't',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xf7, 0x39,
			0x00, 0x01, 0xf7, 0x39, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x74, 0x73, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
			0x73, 0x74, 0x73, 0x63, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14,
			0x73, 0x74, 0x73, 0x7a, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x73, 0x74, 0x63, 0x6f,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x28,
			'm', 'v', 'e', 'x',
			0x00, 0x00, 0x00, 0x20,
			't', 'r', 'e', 'x',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}, byts)
	})
}

func TestInitUnmarshal(t *testing.T) {
	t.Run("video", func(t *testing.T) {
		byts := []byte{
			0x00, 0x00, 0x00, 0x1c,
			'f', 't', 'y', 'p',
			0x64, 0x61, 0x73, 0x68, 0x00, 0x00, 0x00, 0x01,
			0x69, 0x73, 0x6f, 0x6d, 0x61, 0x76, 0x63, 0x31,
			0x64, 0x61, 0x73, 0x68, 0x00, 0x00, 0x02, 0x92,
			'm', 'o', 'o', 'v',
			0x00, 0x00, 0x00, 0x6c,
			'm', 'v', 'h', 'd',
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x98, 0x96, 0x80, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
			0x00, 0x00, 0x01, 0xf6,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x40, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00,
			0x02, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x01, 0x92,
			0x6d, 0x64, 0x69, 0x61, 0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x98, 0x96, 0x80, 0x00, 0x00, 0x00, 0x00,
			0x55, 0xc4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38,
			0x68, 0x64, 0x6c, 0x72, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x76, 0x69, 0x64, 0x65,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x42, 0x72, 0x6f, 0x61,
			0x64, 0x70, 0x65, 0x61, 0x6b, 0x20, 0x56, 0x69,
			0x64, 0x65, 0x6f, 0x20, 0x48, 0x61, 0x6e, 0x64,
			0x6c, 0x65, 0x72, 0x00, 0x00, 0x00, 0x01, 0x32,
			'm', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x14,
			'v', 'm', 'h', 'd',
			0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x24,
			'd', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x1c, 0x64, 0x72, 0x65, 0x66,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x0c, 0x75, 0x72, 0x6c, 0x20,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xf2,
			0x73, 0x74, 0x62, 0x6c, 0x00, 0x00, 0x00, 0xa6,
			0x73, 0x74, 0x73, 0x64, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x96,
			0x61, 0x76, 0x63, 0x31, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x02, 0x1c,
			0x00, 0x48, 0x00, 0x00, 0x00, 0x48, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x04, 0x68,
			0x32, 0x36, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18,
			0xff, 0xff, 0x00, 0x00, 0x00, 0x30, 0x61, 0x76,
			0x63, 0x43, 0x01, 0x42, 0xc0, 0x1f, 0xff, 0xe1,
			0x00, 0x19, 0x67, 0x42, 0xc0, 0x1f, 0xd9, 0x00,
			0xf0, 0x11, 0x7e, 0xf0, 0x11, 0x00, 0x00, 0x03,
			0x00, 0x01, 0x00, 0x00, 0x03, 0x00, 0x30, 0x8f,
			0x18, 0x32, 0x48, 0x01, 0x00, 0x04, 0x68, 0xcb,
			0x8c, 0xb2, 0x00, 0x00, 0x00, 0x10, 0x70, 0x61,
			0x73, 0x70, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x10, 0x73, 0x74,
			0x74, 0x73, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x73, 0x74,
			0x73, 0x63, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x73, 0x74,
			0x73, 0x7a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x10, 0x73, 0x74, 0x63, 0x6f, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x28, 0x6d, 0x76, 0x65, 0x78, 0x00, 0x00,
			0x00, 0x20, 0x74, 0x72, 0x65, 0x78, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}

		var init Init
		err := init.Unmarshal(byts)
		require.NoError(t, err)

		require.Equal(t, Init{
			Tracks: []*InitTrack{
				{
					ID:        256,
					TimeScale: 10000000,
					Format: &format.H264{
						PayloadTyp: 96,
						SPS: []byte{
							0x67, 0x42, 0xc0, 0x1f, 0xd9, 0x00, 0xf0, 0x11,
							0x7e, 0xf0, 0x11, 0x00, 0x00, 0x03, 0x00, 0x01,
							0x00, 0x00, 0x03, 0x00, 0x30, 0x8f, 0x18, 0x32,
							0x48,
						},
						PPS: []byte{
							0x68, 0xcb, 0x8c, 0xb2,
						},
						PacketizationMode: 1,
					},
				},
			},
		}, init)
	})

	t.Run("audio", func(t *testing.T) {
		byts := []byte{
			0x00, 0x00, 0x00, 0x18,
			'f', 't', 'y', 'p',
			0x69, 0x73, 0x6f, 0x35, 0x00, 0x00, 0x00, 0x01,
			0x69, 0x73, 0x6f, 0x35, 0x64, 0x61, 0x73, 0x68,
			0x00, 0x00, 0x02, 0x43,
			'm', 'o', 'o', 'v',
			0x00, 0x00, 0x00, 0x6c,
			'm', 'v', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x01, 0xa7,
			't', 'r', 'a', 'k',
			0x00, 0x00, 0x00, 0x5c,
			't', 'k', 'h', 'd',
			0x00, 0x00, 0x00, 0x07,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x01, 0x43, 0x6d, 0x64, 0x69, 0x61,
			0x00, 0x00, 0x00, 0x20,
			'm', 'd', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80,
			0x00, 0x00, 0x00, 0x00, 0x55, 0xc4, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x38, 0x68, 0x64, 0x6c, 0x72,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x73, 0x6f, 0x75, 0x6e, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x42, 0x72, 0x6f, 0x61, 0x64, 0x70, 0x65, 0x61,
			0x6b, 0x20, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x20,
			0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x00,
			0x00, 0x00, 0x00, 0xe3,
			'm', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x10,
			's', 'm', 'h', 'd',
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x24,
			'd', 'i', 'n', 'f',
			0x00, 0x00, 0x00, 0x1c, 0x64, 0x72, 0x65, 0x66,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x0c, 0x75, 0x72, 0x6c, 0x20,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xa7,
			0x73, 0x74, 0x62, 0x6c, 0x00, 0x00, 0x00, 0x5b,
			0x73, 0x74, 0x73, 0x64, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x4b,
			0x6d, 0x70, 0x34, 0x61, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x10,
			0x00, 0x00, 0x00, 0x00, 0xbb, 0x80, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x27, 0x65, 0x73, 0x64, 0x73,
			0x00, 0x00, 0x00, 0x00, 0x03, 0x19, 0x00, 0x00,
			0x00, 0x04, 0x11, 0x40, 0x15, 0x00, 0x30, 0x00,
			0x00, 0x11, 0x94, 0x00, 0x00, 0x11, 0x94, 0x00,
			0x05, 0x02, 0x11, 0x90, 0x06, 0x01, 0x02, 0x00,
			0x00, 0x00, 0x10, 0x73, 0x74, 0x74, 0x73, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x10, 0x73, 0x74, 0x73, 0x63, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x14, 0x73, 0x74, 0x73, 0x7a, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x73,
			0x74, 0x63, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x6d,
			0x76, 0x65, 0x78, 0x00, 0x00, 0x00, 0x20, 0x74,
			0x72, 0x65, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00,
		}

		var init Init
		err := init.Unmarshal(byts)
		require.NoError(t, err)

		require.Equal(t, Init{
			Tracks: []*InitTrack{
				{
					ID:        257,
					TimeScale: 10000000,
					Format: &format.MPEG4Audio{
						PayloadTyp: 96,
						Config: &mpeg4audio.Config{
							Type:         mpeg4audio.ObjectTypeAACLC,
							SampleRate:   48000,
							ChannelCount: 2,
						},
						SizeLength:       13,
						IndexLength:      3,
						IndexDeltaLength: 3,
					},
				},
			},
		}, init)
	})
}
