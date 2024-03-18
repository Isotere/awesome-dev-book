package main

type ConfigOption func(*Config)

type Config struct {
	dayHours DayHours
}

type DayHours struct {
	begin int
	end   int
}

func WithDayHours(dayHoursBegin, dayHoursEnd int) ConfigOption {
	return func(c *Config) {
		c.dayHours.begin = dayHoursBegin
		c.dayHours.end = dayHoursEnd
	}
}

func NewConfig(opts ...ConfigOption) Config {
	config := Config{}

	for _, opt := range opts {
		opt(&config)
	}

	return config
}
