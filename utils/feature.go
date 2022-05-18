package utils

import (
	"encoding/binary"
	"math"

	log "github.com/sirupsen/logrus"

	commonapis "gitlab.sz.sensetime.com/viper/commonapis/api"
	"gitlab.sz.sensetime.com/viper/gosdkwrapper/feature"
)

// EncodeFeatureFloat2Byte 编码浮点数特征为字节特征
func EncodeFeatureFloat2Byte(feature []float32) []byte {
	var data []byte
	for _, v := range feature {
		bits := math.Float32bits(v)
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, bits)
		data = append(data, bytes...)
	}
	return data
}

// EncodeFeatureByte2Float 编码字节特征为浮点数特征
func EncodeFeatureByte2Float(blob []byte) []float32 {
	off := 0
	var feature []float32
	for i := 0; i < (len(blob) / 4); i++ {
		v := binary.LittleEndian.Uint32(blob[off : off+4])
		feature = append(feature, math.Float32frombits(v))
		off += 4
	}
	return feature
}

// Normalize 特征归一化
func Normalize(feature []float32) []float32 {
	dst := make([]float32, len(feature))
	var dev float32
	for _, v := range feature {
		dev += (v * v)
	}
	std := math.Sqrt(float64(dev))
	for i := range dst {
		dst[i] = (feature[i]) / float32(std)
	}
	return dst
}

// DecodeObjectFeature2RawFeature decodes ObjectFeature to RawFeature
func DecodeObjectFeature2RawFeature(dec *feature.Decoder, feat *commonapis.ObjectFeature) (*feature.RawFeature, error) {
	pfeature, err := feature.NewPersistedFeatureFromBytes(feat.Blob, false)
	if err != nil {
		return nil, err
	}
	raw, err := dec.Decode(pfeature)
	if err != nil {
		return nil, err
	}
	return &raw, nil
}

// EncodeRawFeature2ObjectFeature encodes RawFeature to ObjectFeature
func EncodeRawFeature2ObjectFeature(enc *feature.Encoder, raw *feature.RawFeature) (*commonapis.ObjectFeature, error) {
	pfeature, err := enc.Encode(*raw)
	if err != nil {
		return nil, err
	}
	feat := &commonapis.ObjectFeature{
		Version: pfeature.Header().Version,
		Blob:    pfeature.Bytes(),
	}
	return feat, nil
}

// EncodeFeatureByte2ObjectFeature encodes bytes of RawFeature to ObjectFeature
func EncodeFeatureByte2ObjectFeature(enc *feature.Encoder, data []byte, version int32) (*commonapis.ObjectFeature, error) {
	raw, err := feature.NewRawFeatureFromFloat32Bytes(version, data)
	if err != nil {
		return nil, err
	}
	feat, err := EncodeRawFeature2ObjectFeature(enc, &raw)
	if err != nil {
		return nil, err
	}
	return feat, nil
}

// SafelyEncodeBytesRawFeature2BytesPersistedFeature encodes bytes of RawFeature to bytes of PersistedFeature
func SafelyEncodeBytesRawFeature2BytesPersistedFeature(enc *feature.Encoder, data []byte, version int32) []byte {
	raw, err := feature.NewRawFeatureFromFloat32Bytes(version, data)
	if err != nil {
		log.WithError(err).Error("failed to do feature.NewRawFeatureFromFloat32Bytes")
	}
	pfeature, err := enc.Encode(raw)
	if err != nil {
		log.WithError(err).Error("failed to do enc.Encode")
	}
	return pfeature.Bytes()
}

// SafelyDecodeObjectFeature2BytesRawFeature decodes ObjectFeature to bytes of RawFeature
func SafelyDecodeObjectFeature2BytesRawFeature(dec *feature.Decoder, feat *commonapis.ObjectFeature) []float32 {
	pfeature, err := feature.NewPersistedFeatureFromBytes(feat.Blob, false)
	if err != nil {
		log.WithError(err).Error("failed to do feature.NewPersistedFeatureFromBytes")
	}
	raw, err := dec.Decode(pfeature)
	if err != nil {
		log.WithError(err).Error("failed to do dec.Decode")
	}
	return raw.Raw
}
