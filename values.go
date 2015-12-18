package metrics

func (r *StandardRegistry) Values() interface{} {
	data := make(map[string]map[string]interface{})
	r.Each(func(name string, i interface{}) {
		values := make(map[string]interface{})
		switch metric := i.(type) {
		case Counter:
			values["count"] = metric.Count()
		case Gauge:
			values["value"] = metric.Value()
		case GaugeFloat64:
			values["value"] = metric.Value()
		case Healthcheck:
			values["error"] = nil
			metric.Check()
			if err := metric.Error(); nil != err {
				values["error"] = metric.Error().Error()
			}
		case Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			values["count"] = h.Count()
			values["min"] = h.Min()
			values["max"] = h.Max()
			values["mean"] = h.Mean()
			values["stddev"] = h.StdDev()
			values["median"] = ps[0]
			values["75th"] = ps[1]
			values["95th"] = ps[2]
			values["99th"] = ps[3]
			values["999th"] = ps[4]
		case Meter:
			m := metric.Snapshot()
			values["count"] = m.Count()
			values["rate.1min"] = m.Rate1()
			values["rate.5min"] = m.Rate5()
			values["rate.15min"] = m.Rate15()
			values["rate.all"] = m.RateMean()
		case Timer:
			t := metric.Snapshot()
			ps := t.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			values["count"] = t.Count()
			values["min"] = t.Min()
			values["max"] = t.Max()
			values["mean"] = t.Mean()
			values["stddev"] = t.StdDev()
			values["median"] = ps[0]
			values["75th"] = ps[1]
			values["95th"] = ps[2]
			values["99th"] = ps[3]
			values["999th"] = ps[4]
			values["rate.1min"] = t.Rate1()
			values["rate.5min"] = t.Rate5()
			values["rate.15min"] = t.Rate15()
			values["rate.all"] = t.RateMean()
		}
		data[name] = values
	})
	return data
}
