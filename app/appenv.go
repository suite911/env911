package app

type AppEnv struct {
	vendor, name, path string
}

func New(args ...interface{}) *AppEnv {
	return new(AppEnv).Init(args...)
}

func (env *AppEnv) Init(args ...interface{}) *AppEnv {
	var pathElems []string
	for _, arg := range args {
		if str, ok := arg.(string); ok {
			if len(str) < 1 {
				continue
			}
			pathElems = append(pathElems, str)
			if len(env.vendor) < 1 {
				env.vendor = env.name
			}
			env.name = str
		}
	}
	env.path = filepath.Join(pathElems)
	env.osInit(args...)
	return env
}

func (env *AppEnv) Name() string {
	return env.name
}

func (env *AppEnv) Path() string {
	return env.path
}

func (env *AppEnv) Vendor() string {
	return env.vendor
}
