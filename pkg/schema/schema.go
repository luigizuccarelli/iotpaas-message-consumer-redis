package schema

// IOTPaaS SDK schema
type IOTPaaS struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data [7]VitalSign
}

// VitalSign struct
// this schema can be anything really
// change as needed - remember to update the fields in the consumer.go accordingly
type VitalSign struct {
	DeviceId string `json:"deviceId,omitempty"`
	HR       int    `json:"hr"`   // heart rate
	BPS      int    `json:"bps"`  // blood pressure systolic
	BPD      int    `json:"bpd"`  // blood pressure diastolic
	SPO2     int    `json:"spo2"` // blood oxygen saturation
	Custom   map[string]interface{}
	Date     int64 `json:"date,omitempty"`
}
