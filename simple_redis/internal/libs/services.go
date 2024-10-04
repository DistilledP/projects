package libs

var cachedServices *Services

// Services ...
type Services struct {
	logger *LineLogger
	config *config
}

func init() {
	cachedServices = &Services{}
}

// GetServices ...
func GetServices() *Services {
	return cachedServices
}

func (s *Services) GetLogger() *LineLogger {
	if cachedServices.logger == nil {
		config := s.GetConfig()
		cachedServices.logger = newLogger("[LUNGFISH] ", config.LogLevel)
	}

	return cachedServices.logger
}

func (s *Services) GetConfig() *config {
	if cachedServices.config == nil {
		cachedServices.config = initConfig()
	}

	return cachedServices.config
}
