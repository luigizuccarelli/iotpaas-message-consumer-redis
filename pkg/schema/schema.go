package schema

// IOTPaaS SDK schema
type IOTPaaS struct {
	PatientId string `json:"patient_id"`
	DeviceId  string `json:"device_id"`
	Name      string `json:"name"`
	Data      []VitalSign
}

// VitalSign struct
// this schema can be anything really
// change as needed - remember to update the fields in the consumer.go accordingly
type VitalSign struct {
	HR     int `json:"hr"`   // heart rate
	BPS    int `json:"bps"`  // blood pressure systolic
	BPD    int `json:"bpd"`  // blood pressure diastolic
	SPO2   int `json:"spo2"` // blood oxygen saturation
	Custom map[string]interface{}
	Date   string `json:"date,omitempty"`
}
