package benchmarking

import "sync"

func Init() {
	metrics = sync.Map{}
}

func AddMetric(metric Metric) error {
	if !metric.Validate() {
		return ErrorInvalidMetric
	}
	metric.Save()
	return nil
}

func GetMetric(id string) (Metric, error) {
	metric, ok := metrics.Load(id)
	if !ok {
		return Metric{}, ErrorMetricNotFound
	}

	return metric.(Metric), nil
}

func SetMetric(id string, value interface{}) error {
	metric := Metric{ID: id, Value: value}
	metric.Value = value
	metric.Save()
	return nil
}

func DeleteMetric(id string) error {
	metrics.Delete(id)
	return nil
}

func GetAllMetrics() []Metric {
	var allMetrics []Metric
	metrics.Range(func(key, value interface{}) bool {
		allMetrics = append(allMetrics, value.(Metric))
		return true
	})
	return allMetrics
}

func ShowMetric(id string) {
	metric, err := GetMetric(id)
	if err != nil {
		return
	}
	metric.Show()
}

func ShowMetrics() {
	for _, metric := range GetAllMetrics() {
		metric.Show()
	}
}
